package env

import (
	"fmt"
	"os"
)

// GetEnv returns the value of the environment variable named by the key or default value if missing
func GetEnv(key, defaultValue string) string {
	if key == "" {
		return defaultValue
	}

	if env, ok := os.LookupEnv(key); ok {
		return env
	}

	return defaultValue
}

// GetEnv returns the value of the environment variable named by the key or panic if missing
func GetRequiredEnv(key string) string {
	if key == "" {
		panic("empty key argument")
	}

	if env, ok := os.LookupEnv(key); ok {
		return env
	}

	panic(fmt.Sprintf("key '%s' is required", key))
}
