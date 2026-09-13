package main

import "fmt"

func main() {
	nilai := 80
	absen := 90

	// ==========================================
	// 1. IF SEDERHANA
	// ==========================================
	// Hanya diproses jika kondisi (nilai >= 70) bernilai true
	if nilai >= 70 {
		fmt.Println("1. Selamat! Anda lulus ujian.")
	}

	// ==========================================
	// 2. IF - ELSE
	// ==========================================
	// Pilih salah satu: blok 'if' jika true, blok 'else' jika false
	if nilai >= 90 {
		fmt.Println("2. Nilai Anda sempurna!")
	} else {
		fmt.Println("2. Nilai Anda cukup, tapi belum mencapai angka 90.")
	}

	// ==========================================
	// 3. IF - ELSE IF - ELSE (Multi Kondisi)
	// ==========================================
	// Mengecek kondisi satu per satu dari atas ke bawah
	nilaiUjian := 75

	if nilaiUjian >= 85 {
		fmt.Println("3. Grade: A")
	} else if nilaiUjian >= 75 {
		fmt.Println("3. Grade: B") // Baris ini yang akan dieksekusi
	} else if nilaiUjian >= 60 {
		fmt.Println("3. Grade: C")
	} else {
		fmt.Println("3. Grade: D")
	}

	// ==========================================
	// 4. IF DENGAN LOGIKA KOMBINASI (&& dan ||)
	// ==========================================
	// Menggunakan operator logika AND (&&) untuk mengecek 2 syarat sekaligus
	if nilai >= 75 && absen >= 80 {
		fmt.Println("4. Anda Lulus: Nilai dan Kehadiran memenuhi syarat!")
	} else {
		fmt.Println("4. Anda Tidak Lulus.")
	}

	// ==========================================
	// 5. IF DENGAN SHORT STATEMENT (Khas Golang!)
	// ==========================================
	// Format: if [statement_singkat]; [kondisi] { ... }
	// Variabel 'panjang' HANYA BISA DIAKSES di dalam blok if/else ini saja
	if panjang := len("Golang"); panjang > 5 {
		fmt.Printf("5. Kata memiliki %d karakter (Lebih dari 5)\n", panjang)
	} else {
		fmt.Printf("5. Kata hanya memiliki %d karakter\n", panjang)
	}

	// ==========================================
	// 6. NESTED IF (If Bersarang)
	// ==========================================
	// Ada pengecekan 'if' di dalam 'if'
	umur := 20
	punyaSIM := true

	if umur >= 17 {
		if punyaSIM {
			fmt.Println("6. Boleh mengendarai kendaraan.")
		} else {
			fmt.Println("6. Cukup umur, tetapi belum boleh mengendarai karena tidak punya SIM.")
		}
	} else {
		fmt.Println("6. Belum cukup umur untuk mengendarai.")
	}
}
