/*
generic adalah package yang berisi contoh penggunaan generic di Go. Generic memungkinkan kita untuk menulis fungsi, struct, atau interface yang dapat bekerja dengan berbagai tipe data tanpa harus menulis kode yang sama berulang kali untuk setiap tipe data.
Contoh penggunaan generic di Go:
1. Fungsi Generic: Kita dapat membuat fungsi yang menerima parameter dengan tipe data yang berbeda-beda. Misalnya, kita bisa membuat fungsi untuk menukar dua nilai dari tipe data apa pun.
[T any] adalah sintaks untuk mendefinisikan parameter tipe generik. T adalah nama parameter tipe, dan any menunjukkan bahwa T dapat berupa tipe data apa pun.
2. Struct Generic: Kita juga bisa membuat struct yang memiliki field dengan tipe data generik. Ini memungkinkan kita untuk membuat struct yang fleksibel dan dapat digunakan dengan berbagai tipe data.
3. Interface Generic: Kita dapat mendefinisikan interface yang menggunakan tipe data generik, sehingga implementasi dari interface tersebut dapat bekerja dengan berbagai tipe data.
4. Batasan Tipe (Type Constraints): Kita bisa menetapkan batasan pada parameter tipe generik untuk memastikan bahwa hanya tipe data tertentu yang dapat digunakan. Misalnya, kita bisa membatasi parameter tipe hanya untuk tipe numerik atau tipe yang mengimplementasikan interface tertentu.
5. Keuntungan Menggunakan Generic: Dengan menggunakan generic, kita dapat menulis kode yang lebih bersih, lebih mudah dipelihara, dan mengurangi duplikasi kode. Ini juga meningkatkan fleksibilitas dan kemampuan untuk bekerja dengan berbagai tipe data tanpa harus menulis kode tambahan.
*/
package main

import "fmt"

func swapAngka[T any](a, b T) (T, T) {
	return b, a
}

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}
func (s *Stack[T]) Pop() T {
	if len(s.elements) == 0 {
		var zeroValue T
		return zeroValue
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element
}
func Contains[T comparable](slice []T, element T) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
		//return false
	}
	return false
}
func main() {
	x, y := 10, 20
	x, y = swapAngka(x, y)
	println("x:", x, "y:", y)

	// contoh penggunaan generic dengan tipe data string
	a, b := "Hello", "World"
	a, b = swapAngka(a, b)
	println("a:", a, "b:", b)

	// contioh untuk tipe data struct
	itegrerSatck := Stack[int]{}
	itegrerSatck.Push(1)
	itegrerSatck.Push(2)
	itegrerSatck.Push(4)
	println("Pop:", itegrerSatck.Pop())

	// string
	stringStack := Stack[string]{}
	stringStack.Push("Hello")
	stringStack.Push("World")
	println("Pop:", stringStack.Pop())

	// slice
	angkaSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(Contains(angkaSlice, 7))
	kataSlice := []string{"indonesia", "tanah", "air", "beta"}
	fmt.Println(Contains(kataSlice, "indonesia"))
}
