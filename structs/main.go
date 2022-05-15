package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
}

type personWithContactInfo struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p personWithContactInfo) print() {

	fmt.Printf("\n%+v", p)
}

func (p personWithContactInfo) updateFirstName(newFirstName string) {

	p.firstName = newFirstName
}

func main() {

	alex := person{firstName: "Alex", lastName: "Anderson"}

	fmt.Println(alex)

	var nick person

	nick.firstName = "Nick"
	nick.lastName = "White"

	fmt.Printf("%+v", nick)

	jim := personWithContactInfo{
		firstName: "Jim",
		lastName:  "Sky",
		contact: contactInfo{
			email:   "jim@sky.com",
			zipCode: 43215,
		},
	}

	//fmt.Printf("\n%+v", jim)
	//jim.updateFirstName("Jimmy")
	jim.print()
}
