package main

import (
	"fmt"
	"./api"
	"./texts"
	"./output/output_typewriter"
	"./output/output_standard"
	. "./console"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"sync"
)

var (
	app           = kingpin.New("aareguru", texts.CLI_description)
	standard      = app.Command("standard", texts.CLI_command_standard_description)
	typewriter    = app.Command("schribmaschine", texts.CLI_command_typewriter_description)
	proxy         = app.Flag("proxy", texts.CLI_proxy_description).Short('p').String()
	colorless     = app.Flag("ohni-farb", texts.CLI_colorless_description).Short('f').Bool()
	noprogressbar = app.Flag("ohni-ladebauke", texts.CLI_noprogressbar_description).Short('l').Bool()
	debug         = app.Flag("debug", texts.CLI_proxy_description).Short('d').Hidden().Bool()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	InitConsole(!*colorless)
	ClearConsole()
	fmt.Println()

	aareGuruResponseChannel := make(chan api.AareGuruResponse)
	errChannel := make(chan string)

	var wg sync.WaitGroup

	defer func() {
		if r := recover(); r != nil {
			fmt.Println()
			fmt.Println(CRed(texts.Error_Detail_msg))
			fmt.Println(r)
			fmt.Println()
			fmt.Print(texts.Error_Hints_msg)
			fmt.Print()
		}
		BeforeExitConsole()
	}()

	go func() {
		defer wg.Done()
		wg.Add(1)
		api.AskAareGuru(proxy, aareGuruResponseChannel, errChannel, *debug)
	}()

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case standard.FullCommand():
		output_standard.Init(! *noprogressbar)
		output_standard.RenderAareGuruResponse(aareGuruResponseChannel, errChannel, &wg)
	case typewriter.FullCommand():
		output_typewriter.Init()
		output_typewriter.RenderAareGuruResponse(aareGuruResponseChannel, errChannel, &wg)
	}
}
