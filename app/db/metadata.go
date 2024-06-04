package db

import (
	"log"
)

type metaData struct {
	isTaskActive  bool
	currentTaskID int
}

func (db DB) createMetaDataTable() {
	q := "CREATE TABLE IF NOT EXISTS metadata(id INT, isTaskActive BOOLEAN, currentTaskId INT)"
	_, err := db.db.Exec(q)
	if err != nil {
		log.Fatalln(err)
	}
  db.insertDefaultMetaData()
}

func (db DB) isMetaDataTableEmpty() (bool, error) {
	q := "SELECT id, isTaskActive, currentTaskId FROM metadata"
	rows, err := db.db.Query(q)
	if err != nil {
		return false, err
	}
	i := 0
	for rows.Next() {
		i += 1
	}
	return i == 0, nil
}

func (db DB) insertDefaultMetaData() {
  isTableEmpty, err := db.isMetaDataTableEmpty()
  if err != nil {
    log.Fatalln(err)
  }
  if !isTableEmpty {
    return
  }
	q := "INSERT INTO metadata(id, isTaskActive, currentTaskId) VALUES(0, false, 0)"
	_, err = db.db.Exec(q)
	if err != nil {
    log.Fatalln(err)
	}
}
