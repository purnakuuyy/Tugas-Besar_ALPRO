package main

import "fmt"

// SEQUENTIAL SEARCH BERDASARKAN NAMA
func SequentialSearchNama() {

	var cari string
	var ditemukan bool = false

	fmt.Print("Masukkan Nama Layanan: ")
	fmt.Scan(&cari)

	for i := 0; i < jumlahData; i++ {

		if datalayanan[i].Nama == cari {

			fmt.Println("\n=== DATA DITEMUKAN ===")
			fmt.Println("ID                :", datalayanan[i].ID)
			fmt.Println("Nama              :", datalayanan[i].Nama)
			fmt.Println("Kategori          :", datalayanan[i].Detail.KategoriLayanan)
			fmt.Println("Metode Pembayaran :", datalayanan[i].Detail.MetodeBayar)
			fmt.Println("Biaya             :", datalayanan[i].Detail.BiayaLayanan)
			fmt.Println("Jatuh Tempo       :", datalayanan[i].Detail.JatuhTempo)

			if datalayanan[i].Status {
				fmt.Println("Status            : Aktif")
			} else {
				fmt.Println("Status            : Tidak Aktif")
			}

			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("Data tidak ditemukan")
	}
}

// SORT BERDASARKAN NAMA (Bubble Sort)
// Digunakan sebelum Binary Search
func SortNama() {

	for i := 0; i < jumlahData-1; i++ {

		for j := 0; j < jumlahData-i-1; j++ {

			if datalayanan[j].Nama > datalayanan[j+1].Nama {

				datalayanan[j], datalayanan[j+1] =
					datalayanan[j+1], datalayanan[j]
			}
		}
	}
}

// BINARY SEARCH BERDASARKAN NAMA
func BinarySearchNama() {

	// Data harus diurutkan terlebih dahulu
	SortNama()

	var cari string

	fmt.Print("Masukkan Nama Layanan: ")
	fmt.Scan(&cari)

	left := 0
	right := jumlahData - 1

	for left <= right {

		mid := (left + right) / 2

		if datalayanan[mid].Nama == cari {

			fmt.Println("\n=== DATA DITEMUKAN ===")
			fmt.Println("ID                :", datalayanan[mid].ID)
			fmt.Println("Nama              :", datalayanan[mid].Nama)
			fmt.Println("Kategori          :", datalayanan[mid].Detail.KategoriLayanan)
			fmt.Println("Metode Pembayaran :", datalayanan[mid].Detail.MetodeBayar)
			fmt.Println("Biaya             :", datalayanan[mid].Detail.BiayaLayanan)
			fmt.Println("Jatuh Tempo       :", datalayanan[mid].Detail.JatuhTempo)

			if datalayanan[mid].Status {
				fmt.Println("Status            : Aktif")
			} else {
				fmt.Println("Status            : Tidak Aktif")
			}

			return

		} else if datalayanan[mid].Nama < cari {

			left = mid + 1

		} else {

			right = mid - 1
		}
	}

	fmt.Println("Data tidak ditemukan")
}
