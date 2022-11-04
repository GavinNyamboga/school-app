package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{

			Dialect:  goDotEnvVariable("DIALECT"), //mysql
			Host:     goDotEnvVariable("DB_HOST"), //"127.0.0.1",
			Port:     3306,
			Username: goDotEnvVariable("DB_USERNAME"), //"admin",
			Password: goDotEnvVariable("DB_PASSWORD"), //"admin",
			Name:     goDotEnvVariable("DB_NAME"),     //"school_app",
			Charset:  goDotEnvVariable("CHARSET"),     //"utf8",
		},
	}
}

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load("properties.env")

	if err != nil {
		log.Fatalf("Error loading properties.env file")
	}

	return os.Getenv(key)
}
