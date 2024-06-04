package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

type DB struct {
	db     *sql.DB
	config Config
}

func Connect() *DB {
	db := &DB{
		db:     nil,
		config: getDevConfig(),
	}
  psqlInfo := db.getPsqlInfo()
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db.db = conn
	return db
}

func getDevConfig() Config {
	return Config{
		host:     "localhost",
		port:     5432,
		user:     "root",
		password: "secret",
		dbname:   "task-tracker",
	}
}

func (db *DB) Init() {
  db.createMetaDataTable()
}

func (db *DB) getPsqlInfo() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", db.config.host, db.config.port, db.config.user, db.config.password, db.config.dbname)
}

func (db *DB) TestDB() {
	err := db.db.Ping()
	if err != nil {
		panic(err)
	}
}
