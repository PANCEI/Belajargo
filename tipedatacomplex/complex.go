package main

import (
	"fmt"
	"math/cmplx" // Package khusus untuk operasi matematika bilangan kompleks
)

func main() {
	// 1. Deklarasi Eksplisit complex64
	// Harus dituliskan tipe data 'complex64' secara jelas
	var c64 complex64 = complex(2.5, 4.5)
	fmt.Printf("c64  : %v | Tipe: %T\n", c64, c64)

	// 2. Deklarasi Eksplisit complex128
	var c128 complex128 = complex(2.5, 4.5)
	fmt.Printf("c128 : %v | Tipe: %T\n", c128, c128)

	// 3. Menggunakan Sintaks Literal Biasa (Otomatis jadi complex128)
	cAuto := 10 + 20i
	fmt.Printf("cAuto: %v | Tipe: %T\n", cAuto, cAuto)

	// 4. Konversi Tipe Data (Casting)
	// Go TIDAK MENIZINKAN operasi langsung antara complex64 dan complex128
	// Kita harus mengonversi salah satunya terlebih dahulu
	var hasilKonversi complex128 = complex128(c64) + c128
	fmt.Println("\nHasil Penjumlahan setelah Konversi:", hasilKonversi)

	// 5. Mengambil Nilai Riil dan Imajiner dari complex64
	// Fungsi real() dan imag() mengembalikan float32 jika inputnya complex64
	var r64 float32 = real(c64)
	var i64 float32 = imag(c64)
	fmt.Printf("Komponen c64 -> Riil: %.1f, Imajiner: %.1f\n", r64, i64)

	// 6. Menggunakan Package 'math/cmplx' untuk Fungsi Matematika Lanjutan
	// Catatan: Fungsi di package 'cmplx' hanya menerima tipe 'complex128'
	z := complex(0, 3.141592653589793)             // i * pi
	euler := cmplx.Exp(z)                          // e^(i * pi)
	fmt.Println("\nRumus Euler e^(i*pi) :", euler) // Mendekati (-1 + 0i)
}
