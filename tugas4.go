package main

import "fmt"

func main() {
	var mahasiswa = map[string]int{
		"Aldo":  182,
		"Yosep": 178,
	}

	tampilMahasiswa(mahasiswa)

}

func tampilMahasiswa(mahasiswa map[string]int) {
	for k, v := range mahasiswa {
		fmt.Println(k, " : ", v, " cm")
	}
}
