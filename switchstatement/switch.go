package main

import "fmt"

func main() {
	// ==========================================
	// 1. SWITCH STANDAR
	// ==========================================
	hari := "Selasa"

	switch hari {
	case "Senin":
		fmt.Println("1. Hari pertama kerja")
	case "Selasa":
		fmt.Println("1. Hari kedua kerja") // Baris ini dieksekusi
	case "Rabu":
		fmt.Println("1. Hari tengah minggu")
	default:
		fmt.Println("1. Hari lainnya")
	}

	// ==========================================
	// 2. SWITCH MULTIPLE CASE (Banyak Nilai per Case)
	// ==========================================
	bulan := "Januari"

	switch bulan {
	case "Desember", "Januari", "Februari":
		fmt.Println("2. Musim Hujan / Dingin") // Jan masuk ke sini
	case "Juni", "Juli", "Agustus":
		fmt.Println("2. Musim Kemarau")
	default:
		fmt.Println("2. Musim Pancaroba")
	}

	// ==========================================
	// 3. SWITCH TANPA KONDISI (Mirip If-Else)
	// ==========================================
	nilai := 85

	switch {
	case nilai >= 90:
		fmt.Println("3. Grade A")
	case nilai >= 80:
		fmt.Println("3. Grade B") // 85 >= 80, dieksekusi
	case nilai >= 70:
		fmt.Println("3. Grade C")
	default:
		fmt.Println("3. Grade D")
	}

	// ==========================================
	// 4. SWITCH DENGAN SHORT STATEMENT
	// ==========================================
	// Variabel 'panjang' hanya berlaku di dalam area switch ini saja
	switch panjang := len("Golang"); {
	case panjang > 5:
		fmt.Printf("4. Kata memiliki %d karakter (Lebih dari 5)\n", panjang)
	default:
		fmt.Printf("4. Kata memiliki %d karakter (Kurang dari atau sama dengan 5)\n", panjang)
	}

	// ==========================================
	// 5. SWITCH DENGAN KEYWORD fallthrough
	// ==========================================
	// 'fallthrough' memaksa program mengeksekusi case di bawahnya
	// TANPA memedulikan apakah kondisinya cocok atau tidak.
	nomor := 1

	switch nomor {
	case 1:
		fmt.Println("5. Angka 1")
		fallthrough // Memaksa lanjut ke case 2
	case 2:
		fmt.Println("5. Angka 2 (Muncul karena fallthrough dari case 1)")
	case 3:
		fmt.Println("5. Angka 3")
	}
}
