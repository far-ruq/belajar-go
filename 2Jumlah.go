package main

import (
	"fmt"
)

func main() {
	var angka1, angka2 int

	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&angka1)

	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&angka2)

	hasil := angka1 + angka2

	fmt.Printf("Hasil penjumlahan: %d + %d = %d\n", angka1, angka2, hasil)
}
