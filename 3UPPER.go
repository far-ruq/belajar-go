package main

import (
	"fmt"
	"strings"
)

func main() {
	var kata string

	fmt.Print("Masukkan satu kata: ")
	fmt.Scanln(&kata)

	// ubah jadi huruf besar
	kataBesar := strings.ToUpper(kata)

	fmt.Printf("Hasil: %s\n", kataBesar)
}
