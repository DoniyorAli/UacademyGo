package main

import "fmt"

func main() {

	for i := 1; i <= 100; i+=1 {
		FizzBuzz(i)
	}
}

//* function======================================

func FizzBuzz(i int) int{
	if i % 3 == 0 && i % 5 == 0 {
		fmt.Println(i,"===>","FizzBuzz")
	}else if i % 3 == 0 {
		fmt.Println(i,"===>","Fizz")
	}else if i % 5 == 0 {
		fmt.Println(i,"===>","Buzz")
	}else {
		fmt.Println(i)
	}
	return i
}