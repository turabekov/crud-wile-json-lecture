package storage

import "app/models"

type StorageI interface {
	CloseDB()
	User() UserRepoI
}

type UserRepoI interface {
	Create(*models.CreateUser) (string, error)
	GetUserById(req *models.UserPrimaryKey) (models.User, error)
	GetList(req *models.GetListRequest) (*models.GetListResponse, error)
	UpdateUser(req *models.UpdateUser) (models.User, error)
	DeleteUser(req *models.UserPrimaryKey) (models.User, error)
}
