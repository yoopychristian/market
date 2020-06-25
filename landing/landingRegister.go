//YOOPY CHRISTIAN

package landing

import (
	"challenge/config"
	"challenge/menu"
	"challenge/utils"
	"fmt"
	"log"
	"time"
)

type Register struct {
	username string
	password string
}

func TambahPengguna() {
	var username, password string
	dbEngine, dbSource := config.Env_conn()
	db, err := utils.ConnDB(dbEngine, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Print("Masukkan Username : ")
	fmt.Scan(&username)
	fmt.Print("Masukkan Password : ")
	fmt.Scan(&password)
	stmt, err := db.Prepare("insert into login(username, password) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(username, password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Username %v telah ditambahkan didatabase\n", username)
	time.Sleep(2 * time.Second)
	menu.ClearScreenLoginRegister()
}

func HapusPengguna() {
	var username string
	dbEngine, dbSource := config.Env_conn()
	db, err := utils.ConnDB(dbEngine, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Masukan username yang akan dihapus: ")
	fmt.Scan(&username)

	stmt, err := db.Prepare("delete from login where username = ?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(username)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Username %v telah dihapus\n", username)
	time.Sleep(2 * time.Second)
	menu.ClearScreenLoginRegister()
}

func DetailPengguna() {
	dbEngine, dbSource := config.Env_conn()
	db, err := utils.ConnDB(dbEngine, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("Select username, password from login")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var listUser []*Register
	for rows.Next() {
		var user = Register{}
		var err = rows.Scan(&user.username, &user.password)
		if err != nil {
			log.Fatal(err)
		}
		listUser = append(listUser, &user)
	}
	fmt.Println("===========================================")
	fmt.Println("               User Details                ")
	fmt.Println("===========================================")
	fmt.Printf("%v\t%-10v\t%-10v\n", "no", "username", "password")

	for id, p := range listUser {
		fmt.Printf("%v\t%-10v\t%-10v\n", id+1, p.username, p.password)
	}

}
