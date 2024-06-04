package main

import (
	"fmt"
)

type Mahasiswa struct {
	Nama            string
	NIM             int
	Semester        int
	Jurusan         string
	MataKuliahs     [10]MataKuliah
	JumlahMataKuliah int
}

type MataKuliah struct {
	NamaMataKuliah string
	SKS            int
	UTS            int
	UAS            int
	Quiz           int
	Total          int
	Grade          string
}

var ArrMahasiswa [100]Mahasiswa
var jumlahMahasiswa int = 0
var NIM int = 103000

func main() {
	for {
		fmt.Println("\n === Menu ===")
		fmt.Println("1. Input Data")
		fmt.Println("2. Lihat Data")
		fmt.Println("3. Hapus Data")
		fmt.Println("4. Update Data")
		fmt.Println("5. Transkip Nilai")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih opsi: ")

		var pilihan int
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			inputDataMahasiswa()
		} else if pilihan == 2 {
			tampilkanData()
		} else if pilihan == 3 {
			hapusData()
		} else if pilihan == 4 {
			updateData()
		} else if pilihan == 5 {
			transkipNilai()
		} else if pilihan == 6 {
			fmt.Println("Terima Kasih")
			return
		} else {
			fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
		}
	}
}

func inputDataMahasiswa() {	
	var pilihanInputData int
	fmt.Println("\n === Menu Input Data ===")
	fmt.Println("1. Data Mahasiswa")
	fmt.Println("2. Data Mata Kuliah Mahasiswa")
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&pilihanInputData)

	if pilihanInputData == 1 {
		fmt.Println("===================================")
		fmt.Printf("Masukkan nama mahasiswa %d: ", jumlahMahasiswa+1)
		fmt.Scan(&ArrMahasiswa[jumlahMahasiswa].Nama)

		fmt.Printf("Masukkan Semester mahasiswa %s: ", ArrMahasiswa[jumlahMahasiswa].Nama)
		fmt.Scan(&ArrMahasiswa[jumlahMahasiswa].Semester)

		fmt.Printf("Masukkan Jurusan mahasiswa %s: ", ArrMahasiswa[jumlahMahasiswa].Nama)
		fmt.Scan(&ArrMahasiswa[jumlahMahasiswa].Jurusan)

		NIM++
		ArrMahasiswa[jumlahMahasiswa].NIM = NIM
		
		jumlahMahasiswa++
	} else if pilihanInputData == 2 {
		var NIM int
		fmt.Print("Masukkan NIM mahasiswa yang ingin ditambahkan mata kuliah: ")
		fmt.Scan(&NIM)

		found := false
		for i := 0; i < jumlahMahasiswa; i++ {
			if ArrMahasiswa[i].NIM == NIM {
				found = true
				mkIndex := ArrMahasiswa[i].JumlahMataKuliah
				if mkIndex < 10 { // Pastikan tidak melebihi kapasitas array
					fmt.Println("  ===================================")
					fmt.Printf("  Masukkan nama mata kuliah %d untuk mahasiswa %s: ", mkIndex+1, ArrMahasiswa[i].Nama)
					fmt.Scan(&ArrMahasiswa[i].MataKuliahs[mkIndex].NamaMataKuliah)

					fmt.Printf("  Masukkan SKS mata kuliah %s untuk mahasiswa %s: ", ArrMahasiswa[i].MataKuliahs[mkIndex].NamaMataKuliah, ArrMahasiswa[i].Nama)
					fmt.Scan(&ArrMahasiswa[i].MataKuliahs[mkIndex].SKS)

					fmt.Printf("  Masukkan nilai UTS mata kuliah %s untuk mahasiswa %s: ", ArrMahasiswa[i].MataKuliahs[mkIndex].NamaMataKuliah, ArrMahasiswa[i].Nama)
					fmt.Scan(&ArrMahasiswa[i].MataKuliahs[mkIndex].UTS)

					fmt.Printf("  Masukkan nilai UAS mata kuliah %s untuk mahasiswa %s: ", ArrMahasiswa[i].MataKuliahs[mkIndex].NamaMataKuliah, ArrMahasiswa[i].Nama)
					fmt.Scan(&ArrMahasiswa[i].MataKuliahs[mkIndex].UAS)

					fmt.Printf("  Masukkan nilai Quiz mata kuliah %s untuk mahasiswa %s: ", ArrMahasiswa[i].MataKuliahs[mkIndex].NamaMataKuliah, ArrMahasiswa[i].Nama)
					fmt.Scan(&ArrMahasiswa[i].MataKuliahs[mkIndex].Quiz)

					ArrMahasiswa[i].MataKuliahs[mkIndex].Total = (ArrMahasiswa[i].MataKuliahs[mkIndex].UTS + ArrMahasiswa[i].MataKuliahs[mkIndex].UAS + ArrMahasiswa[i].MataKuliahs[mkIndex].Quiz) / 3

					switch {
					case ArrMahasiswa[i].MataKuliahs[mkIndex].Total >= 90:
						ArrMahasiswa[i].MataKuliahs[mkIndex].Grade = "A"
					case ArrMahasiswa[i].MataKuliahs[mkIndex].Total >= 80:
						ArrMahasiswa[i].MataKuliahs[mkIndex].Grade = "B"
					case ArrMahasiswa[i].MataKuliahs[mkIndex].Total >= 70:
						ArrMahasiswa[i].MataKuliahs[mkIndex].Grade = "C"
					case ArrMahasiswa[i].MataKuliahs[mkIndex].Total >= 60:
						ArrMahasiswa[i].MataKuliahs[mkIndex].Grade = "D"
					default:
						ArrMahasiswa[i].MataKuliahs[mkIndex].Grade = "E"
					}

					ArrMahasiswa[i].JumlahMataKuliah++
				} else {
					fmt.Println("Mahasiswa ini sudah mencapai jumlah maksimum mata kuliah.")
				}
				break
			}
		}
		if !found {
			fmt.Println("  Data tidak ditemukan!")
		}
	}
}

