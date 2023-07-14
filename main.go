package main

import (
	"autobook/query"
	"time"
)

func main() {

	for {
		err := query.HttpSeatMap("1233748", "11243")
		if err != nil {
			continue
		}
		time.Sleep(time.Duration(10) * time.Second)
	}
	//times := query.FnPerfTime("20230812", "46450")
	//
	//pidSeat, class := query.HttpQuerySeat(times.PidTime, times.IdHall)
	//query.HttpQueryLock(times.PidTime, pidSeat)
	//pSeat := query.HttpQuerySeatFlashEnd(times.PidTime, class)
	//query.HttpGetCart("46450", pidSeat, times.PidTime, pSeat)

}
