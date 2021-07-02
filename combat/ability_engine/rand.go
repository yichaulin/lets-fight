package ability_engine

import (
	"fmt"
	"math/rand"
	"time"
)

var randSource rand.Source

func init() {
	seed := time.Now().Unix()
	randSource = rand.NewSource(seed)
}

func getIntn(n int) int {
	if n <= 0 {
		fmt.Println("input of getIntn is <= 0")
	}

	random := rand.New(randSource)
	return random.Intn(n)
}
