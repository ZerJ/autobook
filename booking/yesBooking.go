/**
* @Description 功能介绍
* @Author zhengjili
* @Date  2023/7/14  17:26
**/

package booking

import (
	"autobook/logging"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/anaskhan96/soup"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var cookie = "_gcl_au=1.1.1231468339.1682498612; PCID=16824986121042670196551; _fbp=fb.1.1682498612158.2133767039; _ga=GA1.2.950941543.1684314475; __utmz=186092716.1684314477.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); __utmz=12748607.1684742639.7.6.utmcsr=ticket.yes24.com|utmccn=(referral)|utmcmd=referral|utmcct=/; _ga_D2E1MGEBZ0=GS1.2.1688973343.1.0.1688973343.0.0.0; _ga_FJT6RQ6VPQ=GS1.1.1688973343.1.0.1688973347.56.0.0; _ga_719LSSZFC3=GS1.2.1688973350.2.0.1688973350.60.0.0; ASP.NET_SessionId=qso151itzfpxluze1pugy34c; __utma=186092716.950941543.1684314475.1689328288.1689557033.23; __utmc=186092716; YesTicketForeign=UserNO=110,102,26,136,66,40,120,110,163,226,100,218,234,4,189,224,179,184,98,94,173,28,97,118&UserName=240,182,179,206,18,126,109,145,240,119,125,254,191,231,103,112&Email=151,119,148,171,46,152,30,109,34,43,84,119,228,11,89,115,46,204,50,88,151,16,88,210&UserIdentiNumber=34,252,37,77,253,140,131,5,119,19,107,151,32,170,87,20&Phone=245,54,43,158,39,91,228,116&Mobile=245,54,43,158,39,91,228,116&IdType=60,171,81,93,255,57,213,127&MobileType=195,31,27,161,165,198,213,175&ServiceCookie=19,241,128,239,156,57,68,1,5,233,199,48,224,104,94,18,213,148,190,111,214,23,113,158,73,238,255,156,69,181,213,147,84,37,178,209,147,174,192,74,106,170,8,161,150,128,13,40; HTTP_REFERER=http://ticket.yes24.com/; RecentViewGoods=; RecentViewInfo=NotCookie%3DY%26Interval%3D5; __utmt=1; __utmb=186092716.28.10.1689557033; WaitKey=D99E66092D0F5997A926BA53D895ECD7763C590C43AD0881FF848C5A7871888EC6A53C9157644DB6F28046AC6D5A822DCE765AEEA99B4B29A86A1ED1D7312B9082022A2343795F167EE071C5AFBF38F875321161929FA87411C5FCAE33BF853600390DD2AC19D3F88653708CDEFE56EF302C30; NetFunnel_ID=5003%3A201%3Akey%3DD99E66092D0F5997A926BA53D895ECD7763C590C43AD0881FF848C5A7871888E132D9FBA7B7802EB7284171CB6164440CE765AEEA99B4B29A86A1ED1D7312B9082022A2343795F167EE071C5AFBF38F875321161929FA87411C5FCAE33BF85363ABC0CF91CAD74559E4AAEBE23F15028302C30%26nwait%3D0%26nnext%3D0%26tps%3D0.000000%26ttl%3D10%26ip%3Dtkwait.yes24.com%26port%3D443"
var idCustomer = "N202305102044370d4"
var userNo = "UserNO=110,102,26,136,66,40,120,110,163,226,100,218,234,4,189,224,179,184,98,94,173,28,97,118&UserName=240,182,179,206,18,126,109,145,240,119,125,254,191,231,103,112&Email=151,119,148,171,46,152,30,109,34,43,84,119,228,11,89,115,46,204,50,88,151,16,88,210&UserIdentiNumber=34,252,37,77,253,140,131,5,119,19,107,151,32,170,87,20&Phone=245,54,43,158,39,91,228,116&Mobile=245,54,43,158,39,91,228,116&IdType=60,171,81,93,255,57,213,127&MobileType=195,31,27,161,165,198,213,175&ServiceCookie=42,188,239,213,123,124,229,110,132,83,97,115,6,141,181,147,184,4,233,206,43,136,157,144,40,140,4,122,54,46,241,86,215,42,118,128,11,247,100,124,135,23,204,210,226,29,59,247"

