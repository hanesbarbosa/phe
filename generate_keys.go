package phe

import (
	"crypto/rand"
	"math"
	"math/big"
)

// SK is the secret key structure
type SK struct {
	k1 Multivector
	k2 Multivector
	g  *big.Int
}

// PK is the public key structure
type PK struct {
	b int64
	q *big.Int
}

// GenerateKeys generates a secret key tuple SK and a public key tuple PK
func GenerateKeys(l int64) (SK, PK) {
	var b int64
	var g = new(big.Int)

	b = l / 8

	g.SetString(GenerateIntegers(b, 1)[0], 10)
	q := GeneratePrime(b)

	k1m := NewMultivector(GenerateIntegers(b, 8))
	k2m := NewMultivector(GenerateIntegers(b, 8))

	sk := SK{k1: k1m, k2: k2m, g: g}
	pk := PK{b: b, q: q}

	return sk, pk
}

// GenerateIntegers ...
func GenerateIntegers(b int64, n int) []string {
	// n cannot be 0
	var s []string
	var i *big.Int
	max := int64(math.Pow(float64(2), float64(b)))

	for len(s) < n {
		i, _ = rand.Int(rand.Reader, big.NewInt(max))
		s = append(s, i.String())
	}

	return s
}

// GeneratePrime ...
func GeneratePrime(b int64) *big.Int {
	// q > 2^b
	min := int(b) + 1
	p, err := rand.Prime(rand.Reader, min)
	em := "Error generating a prime number"
	if err != nil {
		em += err.Error()
		panic(em)
	}

	return p
}
