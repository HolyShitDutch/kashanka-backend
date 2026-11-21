package config

import (
	"encoding/base64"
	"encoding/hex"
	"log"
	"net"
	"os"
	"slices"
	"strings"
	"time"

	"golang.org/x/crypto/chacha20"
	"lukechampine.com/blake3"
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

var keyDerivationSalt, err = hex.DecodeString("053754807554f8a549896318ffba62abea48aa4c5006f949fcf9cf918270ef8cd57dfbdda5432cbb06867a3d8e2eb074c71febac5ba0f6efb357334b58e98801")

func MustDeriveKey(secretKeyBase, salt []byte, info string, keySize int) []byte {
	key := make([]byte, keySize)
	ikm := slices.Concat(secretKeyBase, salt)
	blake3.DeriveKey(key[:], info, ikm)
	return key
}

const SessionCookieName = "_kashanka_session"
const SessionMaxAge = 24 * time.Hour

var AppEnv = GetenvWithDefault("APP_ENV", "dev")
var AppHost = GetenvWithDefault("HOST", "::")
var AppPort = GetenvWithDefault("PORT", "5000")
var AppHostPort = net.JoinHostPort(AppHost, AppPort)
var DatabaseUrl = MustGetenv("DATABASE_URL")

var IsProd = AppEnv == "prod"
var IsDev = AppEnv == "dev"

var TestDatabaseUrl = os.Getenv("TEST_DATABASE_URL")
var ViteDevServer = GetenvWithDefault("VITE_DEV_SERVER", "https://assets.localhost")
var SecretKeyBase = MustGetenvBase64("SECRET_KEY_BASE")
var SessionEncryptionKey = MustDeriveKey(SecretKeyBase, keyDerivationSalt, "github.com/HolyShitDutch/kashanka-backend 2026-11 session encryption key", chacha20.KeySize)

// Webauthn Settings
const RPDisplayName = "卡閃卡 CardFlashCard"

var RPID = GetenvWithDefault("RELYING_PARTY_ID", "kashanka.localhost")
var RPOrigins = strings.Split(GetenvWithDefault("RELYING_PARTY_ORIGINS", "https://kashanka.localhost"), ",")
