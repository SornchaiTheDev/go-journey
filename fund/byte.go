package main

import "fmt"

func main() {
	var charA byte = 'A'
	var hi []byte = []byte{'H', 'i'}


	// maybe the way string(X) implemented 
	// result := ""
	// for _, c := range hi {
	// 	result += fmt.Sprintf("%c" , c)
	// }

	fmt.Println(charA)
	fmt.Println(string(hi))
}
