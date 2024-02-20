package use_case

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/use-case/adapters/repository"
)

type contactUseCase struct {
	contactRepo repository.ContactRepository
	db          repository.DBRepository
}

func (c2 *contactUseCase) CreateCon(c *contact.Contact) (*contact.Contact, error) {
	c, err := c2.contactRepo.Create(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c2 *contactUseCase) DeleteCon(telephoneNumber string) error {
	err := c2.contactRepo.Delete(telephoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func (c2 *contactUseCase) GetCon(telephone string) (*contact.Contact, error) {
	c, err := c2.contactRepo.Get(telephone)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c2 *contactUseCase) UpdateCon(c *contact.Contact) (*contact.Contact, error) {
	c, err := c2.contactRepo.Update(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

type Contact interface {
	CreateCon(c *contact.Contact) (*contact.Contact, error)
	DeleteCon(telephoneNumber string) error
	GetCon(telephone string) (*contact.Contact, error)
	UpdateCon(c *contact.Contact) (*contact.Contact, error)
}

func NewContactUseCase(c repository.ContactRepository, db repository.DBRepository) Contact {
	return &contactUseCase{c, db}
}
