package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// 1. Input Nama Barang
	fmt.Print("Nama Barang : ")
	namaBarang, _ := reader.ReadString('\n')
	namaBarang = strings.TrimSpace(namaBarang)

	// 2. Input Harga Barang
	var harga float64
	fmt.Print("Harga Barang: ")
	fmt.Scanln(&harga)

	// 3. Input Jumlah
	var jumlah int
	fmt.Print("Jumlah Beli : ")
	fmt.Scanln(&jumlah)

	// 4. Hitung Total Awal
	totalAwal := harga * float64(jumlah)

	// 5. Hitung Diskon
	var diskon float64
	if totalAwal >= 1000000 {
		diskon = totalAwal * 0.10 // 10%
	} else if totalAwal >= 500000 {
		diskon = totalAwal * 0.05 // 5%
	} else if totalAwal <500000{
		diskon = totalAwal * 0.01 // 1%
	}

	// 6. Hitung Total Akhir
	totalAkhir := totalAwal - diskon

	// 7. Tampilkan Struk Pembayaran
	fmt.Println("\n==============================")
	fmt.Printf("Nama Barang : %s\n", namaBarang)
	fmt.Printf("Harga Satuan: Rp %.0f\n", harga)
	fmt.Printf("Jumlah      : %d\n", jumlah)
	fmt.Println("------------------------------")
	fmt.Printf("Total Awal  : Rp %.0f\n", totalAwal)
	fmt.Printf("Diskon      : Rp %.0f\n", diskon)
	fmt.Printf("TOTAL AKHIR : Rp %.0f\n", totalAkhir)
	fmt.Println("==============================")
}
