package config

type Config struct {
	Lusername      string `env:"Lusername"`
	Lpassword      string `env:"Lpassword"`
	GoogleClientId string `env:"Google_client_id"`
}
