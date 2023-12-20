package main

import "fmt"

func main() {
	movie := "Avengers: Endgame"
	year := 2019
	rating := 8.4
	moivetype := "Sci-Fi"
	superhero := true
	r := '😁'

	fmt.Printf("เรื่อง: %s\n", movie)
	fmt.Printf("ปี: %d\n", year)
	fmt.Printf("เรตติ้ง: %.1f\n", rating)
	fmt.Printf("ประเภท %s\n", moivetype)
	fmt.Printf("ซุปเปอร์ฮีโร่ %t\n", superhero)
	fmt.Printf("ซุปเปอร์ฮีโร่ %c", r)
}
