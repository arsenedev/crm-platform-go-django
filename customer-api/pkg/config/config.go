package config

import "os"

type Config struct {
	Port             string
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	JWTSecret        string
	GIN_MODE         string
}

func Load() Config {

	// a utiliser sur linux
	return Config{
		Port:             getEnv("PORT", "8080"),
		PostgresHost:     getEnv("POSTGRES_HOST", "postgres"),
		PostgresPort:     getEnv("POSTGRES_PORT", "5432"),
		PostgresUser:     getEnv("POSTGRES_USER", "postgres"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", "postgres"),
		PostgresDB:       getEnv("POSTGRES_DB", "customerdb"),
		JWTSecret:        getEnv("JWT_SECRET", "mysecretkey"),
		GIN_MODE:         getEnv("GIN_MODE", "release"),
	}

	// return Config{
	// 	Port:             "8080",
	// 	PostgresHost:     "host.docker.internal", // "localhost",
	// 	PostgresPort:     "5435",
	// 	PostgresUser:     "postgres",
	// 	PostgresPassword: "postgresql16",
	// 	PostgresDB:       "postgres",
	// 	JWTSecret:        "my-super-secret-key",
	// }
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
