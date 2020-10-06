package phe

import "math/big"

// ModularMultivector calculates the modulus of a multivector.
func ModularMultivector(m Multivector, q *big.Int) Multivector {
	m.E0.Mod(m.E0, q)
	m.E1.Mod(m.E1, q)
	m.E2.Mod(m.E2, q)
	m.E3.Mod(m.E3, q)
	m.E12.Mod(m.E12, q)
	m.E13.Mod(m.E13, q)
	m.E23.Mod(m.E23, q)
	m.E123.Mod(m.E123, q)

	return m
}
