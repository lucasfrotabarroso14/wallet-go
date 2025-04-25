package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Client struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Accounts  []*Account `json:"accounts"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func NewClient(name string, email string) (*Client, error) {
	client := &Client{
		ID:        uuid.NewString(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := client.validate(); err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) validate() error {
	if c.Name == "" {
		return errors.New("Name cannot be empty")
	}
	if c.Email == "" {
		return errors.New("Email cannot be empty")
	}
	return nil

}

func (C *Client) Update(name, email string) error {
	if name == "" {
		return errors.New("Name cannot be empty")
	}
	if email == "" {
		return errors.New("Email cannot be empty")
	}
	C.Name = name
	C.Email = email
	C.UpdatedAt = time.Now()
	if err := C.validate(); err != nil {
		return err
	}
	return nil
}

func (c *Client) AddAccount(account *Account) error {
	if account.Client.ID != c.ID {
		return errors.New("Account does not belong to this client")
	}
	c.Accounts = append(c.Accounts, account)
	return nil
}
