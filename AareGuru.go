package main

import (
	"fmt"
	"time"
	"./api"
	"./texts"
	"./output/output_simple"
	. "./console"
	"gopkg.in/cheggaaa/pb.v1"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

const progressBarCount = 100

var (
	app = kingpin.New("aareguru", texts.CLI_description)
	proxy = app.Flag("proxy", texts.CLI_proxy_description).Short('p').String()
	colorless = app.Flag("ohni-faarb", texts.CLI_colorless_description).Short('o').Bool()
	debug = app.Flag("debug", texts.CLI_proxy_description).Short('d').Hidden().Bool()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	InitConsole(!*colorless)
	ClearConsole()

	aareGuruResponseChannel := make(chan api.AareGuruResponse)
	errChannel := make(chan string)

	bar := createBar()

	defer func() {
		if r := recover(); r != nil {
			bar.Finish()
			// check for isFinished synchronized things
			if bar.IsFinished() {
				fmt.Println()
				fmt.Println(CRed(texts.Error_Detail_msg))
				fmt.Println(r)
				fmt.Println()
				fmt.Print(texts.Error_Hints_msg)
				fmt.Print()
			}
		}
		BeforeExitConsole()
	}()

	go api.AskAareGuru(proxy, aareGuruResponseChannel, errChannel, *debug)

	aareGuruResponse := readData(aareGuruResponseChannel, errChannel, bar)

	if bar.IsFinished() {
		output_simple.PrintBanner()
		output_simple.PrintOutput(*aareGuruResponse)
	} else {
		fmt.Println(CRed("Oops"))
	}
}

func readData(aareGuruResponseChannel chan api.AareGuruResponse, errChannel chan string, bar *pb.ProgressBar) *api.AareGuruResponse {
	var aareGuruResponse *api.AareGuruResponse
	i := 0
	// only first part of progress bar
	for ; i < progressBarCount-int(progressBarCount*.75); i++ {
		if aareGuruResponse == nil {
			select {
			case tmp := <-aareGuruResponseChannel:
				aareGuruResponse = &tmp
			case err := <- errChannel:
				panic(err)
			}
		}

		increaseBar(bar, i)
	}
	if aareGuruResponse == nil {
		tmp := <-aareGuruResponseChannel
		aareGuruResponse = &tmp
	}
	// rest of progress bar
	for ; i < progressBarCount; i++ {
		increaseBar(bar, i)
	}
	bar.FinishPrint(CGreen(texts.Success_msg))
	fmt.Println()

	return aareGuruResponse
}

func createBar() *pb.ProgressBar {
	bar := pb.New(progressBarCount).Prefix(texts.Loading_msg)
	bar.SetRefreshRate(50)
	bar.SetWidth(74)
	bar.SetMaxWidth(74)
	bar.ShowCounters = false
	bar.ShowElapsedTime = false
	bar.ShowFinalTime = false
	bar.Start()
	return bar
}

func increaseBar(bar *pb.ProgressBar, i int) {
	bar.Increment()
	if i < progressBarCount*0.5 {
		time.Sleep(time.Millisecond * 4)
	} else if i < progressBarCount*0.7 {
		time.Sleep(time.Millisecond * 6)
	} else {
		time.Sleep(time.Millisecond * 7)
	}
}
