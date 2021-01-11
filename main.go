package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
	//"syscall"
)

var(
	timeFrame string
	setup int
)

func init() {
	flag.Parse()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	loadConfig()
	var err error
	input, err := readLines(cfg.TickersFile)
	if err != nil {
		fmt.Println("Failure loading tickers file: " + cfg.TickersFile)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for true {
		fmt.Println("Please select a trading setup to scan for (1-2):")
		fmt.Println("    1.  Inside Bars")
		fmt.Println("    2.  Oversold Bounce")
		fmt.Println()

		userInput := getUserInput(reader)
		selection, err := strconv.Atoi(userInput)
		if err != nil || selection != 1 && selection != 2 {
			fmt.Println("Invalid entry. Please try again " + err.Error())
			continue
		}

		setup = selection
		fmt.Println("Selection Confirmed!")

		fmt.Println("Please select a time frame to scan (1-4):")
		fmt.Println("    1.  Weekly")
		fmt.Println("    2.  Daily")
		fmt.Println("    3.  Hourly")
		fmt.Println("    4.  15-minute")
		fmt.Println("    5.  5-minute")
		fmt.Println()

		userInput = getUserInput(reader)
		selection, err = strconv.Atoi(userInput)
		if err != nil || selection < 1 || selection > 5 {
			fmt.Println("Invalid entry. Please try again")
			continue
		}
		switch selection {
		case 1:
			timeFrame = "1W"
			break
		case 2:
			timeFrame = "1D"
			break
		case 3:
			timeFrame = "1H"
			break
		case 4:
			timeFrame = "15Min"
			break
		case 5:
			timeFrame = "5Min"
			break
		}

		fmt.Println("Selection Confirmed!")
		fmt.Println()
		break
	}

	client := &http.Client{
		Timeout: time.Second * 30, // time.Second * <sec for req timeout>
	}

	var tickers []string
	tickersCount := 0
	for _, inputLine := range input {
		if tickersCount == 199 {
			tickersCount = 0

			ProcessTickers(client, tickers)
			tickers = []string{}
		}

		tickers = append(tickers, inputLine)

		tickersCount++
	}

	// Process leftover tickers (if any)
	if len(tickers) > 0 {
		ProcessTickers(client, tickers)
	}


}

func getUserInput(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	text = strings.TrimSpace(text)

	return text
}

func ProcessTickers(client *http.Client, tickers []string) {
	var bars Bars
	if timeFrame == "1W" {
		bars = GetBars(client, "day", tickers, 10)
		bars = CondenseBars(bars, 5, "M")
	} else if timeFrame == "1H" {
		bars = GetBars(client, "15Min", tickers, 8)
		bars = CondenseBars(bars, 4, "")
	} else {
		bars = GetBars(client, timeFrame, tickers, 2)
	}

	if setup == 1 {
		DetectInsideBar(bars)
	}
}

func CondenseBars(bars Bars, factor int, stopper string) Bars {
	newBars := Bars{}
	for _, ticker := range bars.Tickers {
		newTicker := Ticker{Name: ticker.Name}
		progress := 0
		if len(ticker.Bars) == 0 {
			continue
		}
		for i := 0; i < len(ticker.Bars)/factor; i++ {
			b := Bar{Low: math.MaxFloat64}
			for x := 0; x < factor; x++ {
				if progress >= len(ticker.Bars) {
					break
				}
				b.High = math.Max(b.High, ticker.Bars[progress].High)
				b.Low = math.Min(b.Low, ticker.Bars[progress].Low)
				if x == 0 {
					b.Open = ticker.Bars[progress].Open
					b.Time = ticker.Bars[progress].Time
				}
				if x == factor-1 {
					b.Close = ticker.Bars[progress].Close
				}
				b.Volume += ticker.Bars[progress].Volume

				progress++
			}
			//fmt.Printf("Candle for %v at %d is C: %f O: %f L: %f H: %f V: %d\n", ticker.Name, b.Time, b.Close, b.Open, b.Low, b.High, b.Volume)
			newTicker.Bars = append(newTicker.Bars, b)
		}
		newBars.Tickers = append(newBars.Tickers, newTicker)
	}

	return newBars

}


// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			lines = append(lines, scanner.Text())
		}
	}
	return lines, scanner.Err()
}

