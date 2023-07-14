/**
* @Description 功能介绍
* @Author zhengjili
* @Date  2023/7/14  17:26
**/

package query

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
	"strings"
)

var httpHeader map[string][]string

func HttpQuerySeat(idTime string, idHall string) (pIdSeat string, class string) {
	params := url.Values{}
	params.Add("idHall", idHall)
	params.Add("idTime", idTime)
	params.Add("block", `0`)
	params.Add("channel", `1`)
	params.Add("idCustomer", `N202305102044370d4`)
	params.Add("idOrg", `1`)
	body := strings.NewReader(params.Encode())

	respByte, _ := httpQuery("application/x-www-form-urlencoded", "http://ticket.yes24.com/OSIF/Book.asmx/GetBookWholeFN", body)
	b := BookWhole{}
	err := xml.Unmarshal(respByte, &b)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(b.BlockSeat)
	//defer resp.Body.Close()
	t := strings.Split(b.BlockSeat.Text, "^")
	for _, v := range t {
		if len(strings.Split(v, "@")) == 8 {
			//seat := strings.Split(v, "@")[6]
			pIdSeat = strings.Split(v, "@")[0]
			class = strings.Split(v, "@")[5]

			//fmt.Println("座位：" + seat + "可选")
		}

	}
	return pIdSeat, class
}

func HttpQueryLock(idTime string, token string) {
	fmt.Println(token)
	params := url.Values{}
	params.Add("name", `N202305102044370d4`)
	params.Add("idTime", idTime)
	params.Add("token", token)
	params.Add("Block", `0`)
	body := strings.NewReader(params.Encode())
	//payloadBytes, err := json.Marshal(data)
	//if err != nil {
	//	// handle err
	//}
	//body := bytes.NewReader(payloadBytes)
	respByte, err := httpQuery("application/x-www-form-urlencoded", "http://ticket.yes24.com/OSIF/Book.asmx/Lock", body)

	b := ServiceResponse{}
	err = xml.Unmarshal(respByte, &b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b.Message)
	//fmt.Println(b.BlockSeat)
	//defer resp.Body.Close()

}

func HttpGetCart(idPerf string, pIdSeat string, idTime string, pSeat string) {
	var jsonData JSONData
	var orderData OrderData
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
	orderData.YesTicketForeign = "UserNO=110,102,26,136,66,40,120,110,163,226,100,218,234,4,189,224,179,184,98,94,173,28,97,118&UserName=240,182,179,206,18,126,109,145,240,119,125,254,191,231,103,112&Email=151,119,148,171,46,152,30,109,34,43,84,119,228,11,89,115,46,204,50,88,151,16,88,210&UserIdentiNumber=34,252,37,77,253,140,131,5,119,19,107,151,32,170,87,20&Phone=245,54,43,158,39,91,228,116&Mobile=245,54,43,158,39,91,228,116&IdType=60,171,81,93,255,57,213,127&MobileType=195,31,27,161,165,198,213,175&ServiceCookie=42,188,239,213,123,124,229,110,132,83,97,115,6,141,181,147,184,4,233,206,43,136,157,144,40,140,4,122,54,46,241,86,215,42,118,128,11,247,100,124,135,23,204,210,226,29,59,247"
	orderData.PayMethod = "100000000000"
	orderData.GoodName = "공연티켓상품"
	orderData.Amount = 79000
	orderData.SaleChannel = "1024"
	jsonData.OrderData = orderData
	//body := strings.NewReader("{orderData:{\"IdPerf\":\"46309\",\"IsMPoint\":\"0\",\"IsHcardOpt\":\"0\",\"IdTime\":\"1233538\",\"TimeOption\":\"2\",\"IdSeat\":\"2400031\",\"IdNonSeat\":\"T86$73$80$188$174-1,\",\"SeatCnt\":1,\"Discounts\":\"\",\"Coupons\":\"\",\"YesMoney\":\"\",\"YesDeposit\":\"\",\"YesGift\":\"\",\"YesGiftAmt\":0,\"GiftTicket\":\"\",\"GiftTicketAmt\":0,\"BenepiaPoint\":0,\"BenepiaPointInfo\":\"\",\"CaptchaText\":\"\",\"CaptchaKey\":\"\",\"SMFanclubNo\":\"\",\"EtcFee\":\"69#2000^\",\"EtcValidTicketCnt\":\"1\",\"Delivery\":\"\",\"Receiver\":\"\",\"Mobile\":\"\",\"ZipCode\":\"\",\"Address\":\"\",\"Paypal_CountryCode\":\"\",\"Paypal_AdminArea1\":\"\",\"Paypal_AdminArea2\":\"\",\"Paypal_AddressLine1\":\"\",\"EmergencyName\":\"HONG BINLIN\",\"EmergencyMobile\":\"--\",\"EmergencyMail\":\"@\",\"PayBank\":\"\",\"Receipt\":\"\",\"ReceiptNo\":\"\",\"OKCashbagCardNo\":\"\",\"IsYesPointYN\":\"N\",\"YesTicketForeign\":\"UserNO=110,102,26,136,66,40,120,110,163,226,100,218,234,4,189,224,179,184,98,94,173,28,97,118&UserName=240,182,179,206,18,126,109,145,240,119,125,254,191,231,103,112&Email=151,119,148,171,46,152,30,109,34,43,84,119,228,11,89,115,46,204,50,88,151,16,88,210&UserIdentiNumber=34,252,37,77,253,140,131,5,119,19,107,151,32,170,87,20&Phone=245,54,43,158,39,91,228,116&Mobile=245,54,43,158,39,91,228,116&IdType=60,171,81,93,255,57,213,127&MobileType=195,31,27,161,165,198,213,175&ServiceCookie=42,188,239,213,123,124,229,110,132,83,97,115,6,141,181,147,184,4,233,206,43,136,157,144,40,140,4,122,54,46,241,86,215,42,118,128,11,247,100,124,135,23,204,210,226,29,59,247\",\"PayMethod\":\"100000000000\",\"GoodName\":\"공연티켓상품\",\"Amount\":79000,\"SaleChannel\":\"1024\"}}")
	b, err := json.Marshal(jsonData)
	fmt.Println(err)
	buf := bytes.NewBufferString(string(b))
	respByte, _ := httpQuery("application/json; charset=UTF-8", "http://ticket.yes24.com/Pages/PaypalFnPG/MakeCart_TicketSale.aspx/GetCart", buf)

	fmt.Println(string(respByte))
	//fmt.Println(b.BlockSeat)
	//defer resp.Body.Close()

}

