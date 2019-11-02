package main

import "fmt"

type Manusia struct {
	tangan int
	kepala int
	kaki   int
}

type Proyektor interface {
	menampilkan()
	menyembunyikan()
}

func (Manusia) menampilkan() {
	fmt.Println("Manusia ngeprint depan UNIKOM")
}

func main() {
	var proyektor Proyektor
	proyektor = Manusia{2, 2, 1}
	proyektor.menampilkan()

	// var ivan Manusia
	// ivan.tangan = 2
	// ivan.kepala = 1
	// ivan.kaki = 10

	// fmt.Print("Si ivan punya ", ivan.kepala, " buah kepala")

	// var nilai map[string]int
	// nilai = make(map[string]int)

	// nilai["Dandi"] = 22
	// nilai["Juliadit"] = 53
	// nilai["Rakha"] = 45
	// nilai["James"] = 21

	// fmt.Println("Nilainya Dandi ", nilai["Dandi"])
}
