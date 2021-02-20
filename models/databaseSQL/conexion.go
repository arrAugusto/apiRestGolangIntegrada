package databaseSQL

import (
	"database/sql" //importando el manejador de bases de datos
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb" //driver para conectarse ala base de datos SQLSERVER 2014
)

var (
	server   = "localhost"
	port     = 1433
	user     = "logUser"
	password = "Contra$2019#"
	database = "Integrada"
)

// Puntero a la estructura DB, nos permite manejar la
// base de datos
var db *sql.DB

func ConectionSQL() *sql.DB {
	// Connect to database
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	return conn

}
