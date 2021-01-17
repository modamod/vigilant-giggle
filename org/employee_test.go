package org

import (
	"log"
	"testing"

	"cloud.google.com/go/civil"

	"github.com/stretchr/testify/assert"
)

func TestIsEqual(t *testing.T) {
	p1 := NewEmployee("Employee1", 1234567, "01/01/1970")
	p2 := NewEmployee("Employee2", 1234567, "01/01/1970")
	if p1.IsEqual(*p2) {
		t.Error("Those two persons are not the same")
	} else {

		t.Log("Those two persons are not the same! Success...")
	}
}

func TestToDate(t *testing.T) {

	p1 := NewEmployee("Employee1", 1234567, "01/01/1970")
	date, _ := ToDate(p1.DateOfBirth.String())
	log.Printf("%T\n", date)
	assert.IsType(t, civil.Date{}, date)
}
