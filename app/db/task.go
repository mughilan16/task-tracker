package db

import (
	"log"
	"strings"
	"time"
)

func (db DB) createTaskTable() {
	q := "CREATE TABLE IF NOT EXISTS tasks(id SERIAL, message TEXT, date DATE, start_time TIME, stop_time TIME, total INT)"
	_, err := db.db.Exec(q)
	if err != nil {
		log.Fatalln(err)
	}
}

func (db DB) addNewTask(message string) (id int) {
	q := "INSERT INTO tasks(message, date, start_time) VALUES($1, $2, $3) RETURNING id"
	dateTime := strings.Split(time.Now().String(), " ")
	date := dateTime[0]
	start_time := strings.Split(dateTime[1], ".")[0]
  rows, err := db.db.Query(q, message, date, start_time)
  if err != nil {
    log.Fatalln(err)
  }
  if rows.Next() {
    rows.Scan(&id)
  }
  return id
}

func (db DB) completeTask(taskId int) {
  q := "UPDATE tasks SET stop_time=$1 WHERE id=$2"
	dateTime := strings.Split(time.Now().String(), " ")
	stop_time := strings.Split(dateTime[1], ".")[0]
  _, err := db.db.Exec(q, stop_time, taskId)
  if err != nil {
    log.Fatalln(err)
  }
}
