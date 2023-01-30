package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Ordspilleren/homeautomation/nordpool"
)

func NordpoolToVM() {
	nordpool := nordpool.Nordpool{
		Area:     "DK1",
		Currency: "DKK",
		Unit:     nordpool.KWH,
		PriceModifier: func(price float64, time time.Time) float64 {
			elafgift := 0.008
			systemtarif := 0.054
			transmissionsnettarif := 0.058
			elselskabAbonnement := 0.0
			nettarifLav := 0.1862
			nettarifSpids := 0.4887
			vat := 1.25

			if time.Month() >= 10 || time.Month() < 4 {
				if time.Hour() >= 17 && time.Hour() < 20 {
					return (price + elafgift + systemtarif + transmissionsnettarif + elselskabAbonnement + nettarifSpids) * vat
				} else {
					return (price + elafgift + systemtarif + transmissionsnettarif + elselskabAbonnement + nettarifLav) * vat
				}
			} else {
				return (price + elafgift + systemtarif + transmissionsnettarif + elselskabAbonnement + nettarifLav) * vat
			}
		},
	}

	nordpoolSpotPrices := nordpool.GetSpotPricesForDate(time.Now().AddDate(0, 0, 1))

	var lines []string

	for _, priceSegment := range nordpoolSpotPrices.HourlyPrices {
		lines = append(lines, fmt.Sprintf("nordpool{area=\"%s\",currency=\"%s\"} %f %d\n", nordpoolSpotPrices.Area, nordpoolSpotPrices.Currency, priceSegment.Price, priceSegment.StartTime.UnixMilli()))
	}

	log.Print(strings.Join(lines, ""))

	payload := strings.NewReader(strings.Join(lines, ""))

	response, err := http.Post("http://localhost:8428/api/v1/import/prometheus", "text/plain", payload)
	if err != nil {
		log.Panic(err)
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(bodyBytes)
}
