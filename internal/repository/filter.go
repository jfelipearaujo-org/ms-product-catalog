package repository

type GetAllCategoriesFilter struct {
	Title string
}

type GetAllProductsFilter struct {
	Title         string
	CategoryTitle string
}
