package main

import (
	"fmt"
)

func main() {
	fmt.Println("Program okeoce (1-20):")

	for i := 1; i <= 20; i++ {
		if i % 2 == 0 && i % 3 == 0 {
			// Habis dibagi 2 DAN 3
			fmt.Println("okeoce")
		} else if i % 2 == 0 {
			// Hanya habis dibagi 2
			fmt.Println("oke")
		} else if i % 3 == 0 {
			// Hanya habis dibagi 3
			fmt.Println("oce")
		} else {
			// Tidak habis dibagi keduanya
			fmt.Println(i)
		}
	}
}
