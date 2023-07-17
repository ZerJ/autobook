package main

import (
	"autobook/logging"
	"autobook/query"
	"fmt"
	"strconv"
	"time"
)

func main() {
	//yes24 订票
	url := bookYes24("20230722", "46237")
	fmt.Println(url)
	//付款回执
	//query.PaypalPayResponse(url.EncryptCartID, url.EncryptPaypalOrderID, url.EncryptPaypalOrderID)
}

func bookYes24(day string, idPerf string) query.PaypalUrl {
	var payData query.PaypalUrl
	times, err := query.FnPerfTime(day, idPerf)
	if err != nil {
		logging.Error(err)
	}
	for {
		blocks, err := query.HttpSeatMap(times.IdTime, times.IdHall)
		if err != nil {
			continue
		}
		if len(blocks) > 0 {
			pidSeat, class := query.HttpQuerySeat(times.IdTime, times.IdHall, blocks[0])

			code, message, err := query.HttpQueryLock(times.IdTime, pidSeat)
			if err != nil {
				logging.Error(err)
				continue
			}
			if code != "None" || message != "요청하신 작업이 정상적으로 처리 되었습니다" {
				continue
			}
			pSeat, price, err := query.HttpQuerySeatFlashEnd(times.IdTime, class)
			if err != nil {
				logging.Error(err)
				continue
			}
			fee, err := query.FnEtcFree(times.IdTime)
			if err != nil {
				logging.Error(err)
				continue
			}
			amountFee, _ := strconv.Atoi(fee)
			amountPrice, _ := strconv.Atoi(price)
			amount := amountFee + amountPrice
			payData, err = query.HttpGetCart(idPerf, pidSeat, times.IdTime, pSeat, amount)
			if err != nil {
				logging.Error(err)
				continue
			}
			if len(payData.PaymentRedirectUrl) != 0 {
				break
			}

		}
		time.Sleep(time.Duration(10) * time.Second)
	}
	return payData
}
