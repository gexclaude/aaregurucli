package main

import (
	"fmt"
	"time"
    "math/rand"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"./asciiart"
	"./model"
)

func main() {
	fmt.Print(asciiart.Banner)
	
	printOutput(readAareGuru())
}

func printOutput(aareGuruResponse model.AareGuruResponse) {
	t := time.Unix(aareGuruResponse.Aare.Timestamp, 0)
	glaeser := liter_to_glas(m3_to_liter(aareGuruResponse.Aare.Flow))
	fmt.Printf("%7.0f - %s\n", glaeser, random_flow_in_glas_text())

	fmt.Printf("Itz grad (%02d:%02d - %02d.%02d.%04d)\n", t.Hour(), t.Minute(), t.Day(), t.Month(), t.Year())
	fmt.Printf("%4.1f %-4s - %s\n", aareGuruResponse.Aare.Temperature, "C°", aareGuruResponse.Aare.Temperature_text)
	fmt.Printf("%4.0f %4s - %s\n", aareGuruResponse.Aare.Flow, "m3/s", aareGuruResponse.Aare.Flow_text)
	fmt.Println()
	fmt.Printf("I öpe 2h\n")
	fmt.Printf("%4.1f %-4s - %s\n", aareGuruResponse.Aare.Forecast2h, "C°", aareGuruResponse.Aare.Forecast2h_text)
}

func readAareGuru() model.AareGuruResponse {
	var aareGuruResponse model.AareGuruResponse
	
	response, err := http.Get("http://aareguru.existenz.ch/currentV2.php?app=aare.guru.CLI")
	if err != nil {
		fmt.Println("Dr aare.guru isch verärgeret. Är git üs kei antwort meh.")
	} else {
		data, err := ioutil.ReadAll(response.Body)
		if (err != nil) {
			panic(err)
		}

		json.Unmarshal(data, &aareGuruResponse)
	}
	
	return aareGuruResponse
}


func m3_to_liter(m3 float32) float32 {
	return m3 * (10 * 10 * 10)
}

func liter_to_glas(liter float32) float32 {
	return liter / 0.3 // 3 dl
} 

func random_flow_in_glas_text() string {
	if(rand_bool()) {
		return "Stange Bier/Sec. "
	} else {
		return "Sirupgleser/Sec"
	}
}

func rand_bool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Float32() < 0.5
}