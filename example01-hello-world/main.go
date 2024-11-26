package main

import (
	"fmt"
	"math/rand"
	"os"
)

const MaxRand = 16 //声明一个具名整型常量
// 一个函数声明
/*
 StatRandomNumbers 生成一些不大于MaxRand的非负随机整数
，并且统计和返回小于和大于MaxRand/2的随机数的个数。输入参数numRands指定了要生成的随机数的总数。
*/
func StatRandomNumbers(numRands int) (int, int) {
  //声明两个整型变量
	var a, b int
  // 一个for循环
	for i := 0; i < numRands; i++ {
    // 一个if语句
		if rand.Intn(MaxRand) < MaxRand/2 {
			a = a + 1
		} else {
			b++
		}
	}
	return a, b  // 返回两个整型变量
}
const NativeWordBits = 32 << (^uint(0) >> 63) // 64 or 32
const Is64bitOS = ^uint(0) >> 63 != 0
const Is32bitOS = ^uint(0) >> 32 == 0
const n  = 1 << 64
const r  = 'a'+0x7FFFFFFF
const x  = 3/2*0.1
const y  = 0.1*3/2

func main() {
	a := 50
	fmt.Printf(HelloWorld("appleboy"))
	fmt.Println("一天就學會 Go 語言")
	fmt.Println("南京黄学习 Go 語言")
	fmt.Println("ok puting is very good")
	b := 2
	c := a + b
	fmt.Println(c)

	if a >= 1 {

		fmt.Println("a >= 1")

	}
	if a >= 10 {

		fmt.Println("a >= 10")

	}
	if a >= 100 {

		fmt.Println("a >= 100")

	}
	if a >= 1000 {

		fmt.Println(" a > 1000")

	}
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	// 產生亂數
	var num = 100
	x1, y1 := StatRandomNumbers(num)
	print("Result：", x1, "+", y1, "=", num, "? ")
	println(x1+yf == num)
  	println('a' == 97)
	println('a' == '\141')
	println('a' == '\x61')
	println('a' == '\u0061')
	println('a' == '\U00000061')
	println(0x61 == '\x61')
	println('\u4f17' == '众')
  println("系统多少位",NativeWordBits)
  //声明一个字符串变量
  var lang,website string ="Go","https://golang.org"
  var compiled,dynamic bool = true,false 
  var announceYear int = 2009
  println(lang,website,compiled,dynamic,announceYear)
  // 变量lang 和year都为新声明的变量
  lang,year :="Go language",2007
  //这里，只有变量createdBy是新声明的变量,
  //变量year已经在上面声明过了，所以这里仅仅
  //改变了它的值，或者说它被重新声明了。
  year,createdBy := 2009,"Google Research"

  lang,year="Go",2012

  print(lang,"油",createdBy,"发明")
  print("并发布于",year,"年。")
  println()
  println(x)
  println(y)
}

func HelloWorld(user_name string) string {
	return fmt.Sprintf("Hi, %s ", user_name)
}
