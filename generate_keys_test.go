package phe

import (
	"testing"
)

// TestGenerateKeys...
func TestGenerateKeys(t *testing.T) {
	GenerateKeys(128)
	// for i := 0; i < 3; i++ {
	// 	result, _ := rand.Int(rand.Reader, big.NewInt(100))
	// 	fmt.Println(result)
	// }
	//t.Errorf("Wrong results for the scalar multiplication calculation.")
}
