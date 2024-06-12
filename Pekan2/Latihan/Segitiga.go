package main // untuk program utama yang akan dieksekusi

import (
	"fmt"
) // fmt package untuk operasi input-output (I/O) di Go

func MencariSegitiga(a,b,c int) string {
	if a + b > c && b + c > a && c + a > b{
		hasil := "Segitiga"
		return hasil
	}else{
		hasil := "Bukan Segitiga"
		return hasil
	}
}

func main() { // main function (function utama yang akan dieksekusi)
	var a, b, c int // variabel dengan nama "a, b, dan c" yang bertipe data integer

	print("Masukkan nilai a : ") // Memberikan keterangan / intruksi user apa yang harus di input
	fmt.Scan(&a) // untuk user input di terminal lalu disimpan kedalam variabel "a"

	print("Masukkan nilai b : ") // Memberikan keterangan / intruksi user apa yang harus di input
	fmt.Scan(&b) // untuk user input di terminal lalu disimpan kedalam variabel "b"

	print("Masukkan nilai c : ") // Memberikan keterangan / intruksi user apa yang harus di input
	fmt.Scan(&c) // untuk user input di terminal lalu disimpan kedalam variabel "c"

	segitiga := MencariSegitiga(a, b, c)

	fmt.Println(segitiga) // untuk menampilkan output isi variabel "segitiga"
}