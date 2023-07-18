package main

import (
	"autobook/booking"

	"autobook/logging"

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

func bookYes24(day string, idPerf string) booking.PaypalUrl {
	var payData booking.PaypalUrl
	times, err := booking.YesFnPerfTime(day, idPerf)
	if err != nil {
		logging.Error(err)
	}
	for {
		blocks, err := booking.YesSeatMap(times.IdTime, times.IdHall)
		if err != nil {
			continue
		}
		if len(blocks) > 0 {
			pidSeat, class := booking.YesQuerySeat(times.IdTime, times.IdHall, blocks[0])

			code, message, err := booking.YesQueryLock(times.IdTime, pidSeat)
			if err != nil {
				logging.Error(err)
				continue
			}
			if code != "None" || message != "요청하신 작업이 정상적으로 처리 되었습니다" {
				continue
			}
			pSeat, price, err := booking.YesQuerySeatFlashEnd(times.IdTime, class)
			if err != nil {
				logging.Error(err)
				continue
			}
			fee, err := booking.YesFnEtcFree(times.IdTime)
			if err != nil {
				logging.Error(err)
				continue
			}
			amountFee, _ := strconv.Atoi(fee)
			amountPrice, _ := strconv.Atoi(price)
			amount := amountFee + amountPrice

			payData, err = booking.YesGetCart(idPerf, pidSeat, times.IdTime, pSeat, amount)
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
