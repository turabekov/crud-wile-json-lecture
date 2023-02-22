package controller

import (
	"app/models"
	"errors"
	"fmt"
)

func (c *Controller) AddShopCart(req *models.Add) (string, error) {
	_, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: req.UserId})
	if err != nil {
		return "", err
	}

	_, err = c.store.Product().GetByID(&models.ProductPrimaryKey{Id: req.ProductId})
	if err != nil {
		return "", err
	}

	id, err := c.store.ShopCart().AddShopCart(req)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (c *Controller) RemoveShopCart(req *models.Remove) error {
	err := c.store.ShopCart().RemoveShopCart(req)
	if err != nil {
		return err
	}
	return err
}

func (c *Controller) CalculateTotal(req *models.UserPrimaryKey, status string, discount float64) (float64, error) {
	_, err := c.store.User().GetByID(req)
	if err != nil {
		return 0, err
	}

	users, err := c.store.ShopCart().GetUserShopCart(req)
	if err != nil {
		return 0, err
	}

	var total float64
	for _, v := range users {
		product, err := c.store.Product().GetByID(&models.ProductPrimaryKey{Id: v.ProductId})
		if err != nil {
			return 0, err
		}
		if status == "fixed" {
			total += float64(v.Count) * (product.Price - discount)
		} else if status == "percent" {
			if discount < 0 || discount > 100 {
				return 0, errors.New("invalid discount range")
			}
			total += float64(v.Count) * (product.Price - (product.Price*discount)/100)
		} else {
			return 0, errors.New("invalid status name")
		}
	}

	if total < 0 {
		return 0, nil
	}
	return total, nil
}
func (c *Controller) StatistikaInShopCart() error {
	// =======================================================================================================================
	// statistika
	shopCarts, err := c.store.ShopCart().GetAllShopCarts()
	if err != nil {
		return err
	}

	hash := map[string]int{}

	fmt.Println("Statistika:")
	for _, v := range shopCarts {
		if v.Status {
			product, e := c.store.Product().GetByID(&models.ProductPrimaryKey{
				Id: v.ProductId,
			})
			if e != nil {
				return e
			}

			category, e := c.store.Category().GetByID(&models.CategoryPrimaryKey{
				Id: product.CategoryID,
			})
			if e != nil {
				return e
			}

			val, ok := hash[category.Name]
			if ok {
				hash[category.Name] = val + v.Count
			} else {
				hash[category.Name] = v.Count
			}
		}
	}
	// iterate hash
	for key, v := range hash {
		fmt.Printf("%s - %d\n", key, v)
	}

	// =======================================================================================================================
	return nil
}
