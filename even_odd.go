package main

import (
	"fmt"
)

func main() {

	fmt.Println("Enter digit:")
	var digit int
	fmt.Scan(&digit)
	if digit%2 == 0 {
		fmt.Printf("Число %d четное.", digit)
	} else {
		fmt.Printf("Число %d нечетное.", digit)
	}

}
