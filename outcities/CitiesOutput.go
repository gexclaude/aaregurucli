package outcities

import (
	"../api"
	"../asciiart"
	"../console"
	"../texts"
	"fmt"
	"strconv"
	"sync"
)

// RenderCities renders a table of cities
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
	fmt.Println(boxHorizontalLine())
	fmt.Print(console.CBlue(asciiart.Banner))

	fmt.Println(fmt.Sprintf("|%-11s+%-12s+%6s+%s|", "------------", "------------", "------", "---------------------------------------"))
	fmt.Println(box(fmt.Sprintf("%-10s | %-10s | %4s | %s", "City *", "Name", texts.DegreeCelsiusLabel, "Gnaue Standort")))
	fmt.Println(fmt.Sprintf("|%-11s+%-12s+%6s+%s|", "------------", "------------", "------", "---------------------------------------"))
	for _, city := range citiesResponse.Cities {
		fmt.Println(box(
			fmt.Sprintf("%-10s | %-10s | %s | %s",
				city.City,
				city.Name,
				console.CBlue(fmt.Sprintf("%4.1f", city.Aare)),
				city.Longname),
			console.CBlue("")))
	}

	fmt.Println(boxHorizontalLine())
	fmt.Println(box("* gib das im Command aus `city` argument a"))
	fmt.Println(boxHorizontalLine())
	fmt.Println()
}

func boxHorizontalLine() string {
	return "+------------------------------------------------------------------------+"
}

func box(str string, colorChars ...string) string {
	return fmt.Sprintf("| %-"+strconv.Itoa(70+colorCharsLength(colorChars...))+"s |", str)
}

func colorCharsLength(colorChars ...string) int {
	var colorCharsLen = 0
	for _, element := range colorChars {
		colorCharsLen += len(element)
	}
	return colorCharsLen
}
