package api

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"../config"
	"net/url"
	"fmt"
)

type AareGuruResponse struct {
	Status  string
	Aare    Aare
	Weather Weather
}

type Aare struct {
	Timestamp        int64
	Timestring       string
	Temperature      float32
	Temperature_text string
	Flow             float32
	Flow_text        string
}

type Weather struct {
	Current  WeatherInfos
	Today    WeatherToday
	Location string
}

type WeatherToday struct {
	V WeatherInfos
	N WeatherInfos
	A WeatherInfos
}

type WeatherInfos struct {
	Sy    string
	Syt   string
	Symt  int16
	Tt    float32
	Rr    int16
	Rrisk int16
}

func AskAareGuru(proxy *string, aareGuruResponseChannel chan<- AareGuruResponse, errChannel chan<- string, debug bool) {
	defer func() {
		if r := recover(); r != nil {
			errChannel <- fmt.Sprintf("%s", r)
		}
	}()

	var aareGuruResponse AareGuruResponse

	client := createHttpClient(proxy)

	response, err := client.Get(config.Endpoint_url)
	if err != nil { 
		panic(err)
	} else {
		if (debug) {
			fmt.Printf("Status: %s\n", response.Status)
		}
		data, err := ioutil.ReadAll(response.Body)
		if (err != nil) {
			panic(err)
		}

		json.Unmarshal(data, &aareGuruResponse)

		if (debug) {
			fmt.Printf("Raw: %s\n", string(data))
			fmt.Printf("Response: %s\n", aareGuruResponse)
		}
	}

	aareGuruResponseChannel <- aareGuruResponse
}
func createHttpClient(proxy *string) *http.Client {
	var myHttpClient *http.Client
	if *proxy == "" {
		myHttpClient = &http.Client{}
	} else {
		proxyUrl, err := url.Parse(*proxy)
		if err != nil {
			panic(err)
		}
		myHttpClient = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	}
	return myHttpClient
}
