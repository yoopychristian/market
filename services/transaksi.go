//YOOPY CHRISTIAN

package services

import (
	"challenge/config"
	"challenge/utils"
	"fmt"
	"log"
)

type transaksi struct {
	id         int
	total      int
	namaProduk string
	banyak     int
	harga      int
	created    string
	updated    string
}

func TambahTransaksi() {
	var inputCode, qty int
	DetailProduct()
	fmt.Println("======================================")
	fmt.Println("           TAMBAH TRANSAKSI           ")
	fmt.Println("======================================")
	fmt.Print("Masukan ID Produk: ")
	fmt.Scanln(&inputCode)
	fmt.Print("Banyak: ")
	fmt.Scanln(&qty)
	tambahTransaksi(inputCode, qty)

}

func tambahTransaksi(productId, qty int) {
	dbEngine, dbSource := config.Env_conn()
	db, err := utils.ConnDB(dbEngine, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("insert into purchase_order (order_date, created_at, updated_at) values (now(),now(),now());")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
	lastID, err := res.LastInsertId()

	stmt2, err := db.Prepare("insert into purchase_order_item (qty, created_at, updated_at, order_id, product_id) values (?,now(),now(),?,?);")
	if err != nil {
		log.Fatal(err)
	}
	res2, err := stmt2.Exec(qty, lastID, productId)
	if err != nil {
		log.Fatal(err)
	}
	lastID2, err := res2.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastID2, rowCnt)

}

func HapusTransaksi() {
	var id int
	DetailTransaksi()
	fmt.Println("======================================")
	fmt.Println("           HAPUS TRANSAKSI            ")
	fmt.Println("======================================")
	fmt.Print("Masukan Order ID transaksi yang akan di Hapus : ")
	fmt.Scanln(&id)
	hapusTransaksi(id)

}

func hapusTransaksi(id int) {
	dbEngine, dbSource := config.Env_conn()
	db, err := utils.ConnDB(dbEngine, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("delete from purchase_order_item where order_id=?;")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	lastID2, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastID2, rowCnt)

}

func DetailTransaksi() {
	daftarOrder := detailTransaksi()
	fmt.Println("===================================================================================================================================")
	fmt.Println("                                                      TRANSAKSI                                                                    ")
	fmt.Println("===================================================================================================================================")
	fmt.Printf("\t%-10v\t%-10v\t%-10v\t%-10v\t%-10v\t%-20v\t%-10v\n", "Order ID", "Nama Produk", "Banyak Produk", "Harga", "Total", "Dibuat", "Disunting")
	for _, i := range daftarOrder {
		fmt.Printf("\t%-10v\t%-10v\t%-10v\t%-10v\t%-10v\t%-20v\t%-10v\n", i.id, i.namaProduk, i.banyak, i.harga, i.total, i.created, i.updated)

	}

}

func detailTransaksi() []transaksi {
	dbEngine, dbSource := config.Env_conn()
	db, err := utils.ConnDB(dbEngine, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	data, err := db.Query("select po.id, poi.qty*p.price, p.product_name, poi.qty,p.price, po.created_at, po.updated_at from purchase_order po join purchase_order_item poi on po.id = poi.order_id join product p on p.id = poi.product_id;")
	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()
	var result []transaksi
	for data.Next() {
		var order = transaksi{}
		var err = data.Scan(&order.id, &order.total, &order.namaProduk, &order.banyak, &order.harga, &order.created, &order.updated)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, order)
	}
	if err = data.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
