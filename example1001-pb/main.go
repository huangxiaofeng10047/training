package main

import (
	"time"

	"github.com/cheggaaa/pb/v3"
)

func main() {
	count := 100000

	// 创建进度条并开始
	bar := pb.StartNew(count)

	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(50 * time.Microsecond)
	}

	// 结束进度条
	bar.Finish()
}
