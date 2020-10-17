package phe

import (
	"math/big"
)

// Mean calculates the average amount between multivectors
func Mean(pk *PublicKey, ms []*Multivector) *Multivector {
	n := len(ms)
	d := big.NewInt(int64(n))
	tm := NewMultivector([]string{"0", "0", "0", "0", "0", "0", "0", "0"})

	for i := 0; i < n; i++ {
		tm = Addition(pk, tm, ms[i])
	}

	return ScalarDivision(pk, tm, d)
}

// MeanFromString parses strings and invokes Mean.
func MeanFromString(m string, ms []string) string {
	q := big.NewInt(0)
	q.SetString(m, 10)
	pk := &PublicKey{Q: q}
	var mv []*Multivector
	n := len(ms)

	for i := 0; i < n; i++ {
		mv = append(mv, StringToMultivector(ms[i]))
	}

	return Mean(pk, mv).ToString()
}
