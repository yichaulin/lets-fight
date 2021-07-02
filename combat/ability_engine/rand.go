package ability_engine

import (
	"math/rand"
	"time"
)

var randSource rand.Source

func init() {
	seed := time.Now().Unix()
	randSource = rand.NewSource(seed)
}

func getIntn(n int) int {
	random := rand.New(randSource)

	return random.Intn(n)
}
