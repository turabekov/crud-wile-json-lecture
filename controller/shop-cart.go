package controller

import (
	"app/models"
)

func (c *Controller) AddShopCart(req *models.AddShopCart) (models.ShopCart, error) {
	// check existing of current user
	userId := models.UserPrimaryKey{
		Id: req.UserId,
	}
	_, err := c.store.User().GetUserById(&userId)
	if err != nil {
		return models.ShopCart{}, err
	}

	// check existing of current product
	productId := models.ProductPrimaryKey{
		Id: req.ProductId,
	}
	_, err = c.store.Product().GetProductById(&productId)
	if err != nil {
		return models.ShopCart{}, err
	}

	p, err := c.store.ShopCart().AddShopCart(req)
	if err != nil {
		return models.ShopCart{}, err
	}

	return p, nil
}

func (c *Controller) RemoveShopCart(req *models.RemoveShopCart) (models.ShopCart, error) {
	// check existing of current user
	userId := models.UserPrimaryKey{
		Id: req.UserId,
	}
	_, err := c.store.User().GetUserById(&userId)
	if err != nil {
		return models.ShopCart{}, err
	}

	// check existing of current product
	productId := models.ProductPrimaryKey{
		Id: req.ProductId,
	}
	_, err = c.store.Product().GetProductById(&productId)
	if err != nil {
		return models.ShopCart{}, err
	}

	p, err := c.store.ShopCart().RemoveShopCart(req)
	if err != nil {
		return models.ShopCart{}, err
	}

	return p, nil
}

func (c *Controller) GetUserShopCarts(req *models.UserPrimaryKey) ([]models.ShopCart, error) {
	// check existing of current user
	userId := models.UserPrimaryKey{
		Id: req.Id,
	}
	_, err := c.store.User().GetUserById(&userId)
	if err != nil {
		return []models.ShopCart{}, err
	}

	p, e := c.store.ShopCart().GetUserShopCarts(req)
	if e != nil {
		return []models.ShopCart{}, e
	}

	return p, nil
}

func (c *Controller) CalcTotalPrice(req models.UserPrimaryKey) (int, error) {
	// check existing of current user
	userId := models.UserPrimaryKey{
		Id: req.Id,
	}
	_, err := c.store.User().GetUserById(&userId)
	if err != nil {
		return 0, err
	}

	shopCarts, e := c.store.ShopCart().GetUserShopCarts(&userId)
	if e != nil {
		return 0, e
	}

	total := 0

	for _, v := range shopCarts {
		pId := models.ProductPrimaryKey{
			Id: v.ProductId,
		}
		p, e := c.store.Product().GetProductById(&pId)
		if e != nil {
			return 0, e
		}

		total += v.Count * p.Price
	}

	return total, nil
}
