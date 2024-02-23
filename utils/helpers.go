package helpers

import (
	"math/rand"
	// "time"
)

func GetRandomNumber(maxi int) int {

	// rand.NewSource(time.Now().UnixNano())

	return rand.Intn(maxi)
}
