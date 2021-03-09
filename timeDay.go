package main

import (
	"fmt"
	"time"
)

func main() {

	switch time.Now().Weekday() {
	case time.Friday:
		fmt.Println("week-end!")

	case time.Tuesday:
		fmt.Println("today is tuesday")

	default:
		fmt.Println("rendom day!")
	}

	fmt.Println(time.Now().Clock())
}
