package output_simple

import (
	"fmt"
	"time"
	"math/rand"
	"../../asciiart"
	"../../api"
	"../../texts"
	. "github.com/logrusorgru/aurora"
	"strconv"
)

func PrintBanner() {
	fmt.Print(BgBlue(Gray(asciiart.Banner)))
	fmt.Println()
}

func PrintOutput(aareGuruResponse api.AareGuruResponse) {
	aare := aareGuruResponse.Aare
	weather := aareGuruResponse.Weather
	weatherToday := weather.Today

	t := time.Unix(aare.Timestamp, 0)
	glasses := liter_to_glasses(m3_to_liter(aare.Flow))

	printCurrentWeather(t, aare, glasses)
	fmt.Println()
	printForecastWeather(aare)
	fmt.Println()
	printNVA(weatherToday)
	fmt.Println()
	fmt.Println(BgBlue((Gray((texts.Footer)))))
	fmt.Println()
}

func printCurrentWeather(t time.Time, aare api.Aare, glasses float32) {
	fmt.Println(Bold(box_horizontal_line()))
	fmt.Println(Bold(box(fmt.Sprintf("%s (%02d:%02d - %02d.%02d.%04d)", texts.Current_title, t.Hour(), t.Minute(), t.Day(), t.Month(), t.Year()))))
	fmt.Println(Bold(box_horizontal_line()))
	fmt.Println(Bold(box(texts.Water_temperature_label)))
	fmt.Println(Bold(box(fmt.Sprintf("%5.1f %-4s - %s", aare.Temperature, texts.Degree_celsius_label, aare.Temperature_text))))
	fmt.Println(Bold(box("")))
	fmt.Println(Bold(box(texts.Water_flow_label)))
	fmt.Println(Bold(box(fmt.Sprintf("%5.0f %4s - %s (%.0f %s)", aare.Flow, texts.Cubic_metre_per_second_label, aare.Flow_text, glasses, random_flow_in_glasses_text()))))
	fmt.Println(Bold(box_horizontal_line()))
}

func printForecastWeather(aare api.Aare) {
	fmt.Println(Bold(box_horizontal_line()))
	fmt.Println(Bold(box(texts.Forecast2h_title)))
	fmt.Println(Bold(box_horizontal_line()))
	fmt.Println(Bold(box(texts.Water_temperature_label)))
	fmt.Println(Bold(box(fmt.Sprintf("%5.1f %-4s - %s", aare.Forecast2h, texts.Degree_celsius_label, aare.Forecast2h_text))))
	fmt.Println(Bold(box_horizontal_line()))
}

func printNVA(weatherToday api.WeatherToday) {
	fmt.Println(Bold(box_horizontal_line()))
	fmt.Println(Bold(nva_row(texts.Nva_title_1st_row, texts.Nva_morning, weatherToday.V)))
	fmt.Println(Bold(nva_row(texts.Nva_title_2nd_row, texts.Nva_afternoon, weatherToday.N)))
	fmt.Println(Bold(nva_row("", texts.Nva_evening, weatherToday.A)))
	fmt.Println(Bold(box_horizontal_line()))
	fmt.Println(Bold(box(fmt.Sprintf(texts.Nva_caption))))
	fmt.Println(Bold(box_horizontal_line()))
}

func nva_row(col1_text string, col2_text string, info api.WeatherInfos) string {
	var bla string

	// TODO replace with bärndütsch
	switch info.Symt {
	case 1:
		bla = "sonnig"
	case 2:
		bla = "ziemlich sonnig"
	case 3:
		bla = "bewölkt"
	case 4:
		bla = "stark bewölkt"
	case 5:
		bla = "Wärmegewitter"
	case 6:
		bla = "starker Regen"
	case 7:
		bla = "Schneefall"
	case 8:
		bla = "Nebel"
	case 9:
		bla = "Schneeregen"
	case 10:
		bla = "Regenschauer"
	case 11:
		bla = "leichter Regen"
	case 12:
		bla = "Schneeschauer"
	case 13:
		bla = "Frontgewitter"
	case 14:
		bla = "Hochnebel"
	case 15:
		bla = "Schneeregenschauer"
	}

	col1 := col1_text
	col2 := pad(fmt.Sprintf("%-6s: %4.1f° / %2dmm / %2d%%", col2_text, info.Tt, info.Rr, info.Rrisk), 20)
	col3 := bla
	return box(fmt.Sprintf("%-13s | %s | %s", col1, col2, col3))
}

func box_horizontal_line() string {
	return "+------------------------------------------------------------------------+"
}

func box(str string) string {
	return fmt.Sprintf("| %-70s |", str)
}

func pad(str string, len int) string {
	return fmt.Sprintf(" %-" + strconv.Itoa(len) + "s ", str)
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
