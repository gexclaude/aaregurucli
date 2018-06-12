package outtypewrt

import (
	"../api"
	"../outcmn"
	"../texts"
	"bytes"
	"fmt"
	"strconv"
	"sync"
	"time"
)

// Init initializes the typewriter output
func Init() {
}

// RenderAareGuruResponse renders an AareGuruResponse with the current output implementation style
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

	buffer.WriteString(texts.LocationAndTimeTitle)
	buffer.WriteString("\n")
	buffer.WriteString("---\n")
	printCityLastUpdateInformation(t, aare, &buffer)

	printAareTemperatureAndFlow(aare, &buffer)
	buffer.WriteString("\n")
	buffer.WriteString("\n")
	printNVA(weather.Today, &buffer)

	typewriter([]rune(buffer.String()), true)
	fmt.Println()

	var updateBuf bytes.Buffer
	updateBuf.WriteString("\n")
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

func printCityLastUpdateInformation(t time.Time, aare api.Aare, buffer *bytes.Buffer) {
	if aare.Location != aare.LocationLong {
		buffer.WriteString(fmt.Sprintf("%-14s: %s (%s)\n", texts.CityTitle, aare.Location, aare.LocationLong))
	} else {
		buffer.WriteString(fmt.Sprintf("%-14s: %s\n", texts.CityTitle, aare.Location))
	}

	buffer.WriteString(fmt.Sprintf("%-14s: %02d:%02d - %02d.%02d.%04d\n", texts.CurrentTitle, t.Hour(), t.Minute(), t.Day(), t.Month(), t.Year()))

	buffer.WriteString("\n")
}

func printAareTemperatureAndFlow(aare api.Aare, buffer *bytes.Buffer) {
	glasses := outcmn.LiterToGlasses(outcmn.M3toLiter(aare.Flow))
	glassesText := strconv.Itoa(glasses)
	if len(glassesText) > 3 {
		// 123456 -> 123'456
		glassesText = glassesText[:len(glassesText)-3] + "'" + glassesText[len(glassesText)-3:]
	}

	buffer.WriteString(texts.WaterLabel)
	buffer.WriteString("\n")
	buffer.WriteString("---\n")
	buffer.WriteString(
		fmt.Sprintf("%-14s: %s %-3s <- %s",
			texts.WaterTemperatureLabel,
			fmt.Sprintf("%-4.1f", aare.Temperature),
			texts.DegreeCelsiusLabel,
			aare.TemperatureText))

	buffer.WriteString("\n")

	buffer.WriteString(
		fmt.Sprintf("%-14s: %s %s <- %s (%s %s)",
			texts.WaterFlowLabel,
			fmt.Sprintf("%3.0f", aare.Flow),
			texts.CubicMetrePerSecondLabel,
			aare.FlowText, glassesText,
			outcmn.RandomFlowInGlassesText()))
}

func printNVA(weatherToday api.WeatherToday, buffer *bytes.Buffer) {
	buffer.WriteString(texts.WeatherLabel)
	buffer.WriteString("\n")
	buffer.WriteString("---\n")
	buffer.WriteString(nva(texts.NvaMorning, weatherToday.V))
	buffer.WriteString("\n")
	buffer.WriteString(nva(texts.NvaAfternoon, weatherToday.N))
	buffer.WriteString("\n")
	buffer.WriteString(nva(texts.NvaEvening, weatherToday.A))
	buffer.WriteString("\n")
	buffer.WriteString("\n")
	buffer.WriteString(texts.NvaCaption)
}

func nva(col2Text string, info api.WeatherInfos) string {
	col2 := fmt.Sprintf("%-14s: %s° / %smm / %s%%", col2Text, fmt.Sprintf("%4.1f", info.Tt), fmt.Sprintf("%2.0f", info.Rr), fmt.Sprintf("%2d", info.Rrisk))
	col3 := info.Syt
	return fmt.Sprintf("%s <- %s", col2, col3)
}
