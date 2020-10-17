package phe

import (
	"math/big"
	"strings"
	"testing"
)

func TestEncDec(t *testing.T) {
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

func TestTokGenKeyUpd(t *testing.T) {
	sk1, pk1 := GenerateKeys(256)
	m1 := big.NewInt(100)
	c1 := Encrypt(sk1, pk1, m1)

	sk2, pk2 := GenerateKeys(256)
	tk := GenerateToken(sk1, sk2, pk1, pk2)
	c2 := KeyUpdate(pk1, tk, c1)

	rm1 := Decrypt(sk2, pk2, c2)

	if strings.Compare(rm1.Num().String(), "100") != 0 {
		t.Errorf("the generated token or key update mechanism changed the original message")
	}
}

func TestTokGenKeyUpdString(t *testing.T) {
	sk1, pk1 := GenerateKeys(256)
	m1 := big.NewInt(100)
	c1 := Encrypt(sk1, pk1, m1)

	sk2, pk2 := GenerateKeys(256)
	tk := GenerateToken(sk1, sk2, pk1, pk2)
	c2 := KeyUpdateFromString(pk1.Q.String(), tk.T1.ToString(), tk.T2.ToString(), c1.ToString())

	rm1 := Decrypt(sk2, pk2, c2)

	if strings.Compare(rm1.Num().String(), "100") != 0 {
		t.Errorf("the generated token or key update mechanism changed the original message")
	}
}
