package models

import "time"

type Category struct {
	ID      int       `db:"id"`
	Name    string    `db:"name"`
	Created time.Time `db:"createdAt"`
	Updated time.Time `db:"updatedAt"`
}

func (c Category) ToDTO() CategoryDTO {
	return CategoryDTO{
		ID:   c.ID,
		Name: c.Name,
	}
}

type Product struct {
	ID      int       `db:"id"`
	CatID   int       `db:"categoryId"`
	Name    string    `db:"name"`
	Price   int       `db:"price"`
	Count   int       `db:"count"`
	Created time.Time `db:"createdAt"`
	Updated time.Time `db:"updatedAt"`
}

func (p Product) ToDTO() ProductDTO {
	return ProductDTO{
		ID:    p.ID,
		CatID: p.CatID,
		Name:  p.Name,
		Price: p.Price,
		Count: p.Count,
	}
}
