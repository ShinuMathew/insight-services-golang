package DBConnect

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Execute_Query(db_name, query string) *sql.Rows {

	log.Printf("Attempting to connect to %s DB at %v", db_name, time.Now())

	con_string := "root:shinz9474@tcp(127.0.0.1:3306)/" + db_name

	db, err := sql.Open("mysql", con_string)

	if err != nil {
		panic(err.Error())
	} else {
		log.Printf("Successfully connected to %s DB at %v", db_name, time.Now())
	}

	defer db.Close()

	log.Printf("Started executing the below query:\n%s\n at %v", query, time.Now())

	result, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	} else {
		log.Printf("Query completed execution succesfully at %v", time.Now())
	}

	return result
}
