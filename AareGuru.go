package main

import (
	"./api"
	"./console"
	"./outstd"
	"./outtypewrt"
	"./texts"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"sync"
)

var (
	app           = kingpin.New("aareguru", texts.CliDescription)
	standard      = app.Command("standard", texts.CliCommandStandardDescription).Default()
	typewriter    = app.Command("schribmaschine", texts.CliCommandTypewriterDescription)
	proxy         = app.Flag("proxy", texts.CliProxyDescription).Short('p').String()
	colorless     = app.Flag("ohni-farb", texts.CliColorlessDescription).Short('f').Bool()
	noprogressbar = app.Flag("ohni-ladebauke", texts.CliNoprogressbarDescription).Short('l').Bool()
	debug         = app.Flag("debug", texts.CliProxyDescription).Short('d').Hidden().Bool()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	console.InitConsole(!*colorless)
	console.ClearConsole()
	fmt.Println()

	aareGuruResponseChannel := make(chan api.AareGuruResponse)
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

	go func() {
		defer wg.Done()
		wg.Add(1)
		api.AskAareGuru(proxy, aareGuruResponseChannel, errChannel, *debug)
	}()

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case standard.FullCommand():
		outstd.Init(!*noprogressbar)
		outstd.RenderAareGuruResponse(aareGuruResponseChannel, errChannel, &wg)
	case typewriter.FullCommand():
		outtypewrt.Init()
		outtypewrt.RenderAareGuruResponse(aareGuruResponseChannel, errChannel, &wg)
	}
}
