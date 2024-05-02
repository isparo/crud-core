package model

type Client struct {
	ID    *int
	Name  string
	Email string
}

func NewClient(
	id *int,
	name string,
	email string,
) Client {
	return Client{
		ID:    id,
		Name:  name,
		Email: email,
	}
}
