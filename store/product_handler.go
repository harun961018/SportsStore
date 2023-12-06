package store

import (
	"math"
	"platform/http/actionresults"
	"platform/http/handling"
	"sportsstore/models"
)

const pageSize = 4

type ProductHandler struct {
	Repository   models.Repository
	URLGenerator handling.URLGenerator
}
type ProductTemplateContext struct {
	Products    []models.Product
	Page        int
	PageCount   int
	PageNumbers []int
	PageUrlFunc func(int) string
}

func (handler ProductHandler) GetProducts(page int) actionresults.ActionResult {
	prods, total := handler.Repository.GetProductPage(page, pageSize)
	pageCount := int(math.Ceil(float64(total) / float64(pageSize)))
	return actionresults.NewTemplateAction("product_list.html",
		ProductTemplateContext{
			Products:    prods,
			Page:        page,
			PageCount:   pageCount,
			PageNumbers: handler.generatePageNumbers(pageCount),
			PageUrlFunc: handler.createPageUrlFunction(),
		})
}
func (handler ProductHandler) createPageUrlFunction() func(int) string {
	return func(page int) string {
		url, _ := handler.URLGenerator.GenerateUrl(ProductHandler.GetProducts, page)
		return url
	}
}
func (handler ProductHandler) generatePageNumbers(pageCount int) (pages []int) {
	pages = make([]int, pageCount)
	for i := 0; i < pageCount; i++ {
		pages[i] = i + 1
	}
	return
}
