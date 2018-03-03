package model

type Aare struct {
	Timestamp int64
	Timestring string
	Temperature float32
	Temperature_text string
	Flow float32
	Flow_text string
	Forecast2h float32
	Forecast2h_text string
}

type AareGuruResponse struct {
	Status string
	Aare   Aare
}
