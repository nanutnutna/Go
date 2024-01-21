package main

import (
	"fmt"
	"math"
)

func Devisable(start, end int) []int {
	var result = []int{}
	for i := start; i <= end; i++ {
		if i%7 == 0 && i%5 != 0 {
			result = append(result, i)
		}
	}
	return result
}

func containsDuplicate(nums []int) bool {
	check := map[int]int{}
	for _, val := range nums {
		check[val] += 1
	}
	for _, v := range check {
		if v > 1 {
			return false
		}
	}
	return true
}

func solution(nums []int) int {
	count := map[int]int{}
	for _, val := range nums {
		count[val] += 1
	}

	result := 0
	for key, val := range count {
		if key == val && val > result {
			result = val
		}
	}

	if result == 0 {
		return 0
	}
	return result
}

func twoSum(nums []int, target int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i, j)
			}
		}
	}
	return result
}

type volumer interface {
	Volume() float64
}

type cube struct {
	edge float64
} // edge x edge x edge

type triangularPrism struct {
	base     float64
	attitude float64
	height   float64
} // 0.5 x base x attitude x height

func VolumeOf(v volumer) float64 {
	return v.Volume()
}

func (c cube) Volume() float64 {
	return c.edge * c.edge * c.edge
}

func (t triangularPrism) Volume() float64 {
	return 0.5 * t.base * t.attitude * t.height
}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxarea := 0
	for t := 0; t < len(height); t++ {
		tall := min(height[left], height[right])
		area := tall * (right - left)
		if area > maxarea {
			maxarea = area
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return maxarea
}

func min(a, b int) int {
	if a == b {
		return a
	}
	if a < b {
		return a
	} else {
		return b
	}
}

//------------------------------------------------------------------------------------------------------------------------------------------------------

func findWordsContaining(words []string, x byte) map[int]int {
	result := map[int]int{}
	for idx, v := range words {
		for i := 0; i < len(words); i++ {
			if string(x) == string([]rune(v)[i]) {
				result[idx] = idx
			}
		}
	}
	return result
}

func compareString(word string) {
	converter := map[rune]rune{}
	for _, runeValue := range word {
		converter[runeValue] = runeValue
	}
}

func test(word string) {
	for _, v := range word {
		fmt.Println(string(v))
	}
}

func isValid(s string) bool {
	parentheses := map[string]int{
		"(": 0,
		")": 0,
		"{": 0,
		"}": 0,
		"[": 0,
		"]": 0,
	}

	for _, v := range s {
		parentheses[string(v)] += 1
	}

	if parentheses["("] == parentheses[")"] {
		if parentheses["{"] == parentheses["}"] {
			if parentheses["["] == parentheses["]"] {
				return true
			}
		}
	}
	return false
}

func myPow(x float64, n int) float64 {
	up := float64(n)
	result := math.Pow(x, up)
	return float64(result)

}

type course struct {
	name, instructor string
	price            int
}

func discount(c *course) int {
	c.price = c.price - 599
	fmt.Println("discount:", c.price)
	return c.price
}

type movie struct {
	title       string
	year        int
	rating      float32
	votes       []float64
	genres      []string
	isSuperHero bool
}

func (m *movie) addVote(num float64) {
	m.votes = append(m.votes, num)
}

func main() {

	eg := &movie{
		title:       "Avengers: Endgame",
		year:        2019,
		rating:      8.4,
		votes:       []float64{7, 8, 9, 10, 6, 9, 9, 10, 8},
		genres:      []string{"Action", "Drama"},
		isSuperHero: true,
	}

	eg.addVote(8)

	fmt.Println("Votes:", eg.votes)
	/*

		fmt.Printf("%#v\n", Devisable(2000, 3000))

		body := `[
		{
			"imdbID": "tt4154796",
			"title": "Avengers: Endgame",
			"year": 2019,
			"rating": 8.4,
			"genres": ["Action","Drama"],
			"isSuperHero": true
		},
		{
			"imdbID": "tt4154756",
			"title" : "Avengers: Infinity War",
			"year" : 2018,
			"rating": 8.4,
			"genres": ["Action", "Sci-Fi"],
			"isSuperHero": true
		}
		]`

		ms := []movie{}
		err := json.Unmarshal([]byte(body), &ms)
		fmt.Printf("%#v\n", err)
		fmt.Printf("%#v\n", ms)

		for _, v := range ms {
			fmt.Println(v.Title)
		}


		nums := []int{1, 2, 3, 1}
		fmt.Println(containsDuplicate(nums))
	*/

	/*
		word := []string{"leet", "code"}
		x := byte('e')
		fmt.Println(findWordsContaining(word, x))
	*/

	/*
		word := "เทส"
		fmt.Print(utf8.RuneCountInString(word))
		test(word)
	*/
}
