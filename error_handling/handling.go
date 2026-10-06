/*
*
type error interface {
Error () string
}
panic adalah kondisi yang tidak dapat diatasi oleh program, biasanya terjadi karena kesalahan yang fatal, seperti pembagian dengan nol atau akses ke indeks di luar batas array. Ketika panic terjadi, program akan berhenti dan menampilkan pesan kesalahan.
recover adalah fungsi yang digunakan untuk menangani panic. Fungsi ini dapat dipanggil di dalam defer untuk memulihkan program dari kondisi panic dan melanjutkan eksekusi program. Jika recover dipanggil di luar defer, maka nil akan dikembalikan.
*/
package main

import (
	"errors"
	"fmt"
)

func pembagian(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("tidak bisa membagi dengan nol")
	}
	return a / b, nil
}
func bukaFile(namaFile string) error {
	if namaFile == "" {
		return errors.New("nama file tidak boleh kosong")
	}
	return nil
}

type myError struct {
	kode  int
	pesan string
}

func (e *myError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.kode, e.pesan)

}
func Permasalahan() error {
	return &myError{
		kode:  404,
		pesan: "Data tidak ditemukan",
	}
}

// Simulasi fungsi yang melakukan proses berbahaya (misal: memicu panic karena index out of range atau error fatal)
func hubungkanDatabase(statusKoneksi bool) {
	fmt.Println("--> Memulai koneksi database...")

	if !statusKoneksi {
		// Memicu panic secara manual (bisa juga terjadi otomatis seperti pembagian dengan nol)
		panic("Koneksi ke database gagal total! Server menolak akses.")
	}

	fmt.Println("--> Koneksi database berhasil!")
}

// Fungsi utama yang menjalankan proses dan menggunakan recover
func prosesAplikasi(statusKoneksi bool) {
	// DEFER + RECOVER:
	// Blok ini akan selalu dieksekusi ketika fungsi prosesAplikasi selesai (baik sukses maupun saat terjadi panic).
	defer func() {
		// recover() menangkap nilai yang dikirim oleh panic()
		if r := recover(); r != nil {
			fmt.Println("\n[RECOVER BERHASIL] Program aman dari crash!")
			fmt.Println("Detail Error yang ditangkap:", r)
		}
	}()

	fmt.Println("Aplikasi mulai berjalan...")

	// Memanggil fungsi yang berpotensi panic
	hubungkanDatabase(statusKoneksi)

	// Baris ini TIDAK AKAN DIEKSEKUSI jika terjadi panic di hubungkanDatabase
	fmt.Println("Aplikasi selesai memproses data.")
}
func main() {
	//   contoh membuat error costume
	error_coba := errors.New("ini adalah error")
	fmt.Println(error_coba.Error())
	// buat error dengan fmt.Errorf
	error_coba2 := fmt.Errorf("ini adalah error kedua")
	fmt.Println(error_coba2.Error())

	// contoh penggunaan error handling
	hasil, err := pembagian(10, 0)
	if err != nil {
		fmt.Println("Terjadi error:", err)
		//return
	}
	fmt.Println("Hasil pembagian:", hasil)
	erorFile := bukaFile("")
	if erorFile != nil {
		fmt.Println("Terjadi buka file error:", erorFile)
		//	return
	} else {
		fmt.Println("File berhasil dibuka")
	}
	errorCostume := Permasalahan()
	if errorCostume != nil {
		fmt.Println("Terjadi error:", errorCostume)
	} else {
		fmt.Println("Tidak ada error")
	}
	fmt.Println("=== SKenario 1: Kondisi Normal (Tanpa Panic) ===")
	prosesAplikasi(true) // statusKoneksi = true (aman)

	fmt.Println("\n---------------------------------------------------\n")

	fmt.Println("=== SKenario 2: Kondisi Terjadi Panic & Direcover ===")
	prosesAplikasi(false) // statusKoneksi = false (memicu panic)

	fmt.Println("\n=== Program Utama Selesai (Tidak Crash!) ===")
}
