package phe

import (
	"math/big"
	"testing"
)

// Tests if the reverse is calculated properly.
func TestInverse(t *testing.T) {
	// Assuring that m1 has inverse and m2 has no inverse.
	q := big.NewInt(257)
	m1 := NewMultivector([]string{"188", "157", "248", "157", "213", "18", "151", "48"})
	m2 := NewMultivector([]string{"214", "142", "106", "107", "8", "130", "137", "116"})

	if !HasInverse(m1, q) {
		t.Errorf("Multivector m1 should have inverse.")
	}

	if HasInverse(m2, q) {
		t.Errorf("Multivector m2 should not have inverse.")
	}
	// fmt.Println("Inverting m1...")
	// // m1 = Inverse(m1, q)
	// fmt.Println("Rationalizing m2: " + Rationalize(m2, q).ToString())
	// fmt.Println("Inverting m2...")
	// m2 = Inverse(m2, q)

	// if strings.Compare(m.E0.String(), "2") != 0 ||
	// 	strings.Compare(m.E1.String(), "254") != 0 ||
	// 	strings.Compare(m.E2.String(), "253") != 0 ||
	// 	strings.Compare(m.E3.String(), "252") != 0 ||
	// 	strings.Compare(m.E12.String(), "6") != 0 ||
	// 	strings.Compare(m.E13.String(), "7") != 0 ||
	// 	strings.Compare(m.E23.String(), "8") != 0 ||
	// 	strings.Compare(m.E123.String(), "248") != 0 {
	// 	t.Errorf("Wrong results for the reverse calculation.")
	// }
}
