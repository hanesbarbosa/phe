package phe

import (
	"math/big"
	"strings"
	"testing"
)

// Tests the addition of two multivectors.
func TestAddition(t *testing.T) {
	pk := PublicKey{q: big.NewInt(10)}

	m1 := NewMultivector([]string{"2", "3", "4", "5", "6", "7", "8", "9"})
	m2 := NewMultivector([]string{"2", "3", "4", "5", "6", "7", "8", "9"})

	m := Addition(pk, m1, m2)

	if strings.Compare(m.E0.String(), "4") != 0 ||
		strings.Compare(m.E1.String(), "6") != 0 ||
		strings.Compare(m.E2.String(), "8") != 0 ||
		strings.Compare(m.E3.String(), "0") != 0 ||
		strings.Compare(m.E12.String(), "2") != 0 ||
		strings.Compare(m.E13.String(), "4") != 0 ||
		strings.Compare(m.E23.String(), "6") != 0 ||
		strings.Compare(m.E123.String(), "8") != 0 {
		t.Errorf("The addition is not correct!")
	}
}
