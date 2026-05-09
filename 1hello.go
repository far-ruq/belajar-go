package main

import (
	"fmt"
)

func main() {
	var name string

	fmt.Print("Masukkan nama Anda: ")
	fmt.Scanln(&name)

	fmt.Printf("Halo %s, selamat belajar Go!\n", name)
}
