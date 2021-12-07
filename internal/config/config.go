package config

import "os"

const (
	HeaderKey = "X-Greenhouse-Key"

	envPrefix       = "GREENHOUSE_"
	EnvKey          = envPrefix + "KEY"
	EnvInfluxUrl    = envPrefix + "INFLUX_URL"
	EnvInfluxKey    = envPrefix + "INFLUX_KEY"
	EnvInfluxOrg    = envPrefix + "INFLUX_ORG"
	EnvInfluxBucket = envPrefix + "INFLUX_BUCKET"
)

func GetKey() string {
	return os.Getenv(EnvKey)
}

func GetInfluxUrl() string {
	return os.Getenv(EnvInfluxUrl)
}

func GetInfluxKey() string {
	return os.Getenv(EnvInfluxKey)
}

func GetInfluxOrg() string {
	return os.Getenv(EnvInfluxOrg)
}

func GetInfluxBucket() string {
	return os.Getenv(EnvInfluxBucket)
}
