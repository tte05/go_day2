package storage

import "architecture_go/services/contact/internal/domain/contact"

type Contact interface {
	CreateContact(c *contact.Contact) (*contact.Contact, error)
	DeleteContact(telephoneNumber string) error
	GetContact(telephone string) (*contact.Contact, error)
	UpdateContact(c *contact.Contact) (*contact.Contact, error)
}

type Group interface {
	CreatGroup(g *contact.Group) (*contact.Group, error)
	ReadGroup(ID int) (*contact.Group, error)
	AddContactToGroup(telephone string, id int) error
}
