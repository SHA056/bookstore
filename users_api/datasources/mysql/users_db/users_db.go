package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

// const (
// 	mysql_username = "mysql_users_username"
// 	mysql_password = "mysql_users_password"
// 	mysql_host     = "mysql_users_host"
// 	mysql_schema   = "mysql_users_schema"
// )

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load("../../../db.env")
	if err != nil {
		panic(err)
	}

	return os.Getenv(key)
}

var (
	Client *sql.DB

	username = goDotEnvVariable("mysql_username")
	password = goDotEnvVariable("mysql_password")
	host     = goDotEnvVariable("mysql_host")
	schema   = goDotEnvVariable("mysql_schema")

	// username = os.Getenv(mysql_username)
	// password = os.Getenv(mysql_password)
	// host     = os.Getenv(mysql_host)
	// schema   = os.Getenv(mysql_schema)
)

func init() {
	// username:password@tcp(host)/dbName
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		schema,
	)
	var err error
	Client, err := sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	// mysql.SetLogger()
	log.Println("dataabse successfully configured")
}
