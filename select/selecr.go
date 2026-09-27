package main

import (
	"fmt"
	"time"
)

// Simulasi Server A (proses agak lambat: 2 detik)
func ambilDataDariServerA(ch chan<- string) {
	time.Sleep(2 * time.Second)
	ch <- "Respon dari Server A"
}

// Simulasi Server B (proses lebih cepat: 1 detik)
func ambilDataDariServerB(ch chan<- string) {
	time.Sleep(1 * time.Second)
	ch <- "Respon dari Server B"
}

func main() {
	// Membuat dua channel string untuk menampung data dari masing-masing server
	chA := make(chan string)
	chB := make(chan string)

	fmt.Println("🔍 Sedang mencari data dari server...")

	// Menjalankan kedua fungsi server secara bersamaan (konkuren) di background
	go ambilDataDariServerA(chA)
	go ambilDataDariServerB(chB)

	// Menggunakan SELECT untuk menunggu respon tercepat atau timeout
	select {
	case hasilA := <-chA:
		fmt.Println("✅ Berhasil mendapatkan:", hasilA)

	case hasilB := <-chB:
		fmt.Println("✅ Berhasil mendapatkan:", hasilB)

	// Fitur tambahan: Timeout jika kedua server tidak merespons dalam waktu 1.5 detik
	case <-time.After(1500 * time.Millisecond):
		fmt.Println("⏰ Waktu habis (Timeout): Server terlalu lama merespons!")
	}

	fmt.Println("Program selesai.")
}
