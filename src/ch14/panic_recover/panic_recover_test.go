package panicrecover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	// defer func() {
	// 	fmt.Println("Finally!")
	// }()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("something wrong!"))
	// os.Exit(-1)
	// fmt.Println("End")
}

// panic 用于不可恢复的错误
// panic 退出前会执行defer指定的内容

// panic vs. os.Exit
// os.Exit 退出时不会调用defer指定的函数
// os.Exit 退出时不输出当前调用栈的信息
