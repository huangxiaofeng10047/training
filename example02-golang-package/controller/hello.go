package controller

import (
	"fmt"
	"time"
)

// HelloWorld func say hi
func HelloWorld(name string) string {
	//go 时间诞生之日
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("Hi, %s,today is %s", name, currentTime)
}
