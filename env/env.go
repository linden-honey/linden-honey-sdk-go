package env

import (
	"fmt"
	"os"
)

func GetEnv(key, defaultValue string) string {
	if key == "" {
		return defaultValue
	}

	if env, ok := os.LookupEnv(key); ok {
		return env
	}

	return defaultValue
}

func GetRequiredEnv(key string) string {
	if key == "" {
		panic("empty key argument")
	}

	if env, ok := os.LookupEnv(key); ok {
		return env
	}

	panic(fmt.Sprintf("key '%s' is required", key))
}
