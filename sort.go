package main

import "fmt"

func insertionSort(dataN *DataNegara, dataM *DataMedali, n int) {
	for pass := 1; pass < n; pass++ {
		tempN := dataN[pass]
		tempM := dataM[pass]
		i := pass

		for i > 0 && (tempM.Emas > dataM[i-1].Emas ||
			(tempM.Emas == dataM[i-1].Emas && tempM.Perak > dataM[i-1].Perak) ||
			(tempM.Emas == dataM[i-1].Emas && tempM.Perak == dataM[i-1].Perak && tempM.Perunggu > dataM[i-1].Perunggu)) {

			dataN[i] = dataN[i-1]
			dataM[i] = dataM[i-1]
			i--
		}

		dataN[i] = tempN
		dataM[i] = tempM
	}
}

func tampilkan(dataN DataNegara, dataM DataMedali, n int) {
	insertionSort(&dataN, &dataM, n)
	fmt.Println("\nPeringkat Negara:")
	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s | Emas: %d, Perak: %d, Perunggu: %d\n",
			i+1, dataN[i].Nama, dataM[i].Emas, dataM[i].Perak, dataM[i].Perunggu)
	}
}
