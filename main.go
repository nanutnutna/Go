package main

import (
	"fmt"
)

func emote(rating float64) string {
	switch {
	case rating < 5.0:
		return "Dis"
	case rating >= 5.0 && rating < 7.0:
		return "Nor"
	case rating >= 7.0 && rating < 10.0:
		return "Good"
	default:
		return "Nil"
	}
}

func main() {

	fmt.Println(emote(4.9))
	fmt.Println(emote(5.0))
	fmt.Println(emote(7.0))
	fmt.Println(emote(15.0))

}
