package config

import (
	"crypto/hkdf"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/crypto/chacha20"
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

func MustGetenvBase64(name string) []byte {
	value := MustGetenv(name)
	binary, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		log.Fatalf("FATAL: Failed to decode environment variable %s from Base64!", name)
	}
	return binary
}

const KeyDerivationSalt = "053754807554f8a549896318ffba62abea48aa4c5006f949fcf9cf918270ef8cd57dfbdda5432cbb06867a3d8e2eb074c71febac5ba0f6efb357334b58e98801"

func MustDeriveKey(info string, keySize int) []byte {
	salt, _ := hex.DecodeString(KeyDerivationSalt)
	key, err := hkdf.Key(sha256.New, SecretKeyBase, salt, info, keySize)
	if err != nil {
		panic(err) // This should never happen
	}
	return key
}

const SessionCookieName = "_kashanka_session"
const SessionMaxAge = 24 * time.Hour

var AppHost = GetenvWithDefault("HOST", "::")
var AppPort = GetenvWithDefault("PORT", "3000")
var AppHostPort = net.JoinHostPort(AppHost, AppPort)
var DatabaseUrl = MustGetenv("DATABASE_URL")

var TestDatabaseUrl = os.Getenv("TEST_DATABASE_URL")
var ViteDevServer = GetenvWithDefault("VITE_DEV_SERVER", "https://assets.localhost")
var SecretKeyBase = MustGetenvBase64("SECRET_KEY_BASE")
var SessionSigner = MustDeriveKey("kashanka/v1/session-signer", chacha20.KeySize)
