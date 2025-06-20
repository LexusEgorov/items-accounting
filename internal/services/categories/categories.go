package categories

import "github.com/LexusEgorov/items-accounting/internal/models"

type Storager interface {
	Add(string) (models.Category, error)
	Get(int) (models.Category, error)
	Set(int, string) (models.Category, error)
	Delete(int) error
}

type Categories struct {
	storage Storager
}

//TODO: add validate data (for example)
func New(storage Storager) *Categories {
	return &Categories{
		storage: storage,
	}
}

func (c Categories) Add(name string) (models.Category, error) {
	return c.storage.Add(name)
}

func (c Categories) Set(ID int, name string) (models.Category, error) {
	return c.storage.Set(ID, name)
}

func (c Categories) Get(ID int) (models.Category, error) {
	return c.storage.Get(ID)
}

func (c Categories) Delete(ID int) error {
	return c.storage.Delete(ID)
}
