package phe

import (
	"math/big"
)

// Rationalize is a function that rationalizes a multivector.
func Rationalize(m *Multivector, q *big.Int) *Multivector {
	as := AmplitudeSquared(m, q)
	asc := CloneMultivector(as)
	asr := Reverse(as, q)
	gp := GeometricProduct(asc, asr, q)

	return gp
}
