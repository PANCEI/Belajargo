package main

import "fmt"

type Siswa struct {
	Nama  string
	Nilai int
}

// Fungsi biasa (Pass by Value): Data di-copy, variabel asli TIDAK berubah
func ubahNilaiBiasa(n int) {
	n = 100
}

// Fungsi dengan Pointer (Pass by Reference): Mengubah variabel asli via alamat memori
func ubahNilaiPointer(n *int) {
	*n = 100 // Dereferencing: Mengubah nilai di alamat memori n
}

// Pointer pada Struct (Method/Fungsi)
func ubahSiswa(s *Siswa, namaBaru string, nilaiBaru int) {
	s.Nama = namaBaru // Go otomatis melakukan dereferencing pada struct (tidak perlu (*s).Nama)
	s.Nilai = nilaiBaru
}

func main() {
	// ==========================================
	// 1. DASAR POINTER (& dan *)
	// ==========================================
	angka := 42
	var ptr *int = &angka // ptr menyimpan alamat memori dari 'angka'

	fmt.Println("1. Nilai variabel 'angka'   :", angka)
	fmt.Println("   Alamat memori 'angka' (&):", &angka)
	fmt.Println("   Nilai variabel 'ptr'     :", ptr)  // Berisi alamat memori yang sama
	fmt.Println("   Nilai via Pointer (*)    :", *ptr) // Dereferencing: Mengambil nilai 42

	// Mengubah nilai via pointer
	*ptr = 77
	fmt.Println("\n   [Setelah *ptr = 77]")
	fmt.Println("   Nilai 'angka' ikut berubah :", angka) // Output: 77

	// ==========================================
	// 2. POINTER PADA FUNGSI (Pass by Reference)
	// ==========================================
	val := 10

	ubahNilaiBiasa(val)
	fmt.Println("\n2. Setelah ubahNilaiBiasa()  :", val) // Masih 10

	ubahNilaiPointer(&val)                             // Mengirimkan alamat memori &val
	fmt.Println("   Setelah ubahNilaiPointer():", val) // Berubah jadi 100

	// ==========================================
	// 3. POINTER PADA STRUCT
	// ==========================================
	s1 := Siswa{Nama: "Budi", Nilai: 75}
	fmt.Println("\n3. Data Siswa Awal           :", s1)

	// Mengirimkan pointer struct &s1
	ubahSiswa(&s1, "Budi Santoso", 95)
	fmt.Println("   Data Siswa Setelah Diubah :", s1)

	// ==========================================
	// 4. POINTER KOSONG (Nil Pointer)
	// ==========================================
	var ptrKosong *int // Zero value dari pointer adalah nil
	fmt.Println("\n4. Nilai Pointer Kosong      :", ptrKosong)

	if ptrKosong == nil {
		fmt.Println("   Pointer belum menunjuk ke alamat memori manapun.")
	}
}
