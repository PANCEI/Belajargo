package main

import (
	"fmt"
	"time"
)

// Fungsi Koki (berjalan sebagai Goroutine)
// ch <-chan string artinya channel ini KHUSUS untuk MENERIMA/MEMBACA data (receive-only)
func koki(namaKoki string, ch <-chan string, selesai chan<- bool) {
	// Loop akan terus berjalan mengambil pesanan dari channel selama channel belum ditutup
	for pesanan := range ch {
		fmt.Printf("👨‍🍳 [%s] Mulai memasak: %s...\n", namaKoki, pesanan)

		// Mensimulasikan waktu memasak selama 1.5 detik
		time.Sleep(1500 * time.Millisecond)

		fmt.Printf("✅ [%s] Selesai memasak: %s!\n", namaKoki, pesanan)
	}

	// Memberi tahu bahwa koki ini sudah selesai bekerja setelah channel ditutup
	selesai <- true
}

func main() {
	// 1. Membuat Unbuffered Channel untuk data pesanan makanan (tipe string)
	// Kapasitas channel ini adalah 0 (harus ada yang terima secara langsung saat dikirim)
	channelPesanan := make(chan string)

	// 2. Membuat channel khusus sinyal selesai (tipe bool)
	sinyalSelesai := make(chan bool)

	fmt.Println(" restoran buka! Menerima pesanan...\n")

	// 3. Menjalankan Goroutine Koki di background
	go koki("Koki Budi", channelPesanan, sinyalSelesai)

	// 4. Mengirim data pesanan ke dalam channel dari fungsi utama (main)
	daftarMenu := []string{"Nasi Goreng Spesial", "Mie Ayam Pangsit", "Es Teh Manis"}

	for _, menu := range daftarMenu {
		fmt.Printf("📝 Pelanggan memesan: %s\n", menu)

		// MENGIRIM data ke channel menggunakan operator `<-`
		channelPesanan <- menu

		// Jeda sejenak sebelum membuat pesanan berikutnya
		time.Sleep(500 * time.Millisecond)
	}

	// 5. PENTING: Menutup channel pesanan setelah semua pesanan selesai dikirim
	// Ini memberi tahu goroutine koki bahwa antrean pesanan sudah habis dan loop `range` bisa berhenti.
	close(channelPesanan)

	// 6. Menunggu sinyal bahwa koki benar-benar telah selesai memproses semuanya
	<-sinyalSelesai

	fmt.Println("\n Semua pesanan selesai! Restoran tutup.")
}
