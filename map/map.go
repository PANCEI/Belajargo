package main

import "fmt"

func main() {
	// ==========================================
	// 1. DEKLARASI MAP LITERALS
	// ==========================================
	// Key bertipe string, Value bertipe int
	hargaBuah := map[string]int{
		"Apel":   15000,
		"Jeruk":  12000,
		"Mangga": 20000,
	}
	fmt.Println("1. Map Harga Buah      :", hargaBuah)

	// ==========================================
	// 2. MENAMBAH DAN MENGUBAH ELEMEN MAP
	// ==========================================
	hargaBuah["Pisang"] = 10000 // Menambah Key-Value baru
	hargaBuah["Apel"] = 18000   // Mengubah nilai dari Key "Apel" yang sudah ada

	fmt.Println("\n2. Setelah Tambah/Ubah :", hargaBuah)

	// ==========================================
	// 3. MENERIMA/MEMBACA ELEMEN MAP
	// ==========================================
	fmt.Println("\n3. Harga Jeruk         : Rp", hargaBuah["Jeruk"])

	// ==========================================
	// 4. MEMERIKSA KEBERADAAN KEY (Comma Ok Idiom)
	// ==========================================
	// Mengambil nilai sekaligus status ada/tidaknya key tersebut
	val, isExist := hargaBuah["Durian"]
	if isExist {
		fmt.Println("4. Durian ditemukan dengan harga:", val)
	} else {
		fmt.Println("4. Key 'Durian' TIDAK DITEMUKAN di dalam map!")
	}

	// ==========================================
	// 5. MENGHAPUS ELEMEN (Fungsi delete())
	// ==========================================
	// Format: delete(namaMap, keyYangDihapus)
	delete(hargaBuah, "Jeruk")
	fmt.Println("\n5. Setelah 'Jeruk' Dihapus:", hargaBuah)

	// ==========================================
	// 6. DEKLARASI MAP DENGAN make()
	// ==========================================
	userAge := make(map[string]int)
	userAge["Budi"] = 25
	userAge["Siti"] = 22
	fmt.Println("\n6. Map userAge (make())   :", userAge)

	// ==========================================
	// 7. PERULANGAN PADA MAP (FOR RANGE)
	// ==========================================
	// Catatan: Urutan iterasi map di Go bersifat ACAK (Random)
	fmt.Println("\n7. Iterasi Map hargaBuah:")
	for key, value := range hargaBuah {
		fmt.Printf("   Buah: %-7s | Harga: Rp %d\n", key, value)
	}
}
