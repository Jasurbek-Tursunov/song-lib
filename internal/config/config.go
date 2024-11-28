package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	Port       int    `env:"PORT"`
	DbHost     string `env:"DB_HOST"`
	DbPort     int    `env:"DB_PORT"`
	DbName     string `env:"DB_NAME"`
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
}

func MustLoad() *Config {
	const op = "config.MustLoad"

	if err := godotenv.Load(); err != nil {
		panic(op + " -> cannot read env:" + err.Error())
	}

	port, err := strconv.ParseInt(os.Getenv("PORT"), 0, 64)
	if err != nil {
		panic(op + " -> env data uncorrected!")
	}

	dbPort, err := strconv.ParseInt(os.Getenv("DB_PORT"), 0, 64)
	if err != nil {
		panic(op + " -> env data uncorrected!")
	}

	cfg := Config{
		Port:       int(port),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     int(dbPort),
		DbName:     os.Getenv("DB_NAME"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
	}

	return &cfg
}
