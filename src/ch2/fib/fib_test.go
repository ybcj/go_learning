package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	// var a = 1
	// var b = 1
	var (
		a int = 1
		b int = 1
	)
	// fmt.Println(" ", a)
	t.Log(a)

	for i := 0; i < 5; i++ {
		// fmt.Println(" ", b)
		t.Log(b)
		tmp := a
		a = b
		b = tmp + a
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2

	// tmp := a
	// a = b
	// b = tmp
	a, b = b, a
	t.Log(a, b)
}
