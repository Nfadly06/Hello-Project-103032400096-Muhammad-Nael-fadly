package main

import "fmt"

func editMedali(dataN DataNegara, dataM *DataMedali, n int) {
	var nama string
	var e, p, pr int
	fmt.Print("Nama negara: ")
	fmt.Scan(&nama)
	idx := cariNegara(dataN, n, nama)
	if idx != -1 {
		fmt.Print("Emas: ")
		fmt.Scan(&e)
		fmt.Print("Perak: ")
		fmt.Scan(&p)
		fmt.Print("Perunggu: ")
		fmt.Scan(&pr)
		dataM[idx] = Medali{e, p, pr}
		fmt.Println("Medali berhasil diubah.")
	} else {
		fmt.Println("Negara tidak ditemukan.")
	}
}

func hapusMedali(dataN DataNegara, dataM *DataMedali, n int) {
	var nama string
	fmt.Print("Nama negara: ")
	fmt.Scan(&nama)
	idx := cariNegara(dataN, n, nama)
	if idx != -1 {
		dataM[idx] = Medali{0, 0, 0}
		fmt.Println("Data medali berhasil direset.")
	} else {
		fmt.Println("Negara tidak ditemukan.")
	}
}
