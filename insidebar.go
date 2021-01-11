package main

import (
	"fmt"
)

func DetectInsideBar(bars Bars) {
	for _, ticker := range bars.Tickers {
		if len(ticker.Bars) == 0 {
			continue

		}
		motherHigh := ticker.Bars[0].High
		motherLow := ticker.Bars[0].Low

		childHigh := ticker.Bars[1].High
		childLow := ticker.Bars[1].Low
		//fmt.Println(cfg.Deviance)


		if childHigh <= motherHigh + (motherHigh * cfg.Deviance) {
			fmt.Println(fmt.Sprintf(ticker.Name + " is %f <= %f + %f", childHigh, motherHigh, motherHigh * cfg.Deviance))
			if childLow >= motherLow - (motherLow * cfg.Deviance) {
				fmt.Println(fmt.Sprintf(ticker.Name + " is %f >= %f - %f", childLow, motherLow, motherLow * cfg.Deviance))
				fmt.Println(ticker.Name)
			}
		}
	}

}