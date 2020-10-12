package phe

import "math/big"

// ScalarMultiplication multiplies a multivector by a scalar.
func ScalarMultiplication(m *Multivector, s *big.Int, q *big.Int) *Multivector {
	m.E0.Mul(m.E0, s)
	m.E1.Mul(m.E1, s)
	m.E2.Mul(m.E2, s)
	m.E3.Mul(m.E3, s)
	m.E12.Mul(m.E12, s)
	m.E13.Mul(m.E13, s)
	m.E23.Mul(m.E23, s)
	m.E123.Mul(m.E123, s)

	return ModularMultivector(m, q)
}
