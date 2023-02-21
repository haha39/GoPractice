package main

import "fmt"

func main() {
	//
	fmt.Println("Hello world")
	//
	var num1 int8
	num1 = 127
	num2 := 20
	var num3 int = 10
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num2 / num3)
	//
	num4 := 8.7
	fmt.Println(num4)
	//
	var a rune = '小'
	var b rune = '櫻'
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%c %c \n", a, b)

	//
	var release string = "封印解除"
	fmt.Println(release)
	var c string = "小櫻今年"
	fmt.Print(c, 4, "歲\n")
	//
	num5 := 55
	if num5 > 10 {
		fmt.Print(num5, "\n")
	}
	//
	fmt.Println(false && false)
	//

}
