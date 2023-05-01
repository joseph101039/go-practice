package env

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func init() {
	loadEnv()
}

// Get retrieves the Env value from .env file
func Get(key string) string {
	return os.Getenv(key)
}

func GetDefault(key string, defaultValue string) string {
	value, existed := os.LookupEnv(key)
	if !existed {
		return defaultValue
	}

	return value
}

// GetMany retrieves many Env value from .env file
func GetMany(envKeys []string) map[string]string {
	envs := make(map[string]string)
	for _, key := range envKeys {
		envs[key] = os.Getenv(key)
	}

	return envs
}

// Environment 取得環境別
func Environment() string {
	return os.Getenv("APP_ENV")
}

// IsLocal 是否為本地環境
func IsLocal() bool {
	return Environment() == "local"
}

func loadEnv() {

	_, currentFilePath, _, _ := runtime.Caller(1)
	dir := filepath.Dir(currentFilePath)
	err := godotenv.Load(dir + "/../../.env")
	if err != nil {
		panic(err)
	}
}
