package controller

import "courseonline/services"

type CartController struct {
	storedb services.Store
}

func NewCartController(store services.Store) *CartController {
	return &CartController{
		storedb: store,
	}
}
