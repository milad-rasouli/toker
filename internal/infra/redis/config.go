package redis

type Config struct {
	Host string `json:"host" koanf:"host"`
	Name string `json:"name" koanf:"name"`
}
