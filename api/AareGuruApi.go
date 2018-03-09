package api

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"../config"
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
	Current WeatherInfos
	Today   WeatherToday
}

type WeatherToday struct {
	V WeatherInfos
	N WeatherInfos
	A WeatherInfos
}

type WeatherInfos struct {
	Sy    string
	Symt  int16
	Tt    float32
	Rr    int16
	Rrisk int16
}

func AskAareGuru(aareGuruResponseChannel chan<-AareGuruResponse, errChannel chan<-string) {
	defer func() {
		if r := recover(); r != nil {
			errChannel <- ""
		}
	}()
	
	var aareGuruResponse AareGuruResponse

	response, err := http.Get(config.Endpoint_url)
	if err != nil {
		panic(err)
	} else {
		data, err := ioutil.ReadAll(response.Body)
		if (err != nil) {
			panic(err)
		}

		json.Unmarshal(data, &aareGuruResponse)
	}

	aareGuruResponseChannel <- aareGuruResponse
}
