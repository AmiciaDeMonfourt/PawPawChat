package config

type config struct {
	UsersAddr string
	AuthAddr  string
	AppAddr   string
}

func App() *config {
	return &config{
		UsersAddr: ":50051",
		AuthAddr:  ":50052",
		AppAddr:   ":8080",
	}
}
