package dto

type Client struct {
	ID    *int   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
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
