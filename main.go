package main

import (
	"autobook/booking"
	"strconv"
	"sync"
	"time"

	"autobook/logging"

	"fmt"
)

var Bs = []booking.BlockInfo{{"101", "GROUND석", ""}, {"102", "GROUND석", ""}, {"103", "GROUND석", ""}, {"104", "GROUND석", ""}, {"105", "GROUND석", ""}, {"106", "GROUND석", ""}, {"107", "GROUND석", ""},
	{"108", "GROUND석", ""}, {"109", "GROUND석", ""}, {"110", "GROUND석", ""}, {"111", "GROUND석", ""}, {"112", "GROUND석", ""}, {"113", "GROUND석", ""}, {"114", "GROUND석", ""}, {"115", "GROUND석", ""}, {"116", "GROUND석", ""}, {"117", "GROUND석", ""}, {"118", "GROUND석", ""}, {"119", "GROUND석", ""}, {"120", "GROUND석", ""}, {"121", "GROUND석", ""}, {"120", "GROUND석", ""}, {"121", "GROUND석", ""}, {"122", "GROUND석", ""}, {"123", "GROUND석", ""}, {"124", "GROUND석", ""}, {"125", "GROUND석", ""},
	{"201", "2층지정석", ""}, {"202", "2층지정석", ""}, {"203", "2층지정석", ""}, {"204", "2층지정석", ""}, {"205", "2층지정석", ""}, {"206", "2층지정석", ""}, {"207", "2층지정석", ""}, {"208", "2층지정석", ""}, {"209", "2층지정석", ""}, {"210", "2층지정석", ""}, {"211", "2층지정석", ""}, {"212", "2층지정석", ""}, {"213", "2층지정석", ""}, {"214", "2층지정석", ""}, {"215", "2층지정석", ""}, {"216", "2층지정석", ""}, {"217", "2층지정석", ""}, {"218", "2층지정석", ""}, {"219", "2층지정석", ""}, {"220", "2층지정석", ""},
	{"221", "2층지정석", ""}, {"222", "2층지정석", ""}, {"223", "2층지정석", ""}, {"224", "2층지정석", ""}, {"225", "2층지정석", ""}}

//var Bs =[]booking.BlockInfo{{"3","S석",""},{"1","S석",""}}
func main() {
	booking.InitRedis()
	var url booking.PaypalUrl
	//
	//for {
	//	url = book("20230826", "46706")
	//	url := bookYes24("20230812", "46560")
	//	if len(url.EncryptPaypalOrderID) > 0 {
	//		fmt.Println(url)
	//		break
	//	}
	//	time.Sleep(time.Duration(1) * time.Second)
	//}
	url = book("20230826", "46706")
	cmd := booking.Redisclient.Set("encryptCartID", url.EncryptCartID, time.Duration(60)*time.Second)
	fmt.Println(cmd.Err(), url.EncryptCartID)
	c := booking.Redisclient.Set("encryptPaypalOrderID", url.EncryptPaypalOrderID, time.Duration(60)*time.Second)
	fmt.Println(c.Err(), url.EncryptPaypalOrderID)
	//payPal()

}
func payPal() {
	encryptCartID, _ := booking.Redisclient.Get("encryptCartID").Result()
	encryptPaypalOrderID, _ := booking.Redisclient.Get("encryptPaypalOrderID").Result()
	if len(encryptCartID) > 0 {
		booking.YesPaypalPayResponse(encryptCartID, encryptPaypalOrderID, encryptPaypalOrderID)
	} else {
		fmt.Println("空")
	}

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
			if len(pidSeat) == 0 {
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
	var wg sync.WaitGroup

	var payData booking.PaypalUrl
	times, err := booking.YesFnPerfTime(day, idPerf)
	fmt.Println(times)
	//if err != nil {
	//	logging.Error(err)
	//}
	//info, err := booking.GetBlockInfo(times.IdTime, times.IdHall)
	fmt.Println(err)
	wg.Add(1)
	for _, v := range Bs {

		go wgBooking(idPerf, times.IdTime, v.Class, times.IdHall, v.Block, &wg)

	}
	fmt.Println("-----1")
	wg.Wait()
	return payData
}

func wgBooking(idPerf string, pIdTime string, pCntClass string, idHall string, block string, wg *sync.WaitGroup) {
	fmt.Println(block)
	for {
		pSeat, price, err := booking.YesQuerySeatFlashEnd(pIdTime, pCntClass)
		if err != nil {
			return
		}
		amountPrice, _ := strconv.Atoi(price)
		pidSeat, _ := booking.YesQuerySeat(pIdTime, idHall, block)

		fmt.Println(pidSeat, amountPrice)

		if len(pidSeat) > 0 {
			payData, err := booking.YesGetCart(idPerf, pidSeat, pIdTime, pSeat, amountPrice+3000)
			if err != nil {
				logging.Error(err)
				continue
			}
			if len(payData.EncryptPaypalOrderID) > 0 {
				fmt.Println(block)
				logging.Info(payData)
				fmt.Println(payData)
				break

			} else {
				continue
			}
		}
		time.Sleep(time.Duration(10) * time.Second)
	}
	wg.Done()
	return

}
