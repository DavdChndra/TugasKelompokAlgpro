package main

import (
	"fmt"
)

// Fungsi untuk memecah uang ke dalam pecahan Rp10.000, Rp2.000, dan Rp1.000
func pecahuang(uang int, k10, k2, k1 *int) {
	*k10 = uang / 10000  // menghitung jumlah lembar 10.000
	sisa := uang % 10000 // sisa setelah dibagi 10.000

	*k2 = sisa / 2000 // menghitung jumlah lembar 2.000
	sisa = sisa % 2000  // sisa setelah dibagi 2.000

	*k1 = sisa / 1000 // menghitung jumlah lembar 1.000
}

func main() {
	var Uang int

	fmt.Print("Masukkan Nominal Uang: ")
	fmt.Scan(&Uang)

	var SepuluhRibu, DuaRibu, Seribu int

	pecahuang(Uang, &SepuluhRibu, &DuaRibu, &Seribu)
	fmt.Println("Uang pertama:", Uang)
	fmt.Println("Lembar Rp10.000:", SepuluhRibu)
	fmt.Println("Lembar Rp2.000:", DuaRibu)
	fmt.Println("Lembar Rp1.000:", Seribu)
}
