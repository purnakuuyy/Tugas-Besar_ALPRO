package main

import "fmt"

// PENGINGAT JATUH TEMPO
func PengingatJatuhTempo() {

	if jumlahData == 0 {
		fmt.Println("Data kosong")
		return
	}

	var temp [MAX]LayananLangganan

	// Copy data
	for i := 0; i < jumlahData; i++ {
		temp[i] = datalayanan[i]
	}

	// Selection Sort berdasarkan tanggal jatuh tempo
	for i := 0; i < jumlahData-1; i++ {

		min := i

		for j := i + 1; j < jumlahData; j++ {

			if temp[j].Detail.JatuhTempo < temp[min].Detail.JatuhTempo {
				min = j
			}
		}

		temp[i], temp[min] = temp[min], temp[i]
	}

	fmt.Println("\n===== DAFTAR LAYANAN =====")

	for i := 0; i < jumlahData; i++ {

		fmt.Println("----------------------------")
		fmt.Println("Nama           :", datalayanan[i].Nama)
		fmt.Println("Kategori       :", datalayanan[i].Detail.KategoriLayanan)
		fmt.Println("Jatuh Tempo    :", datalayanan[i].Detail.JatuhTempo)

		if datalayanan[i].Status {
			fmt.Println("Status         : Aktif")
		} else {
			fmt.Println("Status         : Tidak Aktif")
		}
	}

	fmt.Println("\n===== 3 LAYANAN DENGAN JATUH TEMPO TERDEKAT =====")

	batas := 3
	if jumlahData < 3 {
		batas = jumlahData
	}

	for i := 0; i < batas; i++ {

		fmt.Println("----------------------------")
		fmt.Println("Nama           :", temp[i].Nama)
		fmt.Println("Kategori       :", temp[i].Detail.KategoriLayanan)
		fmt.Println("Jatuh Tempo    :", temp[i].Detail.JatuhTempo)
	}

	fmt.Println("----------------------------")
	fmt.Println("Segera lakukan pembayaran sebelum jatuh tempo.")
}

// LAPORAN TOTAL PENGELUARAN
func LaporanPengeluaran() {

	if jumlahData == 0 {
		fmt.Println("Data kosong")
		return
	}

	var stat Statistik

	stat.TotalLayanan = jumlahData

	for i := 0; i < jumlahData; i++ {

		if datalayanan[i].Status {
			stat.TotalPengeluaran += datalayanan[i].Detail.BiayaLayanan
		}
	}

	fmt.Println("\n===== LAPORAN LANGGANAN =====")
	fmt.Println("Total Layanan Aktif :", stat.TotalLayanan)
	fmt.Println("Total Pengeluaran   : Rp", stat.TotalPengeluaran)
}

// REKOMENDASI PENGHEMATAN
func RekomendasiPenghematan() {

	if jumlahData == 0 {
		fmt.Println("Data kosong")
		return
	}

	max := 0

	for i := 1; i < jumlahData; i++ {

		if datalayanan[i].Detail.BiayaLayanan >
			datalayanan[max].Detail.BiayaLayanan {

			max = i
		}
	}

	fmt.Println("\n===== REKOMENDASI PENGHEMATAN =====")
	fmt.Println("Nama Layanan :", datalayanan[max].Nama)
	fmt.Println("Kategori     :", datalayanan[max].Detail.KategoriLayanan)
	fmt.Println("Biaya        : Rp", datalayanan[max].Detail.BiayaLayanan)
	fmt.Println("Saran        : Pertimbangkan untuk menghentikan langganan ini agar pengeluaran bulanan lebih hemat.")
}
