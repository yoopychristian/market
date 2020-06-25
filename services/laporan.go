//YOOPY CHRISTIAN

package services

import (
	"challenge/config"
	"challenge/utils"
	"fmt"
	"log"
)

type laporan struct {
	grup       int
	namaProduk string
	harga      int
	banyak     int
	total      int
}

func detailHarian() []laporan {
	dbEngine, dbSource := config.Env_conn()
	db, err := utils.ConnDB(dbEngine, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	data, err := db.Query("select day(po.created_at) ,p.product_name, p.price, sum(poi.qty), sum(poi.qty * p.price) from purchase_order po join purchase_order_item poi on po.id = poi.order_id join product p on p.id = poi.product_id group by p.id; ")

	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()
	var result []laporan
	for data.Next() {
		var daily = laporan{}
		var err = data.Scan(&daily.grup, &daily.namaProduk, &daily.harga, &daily.banyak, &daily.total)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, daily)
	}
	if err = data.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func LaporanHarian() {
	laporanHarian := detailHarian()
	fmt.Println("================================================================================")
	fmt.Println("                                 Laporan Harian                                 ")
	fmt.Println("================================================================================")
	fmt.Printf("\t%-10v\t%-10v\t%-10v\t%-10v\t%-10v\n", "Tanggal", "Nama Produk", "Banyak", "Harga", "Total")
	for _, i := range laporanHarian {
		fmt.Printf("\t%-10v\t%-10v\t%-10v\t%-10v\t%-10v\n", i.grup, i.namaProduk, i.banyak, i.harga, i.total)
	}
}

func detailBulanan() []laporan {
	dbEngine, dbSource := config.Env_conn()
	db, err := utils.ConnDB(dbEngine, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	data, err := db.Query("select month(po.created_at) ,p.product_name, p.price, sum(poi.qty), SUM(poi.qty * p.price) from purchase_order po join purchase_order_item poi on po.id = poi.order_id join product p on p.id = poi.product_id group by p.id; ")
	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()
	var result []laporan
	for data.Next() {
		var monthly = laporan{}
		var err = data.Scan(&monthly.grup, &monthly.namaProduk, &monthly.harga, &monthly.banyak, &monthly.total)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, monthly)
	}
	if err = data.Err(); err != nil {
		log.Fatal(err)
	}
	return result

}

func LaporanBulanan() {
	laporanBulanan := detailBulanan()
	fmt.Println("================================================================================")
	fmt.Println("                              Laporan Bulanan                                   ")
	fmt.Println("================================================================================")
	fmt.Printf("\t%-10v\t%-10v\t%-10v\t%-10v\t%-10v\n", "Bulan", "Nama Produk", "Banyak", "Harga", "Total")
	for _, i := range laporanBulanan {
		fmt.Printf("\t%-10v\t%-10v\t%-10v\t%-10v\t%-10v\n", i.grup, i.namaProduk, i.banyak, i.harga, i.total)
	}

}
func detailLaporan() []laporan {
	dbEngine, dbSource := config.Env_conn()
	db, err := utils.ConnDB(dbEngine, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	data, err := db.Query("select po.id ,p.product_name, p.price, sum(poi.qty), sum(poi.qty * p.price) from purchase_order po join purchase_order_item poi on po.id = poi.order_id join product p on p.id = poi.product_id group by po.id")
	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()
	var result []laporan
	for data.Next() {
		var all = laporan{}
		var err = data.Scan(&all.grup, &all.namaProduk, &all.harga, &all.banyak, &all.total)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, all)
	}
	if err = data.Err(); err != nil {
		log.Fatal(err)
	}
	return result

}

func SemuaLaporan() {
	allReport := detailLaporan()
	fmt.Println("================================================================================")
	fmt.Println("                                  Semua Laporan                                 ")
	fmt.Println("================================================================================")
	fmt.Printf("\t%-10v\t%-10v\t%-10v\t%-10v\t%-10v\n", "Order ID", "Nama Produk", "Banyak", "Harga", "Total")
	for _, i := range allReport {
		fmt.Printf("\t%-10v\t%-10v\t%-10v\t%-10v\t%-10v\n", i.grup, i.namaProduk, i.banyak, i.harga, i.total)
	}

}
