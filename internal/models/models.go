package models

type Category struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Created string `json:"createdAt"`
	Updated string `json:"updatedAt"`
}

type Product struct {
	ID      int    `json:"id"`
	CatID   string `json:"categoryId"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
	Created string `json:"createdAt"`
	Updated string `json:"updatedAt"`
}
