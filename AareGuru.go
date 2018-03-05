package main

import (
	"fmt"
	"time"
    "math/rand"
	"./asciiart"
	"./api"
	"./console"
	"./texts"
	. "github.com/logrusorgru/aurora"
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
	fmt.Print(BgBlue(Gray(asciiart.Banner)))
	fmt.Println()
	printOutput(api.AskAareGuru())
	console.BeforeExit()
}

func printOutput(aareGuruResponse api.AareGuruResponse) {
	t := time.Unix(aareGuruResponse.Aare.Timestamp, 0)
	glasses := liter_to_glasses(m3_to_liter(aareGuruResponse.Aare.Flow))

	fmt.Println(Bold(box_horizontal_line()))
	fmt.Println(Bold(box(fmt.Sprintf("%s (%02d:%02d - %02d.%02d.%04d)", texts.Current_title, t.Hour(), t.Minute(), t.Day(), t.Month(), t.Year()))))
	fmt.Println(Bold(box_horizontal_line()))
	fmt.Println(Bold(box(texts.Water_temperature_label)))
	fmt.Println(Bold(box(fmt.Sprintf("%5.1f %-4s - %s", aareGuruResponse.Aare.Temperature, texts.Degree_celsius_label, aareGuruResponse.Aare.Temperature_text))))
	fmt.Println(Bold(box("")))
	fmt.Println(Bold(box(texts.Water_flow_label)))
	fmt.Println(Bold(box(fmt.Sprintf("%5.0f %4s - %s (%.0f %s)", aareGuruResponse.Aare.Flow, texts.Cubic_metre_per_second_label, aareGuruResponse.Aare.Flow_text, glasses, random_flow_in_glasses_text()))))
	fmt.Println(Bold(box_horizontal_line()))
	fmt.Println()
	fmt.Println(Bold(box_horizontal_line()))
	fmt.Println(Bold(box(texts.Forecast2h_title)))
	fmt.Println(Bold(box_horizontal_line()))
	fmt.Println(Bold(box(texts.Water_temperature_label)))
	fmt.Println(Bold(box(fmt.Sprintf("%5.1f %-4s - %s", aareGuruResponse.Aare.Forecast2h, texts.Degree_celsius_label, aareGuruResponse.Aare.Forecast2h_text))))
	fmt.Println(Bold(box_horizontal_line()))
	fmt.Println()
	fmt.Println(BgBlue((Gray((texts.Footer)))))
	fmt.Println()
}


func box_horizontal_line() string {
	return "+--------------------------------------------------------------+"
}

func box(str string) string {
	return fmt.Sprintf("| %-60s |", str)
}

func m3_to_liter(m3 float32) float32 {
	return m3 * (10 * 10 * 10)
}

func liter_to_glasses(liter float32) float32 {
	return liter / 0.3 // 3 dl
} 

func random_flow_in_glasses_text() string {
	if(rand_bool()) {
		return texts.Flow_beer_label
	} else {
		return texts.Flow_siroop_label
	}
}

func rand_bool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Float32() < 0.5
}



