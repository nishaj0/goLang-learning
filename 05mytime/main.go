package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello time") // hello time

	presentTime := time.Now()
	fmt.Println(presentTime) // 2024-05-29 13:08:45.4188847 +0530 IST m=+0.003253001

	// formate
	fmt.Println(presentTime.Format("01-02-2006")) // 05-29-2024

	createdDate := time.Date(2020, time.June, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)                      // 2020-06-10 23:23:00 +0000 UTC
	fmt.Println(createdDate.Format("01-02-2006")) // 06-10-2020
}
