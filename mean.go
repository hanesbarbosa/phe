package phe

import (
	"math/big"
)

// Mean calculates the average amount between multivectors
func Mean(pk PublicKey, ms []Multivector) Multivector {
	n := len(ms)
	d := big.NewInt(int64(n))
	var m0 = []string{"0", "0", "0", "0", "0", "0", "0", "0"}

	tm := NewMultivector(m0)

	for i := 0; i < n; i++ {
		tm = Addition(pk, tm, ms[i])
	}

	return ScalarDivision(pk, tm, d)
}
