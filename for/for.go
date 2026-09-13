package main

import "fmt"

func main() {
	// ==========================================
	// 1. FOR LOOP STANDAR
	// ==========================================
	// Format: for [init]; [kondisi]; [post]
	fmt.Println("--- 1. For Loop Standar ---")
	for i := 1; i <= 3; i++ {
		fmt.Println("Angka ke-", i)
	}

	// ==========================================
	// 2. FOR LOOP MIRIP WHILE
	// ==========================================
	// Hanya berisi kondisi saja
	fmt.Println("\n--- 2. For Loop Mirip While ---")
	counter := 1
	for counter <= 3 {
		fmt.Println("Counter ke-", counter)
		counter++
	}

	// ==========================================
	// 3. FOR RANGE (Perulangan Array / Slice)
	// ==========================================
	// Mengambil indeks dan nilainya dari sebuah data koleksi
	fmt.Println("\n--- 3. For Range pada Slice ---")
	buah := []string{"Apel", "Jeruk", "Mangga"}

	for index, value := range buah {
		fmt.Printf("Indeks %d : %s\n", index, value)
	}

	// Trik: Gunakan blank identifier '_' jika tidak butuh indeksnya
	fmt.Println("\n--- For Range (Tanpa Indeks) ---")
	for _, v := range buah {
		fmt.Println("Nama Buah:", v)
	}

	// ==========================================
	// 4. KONTROL LOOP: CONTINUE & BREAK
	// ==========================================
	// 'continue' = Lewati iterasi saat ini, lanjut ke iterasi berikutnya
	// 'break'    = Hentikan perulangan sepenuhnya
	fmt.Println("\n--- 4. For dengan Continue dan Break ---")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // Jika genap, lewati (tidak dicetak)
		}

		if i > 5 {
			break // Jika angka melebihi 5, hentikan perulangan
		}

		fmt.Println("Angka ganjil:", i) // Hanya akan mencetak 1, 3, 5
	}

	// ==========================================
	// 5. INFINITE LOOP (Perulangan Tak Terbatas)
	// ==========================================
	// for tanpa kondisi akan berjalan terus sampai dipaksa keluar via break
	fmt.Println("\n--- 5. Infinite Loop dengan Break ---")
	n := 1
	for {
		fmt.Println("Loop tak terbatas ke-", n)
		if n == 2 {
			fmt.Println("Dihentikan secara paksa dengan break!")
			break
		}
		n++
	}
}
