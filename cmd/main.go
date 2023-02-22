package main

import (
	"app/config"
	"app/controller"
	"app/models"
	"app/storage/jsonDb"
	"fmt"
	"log"
)

func main() {
	cfg := config.Load()

	jsonDb, err := jsonDb.NewFileJson(&cfg)
	if err != nil {
		log.Fatal("error while connecting to database")
	}
	defer jsonDb.CloseDb()

	c := controller.NewController(&cfg, jsonDb)

	// Category(c)
	shopCart(c)
	// User(c)

}

func Category(c *controller.Controller) {
	// Create category
	// c.CreateCategory(&models.CreateCategory{
	// 	Name:     "Kompyuterlar",
	// 	ParentID: "eed2e676-1f17-429f-b75c-899eda296e65",
	// })

	// Get category by Id
	// category, err := c.GetByIdCategory(&models.CategoryPrimaryKey{Id: "eed2e676-1f17-429f-b75c-899eda296e65"})
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// fmt.Println(category)

}

func User(c *controller.Controller) {

	// sender := "bbda487b-1c0f-4c93-b17f-47b8570adfa6"
	// receiver := "657a41b6-1bdc-47cc-bdad-1f85eb8fb98c"
	// err := c.MoneyTransfer(sender, receiver, 500_000)
	// if err != nil {
	// 	log.Println(err)
	// }
	err := c.WithdrawCheque(11200, "a87b3ac7-1579-4cb9-958b-31118acffd56")
	if err != nil {
		log.Fatal(err)
	}
}

func Product(c *controller.Controller) {
	// c.CreateProduct(&models.CreateProduct{
	// 	Name:       "Smartfon vivo V25 8/256 GB",
	// 	Price:      4_860_000,
	// 	CategoryID: "6325b81f-9a2b-48ef-8d38-5cef642fed6b",
	// })

	// product, err := c.GetByIdProduct(&models.ProductPrimaryKey{Id: "38292285-4c27-497b-bc5f-dfe418a9f959"})

	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// ps, err := c.GetAllProduct(&models.ReqGetListProduct{
	// 	Offset:     0,
	// 	Limit:      1000,
	// 	CategoryID: "6325b81f-9a2b-48ef-8d38-5cef642fed6b",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v\n", ps)
}

func shopCart(c *controller.Controller) {
	// Add data to shop cart
	// sh, e := c.AddShopCart(&models.Add{
	// 	ProductId: "ce06cf2e-6577-46cb-96a3-1cea379bde4b",
	// 	UserId:    "a87b3ac7-1579-4cb9-958b-31118acffd56",
	// 	Count:     4,
	// })
	// if e != nil {
	// 	log.Fatal(e)
	// }
	// fmt.Println("Shop cart added", sh)
	total, _ := c.CalculateTotal(&models.UserPrimaryKey{
		Id: "a87b3ac7-1579-4cb9-958b-31118acffd56",
	}, "fixed", 0)
	fmt.Println(total)

	// Statistika
	err := c.StatistikaInShopCart()
	if err != nil {
		log.Fatal(err)
	}
}
