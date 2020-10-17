package phe

import "math/big"

// KeyUpdate substitutes the key pair s1 and s2 by s3 and s4 through tokens t1 and t2.
func KeyUpdate(pk *PublicKey, t *Token, c1 *Multivector) *Multivector {
	c2 := GeometricProduct(t.T1, c1, pk.Q)
	c2 = GeometricProduct(c2, t.T2, pk.Q)

	return c2
}

// KeyUpdateFromString executes function KeyUpdate by string inputs.
func KeyUpdateFromString(m string, t1 string, t2 string, c string) string {
	q := big.NewInt(0)
	q.SetString(m, 10)
	t1m := StringToMultivector(t1)
	t2m := StringToMultivector(t2)
	c1m := StringToMultivector(c)

	pk := &PublicKey{Q: q}
	t := &Token{T1: t1m, T2: t2m}

	return KeyUpdate(pk, t, c1m).ToString()
}
