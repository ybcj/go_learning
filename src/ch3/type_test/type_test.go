package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	// b = a   // 必须显示类型转换
	b = int64(a)
	t.Log(a, b)

	var c MyInt = 3
	//  b= c // 必须显示类型转换
	b = int64(c)
	t.Log(b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	if s == "" {
		t.Logf("字符串为空，只能用s == \"\"来判断，不能用 s == nil，字符串是值类型")
	}
}
