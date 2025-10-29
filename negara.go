package main

import "fmt"

const MAX = 100

type Negara struct {
	Nama string
}

type Medali struct {
	Emas, Perak, Perunggu int
}

type DataNegara [MAX]Negara
type DataMedali [MAX]Medali

func cariNegara(data DataNegara, n int, nama string) int {
	for i := 0; i < n; i++ {
		if data[i].Nama == nama {
			return i
		}
	}
	return -1
}

func tambahNegara(dataN *DataNegara, dataM *DataMedali, n *int) {
	var nama string
	fmt.Print("Nama Negara: ")
	fmt.Scan(&nama)
	if cariNegara(*dataN, *n, nama) == -1 {
		dataN[*n].Nama = nama
		dataM[*n] = Medali{0, 0, 0}
		*n++
		fmt.Println("Negara berhasil ditambahkan.")
	} else {
		fmt.Println("Negara sudah terdaftar.")
	}
}

func editNegara(dataN *DataNegara, n int) {
	var lama, baru string
	fmt.Print("Nama negara yang ingin diedit: ")
	fmt.Scan(&lama)
	idx := cariNegara(*dataN, n, lama)
	if idx != -1 {
		fmt.Print("Nama baru: ")
		fmt.Scan(&baru)
		dataN[idx].Nama = baru
		fmt.Println("Nama negara berhasil diubah.")
	} else {
		fmt.Println("Negara tidak ditemukan.")
	}
}

func hapusNegara(dataN *DataNegara, dataM *DataMedali, n *int) {
	var nama string
	fmt.Print("Nama negara yang ingin dihapus: ")
	fmt.Scan(&nama)
	idx := cariNegara(*dataN, *n, nama)
	if idx != -1 {
		for i := idx; i < *n-1; i++ {
			dataN[i] = dataN[i+1]
			dataM[i] = dataM[i+1]
		}
		*n--
		fmt.Println("Negara berhasil dihapus.")
	} else {
		fmt.Println("Negara tidak ditemukan.")
	}
}
