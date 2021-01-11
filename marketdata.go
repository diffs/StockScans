package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Bar struct {
	Time   int     `json:"t"`
	Open   float64 `json:"o"`
	High   float64 `json:"h"`
	Low    float64 `json:"l"`
	Close  float64 `json:"c"`
	Volume int     `json:"v"`
}

type Ticker struct {
	Name string
	Bars []Bar
}

type Bars struct {
	Tickers []Ticker
}

func GetBars(client *http.Client, timeFrame string, symbols[] string, barcount int) Bars {
	request, _ := http.NewRequest("GET", "https://data.alpaca.markets/v1/bars/" + timeFrame + "?symbols=" + strings.Join(symbols, ",") + "&limit=" + strconv.Itoa(barcount), nil)
	request.Header.Add("User-Agent", cfg.UserAgent)
	request.Header.Add("APCA-API-KEY-ID", cfg.AlpacaKeyID)
	request.Header.Add("APCA-API-SECRET-KEY", cfg.AlpacaSecretKey)

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("Error while getting bars: " + err.Error())
		return Bars{}
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var f interface{}
	if err := json.Unmarshal(body, &f); err != nil {
		fmt.Println("Json Unmarshal error: " + err.Error())
		return Bars{}
	}

	var bars = Bars{}
	for idek2, val := range f.(map[string]interface{}) {
		ticker := Ticker{}
		if x, ok := val.([]interface{}); ok {
			for _, iface := range x {
				bar := Bar{}
				for ifaceKey, ifaceVal := range iface.(map[string]interface{}) {

					switch ifaceKey {
					case "t":
						bar.Time = int(ifaceVal.(float64))
					case "v":
						bar.Volume = int(ifaceVal.(float64))
					case "o":
						bar.Open = ifaceVal.(float64)
					case "c":
						bar.Close = ifaceVal.(float64)
					case "h":
						bar.High = ifaceVal.(float64)
					case "l":
						bar.Low = ifaceVal.(float64)
					}
				}
				ticker.Bars = append(ticker.Bars, bar)
			}
		}
		ticker.Name = idek2
		bars.Tickers = append(bars.Tickers, ticker)

	}

	return bars

}