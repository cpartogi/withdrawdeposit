package db

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	log "go.uber.org/zap"
)

// PgDB defines pgdb type
type MySqlDB struct {
	DB *sql.DB
}

// CreatePGConnection return db connection instance
func CreateMySqlConnection(opts map[string]string) (*MySqlDB, error) {
	port, err := strconv.Atoi(opts["port"])
	if err != nil {
		log.S().Fatal("Invalid port number : ", opts["port"])
	}

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", opts["user"], opts["password"], opts["host"], port, opts["dbname"])

	db, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		log.S().DPanic(err)
	}

	err = db.Ping()
	if err != nil {
		log.S().DPanic(err)
	}

	log.S().Info("Connected to MySql Server: ", opts["host"], " at port:", opts["port"], " successfully!")

	return &MySqlDB{DB: db}, nil
}
