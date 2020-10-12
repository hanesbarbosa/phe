package phe

import (
	"math/big"
	"strings"
	"testing"
)

func TestMean(t *testing.T) {
	sk, pk := GenerateKeys(256)

	m1 := big.NewInt(10)
	m2 := big.NewInt(20)
	m3 := big.NewInt(30)
	m4 := big.NewInt(40)

	c1 := Encrypt(sk, pk, m1)
	c2 := Encrypt(sk, pk, m2)
	c3 := Encrypt(sk, pk, m3)
	c4 := Encrypt(sk, pk, m4)

	ms := []*Multivector{c1, c2, c3, c4}
	m := Mean(pk, ms)

	rm := Decrypt(sk, pk, m)

	if strings.Compare(rm.String(), "25/1") != 0 {
		t.Errorf("The average is not correct!")
	}
}
