package output_typewriter

import (
	"fmt"
	"strconv"
	"time"
	"../../api"
	"../../texts"
	. "../output_common"
	"sync"
	"bytes"
)

func Init() {
}

func RenderAareGuruResponse(aareGuruResponseChannel chan api.AareGuruResponse, errChannel chan string, wg *sync.WaitGroup) {
	aareGuruResponse := readData(aareGuruResponseChannel, errChannel, wg)
	printOutput(*aareGuruResponse)
}

func readData(aareGuruResponseChannel chan api.AareGuruResponse, errChannel chan string, wg *sync.WaitGroup) *api.AareGuruResponse {
	var aareGuruResponse *api.AareGuruResponse

	select {
	case tmp := <-aareGuruResponseChannel:
		aareGuruResponse = &tmp
	case err := <-errChannel:
		panic(err)
	}

	wg.Wait()

	return aareGuruResponse
}

func printOutput(aareGuruResponse api.AareGuruResponse) {
	aare := aareGuruResponse.Aare
	weather := aareGuruResponse.Weather

	t := time.Unix(aare.Timestamp, 0)

	var buffer bytes.Buffer

	printAareTemperatureAndFlow(aare, &buffer)
	buffer.WriteString("\n")
	buffer.WriteString("\n")
	printNVA(weather.Today, &buffer)
	buffer.WriteString("\n")
	buffer.WriteString("\n")

	typewriter([]rune(buffer.String()), true)
	fmt.Println()

	var updateBuf bytes.Buffer
	printLastUpdateInformation(t, weather, &updateBuf)
	updateBuf.WriteString(texts.Footer)
	typewriter([]rune(updateBuf.String()), false)
	fmt.Println()
}
func typewriter(converted []rune, specialCharSleep bool) {
	for _, c := range converted {
		chr := string(c)
		fmt.Print(chr)
		if specialCharSleep && chr == "\n" {
			time.Sleep(time.Millisecond * 50)
		}
		if specialCharSleep && chr == ":" {
			time.Sleep(time.Millisecond * 200)
		}
		time.Sleep(time.Millisecond * 15)
	}
}

func printLastUpdateInformation(t time.Time, weather api.Weather, buffer *bytes.Buffer) {
	buffer.WriteString(fmt.Sprintf("%s | %02d:%02d - %02d.%02d.%04d (%s)\n", texts.Current_title, t.Hour(), t.Minute(), t.Day(), t.Month(), t.Year(), weather.Location))
}

func printAareTemperatureAndFlow(aare api.Aare, buffer *bytes.Buffer) {
	glasses := LiterToGlasses(M3toLiter(aare.Flow))
	glasses_text := strconv.Itoa(glasses)
	if len(glasses_text) > 3 {
		// 123456 -> 123'456
		glasses_text = glasses_text[:len(glasses_text)-3] + "'" + glasses_text[len(glasses_text)-3:]
	}

	buffer.WriteString(texts.Water_label)
	buffer.WriteString("\n")
	buffer.WriteString("---\n")
	buffer.WriteString(
		fmt.Sprintf("%-11s: %s %-4s <- %s",
			texts.Water_temperature_label,
			fmt.Sprintf("%.1f", aare.Temperature),
			texts.Degree_celsius_label,
			aare.Temperature_text))

	buffer.WriteString("\n")

	buffer.WriteString(
		fmt.Sprintf("%-11s: %s %s <- %s (%s %s)",
			texts.Water_flow_label,
			fmt.Sprintf("%3.0f", aare.Flow),
			texts.Cubic_metre_per_second_label,
			aare.Flow_text, glasses_text,
			RandomFlowInGlassesText()))
}

func printNVA(weatherToday api.WeatherToday, buffer *bytes.Buffer) {
	buffer.WriteString(texts.Weather_label)
	buffer.WriteString("\n")
	buffer.WriteString("---\n")
	buffer.WriteString(nva(texts.Nva_morning, weatherToday.V))
	buffer.WriteString("\n")
	buffer.WriteString(nva(texts.Nva_afternoon, weatherToday.N))
	buffer.WriteString("\n")
	buffer.WriteString(nva(texts.Nva_evening, weatherToday.A))
	buffer.WriteString("\n")
	buffer.WriteString("\n")
	buffer.WriteString(texts.Nva_caption)
}

func nva(col2_text string, info api.WeatherInfos) string {
	col2 := fmt.Sprintf("%-7s: %s° / %smm / %s%%", col2_text, fmt.Sprintf("%4.1f", info.Tt), fmt.Sprintf("%2d", info.Rr), fmt.Sprintf("%2d", info.Rrisk))
	col3 := info.Syt
	return fmt.Sprintf("%s <- %s", col2, col3)
}
