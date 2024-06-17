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
	config Config
}

func New() *DB {
	db := &DB{
		config: getDevConfig(),
	}
	return db
}

func (db *DB) getConnection() *sql.DB {
	psqlInfo := db.getPsqlInfo()
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return conn
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
	db.createTaskTable()
	db.createProjectTable()
}

func (db *DB) getPsqlInfo() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", db.config.host, db.config.port, db.config.user, db.config.password, db.config.dbname)
}

func (db *DB) TestDB() {
  conn := db.getConnection()
	err := conn.Ping()
  defer conn.Close()
	if err != nil {
		panic(err)
	}
}
