package main

import "fmt"

const (
	monday = iota + 1
	tuesday // 2
	wednesday //3
	thursday //4
	friday //5
	saturday //6 
	sunday  //7
)

// const (
// 	a = iota
// 	b
// )

const (
	a = iota //0 iota=0
	b // 1 iota=1
	c = iota + 2 //3 iota=2
	d // 4 iota=3
	e = iota + 1  //iota=4
)
type myConst int

const (
	zero myConst = iota // iota = 0
	one                 // iota = 1
	three = iota + 1    // iota = 2
	foure               // iota = 3
	five =iota               // iota = 4
	six      // iota = 5

)

func calc() (int, int) {
	return 1, 2
}

func main() {
	fmt.Println(one,three, foure, five,six)
	fmt.Println(monday)
	fmt.Println(tuesday)
	fmt.Println(wednesday)
	fmt.Println(thursday)
	fmt.Println(friday)
	fmt.Println(saturday)
	fmt.Println(sunday)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)


	test := true

	if test, test2 := calc(); test+test2 < 10 {
		fmt.Println("test:", test)
		fmt.Println("test2:", test2)
		fmt.Println("test + test2 < 10")
	}

	fmt.Println("test:", test)

	switch a := 1; {
	case a >= 0:
		fmt.Println("a is true")
		fallthrough
	case a > 200:
		fmt.Println("a is false")
	}

	stringSlice1 := []string{"1", "2", "3", "4"}
	fmt.Println(stringSlice1[:2])
	fmt.Println(stringSlice1[2:])

	stringSlice2 := make([]string, 5, 10)
	copy(stringSlice2, stringSlice1)
	fmt.Println(stringSlice2[0])

	stringSlice2 = append(stringSlice2, "5", "6")
	fmt.Println(stringSlice2)
	fmt.Println(stringSlice2[5])
	fmt.Println(stringSlice2[6])
	stringSlice2 = append(stringSlice2, []string{"7", "8"}...)
	fmt.Println(stringSlice2)
	fmt.Println(stringSlice2[7])
	fmt.Println(stringSlice2[8])


	const (
		a = iota   //0
		b          //1
		c          //2
		d = "ha"   //独立值，iota += 1
		e          //"ha"   iota += 1
		f = 100    //iota +=1
		g          //100  iota +=1
		h = iota   //7,恢复计数
		i          //8
)
fmt.Println(a,b,c,d,e,f,g,h,i)
}
