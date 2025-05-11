package mailstorage

import "github.com/psenna/go-mail/model"

type MailStorage interface {
	// Init initializes the storage connection
	Init() error
	// Close closes the storage connection
	Close() error
	// CreateMail creates a new mail in the storage
	CreateMail(mail model.Mail) error
	// GetNextMail retrieves next mail to be sent from the storage
	GetNextMail() (model.Mail, error)
}
