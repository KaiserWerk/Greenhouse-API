package config

import "os"

const (
	HeaderKey = "X-Greenhouse-Key"

	envPrefix = "GREENHOUSE_"
	EnvKey    = envPrefix + "KEY"
)

func GetKey() string {
	return os.Getenv(EnvKey)
}
