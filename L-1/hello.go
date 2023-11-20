package main

import (
	"fmt"
)

var myName string = "Narongrit"
var age int = 20
var t bool
var f bool = true
var favNum float64 = 2.822222
var complex complex128 = 5 + 5i
var r rune = 10

func test() {
	fmt.Println("Hello My name is " + myName)
	fmt.Println("age =", age)
	fmt.Println("I have a cat", f)
	fmt.Println("I have a dog", t)
	fmt.Println("My GPA is", favNum)
	fmt.Println("complex128 is", complex)
	fmt.Println("rune is", r)
	fmt.Println("Code ASCII = ", myName[0])
	fmt.Println("length of myName is", len(myName))
}

func main() {
	test()

	var arry5 [6]float64
	arry5[0] = 95.7
	arry5[1] = 93.2
	arry5[2] = 77.2
	arry5[3] = 83.7
	arry5[4] = 83.2
	arry5[5] = 99.2

	fmt.Println("length of arry is", len(arry5))
	fmt.Println("arry[3] is ", arry5[3])

	arry3 := [3]float64{98, 93, 77}
	fmt.Println(arry3)

	var d []float64 = arry5[0:2]
	fmt.Println(d)
}
