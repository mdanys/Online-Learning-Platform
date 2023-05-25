package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort  int
	DBPort      string
	DBUser      string
	DBPass      string
	DBHost      string
	DBName      string
	PathMigrate string
	Database    string
	JWTSecret   string
}

// FUNC TO CREATE NEW CONFIG
func NewConfig() *AppConfig {
	cfg := initConfig()
	if cfg == nil {
		log.Fatal("cannot run configuration setup")
		return nil
	}

	return cfg
}

// FUNC TO INITIALIZE CONFIG FROM ENVIRONMENT
func initConfig() *AppConfig {
	var app AppConfig

	godotenv.Load("config.env")

	serverPortConv, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatal("error parse server port")
		return nil
	}
	app.ServerPort = serverPortConv

	app.DBPort = os.Getenv("DB_PORT")
	app.DBUser = os.Getenv("DB_USERNAME")
	app.DBPass = os.Getenv("DB_PASSWORD")
	app.DBHost = os.Getenv("DB_HOST")
	app.DBName = os.Getenv("DB_NAME")
	app.PathMigrate = os.Getenv("PATH_MIGRATE")
	app.Database = os.Getenv("DATABASE")
	app.JWTSecret = os.Getenv("JWT_SECRET")

	return &app
}
