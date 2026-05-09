package main

import (
	"fmt"
)

func main() {
	var total float64

	fmt.Print("Masukkan Total Belanja: ")
	fmt.Scanln(&total)

	var diskon float64

	if total > 500000 {
		diskon = 0.25 // 25%
	} else if total > 100000 {
		diskon = 0.10 // 10%
	} else {
		diskon = 0.0
	}

	potongan := total * diskon
	hargaAkhir := total - potongan

	fmt.Printf("Total Belanja  : Rp %.0f\n", total)
	fmt.Printf("Diskon (%.0f%%)  : Rp %.0f\n", diskon*100, potongan)
	fmt.Printf("Harga Akhir    : Rp %.0f\n", hargaAkhir)
}
