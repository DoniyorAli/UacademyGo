package main

import "fmt"

func main() {

	var x string = "#"

	for i := 1; i < 10; i++ {
		for i := 1; i < 10; i++ {
			fmt.Printf("%v ", x)
		}
		fmt.Println()
	}
}
/* 	
	# # # # # # # # # 
	# # # # # # # # # 
	# # # # # # # # # 
	# # # # # # # # # 
	# # # # # # # # # 
	# # # # # # # # # 
	# # # # # # # # # 
	# # # # # # # # # 
	# # # # # # # # # 
*/