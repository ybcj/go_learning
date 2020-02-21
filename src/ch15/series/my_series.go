package series

import "fmt"

/*
1、首字母大写来表示可以被包外代码访问；
2、代码的package可以和所在目录不一致
3、同一目录里的Go代码的package要保持一致
*/

// GetFibonacci 获取n位的Fibonacci序列
func GetFibonacci(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}

// 该函数不能被包外访问
func square(n int) int {
	return n * n
}

/*
init 方法
1、在main被执行前，所有依赖的package的init方法都会被执行
2、不同包的init函数按照包导入的依赖关系决定执行顺序
3、每个包可以有过个init函数
4、包的每个源文件也可以有过个init函数
*/

func init() {
	fmt.Println("init1")
}
func init() {
	fmt.Println("init2")
}
