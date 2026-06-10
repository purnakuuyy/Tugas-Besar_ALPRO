package main

const MAX = 100
const MAX_KATEGORI = 10

// DETAIL LAYANAN
type DetailLayanan struct {
	KategoriLayanan string
	MetodeBayar     string
	BiayaLayanan    int
	JatuhTempo      int
}

// DATA UTAMA LAYANAN
type LayananLangganan struct {
	ID     int
	Nama   string
	Detail DetailLayanan
	Status bool // true = Aktif, false = Tidak Aktif
}

// STATISTIK
type Statistik struct {
	TotalPengeluaran int
	TotalLayanan     int
}

// DAFTAR KATEGORI
var kategoridatalayanan = [MAX_KATEGORI]string{
	"Hiburan Video & Film (Streaming Video)",
	"Streaming Musik & Audio",
	"Produktivitas & Aplikasi Kerja (SaaS)",
	"Game & Hiburan Interaktif",
	"E-Book, Audiobooks & Literasi (Buku Digital)",
	"Pembelajaran & Pengembangan Diri (E-Learning)",
	"Berita & Jurnalisme Digital (Premium News)",
	"Kreator Konten & Komunitas Fans (Fan-Funding)",
	"Aplikasi Kesehatan & Kebugaran (Health & Wellness)",
	"Lainnya",
}

// ARRAY DATA
var datalayanan [MAX]LayananLangganan

// JUMLAH DATA
var jumlahData int

// VARIABEL STATISTIK
var statistik Statistik
