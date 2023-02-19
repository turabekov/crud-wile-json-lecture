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

// Homework
func (c *Controller) ExchangeMoney(req models.ReqExchangeMoney) error {
	sender, err := c.store.User().GetUserById(&models.UserPrimaryKey{
		Id: req.SenderId,
	})
	if err != nil {
		return err
	}
	receiver, err := c.store.User().GetUserById(&models.UserPrimaryKey{
		Id: req.ReceiverId,
	})
	if err != nil {
		return err
	}

	komissiya, err := c.store.Komissiya().GetKomissiya()
	if err != nil {
		return err
	}

	moneyWithKomissiya := float64(req.Amount) + (float64(req.Amount) * komissiya.Amount / 100)

	if sender.Balance > moneyWithKomissiya {
		_, err = c.store.User().UpdateUser(&models.UpdateUser{
			Id:      sender.Id,
			Name:    sender.Name,
			Surname: sender.Surname,
			Balance: sender.Balance - moneyWithKomissiya,
		})
		if err != nil {
			return err
		}
		_, err = c.store.User().UpdateUser(&models.UpdateUser{
			Id:      receiver.Id,
			Name:    receiver.Name,
			Surname: receiver.Surname,
			Balance: receiver.Balance + float64(req.Amount),
		})
		if err != nil {
			return err
		}
		err := c.store.Komissiya().UpdateBalanceKomissiya(models.Komissiya{
			Balance: komissiya.Balance + (float64(req.Amount) * komissiya.Amount / 100),
			Amount:  komissiya.Amount,
		})
		if err != nil {
			return nil
		}

	} else {
		return errors.New("not enough money for sending")
	}

	return nil
}
