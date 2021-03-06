package phe

import (
	"crypto/rand"
	"math/big"
)

// SecretKey is the secret key structure
type SecretKey struct {
	K1 *Multivector
	K2 *Multivector
	G  *big.Int
}

// PublicKey is the public key structure
type PublicKey struct {
	B int64
	Q *big.Int
}

// KeysToStrings ...
func (sk *SecretKey) KeysToStrings() string {
	return sk.K1.ToString() + "," + sk.K2.ToString()
}

// GenerateKeys generates a secret key tuple SK and a public key tuple PK
func GenerateKeys(l int64) (*SecretKey, *PublicKey) {
	var b int64
	var k []*Multivector
	var ki []*big.Int
	var km *Multivector
	var g = new(big.Int)

	b = l / 8
	g = GenerateIntegers(b, 1)[0]
	q := GenerateModulus(b)

	// Until both randomized keys k1 and k2 have inverse
	for len(k) < 2 {
		ki = GenerateIntegers(b, 8)
		km = NewMultivector([]string{ki[0].String(), ki[1].String(), ki[2].String(), ki[3].String(), ki[4].String(), ki[5].String(), ki[6].String(), ki[7].String()})
		if HasInverse(km, q) {
			k = append(k, km)
		}
	}

	sk := &SecretKey{K1: k[0], K2: k[1], G: g}
	pk := &PublicKey{B: b, Q: q}

	return sk, pk
}

// GenerateIntegers generates a big integer slice of b bits integers
func GenerateIntegers(b int64, n int) []*big.Int {
	// n cannot be 0
	var s []*big.Int
	var i *big.Int

	max := big.NewInt(0)
	e := big.NewInt(b)

	max.Exp(big.NewInt(2), e, nil)

	for len(s) < n {
		i, _ = rand.Int(rand.Reader, max)
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
func GenerateModulus(b int64) *big.Int {
	// q > 2^b
	e := big.NewInt(b)
	p := big.NewInt(0)
	p.Exp(big.NewInt(2), e, nil)
	i := big.NewInt(1)

	for !p.ProbablyPrime(0) {
		p.Add(p, i)
	}

	return p
}
