package phe

import (
	"math/big"
	"strings"
	"testing"
)

// Tests if the clifford conjugation is calculated properly.
func TestCliffordConjugation(t *testing.T) {
	c := []string{"2", "3", "4", "5", "6", "7", "8", "9"}

	m := New(c)
	q := big.NewInt(257)
	m = CliffordConjugation(m, q)

	if strings.Compare(m.E0.String(), "2") != 0 ||
		strings.Compare(m.E1.String(), "254") != 0 ||
		strings.Compare(m.E2.String(), "253") != 0 ||
		strings.Compare(m.E3.String(), "252") != 0 ||
		strings.Compare(m.E12.String(), "251") != 0 ||
		strings.Compare(m.E13.String(), "250") != 0 ||
		strings.Compare(m.E23.String(), "249") != 0 ||
		strings.Compare(m.E123.String(), "9") != 0 {
		t.Errorf("Wrong results for the clifford conjugation calculation.")
	}
}
