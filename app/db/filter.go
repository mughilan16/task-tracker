package db

import (
	"fmt"
	"strings"
	"task-tracker/app/util"
	"time"
)

type taskWithFilter struct {
	s_no       int
	message    string
	month      time.Month
	year       int
	date       string
	start_time string
	stop_time  string
	total      int
}

func (db DB) getTotalForMonth(month int, year int) (total int) {
	tasks := db.filterMonthYear(month, year)
	total = 0
	for _, task := range tasks {
		total += task.total
	}
	return
}

func (db DB) filterMonthYear(month int, year int) (result []taskWithFilter) {
	tasks := db.getFilteredTask()
	for _, task := range tasks {
		if time.Month(month) == task.month && year == task.year {
			result = append(result, task)
		}
	}
	return
}

func (db DB) getFilteredTask() (tasks []taskWithFilter) {
	isTaskActive, _ := db.isTaskActive()
	if isTaskActive {
		fmt.Println("A task is currently active! complete to calculate total")
		return
	}
	tasks = convertTaskToFiltered(db.getTasks())
	return
}

func convertTaskToFiltered(tasks []Task) (result []taskWithFilter) {
	for i, task := range tasks {
		newTask := taskWithFilter{
			s_no:       i + 1,
			message:    task.message,
			date:       task.date[:10],
			start_time: task.start_time[11:19],
			stop_time:  task.stop_time[11:19],
			month:      time.Month(util.StringToInt(strings.Split(task.date[:10], "-")[1])),
			year:       util.StringToInt(task.date[:4]),
			total:      task.total,
		}
		result = append(result, newTask)
	}
	return
}
