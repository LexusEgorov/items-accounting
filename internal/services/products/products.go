package products

import "github.com/LexusEgorov/items-accounting/internal/models"

type Storager interface {
	Add(string) (models.Product, error)
	Get(int) (models.Product, error)
	Set(models.Product) (models.Product, error)
	Delete(int) error
}

type Products struct {
	storage Storager
}

//TODO: add validate data (for example)
func New(storage Storager) *Products {
	return &Products{
		storage: storage,
	}
}

func (c Products) Add(name string) (models.Product, error) {
	return c.storage.Add(name)
}

func (c Products) Set(product models.Product) (models.Product, error) {
	return c.storage.Set(product)
}

func (c Products) Get(ID int) (models.Product, error) {
	return c.storage.Get(ID)
}

func (c Products) Delete(ID int) error {
	return c.storage.Delete(ID)
}
