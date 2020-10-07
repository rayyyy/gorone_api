package util

import (
	"os"
)

// GetEnv ENVを取得する、ない場合はデフォルトを返す
func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

// IsDevelopment 開発環境のときにtrueを返す
func IsDevelopment() bool {
	appEnv := GetEnv("APP_ENV", "")

	if appEnv == "development" {
		return true
	}

	return false
}

// IsProduction 本番環境のときにtrueを返す
func IsProduction() bool {
	appEnv := GetEnv("APP_ENV", "")

	if appEnv == "production" {
		return true
	}

	return false
}
