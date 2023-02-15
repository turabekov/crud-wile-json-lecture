package controller

import (
	"app/models"
)

func (c *Controller) CreateUser(req *models.CreateUser) (id int, err error) {

	id, err = c.store.User.Create(req)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (c *Controller) GetList(req *models.GetListRequest) (*models.GetListResponse, error) {

	users, err := c.store.User.GetList(req)
	if err != nil {
		return &models.GetListResponse{}, err
	}

	return users, nil
}

func (c *Controller) GetUserByIdController(req *models.UserPrimaryKey) (models.User, error) {
	user, err := c.store.User.GetUserById(req)
	if err != nil {
		return models.User{}, err
	}

	return user, nil

}

func (c *Controller) UpdateUserController(req *models.UpdateUser) (models.User, error) {
	user, err := c.store.User.UpdateUser(req)
	if err != nil {
		return models.User{}, err
	}

	return user, nil

}
func (c *Controller) DeleteUserController(req *models.UserPrimaryKey) (models.User, error) {
	user, err := c.store.User.DeleteUser(req)
	if err != nil {
		return models.User{}, err
	}

	return user, nil

}
