package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		time.Sleep(50 * time.Millisecond)
		h := strings.Repeat("=", i) + strings.Repeat(" ", 99-i)
		fmt.Printf("\r%.0f%%[%s]", float64(i)/99*100, color(h, i))

		os.Stdout.Sync()
	}
	fmt.Println("\nAll System Go!")
}

func color(str string, i int) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, 1, 0, 30+(i%7), str, 0x1B)
}
