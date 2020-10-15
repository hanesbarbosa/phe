package phe

import (
	"fmt"
	"testing"
)

var lambda = [8]int64{64, 128, 256, 512, 1024, 2048, 4096, 8192}

func BenchmarkKeyGeneration(b *testing.B) {
	for index := 0; index < len(lambda); index++ {
		b.Run(fmt.Sprintf("%s %d", "SecurityParameter ", lambda[index]), func(b *testing.B) {
			for round := 0; round < b.N; round++ {
				GenerateKeys(lambda[index])
			}
		})
	}
}
