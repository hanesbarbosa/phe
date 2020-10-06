package phe

import (
	"math/big"
)

// ScalarDivision makes the division between a multivector and a scalar.
func ScalarDivision(m Multivector, s *big.Int, q *big.Int) Multivector {
	i := new(big.Int).ModInverse(s, q)

	m.E0.Mul(m.E0, i)
	m.E1.Mul(m.E1, i)
	m.E2.Mul(m.E2, i)
	m.E3.Mul(m.E3, i)
	m.E12.Mul(m.E12, i)
	m.E13.Mul(m.E13, i)
	m.E23.Mul(m.E23, i)
	m.E123.Mul(m.E123, i)

	return ModularMultivector(m, q)
}
