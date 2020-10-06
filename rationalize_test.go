package phe

import (
	"strings"
	"testing"
)

// Tests if the rationalized multivector is calculated properly.
func TestRationalize(t *testing.T) {
	c := []string{"40", "-29", "29", "-25", "33", "-29", "32", "-28"}

	m := New(c)
	m = Rationalize(m)

	if strings.Compare(m.E0.String(), "2313425") != 0 ||
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
