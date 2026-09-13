package main

import (
	"fmt"
	"math"
)

// ==========================================
// 1. DEFINISI INTERFACE (KONTRAK)
// ==========================================
// Interface berfungsi sebagai 'kontrak'.
// Siapa pun struct yang ingin dikategorikan sebagai 'HitungDatar',
// ia WAJIB memiliki dua method: Luas() dan Keliling() dengan kembalian float64.
type HitungDatar interface {
	Luas() float64
	Keliling() float64
}

// ==========================================
// 2. STRUCT PERTAMA: PERSEGI
// ==========================================
// Struct Persegi menyimpan data 'Sisi'.
type Persegi struct {
	Sisi float64
}

// Method Luas() untuk Persegi.
// Karena Persegi punya Luas() dan Keliling(), otomatis Persegi
// dianggap mengimplementasikan interface HitungDatar (Implicit).
func (p Persegi) Luas() float64 {
	return p.Sisi * p.Sisi
}

// Method Keliling() untuk Persegi.
func (p Persegi) Keliling() float64 {
	return 4 * p.Sisi
}

// ==========================================
// 3. STRUCT KEDUA: LINGKARAN
// ==========================================
// Struct Lingkaran menyimpan data 'JariJari'.
type Lingkaran struct {
	JariJari float64
}

// Method Luas() untuk Lingkaran menggunakan rumus π * r^2.
func (l Lingkaran) Luas() float64 {
	return math.Pi * l.JariJari * l.JariJari
}

// Method Keliling() untuk Lingkaran menggunakan rumus 2 * π * r.
func (l Lingkaran) Keliling() float64 {
	return 2 * math.Pi * l.JariJari
}

// ==========================================
// 4. FUNGSI POLIMORFISME (Menerima Interface)
// ==========================================
// Parameter 'h' bertipe interface 'HitungDatar'.
// Fungsi ini bersifat fleksibel: bisa menerima Persegi, Lingkaran,
// atau bentuk lain selama memenuhi kontrak interface HitungDatar.
func CetakInfoBangun(h HitungDatar) {
	// Memanggil method Luas() dan Keliling() dari instance yang dikirim
	fmt.Printf("Luas     : %.2f\n", h.Luas())
	fmt.Printf("Keliling : %.2f\n", h.Keliling())
	fmt.Println("------------------------------")
}

// ==========================================
// 5. EMPTY INTERFACE (any) & TYPE SWITCH
// ==========================================
// Parameter 'data' bertipe 'any' (alias untuk interface{}).
// Artinya fungsi ini bisa menerima masukan dalam tipe data APA PUN (int, string, struct, bool, dll).
func CekTipeData(data any) {
	// Type Switch: Memeriksa tipe data asli dari variabel 'data' saat runtime
	switch v := data.(type) {
	case int:
		fmt.Printf("Data ini adalah Integer : %d\n", v)
	case string:
		fmt.Printf("Data ini adalah String  : %s\n", v)
	case Persegi:
		// Jika tipenya Persegi, kita bisa mengakses field .Sisi milik Persegi
		fmt.Printf("Data ini adalah Struct Persegi dengan Sisi: %.1f\n", v.Sisi)
	default:
		// Menangani tipe data lain yang tidak didefinisikan secara khusus di atas
		fmt.Printf("Tipe data tidak dikenal : %T\n", v)
	}
}

// ==========================================
// 6. FUNGSI UTAMA (MAIN)
// ==========================================
func main() {
	fmt.Println("=== 1. DEMO CUSTOM INTERFACE ===")

	// Inisialisasi instance dari struct Persegi dan Lingkaran
	p := Persegi{Sisi: 5}
	l := Lingkaran{JariJari: 7}

	// Keduanya bisa dimasukkan ke fungsi CetakInfoBangun() yang sama
	// karena keduanya sudah memenuhi spesifikasi interface HitungDatar.
	fmt.Println("Bangun 1 (Persegi):")
	CetakInfoBangun(p)

	fmt.Println("Bangun 2 (Lingkaran):")
	CetakInfoBangun(l)

	fmt.Println("=== 2. DEMO EMPTY INTERFACE (any) & TYPE SWITCH ===")
	// Memanggil fungsi CekTipeData() dengan berbagai jenis nilai/tipe data
	CekTipeData(100)                // int
	CekTipeData("Golang Interface") // string
	CekTipeData(p)                  // Struct Persegi
	CekTipeData(true)               // bool (masuk ke default)
}
