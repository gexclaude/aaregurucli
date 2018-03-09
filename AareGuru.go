package main

import (
	"fmt"
	"./api"
	"./console"
	"./texts"
	"./output/output_simple"
	. "github.com/logrusorgru/aurora"
	"gopkg.in/cheggaaa/pb.v1"
	"time"
)

const progressBarCount = 100

func main() {
	console.Init()
	console.CallClear()

	aareGuruResponseChannel := make(chan api.AareGuruResponse)
	errChannel := make(chan string)

	bar := createBar()

	defer func() {
		if r := recover(); r != nil {
			bar.Finish()
			fmt.Println()
			fmt.Println(Red(texts.Error_Detail_msg), r)
			fmt.Print(texts.Error_Hints_msg)
			fmt.Print()
		}
		console.BeforeExit()
	}()

	go api.AskAareGuru(aareGuruResponseChannel, errChannel)

	aareGuruResponse := readData(aareGuruResponseChannel, errChannel, bar)

	output_simple.PrintBanner()
	output_simple.PrintOutput(*aareGuruResponse)
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
			case <-errChannel:
				panic("")
			}
		}

		increaseBar(bar, i)
	}
	if aareGuruResponse == nil {
		fmt.Println("wtf")
		fmt.Println("wtf")
		fmt.Println("wtf")
		tmp := <-aareGuruResponseChannel
		aareGuruResponse = &tmp
	}
	// rest of progress bar
	for ; i < progressBarCount; i++ {
		increaseBar(bar, i)
	}
	bar.FinishPrint(Green(texts.Success_msg).String())
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
	if i < progressBarCount * 0.5 {
		time.Sleep(time.Millisecond * 4)
	} else if i < progressBarCount * 0.7 {
			time.Sleep(time.Millisecond * 6)
	} else {
		time.Sleep(time.Millisecond * 7)
	}
}

