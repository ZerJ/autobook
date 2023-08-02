package main

import (
	"autobook/booking"
	"strconv"
	"time"

	"autobook/logging"

	"fmt"
)

func main() {
	//yes24 订票
	for {
		//url := bookYes24("20230812", "46558")
		url := bookYes24("20231025", "46305")
		if len(url.EncryptPaypalOrderID) > 0 {
			fmt.Println(url)
			break
		}
		time.Sleep(time.Duration(10) * time.Millisecond)
	}

	//booking.InitRedis()

	//cmd:=booking.Redisclient.Set("encryptCartID",url.EncryptCartID,time.Duration(60) * time.Second)
	//fmt.Println(cmd.Err())
	//c:=booking.Redisclient.Set("encryptPaypalOrderID",url.EncryptPaypalOrderID,time.Duration(60) * time.Second)
	//fmt.Println(c.Err())

	//付款回执
	//encryptCartID,_:=booking.Redisclient.Get("encryptCartID").Result()
	//encryptPaypalOrderID,_:=booking.Redisclient.Get("encryptPaypalOrderID").Result()
	//booking.YesPaypalPayResponse(encryptCartID, encryptPaypalOrderID, encryptPaypalOrderID)
}

func bookYes24(day string, idPerf string) booking.PaypalUrl {
	var payData booking.PaypalUrl
	times, err := booking.YesFnPerfTime(day, idPerf)
	if err != nil {
		logging.Error(err)
	}

	for len(times.IdTime) > 0 {
		blocks, err := booking.YesSeatMap(times.IdTime, times.IdHall)

		if err != nil {
			continue
		}
		if len(blocks) > 0 {
			pidSeat, class := booking.YesQuerySeat(times.IdTime, times.IdHall, blocks[0])
			logging.Info("选择的座位：" + pidSeat)
			code, message, err := booking.YesQueryLock(times.IdTime, pidSeat, blocks[0])
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
			fmt.Println(price)
			//fee, err := booking.YesFnEtcFree(times.IdTime)
			//if err != nil {
			//	logging.Error(err)
			//	continue
			//}
			//amountFee, _ := strconv.Atoi(fee)
			amountPrice, _ := strconv.Atoi(price)
			//amount := amountFee + amountPrice
			//fmt.Println(pSeat, pidSeat)
			//pSeat:="T192$252$188$174-1,"
			amount := amountPrice + 2000
			fmt.Println(amount)
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
func book(day string, idPerf string) booking.PaypalUrl {
	var payData booking.PaypalUrl
	times, err := booking.YesFnPerfTime(day, idPerf)
	if err != nil {
		logging.Error(err)
	}
	info, err := booking.GetBlockInfo(times.IdTime, times.IdHall)
	for _, v := range info {
		pSeat, price, err := booking.YesQuerySeatFlashEnd(times.IdTime, v.Class)
		fmt.Println(pSeat, price)
		if err != nil {
			fmt.Println(v)
			fmt.Println(err)
		}
		amountPrice, _ := strconv.Atoi(price)
		pidSeat, _ := booking.YesQuerySeat(times.IdTime, times.IdHall, v.Block)
		fmt.Println(pidSeat, amountPrice)
		//for len(pidSeat) > 0 {
		//	payData, err = booking.YesGetCart(idPerf, pidSeat, times.IdTime, pSeat, amountPrice+2000)
		//	if err != nil {
		//		logging.Error(err)
		//	}
		//}
	}
	return payData
}
