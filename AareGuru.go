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

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println()
			fmt.Println(Red(texts.Error_Detail_msg), r)
			fmt.Print(texts.Error_Hints_msg)
			fmt.Print()
		}
	}()
	
	console.Init()
	console.CallClear()

	count := 700
	bar := pb.New(count).Prefix(texts.Loading_msg)
	bar.SetWidth(74)
	bar.SetMaxWidth(74)
	bar.ShowCounters = false
	bar.ShowElapsedTime = false
	bar.ShowFinalTime = false
	bar.Start()

	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.FinishPrint(Green(texts.Success_msg).String())

	fmt.Println()
	aareGuruResponse := api.AskAareGuru()

	output_simple.PrintBanner()
	output_simple.PrintOutput(aareGuruResponse)
	
	console.BeforeExit()
}




