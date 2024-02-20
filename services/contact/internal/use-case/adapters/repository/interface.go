package repository

import "architecture_go/services/contact/internal/domain/contact"

type ContactRepository interface {
	Create(c *contact.Contact) (*contact.Contact, error)
	Delete(telephoneNumber string) error
	Get(telephone string) (*contact.Contact, error)
	Update(c *contact.Contact) (*contact.Contact, error)
}

type GroupRepository interface {
	Create(g *contact.Group) (*contact.Group, error)
	Read(ID int) (*contact.Group, error)
	AddContactToGroup(telephone string, id int) error
}
