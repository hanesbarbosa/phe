package phe

import (
	"math/big"
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
	// m1 := CloneMultivector(m)
	// mi := Inverse(m1, q)
	// gp := GeometricProduct(m, mi, q)
	e0 := Rationalize(m, q).E0

	// After rationalizing m, if E0 is zero there is no inverse.
	return !(e0.Cmp(big.NewInt(0)) == 0)
}
