package phe

// KeyUpdate substitutes the key pair s1 and s2 by s3 and s4 through tokens t1 and t2.
func KeyUpdate(pk PublicKey, t Token, c1 Multivector) Multivector {
	c2 := GeometricProduct(t.t1, c1, pk.Q)
	c2 = GeometricProduct(c2, t.t2, pk.Q)

	return c2
}
