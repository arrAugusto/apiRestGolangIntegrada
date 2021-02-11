package conexionSQL

/**
*El uso de la libreria github.com/denisenkom/go-mssqldb es para uso de manejo de la conexion a
*base de datos con en SQLSERVER
**/

import (
	"database/sql"
	"fmt"
	"log"
)

//Declarando variables de conexion a sqlserver
var (
	server   = "DESKTOP-2MO9KA5/SQLEXPRESS"
	port     = 1433
	user     = "logUser"
	password = "Contra$2019#"
	database = "Integrada"
)

var db *sql.DB

func conexionDB() {
	var err error

	// Create connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d",
		server, user, password, port)

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
	log.Printf("Connected!\n")

	// Close the database connection pool after program executes
	defer db.Close()

}