// YesQuerySeat 查询区域是否有票
func YesQuerySeat(idTime string, idHall string, block string) (pIdSeat string, class string) {
	if block == "" {
		block = "0"
	}
	params := url.Values{}
	params.Add("idHall", idHall)
	params.Add("idTime", idTime)
	params.Add("block", block)
	params.Add("channel", `1`)
	params.Add("idCustomer", idCustomer)
	params.Add("idOrg", `1`)
	body := strings.NewReader(params.Encode())
	logging.Info("开始查询区域是否有票")
	respByte, _ := httpQuery("application/x-www-form-urlencoded", "http://ticket.yes24.com/OSIF/Book.asmx/GetBookWholeFN", body)
	b := BookWhole{}
	err := xml.Unmarshal(respByte, &b)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(b.BlockSeat)
	//defer resp.Body.Close()
	minSeat := 99999999
	t := strings.Split(b.BlockSeat.Text, "^")
	for k, _ := range t {
		if len(strings.Split(t[k], "@")) == 8 {
			//seat := strings.Split(v, "@")[6]
			seat, _ := strconv.Atoi(strings.Split(t[k], "@")[0])
			if seat < minSeat {
				minSeat = seat
				logging.Info(seat)
			}
			class = strings.Split(t[k], "@")[5]
			//fmt.Println("座位：" + seat + "可选")
		}

	}
	if minSeat != 99999999 {
		pIdSeat = strconv.Itoa(minSeat)
	}

	logging.Info(pIdSeat)
	return pIdSeat, class
}

// YesQueryLock 锁票
func YesQueryLock(idTime string, token string, block string) (code string, message string, err error) {
	params := url.Values{}
	params.Add("name", idCustomer)
	params.Add("idTime", idTime)
	params.Add("token", token)
	params.Add("Block", block)
	body := strings.NewReader(params.Encode())
	//payloadBytes, err := json.Marshal(data)
	//if err != nil {
	//	// handle err
	//}
	//body := bytes.NewReader(payloadBytes)
	logging.Info("尝试锁票")
	respByte, err := httpQuery("application/x-www-form-urlencoded", "http://ticket.yes24.com/OSIF/Book.asmx/Lock", body)
	if err != nil {
		return
	}
	b := ServiceResponse{}
	err = xml.Unmarshal(respByte, &b)
	if err != nil {
		return
	}
	return b.Code.Text, b.Message.Text, err
	//fmt.Println(b.BlockSeat)
	//defer resp.Body.Close()

}

