package main

import ("fmt")

func LuasLingkaran(jari float64) float64 {
	return 3.14*(jari*jari)
}

func LuasSelimut(jari, tinggi float64) float64 {
	return 2*3.14*jari*tinggi
}

func main() {
	var jari, tinggi, luas float64

	print("Masukkan Jari-Jari : ")
	fmt.Scan(&jari)

	print("Masukkan Tinggi : ")
	fmt.Scan(&tinggi)

	luas = 2 * LuasLingkaran(jari) + LuasSelimut(jari, tinggi)

	fmt.Println(luas)
}