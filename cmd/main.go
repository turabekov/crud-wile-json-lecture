package main

import (
	"fmt"
	"log"

	"app/config"
	"app/controller"
	"app/models"
	"app/storage/jsondb"
)

func main() {

	cfg := config.Load()

	jsondb, err := jsondb.NewFileJson(&cfg)
	if err != nil {
		panic("error while connect to json file: " + err.Error())
	}

	c := controller.NewController(&cfg, jsondb)

	// Create User
	// id, err := c.CreateUser(
	// 	&models.CreateUser{
	// 		Name:    "Khumoyun",
	// 		Surname: "Turaekov",
	// 	},
	// )

	// if err != nil {
	// 	log.Println("error while CreateUser:", err.Error())
	// 	return
	// }

	// fmt.Println(id)

	// GetList of user
	// res, err := c.GetList(
	// 	&models.GetListRequest{
	// 		Offset: 0,
	// 		Limit:  100,
	// 	},
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(res.Users)

	// Get user by id
	// user, err := c.GetUserByIdController(
	// 	&models.UserPrimaryKey{
	// 		Id: 5,
	// 	},
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Get by id", user)

	// update user
	// user, err := c.UpdateUserController(
	// 	&models.UpdateUser{
	// 		Id:      "",
	// 		Name:    "Wayne",
	// 		Surname: "Rooney",
	// 	},
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("updated user", user)

	// Delete user
	// user, err = c.DeleteUserController(
	// 	&models.UserPrimaryKey{
	// 		Id: "18",
	// 	},
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("deleted user", user)

	// ==========Product========================================================================================================================
	// Create Product
	// id, err := c.CreateProduct(&models.CreateProduct{
	// 	Name:  "Sprite",
	// 	Price: 10000,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(id)

	// Get all products
	products, err := c.GetListProducts(&models.GetListProductRequest{
		Offset: 0,
		Limit:  2,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Get all products", products.Products)

	// Get one product
	p, e := c.GetProductByIdController(&models.ProductPrimaryKey{
		Id: "48b934e9-ed15-4779-8d0d-e45c61c7a089",
	})
	if e != nil {
		log.Fatal(err)
	}
	fmt.Println("Get one by id", p)

	// Update products
	product, err := c.UpdateProductController(&models.UpdateProduct{
		Id:    "ec529cd6-dbb8-4982-a984-017b6a042378",
		Name:  "Dena",
		Price: 15000,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated product", product)
	// Delete product
	product, err = c.DeleteProductController(&models.ProductPrimaryKey{
		Id: "cba2bbf9-4893-409b-be52-20ad631330fe",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted product", product)

}
