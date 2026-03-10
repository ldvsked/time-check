package main

import (
	"strings"
	"time"
	"unicode/utf8"
	"errors"
	"fmt"
)

func currentDayOfTheWeek() string {
	var currentDay int = int(time.Now().Weekday())
	var currentDayStringRussian string 
	switch {
	case currentDay == 0 :
		currentDayStringRussian = "Воскресенье"
	case currentDay == 1 :
		currentDayStringRussian = "Понедельник"
	case currentDay == 2 :
		currentDayStringRussian = "Вторник"
	case currentDay == 3 :
		currentDayStringRussian = "Среда"
	case currentDay == 4 :
		currentDayStringRussian = "Четверг"
	case currentDay == 5 :
		currentDayStringRussian = "Пятница"
	case currentDay == 6 :
		currentDayStringRussian = "Суббота"
	default :
	}
	return currentDayStringRussian
}

func dayOrNight() string {
	var currentHour int = time.Now().Hour()
	if currentHour >= 10 && currentHour <= 22 {
		return "День"
	}
	return "Ночь"
}

func nextFriday() int {
	var currentDay int = int(time.Now().Weekday())
	return (5 - currentDay + 7) % 7 
}

func CheckCurrentDayOfTheWeek(answer string) bool {
	var answerLower string = strings.ToLower(answer)
	return answerLower == strings.ToLower(currentDayOfTheWeek())
}

func CheckNowDayOrNight(answer string) (bool, error) {
	if utf8.RuneCountInString(answer) != 4 {
		return false, errors.New("исправь свой ответ, а лучше ложись спать")
	}
	answerLower := strings.ToLower(answer)
	return answerLower == strings.ToLower(dayOrNight()), nil
}

func main() {
	fmt.Println(currentDayOfTheWeek())
	fmt.Println(dayOrNight())
	fmt.Println(nextFriday())
}