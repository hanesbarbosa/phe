package phe

import (
	"strings"
	"testing"
)

// Tests if the clifford conjugation is calculated properly.
func TestCliffordConjugation(t *testing.T) {
	c := []string{"40", "-29", "29", "-25", "33", "-29", "32", "-28"}

	m := New(c)
	m = CliffordConjugation(m)

	if strings.Compare(m.E0.String(), "40") != 0 ||
		strings.Compare(m.E1.String(), "29") != 0 ||
		strings.Compare(m.E2.String(), "-29") != 0 ||
		strings.Compare(m.E3.String(), "25") != 0 ||
		strings.Compare(m.E12.String(), "-33") != 0 ||
		strings.Compare(m.E13.String(), "29") != 0 ||
		strings.Compare(m.E23.String(), "-32") != 0 ||
		strings.Compare(m.E123.String(), "-28") != 0 {
		t.Errorf("Wrong results for the clifford conjugation calculation.")
	}
}
