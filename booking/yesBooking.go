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
	"sort"
	"strings"
)

// var cookie = "_fbp=fb.1.1682745972608.1205556132; PCID=16827459728262481361472; _ga=GA1.2.1395600857.1682745973; __utmz=186092716.1684847279.26.2.utmcsr=tkfile.yes24.com|utmccn=(referral)|utmcmd=referral|utmcct=/; _ga_719LSSZFC3=GS1.2.1690547349.2.0.1690547349.60.0.0; ASP.NET_SessionId=w3k4jwntznno12ynw5wegv20; __utma=186092716.1395600857.1682745973.1690547345.1690814611.37; __utmc=186092716; __utmt=1; YesTicketForeign=UserNO=144,76,33,219,225,39,126,104,6,62,231,219,144,12,19,188,252,202,32,122,112,205,170,251&UserName=213,226,177,199,50,100,215,91,211,122,154,123,51,253,55,138&Email=229,92,65,77,21,108,20,95,124,140,167,27,36,95,167,80,115,225,218,212,209,182,171,68&UserIdentiNumber=140,92,250,234,198,85,79,14,132,140,202,250,134,122,26,240&Phone=245,54,43,158,39,91,228,116&Mobile=245,54,43,158,39,91,228,116&IdType=60,171,81,93,255,57,213,127&MobileType=195,31,27,161,165,198,213,175&ServiceCookie=227,34,38,52,198,215,195,15,218,240,254,28,93,77,134,135,171,196,7,210,129,58,222,206,244,234,131,209,233,250,90,71,56,237,32,251,212,3,3,27,174,236,132,200,123,217,238,192; __utmb=186092716.4.10.1690814611"
// var idCustomer = "N2023050721225785a"

// var userNo = "UserNO=144,76,33,219,225,39,126,104,6,62,231,219,144,12,19,188,252,202,32,122,112,205,170,251&UserName=213,226,177,199,50,100,215,91,211,122,154,123,51,253,55,138&Email=229,92,65,77,21,108,20,95,124,140,167,27,36,95,167,80,115,225,218,212,209,182,171,68&UserIdentiNumber=140,92,250,234,198,85,79,14,132,140,202,250,134,122,26,240&Phone=245,54,43,158,39,91,228,116&Mobile=245,54,43,158,39,91,228,116&IdType=60,171,81,93,255,57,213,127&MobileType=195,31,27,161,165,198,213,175&ServiceCookie=227,34,38,52,198,215,195,15,218,240,254,28,93,77,134,135,171,196,7,210,129,58,222,206,244,234,131,209,233,250,90,71,56,237,32,251,212,3,3,27,174,236,132,200,123,217,238,192"
var cookie = "PCID=16824986121042670196551; _fbp=fb.1.1682498612158.2133767039; __utmz=186092716.1682500477.1.1.utmcsr=yes24.com|utmccn=(referral)|utmcmd=referral|utmcct=/Main/default.aspx; __utmz=12748607.1684742639.7.6.utmcsr=ticket.yes24.com|utmccn=(referral)|utmcmd=referral|utmcct=/; __utma=12748607.921202228.1682498612.1684742639.1684824539.8; _ga_D2E1MGEBZ0=GS1.2.1688973343.1.0.1688973343.0.0.0; _abx__ioBYurVi4UKyZbh859fByw={%22firstPartyId%22:%224872585e-9334-4c15-8f05-330a8954a5d1%22%2C%22isSendFirstParty%22:true%2C%22lastFirstOpenId%22:%221682498612428:3defdb05-8abc-4da8-b22d-d8c6714b31e5%22%2C%22lastEventLogId%22:%221689662651490:818e5a83-35be-45f4-8569-5d5a102c418f%22%2C%22lastDailyFirstOpenTime%22:1689662634132%2C%22session%22:{%22sessionId%22:%22b7b2af63-cf1e-43e2-a4d0-d7c96271d717%22%2C%22lastUpdate%22:1689662635577}%2C%22userId%22:null%2C%22userProperty%22:{%22userProperty%22:[]%2C%22snapshot%22:%228ad504da-ca26-4224-9a96-2bc941bcf546%22}}; cartCookieCnt=0; _ga_FJT6RQ6VPQ=GS1.1.1689662629.2.1.1689662664.25.0.0; _ga=GA1.2.921202228.1682498612; ASP.NET_SessionId=czpjieprasfrvgvulzlei4gf; _gid=GA1.2.854771177.1691144868; __utma=186092716.921202228.1682498612.1690427536.1691144868.28; __utmc=186092716; YesTicketForeign=UserNO=144,76,33,219,225,39,126,104,6,62,231,219,144,12,19,188,252,202,32,122,112,205,170,251&UserName=213,226,177,199,50,100,215,91,211,122,154,123,51,253,55,138&Email=229,92,65,77,21,108,20,95,124,140,167,27,36,95,167,80,115,225,218,212,209,182,171,68&UserIdentiNumber=140,92,250,234,198,85,79,14,132,140,202,250,134,122,26,240&Phone=245,54,43,158,39,91,228,116&Mobile=245,54,43,158,39,91,228,116&IdType=60,171,81,93,255,57,213,127&MobileType=195,31,27,161,165,198,213,175&ServiceCookie=18,109,33,26,241,108,185,223,65,232,77,207,42,170,21,118,26,19,87,112,177,230,156,24,12,199,233,194,80,190,236,74,6,183,118,39,29,88,159,223,131,163,176,20,21,233,74,83; __utmt=1; _gat_UA-166644337-1=1; _ga_719LSSZFC3=GS1.2.1691144868.5.1.1691146410.60.0.0; __utmb=186092716.20.10.1691144868"

