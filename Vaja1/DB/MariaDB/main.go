package MariaDB

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

type MariaDB struct {
	database *sql.DB
	User     string
	Pass     string
	IP       string
	Port     int
	Database string
}

func (db *MariaDB) Init() (err error) {

	// Open a connection to the database
	db.database, err = sql.Open("mysql", db.User+":"+db.Pass+"@tcp("+db.IP+":"+strconv.Itoa(db.Port)+")/"+db.Database+"?multiStatements=true&parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Check the connection. Retry every 5 seconds until it succeeds
	for {
		err = db.database.Ping()
		if err == nil {
			break
		}
		fmt.Println("Database not ready, retrying in 5 seconds...")
		<-time.After(5 * time.Second)
	}
	// Unused connections are closed automatically
	db.database.SetMaxIdleConns(0)

	// MariaDB default `MaxOpenConns == 100`. Set this value to the needs of the application
	db.database.SetMaxOpenConns(50)

	// Database migrations
	m := migrator{database: db.database, definitions: definitions}
	err = m.Migrate()
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}