// YesGetCart 获取付款信息
func YesGetCart(idPerf string, pIdSeat string, idTime string, pSeat string, amount int) (r PaypalUrl, err error) {
	var jsonData JSONData
	var orderData OrderData
	var d Paypal
	orderData.IDPerf = idPerf
	orderData.IsMPoint = "0"
	orderData.IsHcardOpt = "0"
	orderData.IDTime = idTime
	orderData.TimeOption = "2"
	orderData.IDSeat = pIdSeat
	orderData.IDNonSeat = pSeat
	orderData.SeatCnt = 1
	orderData.Discounts = ""
	orderData.Coupons = ""
	orderData.YesMoney = ""
	orderData.YesDeposit = ""
	orderData.YesGift = ""
	orderData.YesGiftAmt = 0
	orderData.GiftTicket = ""
	orderData.GiftTicketAmt = 0
	orderData.BenepiaPoint = 0
	orderData.BenepiaPointInfo = ""
	orderData.CaptchaText = ""
	orderData.CaptchaKey = ""
	orderData.SMFanclubNo = ""
	orderData.EtcFee = "69#2000^"
	orderData.EtcValidTicketCnt = "1"
	orderData.Delivery = ""
	orderData.Receiver = ""
	orderData.Mobile = ""
	orderData.ZipCode = ""
	orderData.Address = ""
	orderData.PaypalCountryCode = ""
	orderData.PaypalAdminArea1 = ""
	orderData.PaypalAdminArea2 = ""
	orderData.PaypalAddressLine1 = ""
	orderData.EmergencyName = "HONG BINLIN"
	orderData.EmergencyMobile = "--"
	orderData.EmergencyMail = "@"
	orderData.PayBank = ""
	orderData.Receipt = ""
	orderData.ReceiptNo = ""
	orderData.OKCashbagCardNo = ""
	orderData.IsYesPointYN = "N"
	orderData.YesTicketForeign = userNo
	orderData.PayMethod = "100000000000"
	orderData.GoodName = "공연티켓상품"
	orderData.Amount = amount
	orderData.SaleChannel = "1024"
	jsonData.OrderData = orderData
	//body := strings.NewReader("{orderData:{\"IdPerf\":\"46309\",\"IsMPoint\":\"0\",\"IsHcardOpt\":\"0\",\"IdTime\":\"1233538\",\"TimeOption\":\"2\",\"IdSeat\":\"2400031\",\"IdNonSeat\":\"T86$73$80$188$174-1,\",\"SeatCnt\":1,\"Discounts\":\"\",\"Coupons\":\"\",\"YesMoney\":\"\",\"YesDeposit\":\"\",\"YesGift\":\"\",\"YesGiftAmt\":0,\"GiftTicket\":\"\",\"GiftTicketAmt\":0,\"BenepiaPoint\":0,\"BenepiaPointInfo\":\"\",\"CaptchaText\":\"\",\"CaptchaKey\":\"\",\"SMFanclubNo\":\"\",\"EtcFee\":\"69#2000^\",\"EtcValidTicketCnt\":\"1\",\"Delivery\":\"\",\"Receiver\":\"\",\"Mobile\":\"\",\"ZipCode\":\"\",\"Address\":\"\",\"Paypal_CountryCode\":\"\",\"Paypal_AdminArea1\":\"\",\"Paypal_AdminArea2\":\"\",\"Paypal_AddressLine1\":\"\",\"EmergencyName\":\"HONG BINLIN\",\"EmergencyMobile\":\"--\",\"EmergencyMail\":\"@\",\"PayBank\":\"\",\"Receipt\":\"\",\"ReceiptNo\":\"\",\"OKCashbagCardNo\":\"\",\"IsYesPointYN\":\"N\",\"YesTicketForeign\":\"UserNO=110,102,26,136,66,40,120,110,163,226,100,218,234,4,189,224,179,184,98,94,173,28,97,118&UserName=240,182,179,206,18,126,109,145,240,119,125,254,191,231,103,112&Email=151,119,148,171,46,152,30,109,34,43,84,119,228,11,89,115,46,204,50,88,151,16,88,210&UserIdentiNumber=34,252,37,77,253,140,131,5,119,19,107,151,32,170,87,20&Phone=245,54,43,158,39,91,228,116&Mobile=245,54,43,158,39,91,228,116&IdType=60,171,81,93,255,57,213,127&MobileType=195,31,27,161,165,198,213,175&ServiceCookie=42,188,239,213,123,124,229,110,132,83,97,115,6,141,181,147,184,4,233,206,43,136,157,144,40,140,4,122,54,46,241,86,215,42,118,128,11,247,100,124,135,23,204,210,226,29,59,247\",\"PayMethod\":\"100000000000\",\"GoodName\":\"공연티켓상품\",\"Amount\":79000,\"SaleChannel\":\"1024\"}}")
	b, err := json.Marshal(jsonData)
	if err != nil {
		return
	}
	buf := bytes.NewBufferString(string(b))
	logging.Info("获取付款信息")
	respByte, _ := httpQuery("application/json; charset=UTF-8", "http://ticket.yes24.com/Pages/PaypalFnPG/MakeCart_TicketSale.aspx/GetCart", buf)

	err = json.Unmarshal(respByte, &d)
	if err != nil {
		return
	}
	logging.Info("获取付款信息成功：" + string(respByte))
	return d.D, err
	//fmt.Println(b.BlockSeat)
	//defer resp.Body.Close()

}

// YesSeatMap 查询全图是否有票
func YesSeatMap(idTime string, idHall string) (blocks []string, err error) {

	var bookSeatMap BookSeatMap
	params := url.Values{}
	params.Add("idHall", idHall)
	params.Add("idTime", idTime)
	body := strings.NewReader(params.Encode())
	logging.Info("开始查询全图是否有票")
	respByte, err := httpQuery("application/x-www-form-urlencoded", "http://ticket.yes24.com/OSIF/Book.asmx/GetHallMapRemainFN", body)

	err = xml.Unmarshal(respByte, &bookSeatMap)
	if err != nil {
		fmt.Println(err)
	}
	if len(bookSeatMap.BlockRemain.Text) > 0 {
		blockStr := strings.Split(bookSeatMap.BlockRemain.Text, "^")
		for _, v := range blockStr {
			if strings.Contains(v, "@") {
				fmt.Println(v)
				if strings.Split(v, "@")[1] != "0" {
					blocks = append(blocks, strings.Split(v, "@")[0])
				}
			}

		}
	}
	if len(blocks) == 0 {
		logging.Info("无票")
	}
	return blocks, nil
}

