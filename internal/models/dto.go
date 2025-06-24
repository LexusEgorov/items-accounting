package models

type CategoryDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProductDTO struct {
	ID    int    `json:"id"`
	CatID int    `json:"categoryId"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Count int    `json:"count"`
}
