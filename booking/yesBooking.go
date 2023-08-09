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

//ling dahao
//var idCustomer = "N2023050721225785a"
//var userNo = "UserNO=144,76,33,219,225,39,126,104,6,62,231,219,144,12,19,188,252,202,32,122,112,205,170,251&UserName=213,226,177,199,50,100,215,91,211,122,154,123,51,253,55,138&Email=229,92,65,77,21,108,20,95,124,140,167,27,36,95,167,80,115,225,218,212,209,182,171,68&UserIdentiNumber=140,92,250,234,198,85,79,14,132,140,202,250,134,122,26,240&Phone=245,54,43,158,39,91,228,116&Mobile=245,54,43,158,39,91,228,116&IdType=60,171,81,93,255,57,213,127&MobileType=195,31,27,161,165,198,213,175&ServiceCookie=184,255,176,225,55,227,53,30,99,215,7,57,7,137,109,218,0,229,97,106,183,125,225,252,156,128,102,132,194,158,1,186,50,48,182,138,52,84,133,26,103,201,153,136,77,239,144,164"
//var cookie = "_fbp=fb.1.1682745972608.1205556132; PCID=16827459728262481361472; __utmz=186092716.1684847279.26.2.utmcsr=tkfile.yes24.com|utmccn=(referral)|utmcmd=referral|utmcct=/; RecentItems=6RrnjbUVjn849e0QBPGlXg2FhviAfX4xsTLjv63z9ItZzlBFMvCGOVWe7rwWVGEwfOYyKfAyLzmaM9JQwAinCnpFnJ93w+0Szd3qV0SO9+AEo6WJ+brLYXS0OE8HtKTg; _abx__ioBYurVi4UKyZbh859fByw={%22firstPartyId%22:%22d30868c4-4bee-42df-90d4-812e52850543%22%2C%22isSendFirstParty%22:true%2C%22lastFirstOpenId%22:%221691237626064:d92d8644-204b-4e5c-bbcd-04a589eac8cd%22%2C%22lastEventLogId%22:%221691237626586:5e96903b-29f5-4ad6-b377-79fa7225a511%22%2C%22lastDailyFirstOpenTime%22:1691237626064%2C%22session%22:{%22sessionId%22:%22b29fb4e3-4faf-452d-a7fb-db8fd23a79c6%22%2C%22lastUpdate%22:1691243026585}%2C%22userId%22:null%2C%22userProperty%22:{%22userProperty%22:[]%2C%22snapshot%22:null}}; RushTime=OK; _gid=GA1.2.152617985.1691405918; _ga_FJT6RQ6VPQ=GS1.1.1691413725.2.0.1691413725.60.0.0; RecentViewGoods=; RecentViewInfo=NotCookie%3DY%26Interval%3D5; _ga=GA1.2.1395600857.1682745973; ASP.NET_SessionId=nxsphntakspgx1jx4qrspmat; __utmc=186092716; YesTicketForeign=UserNO=144,76,33,219,225,39,126,104,6,62,231,219,144,12,19,188,252,202,32,122,112,205,170,251&UserName=213,226,177,199,50,100,215,91,211,122,154,123,51,253,55,138&Email=229,92,65,77,21,108,20,95,124,140,167,27,36,95,167,80,115,225,218,212,209,182,171,68&UserIdentiNumber=140,92,250,234,198,85,79,14,132,140,202,250,134,122,26,240&Phone=245,54,43,158,39,91,228,116&Mobile=245,54,43,158,39,91,228,116&IdType=60,171,81,93,255,57,213,127&MobileType=195,31,27,161,165,198,213,175&ServiceCookie=184,255,176,225,55,227,53,30,99,215,7,57,7,137,109,218,0,229,97,106,183,125,225,252,156,128,102,132,194,158,1,186,50,48,182,138,52,84,133,26,103,201,153,136,77,239,144,164; __utma=186092716.1395600857.1682745973.1691455392.1691460333.59; _ga_719LSSZFC3=GS1.2.1691460333.17.0.1691460333.60.0.0; WaitKey=CE6C5DD8DDC1D2DA9176D4D55791276D65F44B9651316C5C327995F0B1741E11507453E24F4917303A505D119F1D5B07181A51E9E2AD2A4F038487A3D75AF73B76A8AC456210B5B5FD72EE1B4201E52B801AC69A91510CA86C8834AFD3C01A6A191F785C209AEC42A7EC9B9465BEA8BF30"
//ling xiaohao
//var cookie = "__utmz=186092716.1684598271.1.1.utmcsr=bing|utmccn=(organic)|utmcmd=organic|utmctr=(not%20provided); PCID=16845983532992452318367; _ga=GA1.2.454869988.1684598271; _fbp=fb.1.1690815877304.1908226884; ASP.NET_SessionId=k55er2xiqpp0vxspmb1jwctx; __utma=186092716.454869988.1684598271.1691068509.1691306172.7; __utmc=186092716; __utmt=1; _gid=GA1.2.522625886.1691306176; _gat_UA-166644337-1=1; _ga_719LSSZFC3=GS1.2.1691306178.5.0.1691306178.60.0.0; YesTicketForeign=UserNO=110,102,26,136,66,40,120,110,163,226,100,218,234,4,189,224,179,184,98,94,173,28,97,118&UserName=240,182,179,206,18,126,109,145,240,119,125,254,191,231,103,112&Email=151,119,148,171,46,152,30,109,34,43,84,119,228,11,89,115,46,204,50,88,151,16,88,210&UserIdentiNumber=34,252,37,77,253,140,131,5,119,19,107,151,32,170,87,20&Phone=245,54,43,158,39,91,228,116&Mobile=245,54,43,158,39,91,228,116&IdType=60,171,81,93,255,57,213,127&MobileType=195,31,27,161,165,198,213,175&ServiceCookie=240,171,25,176,166,12,251,175,248,61,106,100,150,7,116,71,211,189,53,90,65,169,124,94,160,88,205,189,129,12,133,135,10,50,46,20,56,168,215,161,28,215,10,54,105,52,119,33; __utmb=186092716.5.10.1691306172"
//var idCustomer = "N202305102044370d4"
//var userNo = "UserNO=110,102,26,136,66,40,120,110,163,226,100,218,234,4,189,224,179,184,98,94,173,28,97,118&UserName=240,182,179,206,18,126,109,145,240,119,125,254,191,231,103,112&Email=151,119,148,171,46,152,30,109,34,43,84,119,228,11,89,115,46,204,50,88,151,16,88,210&UserIdentiNumber=34,252,37,77,253,140,131,5,119,19,107,151,32,170,87,20&Phone=245,54,43,158,39,91,228,116&Mobile=245,54,43,158,39,91,228,116&IdType=60,171,81,93,255,57,213,127&MobileType=195,31,27,161,165,198,213,175&ServiceCookie=42,188,239,213,123,124,229,110,132,83,97,115,6,141,181,147,184,4,233,206,43,136,157,144,40,140,4,122,54,46,241,86,215,42,118,128,11,247,100,124,135,23,204,210,226,29,59,247"
var cookie = "_fbp=fb.1.1682745972608.1205556132; PCID=16827459728262481361472; __utmz=186092716.1684847279.26.2.utmcsr=tkfile.yes24.com|utmccn=(referral)|utmcmd=referral|utmcct=/; RecentItems=6RrnjbUVjn849e0QBPGlXg2FhviAfX4xsTLjv63z9ItZzlBFMvCGOVWe7rwWVGEwfOYyKfAyLzmaM9JQwAinCnpFnJ93w+0Szd3qV0SO9+AEo6WJ+brLYXS0OE8HtKTg; _abx__ioBYurVi4UKyZbh859fByw={%22firstPartyId%22:%22d30868c4-4bee-42df-90d4-812e52850543%22%2C%22isSendFirstParty%22:true%2C%22lastFirstOpenId%22:%221691237626064:d92d8644-204b-4e5c-bbcd-04a589eac8cd%22%2C%22lastEventLogId%22:%221691237626586:5e96903b-29f5-4ad6-b377-79fa7225a511%22%2C%22lastDailyFirstOpenTime%22:1691237626064%2C%22session%22:{%22sessionId%22:%22b29fb4e3-4faf-452d-a7fb-db8fd23a79c6%22%2C%22lastUpdate%22:1691243026585}%2C%22userId%22:null%2C%22userProperty%22:{%22userProperty%22:[]%2C%22snapshot%22:null}}; _gid=GA1.2.152617985.1691405918; _ga_FJT6RQ6VPQ=GS1.1.1691413725.2.0.1691413725.60.0.0; _ga=GA1.2.1395600857.1682745973; ASP.NET_SessionId=0z32dudwkbydno0ertgod4ny; __utmc=186092716; _ga_719LSSZFC3=GS1.2.1691585077.20.0.1691585077.60.0.0; __utma=186092716.1395600857.1682745973.1691585075.1691591502.64; __utmt=1; YesTicketForeign=UserNO=123,245,187,233,124,52,197,44,44,34,206,75,199,174,246,101,30,218,25,194,199,240,137,62&UserName=153,44,129,35,117,173,14,216&Email=142,187,16,35,164,82,104,215,248,190,111,249,91,104,173,194,16,91,3,66,150,46,71,64&UserIdentiNumber=242,241,53,77,173,152,69,131,119,49,185,94,56,31,251,120&Phone=245,54,43,158,39,91,228,116&Mobile=245,54,43,158,39,91,228,116&IdType=60,171,81,93,255,57,213,127&MobileType=195,31,27,161,165,198,213,175&ServiceCookie=110,29,150,139,21,59,38,103,180,87,197,93,127,119,204,181,250,34,185,126,57,77,198,111,115,25,196,99,221,35,180,51,165,97,53,229,92,77,202,0,54,191,221,28,182,63,146,149; __utmb=186092716.9.10.1691591502; WaitKey=0DB7D4ABFC5C154054B6A99785877BD15D85A3F3B7167F66587AC79A07FECFFEBA37AA7DD46E0944F2EFEBC6F3923393FE375591DA8BFD871FF31458AAF1EB2B02504C16B6E885373061F8E61CBB561AC7EA17F17FD0E48C6F60005E029713C7B1A4DEBBAEDCC320D7FAFEA74316F7502C312C302C30; NetFunnel_ID=5003%3A201%3Akey%3D0DB7D4ABFC5C154054B6A99785877BD15D85A3F3B7167F66587AC79A07FECFFE2413B69E9A2F3A8FF76F981E13BDC065FE375591DA8BFD871FF31458AAF1EB2B02504C16B6E885373061F8E61CBB561AC7EA17F17FD0E48C6F60005E029713C7B1A4DEBBAEDCC320D7FAFEA74316F7502C312C302C30%26nwait%3D0%26nnext%3D0%26tps%3D0.000000%26ttl%3D10%26ip%3Dtkwait.yes24.com%26port%3D443"
var idCustomer = "N20230429143257823"
var userNo = "UserNO=123,245,187,233,124,52,197,44,44,34,206,75,199,174,246,101,30,218,25,194,199,240,137,62&UserName=153,44,129,35,117,173,14,216&Email=142,187,16,35,164,82,104,215,248,190,111,249,91,104,173,194,16,91,3,66,150,46,71,64&UserIdentiNumber=242,241,53,77,173,152,69,131,119,49,185,94,56,31,251,120&Phone=245,54,43,158,39,91,228,116&Mobile=245,54,43,158,39,91,228,116&IdType=60,171,81,93,255,57,213,127&MobileType=195,31,27,161,165,198,213,175&ServiceCookie=110,29,150,139,21,59,38,103,180,87,197,93,127,119,204,181,250,34,185,126,57,77,198,111,115,25,196,99,221,35,180,51,165,97,53,229,92,77,202,0,54,191,221,28,182,63,146,149"

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
	//logging.Info("开始查询区域是否有票")
	respByte, _ := httpQuery("application/x-www-form-urlencoded", "http://ticket.yes24.com/OSIF/Book.asmx/GetBookWholeFN", body)
	b := BookWhole{}
	err := xml.Unmarshal(respByte, &b)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(b.BlockSeat)
	//defer resp.Body.Close()
	var pidSeats []string
	t := strings.Split(b.BlockSeat.Text, "^")
	for k, _ := range t {
		if len(strings.Split(t[k], "@")) == 8 {
			//seat := strings.Split(v, "@")[6]
			pidSeats = append(pidSeats, strings.Split(t[k], "@")[0])

			class = strings.Split(t[k], "@")[5]
			//fmt.Println("座位：" + seat + "可选")
		}

	}
	sort.Strings(pidSeats)
	for k, _ := range pidSeats {
		code, message, err := YesQueryLock(idTime, pidSeats[k], block)
		if err != nil {
			logging.Error(err)
			continue
		}
		if code != "None" || message != "요청하신 작업이 정상적으로 처리 되었습니다" {
			fmt.Println(block)
			fmt.Println(pidSeats[k], message)
			continue
		}
		pIdSeat = pidSeats[k]
		break
	}
	if len(pIdSeat) > 0 {
		logging.Info(pIdSeat)
	}

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
	orderData.EtcFee = "69#3000^"
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
	orderData.EmergencyName = "z z"
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
	proxy := "http://127.0.0.1:1439"
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
