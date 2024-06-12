package main

import ("fmt")

func PerhitunganDiskon(TotalBelanjaAwal int, Membership bool) int{
	var TotalBelanjaAkhir, Diskon int
	if TotalBelanjaAwal > 100000 && Membership == true{
		Diskon = TotalBelanjaAwal * 10 / 100
		TotalBelanjaAkhir = TotalBelanjaAwal - Diskon
	}else if TotalBelanjaAwal > 100000 && Membership == false{
		Diskon = TotalBelanjaAwal * 5 / 100
		TotalBelanjaAkhir = TotalBelanjaAwal - Diskon
	} else{
		TotalBelanjaAkhir = TotalBelanjaAwal
	}
	return TotalBelanjaAkhir
}

func main() {
	var TotalBelanjaAwal int
	var Membership bool

	print("Masukkan Total Belanja Awal : ")
	fmt.Scan(&TotalBelanjaAwal)

	print("Apakah Sudah Menjadi Membership? (true/false) : ")
	fmt.Scan(&Membership)

	Perhitungan := PerhitunganDiskon(TotalBelanjaAwal, Membership)

	fmt.Println(Perhitungan)
}