/*
buffeed channel digunakan untuk komunikasi antar go routine yang datanya bersifat sementara
bytes.buffer untuk menangangi operasi string dank byte secara efisien
bytes dalam i/o digunakan untuk membaca dile atau menulis file atau koneksi jaringan
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// channel := make(chan int, 4)
	// channel <- 1
	// channel <- 2
	// channel <- 3
	// channel <- 4

	// // jalankan buffer  chhannel
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)

	// penanaganan bytes buffer
	// var bytes bytes.Buffer
	// bytes.WriteString("chua")
	// bytes.WriteString("baci")
	// bytes.WriteString("bac0")
	// fmt.Println(bytes.String())

	// penanganana bytes dalam io
	file, error := os.Open("file_kita.text")
	if error != nil {
		fmt.Println("error", error)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, error := reader.ReadString('\n')
		if error != nil {
			fmt.Println("error", error)
			break
		}
		fmt.Println(line)
	}
}
