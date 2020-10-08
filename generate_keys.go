package phe

import (
	"crypto/rand"
	"math"
	"math/big"
)

// SecretKey is the secret key structure
type SecretKey struct {
	k1 Multivector
	k2 Multivector
	g  *big.Int
}

// PublicKey is the public key structure
type PublicKey struct {
	b int64
	q *big.Int
}

// GenerateKeys generates a secret key tuple SK and a public key tuple PK
func GenerateKeys(l int64) (SecretKey, PublicKey) {
	var b int64
	var k []Multivector
	var ki []*big.Int
	var km Multivector
	var g = new(big.Int)

	b = l / 8
	g = GenerateIntegers(b, 1)[0]
	q := GenerateModulus(float64(b))

	// Until both randomized keys k1 and k2 have inverse
	for len(k) < 2 {
		ki = GenerateIntegers(b, 8)
		km = NewMultivector([]string{ki[0].String(), ki[1].String(), ki[2].String(), ki[3].String(), ki[4].String(), ki[5].String(), ki[6].String(), ki[7].String()})
		if HasInverse(km, q) {
			k = append(k, km)
		}
	}

	sk := SecretKey{k1: k[0], k2: k[1], g: g}
	pk := PublicKey{b: b, q: q}

	return sk, pk
}

// GenerateIntegers generates a big integer slice of b bits integers
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

// GeneratePrime generates a prime of b bits
func GeneratePrime(b int64) *big.Int {
	min := int(b) + 1
	p, err := rand.Prime(rand.Reader, min)
	if err != nil {
		panic(err.Error())
	}

	return p
}

// GenerateModulus generates the smallest prime > 2^b
func GenerateModulus(b float64) *big.Int {
	// q > 2^b
	min := math.Pow(2, b)
	p := big.NewInt(int64(min))
	i := big.NewInt(1)

	for !p.ProbablyPrime(0) {
		p.Add(p, i)
	}

	return p
}
