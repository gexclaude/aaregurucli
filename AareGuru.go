package main

import (
	"fmt"
	"time"
    "math/rand"
	"./asciiart"
	"./api"
	"./texts"
	. "github.com/logrusorgru/aurora"
	"os/exec"
	"os"
	"runtime"
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
	
	CallClear()
	fmt.Print(BgBlue(Gray(asciiart.Banner)))
	fmt.Println()
	fmt.Println()
	printOutput(api.AskAareGuru())
}

func printOutput(aareGuruResponse api.AareGuruResponse) {
	t := time.Unix(aareGuruResponse.Aare.Timestamp, 0)
	glasses := liter_to_glasses(m3_to_liter(aareGuruResponse.Aare.Flow))

	fmt.Println(box_horizontal_line())
	fmt.Println(Bold(box(fmt.Sprintf("%s (%02d:%02d - %02d.%02d.%04d)", texts.Current_title, t.Hour(), t.Minute(), t.Day(), t.Month(), t.Year()))))
	fmt.Println(box_horizontal_line())
	fmt.Println(Bold(box(texts.Water_temperature_label)))
	fmt.Println(box(fmt.Sprintf("%5.1f %-4s - %s", aareGuruResponse.Aare.Temperature, texts.Degree_celsius_label, aareGuruResponse.Aare.Temperature_text)))
	fmt.Println(box(""))
	fmt.Println(Bold(box(texts.Water_flow_label)))
	fmt.Println(box(fmt.Sprintf("%5.0f %4s - %s (%.0f %s)", aareGuruResponse.Aare.Flow, texts.Cubic_metre_per_second_label, aareGuruResponse.Aare.Flow_text, glasses, random_flow_in_glasses_text())))
	fmt.Println(box_horizontal_line())
	fmt.Println()
	fmt.Println(box_horizontal_line())
	fmt.Println(Bold(box(texts.Forecast2h_title)))
	fmt.Println(box_horizontal_line())
	fmt.Println(Bold(box(texts.Water_temperature_label)))
	fmt.Println(box(fmt.Sprintf("%5.1f %-4s - %s", aareGuruResponse.Aare.Forecast2h, texts.Degree_celsius_label, aareGuruResponse.Aare.Forecast2h_text)))
	fmt.Println(box_horizontal_line())
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


// source: https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = clear["linux"]
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok { //if we defined a clear func for that platform:
		value()  //we execute it
	} else { //unsupported platform
		fmt.Println(runtime.GOOS)
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
