package phe

import (
	"math/big"
	"strings"
	"testing"
)

// Tests if the scalar multiplication is calculated properly.
func TestModularMultivector(t *testing.T) {
	c := []string{"40", "29", "29", "25", "33", "9", "32", "28"}

	m := NewMultivector(c)

	q := new(big.Int)
	q.SetString("20", 10)

	m = ModularMultivector(m, q)

	if strings.Compare(m.E0.String(), "0") != 0 ||
		strings.Compare(m.E1.String(), "9") != 0 ||
		strings.Compare(m.E2.String(), "9") != 0 ||
		strings.Compare(m.E3.String(), "5") != 0 ||
		strings.Compare(m.E12.String(), "13") != 0 ||
		strings.Compare(m.E13.String(), "9") != 0 ||
		strings.Compare(m.E23.String(), "12") != 0 ||
		strings.Compare(m.E123.String(), "8") != 0 {
		t.Errorf("Wrong results for the calculation of the modulus for each coefficient.")
	}
}
