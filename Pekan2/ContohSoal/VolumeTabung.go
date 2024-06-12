package main

import ("fmt")

//Perhitungan
func VolumeTabung(jari, tinggi float64) float64{
	var hasil float64
	hasil = 3.14 * (jari*jari) * tinggi
	return hasil
}

//Input dan Menampilkan Output
func main() {
	var jari, tinggi float64
	var volume float64

	print("Masukkan Jari-jari : ")
	fmt.Scan(&jari)

	print("Masukkan tinggi : ")
	fmt.Scan(&tinggi)

	volume = VolumeTabung(jari, tinggi) //perhitungan ke func VolumeTabung

	fmt.Println(volume)
}