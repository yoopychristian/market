//YOOPY CHRISTIAN

package services

import (
	"challenge/config"
	"challenge/utils"
	"fmt"
	"log"
)

type produk struct {
	id         int
	kodeProduk string
	namaProduk string
	harga      int
	created    string
	updated    string
}

func TambahProduct() {
	var code, namaProduk string
	var hargaProduk int
	fmt.Println("======================================")
	fmt.Println("TAMBAH PRODUK")
	fmt.Println("======================================")
	fmt.Print("Kode Produk: ")
	fmt.Scanln(&code)
	fmt.Print("Nama Produk: ")
	fmt.Scanln(&namaProduk)
	fmt.Print("Harga : ")
	fmt.Scanln(&hargaProduk)
	tambahProduct(code, namaProduk, hargaProduk)

}
func tambahProduct(code, nama string, harga int) {
	dbEngine, dbSource := config.Env_conn()
	db, err := utils.ConnDB(dbEngine, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("insert into product(product_code, product_name, price) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(code, nama, harga)
	if err != nil {
		log.Fatal(err)
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)
}

func HapusProduct() {
	var id int
	DetailProduct()
	fmt.Println("======================================")
	fmt.Println("HAPUS PRODUK")
	fmt.Println("======================================")
	fmt.Print("Masukan ID produk yang akan anda Hapus : ")
	fmt.Scanln(&id)
	hapusProduct(id)

}

func hapusProduct(id int) {
	dbEngine, dbSource := config.Env_conn()
	db, err := utils.ConnDB(dbEngine, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("delete from product where id=?;")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)

}

func DetailProduct() {
	daftar := detailProduct()
	fmt.Println("===========================================================================================================")
	fmt.Println("                                             SEMUA PRODUK                                                  ")
	fmt.Println("===========================================================================================================")
	fmt.Printf("\t%-5v\t%-10v\t%-10v\t%-8v\t%-20v\t%-10v\n", "ID.", "Kode Produk", "Nama Produk", "Harga", "Dibuat", "Disunting")
	for _, i := range daftar {
		fmt.Printf("\t%-5v\t%-10v\t%-10v\t%-8v\t%-20v\t%-10v\n", i.id, i.kodeProduk, i.namaProduk, i.harga, i.created, i.updated)

	}
}

func detailProduct() []produk {
	dbEngine, dbSource := config.Env_conn()
	db, err := utils.ConnDB(dbEngine, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	data, err := db.Query("select * from product;")
	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()
	var result []produk
	for data.Next() {
		var barang = produk{}
		var err = data.Scan(&barang.id, &barang.kodeProduk, &barang.namaProduk, &barang.harga, &barang.created, &barang.updated)

		if err != nil {
			log.Fatal(err)
		}

		result = append(result, barang)
	}
	if err = data.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
