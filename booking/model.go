/**
* @Description 功能介绍
* @Author zhengjili
* @Date  2023/7/14  17:26
**/

package booking

import "encoding/xml"

// BookWhole was generated 2023-07-12 10:56:33 by www on server.
type BookWhole struct {
	XMLName xml.Name `xml:"BookWhole"`
	Text    string   `xml:",chardata"`
	Xsi     string   `xml:"xsi,attr"`
	Xsd     string   `xml:"xsd,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	IdTime  struct {
		Text string `xml:",chardata"`
	} `xml:"IdTime"`
	Block struct {
		Text string `xml:",chardata"`
	} `xml:"Block"`
	Layout struct {
		Text string `xml:",chardata"`
	} `xml:"Layout"`
	Background struct {
		Text string `xml:",chardata"`
	} `xml:"Background"`
	BlockSeat struct {
		Text string `xml:",chardata"`
	} `xml:"BlockSeat"`
	BlockRemain struct {
		Text string `xml:",chardata"`
	} `xml:"BlockRemain"`
	Regend struct {
		Text string `xml:",chardata"`
	} `xml:"Regend"`
	Section struct {
		Text string `xml:",chardata"`
	} `xml:"Section"`
}

// ServiceResponse was generated 2023-07-12 15:17:07 by www on server.
type ServiceResponse struct {
	XMLName xml.Name `xml:"ServiceResponse"`
	Text    string   `xml:",chardata"`
	Xsi     string   `xml:"xsi,attr"`
	Xsd     string   `xml:"xsd,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	Code    struct {
		Text string `xml:",chardata"`
	} `xml:"Code"`
	Message struct {
		Text string `xml:",chardata"`
	} `xml:"Message"`
}
type JSONData struct {
	OrderData OrderData `json:"orderData"`
}
type OrderData struct {
	IDPerf             string `json:"IdPerf"`
	IsMPoint           string `json:"IsMPoint"`
	IsHcardOpt         string `json:"IsHcardOpt"`
	IDTime             string `json:"IdTime"`
	TimeOption         string `json:"TimeOption"`
	IDSeat             string `json:"IdSeat"`
	IDNonSeat          string `json:"IdNonSeat"`
	SeatCnt            int    `json:"SeatCnt"`
	Discounts          string `json:"Discounts"`
	Coupons            string `json:"Coupons"`
	YesMoney           string `json:"YesMoney"`
	YesDeposit         string `json:"YesDeposit"`
	YesGift            string `json:"YesGift"`
	YesGiftAmt         int    `json:"YesGiftAmt"`
	GiftTicket         string `json:"GiftTicket"`
	GiftTicketAmt      int    `json:"GiftTicketAmt"`
	BenepiaPoint       int    `json:"BenepiaPoint"`
	BenepiaPointInfo   string `json:"BenepiaPointInfo"`
	CaptchaText        string `json:"CaptchaText"`
	CaptchaKey         string `json:"CaptchaKey"`
	SMFanclubNo        string `json:"SMFanclubNo"`
	EtcFee             string `json:"EtcFee"`
	EtcValidTicketCnt  string `json:"EtcValidTicketCnt"`
	Delivery           string `json:"Delivery"`
	Receiver           string `json:"Receiver"`
	Mobile             string `json:"Mobile"`
	ZipCode            string `json:"ZipCode"`
	Address            string `json:"Address"`
	PaypalCountryCode  string `json:"Paypal_CountryCode"`
	PaypalAdminArea1   string `json:"Paypal_AdminArea1"`
	PaypalAdminArea2   string `json:"Paypal_AdminArea2"`
	PaypalAddressLine1 string `json:"Paypal_AddressLine1"`
	EmergencyName      string `json:"EmergencyName"`
	EmergencyMobile    string `json:"EmergencyMobile"`
	EmergencyMail      string `json:"EmergencyMail"`
	PayBank            string `json:"PayBank"`
	Receipt            string `json:"Receipt"`
	ReceiptNo          string `json:"ReceiptNo"`
	OKCashbagCardNo    string `json:"OKCashbagCardNo"`
	IsYesPointYN       string `json:"IsYesPointYN"`
	YesTicketForeign   string `json:"YesTicketForeign"`
	PayMethod          string `json:"PayMethod"`
	GoodName           string `json:"GoodName"`
	Amount             int    `json:"Amount"`
	SaleChannel        string `json:"SaleChannel"`
}

// BookSeatMap was generated 2023-07-12 16:05:12 by www on server.
type BookSeatMap struct {
	XMLName xml.Name `xml:"BookSeatMap"`
	Text    string   `xml:",chardata"`
	Xsi     string   `xml:"xsi,attr"`
	Xsd     string   `xml:"xsd,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	IdTime  struct {
		Text string `xml:",chardata"`
	} `xml:"IdTime"`
	SeatMap struct {
		Text string `xml:",chardata"`
	} `xml:"SeatMap"`
	BlockRemain struct {
		Text string `xml:",chardata"`
	} `xml:"BlockRemain"`
	Regend struct {
		Text string `xml:",chardata"`
	} `xml:"Regend"`
	Section struct {
		Text string `xml:",chardata"`
	} `xml:"Section"`
}
type PerTime struct {
	IdTime string `json:"idTime"`
	IdHall string `json:"idHall"`
}

type Paypal struct {
	D PaypalUrl `json:"d"`
}
type PaypalUrl struct {
	Type                 string `json:"__type"`
	Code                 int    `json:"code"`
	Message              string `json:"message"`
	EncryptCartID        string `json:"encryptCartID"`
	EncryptPaypalOrderID string `json:"encryptPaypalOrderID"`
	PaymentRedirectUrl   string `json:"paymentRedirectUrl"`
}
