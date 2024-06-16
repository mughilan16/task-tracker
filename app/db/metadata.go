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
  conn := db.getConnection()
  defer conn.Close()
	_, err := conn.Exec(q)
	if err != nil {
		log.Fatalln(err)
	}
  db.insertDefaultMetaData()
}

func (db DB) isMetaDataTableEmpty() (bool, error) {
	q := "SELECT id, isTaskActive, currentTaskId FROM metadata"
  conn := db.getConnection()
  defer conn.Close()
	rows, err := conn.Query(q)
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
  conn := db.getConnection()
  defer conn.Close()
  if err != nil {
    log.Fatalln(err)
  }
  if !isTableEmpty {
    return
  }
	q := "INSERT INTO metadata(id, isTaskActive, currentTaskId) VALUES(0, false, 0)"
	_, err = conn.Exec(q)
	if err != nil {
    log.Fatalln(err)
	}
}

func (db DB) resetMetaData() {
  conn := db.getConnection()
  defer conn.Close()
  q := "UPDATE metadata SET isTaskActive=false, currentTaskId=0 WHERE id = 0"
  _, err := conn.Exec(q)
  if err != nil {
    log.Fatalln(err)
  }
}

func (db DB) setMetaDataActive(id int) {
  conn := db.getConnection()
  defer conn.Close()
  _, err := conn.Exec("UPDATE metadata SET isTaskActive=true, currentTaskId = $1 WHERE id = 0", id)
  if err != nil {
    log.Fatalln(err)
  }
}

func (db DB) isTaskActive() (isTaskActive bool, id int) {
  conn := db.getConnection()
  defer conn.Close()
  rows, err := conn.Query("SELECT isTaskActive, currenttaskid from metadata where id=0")
  if err != nil {
    log.Println(err)
  }
  if rows.Next() {
    rows.Scan(&isTaskActive, &id)
  }
  return isTaskActive, id
}

