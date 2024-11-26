package main

import (
	"fmt"

	"github.com/go-training/training/example02-golang-package/controller"
)
/*
* 函数声明和调用学习
* init 函数 可以在 main 函数执行之前执行 可以声明多个init函数
*/
var bob, smith = titledName("Bob"), titledName("Smith")
func init() {
	fmt.Println("hi,", bob)
}
func titledName(who string) string {
	return "Mr. " + who
}

func init() {
	fmt.Println("hello,", smith)
}



func SquaresOfSumAndDiff(a int64, b int64) (s int64, d int64) {
  //x,y :=a+b,a-b
  //s= x*x 
  //d = y*y
  //return // 等价于 return s,d
  return (a+b)*(a+b),(a-b)*(a-b)
}
func CompareLower4bit(m,n uint32) (r bool)  {
  r = m&0xF > n&0xf
  return
}
var v = VersionString()

func main() {
	fmt.Println("一天就學會 Go 語言")

	hi := controller.HelloWorld("hxf168482")
	fmt.Println(hi)
  
  println(v)
  x,y := SquaresOfSumAndDiff(3,6)
  println(x,y)
  b := CompareLower4bit(uint32(x),uint32(y))
  println(b)

  doNoting("GO",1)

  x2,y2 := func() (s int, d int) {
    println("匿名函数")
    return 3,4
}()// 一对小括号表示立即调用该匿名函数

func (a,b int)  {
  println("a*a + b*b  = ",a*a+b*b)

}(x2,y2)

func (x int)  {
    //形参x 遮挡了外层声明的变量x
    println("x*x +y*y = ",x*x+y2*y2)
  }  (y2)

  func ()  {
    println("x*x +y*y = ",x2*x2+y2*y2)
  }()
}


func VersionString() string {
  return "v1.0"
}
func doNoting(string,int32){
  //doNoth

}
