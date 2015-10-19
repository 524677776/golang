package main 

import(
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/mysql")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close();

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("The square number of 1 is: %d", 100)
}