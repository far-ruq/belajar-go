package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// gunakan bufio.NewReader atau bufio.NewScanner untuk nama yg pake spasi
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan nama lengkap: ")
	nama, _ := reader.ReadString('\n')

	// String slicing untuk mengambil karakter pertama
	// Indeks 0 sampai sebelum indeks 1
	inisial := nama[0:1]

	fmt.Printf("Inisial dari nama Anda adalah: %s\n", inisial)
}
