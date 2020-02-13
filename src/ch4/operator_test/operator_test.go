package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executeable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1,2,3,4}
	b := [...]int{2,1,4,3}
	// c := [...]int{1,2,3,4,5}
	d := [...]int{1,2,3,4}

	t.Log(a==b)
	// t.Log(a==c)
	t.Log(a==d)
}

func TestBitClear(t *testing.T) {
	a := 7 // 0111
	a = a &^ Readable
	t.Logf("Readable = %04b, Writable = %04b, Executeable = %04b", Readable, Writable, Executeable)
	t.Log(a & Readable == Readable, a & Writable == Writable, a & Executeable == Executeable)
}
