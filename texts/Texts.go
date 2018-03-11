package texts

const Current_title = "Latest Öpdate"
const Forecast2h_title = "I öpe 2h"

const Degree_celsius_label = "C°"
const Cubic_metre_per_second_label = "m3/s"

const Flow_beer_label = "Stange Bier/Sec."
const Flow_siroop_label = "Sirupgleser/Sec"
const Water_temperature_label = "Wasser"
const Water_flow_label = "Mängi"

const Nva_title_1st_row = "Ds wätter"
const Nva_title_2nd_row = "hüt düre Tag"

const Nva_morning = "Morge"
const Nva_afternoon = "Nami"
const Nva_evening = "Abe"
const Nva_caption = "X° = Temp. / Ymm = Rägemängi / Z%% = Rägewahrschindlechkeit"

const Footer = " https://aare.guru 👍  "

const Loading_msg = "Louding "
const Success_msg = "Söögsess!"

const Error_msg = "Dr aare.guru isch verärgeret. Är git üs kei antwort meh."
const Error_Detail_msg = "Nachfolgends isch dr detailliert Fehler: "
const Error_Hints_msg = `* Bisch auefaus hingeremne proxy? Probiere mau d umgäbigsvariable 'HTTP_PROXY=di-leidig-proxy:80' ds setze
* Oder hesch no e auti Version? Due mau aktualisiere
`
var WeatherSympolTexts map[int16]string = createWeatherSymbolTextMapping()
func createWeatherSymbolTextMapping() map[int16]string {
	var m map[int16]string = make(map[int16]string)
	// TODO replace with bärndütsch
	m[0] = "n/a"
	m[1] = "sonnig"
	m[2] = "ziemlich sonnig"
	m[3] = "bewölkt"
	m[4] = "stark bewölkt"
	m[5] = "Wärmegewitter"
	m[6] = "starker Regen"
	m[7] = "Schneefall"
	m[8] = "Nebel"
	m[9] = "Schneeregen"
	m[10] = "Regenschauer"
	m[11] = "leichter Regen"
	m[12] = "Schneeschauer"
	m[13] = "Frontgewitter"
	m[14] = "Hochnebel"
	m[15] = "Schneeregenschauer"
	return m
}