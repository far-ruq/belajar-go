package main

import (
	"fmt"
)

func main() {
	// Menentukan tinggi atau total baris bintang
	tinggi := 10

	fmt.Println("Pola Bintang:")

	// Outer loop untuk baris
	for i := 1; i <= tinggi; i++ {
		// Inner loop untuk mencetak bintang di setiap baris
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		// Pindah ke baris baru
		fmt.Println()
	}
}
