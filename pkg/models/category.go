package models

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var Categories = []Category{
	{ID: "1", Name: "Makanan"},
	{ID: "2", Name: "Cemilan"},
	{ID: "3", Name: "Minuman"},
}
