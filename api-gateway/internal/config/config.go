package config

type Config struct {
	HTTPPort             string
	InventoryServiceAddr string
	OrderServiceAddr     string
	UserServiceAddr      string
}

func NewConfig() *Config {
	return &Config{
		HTTPPort:             "8000",
		InventoryServiceAddr: "localhost:8080",
		OrderServiceAddr:     "localhost:8081",
		UserServiceAddr:      "localhost:50051",
	}
}
