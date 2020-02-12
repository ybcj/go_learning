/*
	编写测试程序
	1.源码文件以_test结尾：XXX_test.go
	2.测试方法名以Test开头：func TestXXX(t *testing.T) {...}
*/

package try_test

import "testing"

func TestFirstTry(t *testing.T) {
	t.Log("My first try!")
}
