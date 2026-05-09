package main

import (
	"fmt"
)

func main() {
	var angka int

	fmt.Print("Masukkan sebuah angka: ")
	fmt.Scanln(&angka)

	// Mengecek apakah angka habis dibagi 2 menggunakan modulo (%)
	if angka % 2 == 0 {
		fmt.Printf("Angka %d adalah Genap\n", angka)
	} else {
		fmt.Printf("Angka %d adalah Ganjil\n", angka)
	}
}
