package org

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/civil"
)

//Employee is a class used to describe a person
type Employee struct {
	Name        string
	Number      int64
	DateOfBirth civil.Date
}

//NewEmployee constructor for person
func NewEmployee(name string, number int64, dateOfBirth string) *Employee {
	if dateOfBirth == "" {
		dateOfBirth = "01/01/1970"
	}
	dateOfBirthObject, err := ToDate(dateOfBirth)
	if err != nil {
		log.Fatal(err)
	}

	return &Employee{Name: name, Number: number, DateOfBirth: dateOfBirthObject}
}

//String string format
func (p Employee) String() string {

	return fmt.Sprintf("Name: [%s]\nNumber: [%d]\nDate Of Birth: [%d/%d/%d]\n", p.Name, p.Number, p.DateOfBirth.Day, p.DateOfBirth.Month, p.DateOfBirth.Year)
}

//Call is a method to call a person
func (p Employee) Call() {
	log.Printf("Calling %s on their phone number %d\n", p.Name, p.Number)
}

//SendSMS is a method to send SMS to a person
func (p Employee) SendSMS(msg string) {

	log.Printf("Sending [%s] to [%d]\n", msg, p.Number)
}

//ToDate is a method to convert date string to Date Object
func ToDate(strdate string) (civil.Date, error) {
	var dateSlice []string
	if strings.Contains(strdate, "/") {

		dateSlice = strings.Split(strdate, "/")
	} else if strings.Contains(strdate, "-") {
		dateSlice = strings.Split(strdate, "-")
	} else {
		return civil.Date{}, errors.New("unhandled date format, please make sure your date is using the following formats\n dd/mm/yyyy or dd-mm-yyyy")
	}
	day, _ := strconv.Atoi(dateSlice[0])
	month, _ := strconv.Atoi(dateSlice[1])
	year, _ := strconv.Atoi(dateSlice[2])
	return civil.Date{Year: year, Month: time.Month(month), Day: day}, nil
}

//IsEqual method to determin if two objects are the same
func (p Employee) IsEqual(v Employee) bool {

	return p.Name == v.Name
}

//RequestTimeOff is a method for person to request time off.
func (p Employee) RequestTimeOff(startDate civil.Date, endDate civil.Date) (ok bool) {
	ok = true
	log.Printf("[%v] is requesting time off from [%v] to [%v]...\n", p.Name, startDate, endDate)

	return ok
}
