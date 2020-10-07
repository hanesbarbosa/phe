package phe

import (
	"math/big"
	"strings"
	"testing"
)

func TestEncryptionDecryption(t *testing.T) {
	sk, pk := GenerateKeys(256)

	m1 := big.NewInt(300)

	c1 := Encrypt(sk, pk, m1)

	rm1 := Decrypt(sk, pk, c1)

	if strings.Compare(m1.String(), rm1.Num().String()) != 0 {
		t.Errorf("decyphered message does not match original message")
	}
}

func TestAdd(t *testing.T) {
	sk, pk := GenerateKeys(256)

	m1 := big.NewInt(110)
	m2 := big.NewInt(190)

	c1 := Encrypt(sk, pk, m1)
	c2 := Encrypt(sk, pk, m2)

	c3 := Addition(pk, c1, c2)

	rm3 := Decrypt(sk, pk, c3)

	if strings.Compare(rm3.Num().String(), "300") != 0 {
		t.Errorf("decyphered result does not match the addition of the original numbers")
	}
}

func TestSDiv(t *testing.T) {
	sk, pk := GenerateKeys(256)

	m1 := big.NewInt(110)
	m2 := big.NewInt(190)
	s := big.NewInt(7)

	c1 := Encrypt(sk, pk, m1)
	c2 := Encrypt(sk, pk, m2)

	c3 := Addition(pk, c1, c2)

	c4 := ScalarDivision(pk, c3, s)

	rm4 := Decrypt(sk, pk, c4)

	if strings.Compare(rm4.String(), "300/7") != 0 {
		t.Errorf("decyphered result does not match the scalar division of the addition result")
	}
}
