package main

import "fmt"

func main() {
	// ==========================================
	// 1. DEFER STATEMENT (Dieksekusi di Akhir Fungsi)
	// ==========================================
	// 'defer' menunda eksekusi perintah ini sampai fungsi main() selesai.
	defer fmt.Println("--- [Selesai] Ini dicetak paling akhir karena 'defer' ---")

	fmt.Println("=== DEMO CONTROL STATEMENTS DI GO ===")

	// ==========================================
	// 2. BRANCHING: IF - ELSE IF - ELSE
	// ==========================================
	nilai := 85

	// Short statement di dalam if
	if lulus := nilai >= 75; lulus {
		fmt.Println("1. Status: Lulus Ujian")
	} else {
		fmt.Println("1. Status: Tidak Lulus Ujian")
	}

	// ==========================================
	// 3. BRANCHING: SWITCH
	// ==========================================
	kategori := "B"

	switch kategori {
	case "A":
		fmt.Println("2. Kinerja: Sangat Baik")
	case "B":
		fmt.Println("2. Kinerja: Baik") // Baris ini dieksekusi
	case "C":
		fmt.Println("2. Kinerja: Cukup")
	default:
		fmt.Println("2. Kinerja: Kurang")
	}

	// ==========================================
	// 4. LOOPING: FOR WITH CONTINUE & BREAK
	// ==========================================
	fmt.Println("\n--- Perulangan Angka Ganjil (Max 5) ---")
	for i := 1; i <= 10; i++ {
		// Mengabaikan angka genap
		if i%2 == 0 {
			continue // Lanjut ke iterasi berikutnya
		}

		// Menghentikan loop jika angka lebih besar dari 5
		if i > 5 {
			break // Keluar dari loop 'for'
		}

		fmt.Printf("Angka: %d\n", i) // Mencetak 1, 3, 5
	}

	// ==========================================
	// 5. JUMP STATEMENT: GOTO
	// ==========================================
	// 'goto' melompat langsung ke label tertentu di dalam kode
	fmt.Println("\n--- Menggunakan Goto ---")
	statusServer := "ERROR"

	if statusServer == "ERROR" {
		goto HandleError // Melompat ke label HandleError
	}

	fmt.Println("Proses server normal...") // Baris ini dilompati

HandleError:
	fmt.Println("3. Terjadi kesalahan pada server! Menangani error...")
}
