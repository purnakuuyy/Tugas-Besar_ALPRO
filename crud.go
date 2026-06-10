package main

import "fmt"

// TAMBAH DATA
func TambahLayanan() {

	if jumlahData >= MAX {
		fmt.Println("Data penuh")
		return
	}

	fmt.Println("\n=== Tambah Layanan Baru ===")

	fmt.Print("ID: ")
	fmt.Scan(&datalayanan[jumlahData].ID)

	fmt.Print("Nama Layanan: ")
	fmt.Scan(&datalayanan[jumlahData].Nama)

	var pilihanKategori int

	fmt.Println("Kategori yang tersedia:")
	for i := 0; i < MAX_KATEGORI; i++ {
		fmt.Printf("%d. %s\n", i+1, kategoridatalayanan[i])
	}

	fmt.Print("Pilih Kategori (1-10): ")
	fmt.Scan(&pilihanKategori)

	if pilihanKategori >= 1 && pilihanKategori <= MAX_KATEGORI {
		datalayanan[jumlahData].Detail.KategoriLayanan =
			kategoridatalayanan[pilihanKategori-1]
	} else {
		fmt.Println("Kategori tidak valid")
		return
	}

	fmt.Print("Metode Pembayaran: ")
	fmt.Scan(&datalayanan[jumlahData].Detail.MetodeBayar)

	fmt.Print("Biaya Langganan: ")
	fmt.Scan(&datalayanan[jumlahData].Detail.BiayaLayanan)

	fmt.Print("Jatuh Tempo (YYYYMMDD): ")
	fmt.Scan(&datalayanan[jumlahData].Detail.JatuhTempo)

	var status int
	fmt.Print("Status (1=Aktif, 0=Tidak Aktif): ")
	fmt.Scan(&status)

	datalayanan[jumlahData].Status = (status == 1)

	jumlahData++

	fmt.Println("Layanan berhasil ditambahkan!")
}

// TAMPIL DATA
func TampilLayanan() {

	if jumlahData == 0 {
		fmt.Println("Data kosong")
		return
	}

	fmt.Println("\n=== Data Layanan Langganan ===")

	for i := 0; i < jumlahData; i++ {

		fmt.Println("--------------------------")
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
	}

	fmt.Println("--------------------------")
}

// UBAH DATA
func UbahLayanan() {

	var id int

	fmt.Print("Masukkan ID yang ingin diubah: ")
	fmt.Scan(&id)

	for i := 0; i < jumlahData; i++ {

		if datalayanan[i].ID == id {

			fmt.Print("Nama Baru: ")
			fmt.Scan(&datalayanan[i].Nama)

			var pilihanKategori int

			fmt.Println("Kategori yang tersedia:")
			for j := 0; j < MAX_KATEGORI; j++ {
				fmt.Printf("%d. %s\n", j+1, kategoridatalayanan[j])
			}

			fmt.Print("Pilih Kategori Baru (1-10): ")
			fmt.Scan(&pilihanKategori)

			if pilihanKategori >= 1 && pilihanKategori <= MAX_KATEGORI {
				datalayanan[i].Detail.KategoriLayanan =
					kategoridatalayanan[pilihanKategori-1]
			} else {
				fmt.Println("Kategori tidak valid")
				return
			}

			fmt.Print("Metode Pembayaran Baru: ")
			fmt.Scan(&datalayanan[i].Detail.MetodeBayar)

			fmt.Print("Biaya Langganan Baru: ")
			fmt.Scan(&datalayanan[i].Detail.BiayaLayanan)

			fmt.Print("Jatuh Tempo Baru (YYYYMMDD): ")
			fmt.Scan(&datalayanan[i].Detail.JatuhTempo)

			var status int
			fmt.Print("Status Baru (1=Aktif, 0=Tidak Aktif): ")
			fmt.Scan(&status)

			datalayanan[i].Status = (status == 1)

			fmt.Println("Data berhasil diubah!")
			return
		}
	}

	fmt.Println("Data tidak ditemukan")
}

// HAPUS DATA
func HapusLayanan() {

	var id int

	fmt.Print("Masukkan ID yang ingin dihapus: ")
	fmt.Scan(&id)

	for i := 0; i < jumlahData; i++ {

		if datalayanan[i].ID == id {

			for j := i; j < jumlahData-1; j++ {
				datalayanan[j] = datalayanan[j+1]
			}

			jumlahData--

			fmt.Println("Data berhasil dihapus!")
			return
		}
	}

	fmt.Println("Data tidak ditemukan")
}
