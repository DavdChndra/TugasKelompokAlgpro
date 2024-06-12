package main

import ("fmt")

func Konversi(jam, menit, detik int) int {
	return (jam*3600) + (menit*60) + detik
}

func main() {
	var jam, menit, detik int
	var HasilKonversi int

	print("Masukkan Jam : ")
	fmt.Scan(&jam)

	print("Masukkan Menit : ")
	fmt.Scan(&menit)

	print("Masukkan detik : ")
	fmt.Scan(&detik)

	HasilKonversi = Konversi(jam, menit, detik)

	fmt.Println("Hasil Konversi : ", HasilKonversi)
}