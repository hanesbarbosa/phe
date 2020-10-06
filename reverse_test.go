package phe

import (
	"strings"
	"testing"
)

// Tests if the reverse is calculated properly.
func TestReverse(t *testing.T) {
	c := []string{"40", "-29", "29", "-25", "33", "-29", "32", "-28"}

	m := New(c)
	m = Reverse(m)

	// 40e0 + 29e1 + -29e2 + 25e3 + 33e12 + -29e13 + 32e23 + 28e123
	if strings.Compare(m.E0.String(), "40") != 0 ||
		strings.Compare(m.E1.String(), "29") != 0 ||
		strings.Compare(m.E2.String(), "-29") != 0 ||
		strings.Compare(m.E3.String(), "25") != 0 ||
		strings.Compare(m.E12.String(), "33") != 0 ||
		strings.Compare(m.E13.String(), "-29") != 0 ||
		strings.Compare(m.E23.String(), "32") != 0 ||
		strings.Compare(m.E123.String(), "28") != 0 {
		t.Errorf("Wrong results for the reverse calculation.")
	}
}
