package main

import "fmt"

func main() {
	movie := "Avengers: Endgame"
	year := 2019
	rating := 8.4
	moivetype := "Sci-Fi"
	superhero := true
	fmt.Println("เรื่อง:", movie)
	fmt.Println("ปี:", year)
	fmt.Println("เรตติ้ง:", rating)
	fmt.Println("ประเภท:", moivetype)
	fmt.Println("ซุปเปอร์ฮีโร่:", superhero)

	var r rune = '😁'
	fmt.Println("r:", r)
	fmt.Printf("r: %c\n", r)
	fmt.Printf("r: %#v\n", r)
}
