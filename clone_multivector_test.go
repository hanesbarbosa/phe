package phe

import (
	"strings"
	"testing"
)

// Tests if the cloned multivector keeps separated values.
func TestCloneMultivector(t *testing.T) {
	c := []string{"40", "-29", "29", "-25", "33", "-29", "32", "-28"}

	m := NewMultivector(c)
	mc := CloneMultivector(m)

	if strings.Compare(m.E0.String(), mc.E0.String()) != 0 ||
		strings.Compare(m.E1.String(), mc.E1.String()) != 0 ||
		strings.Compare(m.E2.String(), mc.E2.String()) != 0 ||
		strings.Compare(m.E3.String(), mc.E3.String()) != 0 ||
		strings.Compare(m.E12.String(), mc.E12.String()) != 0 ||
		strings.Compare(m.E13.String(), mc.E13.String()) != 0 ||
		strings.Compare(m.E23.String(), mc.E23.String()) != 0 ||
		strings.Compare(m.E123.String(), mc.E123.String()) != 0 {
		t.Errorf("Different results for cloned multivector when it should be equal.")
	}

	m.E0.SetString("1", 10)
	m.E1.SetString("1", 10)
	m.E2.SetString("1", 10)
	m.E3.SetString("1", 10)
	m.E12.SetString("1", 10)
	m.E13.SetString("1", 10)
	m.E23.SetString("1", 10)
	m.E123.SetString("1", 10)

	if strings.Compare(m.E0.String(), mc.E0.String()) == 0 ||
		strings.Compare(m.E1.String(), mc.E1.String()) == 0 ||
		strings.Compare(m.E2.String(), mc.E2.String()) == 0 ||
		strings.Compare(m.E3.String(), mc.E3.String()) == 0 ||
		strings.Compare(m.E12.String(), mc.E12.String()) == 0 ||
		strings.Compare(m.E13.String(), mc.E13.String()) == 0 ||
		strings.Compare(m.E23.String(), mc.E23.String()) == 0 ||
		strings.Compare(m.E123.String(), mc.E123.String()) == 0 {
		t.Errorf("Equal results for cloned multivector when it should be different.")
	}
}
