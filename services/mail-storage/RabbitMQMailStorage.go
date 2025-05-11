package mailstorage

import (
	"log"

	"github.com/psenna/go-mail/model"
)

type rabbitMQConfig struct {
	Server    string `env:"RABBITMQ_SERVER"`
	Port      int    `env:"RABBITMQ_PORT"`
	User      string `env:"RABBITMQ_USER"`
	Password  string `env:"RABBITMQ_PASS"`
	QueueName string `env:"RABBITMQ_QUEUE_NAME"`
}

type RabbitMQMailStorage struct {
}

func (r *RabbitMQMailStorage) Init() error {
	// Dummy implementation
	log.Println("RabbitMQMailStorage initialized")
	return nil
}

func (r *RabbitMQMailStorage) Close() error {
	// Dummy implementation
	log.Println("RabbitMQMailStorage closed")
	return nil
}

func (r *RabbitMQMailStorage) CreateMail(mail model.Mail) error {
	// Dummy implementation
	log.Println("RabbitMQMailStorage closed")
	return nil
}
