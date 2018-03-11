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

const Footer = " https://aare.guru "

const Loading_msg = "Louding "
const Success_msg = "Söögsess!"

const Error_msg = "Dr aare.guru isch verärgeret. Är git üs kei antwort meh."
const Error_Detail_msg = "Nachfolgends isch dr detailliert Fehler: "
const Error_Hints_msg = `* Bisch auefaus hingeremne proxy? Probiere mau dr Proxy mit --proxy http://di-leidig-proxy:8080 adsgä
* Oder hesch no e auti Version? Due mau aktualisiere
`

var WeatherSympolTexts map[int16]string = createWeatherSymbolTextMapping()
func createWeatherSymbolTextMapping() map[int16]string {
	var m map[int16]string = make(map[int16]string)
	m[0] = "n/a"
	m[1] = "sunnig"
	m[2] = "zimlech sunnig"
	m[3] = "bewöukt"
	m[4] = "starch bewöukt"
	m[5] = "wärmegwitter"
	m[6] = "starche räge"
	m[7] = "schneefau"
	m[8] = "näbu"
	m[9] = "schneeräge"
	m[10] = "e gutsch"
	m[11] = "liechts rägeli"
	m[12] = "es flöcklet e chli"
	m[13] = "frontgwitter"
	m[14] = "hochnäbu"
	m[15] = "es schneerägeli"
	return m
}

// CLI help

const CLI_description = `Mit däm aare guru CLI Wärchzüg chasch ganz komod d Aare-Tämperatur, -Wassermängi u ds aktuelle bärner Wätter i dire Befählszyle abfragä - u das i gwaneter aare.guru Qualität.

Obenuse, nid?

https://aare.guru`

const CLI_proxy_description = "Proxy URL - bspw. http://di-leidig-proxy:8080"
const CLI_colorless_description = "Faarbloosi Usgaab"