package client

type ClientRepository interface {
	Create(client Client) error
	GetByID(id int) (*Client, error)
	Delete(id int) error
	Update(client Client) error
}
