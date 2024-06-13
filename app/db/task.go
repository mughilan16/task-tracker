package db

import (
	"log"
	"strings"
	"task-tracker/app/util"
	"time"
)

type Task struct {
  id int
  message string
  date string
  start_time string
  stop_time string
  total int
}

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
	return
}

func (db DB) getStartTime(taskId int) time.Time {
	q := "SELECT start_time, date FROM tasks WHERE id=$1"
	rows, err := db.db.Query(q, taskId)
	var startTime, date string
	if err != nil {
		log.Fatalln(err)
	}
	if rows.Next() {
		rows.Scan(&startTime, &date)
	}
	year, month, day := util.FormatDateAndTime(date[0:10], "-")
	hour, minu, sec := util.FormatDateAndTime(startTime[11:19], ":")
	return time.Date(year, time.Month(month), day, hour, minu, sec, 0, time.Now().Location())
}

func (db DB) completeTask(taskId int) (message string, total int) {
	q := "UPDATE tasks SET stop_time=$1, total=$2 WHERE id=$3 RETURNING message"
	dateTime := strings.Split(time.Now().String(), " ")
	stop_time := strings.Split(dateTime[1], ".")[0]
	start_time := db.getStartTime(taskId)
	total = int(time.Now().Sub(start_time).Minutes())
	rows, err := db.db.Query(q, stop_time, total, taskId)
	if err != nil {
		log.Fatalln(err)
	}
	if rows.Next() {
		rows.Scan(&message)
	}
	return
}

func (db DB) getTasks() (tasks []Task) {
  q := "SELECT id, message, date, start_time, stop_time, total FROM tasks"
  rows, err := db.db.Query(q)
  if err != nil {
    log.Fatalln(err)
  }
  for rows.Next() {
    var task Task
    rows.Scan(&task.id, &task.message, &task.date, &task.start_time, &task.stop_time, &task.total)
    tasks = append(tasks, task)
  }
  return
}

