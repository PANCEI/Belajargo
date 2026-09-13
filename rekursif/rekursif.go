package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ============================================================================
// 1. FUNGSI REKURSIF FAKTORIAL
// ============================================================================
// Menghitung n! secara matematis (misal: 4! = 4 x 3 x 2 x 1 = 24).
func Faktorial(n int) int {
	// BASE CASE (Syarat Berhenti):
	// Jika n = 0 atau 1, langsung kembalikan angka 1 agar rekursi tidak berjalan selamanya.
	if n <= 1 {
		return 1
	}

	// RECURSIVE CASE (Pemanggilan Diri):
	// n dikalikan dengan hasil pemanggilan Faktorial(n - 1)
	return n * Faktorial(n-1)
}

// ============================================================================
// 2. FUNGSI REKURSIF HITUNG MUNDUR (TRACER)
// ============================================================================
// Mencetak angka secara menurun dari n sampai 1 untuk melacak alur rekursi.
func HitungMundur(n int) {
	// BASE CASE:
	// Berhenti ketika angka menyentuh 0 atau minus
	if n <= 0 {
		fmt.Println("Selesai! 🚀")
		return // Keluar dari fungsi
	}

	// Tampilkan angka saat ini
	fmt.Printf("Angka: %d\n", n)

	// RECURSIVE CASE:
	// Memanggil fungsi dirinya sendiri dengan mengurangi nilai n sebesar 1
	HitungMundur(n - 1)
}

// ============================================================================
// 3. FUNGSI REKURSIF TRAVERSAL FOLDER (REAL CASE)
// ============================================================================
// Membaca direktori beserta seluruh isi sub-folder di dalamnya secara hierarki.
func TampilkanStrukturFolder(pathFolder string, kedalaman int) {
	// Step A: Baca isi direktori dari path yang diberikan
	entri, err := os.ReadDir(pathFolder)
	if err != nil {
		fmt.Printf("Gagal membaca direktori %s: %v\n", pathFolder, err)
		return
	}

	// BASE CASE (Implicit):
	// Jika folder kosong (len(entri) == 0), perulangan for di bawah tidak akan
	// berjalan dan fungsi otomatis selesai (kembali ke caller).
	if len(entri) == 0 {
		return
	}

	// Buat spasi/indentasi visual berdasarkan tingkat kedalaman sub-folder
	indentasi := strings.Repeat("│   ", kedalaman)

	// Step B: Iterasi setiap elemen yang ada di dalam folder tersebut
	for _, item := range entri {
		// Gabungkan path asal dengan nama file/folder
		pathLengkap := filepath.Join(pathFolder, item.Name())

		if item.IsDir() {
			// Jika elemen adalah FOLDER:
			fmt.Printf("%s├── 📁 [%s]\n", indentasi, item.Name())

			// RECURSIVE CASE:
			// Panggil fungsi ini lagi untuk "masuk" ke dalam sub-folder tersebut.
			// Nilai 'kedalaman' ditambah 1 untuk menggeser cetakan ke kanan.
			TampilkanStrukturFolder(pathLengkap, kedalaman+1)
		} else {
			// Jika elemen adalah FILE biasa:
			fmt.Printf("%s├── 📄 %s\n", indentasi, item.Name())
		}
	}
}

// ============================================================================
// FUNGSI UTAMA (MAIN)
// ============================================================================
func main() {
	// --------------------------------------------------
	// BAGIAN 1: DEMO HITUNG MUNDUR
	// --------------------------------------------------
	fmt.Println("=== 1. DEMO REKURSIF HITUNG MUNDUR ===")
	HitungMundur(3)

	// --------------------------------------------------
	// BAGIAN 2: DEMO FAKTORIAL
	// --------------------------------------------------
	fmt.Println("\n=== 2. DEMO REKURSIF FAKTORIAL ===")
	angka := 4
	hasil := Faktorial(angka)
	fmt.Printf("Hasil faktorial dari %d! adalah : %d\n", angka, hasil)

	// --------------------------------------------------
	// BAGIAN 3: DEMO TRAVERSAL STRUKTUR FOLDER
	// --------------------------------------------------
	fmt.Println("\n=== 3. DEMO TRAVERSAL STRUKTUR FOLDER ===")

	// Menyiapkan folder sementara untuk keperluan uji coba
	folderDemo := "./folder_proyek"
	os.MkdirAll(folderDemo+"/src/controllers", 0755)
	os.WriteFile(folderDemo+"/main.go", []byte("package main"), 0644)
	os.WriteFile(folderDemo+"/src/controllers/user.go", []byte("package controllers"), 0644)
	os.WriteFile(folderDemo+"/README.md", []byte("# Documentation"), 0644)

	// Panggil fungsi rekursif direktori dimulai dari tingkat akar (kedalaman 0)
	fmt.Printf("📁 [%s]\n", folderDemo)
	TampilkanStrukturFolder(folderDemo, 0)

	// Bersihkan folder sementara setelah eksekusi selesai
	os.RemoveAll(folderDemo)
}
