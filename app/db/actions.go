package db

import (
	"fmt"
	"task-tracker/app/util"
	"time"
)

func (db DB) CreateNewProject(name, tag string) {
  db.addNewProject(name, tag)
  fmt.Printf("Created a new project: %s (%s)", name, tag)
}

func (db DB) AddNewTask(message string, tag string) {
	isTaskActive, _ := db.isTaskActive()
	if isTaskActive {
		fmt.Println("Another task is currently active! Complete it before starting new task")
		return
	}
	id := db.addNewTask(message, tag)
	db.setMetaDataActive(id)
	fmt.Println("Started new task :", message)
}

func (db DB) CompleteTask() {
	isTaskActive, id := db.isTaskActive()
	if !isTaskActive {
		fmt.Println("No task is current active")
		return
	}
	message, total_time := db.completeTask(id)
	total := util.MinuteToHour(total_time)
	fmt.Println("Complete task :", message, "taken", total)
	db.resetMetaData()
}

func (db DB) TotalForMonth(month int, year int, tag string) {
	total := db.getTotalForMonth(month, year, tag)
	fmt.Println(util.HourInfo(total))
}

func (db DB) Export(month int, year int, tag string) {
	tasks := db.filterMonthYear(month, year, tag)
	fmt.Printf("S. No.,Date,Start Time,End Time,Total,Job\n")
	for _, task := range tasks {
		fmt.Printf("%d,%s,%s,%s,%s,%s\n", task.s_no, task.date, task.start_time, task.stop_time, util.MinuteToHour(task.total), task.message)
	}
	total := db.getTotalForMonth(month, year, tag)
	fmt.Printf(",,,,%s\n", util.HourInfo(total))
}

func (db DB) ThisMonthTotal(tag string) {
	tasks := filterTag(convertTaskToFiltered(db.getTasks()), tag)
	total := 0
	currentMonth := time.Now().Month()
	for _, task := range tasks {
		if task.month == currentMonth {
			total += task.total
		}
	}
	currentTaskTime, _, currentTag, isTaskActive := db.getCurrentTask()
	if isTaskActive && tag == currentTag {
		total += currentTaskTime
	}
	fmt.Println(util.HourInfo(total))
}

func (db DB) TodayTotal(tag string) {
	newtasks := filterTag(convertTaskToFiltered(db.getTasks()), tag)
	tasks := filterDay(newtasks, time.Now().Day(), int(time.Now().Month()), time.Now().Year())
	total := 0
	for _, task := range tasks {
		total += task.total
	}
	currentTaskTime, _, currentTag, isTaskActive := db.getCurrentTask()
	if isTaskActive && currentTag == tag {
		total += currentTaskTime
	}
	fmt.Println(util.HourInfo(total))
}

func (db DB) ActiveTask(input_tag string) {
	total, message, tag, isActive := db.getCurrentTask()
	if !isActive {
		total, message, tag := db.getLastTask(input_tag)
    fmt.Printf("Worked on \"%s - %s\" %s ago", message, tag, util.HourInfo(total))
		return
	}
  fmt.Printf("Working on \"%s - %s\" for %s", message, tag, util.HourInfo(total))
}
