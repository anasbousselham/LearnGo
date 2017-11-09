package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	
	db, err := sql.Open("mysql", "root:anasoft@/ivgcafe?charset=utf8")
	if err != nil {
		checkErr(err)  // Just for example purpose. You should use proper error handling instead of panic
	} else {
     fmt.Println("vous etes connecté")
	}
	defer db.Close()

	row, err := db.Query("select username,ip from logins") 
	checkErr(err)
	for row.Next(){
		var username string
		var ip string
		err := row.Scan(&username,&ip)
		checkErr(err)
		fmt.Printf(username)
		fmt.Printf("|")
		fmt.Println(ip)
	}
	
	
	
	}

	func checkErr(err error) {
        if err != nil {
            panic(err)
        }
    }


