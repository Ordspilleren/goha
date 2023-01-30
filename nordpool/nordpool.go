package nordpool

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var nordpoolUrl = url.URL{
	Scheme: "https",
	Host:   "www.nordpoolgroup.com",
	Path:   "api/marketdata/page/10",
}

type Unit int

const (
	MWH Unit = iota
	KWH
)

type Nordpool struct {
	Area          string
	Currency      string
	Unit          Unit
	PriceModifier func(price float64, time time.Time) float64
}

type SpotPrices struct {
	StartTime    time.Time
	EndTime      time.Time
	Updated      time.Time
	Currency     string
	Area         string
	HourlyPrices []HourlyPrice
}

type HourlyPrice struct {
	StartTime time.Time
	EndTime   time.Time
	Price     float64
}

func (n *Nordpool) GetSpotPricesForDate(endDate time.Time) *SpotPrices {
	date := endDate.Format("02-01-2006")
	q := nordpoolUrl.Query()
	q.Add("endDate", date)
	nordpoolUrl.RawQuery = q.Encode()
	response := n.fetch(nordpoolUrl)
	return n.parseResponse(response)
}

func (n *Nordpool) fetch(apiUrl url.URL) *NordpoolResponse {
	q := apiUrl.Query()
	q.Add("currency", n.Currency)
	apiUrl.RawQuery = q.Encode()
	request, err := http.Get(apiUrl.String())
	if err != nil {
		log.Fatal(err)
	}

	var data NordpoolResponse

	err = json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return &data
}

func (n *Nordpool) parseResponse(response *NordpoolResponse) *SpotPrices {
	startTime, _ := time.Parse("2006-01-02T15:04:05", response.Data.DataStartdate)
	endTime, _ := time.Parse("2006-01-02T15:04:05", response.Data.DataEnddate)
	updated, _ := time.Parse("2006-01-02T15:04:05", response.Data.DateUpdated)

	spotPrices := SpotPrices{
		StartTime: startTime,
		EndTime:   endTime,
		Updated:   updated,
		Currency:  n.Currency,
		Area:      n.Area,
	}

	areaData := []HourlyPrice{}

	for _, row := range response.Data.Rows {
		rowStartTime, _ := time.Parse("2006-01-02T15:04:05", row.StartTime)
		rowEndTime, _ := time.Parse("2006-01-02T15:04:05", row.EndTime)

		for _, column := range row.Columns {
			if column.Name != n.Area || row.IsExtraRow {
				continue
			}

			value, err := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(column.Value, " ", ""), ",", "."), 32)
			if err != nil {
				continue
			}

			if n.Unit == KWH {
				value = value / 1000
			}

			areaData = append(areaData, HourlyPrice{
				StartTime: rowStartTime,
				EndTime:   rowEndTime,
				Price:     n.PriceModifier(value, rowStartTime),
			})

		}
	}
	spotPrices.HourlyPrices = areaData

	return &spotPrices
}