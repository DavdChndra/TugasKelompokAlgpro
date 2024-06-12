package main

import ( "fmt" )

// Faktorial menghitung nilai faktorial dari suatu bilangan.
func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * faktorial(n-1)
}

// Permutasi menghitung nilai permutasi dari x dan y.
func Permutasi(x, y int) int {
	if x >= y {
		return faktorial(x) / faktorial(x-y)
	}
	return faktorial(y) / faktorial(y-x)
}

func main() {
	var x, y int

	// Meminta input dari pengguna
	print("Masukkan nilai pertama: ")
	fmt.Scanln(&x)
	print("Masukkan nilai kedua: ")
	fmt.Scanln(&y)

	// Menghitung faktorial x dan y
	xFaktorial := faktorial(x)
	yFaktorial := faktorial(y)

	// Menghitung permutasi x dan y
	var HasilPermutasi int
	if x >= y {
		HasilPermutasi = Permutasi(x, y)
	} else {
		HasilPermutasi = Permutasi(y, x)
	}

	// Menampilkan hasil
	fmt.Println(xFaktorial, yFaktorial, HasilPermutasi)
}
