package db


import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/go-sql-driver/mysql"
)

// You'll use a pointer to an sql.DB struct,
// which represents access to a specific database.
// Declare a db variable of type *sql.DB. This is your database handle.
var DB *sql.DB

func Connect() {
	// Capture connection properties.
	// Use the MySQL driver’s Config – and the type’s FormatDSN -– 
	// to collect connection properties and format them into a DSN 
	// for a connection string.
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "recordings"

	// Get a database handle.
	var err error
	// Call sql.Open to initialize the DB variable, passing the return value of FormatDSN.
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	// Check for an error from sql.Open. It could fail if, for example, 
	// your database connection specifics weren’t well-formed.
	if err != nil {
		// To simplify the code, you’re calling log.Fatal to end execution and print the error to the console. 
		// In production code, you’ll want to handle errors in a more graceful way.
		log.Fatal(err)
	}
	// Call DB.Ping to confirm that connecting to the database works. 
	// At run time, sql.Open might not immediately connect, depending on the driver. 
	// You’re using Ping here to confirm that the database/sql package can connect when it needs to.
	pingErr := DB.Ping()
	// Check for an error from Ping, in case the connection failed.
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	// Print a message if Ping connects successfully.
	fmt.Println("Connected!")

}