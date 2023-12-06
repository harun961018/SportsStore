package models

type Repository interface {
	GetProduct(id int) Product
	GetProducts() []Product
	GetCategories() []Category
	Seed()
}
