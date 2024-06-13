package db

import (
	"fmt"
	"task-tracker/app/util"
)

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
	fmt.Println("Total Time: ", util.MinuteToHour(total))
}

func (db DB) Export(month int, year int, tag string) {
	tasks := db.filterMonthYear(month, year, tag)
	fmt.Printf("S. No.,Date,Start Time,End Time,Total,Job\n")
	for _, task := range tasks {
		fmt.Printf("%d,%s,%s,%s,%s,%s\n", task.s_no, task.date, task.start_time, task.stop_time, util.MinuteToHour(task.total), task.message)
	}
	total := db.getTotalForMonth(month, year, tag)
	fmt.Printf(",,,,%s\n", util.MinuteToHour(total))
}
