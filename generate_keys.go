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
	b = l / 8

	k1i := GenerateIntegers(b, 8)
	k2i := GenerateIntegers(b, 8)
	g := GenerateIntegers(b, 1)[0]
	q := GeneratePrime(b)

	k1m := NewMultivector([]string{k1i[0].String(), k1i[1].String(), k1i[2].String(), k1i[3].String(), k1i[4].String(), k1i[5].String(), k1i[6].String(), k1i[7].String()})
	k2m := NewMultivector([]string{k2i[0].String(), k2i[1].String(), k2i[2].String(), k2i[3].String(), k2i[4].String(), k2i[5].String(), k2i[6].String(), k2i[7].String()})

	sk := SK{k1: k1m, k2: k2m, g: g}
	pk := PK{b: b, q: q}

	// fmt.Println("SK = ", sk)
	// fmt.Println("PK = ", pk)
	// fmt.Println("b=", b)
	// fmt.Println("g=", g)
	// fmt.Println("q=", q)
	// fmt.Println("k1=", k1)
	// fmt.Println("k2=", k2)
	// sk := new(SK)
	// pk := new(PK)
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
