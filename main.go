package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
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

func isHappy(n int) bool {
	tostr := strconv.Itoa(n)
	set := map[string]int{}
	for true {
		result := 0
		set[tostr] += 1
		if set[tostr] > 1 {
			break
		}
		for _, v := range tostr {
			num, _ := strconv.Atoi(string(v))
			result += num * num
		}
		if result == 1 {
			return true
		}
		tostr = strconv.Itoa(result)
	}
	return false
}

func isValidP(s string) bool {
	parentheses := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	stack := []string{}

	for i := 0; i < len(s)/2; i++ {
		stack = append(stack, string(s[i]))
	}

	for i := len(s) / 2; i < len(s); i++ {
		if stack[len(stack)-1] == parentheses[string(s[i])] {
			stack = slices.Delete(stack, len(stack)-2, len(stack)-1)
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func kClosest(points [][]int, k int) [][]int {
	dis := map[int]int{}
	result := [][]int{}
	for i, v := range points {
		ans := 0
		ans += (v[0] * v[0]) + (v[1] * v[1])
		dis[i] = ans
	}
	sort.SliceStable(dis, func(i, j int) bool { return dis[dis[i]] > dis[dis[j]] })
	for i := 0; i < k; i++ {
		result[i] = append(result[i], dis[i])
	}
	return result
}

func canConstruct(ransomNote string, magazine string) bool {
	map_ransom := map[string]int{}
	for _, v := range ransomNote {
		map_ransom[string(v)] += 1
	}

	for _, v := range magazine {
		_, ok := map_ransom[string(v)]
		if ok && map_ransom[string(v)] != 0 {
			map_ransom[string(v)] -= 1
		}
	}

	for _, v := range map_ransom {
		if v > 0 {
			return false
		}
	}
	return true
}

// 389
func findTheDifference(s string, t string) byte {
	init := make(map[rune]int)
	var result byte
	for _, v := range s {
		init[v] += 1
	}

	for _, v := range t {
		_, ok := init[v]
		if !ok {
			result = byte(v)
		}
	}
	return result
}

func staircase(n int32) {
	// Write your code here
	for i := 1; i <= int(n); i++ {
		fmt.Printf("%s%s\n", strings.Repeat(" ", int(n)-i), strings.Repeat("#", i))
	}

}

func Solution_(A []int) int {
	// Implement your solution here
	sort.Ints(A)
	if A[len(A)-1] < 0 {
		return 1
	}
	mapping := make(map[int]int)
	result := []int{}
	for _, v := range A {
		_, ok := mapping[v]
		if !ok {
			mapping[v] += 1
			result = append(result, v)
		}
	}

	for i := 0; i < len(result)-1; i++ {
		if result[i+1] != result[i]+1 {
			return result[i] + 1
		}
	}
	return A[len(A)-1] + 1
}

func Solution(A []int) int {
	// Implement your solution here
	mapping := map[string]int{}
	for _, v := range A {
		mapping[strconv.Itoa(v)] += 1
	}

	for key, val := range mapping {
		if val == 1 {
			ans, _ := strconv.Atoi(key)
			return ans
		}
	}
	return -1
}

/*
func isSubsequence(s string, t string) bool {
    // consider i for "s" and j for "t"
    i, j := 0, 0
    length := len(t)

    for j < length {
        if i < len(s) && s[i] == t[j] {
            i++
        }
        j++
    }

    return len(s) == i
}
*/

func isSubsequence(s string, t string) bool {
	stack := []rune{}
	check := map[rune]int{}
	for _, val := range s {
		check[val]++
		stack = append(stack, val)
	}

	if len(stack) == 0 {
		return true
	}

	for i := len(t) - 1; i >= 0; i-- {
		_, ok := check[rune(t[i])]
		if ok {
			if rune(t[i]) == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
				if len(stack) == 0 {
					return true
				}
			}
		}
	}
	return false
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

	/*
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
	s := ""
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}
