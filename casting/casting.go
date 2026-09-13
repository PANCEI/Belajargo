package main

import (
	"fmt"
	"strconv" // Package khusus konversi String ke Tipe Data Lain & Sebaliknya
)

func main() {
	// ==========================================
	// 1. KONVERSI SESAMA INTEGER (Beda Ukuran Bit)
	// ==========================================
	var angka32 int32 = 100
	// Mengubah int32 menjadi int64
	var angka64 int64 = int64(angka32)
	fmt.Printf("1. Int32 ke Int64 : %d | Tipe: %T\n", angka64, angka64)

	// HATI-HATI (Overflow): Jika mengubah angka besar ke tipe yang lebih kecil
	var angkaBesar int16 = 300
	var angkaKecil int8 = int8(angkaBesar)          // int8 maksimal 127
	fmt.Println("   Hasil Overflow  :", angkaKecil) // Hasilnya akan kacau/berputar (-44)

	// ==========================================
	// 2. KONVERSI INTEGER DAN FLOAT
	// ==========================================
	var nilaiInt int = 42
	var nilaiFloat float64 = float64(nilaiInt) // int ke float64
	fmt.Printf("2. Int ke Float   : %.2f | Tipe: %T\n", nilaiFloat, nilaiFloat)

	var piFloat float64 = 3.99
	var piInt int = int(piFloat) // float64 ke int (Desimal dibuang, BUKAN dibulatkan)
	fmt.Printf("   Float ke Int   : %d | Tipe: %T (Angka di belakang koma hilang)\n", piInt, piInt)

	// ==========================================
	// 3. KONVERSI BYTE / RUNE KE STRING
	// ==========================================
	// Di Go, string adalah kumpulan byte / rune
	var byteVal byte = 'A' // ASCII 65
	var strVal string = string(byteVal)
	fmt.Printf("3. Byte/Rune ke String : %s\n", strVal)

	// Slice of bytes ke String
	bytes := []byte{71, 111, 108, 97, 110, 103} // ASCII untuk "Golang"
	teks := string(bytes)
	fmt.Printf("   Byte Slice ke String: %s\n", teks)

	// ==========================================
	// 4. KONVERSI STRING KE ANGKA & SEBALIKNYA (Package strconv)
	// ==========================================
	// a. String ke Integer (strconv.Atoi)
	teksAngka := "123"
	angkaBiasa, err := strconv.Atoi(teksAngka)
	if err == nil {
		fmt.Println("4. String ke Int  :", angkaBiasa+7) // 123 + 7 = 130
	}

	// b. Integer ke String (strconv.Itoa)
	nomor := 500
	teksTujuan := strconv.Itoa(nomor)
	fmt.Printf("   Int ke String  : %s | Tipe: %T\n", teksTujuan, teksTujuan)

	// c. String ke Float (strconv.ParseFloat)
	teksFloat := "3.14"
	hasilFloat, _ := strconv.ParseFloat(teksFloat, 64)
	fmt.Printf("   String ke Float: %.2f | Tipe: %T\n", hasilFloat, hasilFloat)
}
