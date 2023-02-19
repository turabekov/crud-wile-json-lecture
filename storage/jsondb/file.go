package jsondb

import (
	"app/config"
	"app/storage"
	"os"
)

type Store struct {
	user      *userRepo
	product   *productRepo
	shopCart  *shopCartRepo
	komissiya *komissiyaRepo
}

func NewFileJson(cfg *config.Config) (storage.StorageI, error) {

	userFile, err := os.Open(cfg.Path + cfg.UserFileName)
	if err != nil {
		return nil, err
	}

	productFile, err := os.Open(cfg.Path + cfg.ProductFileName)
	if err != nil {
		return nil, err
	}

	shopCartFile, err := os.Open(cfg.Path + cfg.ProductFileName)
	if err != nil {
		return nil, err
	}

	komissiyaFile, err := os.Open(cfg.Path + cfg.KomissiyaFileName)
	if err != nil {
		return nil, err
	}

	return &Store{
		user:      NewUserRepo(cfg.Path+cfg.UserFileName, userFile),
		product:   NewProductRepo(cfg.Path+cfg.ProductFileName, productFile),
		shopCart:  NewShopCartRepo(cfg.Path+cfg.ShopCartFileName, shopCartFile),
		komissiya: NewKomissiyaRepo(cfg.Path+cfg.KomissiyaFileName, komissiyaFile),
	}, nil
}

func (s *Store) CloseDB() {
	s.user.file.Close()
	s.product.file.Close()
	s.shopCart.file.Close()
	s.komissiya.file.Close()
}

func (s *Store) User() storage.UserRepoI {
	return s.user
}

func (s *Store) Product() storage.ProductRepoI {
	return s.product
}

func (s *Store) ShopCart() storage.ShopCartRepoI {
	return s.shopCart
}

func (s *Store) Komissiya() storage.KomissiyaRepoI {
	return s.komissiya
}
