package main

import "fmt"

func main() {

	var x string = "#"
	num := 0
	fmt.Printf("Enter the num:")
	fmt.Scanf("%d", &num)

	for i := 1; i <= num; i++ {
		for j := num; j >= 1; j-- {
			if i <= j {
				fmt.Printf("%v ", x)
			}else {
				fmt.Print("	")
			}
		}
		fmt.Println()
	}
}
/*
	# # # # # # # # # # 
	# # # # # # # # # 
	# # # # # # # # 
	# # # # # # # 
	# # # # # # 
	# # # # # 
	# # # # 
	# # # 
	# # 
	# 
*/