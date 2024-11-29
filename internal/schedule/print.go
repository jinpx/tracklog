package schedule

import (
	"fmt"
	"time"
)

// 自定义 Println 函数
func MyPrintln(v ...interface{}) {
	// 获取当前时间
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	// 打印时间戳和传入的内容
	fmt.Printf("[%s] ", currentTime)
	fmt.Println(v...)
}
