package db

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq" // postgres driver
	log "go.uber.org/zap"
)

// PgDB defines pgdb type
type PgDB struct {
	DB *sql.DB
}

// CreatePGConnection return db connection instance
func CreatePGConnection(opts map[string]string) (*PgDB, error) {
	port, err := strconv.Atoi(opts["port"])
	if err != nil {
		log.S().Fatal("Invalid port number : ", opts["port"])
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		opts["host"], port, opts["user"], opts["password"], opts["dbname"], opts["sslmode"])

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.S().DPanic(err)
	}

	err = db.Ping()
	if err != nil {
		log.S().DPanic(err)
	}

	log.S().Info("Connected to PG DB Server: ", opts["host"], " at port:", opts["port"], " successfully!")

	return &PgDB{DB: db}, nil
}
