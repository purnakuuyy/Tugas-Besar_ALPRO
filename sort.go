package main

import "fmt"

// ARRAY SEMENTARA UNTUK SORTING
var temp [MAX]LayananLangganan

// SELECTION SORT BERDASARKAN BIAYA (DESCENDING)
func SelectionSortBiaya() {

	if jumlahData == 0 {
		fmt.Println("Data kosong")
		return
	}

	// Copy data
	for i := 0; i < jumlahData; i++ {
		temp[i] = datalayanan[i]
	}

	for i := 0; i < jumlahData-1; i++ {

		max := i

		for j := i + 1; j < jumlahData; j++ {

			if temp[j].Detail.BiayaLayanan > temp[max].Detail.BiayaLayanan {
				max = j
			}
		}

		temp[i], temp[max] = temp[max], temp[i]
	}

	fmt.Println(">>> Data berhasil diurutkan berdasarkan biaya terbesar! <<<")

	fmt.Println("\n=== DATA BERDASARKAN BIAYA ===")

	for i := 0; i < jumlahData; i++ {

		fmt.Println("----------------------------")
		fmt.Println("ID                :", temp[i].ID)
		fmt.Println("Nama              :", temp[i].Nama)
		fmt.Println("Kategori          :", temp[i].Detail.KategoriLayanan)
		fmt.Println("Metode Pembayaran :", temp[i].Detail.MetodeBayar)
		fmt.Println("Biaya             :", temp[i].Detail.BiayaLayanan)
		fmt.Println("Jatuh Tempo       :", temp[i].Detail.JatuhTempo)

		if temp[i].Status {
			fmt.Println("Status            : Aktif")
		} else {
			fmt.Println("Status            : Tidak Aktif")
		}
	}

	fmt.Println("----------------------------")
}

// INSERTION SORT BERDASARKAN JATUH TEMPO (ASCENDING)
func InsertionSortJatuhTempo() {

	if jumlahData == 0 {
		fmt.Println("Data kosong")
		return
	}

	// Copy data
	for i := 0; i < jumlahData; i++ {
		temp[i] = datalayanan[i]
	}

	for i := 1; i < jumlahData; i++ {

		key := temp[i]
		j := i - 1

		for j >= 0 &&
			temp[j].Detail.JatuhTempo > key.Detail.JatuhTempo {

			temp[j+1] = temp[j]
			j--
		}

		temp[j+1] = key
	}

	fmt.Println(">>> Data berhasil diurutkan berdasarkan jatuh tempo! <<<")

	fmt.Println("\n=== DATA BERDASARKAN JATUH TEMPO ===")

	for i := 0; i < jumlahData; i++ {

		fmt.Println("----------------------------")
		fmt.Println("ID                :", temp[i].ID)
		fmt.Println("Nama              :", temp[i].Nama)
		fmt.Println("Kategori          :", temp[i].Detail.KategoriLayanan)
		fmt.Println("Metode Pembayaran :", temp[i].Detail.MetodeBayar)
		fmt.Println("Biaya             :", temp[i].Detail.BiayaLayanan)
		fmt.Println("Jatuh Tempo       :", temp[i].Detail.JatuhTempo)

		if temp[i].Status {
			fmt.Println("Status            : Aktif")
		} else {
			fmt.Println("Status            : Tidak Aktif")
		}
	}

	fmt.Println("----------------------------")
}
