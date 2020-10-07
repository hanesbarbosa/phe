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

	g = GenerateIntegers(b, 1)[0]
	q := GeneratePrime(b)

	k1s := GenerateIntegers(b, 8)
	k2s := GenerateIntegers(b, 8)
	k1m := NewMultivector([]string{k1s[0].String(), k1s[1].String(), k1s[2].String(), k1s[3].String(), k1s[4].String(), k1s[5].String(), k1s[6].String(), k1s[7].String()})
	k2m := NewMultivector([]string{k2s[0].String(), k2s[1].String(), k2s[2].String(), k2s[3].String(), k2s[4].String(), k2s[5].String(), k2s[6].String(), k2s[7].String()})

	sk := SK{k1: k1m, k2: k2m, g: g}
	pk := PK{b: b, q: q}

	return sk, pk
}

// GenerateIntegers ...
func GenerateIntegers(b int64, n int) []*big.Int {
	// n cannot be 0
	var s []*big.Int
	var i *big.Int
	max := int64(math.Pow(float64(2), float64(b)))

	for len(s) < n {
		i, _ = rand.Int(rand.Reader, big.NewInt(max))
		s = append(s, i)
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