func tampilkanData() {
	var pilihanTampilkanData int
	fmt.Println("\n === Menu Tampilkan Data ===")
	fmt.Println("1. Kumpulan data mahasiswa ") 
	fmt.Println("2. Daftar matakuliah yang diambil oleh mahasiswa ") 
	fmt.Println("3. Daftar mahasiswa yang mengambil suatu matakuliah ")
	fmt.Println("4. Jumlah SKS yang diambil oleh mahasiswa ")
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&pilihanTampilkanData)

	if pilihanTampilkanData == 1{
		fmt.Println("  ===== Kumpulan Data Mahasiswa =====")
		for i := 0; i < jumlahMahasiswa; i++ {
			fmt.Println("  ===================================")
			fmt.Printf("  Nama: %s\n", ArrMahasiswa[i].Nama)
			fmt.Printf("  NIM: %d\n", ArrMahasiswa[i].NIM)
			fmt.Printf("  Semester: %d\n", ArrMahasiswa[i].Semester)
			fmt.Printf("  Jurusan: %s\n", ArrMahasiswa[i].Jurusan)
		}
	} else if pilihanTampilkanData == 2{
		var NIM int
		fmt.Print("Masukkan NIM mahasiswa: ")
		fmt.Scan(&NIM)

		found := false
		for i := 0; i < jumlahMahasiswa; i++ {
			if ArrMahasiswa[i].NIM == NIM {
				found = true
				for j := 0; j < len(ArrMahasiswa[i].MataKuliahs); j++ {
					if ArrMahasiswa[i].MataKuliahs[j].NamaMataKuliah != "" {
						fmt.Println("  ===================================")
						fmt.Printf("  Nama Mata Kuliah: %s\n", ArrMahasiswa[i].MataKuliahs[j].NamaMataKuliah)
						fmt.Printf("  SKS: %d\n", ArrMahasiswa[i].MataKuliahs[j].SKS)
						fmt.Printf("  Nilai UTS: %d\n", ArrMahasiswa[i].MataKuliahs[j].UTS)
						fmt.Printf("  Nilai UAS: %d\n", ArrMahasiswa[i].MataKuliahs[j].UAS)
						fmt.Printf("  Nilai Quiz: %d\n", ArrMahasiswa[i].MataKuliahs[j].Quiz)
						fmt.Printf("  Total Nilai: %d\n", ArrMahasiswa[i].MataKuliahs[j].Total)
						fmt.Printf("  Grade: %s\n", ArrMahasiswa[i].MataKuliahs[j].Grade)
					}
				}
			}
		}
		if !found {
			fmt.Println("  Data tidak ditemukan!")
		}
	}else if pilihanTampilkanData == 3{
		var NamaMataKuliah string
		fmt.Print("Masukkan Nama Mata Kuliah: ")
		fmt.Scan(&NamaMataKuliah)

		found := false
		for i := 0; i < jumlahMahasiswa; i++ {
			mkIndex := ArrMahasiswa[i].JumlahMataKuliah
			for j := 0; j < mkIndex; j++ {
				if ArrMahasiswa[i].MataKuliahs[j].NamaMataKuliah == NamaMataKuliah {
					found = true
					fmt.Println("  ===================================")
					fmt.Printf("  Nama: %s\n", ArrMahasiswa[i].Nama)
					fmt.Printf("  NIM: %d\n", ArrMahasiswa[i].NIM)
					fmt.Printf("  Semester: %d\n", ArrMahasiswa[i].Semester)
					fmt.Printf("  Jurusan: %s\n", ArrMahasiswa[i].Jurusan)
				}
			}
		}
		if !found {
			fmt.Println("  Data tidak ditemukan!")
		}
	} else if pilihanTampilkanData == 4 {
		var NIM int
		fmt.Print("Masukkan NIM mahasiswa: ")
		fmt.Scan(&NIM)

		found := false
		for i := 0; i < jumlahMahasiswa; i++ {
			if ArrMahasiswa[i].NIM == NIM {
				found = true
				var TotalSKS int
				mkIndex := ArrMahasiswa[i].JumlahMataKuliah
				for j := 0 ; j < mkIndex; j++ {
					TotalSKS += ArrMahasiswa[i].MataKuliahs[j].SKS
				}
				fmt.Printf("  Total SKS yang di ambil: %d\n", TotalSKS)
			}
		}
		if !found {
			fmt.Println("  Data tidak ditemukan!")
		}
	}
}

