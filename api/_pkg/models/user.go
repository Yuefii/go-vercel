package models

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var Users = []User{
	{ID: "1", Name: "John Doe", Email: "john@doe.com"},
	{ID: "2", Name: "Jane Doe", Email: "jane@doe.com"},
	{ID: "3", Name: "Bob Smith", Email: "bob@smith.com"},
}
