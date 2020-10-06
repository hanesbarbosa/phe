package phe

import (
	"math/big"
	"strings"
	"testing"
)

// Tests if the reverse is calculated properly.
func TestReverse(t *testing.T) {
	c := []string{"2", "3", "4", "5", "6", "7", "8", "9"}
	m := New(c)
	q := big.NewInt(257)
	m = Reverse(m, q)

	if strings.Compare(m.E0.String(), "2") != 0 ||
		strings.Compare(m.E1.String(), "3") != 0 ||
		strings.Compare(m.E2.String(), "-4") != 0 ||
		strings.Compare(m.E3.String(), "5") != 0 ||
		strings.Compare(m.E12.String(), "6") != 0 ||
		strings.Compare(m.E13.String(), "-7") != 0 ||
		strings.Compare(m.E23.String(), "8") != 0 ||
		strings.Compare(m.E123.String(), "9") != 0 {
		t.Errorf("Wrong results for the reverse calculation.")
	}
}
