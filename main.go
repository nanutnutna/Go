package main

import (
	"fmt"
)

func main() {
	rating := 8.4
	/*
		if rating < 5.0 {
			fmt.Println("Dis")
		} else if rating >= 5.0 && rating < 7.0 {
			fmt.Println("Nor")
		} else if rating >= 7.0 && rating < 10.0 {
			fmt.Println("Good")
		} else {
			fmt.Println("Nil")
		}
	*/

	switch {
	case rating < 5.0:
		fmt.Println("Dis")
	case rating >= 5.0 && rating < 7.0:
		fmt.Println("Nor")
	case rating >= 7.0 && rating < 10.0:
		fmt.Println("Good")
	default:
		fmt.Println("Nil")
	}

}
