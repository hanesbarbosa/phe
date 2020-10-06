package phe

import "math/big"

// Inverse gives the inverse of a multivector.
func Inverse(m Multivector, q *big.Int) Multivector {
	mc := CloneMultivector(m)
	n := Numerator(m, q)
	d := DenominatorInverse(mc, q)
	return ScalarMultiplication(n, d, q)
}

// Numerator to be defined.
func Numerator(m Multivector, q *big.Int) Multivector {
	mc := CloneMultivector(m)
	as := AmplitudeSquared(m, q)
	asr := Reverse(as, q)
	cc := CliffordConjugation(mc, q)

	return GeometricProduct(cc, asr, q)
}

// // Denominator to be defined.
// func Denominator(m Multivector) string {
// 	r := Rationalize(m)
// 	return r.E0.Num().String()
// }

// Denominator to be defined.
func DenominatorInverse(m Multivector, q *big.Int) *big.Int {
	d := Rationalize(m, q).E0
	di := d.ModInverse(d, q)
	return di
}
