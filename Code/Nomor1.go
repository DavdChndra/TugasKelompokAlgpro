package main

import "fmt"

var tahun [8]int

func tampilkanMenu() {
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

		switch pilihan {
		case 1:
			inputData()
		case 2:
			tampilkanData()
		case 3:
			updateData()
		case 4:
			tampilkanAscending()
		case 5:
			fmt.Println("Terima Kasih")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
		}
	}
}

func inputData() {
	for i := 0; i < 8; i++ {
		fmt.Printf("Data penjualan tahunan ke-%d: ", i+1)
		fmt.Scan(&tahun[i])
	}
}

func tampilkanData() {
	fmt.Println("Seluruh Data:")
	for i := 0; i < 8; i++ {
		fmt.Println(i+1, ". ", tahun[i])
	}
}

func updateData() {
	var indeks, tahunBaru int
	fmt.Print("Masukkan tahun yang akan di-update (1-8): ")
	fmt.Scan(&indeks)

	if indeks < 1 || indeks > 8 {
		fmt.Println("Nomor item tidak valid.")
		return
	}

	fmt.Print("Masukkan tahun baru: ")
	fmt.Scan(&tahunBaru)

	tahun[indeks-1] = tahunBaru
	fmt.Println("Data berhasil di-update.")
}

func tampilkanAscending() {
	for i := 0; i < len(tahun)-1; i++ {
		for j := 0; j < len(tahun)-i-1; j++ {
			if tahun[j] > tahun[j+1] {
				tahun[j], tahun[j+1] = tahun[j+1], tahun[j]
			}
		}
	}

	fmt.Println("Data yang sudah diurutkan secara ascending:")
	for i, qty := range tahun {
		fmt.Printf("Item %d: %d\n", i+1, qty)
	}
}

func main() {
	var username, sandi string

	fmt.Print("Masukkan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan kata sandi: ")
	fmt.Scan(&sandi)

	if username == "geoleo31" && sandi == "leoleo71" {
		tampilkanMenu()
	} else {
		fmt.Println("Nama username atau kata sandi salah.")
		return
	}
}
