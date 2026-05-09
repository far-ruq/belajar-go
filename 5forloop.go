package main

import (
	"fmt"
)

func main() {
	fmt.Println("Angka kelipatan 2 dari 1 sampai 100:")

	for i := 2; i <= 100; i++ {
		// Jika i habis dibagi 3, maka i adalah kelipatan 3
		if i % 2 == 0 {
			fmt.Printf("%d ", i)
		}
	}
	
	fmt.Println("Angka kelipatan 4 dari 1 sampai 100 :")

	for i := 4; i <= 100; i++ {
		if i % 4 == 0 {
			fmt.Printf("%d ", i)
		}
	}
	
	fmt.Println()


}


