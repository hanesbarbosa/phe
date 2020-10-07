package phe

import (
	"math/big"
)

// Encrypt ...
func Encrypt(sk SecretKey, pk PublicKey, m *big.Int) Multivector {
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
func Decrypt(sk SecretKey, pk PublicKey, c Multivector) *big.Rat {
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

	return ExtendedEuclideanAlgorithm(pk.q, s)
}

// ExtendedEuclideanAlgorithm does...
func ExtendedEuclideanAlgorithm(p, c *big.Int) *big.Rat {
	var a, b []*big.Int
	var i int

	a = append(a, p)
	a = append(a, c)
	b = append(b, big.NewInt(0))
	b = append(b, big.NewInt(1))
	i = 1
	q := big.NewInt(0)

	f := big.NewRat(1, 2)
	f.SetString(p.String() + "/2")

	fl := big.NewFloat(0.0)
	fl.SetRat(f)
	fl.Sqrt(fl)
	var z *big.Int
	z, _ = fl.Int(z)

	// Loop
	// +1 if ai > floor(sqrt(p/2))
	for a[i].Cmp(z) == 1 {
		// q = floor(a[i-1]/a[i])
		q.Div(a[i-1], a[i])

		// a[i+1] = a[i-1] - q * a[i]
		// First: creates a big number on the next position.
		// Second: uses the Sub() function of the early created big number
		a = append(a, big.NewInt(0))
		a[i+1].Sub(a[i-1], a[i+1].Mul(q, a[i]))

		// b[i+1] = b[i-1] + q * b[i]
		b = append(b, big.NewInt(0))
		b[i+1].Add(b[i-1], b[i+1].Mul(q, b[i]))

		i = i + 1
	}

	// Creates a[i]/b[i] big rational
	m := big.NewRat(1, 1)
	m, success := m.SetString(a[i].String() + "/" + b[i].String())
	if !success {
		// Raise exception and abort
		panic("error converting m into big rational")
	}

	return m
}
