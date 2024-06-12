package db

import (
	"fmt"
	"task-tracker/app/util"
)

func (db DB) AddNewTask(message string) {
	isTaskActive, _ := db.isTaskActive()
	if isTaskActive {
		fmt.Println("Another task is currently active! Complete it before starting new task")
		return
	}
	id := db.addNewTask(message)
	db.setMetaDataActive(id)
  fmt.Println("Create new task :", message)
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
