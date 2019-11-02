package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pengguna struct {
	nim      string
	password string
	nama     string
}

var scanner = bufio.NewReader(os.Stdin)
var listPengguna = []Pengguna{}
var menu int

func main() {
	for menu != 4 {
		inputMenu()
		switch menu {
		case 1:
			tambahData()
		case 2:
			cariData()
		case 3:
			login()
		case 4:
			fmt.Println("Selamat jalan ~")
		default:
			fmt.Println("Menu yang dimasukkan tidak tersedia ~")
		}
	}
}

func inputMenu() {
	fmt.Println("1. Tambah data")
	fmt.Println("2. Cari data")
	fmt.Println("3. Login")
	fmt.Println("4. Keluar")
	fmt.Print("Masukkan menu yang dipilih: ")
	fmt.Scanln(&menu)
}

func tambahData() {
	fmt.Print("Masukkan Nim: ")
	nim, _ := scanner.ReadString('\n')
	fmt.Print("Masukkan Password: ")
	password, _ := scanner.ReadString('\n')
	fmt.Print("Masukkan Nama: ")
	nama, _ := scanner.ReadString('\n')

	penggunaBaru := Pengguna{nim, password, nama}
	listPengguna = append(listPengguna, penggunaBaru)
}

func cariData() {
	fmt.Print("Masukkan Nim: ")
	nim, _ := scanner.ReadString('\n')

	ketemu := false
	penggunaYangDicari := Pengguna{}
	for _, pengguna := range listPengguna {
		if pengguna.nim == nim {
			penggunaYangDicari = pengguna
			ketemu = true
			break
		}
	}

	if ketemu {
		fmt.Print("Nim: ", penggunaYangDicari.nim)
		fmt.Print("Nama: ", penggunaYangDicari.nama)
	} else {
		fmt.Print("Tidak ketemu")
	}
}

func login() {
	fmt.Print("Masukkan Nim: ")
	nim, _ := scanner.ReadString('\n')

	ketemu := false
	for _, pengguna := range listPengguna {
		if pengguna.nim == nim {
			ketemu = true
			fmt.Print("Masukkan Password: ")
			password, _ := scanner.ReadString('\n')
			if pengguna.password == password {
				fmt.Println("Sukses login")
			} else {
				fmt.Println("Gagal login")
			}

			return
		}
	}

	if !ketemu {
		fmt.Println("Pengguna tidak ditemukan")
	}
}
