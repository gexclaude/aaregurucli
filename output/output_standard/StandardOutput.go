package output_standard

import (
	"fmt"
	"strconv"
	"time"
	"../../asciiart"
	"../../api"
	"../../texts"
	. "../../console"
	. "../output_common"
	"github.com/gosuri/uiprogress"
	"sync"
)

const progressBarCount = 100

var bar *uiprogress.Bar
var progressbar = true

func Init(withProgressBar bool) {
	progressbar = withProgressBar
	bar = createBar()
}

func RenderAareGuruResponse(aareGuruResponseChannel chan api.AareGuruResponse, errChannel chan string, wg *sync.WaitGroup) {
	fmt.Println(boxHorizontalLine())
	fmt.Print(CBlue(asciiart.Banner))
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
	if (!progressbar) {
		return nil;
	}

	bar := uiprogress.AddBar(progressBarCount).AppendCompleted().PrependElapsed()
	bar.Width = 45
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		msg := texts.Loading_msg
		len := 9
		if b.Current() == progressBarCount {
			msg = CGreen(texts.Success_msg)
			len += colorCharsLength(CGreen(""))
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
	if (progressbar) {
		uiprogress.Stop()
		fmt.Println(boxHorizontalLine())
	}
}

func increaseBar(bar *uiprogress.Bar, i int) {
	if (progressbar) {
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

	printLastUpdateInformation(t, weather)
	printAareTemperatureAndFlow(aare)
	printNVA(weather.Today)
	fmt.Println()
	fmt.Println(CBgBlue(CGray(fmt.Sprintf(" %s ", texts.Footer))))
	fmt.Println()
}

func printLastUpdateInformation(t time.Time, weather api.Weather) {
	fmt.Println(box(fmt.Sprintf("%-13s | %02d:%02d - %02d.%02d.%04d (%s)", texts.Current_title, t.Hour(), t.Minute(), t.Day(), t.Month(), t.Year(), weather.Location)))
	fmt.Println(boxHorizontalLine())
}

func printAareTemperatureAndFlow(aare api.Aare) {
	glasses := LiterToGlasses(M3toLiter(aare.Flow))
	glasses_text := strconv.Itoa(glasses)
	if len(glasses_text) > 3 {
		// 123456 -> 123'456
		glasses_text = glasses_text[:len(glasses_text)-3] + "'" + glasses_text[len(glasses_text)-3:]
	}

	fmt.Println(box(
		fmt.Sprintf("%-13s | %s %-4s - %s",
			texts.Water_label,
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
			RandomFlowInGlassesText()),
		CBlue("")))
}

func printNVA(weatherToday api.WeatherToday) {
	fmt.Println(boxHorizontalLine())
	fmt.Println(nva_row(texts.Nva_title_1st_row, texts.Nva_morning, weatherToday.V))
	fmt.Println(nva_row(texts.Nva_title_2nd_row, texts.Nva_afternoon, weatherToday.N))
	fmt.Println(nva_row("", texts.Nva_evening, weatherToday.A))
	fmt.Println(boxHorizontalLine())
	fmt.Println(box(texts.Nva_caption))
	fmt.Println(boxHorizontalLine())
}

func nva_row(col1_text string, col2_text string, info api.WeatherInfos) string {

	col1 := col1_text
	col2 := fmt.Sprintf("%-6s: %s° / %smm / %s%%", col2_text, CRed(fmt.Sprintf("%4.1f", info.Tt)), CGreen(fmt.Sprintf("%2d", info.Rr)), CBrown(fmt.Sprintf("%2d", info.Rrisk)))
	col3 := info.Syt
	return box(fmt.Sprintf("%-13s | %s | %s", col1, col2, col3), CRed(""), CGreen(""), CBrown(""))
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
