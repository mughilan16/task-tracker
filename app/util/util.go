package util

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func FormatDateAndTime(date string, delimiter string) (int, int, int) {
	dateList := strings.Split(date, delimiter)
	return StringToInt(dateList[0]), StringToInt(dateList[1]), StringToInt(dateList[2])
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln(err)
	}
	return i
}

func MinuteToHour(minu int) string {
	hours := minu / 60
	minutes := minu % 60
  fmt.Println(minutes)
	var hourString, minutesString string
	if hours < 10 {
		hourString = fmt.Sprintf("0%d", hours)
	} else {
		hourString = fmt.Sprintf("%d", hours)
	}
	if minutes < 10 {
		minutesString = fmt.Sprintf("0%d", minutes)
	} else {
		minutesString = fmt.Sprintf("%d", minutes)
	}
	return fmt.Sprintf("%s:%s", hourString, minutesString)
}
