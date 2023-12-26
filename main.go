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

	genres := [...]string{"Action", "Adventure", "Fantasy"}
	fmt.Printf("before for loop: %#v\n", genres)
	for i := 0; i < len(genres); i++ {
		genres[i] = "Movie: " + genres[i]
	}

	for _, v := range genres {
		fmt.Println(v)
	}

	xs := []float64{4, 5, 7, 8, 3, 8, 0}
	ys := []float64{7, 2, 10, 9, 7}
	var vote []float64

	vote = append(xs, ys...)

	fmt.Println(vote[5:9])
}
