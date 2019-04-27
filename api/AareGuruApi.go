package api

import (
	"encoding/json"
	"fmt"
	"github.com/gexclaude/aaregurucli/config"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// AareGuruResponse root type of the aare.guru response
type AareGuruResponse struct {
	Status  string
	Aare    Aare
	Weather Weather
}

// Aare contains all values related to the aare and is contained in the AareGuruResponse
type Aare struct {
	Timestamp       int64
	Timestring      string
	Temperature     float32
	TemperatureText string `json:"temperature_text,omitempty"`
	Flow            float32
	FlowText        string `json:"flow_text,omitempty"`
	Location        string
	LocationLong    string `json:"location_long,omitempty"`
}

// Weather holds weather related information from the AareGuruResponse
type Weather struct {
	Current WeatherInfos
	Today   WeatherToday
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
	Rr    float32
	Rrisk int16
}

// AskAareGuru ask aare.guru for an answer, returns an AareGuruResponse
func AskAareGuru(proxy *string, city *string, aareGuruResponseChannel chan<- AareGuruResponse, errChannel chan<- string, debug bool) {
	defer func() {
		if r := recover(); r != nil {
			errChannel <- fmt.Sprintf("%s", r)
		}
	}()

	var aareGuruResponse AareGuruResponse

	client := createHTTPClient(proxy)

	url := config.EndpointURL + "&city=" + strings.ToLower(*city)

	response, err := client.Get(url)
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
			fmt.Printf("Response: %#v\n", aareGuruResponse)
		}
	}

	aareGuruResponseChannel <- aareGuruResponse
}

// CitiesResponse contains all supported cities
type CitiesResponse struct {
	Cities []City
}

// City represents a single location supported by aareguru
type City struct {
	City     string
	Name     string
	Longname string
	Aare     float32
}

// AskAareGuruForCities ask aare.guru for cities, returns an CitiesResponse
func AskAareGuruForCities(proxy *string, citiesResponseChannel chan<- CitiesResponse, errChannel chan<- string, debug bool) {
	defer func() {
		if r := recover(); r != nil {
			errChannel <- fmt.Sprintf("%s", r)
		}
	}()

	var cities []City

	client := createHTTPClient(proxy)

	response, err := client.Get(config.CitiesEndpointURL)
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

		json.Unmarshal(data, &cities)

		if debug {
			fmt.Printf("Raw: %s\n", string(data))
			fmt.Printf("Response: %#v\n", cities)
		}
	}

	var citiesResponse CitiesResponse
	citiesResponse.Cities = cities
	citiesResponseChannel <- citiesResponse
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
