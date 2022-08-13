package main

import "fmt"

type Barang struct{
	NamaBarang string
	IdBarang, HargaBarang int
}

func (barang Barang) beli (nama string){
	fmt.Println(nama, "mau beli", barang.NamaBarang, "harganya", barang.HargaBarang)
}

func (barang Barang) cek() {
	fmt.Println(barang.NamaBarang, "masih ada")
}

func main() {
	var sabun Barang
	sabun.IdBarang=1
	sabun.NamaBarang="Sabun"
	sabun.HargaBarang=4000

	fmt.Println(sabun)
	fmt.Println(sabun.IdBarang)
	fmt.Println(sabun.NamaBarang)
	fmt.Println(sabun.HargaBarang)

	sikat := Barang{
		IdBarang: 2,
		NamaBarang: "Sikat Gigi",
		HargaBarang: 7000,
	}
	fmt.Println(sikat)

	odol := Barang{"Odol",3,5000}
	fmt.Println(odol)

	sabun.beli("Bayu")
	sikat.beli("Edo")
	odol.beli("Joko")

	sabun.cek()
	sikat.cek()
	odol.cek()
}