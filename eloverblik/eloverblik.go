package eloverblik

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

var eloverblikUrl = url.URL{
	Scheme: "https",
	Host:   "api.eloverblik.dk",
	Path:   "CustomerApi/api",
}

type Eloverblik struct {
}

func getToken() {
	// Get token, provide bearer header
	eloverblikUrl.JoinPath("token")
}

func fetchTimeseries(apiUrl url.URL) *EloverblikResponse {
	// Get payload, provide previously obtained token as bearer header
	apiUrl.JoinPath("meterdata/gettimeseries/2023-01-29/2023-01-30/Hour")
	request, err := http.Get(apiUrl.String())
	if err != nil {
		log.Fatal(err)
	}

	var data EloverblikResponse

	err = json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return &data
}
