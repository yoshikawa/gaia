package config

import (
	"errors"
	"os"
	"strconv"

	dotenv "github.com/joho/godotenv"
)

var (
	// Port config
	Port int
	// DBUser config
	DBUser string
	// DBPassword config
	DBPassword string
	// DBProtocol config
	DBProtocol string
	// DBHost config
	DBHost string
	// DBPort config
	DBPort string
	// DBName config
	DBName string
	// DBCharset config
	DBCharset string
	// DBParseTime config
	DBParseTime string
)

// Load function will load all config from environment variable
func Load() error {
	// load .env
	err := dotenv.Load(".env")
	if err != nil {
		return errors.New(".env is not loaded properly")
	}

	portStr, ok := os.LookupEnv("PORT")
	if !ok {
		return errors.New("PORT env is not loaded")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return err
	}
	Port = port

	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return errors.New("DB_HOST env is not loaded")
	}
	// set DBHost
	DBHost = dbHost

	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return errors.New("DB_NAME env is not loaded")
	}
	// set DBName
	DBName = dbName

	dbUser, ok := os.LookupEnv("DB_USER")
	if !ok {
		return errors.New("DB_USER env is not loaded")
	}
	// set DBUser
	DBUser = dbUser

	dbPassword, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		return errors.New("DB_PASSWORD env is not loaded")
	}
	// set DBPassword
	DBPassword = dbPassword

	dbPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return errors.New("DB_PORT env is not loaded")
	}
	// set DBPort
	DBPort = dbPort

	dbCharset, ok := os.LookupEnv("DB_CHARSET")
	if !ok {
		return errors.New("DB_CHARSET env is not loaded")
	}
	// set DBCharset
	DBCharset = dbCharset

	dbProtocol, ok := os.LookupEnv("DB_PROTOCOL")
	if !ok {
		return errors.New("DB_PROTOCOL env is not loaded")
	}
	// set DBProtocol
	DBProtocol = dbProtocol

	dbParseTime, ok := os.LookupEnv("DB_PARSETIME")
	if !ok {
		return errors.New("DB_PARSETIME env is not loaded")
	}
	// set DBParseTime
	DBParseTime = dbParseTime

	return nil
}
