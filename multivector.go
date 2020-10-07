package phe

import (
	"math/big"
	"regexp"
)

// Multivector object
type Multivector struct {
	E0   *big.Int
	E1   *big.Int
	E2   *big.Int
	E3   *big.Int
	E12  *big.Int
	E13  *big.Int
	E23  *big.Int
	E123 *big.Int
}

// NewMultivector creates a new multivector from string coefficients
func NewMultivector(s []string) Multivector {
	var c [8]*big.Int
	var success bool

	for i := 0; i <= 7; i++ {
		c[i] = new(big.Int)
		_, success = c[i].SetString(s[i], 10)

		if !success {
			panic("Error converting string to big.Int.")
		}
	}

	return Multivector{
		E0:   c[0],
		E1:   c[1],
		E2:   c[2],
		E3:   c[3],
		E12:  c[4],
		E13:  c[5],
		E23:  c[6],
		E123: c[7],
	}
}

// ToString converts multivector to a string
func (m *Multivector) ToString() string {
	return m.E0.String() + "e0" + "+" +
		m.E1.String() + "e1" + "+" +
		m.E2.String() + "e2" + "+" +
		m.E3.String() + "e3" + "+" +
		m.E12.String() + "e12" + "+" +
		m.E13.String() + "e13" + "+" +
		m.E23.String() + "e23" + "+" +
		m.E123.String() + "e123"
}

// StringToMultivector converts strings into multivectors
func StringToMultivector(s string) Multivector {
	c := regexp.MustCompile("e[0-3]+(\\+)?").Split(s, -1)
	return NewMultivector(c)
}
