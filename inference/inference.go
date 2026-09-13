package main

import "fmt"

func main() {
	// ==========================================
	// 1. INFERENSI TIPE DATA DASAR (:=)
	// ==========================================
	// Go otomatis mendeteksi tipe data dari nilai di sebelah kanan '='

	nama := "Budi"     // Otomatis dianggap 'string'
	umur := 25         // Otomatis dianggap 'int'
	tinggi := 175.5    // Otomatis dianggap 'float64'
	isMenikah := false // Otomatis dianggap 'bool'

	// %T pada Printf digunakan untuk mengecek Tipe Data variabel secara langsung
	fmt.Printf("1. nama      : %-7v | Tipe: %T\n", nama, nama)
	fmt.Printf("   umur      : %-7v | Tipe: %T\n", umur, umur)
	fmt.Printf("   tinggi    : %-7v | Tipe: %T\n", tinggi, tinggi)
	fmt.Printf("   isMenikah : %-7v | Tipe: %T\n", isMenikah, isMenikah)

	// ==========================================
	// 2. INFERENSI DENGAN KATA KUNCI 'var'
	// ==========================================
	// Kamu juga bisa memakai 'var' tanpa menuliskan tipe data
	var kota = "Jakarta" // Otomatis dianggap 'string'
	var gaji = 5000000.0 // Otomatis dianggap 'float64'

	fmt.Printf("\n2. kota      : %-7v | Tipe: %T\n", kota, kota)
	fmt.Printf("   gaji      : %-7v | Tipe: %T\n", gaji, gaji)

	// ==========================================
	// 3. INFERENSI PADA KOMPLEKS & STRUKTUR DATA
	// ==========================================
	angkaKompleks := 3 + 4i          // Otomatis 'complex128'
	daftarAngka := []int{10, 20, 30} // Otomatis '[]int' (slice of int)
	dataRune := 'A'                  // Otomatis 'int32' / rune (ASCII 65)

	fmt.Printf("\n3. c128      : %-7v | Tipe: %T\n", angkaKompleks, angkaKompleks)
	fmt.Printf("   slice     : %-7v | Tipe: %T\n", daftarAngka, daftarAngka)
	fmt.Printf("   rune      : %-7v | Tipe: %T\n", dataRune, dataRune)

	// ==========================================
	// 4. INFERENSI HASIL DARI FUNGSI/EKSPRESI
	// ==========================================
	a := 10
	b := 3.5
	// Hasil ekspresi (int di-cast ke float + float) otomatis jadi float64
	total := float64(a) + b
	fmt.Printf("\n4. total     : %-7v | Tipe: %T\n", total, total)
}