// GetBlockInfo 查询全图是否有票
func GetBlockInfo(idTime string, idHall string) (blockInfo []BlockInfo, err error) {

	var bookSeatMap BookSeatMap
	params := url.Values{}
	params.Add("idHall", idHall)
	params.Add("idTime", idTime)
	body := strings.NewReader(params.Encode())
	logging.Info("开始查询全图是否有票")
	respByte, err := httpQuery("application/x-www-form-urlencoded", "http://ticket.yes24.com/OSIF/Book.asmx/GetHallMapRemainFN", body)

	err = xml.Unmarshal(respByte, &bookSeatMap)
	if err != nil {
		fmt.Println(err)
	}
	if len(bookSeatMap.Section.Text) > 0 {
		blockStr := strings.Split(bookSeatMap.Section.Text, "^")
		for k, v := range blockStr {
			if strings.Contains(v, "@") {
				var info BlockInfo
				info.Block = strings.Split(blockStr[k], "@")[1]
				info.Class = strings.Split(blockStr[k], "@")[0]
				info.Seat = strings.Split(blockStr[k], "@")[2]
				blockInfo = append(blockInfo, info)
			}

		}
	}
	return blockInfo, nil
}

// YesQuerySeatFlashEnd 获取pCntClass 为座位号跟着韩文
func YesQuerySeatFlashEnd(pIdTime string, pCntClass string) (pSeat string, price string, err error) {
	params := url.Values{}
	params.Add("pIdTime", pIdTime)
	params.Add("PCntClass", pCntClass)
	body := strings.NewReader(params.Encode())
	//body := strings.NewReader("pIdTime=1234289&PCntClass=\"2600035@전석\"")
	logging.Info("获取价格，座位代码")
	respByte, err := httpQuery("application/x-www-form-urlencoded; charset=UTF-8", "http://ticket.yes24.com/Pages/English/Sale/Ajax/Perf/FnTimeSeatFlashEnd.aspx", body)
	if err != nil {
		return
	}
	doc := soup.HTMLParse(string(respByte))
	root := doc.Find("div")
	if strings.Contains(string(respByte), "classbyte") && strings.Contains(string(respByte), "price") {
		logging.Info(root.Attrs()["classbyte"]+"-1,", doc.Find("select").Attrs()["price"])
		return root.Attrs()["classbyte"] + "-1,", doc.Find("select").Attrs()["price"], nil
	} else {
		fmt.Println(string(respByte))
		fmt.Println(pIdTime)
		return "", "", fmt.Errorf("发生一些错误")
	}

}

// YesFnPerfTime 获取时间id 场馆id
func YesFnPerfTime(pDay string, pIdPerf string) (PerTime, error) {
	params := url.Values{}
	params.Add("pDay", pDay)
	params.Add("pIdPerf", pIdPerf)
	params.Add("pIdCode", "")
	params.Add("pIsMania", "0")
	body := strings.NewReader(params.Encode())
	logging.Info(" 获取时间场馆")
	respByte, err := httpQuery("application/x-www-form-urlencoded; charset=UTF-8", "http://ticket.yes24.com/Pages/English/Sale/Ajax/Perf/FnPerfTime.aspx", body)
	if err != nil {
		return PerTime{}, err
	}
	doc := soup.HTMLParse(string(respByte))
	root := doc.Find("li")
	var vs PerTime
	if strings.Contains(string(respByte), "idhall") && strings.Contains(string(respByte), "value") {
		vs.IdHall = root.Attrs()["idhall"]
		vs.IdTime = root.Attrs()["value"]
	} else {
		fmt.Println(string(respByte))
		return vs, fmt.Errorf("发生一些错误")
	}

	logging.Info(" 获取时间场馆成功 idHall:" + vs.IdHall + ",idTime:" + vs.IdTime)
	return vs, err
}

// YesFnEtcFree 获取手续费fee
func YesFnEtcFree(pIdTime string) (fee string, err error) {

	params := url.Values{}
	params.Add("pIdTime", pIdTime)
	params.Add("pSeatCnt", "1")
	params.Add("pFreeCountOfCoupon", "0")
	params.Add("pFreeCountOfGiftTicket", "0")
	body := strings.NewReader(params.Encode())
	//body := strings.NewReader("pIdTime=1234289&PCntClass=\"2600035@전석\"")
	logging.Info(" 获取手续费fee")
	respByte, err := httpQuery("application/x-www-form-urlencoded; charset=UTF-8", "http://ticket.yes24.com/Pages/English/Sale/Ajax/Perf/FnEtcFee.aspx", body)
	if err != nil {
		return
	}
	doc := soup.HTMLParse(string(respByte))
	root := doc.Find("input", "id", "EtcFeeAmount")
	logging.Info(" 获取手续费fee成功:" + fee)
	return root.Attrs()["value"], err
}

