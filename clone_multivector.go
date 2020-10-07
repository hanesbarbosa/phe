package phe

// CloneMultivector clones a multivector.
func CloneMultivector(m Multivector) Multivector {
	mc := NewMultivector([]string{
		m.E0.String(),
		m.E1.String(),
		m.E2.String(),
		m.E3.String(),
		m.E12.String(),
		m.E13.String(),
		m.E23.String(),
		m.E123.String(),
	})
	return mc
}
