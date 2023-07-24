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

	// Odpremo povezavo do baze
	db.database, err = sql.Open("mysql", db.User+":"+db.Pass+"@tcp("+db.IP+":"+strconv.Itoa(db.Port)+")/"+db.Database+"?multiStatements=true&parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Preverimo povezavo z bazo. Če ni vzpostavljena probamo vsakih 5 sekund
	for {
		err = db.database.Ping()
		if err == nil {
			break // Ce ni napak pomeni, da je povezava vzpostavljena
		}
		fmt.Println("Database not ready")
		<-time.After(5 * time.Second)
	}

	// Neuporabljene povezave se avtomatsko zaprejo
	db.database.SetMaxIdleConns(0)

	// Max število povezav se nastavi glede na bazo in potrebe. MariaDB ima default 100 povezav.
	db.database.SetMaxOpenConns(50)

	// Migracije baze
	m := migrator{database: db.database, definitions: definitions}
	err = m.Migrate()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return

}
