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
	fmt.Print(Blue(asciiart.Banner))
}

func PrintOutput(aareGuruResponse api.AareGuruResponse) {
	aare := aareGuruResponse.Aare
	weather := aareGuruResponse.Weather

	t := time.Unix(aare.Timestamp, 0)

	printLastUpdateInformation(t, weather)
	printAareTemperatureAndFlow(aare)
	printNVA(weather.Today)
	fmt.Println()
	fmt.Println(BgBlue((Gray((texts.Footer)))))
	fmt.Println()
}

func printLastUpdateInformation(t time.Time, weather api.Weather) {
	fmt.Println(box(fmt.Sprintf("%-13s | %02d:%02d - %02d.%02d.%04d (%s)", texts.Current_title, t.Hour(), t.Minute(), t.Day(), t.Month(), t.Year(), weather.Location)))
	fmt.Println(box_horizontal_line())
}

func printAareTemperatureAndFlow(aare api.Aare) {
	glasses := liter_to_glasses(m3_to_liter(aare.Flow))
	glasses_text := strconv.Itoa(glasses)
	// 123456 -> 123'456
	glasses_text = glasses_text[:len(glasses_text)-3] + "'" + glasses_text[len(glasses_text)-3:]

	fmt.Println(box(fmt.Sprintf("%-13s | %5.1f %-4s - %s", texts.Water_temperature_label, Blue(aare.Temperature), texts.Degree_celsius_label, aare.Temperature_text), Blue("")))
	fmt.Println(box(fmt.Sprintf("%-13s | %5.0f %4s - %s (%s %s)", texts.Water_flow_label, Blue(aare.Flow), texts.Cubic_metre_per_second_label, aare.Flow_text, glasses_text, random_flow_in_glasses_text()), Blue("")))
}

func printNVA(weatherToday api.WeatherToday) {
	fmt.Println(box_horizontal_line())
	fmt.Println(nva_row(texts.Nva_title_1st_row, texts.Nva_morning, weatherToday.V))
	fmt.Println(nva_row(texts.Nva_title_2nd_row, texts.Nva_afternoon, weatherToday.N))
	fmt.Println(nva_row("", texts.Nva_evening, weatherToday.A))
	fmt.Println(box_horizontal_line())
	fmt.Println(box(fmt.Sprintf(texts.Nva_caption)))
	fmt.Println(box_horizontal_line())
}

func nva_row(col1_text string, col2_text string, info api.WeatherInfos) string {

	col1 := col1_text
	col2 := fmt.Sprintf("%-6s: %4.1f° / %2dmm / %2d%%", col2_text, Red(info.Tt), Green(info.Rr), Gray(info.Rrisk))
	col3 := texts.WeatherSympolTexts[info.Symt]
	return box(fmt.Sprintf("%-13s | %s | %s", col1, col2, col3), Red(""), Green(""), Gray(""))
}

func box_horizontal_line() string {
	return "+------------------------------------------------------------------------+"
}

func box(str string, colorChars ...fmt.Stringer) string {
	return fmt.Sprintf("| %-"+strconv.Itoa(70+colorCharsLength(colorChars...))+"s |", str)
}

func colorCharsLength(colorChars ...fmt.Stringer) int {
	var colorCharsLen = 0
	for _, element := range colorChars {
		colorCharsLen += len(element.String())
	}
	return colorCharsLen
}

func m3_to_liter(m3 float32) float32 {
	return m3 * (10 * 10 * 10)
}

func liter_to_glasses(liter float32) int {
	return int(liter / 0.3) // 3 dl
}

func random_flow_in_glasses_text() string {
	if (rand_bool()) {
		return texts.Flow_beer_label
	} else {
		return texts.Flow_siroop_label
	}
}

func rand_bool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Float32() < 0.5
}
