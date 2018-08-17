package main

import (
	"fmt"
	"time"
)

type TeamMember interface {
	PrintName()
	PrintDetails()
}
type Employee struct {
	FirstName, LastName string
	Dob                 time.Time
	JobTitle, Location  string
}

func (e Employee) PrintName() {
	fmt.Printf("\n%s %s\n", e.FirstName, e.LastName)
}

func (e Employee) PrintDetails() {
	fmt.Printf("Date of Birth: %s, Job: %s, Location: %s\n", e.Dob.String(), e.JobTitle, e.Location)
}

type Developer struct {
	Employee //type embedding for composition
	Skills   []string
}

func main() {
	d := Developer{
		Employee{
			"Krysta",
			"Boyle",
			time.Date(1986, time.April, 28, 0, 0, 0, 0, time.UTC),
			"Software Engineer",
			"San Francisco",
		},
		[]string{"Go", ".Net Core", "Java"},
	}
	d.PrintName()
	d.PrintDetails()
}
