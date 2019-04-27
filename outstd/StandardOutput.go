package outstd

import (
	"fmt"
	"github.com/gexclaude/aaregurucli/api"
	"github.com/gexclaude/aaregurucli/asciiart"
	"github.com/gexclaude/aaregurucli/console"
	"github.com/gexclaude/aaregurucli/outcmn"
	"github.com/gexclaude/aaregurucli/texts"
	"github.com/gosuri/uiprogress"
	"strconv"
	"sync"
	"time"
)

const progressBarCount = 100

var bar *uiprogress.Bar
var progressbar = true

// Init initializes the typewriter output
func Init(withProgressBar bool) {
	progressbar = withProgressBar
	bar = createBar()
}

// RenderAareGuruResponse renders an AareGuruResponse with the current output implementation style
func RenderAareGuruResponse(aareGuruResponseChannel chan api.AareGuruResponse, errChannel chan string, wg *sync.WaitGroup) {
	fmt.Println(boxHorizontalLine())
	fmt.Print(console.CBlue(asciiart.Banner))
	fmt.Println(boxHorizontalLine())

	defer func() {
		if r := recover(); r != nil {
			uiprogress.Stop()
			panic(r)
		}
	}()

	aareGuruResponse := readData(aareGuruResponseChannel, errChannel, wg)
	printOutput(*aareGuruResponse)
}

func readData(aareGuruResponseChannel chan api.AareGuruResponse, errChannel chan string, wg *sync.WaitGroup) *api.AareGuruResponse {
	var aareGuruResponse *api.AareGuruResponse
	i := 0
	// only first part of progress bar
	for ; progressbar && i < progressBarCount-int(progressBarCount*.75); i++ {
		if aareGuruResponse == nil {
			select {
			case tmp := <-aareGuruResponseChannel:
				aareGuruResponse = &tmp
			case err := <-errChannel:
				panic(err)
			}
		}

		increaseBar(bar, i)
	}
	if aareGuruResponse == nil {
		tmp := <-aareGuruResponseChannel
		aareGuruResponse = &tmp
	}
	// rest of progress bar
	for ; progressbar && i < progressBarCount; i++ {
		increaseBar(bar, i)
	}
	wg.Wait()
	stopBar()

	return aareGuruResponse
}

func createBar() *uiprogress.Bar {
	if !progressbar {
		return nil
	}

	bar := uiprogress.AddBar(progressBarCount).AppendCompleted().PrependElapsed()
	bar.Width = 45
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		msg := texts.LoadingMsg
		len := 9
		if b.Current() == progressBarCount {
			msg = console.CGreen(texts.SuccessMsg)
			len += colorCharsLength(console.CGreen(""))
		}
		return fmt.Sprintf("| %-"+strconv.Itoa(len)+"s %3d", msg, b.Current())
	})
	bar.AppendFunc(func(b *uiprogress.Bar) string {
		return fmt.Sprintf("|")
	})
	uiprogress.Start()
	return bar
}

func stopBar() {
	if progressbar {
		uiprogress.Stop()
		fmt.Println(boxHorizontalLine())
	}
}

func increaseBar(bar *uiprogress.Bar, i int) {
	if progressbar {
		bar.Incr()
		if i < progressBarCount*0.5 {
			time.Sleep(time.Millisecond * 2)
		} else if i < progressBarCount*0.7 {
			time.Sleep(time.Millisecond * 3)
		} else {
			time.Sleep(time.Millisecond * 4)
		}
	}
}

func printOutput(aareGuruResponse api.AareGuruResponse) {
	aare := aareGuruResponse.Aare
	weather := aareGuruResponse.Weather

	t := time.Unix(aare.Timestamp, 0)
	printCityAndLastUpdate(t, aare)
	printAareTemperatureAndFlow(aare)
	printNVA(weather.Today)
	fmt.Println()
	fmt.Println(console.CBgBlue(console.CGray(fmt.Sprintf(" %s ", texts.Footer))))
	fmt.Println()
}

func printCityAndLastUpdate(t time.Time, aare api.Aare) {
	if aare.Location != aare.LocationLong {
		fmt.Println(box(fmt.Sprintf("%-13s | %s (%s)", texts.CityTitle, aare.Location, aare.LocationLong)))
	} else {
		fmt.Println(box(fmt.Sprintf("%-13s | %s", texts.CityTitle, aare.Location)))
	}
	fmt.Println(box(fmt.Sprintf("%-13s | %02d:%02d - %02d.%02d.%04d", texts.CurrentTitle, t.Hour(), t.Minute(), t.Day(), t.Month(), t.Year())))
	fmt.Println(boxHorizontalLine())
}

func printAareTemperatureAndFlow(aare api.Aare) {
	glasses := outcmn.LiterToGlasses(outcmn.M3toLiter(aare.Flow))
	glassesText := strconv.Itoa(glasses)
	if len(glassesText) > 3 {
		// 123456 -> 123'456
		glassesText = glassesText[:len(glassesText)-3] + "'" + glassesText[len(glassesText)-3:]
	}

	fmt.Println(box(
		fmt.Sprintf("%-13s | %s %-4s - %s",
			texts.WaterLabel,
			console.CBlue(fmt.Sprintf("%5.1f", aare.Temperature)),
			texts.DegreeCelsiusLabel,
			aare.TemperatureText),
		console.CBlue("")))

	fmt.Println(box(
		fmt.Sprintf("%-13s | %s %4s - %s (%s %s)",
			texts.WaterFlowLabel,
			console.CBlue(fmt.Sprintf("%5.0f", aare.Flow)),
			texts.CubicMetrePerSecondLabel,
			aare.FlowText, glassesText,
			outcmn.RandomFlowInGlassesText()),
		console.CBlue("")))
}

func printNVA(weatherToday api.WeatherToday) {
	fmt.Println(boxHorizontalLine())
	fmt.Println(nvaRow(texts.NvaTitle1stRow, texts.NvaMorning, weatherToday.V))
	fmt.Println(nvaRow(texts.NvaTitle2ndRow, texts.NvaAfternoon, weatherToday.N))
	fmt.Println(nvaRow("", texts.NvaEvening, weatherToday.A))
	fmt.Println(boxHorizontalLine())
	fmt.Println(box(texts.NvaCaption))
	fmt.Println(boxHorizontalLine())
}

func nvaRow(col1Text string, col2Text string, info api.WeatherInfos) string {

	col1 := col1Text
	col2 := fmt.Sprintf("%-6s: %s° / %smm / %s%%", col2Text, console.CRed(fmt.Sprintf("%4.1f", info.Tt)), console.CGreen(fmt.Sprintf("%2.0f", info.Rr)), console.CBrown(fmt.Sprintf("%2d", info.Rrisk)))
	col3 := info.Syt
	return box(fmt.Sprintf("%-13s | %s | %s", col1, col2, col3), console.CRed(""), console.CGreen(""), console.CBrown(""))
}

func boxHorizontalLine() string {
	return "+------------------------------------------------------------------------+"
}

func box(str string, colorChars ...string) string {
	return fmt.Sprintf("| %-"+strconv.Itoa(70+colorCharsLength(colorChars...))+"s |", str)
}

func colorCharsLength(colorChars ...string) int {
	var colorCharsLen = 0
	for _, element := range colorChars {
		colorCharsLen += len(element)
	}
	return colorCharsLen
}
