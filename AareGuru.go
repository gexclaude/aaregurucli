package main

import (
	"fmt"
	"time"
	"./api"
	"./texts"
	"./output/output_simple"
	. "./console"
	"github.com/gosuri/uiprogress"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"sync"
)

const progressBarCount = 100

var (
	app           = kingpin.New("aareguru", texts.CLI_description)
	proxy         = app.Flag("proxy", texts.CLI_proxy_description).Short('p').String()
	colorless     = app.Flag("ohni-farb", texts.CLI_colorless_description).Short('f').Bool()
	noprogressbar = app.Flag("ohni-ladebauke", texts.CLI_noprogressbar_description).Short('l').Bool()
	debug         = app.Flag("debug", texts.CLI_proxy_description).Short('d').Hidden().Bool()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	InitConsole(!*colorless)
	ClearConsole()

	aareGuruResponseChannel := make(chan api.AareGuruResponse)
	errChannel := make(chan string)

	bar := createBar()
	var wg sync.WaitGroup

	defer func() {
		if r := recover(); r != nil {
			uiprogress.Stop()
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

	aareGuruResponse := readData(aareGuruResponseChannel, errChannel, bar, &wg)

	fmt.Println()
	output_simple.PrintBanner()
	output_simple.PrintOutput(*aareGuruResponse)
}

func readData(aareGuruResponseChannel chan api.AareGuruResponse, errChannel chan string, bar *uiprogress.Bar, wg *sync.WaitGroup) *api.AareGuruResponse {
	var aareGuruResponse *api.AareGuruResponse
	i := 0
	// only first part of progress bar
	for ; isProgressBar() && i < progressBarCount-int(progressBarCount*.75); i++ {
		if aareGuruResponse == nil {
			select {
			case tmp := <-aareGuruResponseChannel:
				aareGuruResponse = &tmp
			case err := <-errChannel:
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
	for ; isProgressBar() && i < progressBarCount; i++ {
		increaseBar(bar, i)
	}
	wg.Wait()
	stopBar()

	fmt.Println(CGreen(texts.Success_msg))
	fmt.Println()

	return aareGuruResponse
}

func isProgressBar() bool {
	return !noProgressBar();
}

func noProgressBar() bool {
	return *noprogressbar;
}

func createBar() *uiprogress.Bar {
	if (noProgressBar()) {
		return nil;
	}

	bar := uiprogress.AddBar(progressBarCount).AppendCompleted().PrependElapsed()
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return fmt.Sprintf("%s (%d/%d)", texts.Loading_msg, b.Current(), progressBarCount)
	})
	uiprogress.Start()
	return bar
}

func stopBar() {
	if (isProgressBar()) {
		uiprogress.Stop()
	}
}

func increaseBar(bar *uiprogress.Bar, i int) {
	if (isProgressBar()) {
		bar.Incr()
		if i < progressBarCount*0.5 {
			time.Sleep(time.Millisecond * 2)
		} else if i < progressBarCount*0.7 {
			time.Sleep(time.Millisecond * 3)
		} else {
			time.Sleep(time.Millisecond * 4)
		}
	}
}
