package main

import "fmt"

func main() {

	var x string = "#"

	for i := 10; i >= 1; i-- {
		for j := 1; j <= 10; j++ {
			if i <= j {
				fmt.Printf("%v ", x)
			}else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

/*                
					  # 
					# # 
				  # # # 
				# # # # 
			  # # # # # 
			# # # # # # 
		  # # # # # # # 
		# # # # # # # # 
	  # # # # # # # # # 
	# # # # # # # # # #
*/