package main

import "fmt"

func main() {
	// ==========================================
	// 1. RANGE PADA SLICE / ARRAY
	// ==========================================
	fmt.Println("=== 1. Range pada Slice ===")
	buah := []string{"Apel", "Jeruk", "Mangga"}

	// Mengambil Indeks dan Nilai
	for index, value := range buah {
		fmt.Printf("Indeks %d: %s\n", index, value)
	}

	// Mengabaikan Indeks menggunakan Blank Identifier (_)
	fmt.Print("Hanya Nilai: ")
	for _, v := range buah {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// ==========================================
	// 2. RANGE PADA MAP
	// ==========================================
	fmt.Println("\n=== 2. Range pada Map ===")
	modal := map[string]int{
		"Budi": 5000,
		"Siti": 7500,
	}

	// Mengambil Key dan Value (Urutan iterasi map acak)
	for key, val := range modal {
		fmt.Printf("Key: %-4s | Value: %d\n", key, val)
	}

	// Hanya mengiterasi Key saja
	fmt.Print("Daftar Nama (Key): ")
	for key := range modal {
		fmt.Print(key, " ")
	}
	fmt.Println()

	// ==========================================
	// 3. RANGE PADA STRING (Iterasi Unicode/Rune)
	// ==========================================
	fmt.Println("\n=== 3. Range pada String ===")
	teks := "Go 🚀"

	// range mengurai string menjadi rune (UTF-8), bukan sekadar byte
	for i, char := range teks {
		fmt.Printf("Indeks Byte: %d | Karakter: %c | Rune: %U\n", i, char, char)
	}

	// ==========================================
	// 4. RANGE PADA CHANNEL
	// ==========================================
	fmt.Println("\n=== 4. Range pada Channel ===")
	ch := make(chan int, 3)
	ch <- 10
	ch <- 20
	ch <- 30
	close(ch) // Channel wajib ditutup agar range berhenti

	// Reading values until channel is closed
	for val := range ch {
		fmt.Println("Diterima dari channel:", val)
	}

	// ==========================================
	// 5. RANGE PADA INTEGER (Fitur Go 1.22+)
	// ==========================================
	fmt.Println("\n=== 5. Range pada Integer ===")
	// Berjalan dari 0 hingga n-1 (0 sampai 2)
	for i := range 3 {
		fmt.Println("Iterasi ke-", i)
	}
}
