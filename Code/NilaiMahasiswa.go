package main

import (
	"fmt"
)

type Mahasiswa struct {
	Nama        string
	NIM         int
	Kelas       int
	MataKuliahs [2]MataKuliah
}

type MataKuliah struct {
	NamaMataKuliah string
	UTS            int
	UAS            int
	Quiz           int
	Total          int
	Grade          string
}

var ArrMahasiswa [2]Mahasiswa

func main() {
	for {
		fmt.Println("\n === Menu ===")
		fmt.Println("1. Input Data")
		fmt.Println("2. Melihat Data")
		fmt.Println("3. Update Data")
		fmt.Println("4. Mengurutkan Data")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih opsi: ")

		var pilihan int
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			inputDataMahasiswa()
		} else if pilihan == 2 {
			tampilkanData()
		} else if pilihan == 3 {
			// updateData()
		} else if pilihan == 4 {
			// tampilkanAscending()
		} else if pilihan == 5 {
			fmt.Println("Terima Kasih")
			return
		} else {
			fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
		}
	}
}

func inputDataMahasiswa() {
	for i := 0; i < len(ArrMahasiswa); i++ {
		fmt.Println("===================================")
		fmt.Printf("Masukkan nama mahasiswa %d: ", i+1)
		fmt.Scan(&ArrMahasiswa[i].Nama)

		fmt.Printf("Masukkan NIM mahasiswa %d: ", i+1)
		fmt.Scan(&ArrMahasiswa[i].NIM)

		fmt.Printf("Masukkan kelas mahasiswa %d: ", i+1)
		fmt.Scan(&ArrMahasiswa[i].Kelas)

		for j := 0; j < len(ArrMahasiswa[i].MataKuliahs); j++ {
			fmt.Println("  ===================================")
			fmt.Printf("  Masukkan nama mata kuliah %d untuk mahasiswa %s: ", j+1, ArrMahasiswa[i].Nama)
			fmt.Scan(&ArrMahasiswa[i].MataKuliahs[j].NamaMataKuliah)

			fmt.Printf("  Masukkan nilai UTS mata kuliah %s untuk mahasiswa %s: ", ArrMahasiswa[i].MataKuliahs[j].NamaMataKuliah, ArrMahasiswa[i].Nama)
			fmt.Scan(&ArrMahasiswa[i].MataKuliahs[j].UTS)

			fmt.Printf("  Masukkan nilai UAS mata kuliah %s untuk mahasiswa %s: ", ArrMahasiswa[i].MataKuliahs[j].NamaMataKuliah, ArrMahasiswa[i].Nama)
			fmt.Scan(&ArrMahasiswa[i].MataKuliahs[j].UAS)

			fmt.Printf("  Masukkan nilai Quiz mata kuliah %s untuk mahasiswa %s: ", ArrMahasiswa[i].MataKuliahs[j].NamaMataKuliah, ArrMahasiswa[i].Nama)
			fmt.Scan(&ArrMahasiswa[i].MataKuliahs[j].Quiz)

			ArrMahasiswa[i].MataKuliahs[j].Total = (ArrMahasiswa[i].MataKuliahs[j].UTS + ArrMahasiswa[i].MataKuliahs[j].UAS + ArrMahasiswa[i].MataKuliahs[j].Quiz)/3

			if ArrMahasiswa[i].MataKuliahs[j].Total >= 90 {
				ArrMahasiswa[i].MataKuliahs[j].Grade = "A"
			} else if ArrMahasiswa[i].MataKuliahs[j].Total >= 80 && ArrMahasiswa[i].MataKuliahs[j].Total < 90 {
				ArrMahasiswa[i].MataKuliahs[j].Grade = "B"
			} else if ArrMahasiswa[i].MataKuliahs[j].Total >= 70 && ArrMahasiswa[i].MataKuliahs[j].Total < 80 {
				ArrMahasiswa[i].MataKuliahs[j].Grade = "C"
			} else if ArrMahasiswa[i].MataKuliahs[j].Total >= 60 && ArrMahasiswa[i].MataKuliahs[j].Total < 70 {
				ArrMahasiswa[i].MataKuliahs[j].Grade = "D"
			} else {
				ArrMahasiswa[i].MataKuliahs[j].Grade = "E"
			}
		}
	}
}

func tampilkanData() {
	for i, mahasiswa := range ArrMahasiswa {
		fmt.Println("===================================")
		fmt.Printf("Nama Mahasiswa %d: %s\n", i+1, mahasiswa.Nama)
		fmt.Printf("NIM Mahasiswa %d: %d\n", i+1, mahasiswa.NIM)
		fmt.Printf("Kelas Mahasiswa %d: %d\n", i+1, mahasiswa.Kelas)
		for j, matakuliah := range mahasiswa.MataKuliahs {
			fmt.Println("  ===================================")
			fmt.Printf("  Nama Mata Kuliah %d: %s\n", j+1, matakuliah.NamaMataKuliah)
			fmt.Printf("  UTS %d: %d\n", j+1, matakuliah.UTS)
			fmt.Printf("  UAS %d: %d\n", j+1, matakuliah.UAS)
			fmt.Printf("  Quiz %d: %d\n", j+1, matakuliah.Quiz)
			fmt.Printf("  Total %d: %d\n", j+1, matakuliah.Total)
			fmt.Printf("  Grade %d: %s\n", j+1, matakuliah.Grade)
		}
	}
}

// func updateData(){
// 	var pilihanUpdateData int
// 	fmt.Println("\n === Menu Update Data ===")
// 	fmt.Println("1. Update Data Mahasiswa") 
// 	fmt.Println("2. Update Data Mata Kuliah Mahasiswa")
// 	fmt.Print("Pilih opsi: ")
// 	fmt.Scan(&pilihanUpdateData)

// 	if pilihanUpdateData == 1{
// 		var
// 	}else if pilihanUpdateData == 2{

// 	}
// }
