package models

type ShopCartPrimaryKey struct {
	Id string `json:"id"`
}

type AddShopCart struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Count     int    `json:"count"`
}

type ShopCart struct {
	Id        string `json:"id"`
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Count     int    `json:"count"`
}

type RemoveShopCart struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
}

type UserProductIds struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
}

type Discount struct {
	Status string  `json:"status"`
	Amount float64 `json:"amount"`
}

// type GetListShopCartRequest struct {
// 	Offset int `json:"offset"`
// 	Limit  int `json:"limit"`
// }

// type GetListShopCartResponse struct {
// 	Count     int        `json:"count"`
// 	ShopCarts []ShopCart `json:"products"`
// }
