package main

import "fmt"

func FizzBuzz(i int) int {
	if i % 3 == 0 {
		return i
	}else if i % 5 == 0 {
		return i
	}else if i % 3 == 0 && i % 5 == 0 {
		return i
	}
}

func main() {

	for i := 1; i <= 100; i+=1 {
		fmt.Println(FizzBuzz(i))
	}
}