func hapusData() {
	var pilihanHapusData int
	fmt.Println("\n === Menu Hapus Data ===")
	fmt.Println("1. Data Mahasiswa") 
	fmt.Println("2. Data Mata Kuliah Mahasiswa")
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&pilihanHapusData)

	if pilihanHapusData == 1{
		var NIM int
		fmt.Print("Masukkan NIM mahasiswa yang akan di hapus : ")
		fmt.Scan(&NIM)

		found := false
		for i := 0; i < jumlahMahasiswa; i++ {
			if ArrMahasiswa[i].NIM == NIM {
				found = true
				for j := i; j < jumlahMahasiswa - 1; i++ {
					ArrMahasiswa[j] = ArrMahasiswa[j+1]
				}
				ArrMahasiswa[jumlahMahasiswa - 1] = Mahasiswa{}
				jumlahMahasiswa--
				fmt.Println("  Data berhasil di-hapus.")
			}
		}
		if !found {
			fmt.Println("  Data tidak ditemukan!")
		}
	} else if pilihanHapusData == 2{
		var NIM int
		fmt.Print("Masukkan NIM mahasiswa : ")
		fmt.Scan(&NIM)

		var NamaMataKuliah string
		fmt.Print("Masukkan Nama Mata Kuliah mahasiswa yang akan di hapus : ")
		fmt.Scan(&NamaMataKuliah)

		found := false
		for i := 0; i < jumlahMahasiswa; i++ {
			if ArrMahasiswa[i].NIM == NIM {
				mkIndex := ArrMahasiswa[i].JumlahMataKuliah
				for j := 0; j < mkIndex; j++ {
					if ArrMahasiswa[i].MataKuliahs[j].NamaMataKuliah == NamaMataKuliah {
						found = true
						for k := j; k < mkIndex-1; k++ {
							ArrMahasiswa[i].MataKuliahs[k] = ArrMahasiswa[i].MataKuliahs[k+1]
						}
						ArrMahasiswa[i].MataKuliahs[mkIndex - 1] = MataKuliah{}
						mkIndex--
						fmt.Println("  Data berhasil di-hapus.")
					}
				}
			}
		}
		if !found {
			fmt.Println("  Data tidak ditemukan!")
		}
	}
}

