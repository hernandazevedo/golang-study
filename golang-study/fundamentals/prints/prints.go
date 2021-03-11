package main

import "fmt"

func main() {
	fmt.Print("Same")
	fmt.Print(" line")

	fmt.Println("New")
	fmt.Println(" line")

	x := 3.141516

	xs := fmt.Sprint(x)

	fmt.Println("The value of xs is ", xs)
	fmt.Println("The value of x is ", x)

	fmt.Printf("The value of x is %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "Oops"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

}
