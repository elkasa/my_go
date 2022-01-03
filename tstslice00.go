package main

import (
	"fmt"
)

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func main() {
	t := []string{"a", "b", "c", "d", "e", "f", "g", "i"}

	fmt.Println(t)     //[0 1 2 3 4 5 6 7 8 9]
	fmt.Println(t[:2]) //[0 1 2 3 4 5 6 7 8 9]
	fmt.Println(t[3:]) //[0 1 2 3 4 5 6 7 8 9]

	t = RemoveIndex(t, 5)
	fmt.Println(t) //[0 1 2 3 4 6 7 8 9]
}
