package repo

import (
	"platform/services"
	"sportsstore/models"
)

func RegisterMemoryRepoService() {
	services.AddSingleton(func() models.Repository {
		repo := &MemoryRepo{}
		repo.Seed()
		return repo
	})
}

type MemoryRepo struct {
	products   []models.Product
	categories []models.Category
}

func (repo *MemoryRepo) GetProduct(id int) (product models.Product) {
	for _, p := range repo.products {
		if p.ID == id {
			product = p
			return
		}
	}
	return
}
func (repo *MemoryRepo) GetProducts() (results []models.Product) {
	return repo.products
}
func (repo *MemoryRepo) GetCategories() (results []models.Category) {
	return repo.categories
}
