package main

import (
	"fmt"
)

func HitungLuas(a,t int)float64{
	var Luas float64
	Luas = (float64(a) * float64(t))/2
	return Luas
}

func main() {
	var N, i, a, t int
	var hasil float64

	print("Masukkan banyaknya segitiga : ")
	fmt.Scan(&N)

	for i = 1; i <= N; i++ {
		print("\n\n")

		print("Masukkan Alas Segitiga ke-", i, " : ")
		fmt.Scan(&a)

		print("Masukkan Tinggi Segitiga ke-", i, " : ")
		fmt.Scan(&t)

		hasil = HitungLuas(a,t)

		fmt.Print(hasil)
	}
}