package phe

import (
	"fmt"
	"testing"
)

// TestGenerateKeys...
func TestGenerateKeys(t *testing.T) {
	pk, sk := GenerateKeys(128)
	// PK
	fmt.Println("pk.k1 = ", pk.k1.ToString())
	fmt.Println("pk.k2 = ", pk.k2.ToString())
	fmt.Println("pk.g = ", pk.g.String())
	// SK
	fmt.Println("sk.b = ", sk.b)
	fmt.Println("sk.q = ", sk.q)
}
