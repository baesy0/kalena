package main

import (
	"fmt"
)

func main() {
	var lastday = [...]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	year := 2004
	month := 3
	result := 0

	//fmt.Print("년:")
	//fmt.Scanf("%d", &year)
	//fmt.Print("월:")
	//fmt.Scanf("%d", &month)

	if month == 2 {
		result = lastDate(year, month) + 28
	} else {
		result = lastday[month]
	}

	fmt.Println(result)
}

func lastDate(year int, month int) int {

	flag := 0

	if year%4 == 0 && year%100 != 0 {
		flag = 1
	} else if year%400 == 0 {
		flag = 1
	} else {
		flag = 0
	}

	return flag
}
