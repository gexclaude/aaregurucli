package main

import (
	"./api"
	"./console"
	"./outcities"
	"./outstd"
	"./outtypewrt"
	"./texts"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"sync"
)

var (
	app            = kingpin.New("aareguru", texts.CliDescription)
	standard       = app.Command("standard", texts.CliCommandStandardDescription).Default()
	cityStandard   = standard.Arg("city", texts.CliCityDescription).String()
	
	typewriter     = app.Command("schribmaschine", texts.CliCommandTypewriterDescription)
	cityTypewriter = typewriter.Arg("city", texts.CliCityDescription).String()
	
	cities         = app.Command("cities", texts.CliCommandCitiesDescription)
	proxy          = app.Flag("proxy", texts.CliProxyDescription).Short('p').String()
	colorless      = app.Flag("ohni-farb", texts.CliColorlessDescription).Short('f').Bool()
	noprogressbar  = app.Flag("ohni-ladebauke", texts.CliNoprogressbarDescription).Short('l').Bool()
	debug          = app.Flag("debug", texts.CliProxyDescription).Short('d').Hidden().Bool()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	console.InitConsole(!*colorless)
	console.ClearConsole()
	fmt.Println()

	errChannel := make(chan string)

	var wg sync.WaitGroup

	defer func() {
		if r := recover(); r != nil {
			fmt.Println()
			fmt.Println(console.CRed(texts.ErrorDetailMsg))
			fmt.Println(r)
			fmt.Println()
			fmt.Print(texts.ErrorHintsMsg)
			fmt.Print()
		}
		console.BeforeExitConsole()
	}()

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case standard.FullCommand():
		outstd.Init(!*noprogressbar)
		aareGuruResponseChannel := askAareGuru(errChannel, &wg, cityStandard)
		outstd.RenderAareGuruResponse(aareGuruResponseChannel, errChannel, &wg)
	case typewriter.FullCommand():
		outtypewrt.Init()
		aareGuruResponseChannel := askAareGuru(errChannel, &wg, cityTypewriter)
		outtypewrt.RenderAareGuruResponse(aareGuruResponseChannel, errChannel, &wg)
    case cities.FullCommand():
        citiesResponseChannel := askAareGuruForCities(errChannel, &wg)
        outcities.RenderCities(citiesResponseChannel, errChannel, &wg)
	}
}

func askAareGuru (errChannel chan<- string, wg *sync.WaitGroup, city *string) chan api.AareGuruResponse {
	aareGuruResponseChannel := make(chan api.AareGuruResponse)
	go func() {
		defer wg.Done()
		wg.Add(1)
		api.AskAareGuru(proxy, city, aareGuruResponseChannel, errChannel, *debug)
	}()
	return aareGuruResponseChannel 
}

func askAareGuruForCities (errChannel chan<- string, wg *sync.WaitGroup) chan api.CitiesResponse {
	citiesResponseChannel := make(chan api.CitiesResponse)
	go func() {
		defer wg.Done()
		wg.Add(1)
		api.AskAareGuruForCities(proxy, citiesResponseChannel, errChannel, *debug)
	}()
	return citiesResponseChannel 
}