package main

import "fmt"

func main() {
	// ==========================================
	// 1. DEKLARASI ARRAY KOSONG (Zero Value)
	// ==========================================
	// Membuat array integer berkapasitas 3. Default nilainya adalah [0, 0, 0]
	var angka [3]int
	angka[0] = 10 // Mengisi indeks ke-0
	angka[1] = 20 // Mengisi indeks ke-1
	angka[2] = 30 // Mengisi indeks ke-2

	fmt.Println("1. Array Angka          :", angka)
	fmt.Println("   Elemen indeks ke-1   :", angka[1])

	// ==========================================
	// 2. DEKLARASI LANGSUNG DENGAN NILAI
	// ==========================================
	buah := [4]string{"Apel", "Jeruk", "Mangga", "Pisang"}
	fmt.Println("\n2. Array Buah           :", buah)

	// Mengetahui panjang/ukuran array menggunakan len()
	fmt.Println("   Panjang Array Buah   :", len(buah))

	// ==========================================
	// 3. DEKLARASI DENGAN TITIK TIGA (...)
	// ==========================================
	// Ukuran array akan dihitung otomatis oleh Go (hasilnya 3 elemen)
	hewan := [...]string{"Kucing", "Anjing", "Kelinci"}
	fmt.Printf("\n3. Array Hewan          : %v | Tipe Data: %T\n", hewan, hewan)

	// ==========================================
	// 4. MENGUBAH ELEMEN ARRAY
	// ==========================================
	buah[0] = "Melon" // Mengubah "Apel" menjadi "Melon"
	fmt.Println("\n4. Buah Setelah Diubah  :", buah)

	// ==========================================
	// 5. PERULANGAN PADA ARRAY (FOR RANGE)
	// ==========================================
	fmt.Println("\n5. Iterasi Array Hewan:")
	for index, value := range hewan {
		fmt.Printf("   Indeks ke-%d : %s\n", index, value)
	}

	// ==========================================
	// 6. ARRAY MULTIDIMENSI (Matrix 2D)
	// ==========================================
	// Membuat matriks 2x3 (2 baris, 3 kolom)
	var matrix [2][3]int = [2][3]int{
		{1, 2, 3}, // Baris ke-0
		{4, 5, 6}, // Baris ke-1
	}

	fmt.Println("\n6. Array 2 Dimensi (Matrix):")
	fmt.Println("   Baris 0, Kolom 2 :", matrix[0][2]) // Output: 3
	fmt.Println("   Seluruh Matrix   :", matrix)
}
