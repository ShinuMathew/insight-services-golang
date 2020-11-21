package DBConnector

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Shinz9474/InsightAps/InsightAPI/Logger"
	_ "github.com/go-sql-driver/mysql"
)

type DB_Con struct {
	Protocol      string `json:"protocol"`
	Host          string `json:"host"`
	Port          string `json:"port"`
	Db_name       string `json:"db_name"`
	Db_username   string `json:"db_username"`
	Db_password   string `json:"db_password"`
	Db_connection string `json:"db_connection"`
	os            struct {
		Os_name     string `json:"os_name"`
		Db_username string `json:"db_username"`
		Db_password string `json:"db_password"`
	}
}

var db_con DB_Con

func ReadDBConfig(fileName string) {

	jsonFile, err := os.Open(fileName)

	defer jsonFile.Close()
	if err != nil {
		log.Printf("Following error occured when trying to read from the %s \n", fileName)
		panic(err.Error())
	}

	jsonFileParser := json.NewDecoder(jsonFile)
	err = jsonFileParser.Decode(&db_con)

	Logger.CheckError(err, "Following error occured when trying to parse the file: "+fileName)
}

func createConnectionString() string {

	Db_connection := db_con.Db_username + ":" + db_con.Db_password + "@" + db_con.Protocol + "(" + db_con.Host +
		":" + db_con.Port + ")/" + db_con.Db_name

	//tcp(127.0.0.1:3306)
	fmt.Printf("\n\n.. DBCON: %v", Db_connection)
	return Db_connection
}
func Open_DBConnection() *sql.DB {

	con_string := createConnectionString()
	log.Printf("DBConnection string: '%s' ", con_string)

	log.Printf("Attempting to connect '%s' DB...\n", db_con.Db_name)
	db, err := sql.Open("mysql", con_string)

	if err != nil {
		log.Printf("Following error occured when trying to connect to %s \n", db_con.Db_name)
		panic(err.Error())
	} else {
		log.Printf("DB Connection opened succesfully at ''%v''", time.Now())
	}
	return db
}

func Execute_query(db *sql.DB, query string) *sql.Rows {

	result, err := db.Query(query)
	if err != nil {
		log.Printf("\n\nFollowing error occured when executing the query:\n %s \nover '%s' DB\n", query, db_con.Db_name)
		db.Close()
		panic(err.Error())
	} else {
		log.Printf("Query completed execution succesfully...")
	}
	//defer db.Close()
	return result
}

func Close_DBConnection(db *sql.DB) {

	err := db.Close()
	if err != nil {
		log.Printf("\nFollowing error occured when closing DB connection\n")
		panic(err.Error())
	} else {
		log.Printf("\nDB Connection closed succesfully....")
	}
}
