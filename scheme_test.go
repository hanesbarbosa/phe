package phe

import (
	"fmt"
	"math/big"
	"testing"
)

// TestGenerateKeys...
func TestScheme(t *testing.T) {
	sk, pk := GenerateKeys(128)
	m1 := big.NewInt(110)
	m2 := big.NewInt(190)
	c1 := Encrypt(sk, pk, m1)
	c2 := Encrypt(sk, pk, m2)
	c3 := Addition(pk, c1, c2)
	rm := Decrypt(sk, pk, c3)

	// SK
	fmt.Println("sk.k1 = ", sk.k1.ToString())
	fmt.Println("sk.k2 = ", sk.k2.ToString())
	fmt.Println("sk.g = ", sk.g.String())
	// PK
	fmt.Println("pk.b = ", pk.b)
	fmt.Println("pk.q = ", pk.q)
	fmt.Println("###################################")
	// m
	// fmt.Println("message = ", m.String())
	// fmt.Println("ciphertext = ", c.ToString())
	fmt.Println("retrieved message = ", rm.String())

}
