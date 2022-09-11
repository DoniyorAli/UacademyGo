package main

import "fmt"

func main() {

	num1, num2 := 0, 0
	fmt.Println("Enter the numbers!")
	fmt.Printf("first number:")
	fmt.Scanf("%d", &num1)
	fmt.Printf("second number:")
	fmt.Scanf("%d", &num2)
	fmt.Println(num1, "is odd: ", isOdd(num1))
	fmt.Println(num2, "is even: ", isEven(num2))
}

//*function===============================================
// Identify is Odd and is Even numbers
func isOdd(a int) bool {
	return a%2 == 0
}

func isEven(b int) bool {
	return b%2 != 0
}
