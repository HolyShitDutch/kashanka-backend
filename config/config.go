package config

import (
	"log"
	"net"
	"os"
)

func MustGetenv(name string) string {
	val := os.Getenv(name)
	if val == "" {
		log.Fatalf("FATAL: Environment variable %s is not set!", name)
	}
	return val
}

func GetenvWithDefault(name, defaultValue string) string {
	val := os.Getenv(name)
	if val == "" {
		return defaultValue
	}
	return val
}

var AppHost = GetenvWithDefault("HOST", "::")
var AppPort = GetenvWithDefault("PORT", "3000")
var AppHostPort = net.JoinHostPort(AppHost, AppPort)
var DatabaseUrl = MustGetenv("DATABASE_URL")

var TestDatabaseUrl = os.Getenv("TEST_DATABASE_URL")

var ViteDevServer = GetenvWithDefault("VITE_DEV_SERVER", "http://localhost:5173")
