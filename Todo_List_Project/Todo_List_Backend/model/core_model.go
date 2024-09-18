package model

type Config struct {
	Port                  int    `yaml:"SERVER_PORT"`
	DBHost                string `yaml:"DB_HOST"`
	DBPort                int    `yaml:"DB_PORT"`
	DBUsername            string `yaml:"DB_USERNAME"`
	DBName                string `yaml:"DB_NAME"`
	DBPassword            string `yaml:"DB_PASSWORD"`
}
