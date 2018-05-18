package outkonjunktiv

import (
	"fmt"
	"regexp"
	// "strconv"
	// "time"
	"../api"
	// "../../texts"
	. "../console"
	// . "../output_common"
	"sync"
)

func Init() {
}

func RenderAareGuruResponse(aareGuruResponseChannel chan api.AareGuruResponse, errChannel chan string, wg *sync.WaitGroup) {
	aareGuruResponse := readData(aareGuruResponseChannel, errChannel, wg)
	printOutput(*aareGuruResponse)
}

func readData(aareGuruResponseChannel chan api.AareGuruResponse, errChannel chan string, wg *sync.WaitGroup) *api.AareGuruResponse {
	var aareGuruResponse *api.AareGuruResponse

	select {
	case tmp := <-aareGuruResponseChannel:
		aareGuruResponse = &tmp
	case err := <-errChannel:
		panic(err)
	}

	wg.Wait()

	return aareGuruResponse
}

func printOutput(aareGuruResponse api.AareGuruResponse) {
	aare := aareGuruResponse.Aare

	fmt.Println("D Aaare wär", CGreen(fmt.Sprint(aare.Temperature)), "Grad warm.")
	fmt.Println("I würd säge", CGreen(fmt.Sprint(konjunktify(aare.TemperatureText))))
	fmt.Println(CGreen(fmt.Sprint(konjunktify("Isch das wahr?"))))
}

func konjunktify(text string) string {
	
	for _, s := range getSubstitutions() {
    // only match entire words, cannot go case insensitiv as some words would match nouns.
		re := regexp.MustCompile("\\b" + s.original + "\\b")
		text = re.ReplaceAllString(text, s.replacement)
	}

	return text
} 
