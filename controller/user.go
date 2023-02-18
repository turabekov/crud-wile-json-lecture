package controller

import (
	"app/models"
	"errors"
)

func (c *Controller) CreateUser(req *models.CreateUser) (id string, err error) {

	id, err = c.store.User().Create(req)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (c *Controller) GetList(req *models.GetListRequest) (*models.GetListResponse, error) {

	users, err := c.store.User().GetList(req)
	if err != nil {
		return &models.GetListResponse{}, err
	}

	return users, nil
}

func (c *Controller) GetUserByIdController(req *models.UserPrimaryKey) (models.User, error) {
	user, err := c.store.User().GetUserById(req)
	if err != nil {
		return models.User{}, err
	}

	return user, nil

}

func (c *Controller) UpdateUserController(req *models.UpdateUser) (models.User, error) {
	user, err := c.store.User().UpdateUser(req)
	if err != nil {
		return models.User{}, err
	}

	return user, nil

}
func (c *Controller) DeleteUserController(req *models.UserPrimaryKey) (models.User, error) {
	user, err := c.store.User().DeleteUser(req)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (c *Controller) WithdrawUserBalance(id string, balance float64) error {

	user, err := c.store.User().GetUserById(&models.UserPrimaryKey{Id: id})
	if err != nil {
		return err
	}

	if user.Balance-balance <= 0 {
		return errors.New("not enough money")
	}

	user.Balance = user.Balance - balance

	_, err = c.store.User().UpdateUser(&models.UpdateUser{
		Id:      user.Id,
		Name:    user.Name,
		Surname: user.Surname,
		Balance: user.Balance,
	})

	if err != nil {
		return err
	}

	p, e := c.store.ShopCart().GetUserShopCarts(&models.UserPrimaryKey{
		Id: user.Id,
	})
	if e != nil {
		return e
	}

	for _, v := range p {
		_, err := c.store.ShopCart().UpdateShopCart(models.ShopCart{
			Id:        v.Id,
			ProductId: v.ProductId,
			UserId:    v.UserId,
			Count:     v.Count,
			Status:    true,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
