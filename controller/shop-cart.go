package controller

import (
	"app/models"
	"errors"
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

func (c *Controller) CalcTotalPrice(req models.UserPrimaryKey, d models.Discount) (int, error) {
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

	var total float64
	var discount float64
	if d.Status == "fixed" {
		discount = d.Amount
		for _, v := range shopCarts {
			pId := models.ProductPrimaryKey{
				Id: v.ProductId,
			}
			p, e := c.store.Product().GetProductById(&pId)
			if e != nil {
				return 0, e
			}

			if float64(p.Price)-discount < 0 {
				total = 0
			} else {
				total += float64(v.Count*p.Price) - float64(v.Count)*discount
			}

		}
	} else if d.Status == "percent" {
		if d.Amount >= 0 && d.Amount <= 100 {
			discount = d.Amount / 100
			for _, v := range shopCarts {
				pId := models.ProductPrimaryKey{
					Id: v.ProductId,
				}

				p, e := c.store.Product().GetProductById(&pId)
				if e != nil {
					return 0, e
				}

				if float64(p.Price)*discount == 0 {
					total = float64(v.Count * p.Price)
				} else {
					total += float64(v.Count*p.Price) - float64(p.Price)*discount*float64(v.Count)
				}
			}
		} else {
			return 0, errors.New("enter correct discount for percent")
		}
	} else {
		return 0, errors.New("please enter correct value for status")
	}

	return int(total), nil
}