func httpQuery(contentType string, urls string, body io.Reader) (respByte []byte, err error) {
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		fmt.Println(err)
	}
	proxy := "http://127.0.0.1:7890"
	proxyAddr, _ := url.Parse(proxy)
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyAddr),
		},
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6,ko;q=0.5")
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Proxy-Connection", "keep-alive")
	req.Header.Set("Cookie", "_gcl_au=1.1.1231468339.1682498612; PCID=16824986121042670196551; _fbp=fb.1.1682498612158.2133767039; _ga=GA1.2.950941543.1684314475; __utmz=186092716.1684314477.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); __utmz=12748607.1684742639.7.6.utmcsr=ticket.yes24.com|utmccn=(referral)|utmcmd=referral|utmcct=/; _ga_D2E1MGEBZ0=GS1.2.1688973343.1.0.1688973343.0.0.0; _ga_FJT6RQ6VPQ=GS1.1.1688973343.1.0.1688973347.56.0.0; _ga_719LSSZFC3=GS1.2.1688973350.2.0.1688973350.60.0.0; ASP.NET_SessionId=qso151itzfpxluze1pugy34c; __utma=186092716.950941543.1684314475.1689328288.1689557033.23; __utmc=186092716; YesTicketForeign=UserNO=110,102,26,136,66,40,120,110,163,226,100,218,234,4,189,224,179,184,98,94,173,28,97,118&UserName=240,182,179,206,18,126,109,145,240,119,125,254,191,231,103,112&Email=151,119,148,171,46,152,30,109,34,43,84,119,228,11,89,115,46,204,50,88,151,16,88,210&UserIdentiNumber=34,252,37,77,253,140,131,5,119,19,107,151,32,170,87,20&Phone=245,54,43,158,39,91,228,116&Mobile=245,54,43,158,39,91,228,116&IdType=60,171,81,93,255,57,213,127&MobileType=195,31,27,161,165,198,213,175&ServiceCookie=19,241,128,239,156,57,68,1,5,233,199,48,224,104,94,18,213,148,190,111,214,23,113,158,73,238,255,156,69,181,213,147,84,37,178,209,147,174,192,74,106,170,8,161,150,128,13,40; HTTP_REFERER=http://ticket.yes24.com/; RecentViewGoods=; RecentViewInfo=NotCookie%3DY%26Interval%3D5; __utmt=1; __utmb=186092716.28.10.1689557033; WaitKey=D99E66092D0F5997A926BA53D895ECD7763C590C43AD0881FF848C5A7871888EC6A53C9157644DB6F28046AC6D5A822DCE765AEEA99B4B29A86A1ED1D7312B9082022A2343795F167EE071C5AFBF38F875321161929FA87411C5FCAE33BF853600390DD2AC19D3F88653708CDEFE56EF302C30; NetFunnel_ID=5003%3A201%3Akey%3DD99E66092D0F5997A926BA53D895ECD7763C590C43AD0881FF848C5A7871888E132D9FBA7B7802EB7284171CB6164440CE765AEEA99B4B29A86A1ED1D7312B9082022A2343795F167EE071C5AFBF38F875321161929FA87411C5FCAE33BF85363ABC0CF91CAD74559E4AAEBE23F15028302C30%26nwait%3D0%26nnext%3D0%26tps%3D0.000000%26ttl%3D10%26ip%3Dtkwait.yes24.com%26port%3D443")
	req.Header.Set("Origin", "http://ticket.yes24.com")
	req.Header.Set("Referer", "http://ticket.yes24.com/Pages/English/Sale/FnPerfSaleProcess.aspx?IdPerf=46309")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.67")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	resp, err := client.Do(req)
	if err != nil {

		return
	}
	respByte, err = ioutil.ReadAll(resp.Body)
	return respByte, nil
}
func YesPaypalPayResponse(cartID string, orderId string, token string) {
	params := url.Values{}
	params.Add("paypalCartID", cartID)
	params.Add("paypalOrderID", orderId)
	params.Add("paypalPGToken", token)
	body := strings.NewReader(params.Encode())
	respByte, err := httpQuery("application/x-www-form-urlencoded; charset=UTF-8", "http://ticket.yes24.com/Pages/PaypalFnPG/PaypalPayResponse.aspx", body)
	fmt.Println(string(respByte))
	fmt.Println(err)
}