func HttpSeatMap(idTime string, idHall string) error {

	params := url.Values{}
	params.Add("idHall", idHall)
	params.Add("idTime", idTime)
	body := strings.NewReader(params.Encode())
	respByte, _ := httpQuery("application/x-www-form-urlencoded", "http://ticket.yes24.com/OSIF/Book.asmx/GetHallMapRemainFN", body)
	b := BookSeatMap{}
	err := xml.Unmarshal(respByte, &b)
	if err != nil {
		fmt.Println(err)
	}
	logging.Info(string(respByte))
	return nil
}

// HttpQuerySeatFlashEnd pCntClass 为座位号跟着韩文
func HttpQuerySeatFlashEnd(pIdTime string, pCntClass string) (pSeat string) {
	params := url.Values{}
	params.Add("pIdTime", pIdTime)
	params.Add("PCntClass", pCntClass)
	body := strings.NewReader(params.Encode())
	//body := strings.NewReader("pIdTime=1234289&PCntClass=\"2600035@전석\"")
	respByte, _ := httpQuery("application/x-www-form-urlencoded; charset=UTF-8", "http://ticket.yes24.com/Pages/English/Sale/Ajax/Perf/FnTimeSeatFlashEnd.aspx", body)

	doc := soup.HTMLParse(string(respByte))
	root := doc.Find("div")

	pSeat = root.Attrs()["classbyte"]
	return pSeat + "-1,"
}

func FnPerfTime(pDay string, pIdPerf string) PerTime {
	params := url.Values{}
	params.Add("pDay", pDay)
	params.Add("pIdPerf", pIdPerf)
	params.Add("pIdCode", "")
	params.Add("pIsMania", "0")
	body := strings.NewReader(params.Encode())
	respByte, _ := httpQuery("application/x-www-form-urlencoded; charset=UTF-8", "http://ticket.yes24.com/Pages/English/Sale/Ajax/Perf/FnPerfTime.aspx", body)
	doc := soup.HTMLParse(string(respByte))
	root := doc.Find("li")
	var vs PerTime
	vs.IdHall = root.Attrs()["idhall"]
	vs.PidTime = root.Attrs()["value"]
	return vs
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
	req.Header.Set("Cookie", "_gcl_au=1.1.1231468339.1682498612; PCID=16824986121042670196551; _fbp=fb.1.1682498612158.2133767039; _ga=GA1.2.950941543.1684314475; __utmz=186092716.1684314477.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); __utmz=12748607.1684742639.7.6.utmcsr=ticket.yes24.com|utmccn=(referral)|utmcmd=referral|utmcct=/; _ga_D2E1MGEBZ0=GS1.2.1688973343.1.0.1688973343.0.0.0; _ga_FJT6RQ6VPQ=GS1.1.1688973343.1.0.1688973347.56.0.0; _ga_719LSSZFC3=GS1.2.1688973350.2.0.1688973350.60.0.0; ASP.NET_SessionId=t2bzhjhmjnikupackpajsqfs; __utmc=186092716; YesTicketForeign=UserNO=110,102,26,136,66,40,120,110,163,226,100,218,234,4,189,224,179,184,98,94,173,28,97,118&UserName=240,182,179,206,18,126,109,145,240,119,125,254,191,231,103,112&Email=151,119,148,171,46,152,30,109,34,43,84,119,228,11,89,115,46,204,50,88,151,16,88,210&UserIdentiNumber=34,252,37,77,253,140,131,5,119,19,107,151,32,170,87,20&Phone=245,54,43,158,39,91,228,116&Mobile=245,54,43,158,39,91,228,116&IdType=60,171,81,93,255,57,213,127&MobileType=195,31,27,161,165,198,213,175&ServiceCookie=42,188,239,213,123,124,229,110,132,83,97,115,6,141,181,147,184,4,233,206,43,136,157,144,40,140,4,122,54,46,241,86,215,42,118,128,11,247,100,124,135,23,204,210,226,29,59,247; __utma=186092716.950941543.1684314475.1689130890.1689143721.12")
	req.Header.Set("Origin", "http://ticket.yes24.com")
	req.Header.Set("Proxy-Connection", "keep-alive")
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