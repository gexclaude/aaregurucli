package api

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"../config"
	"../texts"
)

type Aare struct {
	Timestamp int64
	Timestring string
	Temperature float32
	Temperature_text string
	Flow float32
	Flow_text string
	Forecast2h float32
	Forecast2h_text string
}

type AareGuruResponse struct {
	Status string
	Aare   Aare
}

func AskAareGuru() AareGuruResponse {
	var aareGuruResponse AareGuruResponse

	response, err := http.Get(config.Endpoint_url)
	if err != nil {
		fmt.Println(texts.Error_msg)
	} else {
		data, err := ioutil.ReadAll(response.Body)
		if (err != nil) {
			panic(err)
		}

		json.Unmarshal(data, &aareGuruResponse)
	}

	return aareGuruResponse
}