package phe

import (
	"math/big"
)

// ScalarDivision makes the division between a multivector and a scalar.
func ScalarDivision(pk PublicKey, m Multivector, s *big.Int) Multivector {
	i := new(big.Int).ModInverse(s, pk.Q)

	m.E0.Mul(m.E0, i)
	m.E1.Mul(m.E1, i)
	m.E2.Mul(m.E2, i)
	m.E3.Mul(m.E3, i)
	m.E12.Mul(m.E12, i)
	m.E13.Mul(m.E13, i)
	m.E23.Mul(m.E23, i)
	m.E123.Mul(m.E123, i)

	return ModularMultivector(m, pk.Q)
}
