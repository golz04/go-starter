package main

import (
	"fmt"
	"time"
)

func sampleVariable() {
	var name, jenis_kelamin string

	name = "Muhammad Haidar" //menginisialisasikan diawal
	jenis_kelamin = "Laki-laki"
	place_of_birth := "Bondowoso" //tanpa menginisialisasikan di awal dan tanpa menggunakan key "var"
	date_of_birth := time.Date(2000, 12, 30, 0, 0, 0, 0, time.UTC)
	umur := 23

	fmt.Println("Nama           = ", name)
	fmt.Println("Jenis Kelamin  = ", jenis_kelamin)
	fmt.Println("Tempat Lahir   = ", place_of_birth)
	fmt.Println("Tanggal Lahir  = ", date_of_birth.Format("2006-01-02"))
	fmt.Println("Umur           = ", umur)
}
