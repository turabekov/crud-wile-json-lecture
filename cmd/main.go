package main

import (
	"fmt"
	"log"

	"app/config"
	"app/controller"
	"app/models"
	"app/storage"
)

func main() {

	cfg := config.Load()

	store, err := storage.NewFileJson(&cfg)
	if err != nil {
		panic("error while connect to json file: " + err.Error())
	}

	c := controller.NewController(&cfg, store)

	// Create User
	// id, err := c.CreateUser(
	// 	&models.CreateUser{
	// 		Name:    "Abduqodir",
	// 		Surname: "Musayev",
	// 	},
	// )

	// if err != nil {
	// 	log.Println("error while CreateUser:", err.Error())
	// 	return
	// }

	// fmt.Println(id)

	// GetList of user
	res, err := c.GetList(
		&models.GetListRequest{
			Offset: 0,
			Limit:  100,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.Users)

	// Get user by id
	user, err := c.GetUserByIdController(
		&models.UserPrimaryKey{
			Id: 5,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Get by id", user)

	// update user
	user, err = c.UpdateUserController(
		&models.UpdateUser{
			Id:      7,
			Name:    "Wayne",
			Surname: "Rooney",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("updated user", user)

	// Delete user
	user, err = c.DeleteUserController(
		&models.UserPrimaryKey{
			Id: 19,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted user", user)

}

// GetByPkey() Done
// GetList() Done
// Update(user)
// Delete()
