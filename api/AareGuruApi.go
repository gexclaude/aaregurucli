package api

import (
	"../config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// AareGuruResponse root type of the aare.guru response
type AareGuruResponse struct {
	Status  string
	Aare    Aare
	Weather Weather
}

// Aare contains all values related to the aare and is contained in the AareGuruResponse
type Aare struct {
	Timestamp        int64
	Timestring       string
	Temperature      float32
	TemperatureText string `json:"temperature_text,omitempty"`
	Flow             float32
	FlowText        string `json:"flow_text,omitempty"`
}

// Weather holds weather related information from the AareGuruResponse
type Weather struct {
	Current  WeatherInfos
	Today    WeatherToday
	Location string
}

// WeatherToday represent today's weather split up into morning, afternoon and evening
type WeatherToday struct {
	V WeatherInfos
	N WeatherInfos
	A WeatherInfos
}

// WeatherInfos represents a weather snapshot
type WeatherInfos struct {
	Sy    string
	Syt   string
	Symt  int16
	Tt    float32
	Rr    int16
	Rrisk int16
}

// AskAareGuru ask aare.guru for an answer, returns an AareGuruResponse
func AskAareGuru(proxy *string, aareGuruResponseChannel chan<- AareGuruResponse, errChannel chan<- string, debug bool) {
	defer func() {
		if r := recover(); r != nil {
			errChannel <- fmt.Sprintf("%s", r)
		}
	}()

	var aareGuruResponse AareGuruResponse

	client := createHTTPClient(proxy)

	response, err := client.Get(config.EndpointURL)
	if err != nil {
		panic(err)
	} else {
		if debug {
			fmt.Printf("Status: %s\n", response.Status)
		}
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		json.Unmarshal(data, &aareGuruResponse)

		if debug {
			fmt.Printf("Raw: %s\n", string(data))
			fmt.Printf("Response: %s\n", aareGuruResponse)
		}
	}

	aareGuruResponseChannel <- aareGuruResponse
}
func createHTTPClient(proxy *string) *http.Client {
	var myHTTPClient *http.Client
	if *proxy == "" {
		myHTTPClient = &http.Client{}
	} else {
		proxyURL, err := url.Parse(*proxy)
		if err != nil {
			panic(err)
		}
		myHTTPClient = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	}
	return myHTTPClient
}