var idCustomer = "N2023050721225785a"
var userNo = "UserNO=144,76,33,219,225,39,126,104,6,62,231,219,144,12,19,188,252,202,32,122,112,205,170,251&UserName=213,226,177,199,50,100,215,91,211,122,154,123,51,253,55,138&Email=229,92,65,77,21,108,20,95,124,140,167,27,36,95,167,80,115,225,218,212,209,182,171,68&UserIdentiNumber=140,92,250,234,198,85,79,14,132,140,202,250,134,122,26,240&Phone=245,54,43,158,39,91,228,116&Mobile=245,54,43,158,39,91,228,116&IdType=60,171,81,93,255,57,213,127&MobileType=195,31,27,161,165,198,213,175&ServiceCookie=18,109,33,26,241,108,185,223,65,232,77,207,42,170,21,118,26,19,87,112,177,230,156,24,12,199,233,194,80,190,236,74,6,183,118,39,29,88,159,223,131,163,176,20,21,233,74,83"

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
	t := strings.Split(b.BlockSeat.Text, "^")
	var pidSeats []string
	for k, _ := range t {
		if len(strings.Split(t[k], "@")) == 8 {
			//seat := strings.Split(v, "@")[6]
			//code, message, err := YesQueryLock(idTime, strings.Split(t[k], "@")[0], block)
			//if err != nil {
			//	logging.Error(err)
			//	continue
			//}
			//if code != "None" || message != "요청하신 작업이 정상적으로 처리 되었습니다" {
			//	fmt.Println(code, message)
			//	continue
			//}
			//seat, _ := strconv.Atoi(strings.Split(t[k], "@")[0])
			pidSeats = append(pidSeats, strings.Split(t[k], "@")[0])

			class = strings.Split(t[k], "@")[5]
			//fmt.Println("座位：" + seat + "可选")
		}

	}
	sort.Strings(pidSeats)
	fmt.Println(pidSeats)
	for k, _ := range pidSeats {
		code, message, err := YesQueryLock(idTime, pidSeats[k], block)
		if err != nil {
			logging.Error(err)
			continue
		}
		if code != "None" || message != "요청하신 작업이 정상적으로 처리 되었습니다" {
			fmt.Println(code, message)
			continue
		}
		pIdSeat = pidSeats[k]
		break
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
				if strings.Split(v, "@")[1] != "0" {
					blocks = append(blocks, strings.Split(v, "@")[0])
				}
			}

		}
	}
	if len(blocks) == 0 {
		logging.Info("无票")
	}
	sort.Strings(blocks)
	logging.Info(blocks)
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
	req.Header.Set("Cookie", cookie)
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
