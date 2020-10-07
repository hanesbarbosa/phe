package phe

import "math/big"

// Encrypt ...
func Encrypt(sk SK, pk PK, m *big.Int) Multivector {
	// m0,...,m123 with the exception of m12
	s := GenerateIntegers(pk.b, 8)
	n := []string{s[0].String(), s[1].String(), s[2].String(), s[3].String(), s[4].String(), s[5].String(), s[6].String(), s[7].String()}

	// m12
	s[0].Mul(big.NewInt(-1), s[0])
	s[1].Mul(big.NewInt(-1), s[1])
	s[5].Mul(big.NewInt(-1), s[5])

	s[4] = big.NewInt(0)
	s[4].Add(s[0], s[1])
	s[4].Add(s[4], s[2])
	s[4].Add(s[4], s[3])
	s[4].Add(s[4], s[5])
	s[4].Add(s[4], s[6])
	s[4].Add(s[4], s[7])
	s[4].Add(s[4], m)
	s[4].Mod(s[4], pk.q)

	n[4] = s[4].String()

	mv := NewMultivector(n)
	mvp := ScalarMultiplication(mv, sk.g, pk.q)
	cg := GeometricProduct(sk.k1, mvp, pk.q)
	c := GeometricProduct(cg, sk.k2, pk.q)

	return c
}

// Decrypt ...
func Decrypt(sk SK, pk PK, c Multivector) *big.Int {
	k1i := Inverse(sk.k1, pk.q)
	k2i := Inverse(sk.k2, pk.q)
	cg := GeometricProduct(k1i, c, pk.q)
	mvp := GeometricProduct(cg, k2i, pk.q)
	mv := ScalarDivision(pk, mvp, sk.g)

	s := big.NewInt(0)
	s.Add(mv.E0, mv.E1)
	s.Sub(s, mv.E2)
	s.Sub(s, mv.E3)
	s.Add(s, mv.E12)
	s.Add(s, mv.E13)
	s.Sub(s, mv.E23)
	s.Sub(s, mv.E123)
	s.Mod(s, pk.q)

	return s
}
