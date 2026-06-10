package main

import "fmt"

func Menu() {

	var pilihan int

	for {

		fmt.Println("\n==== APLIKASI MANAJEMEN SUBSKRIPSI DIGITAL ====")
		fmt.Println("1. Tambah Layanan")
		fmt.Println("2. Lihat Data Layanan")
		fmt.Println("3. Ubah Data Layanan")
		fmt.Println("4. Hapus Data Layanan")
		fmt.Println("5. Cari Layanan (Sequential Search)")
		fmt.Println("6. Cari Layanan (Binary Search)")
		fmt.Println("7. Urutkan Berdasarkan Biaya (Selection Sort)")
		fmt.Println("8. Urutkan Berdasarkan Jatuh Tempo (Insertion Sort)")
		fmt.Println("9. Pengingat Jatuh Tempo")
		fmt.Println("10. Laporan Total Pengeluaran")
		fmt.Println("11. Rekomendasi Penghematan")
		fmt.Println("0. Keluar")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {

		case 1:
			TambahLayanan()

		case 2:
			TampilLayanan()

		case 3:
			UbahLayanan()

		case 4:
			HapusLayanan()

		case 5:
			SequentialSearchNama()

		case 6:
			BinarySearchNama()

		case 7:
			SelectionSortBiaya()

		case 8:
			InsertionSortJatuhTempo()

		case 9:
			PengingatJatuhTempo()

		case 10:
			LaporanPengeluaran()

		case 11:
			RekomendasiPenghematan()

		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
