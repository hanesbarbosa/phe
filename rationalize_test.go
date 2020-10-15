package phe

import (
	"math/big"
	"strings"
	"testing"
)

// Tests if the rationalized multivector is calculated properly.
func TestRationalize(t *testing.T) {
	q := big.NewInt(257)

	m1 := NewMultivector([]string{"2", "3", "4", "5", "6", "7", "8", "9"})
	m2 := NewMultivector([]string{"214", "142", "106", "107", "8", "130", "137", "116"})

	m1 = Rationalize(m1, q)
	m2 = Rationalize(m2, q)

	// Multivector has inverse
	if strings.Compare(m1.E0.String(), "226") != 0 ||
		strings.Compare(m1.E1.String(), "0") != 0 ||
		strings.Compare(m1.E2.String(), "0") != 0 ||
		strings.Compare(m1.E3.String(), "0") != 0 ||
		strings.Compare(m1.E12.String(), "0") != 0 ||
		strings.Compare(m1.E13.String(), "0") != 0 ||
		strings.Compare(m1.E23.String(), "0") != 0 ||
		strings.Compare(m1.E123.String(), "0") != 0 {
		t.Errorf("Wrong results for the rationalized calculation.")
	}

	// Multivector has no inverse
	if strings.Compare(m2.E0.String(), "0") != 0 ||
		strings.Compare(m2.E1.String(), "0") != 0 ||
		strings.Compare(m2.E2.String(), "0") != 0 ||
		strings.Compare(m2.E3.String(), "0") != 0 ||
		strings.Compare(m2.E12.String(), "0") != 0 ||
		strings.Compare(m2.E13.String(), "0") != 0 ||
		strings.Compare(m2.E23.String(), "0") != 0 ||
		strings.Compare(m2.E123.String(), "0") != 0 {
		t.Errorf("Wrong results for the rationalized calculation. Expected zeroed multivector.")
	}
}
