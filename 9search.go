package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Ketikkan kalimat: ")
	kalimat, _ := reader.ReadString('\n')
	kalimat = strings.TrimSpace(kalimat) // Menghapus karakter newline di akhir

	fmt.Print("Kata yang dicari: ")
	cari, _ := reader.ReadString('\n')
	cari = strings.TrimSpace(cari)

	// Menggunakan strings.Contains untuk mengecek keberadaan kata
	if strings.Contains(kalimat, cari) {
		fmt.Printf("Ketemu! Kata '%s' ada dalam kalimat.\n", cari)
	} else {
		fmt.Printf("Tidak ada kata '%s' dalam kalimat tersebut.\n", cari)
	}
}
