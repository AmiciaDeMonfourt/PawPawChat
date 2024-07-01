package config

type Config struct {
	App struct {
		AuthServiceAddr string
	}

	Auth struct {
		UsersServiceAddr string
	}
}

func New() *Config {
	var cfg Config

	cfg.App.AuthServiceAddr = ":50052"
	cfg.Auth.UsersServiceAddr = ":50051"

	return &cfg
}
