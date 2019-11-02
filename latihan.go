package main

import (
	"bufio"
	"fmt"
	"os"
)

type Buku struct {
	judul string
	tahun string
	isbn  string
}

func main() {
	scanner := bufio.NewReader(os.Stdin)

	var buku Buku
	fmt.Print("Masukkan judul buku: ")
	buku.judul, _ = scanner.ReadString('\n')

	fmt.Print("Masukkan terbit buku: ")
	buku.tahun, _ = scanner.ReadString('\n')

	fmt.Print("Masukkan isbn buku: ")
	buku.isbn, _ = scanner.ReadString('\n')

	fmt.Println("Data buku yang dimasukkan:")
	fmt.Println("Judul: ", buku.judul)
	fmt.Println("Tahun terbit buku: ", buku.tahun)
	fmt.Println("ISBN buku: ", buku.isbn)
}
