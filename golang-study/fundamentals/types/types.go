package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("The integer literal is ", reflect.TypeOf(32000))

	var b byte = 255

	fmt.Println("The byte is ", reflect.TypeOf(b))

	i1 := math.MaxInt64

	fmt.Println("The max value for int is", i1)
	fmt.Println("The type of i1 is", reflect.TypeOf(i1))

	var i2 rune = 'a' // represents a mapping of unicode table (int32)

	fmt.Println("The rune is ", reflect.TypeOf(i2))
	fmt.Println(i2)

	var x float32 = 49.99

	fmt.Println("The type if x is ", reflect.TypeOf(x))
	fmt.Println("The type of the literal 49.99 is ", reflect.TypeOf(49.99))

	bo := true
	fmt.Println("The type if bo is ", reflect.TypeOf(bo))
	fmt.Println(!bo)

	s1 := "Hello my name is Borracha"

	fmt.Println(s1 + "!")
	fmt.Println("THe size of the string is", len(s1))

	s2 := `Hello
	my
	name
	is
	Borracha
	`
	fmt.Println("THe size of the string is", len(s2))

	//var x char = 'b'|
	char := 'a'
	fmt.Println("The type if char is ", reflect.TypeOf(char))
	fmt.Println(char)
}
