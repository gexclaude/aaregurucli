package outcities

import (
	"../api"
	"fmt"
	"sync"
)

func RenderCities(citiesResponseChannel chan api.CitiesResponse, errChannel chan string, wg *sync.WaitGroup) {

	defer func() {
		if r := recover(); r != nil {
			panic(r)
		}
	}()

	citiesResponse := readData(citiesResponseChannel, errChannel, wg)
	printOutput(*citiesResponse)
}

func readData(citiesResponseChannel chan api.CitiesResponse, errChannel chan string, wg *sync.WaitGroup) *api.CitiesResponse {
	var citiesResponse *api.CitiesResponse

	select {
	case tmp := <-citiesResponseChannel:
		citiesResponse = &tmp
	case err := <-errChannel:
		panic(err)
	}

	wg.Wait()

	return citiesResponse
}

func printOutput(citiesResponse api.CitiesResponse) {
	
    fmt.Printf("%-10s | %-10s | %s\n", "City *", "Name", "Gnaue Standort")
    fmt.Printf("%-10s | %-10s | %s\n", "----------", "----------", "----------------------")
    for _, elem := range citiesResponse.Cities {
        fmt.Printf("%-10s | %-10s | %s\n", elem.City, elem.Name, elem.Longname)
    }
    fmt.Println()
    fmt.Println("* gib das im command aus `city` argument a")
    fmt.Println()
}

