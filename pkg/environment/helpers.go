package environment

import "os"

func getEnvDefault(key string, defaultValue string) string {
	value, found := os.LookupEnv(key)
	if !found {
		return defaultValue
	}
	return value
}
