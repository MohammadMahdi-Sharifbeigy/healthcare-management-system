package config

type Config struct {
	ServerPort    int
	Environment   string
	DBHost        string
	DBPort        int
	DBUser        string
	DBPassword    string
	DBName        string
	DBSSLMode     string
	LogLevel      string
	LogFormat     string
	JWTSecretKey  string
	JWTExpiration int
}

func LoadConfig() (*Config, error) {
	return &Config{
		ServerPort:    8080,
		Environment:   "development",
		DBHost:        "localhost",
		DBPort:        5432,
		DBUser:        "postgres",
		DBPassword:    "postgres",
		DBName:        "healthcare_db",
		DBSSLMode:     "disable",
		LogLevel:      "debug",
		LogFormat:     "json",
		JWTSecretKey:  "your-secret-key",
		JWTExpiration: 24,
	}, nil
}
