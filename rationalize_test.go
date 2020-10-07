package phe

import (
	"math/big"
	"strings"
	"testing"
)

// Tests if the rationalized multivector is calculated properly.
func TestRationalize(t *testing.T) {
	c := []string{"2", "3", "4", "5", "6", "7", "8", "9"}

	m := NewMultivector(c)
	q := big.NewInt(257)
	m = Rationalize(m, q)

	if strings.Compare(m.E0.String(), "226") != 0 ||
		strings.Compare(m.E1.String(), "0") != 0 ||
		strings.Compare(m.E2.String(), "0") != 0 ||
		strings.Compare(m.E3.String(), "0") != 0 ||
		strings.Compare(m.E12.String(), "0") != 0 ||
		strings.Compare(m.E13.String(), "0") != 0 ||
		strings.Compare(m.E23.String(), "0") != 0 ||
		strings.Compare(m.E123.String(), "0") != 0 {
		t.Errorf("Wrong results for the rationalized calculation.")
	}
}
