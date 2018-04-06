package output_common

import (
	"time"
	"math/rand"
	"../../texts"
)

func M3toLiter(m3 float32) float32 {
	return m3 * (10 * 10 * 10)
}

func LiterToGlasses(liter float32) int {
	return int(liter / 0.3) // 3 dl
}

func RandomFlowInGlassesText() string {
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