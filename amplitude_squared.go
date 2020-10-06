package phe

import "math/big"

// AmplitudeSquared gives the amplitude square of a multivector.
func AmplitudeSquared(m Multivector, q *big.Int) Multivector {
	// Clone the multivector
	mc := CloneMultivector(m)
	ccm := CliffordConjugation(m, q)
	return GeometricProduct(mc, ccm, q)
}
