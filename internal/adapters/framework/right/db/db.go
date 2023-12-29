package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	db, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		log.Fatalf("Db Connection Failure: %v", err)
	}

	// test connection
	err = db.Ping()

	if err != nil {
		log.Fatalf("Db Ping Failure: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (dbA Adapter) CloseDBConnection() {
	err := dbA.db.Close()

	if err != nil {
		log.Fatalf("Db Close Failure: %v", err)
	}
}

func (dbA Adapter) Create(answer int32, operation string) error {
	queryString, args, err := sq.Insert("arithmetics").Columns("created_at", "answer", "operation").Values(time.Now().Unix(), answer, operation).ToSql()

	if err != nil {
		return err
	}

	_, err = dbA.db.Exec(queryString, args...)

	if err != nil {
		return err
	}

	return nil
}
