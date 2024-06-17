package db

import (
	"database/sql"
	"log"
	"strings"
	"task-tracker/app/util"
	"time"
)

type Task struct {
	id         int
	message    string
	date       string
	start_time string
	stop_time  string
	total      int
	tag        string
}

func (db DB) createTaskTable() {
	q := "CREATE TABLE IF NOT EXISTS tasks(id SERIAL, message TEXT, date DATE, start_time TIME, stop_time TIME, total INT, tag TEXT)"
	conn := db.getConnection()
	defer conn.Close()
	_, err := conn.Exec(q)
	if err != nil {
		log.Fatalln(err)
	}
}

func (db DB) addNewTask(message string, tag string) (id int) {
	q := "INSERT INTO tasks(message, date, start_time, tag) VALUES($1, $2, $3, $4) RETURNING id"
	conn := db.getConnection()
	defer conn.Close()
	dateTime := strings.Split(time.Now().String(), " ")
	date := dateTime[0]
	start_time := strings.Split(dateTime[1], ".")[0]
	rows, err := conn.Query(q, message, date, start_time, tag)
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
	conn := db.getConnection()
	defer conn.Close()
	rows, err := conn.Query(q, taskId)
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
	conn := db.getConnection()
	defer conn.Close()
	rows, err := conn.Query(q, stop_time, total, taskId)
	if err != nil {
		log.Fatalln(err)
	}
	if rows.Next() {
		rows.Scan(&message)
	}
	return
}

func (db DB) getTasks() (tasks []Task) {
	q := "SELECT id, message, date, start_time, stop_time, total, tag FROM tasks WHERE stop_time IS NOT NULL"
	conn := db.getConnection()
	defer conn.Close()
	rows, err := conn.Query(q)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var task Task
		rows.Scan(&task.id, &task.message, &task.date, &task.start_time, &task.stop_time, &task.total, &task.tag)
		tasks = append(tasks, task)
	}
	return
}

func (db DB) getCurrentTask() (int, string, string, bool) {
	isTaskActive, taskId := db.isTaskActive()
	if !isTaskActive {
		return 0, "", "", false
	}
	q := "SELECT start_time, date, message, tag FROM tasks WHERE id=$1"
	conn := db.getConnection()
	defer conn.Close()
	rows, err := conn.Query(q, taskId)
	var startTime, date, message, tag string
	if err != nil {
		log.Fatalln(err)
	}
	if rows.Next() {
		rows.Scan(&startTime, &date, &message, &tag)
	}
	year, month, day := util.FormatDateAndTime(date[0:10], "-")
	hour, minu, sec := util.FormatDateAndTime(startTime[11:19], ":")
	start_time := time.Date(year, time.Month(month), day, hour, minu, sec, 0, time.Now().Location())
	total := int(time.Now().Sub(start_time).Minutes())
	return total, message, tag, true
}

func (db DB) getLastTask(input_tag string) (int, string, string) {
	conn := db.getConnection()
	defer conn.Close()
  var rows *sql.Rows
  var err error
	if input_tag == "all" {
		q := "SELECT stop_time, date, message, tag FROM tasks ORDER BY ID DESC LIMIT 1"
		rows, err = conn.Query(q)
	} else {
		q := "SELECT stop_time, date, message, tag FROM tasks WHERE tag=$1 ORDER BY ID DESC LIMIT 1"
		rows, err = conn.Query(q, input_tag)
	}
	var stopTime, date, message, tag string
	if err != nil {
		log.Fatalln(err)
	}
	if rows.Next() {
		rows.Scan(&stopTime, &date, &message, &tag)
	}
	year, month, day := util.FormatDateAndTime(date[0:10], "-")
	hour, minu, sec := util.FormatDateAndTime(stopTime[11:19], ":")
	start_time := time.Date(year, time.Month(month), day, hour, minu, sec, 0, time.Now().Location())
	total := int(time.Now().Sub(start_time).Minutes())
	return total, message, tag
}
