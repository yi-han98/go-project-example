package benchmarkTesting

import (
	"github.com/NebulousLabs/fastrand"
	"math/rand"
)

var ServerIndex [10]int

func InitServerIndex() {
	for i := 0; i < 10; i++ {
		ServerIndex[i] = 1 + 100
	}
}

func SelectServer() int {
	return ServerIndex[rand.Intn(10)]
}

func FastSelectServer() int {
	return ServerIndex[fastrand.Intn(10)]

}
