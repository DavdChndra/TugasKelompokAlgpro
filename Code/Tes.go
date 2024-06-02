package main

import (
	"fmt"
	"sort"
)

// MataKuliah struct untuk menyimpan nama dan nilai mata kuliah
type MataKuliah struct {
	nama  string
	nilai []float64
}

// Fungsi untuk mencari nilai terkecil
func min(nilai []float64) float64 {
	min := nilai[0]
	for _, v := range nilai {
		if v < min {
			min = v
		}
	}
	return min
}

// Fungsi untuk mencari nilai terbesar
func max(nilai []float64) float64 {
	max := nilai[0]
	for _, v := range nilai {
		if v > max {
			max = v
		}
	}
	return max
}

// Fungsi untuk menghitung rata-rata
func average(nilai []float64) float64 {
	total := 0.0
	for _, v := range nilai {
		total += v
	}
	return total / float64(len(nilai))
}

// Fungsi untuk mencari nilai tengah
func median(nilai []float64) float64 {
	sort.Float64s(nilai)
	n := len(nilai)
	if n%2 == 0 {
		return (nilai[n/2-1] + nilai[n/2]) / 2
	}
	return nilai[n/2]
}

func main() {
	var pilihan int
	var mataKuliahs []MataKuliah

	for {
		fmt.Println("Aplikasi Sederhana Nilai Mahasiswa\n1. Masukkan Data \n2. Informasi Data \n3. Keluar Program")
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			var jumlahMataKuliah int
			fmt.Print("Jumlah mata kuliah: ")
			fmt.Scan(&jumlahMataKuliah)

			for i := 0; i < jumlahMataKuliah; i++ {
				var mk MataKuliah

				fmt.Printf("Nama mata kuliah ke-%d: ", i+1)
				fmt.Scan(&mk.nama)

				var jumlahMahasiswa int
				fmt.Printf("Jumlah Mahasiswa %s: ", mk.nama)
				fmt.Scan(&jumlahMahasiswa)

				mk.nilai = make([]float64, jumlahMahasiswa)
				for j := 0; j < jumlahMahasiswa; j++ {
					fmt.Printf("Nilai mata kuliah %s absen %d: ", mk.nama, j+1)
					fmt.Scan(&mk.nilai[j])
				}

				mataKuliahs = append(mataKuliahs, mk)
			}
		} else if pilihan == 2 {
			fmt.Println("Melihat Data")
			for _, mk := range mataKuliahs {
				fmt.Printf("\nMata kuliah: %s\n", mk.nama)
				fmt.Printf("Nilai terkecil: %.2f\n", min(mk.nilai))
				fmt.Printf("Nilai terbesar: %.2f\n", max(mk.nilai))
				fmt.Printf("Rata-rata: %.2f\n", average(mk.nilai))
				fmt.Printf("Nilai tengah: %.2f\n", median(mk.nilai))

				sort.Float64s(mk.nilai)
				fmt.Printf("Nilai diurutkan (ascending): %v\n", mk.nilai)

				sort.Sort(sort.Reverse(sort.Float64Slice(mk.nilai)))
				fmt.Printf("Nilai diurutkan (descending): %v\n", mk.nilai)
			}
		} else if pilihan == 3 {
			fmt.Println("Terima Kasih")
			break
		} else {
			fmt.Println("Pilihan tidak valid, silakan pilih lagi.")
		}
	}
}
