package output_simple

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"../../asciiart"
	"../../api"
	"../../texts"
	. "../../console"
)

func PrintBanner() {
	fmt.Print(CBlue(asciiart.Banner))
}

func PrintOutput(aareGuruResponse api.AareGuruResponse) {
	aare := aareGuruResponse.Aare
	weather := aareGuruResponse.Weather

	t := time.Unix(aare.Timestamp, 0)

	printLastUpdateInformation(t, weather)
	printAareTemperatureAndFlow(aare)
	printNVA(weather.Today)
	fmt.Println()
	fmt.Println(CBgBlue(CGray(texts.Footer)))
	fmt.Println()
}

func printLastUpdateInformation(t time.Time, weather api.Weather) {
	fmt.Println(box(fmt.Sprintf("%-13s | %02d:%02d - %02d.%02d.%04d (%s)", texts.Current_title, t.Hour(), t.Minute(), t.Day(), t.Month(), t.Year(), weather.Location)))
	fmt.Println(Box_horizontal_line())
}

func printAareTemperatureAndFlow(aare api.Aare) {
	glasses := liter_to_glasses(m3_to_liter(aare.Flow))
	glasses_text := strconv.Itoa(glasses)
	if len(glasses_text) > 3 {
		// 123456 -> 123'456
		glasses_text = glasses_text[:len(glasses_text)-3] + "'" + glasses_text[len(glasses_text)-3:]
	}

	fmt.Println(box(
		fmt.Sprintf("%-13s | %s %-4s - %s",
			texts.Water_temperature_label,
			CBlue(fmt.Sprintf("%5.1f", aare.Temperature)),
			texts.Degree_celsius_label,
			aare.Temperature_text),
		CBlue("")))
		
	fmt.Println(box(
		fmt.Sprintf("%-13s | %s %4s - %s (%s %s)",
			texts.Water_flow_label,
			CBlue(fmt.Sprintf("%5.0f", aare.Flow)),
			texts.Cubic_metre_per_second_label,
			aare.Flow_text, glasses_text,
			random_flow_in_glasses_text()),
		CBlue("")))
}

func printNVA(weatherToday api.WeatherToday) {
	fmt.Println(Box_horizontal_line())
	fmt.Println(nva_row(texts.Nva_title_1st_row, texts.Nva_morning, weatherToday.V))
	fmt.Println(nva_row(texts.Nva_title_2nd_row, texts.Nva_afternoon, weatherToday.N))
	fmt.Println(nva_row("", texts.Nva_evening, weatherToday.A))
	fmt.Println(Box_horizontal_line())
	fmt.Println(box(fmt.Sprintf(texts.Nva_caption)))
	fmt.Println(Box_horizontal_line())
}

func nva_row(col1_text string, col2_text string, info api.WeatherInfos) string {

	col1 := col1_text
	col2 := fmt.Sprintf("%-6s: %s° / %smm / %s%%", col2_text, CRed(fmt.Sprintf("%4.1f", info.Tt)), CGreen(fmt.Sprintf("%2d", info.Rr)), CBrown(fmt.Sprintf("%2d", info.Rrisk)))
	col3 := info.Syt
	return box(fmt.Sprintf("%-13s | %s | %s", col1, col2, col3), CRed(""), CGreen(""), CBrown(""))
}

func Box_horizontal_line() string {
	return "+------------------------------------------------------------------------+"
}

func box(str string, colorChars ...string) string {
	return fmt.Sprintf("| %-"+strconv.Itoa(70+ColorCharsLength(colorChars...))+"s |", str)
}

func ColorCharsLength(colorChars ...string) int {
	var colorCharsLen = 0
	for _, element := range colorChars {
		colorCharsLen += len(element)
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
