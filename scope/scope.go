package main

import "fmt"

// ==========================================
// 1. PACKAGE SCOPE (Global di tingkat package)
// ==========================================
// Variabel ini bisa diakses di manapun di dalam fungsi-fungsi package main
var appName string = "Aplikasi Kasir"

// Exported Variable (Bisa diimpor package lain jika dipisah file)
var VersiApp string = "v1.0.0"

func main() {
	// Memanggil variabel Package Scope
	fmt.Println("1. App Name (Package Scope) :", appName)

	// ==========================================
	// 2. LOCAL / FUNCTION SCOPE
	// ==========================================
	// Variabel 'pesanLocal' hanya ada di dalam fungsi main()
	pesanLocal := "Halo dari fungsi main"
	fmt.Println("2. Local Scope              :", pesanLocal)

	// ==========================================
	// 3. BLOCK SCOPE (Dalam if / for / switch)
	// ==========================================
	stok := 5

	if stok > 0 {
		// Variabel 'statusStok' dan 'diskon' HANYA HIDUP di dalam blok if ini
		statusStok := "Tersedia"
		diskon := 10

		fmt.Println("3. Block Scope (dalam if)   :", statusStok)
		fmt.Printf("   Diskon: %d%%\n", diskon)
	}

	// ERROR jika baris di bawah diaktifkan!
	// fmt.Println(statusStok) // panic: undefined: statusStok

	// ==========================================
	// 4. VARIABLE SHADOWING (Penutupan Scope)
	// ==========================================
	// Menggunakan nama variabel yang sama dengan Package Scope di dalam fungsi local
	appName := "Aplikasi Toko Online" // Menutup (shadow) appName global khusus di fungsi ini
	fmt.Println("4. Variable Shadowing       :", appName)

	// Memanggil fungsi lain untuk membuktikan Scope
	prosesTransaksi()
}

func prosesTransaksi() {
	// Memanggil Package Scope (Masih bisa diakses)
	fmt.Println("\n--- Di Dalam Fungsi prosesTransaksi() ---")
	// appName di sini kembali mengakses variabel Global "Aplikasi Kasir"
	fmt.Println("Package Scope appName       :", appName)

	// ERROR jika baris di bawah diaktifkan!
	// fmt.Println(pesanLocal) // panic: undefined: pesanLocal (karena milik main)
}
