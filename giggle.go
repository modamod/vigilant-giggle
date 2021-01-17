package main

import (
	"log"

	"github.com/modamod/giggle/org"
)

func main() {
	p1 := org.NewEmployee("modamod", 23945636, "27/12/1989")

	log.Println(p1.DateOfBirth)
	startDate, errStartDate := org.ToDate("01/02/2021")
	endDate, errEndDate := org.ToDate("05/02/2021")
	if errStartDate != nil || errEndDate != nil {
		log.Println(errStartDate)
		log.Println(errEndDate)
	}
	p1.RequestTimeOff(startDate, endDate)
}