func updateData() {
	var pilihanUpdateData int
	fmt.Println("\n === Menu Update Data ===")
	fmt.Println("1. Data Mahasiswa") 
	fmt.Println("2. Data Mata Kuliah Mahasiswa")
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&pilihanUpdateData)

	if pilihanUpdateData == 1{
		var NIM int
		fmt.Print("Masukkan NIM mahasiswa : ")
		fmt.Scan(&NIM)

		found := false
		for i := 0; i < jumlahMahasiswa; i++ {
			if ArrMahasiswa[i].NIM == NIM {
				found = true
				
				var pilihan int
				fmt.Println("\n === Pilih Salah Satu ===")
				fmt.Println("1. Nama")
				fmt.Println("2. NIM")
				fmt.Println("3. Semester")
				fmt.Println("4. Jurusan")
				fmt.Print("Pilih opsi: ")
				fmt.Scan(&pilihan)

				if pilihan == 1 {
					var NamaBaru string
					fmt.Printf("  Nama: %s\n", ArrMahasiswa[i].Nama)

					fmt.Println("  ===================================")

					fmt.Print("  Masukkan Nama baru: ")
					fmt.Scan(&NamaBaru)

					ArrMahasiswa[i].Nama = NamaBaru
					fmt.Println("  Data berhasil di-update.")

				} else if pilihan == 2 {
					var NIMBaru int
					fmt.Printf("  NIM: %d\n", ArrMahasiswa[i].NIM)

					fmt.Println("  ===================================")

					fmt.Print("  Masukkan NIM baru: ")
					fmt.Scan(&NIMBaru)

					ArrMahasiswa[i].NIM = NIMBaru
					fmt.Println("  Data berhasil di-update.")

				} else if pilihan == 3 {
					var SemesterBaru int
					fmt.Printf("  Semester: %s\n", ArrMahasiswa[i].Semester)

					fmt.Println("  ===================================")

					fmt.Print("  Masukkan Semester baru: ")
					fmt.Scan(&SemesterBaru)

					ArrMahasiswa[i].Semester = SemesterBaru
					fmt.Println("  Data berhasil di-update.")

				}else if pilihan == 4 {
					var JurusanBaru string
					fmt.Printf("  Jurusan: %s\n", ArrMahasiswa[i].Jurusan)

					fmt.Println("  ===================================")

					fmt.Print("  Masukkan Jurusan baru: ")
					fmt.Scan(&JurusanBaru)

					ArrMahasiswa[i].Jurusan = JurusanBaru
					fmt.Println("  Data berhasil di-update.")

				}
			}
		}
		if !found {
			fmt.Println("  Data tidak ditemukan!")
		}
	}else if pilihanUpdateData == 2 {
		var NIM int
		fmt.Print("Masukkan NIM mahasiswa : ")
		fmt.Scan(&NIM)

		var NamaMataKuliah string
		fmt.Print("Masukkan Nama Mata Kuliah mahasiswa yang akan di hapus : ")
		fmt.Scan(&NamaMataKuliah)

		found := false
		for i := 0; i < jumlahMahasiswa; i++ {
			if ArrMahasiswa[i].NIM == NIM {
				mkIndex := ArrMahasiswa[i].JumlahMataKuliah
				for j := 0; j < mkIndex; j++ {
					if ArrMahasiswa[i].MataKuliahs[j].NamaMataKuliah == NamaMataKuliah {
						found = true
	
						var NamaMKBaru string
						var SKSBaru int
						var UTSBaru int
						var UASBaru int
						var QuizBaru int
	
						var pilihan int
						fmt.Println("\n === Pilih Salah Satu ===")
						fmt.Println("1. Nama Mata Kuliah")
						fmt.Println("2. SKS")
						fmt.Println("3. UTS")
						fmt.Println("4. UAS")
						fmt.Println("5. Quiz")
						fmt.Print("Pilih opsi: ")
						fmt.Scan(&pilihan)
	
						if pilihan == 1 {
							fmt.Printf("  Nama Mata Kuliah: %s\n", ArrMahasiswa[i].MataKuliahs[j].NamaMataKuliah)
	
							fmt.Println("  ===================================")
	
							fmt.Print("  Masukkan Nama Mata Kuliah baru: ")
							fmt.Scan(&NamaMKBaru)
	
							ArrMahasiswa[i].MataKuliahs[j].NamaMataKuliah = NamaMKBaru
							fmt.Println("  Data berhasil di-update.")
						} else if pilihan == 2 {
							fmt.Printf("  SKS: %d\n", ArrMahasiswa[i].MataKuliahs[j].SKS)
	
							fmt.Println("  ===================================")
	
							fmt.Print("  Masukkan Nilai SKS baru: ")
							fmt.Scan(&SKSBaru)
	
							ArrMahasiswa[i].MataKuliahs[j].UTS = SKSBaru
						} else if pilihan == 3 {
							fmt.Printf("  Nilai UTS: %d\n", ArrMahasiswa[i].MataKuliahs[j].UTS)
	
							fmt.Println("  ===================================")
	
							fmt.Print("  Masukkan Nilai UTS baru: ")
							fmt.Scan(&UTSBaru)
	
							ArrMahasiswa[i].MataKuliahs[j].UTS = UTSBaru
							fmt.Println("  Data berhasil di-update.")
						} else if pilihan == 3 {
							fmt.Printf("  Nilai UAS: %d\n", ArrMahasiswa[i].MataKuliahs[j].UAS)
	
							fmt.Println("  ===================================")
	
							fmt.Print("  Masukkan Nilai UAS baru: ")
							fmt.Scan(&UASBaru)
	
							ArrMahasiswa[i].MataKuliahs[j].UAS = UASBaru
							fmt.Println("  Data berhasil di-update.")
						} else if pilihan == 4 {
							fmt.Printf("  Nilai Quiz: %d\n", ArrMahasiswa[i].MataKuliahs[j].Quiz)
	
							fmt.Println("  ===================================")
	
							fmt.Print("  Masukkan Nilai Quiz baru: ")
							fmt.Scan(&QuizBaru)
	
							ArrMahasiswa[i].MataKuliahs[j].Quiz = QuizBaru
							fmt.Println("  Data berhasil di-update.")
						}
	
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
		}
		if !found {
			fmt.Println("  Data tidak ditemukan!")
		}
	}
}

func transkipNilai() {
	var NIM int
	fmt.Print("Masukkan NIM mahasiswa: ")
	fmt.Scan(&NIM)

	found := false
	for i := 0; i < jumlahMahasiswa; i++ {
		if ArrMahasiswa[i].NIM == NIM {
			found = true
			fmt.Println("  ===================================")
			fmt.Printf("  Nama Mahasiswa: %s\n", ArrMahasiswa[i].Nama)
			fmt.Printf("  NIM: %d\n", ArrMahasiswa[i].NIM)
			fmt.Printf("  Semester: %d\n", ArrMahasiswa[i].Semester)
			fmt.Printf("  Jurusan: %s\n", ArrMahasiswa[i].Jurusan)
			for j := 0; j < ArrMahasiswa[i].JumlahMataKuliah; j++ {
				if ArrMahasiswa[i].MataKuliahs[j].NamaMataKuliah != "" {
					fmt.Println("  ===================================")
					fmt.Printf("  Nama Mata Kuliah: %s\n", ArrMahasiswa[i].MataKuliahs[j].NamaMataKuliah)
					fmt.Printf("  SKS: %d\n", ArrMahasiswa[i].MataKuliahs[j].SKS)
					fmt.Printf("  Nilai UTS: %d\n", ArrMahasiswa[i].MataKuliahs[j].UTS)
					fmt.Printf("  Nilai UAS: %d\n", ArrMahasiswa[i].MataKuliahs[j].UAS)
					fmt.Printf("  Nilai Quiz: %d\n", ArrMahasiswa[i].MataKuliahs[j].Quiz)
					fmt.Printf("  Total Nilai: %d\n", ArrMahasiswa[i].MataKuliahs[j].Total)
					fmt.Printf("  Grade: %s\n", ArrMahasiswa[i].MataKuliahs[j].Grade)
				}
			}
			break
		}
	}
	if !found {
		fmt.Println("Data mahasiswa tidak ditemukan.")
	}
}