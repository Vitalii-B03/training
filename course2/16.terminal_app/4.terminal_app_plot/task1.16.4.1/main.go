package main

import (
	"encoding/json"
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/gosuri/uilive"
	"github.com/guptarohit/asciigraph"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
	//"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Activ struct {
}
type PriceData struct {
	BuyPrice  string `json:"buy_price"`
	SellPrice string `json:"sell_price"`
	LastTrade string `json:"last_trade"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Avg       string `json:"avg"`
	Vol       string `json:"vol"`
	VolCurr   string `json:"vol_curr"`
	Updated   int64  `json:"updated"`
}

func Grafprint(client *http.Client, req *http.Request, data []float64, writer *uilive.Writer, ticket string) ([]float64, error) {

	res, err := client.Do(req)
	if err != nil {
		return data, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return data, err
	}

	var prices map[string]PriceData
	err = json.Unmarshal((body), &prices)
	if err != nil {
		return data, err
	}

	newDataPoint, err := strconv.ParseFloat(prices[ticket].BuyPrice, 64)
	if err != nil {
		return data, err
	}

	data = append(data, newDataPoint)

	// Ограничиваем размер данных, чтобы график не становился слишком большим
	if len(data) > 20 {
		data = data[1:]
	}
	now := time.Now()
	timeFormatted := now.Format("15:04:05")
	dateFormatted := now.Format("2006-01-02")
	graph := asciigraph.Plot(data, asciigraph.Height(10), asciigraph.Width(100))

	fmt.Fprintln(writer, ticket, ":", data[len(data)-1], "\n"+graph, "\nТекущая дата: ", dateFormatted, "\nТекущее время: ", timeFormatted)

	return data, nil
}

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	url := "https://api.exmo.com/v1.1/ticker"
	method := "POST"
	payload := strings.NewReader("")
	client := &http.Client{}

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	data := []float64{}
	rand.Seed(time.Now().UnixNano())
	//________________________________________________________________________
	writer := uilive.New()
	ticket1 := "BTC_USD"
	ticket2 := "LTC_USD"
	ticket3 := "ETH_USD"

	var currentStop chan bool
	var runningGorout bool
	var wg sync.WaitGroup

	runGraph := func(ticket string) {
		if runningGorout {
			currentStop <- true
			<-currentStop
			wg.Wait()
		}
		currentStop = make(chan bool)
		runningGorout = true
		wg.Add(1)
		go func() {
			defer wg.Done()
			writer.Start()
			fmt.Fprintln(writer, ticket)
			for {
				select {
				case <-currentStop:
					writer.Stop()
					currentStop <- true
					runningGorout = false
					return
				default:
					data, err = Grafprint(client, req, data, writer, ticket)
					if err != nil {
						log.Fatal(err)
					}
					time.Sleep(1000 * time.Millisecond)
				}

			}
		}()
	}
	for {
		fmt.Print("\033[H\033[2J")
		fmt.Println("1. BTC_USD")
		fmt.Println("2. LTC_USD")
		fmt.Println("3. ETH_USD")
		fmt.Println()
		fmt.Println("Press 1-3 to change symbol, press ESC to exit")

		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}
		if key == keyboard.KeyEsc {
			fmt.Println("выход из программы...")
			if runningGorout {
				currentStop <- true
				<-currentStop
				wg.Wait()
			}
			break
		}
		if key == keyboard.KeyBackspace2 {
			if runningGorout {
				currentStop <- true
				<-currentStop
				wg.Wait()
			}
			data = []float64{}
			continue
		}
		switch char {
		case '1':
			fmt.Fprintln(writer, "\033[H\033[2J")
			data = []float64{}
			runGraph(ticket1)
		case '2':
			fmt.Fprintln(writer, "\033[H\033[2J")
			data = []float64{}
			runGraph(ticket2)
		case '3':
			fmt.Fprintln(writer, "\033[H\033[2J")
			data = []float64{}
			runGraph(ticket3)
		}
	}
}
