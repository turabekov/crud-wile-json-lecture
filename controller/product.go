package controller

import "app/models"

func (c *Controller) CreateProduct(req *models.CreateProduct) (string, error) {
	id, err := c.store.Product().Create(req)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (c *Controller) DeleteProduct(req *models.ProductPrimaryKey) error {
	err := c.store.Product().Delete(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) UpdateProduct(req *models.UpdateProduct, productId string) error {
	err := c.store.Product().Update(req, productId)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) GetByIdProduct(req *models.ProductPrimaryKey) (models.Product, error) {
	product, err := c.store.Product().GetByID(req)
	if err != nil {
		return models.Product{}, err
	}

	category, err := c.store.Category().GetByID(&models.CategoryPrimaryKey{Id: product.CategoryID})
	if err != nil {
		return models.Product{}, err
	}

	return models.Product{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Category: category,
	}, nil
}

func (c *Controller) GetAllProduct(req *models.ReqGetListProduct) ([]models.Product, error) {
	products, err := c.store.Product().GetAll(req)
	if err != nil {
		return []models.Product{}, err
	}

	category, err := c.store.Category().GetByID(&models.CategoryPrimaryKey{
		Id: req.CategoryID,
	})
	if err != nil {
		return []models.Product{}, err
	}

	arr := []models.Product{}
	for _, v := range products.Products {
		if v.CategoryID == req.CategoryID {
			arr = append(arr, models.Product{
				Id:       v.Id,
				Name:     v.Name,
				Price:    v.Price,
				Category: category,
			})
		}
	}

	if req.Limit+req.Offset > len(arr) {
		if req.Offset > len(arr) {
			return []models.Product{}, nil
		}

		return arr[req.Offset:], nil
	}

	return arr[req.Offset : req.Limit+req.Offset], nil
}
