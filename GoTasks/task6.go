package main

import "fmt"

func main() {
	DisplayNumberInReverseOrderWithDefer()
}

//*function=============================================

func DisplayNumberInReverseOrderWithDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
	}
}
