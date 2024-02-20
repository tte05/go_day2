package postgres

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/use-case/adapters/repository"
	"github.com/jinzhu/gorm"
)

type contactRepository struct {
	db *gorm.DB
}

func (cc contactRepository) Create(c *contact.Contact) (*contact.Contact, error) {
	err := cc.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, err
}

func (cc contactRepository) Delete(telephoneNumber string) error {
	err := cc.db.Delete(&telephoneNumber).Error
	if err != nil {
		return err
	}
	return nil
}

func (cc contactRepository) Get(telephone string) (*contact.Contact, error) {
	err := cc.db.Find(&telephone).Error
	if err != nil {
		return nil, err
	}
	// nuzno uzmenit
	return nil, nil
}

func (cc contactRepository) Update(c *contact.Contact) (*contact.Contact, error) {
	err := cc.db.Update(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func NewContactRepository(db *gorm.DB) repository.ContactRepository {
	return &contactRepository{db}
}
