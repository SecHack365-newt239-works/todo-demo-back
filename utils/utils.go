package utils

import "os"

func GetDBConfig() (string, string, string, string) {
	host := os.Getenv("POSTGRES_HOSTNAME")
	database := os.Getenv("POSTGRES_DB")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	return host, database, user, password
}
