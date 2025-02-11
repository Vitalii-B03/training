package main

import (
	"fmt"
	"github.com/mattevans/dinero"
	"log"
	"time"
)

func currencyPairRate(a, b string, c float64) float64 {
	client := dinero.NewClient(
		"3295ac27d55e42f3959b2521ae5e1303",
		"USD",
		20*time.Minute,
	)
	rsp, err := client.Rates.List()
	if err != nil {
		log.Fatal(err)

	}
	rsp1, err := client.Rates.List()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rsp)
	return rsp.Rates[b] / rsp1.Rates[a] * c
}
func main() {

	rate := currencyPairRate("CAD", "EUR", 100.0)
	fmt.Println(rate)
}
