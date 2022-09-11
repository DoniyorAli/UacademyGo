package main

import (
	"fmt"
	"time"
)

func main() {

	var x string = "#"
	num := 0
	fmt.Printf("Enter the num:")
	fmt.Scanf("%d", &num)

	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			if i <= j {
				fmt.Printf("%s ", x)
				time.Sleep(50*time.Millisecond)
			} else {
				fmt.Print(" ")
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
	