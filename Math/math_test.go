package math

import "testing"

func TestDivide(t *testing.T) {
	divide(15, 2)
}

func TestAddB(t *testing.T) {
	a := "1111"
	b := "11"
	t.Logf("%s + %s = %s", a, b, addBinary(a, b))
}
