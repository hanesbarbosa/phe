package phe

import (
	"math/big"
	"strings"
)

// Inverse gives the inverse of a multivector.
func Inverse(m *Multivector, q *big.Int) *Multivector {
	mc := CloneMultivector(m)
	n := Numerator(m, q)
	d := DenominatorInverse(mc, q)
	return ScalarMultiplication(n, d, q)
}

// Numerator to be defined.
func Numerator(m *Multivector, q *big.Int) *Multivector {
	mc := CloneMultivector(m)
	as := AmplitudeSquared(m, q)
	asr := Reverse(as, q)
	cc := CliffordConjugation(mc, q)

	return GeometricProduct(cc, asr, q)
}

// DenominatorInverse to be defined.
func DenominatorInverse(m *Multivector, q *big.Int) *big.Int {
	d := Rationalize(m, q).E0
	di := d.ModInverse(d, q)
	return di
}

// HasInverse checks if a multivector (i.e: key) has inverse.
// If so, the decryption can occur.
func HasInverse(m *Multivector, q *big.Int) bool {
	m1 := CloneMultivector(m)
	mi := Inverse(m1, q)
	gp := GeometricProduct(m, mi, q)

	return (strings.Compare(gp.E0.String(), "1") == 0 &&
		strings.Compare(gp.E1.String(), "0") == 0 &&
		strings.Compare(gp.E2.String(), "0") == 0 &&
		strings.Compare(gp.E3.String(), "0") == 0 &&
		strings.Compare(gp.E12.String(), "0") == 0 &&
		strings.Compare(gp.E13.String(), "0") == 0 &&
		strings.Compare(gp.E23.String(), "0") == 0 &&
		strings.Compare(gp.E123.String(), "0") == 0)
}
