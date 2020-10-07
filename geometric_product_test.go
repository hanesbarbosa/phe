package phe

import (
	"math/big"
	"strings"
	"testing"
)

// Tests if the geometric product between two multivectors (plaintext and key) are calculated properly.
func TestGeometricProduct(t *testing.T) {
	m0 := []string{"2", "3", "4", "5", "6", "7", "8", "9"}
	s0 := []string{"2", "3", "4", "5", "6", "7", "8", "9"}

	m1 := NewMultivector(m0)
	s1 := NewMultivector(s0)
	q := big.NewInt(257)

	m := GeometricProduct(m1, s1, q)

	if strings.Compare(m.E0.String(), "81") != 0 ||
		strings.Compare(m.E1.String(), "125") != 0 ||
		strings.Compare(m.E2.String(), "142") != 0 ||
		strings.Compare(m.E3.String(), "169") != 0 ||
		strings.Compare(m.E12.String(), "114") != 0 ||
		strings.Compare(m.E13.String(), "213") != 0 ||
		strings.Compare(m.E23.String(), "86") != 0 ||
		strings.Compare(m.E123.String(), "88") != 0 {
		t.Errorf("Wrong results for the geometric product calculation.")
	}
}
