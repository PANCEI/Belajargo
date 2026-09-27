// contect
/*
context.Background
context.TODO()
context.WithCancle
context.WithTimeOut / context.WithDeadline
context.WithValue
**/
package main

import (
	"context"
	"fmt"
	"time"
)

//	func prosesRequest(context_kita context.Context) {
//		username := context_kita.Value("username")
//		if username != nil {
//			fmt.Println("hai", username)
//		} else {
//			fmt.Println("username tidak di temukan")
//		}
//	}
func worker(context_kita context.Context, nama string) {
	for {
		select {
		case <-context_kita.Done():
			fmt.Println(nama, "stop", context_kita.Err())
			return
		default:
			fmt.Println(nama, "bekerja")
			time.Sleep(1 * time.Second)
		}
	}
}
func main() {
	// test_context := context.Background()
	// fmt.Println("Asasa", test_context)
	// context_kita, cancel := context.WithCancel(context.Background())
	// go func() {
	// 	for {
	// 		select {
	// 		case <-context_kita.Done():
	// 			fmt.Println("context selesai")
	// 			return
	// 		default:
	// 			fmt.Println("work")
	// 			time.Sleep(500 * time.Millisecond)
	// 		}
	// 	}
	// }()
	// time.Sleep(3 * time.Second)
	// cancel()
	// time.Sleep(1 * time.Second)
	// context_kita, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()

	// select {
	// case <-time.After(1 * time.Second):
	// 	fmt.Println("operasi berjalan")
	// case <-context_kita.Done():
	// 	fmt.Println("contxt kita time out", context_kita.Err())
	// }

	// context_kita := context.WithValue(context.Background(), "username", "chua")
	// prosesRequest(context_kita)

	// context dengan go routine
	parentContext := context.Background()
	context_kita, cancel := context.WithTimeout(parentContext, 3*time.Second)
	defer cancel()
	go worker(context_kita, "Pekerja 1")
	go worker(context_kita, "Pekerja 2")
	time.Sleep(4 * time.Second)

}
