package config

type Config struct {
	Path string

	UserFileName    string
	ProductFileName string
}

func Load() Config {

	cfg := Config{}

	cfg.Path = "./data"
	cfg.UserFileName = "/customer.json"
	cfg.ProductFileName = "/products.json"

	return cfg
}
