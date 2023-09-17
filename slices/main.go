package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	// whenever we append to a slice a value that exceeds the capacity of the underlying array,
	// the capacity of the slice is doubled!
	s = append(s, 110) // cap=20
	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2])

	s = append(s, 120, 130, 140, 150, 160, 170, 180, 190, 200, 210) // cap=40
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}