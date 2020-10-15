package phe

import (
	"fmt"
	"testing"
)

func BenchmarkKeyGeneration(b *testing.B) {
	var lambda = [9]int64{128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768}
	for index := 0; index < len(lambda); index++ {
		b.Run(fmt.Sprintf("%s %d", "SecurityParameter ", lambda[index]), func(b *testing.B) {
			for round := 0; round < b.N; round++ {
				GenerateKeys(lambda[index])
			}
		})
	}
}
