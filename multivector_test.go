package phe

import (
	"strings"
	"testing"
)

// Tests if the function is returning a multivector with rational coefficients.
func TestNew(t *testing.T) {
	n := []string{"12", "23", "34", "45", "56", "67", "78", "89"}
	m := New(n)
	// m.e0, m.e1, m.e2, m.e3, m.e12, m.e13, m.e23, m.e123
	if strings.Compare(m.E0.String(), n[0]) != 0 ||
		strings.Compare(m.E1.String(), n[1]) != 0 ||
		strings.Compare(m.E2.String(), n[2]) != 0 ||
		strings.Compare(m.E3.String(), n[3]) != 0 ||
		strings.Compare(m.E12.String(), n[4]) != 0 ||
		strings.Compare(m.E13.String(), n[5]) != 0 ||
		strings.Compare(m.E23.String(), n[6]) != 0 ||
		strings.Compare(m.E123.String(), n[7]) != 0 {
		t.Errorf("One or more coefficients are different from the given string.")
	}
}

func TestToString(t *testing.T) {
	m := New([]string{"12", "23", "34", "45", "56", "67", "78", "89"})

	if strings.Compare(m.ToString(), "12e0+23e1+34e2+45e3+56e12+67e13+78e23+89e123") != 0 {
		t.Errorf("Wrong formatting for multivector.")
	}

	m = New([]string{"12", "-23", "34", "45", "56", "-67", "78", "89"})

	if strings.Compare(m.ToString(), "12e0+-23e1+34e2+45e3+56e12+-67e13+78e23+89e123") != 0 {
		t.Errorf("Wrong formatting for multivector.")
	}
}

func TestStringToMultivector(t *testing.T) {
	s := "230e0+-179e1+156e2+-110e3+238e12+-192e13+246e23+-200e123"
	m := StringToMultivector(s)

	if strings.Compare(m.E0.String(), "230") != 0 ||
		strings.Compare(m.E1.String(), "-179") != 0 ||
		strings.Compare(m.E2.String(), "156") != 0 ||
		strings.Compare(m.E3.String(), "-110") != 0 ||
		strings.Compare(m.E12.String(), "238") != 0 ||
		strings.Compare(m.E13.String(), "-192") != 0 ||
		strings.Compare(m.E23.String(), "246") != 0 ||
		strings.Compare(m.E123.String(), "-200") != 0 {
		t.Errorf("One or more coefficients are different from the given string.")
	}

	s = "345801144e0+36816179e1+10367e2+36816179e3+10367e12+36816179e13+10367e23+10367e123"
	m = StringToMultivector(s)

	if strings.Compare(m.E0.String(), "345801144") != 0 ||
		strings.Compare(m.E1.String(), "36816179") != 0 ||
		strings.Compare(m.E2.String(), "10367") != 0 ||
		strings.Compare(m.E3.String(), "36816179") != 0 ||
		strings.Compare(m.E12.String(), "10367") != 0 ||
		strings.Compare(m.E13.String(), "36816179") != 0 ||
		strings.Compare(m.E23.String(), "10367") != 0 ||
		strings.Compare(m.E123.String(), "10367") != 0 {
		t.Errorf("One or more coefficients are different from the given string.")
	}
}
