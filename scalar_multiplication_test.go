package phe

import (
	"math/big"
	"strings"
	"testing"
)

// Tests if the scalar multiplication is calculated properly.
func TestScalarMultiplication(t *testing.T) {
	c := []string{"2", "3", "4", "5", "6", "7", "8", "9"}

	m := NewMultivector(c)
	q := big.NewInt(257)
	s := big.NewInt(12)
	m = ScalarMultiplication(m, s, q)

	if strings.Compare(m.E0.String(), "24") != 0 ||
		strings.Compare(m.E1.String(), "36") != 0 ||
		strings.Compare(m.E2.String(), "48") != 0 ||
		strings.Compare(m.E3.String(), "60") != 0 ||
		strings.Compare(m.E12.String(), "72") != 0 ||
		strings.Compare(m.E13.String(), "84") != 0 ||
		strings.Compare(m.E23.String(), "96") != 0 ||
		strings.Compare(m.E123.String(), "108") != 0 {
		t.Errorf("Wrong results for the scalar multiplication calculation.")
	}
}
