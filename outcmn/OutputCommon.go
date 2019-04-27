package outcmn

import (
	"github.com/toasterson/aaregurucli/texts"
	"math/rand"
	"time"
)

// M3toLiter converts cubic meter to liter
func M3toLiter(m3 float32) float32 {
	return m3 * (10 * 10 * 10)
}

// LiterToGlasses calculates nr of glasses of 3dl for a given liter count
func LiterToGlasses(liter float32) int {
	return int(liter / 0.3) // 3 dl
}

// RandomFlowInGlassesText either returns glasses of beer per second or glasses of siroop per second. decision is random
func RandomFlowInGlassesText() string {
	if randomBool() {
		return texts.FlowBeerLabel
	}
	return texts.FlowSiroopLabel
}

func randomBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Float32() < 0.5
}
