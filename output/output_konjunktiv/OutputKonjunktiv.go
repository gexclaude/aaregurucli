package output_konjunktiv

import (
	"fmt"
	// "strconv"
	// "time"
	"../../api"
	// "../../texts"
	. "../../console"
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
	fmt.Println("I würd säge", CGreen(fmt.Sprint(aare.Temperature_text)))
}
