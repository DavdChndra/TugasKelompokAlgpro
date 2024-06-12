package main

import ("fmt")

func PerhitunganFibonacci(N int) int {
	var x int = 0
	var y int = 1
	var hitung int

	for i:=0;i<N;i++ {
		hitung += x
		x, y = y, x+y
	}
	return hitung
}

func main() {
	var N int

	print("Masukkan jumlah deret angka fibonacci : ")
	fmt.Scan(&N)

	HasilPerhitungan := PerhitunganFibonacci(N)

	fmt.Println(HasilPerhitungan)
}