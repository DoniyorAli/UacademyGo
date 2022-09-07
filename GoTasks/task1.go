package main

import "fmt"

func main() {

	var a, b int
	fmt.Println("Enter the numbers!")
	fmt.Printf("first number:")
	fmt.Scanf("%d", &a)
	fmt.Printf("second number:")
	fmt.Scanf("%d", &b)
	fmt.Println("--->", "a =", a, ",", "b =", b)

	// variables numbers change with simple example
	b = a + b // 7 = 4 + 3
	a = b - a // 3 = 7 - 4
	b = b - a // 4 = 7 - 3
	fmt.Println("--->", "a =", a, ",", "b =", b)
}
