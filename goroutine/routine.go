package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulasi fungsi untuk mengunduh file
// Kita menggunakan channel (ch) untuk mengirim hasil ke goroutine utama
func downloadFile(fileName string, wg *sync.WaitGroup, ch chan<- string) {
	// Pastikan Done() dipanggil saat fungsi selesai agar WaitGroup tahu
	defer wg.Done()

	fmt.Printf("Mulai mengunduh: %s...\n", fileName)

	// Mensimulasikan proses I/O yang memakan waktu (misal: 2 detik)
	time.Sleep(2 * time.Second)

	// Mengirim pesan sukses ke channel
	ch <- fmt.Sprintf("File %s berhasil diunduh!", fileName)
}

func main() {
	// Waktu mulai eksekusi untuk menghitung durasi total
	startTime := time.Now()

	// 1. Membuat WaitGroup untuk menunggu semua Goroutine selesai
	var wg sync.WaitGroup

	// 2. Membuat Channel untuk menerima data/pesan dari goroutine
	// Kapasitas channel disesuaikan dengan jumlah pekerjaan (3 file)
	ch := make(chan string, 3)

	// Daftar file yang akan diunduh
	files := []string{"video.mp4", "dokumen.pdf", "foto.jpg"}

	// 3. Menjalankan Goroutine menggunakan keyword 'go' di dalam perulangan
	for _, file := range files {
		// Menambahkan counter ke WaitGroup setiap kali ada goroutine baru
		wg.Add(1)

		// Keyword 'go' membuat fungsi downloadFile berjalan secara asynchronous (concurrent)
		go downloadFile(file, &wg, ch)
	}

	// 4. Menunggu semua goroutine selesai di background
	// Wait() akan menahan fungsi main agar tidak berhenti sebelum semua wg.Done() dipanggil
	wg.Wait()

	// Menutup channel karena semua proses pengiriman data telah selesai
	close(ch)

	// 5. Membaca hasil dari channel
	fmt.Println("\n--- Hasil Unduhan ---")
	for result := range ch {
		fmt.Println(result)
	}

	// Menghitung total waktu eksekusi
	duration := time.Since(startTime)
	fmt.Printf("\nSemua proses selesai dalam waktu: %v\n", duration)
}
