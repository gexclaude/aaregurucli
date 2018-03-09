package main

import (
	"fmt"
	"./api"
	"./console"
	"./texts"
	"./output/output_simple"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println()
			fmt.Println(texts.Error_Detail_msg, r)
			fmt.Print(texts.Error_Hints_msg)
			fmt.Print()
		}
	}()
	
	console.Init()
	console.CallClear()

	output_simple.PrintBanner()
	aareGuruResponse := api.AskAareGuru()

	output_simple.PrintOutput(aareGuruResponse)
	
	console.BeforeExit()
}




