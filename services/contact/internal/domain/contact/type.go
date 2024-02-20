package contact

import (
	"fmt"
	"regexp"
)

const MaxGroupLettersLength = 250

type Contact struct {
	telephoneNumber string
	Name            name
}

type name struct {
	firstName string
	lastName  string
}

func (c *Contact) NumbersPhone(phoneNumber string) {
	reg := regexp.MustCompile("[^0-9]+")
	c.telephoneNumber = reg.ReplaceAllString(phoneNumber, "")
}

func (n name) FirstName() string {
	return n.firstName
}

func (n name) LastName() string {
	return n.lastName
}

type Group struct {
	ID   int
	Name string
}

func (g *Group) SetName(name string) error {
	if len(name) > MaxGroupLettersLength {
		return fmt.Errorf("max %d charachter in group name", MaxGroupLettersLength)
	}
	g.Name = name
	return nil
}
