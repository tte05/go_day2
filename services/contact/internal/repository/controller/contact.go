package controller

import (
	use_case "architecture_go/services/contact/internal/use-case"
)

type contactController struct {
	contactUseCase use_case.Contact
}

func (cc contactController) GetContacts(ctx Context) error {
	//TODO implement me
	panic("implement me")
}

func (cc contactController) CreateContacts(ctx Context) error {
	//TODO implement me
	panic("implement me")
}

func (cc contactController) DeleteContacts(ctx Context) error {
	//TODO implement me
	panic("implement me")
}

func (cc contactController) UpdateContacts(ctx Context) error {
	//TODO implement me
	panic("implement me")
}

type Contact interface {
	GetContacts(c Context) error
	CreateContacts(c Context) error
	DeleteContacts(c Context) error
	UpdateContacts(c Context) error
}

func NewContactsController(cs use_case.Contact) Contact {
	return &contactController{cs}
}
