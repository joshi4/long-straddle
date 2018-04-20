package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) != 5 {
		fmt.Println("usage: long-straddle sharePrice strikePrice callFees putFees")
		os.Exit(1)
	}
	args = args[1:]
	sp, strikePrice, call, put := args[0], args[1], args[2], args[3]
	sr := calcRange(sp, strikePrice, call, put)
	fmt.Println(sr)
}

type straddleRange struct {
	putBreakEven      float64
	callBreakEven     float64
	putPercentChange  float64
	callPercentChange float64
}

func (s straddleRange) String() string {
	return fmt.Sprintf("putBreakEvenPrice: %f putPercentChange: %f\ncallBreakEven: %f callPercentChange: %f", s.putBreakEven, s.putPercentChange, s.callBreakEven, s.callPercentChange)
}

func calcRange(sharePriceStr, strikePriceStr, callFeesStr, putFeesStr string) straddleRange {
	// strike price is same for both options:
	sharePrice, _ := strconv.ParseFloat(sharePriceStr, 64)
	strikePrice, _ := strconv.ParseFloat(strikePriceStr, 64)
	callFees, _ := strconv.ParseFloat(callFeesStr, 64)
	putFees, _ := strconv.ParseFloat(putFeesStr, 64)

	var s straddleRange
	s.callBreakEven = strikePrice + callFees + putFees
	s.callPercentChange = calcPercent(s.callBreakEven, sharePrice)

	s.putBreakEven = strikePrice - (callFees + putFees)
	s.putPercentChange = calcPercent(s.putBreakEven, sharePrice)
	return s
}

func calcPercent(nr, dr float64) float64 {
	delta := (nr - dr) / dr
	return delta * 100.0
}
