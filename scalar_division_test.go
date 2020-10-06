package phe

import (
	"math/big"
	"strings"
	"testing"
)

// Tests if the scalar division is calculated properly.
func TestScalarDivision(t *testing.T) {
	c := []string{"2", "3", "4", "5", "6", "7", "8", "9"}

	m := New(c)
	q := big.NewInt(257)
	s := big.NewInt(12)
	m = ScalarDivision(m, s, q)

	if strings.Compare(m.E0.String(), "43") != 0 ||
		strings.Compare(m.E1.String(), "193") != 0 ||
		strings.Compare(m.E2.String(), "86") != 0 ||
		strings.Compare(m.E3.String(), "236") != 0 ||
		strings.Compare(m.E12.String(), "129") != 0 ||
		strings.Compare(m.E13.String(), "22") != 0 ||
		strings.Compare(m.E23.String(), "172") != 0 ||
		strings.Compare(m.E123.String(), "65") != 0 {
		t.Errorf("Wrong results for the rational form of coefficients.")
	}
}
