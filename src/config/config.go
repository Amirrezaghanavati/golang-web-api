package config

import (
	"errors"
	"log"
	"os"
	"github.com/spf13/viper"
)


type Config struct{
	Server ServerConfig
	Postgres PostgresConfig
	Redis RedisConfig
}

type ServerConfig struct{
	Port string
	RunMode string
}

type PostgresConfig struct{
	Host string
	Port string
	User string
	Password string
	DbName string
	SSLMode bool
}

type RedisConfig struct {
	Host string
	Post string
	Password string
	Db string
	MinIdleConnections int
	PoolSice int
	PoolTimeout int
}

func GetConfig() *Config{
	configPath := getConfigPath(os.Getenv("APP_ENV"))
	viper, error := LoadConfig(configPath, "yml")
	if error != nil {
		log.Fatalf("Error in load config: %v", error)
	}

	config, error := ParseConfig(viper)
	if error != nil {
		log.Fatalf("Error in parse config: %v", error)
	}

	return config
}

func ParseConfig(v *viper.Viper) (*Config, error){
	var config Config
	error := v.Unmarshal(&config)
	if error != nil {
		log.Printf("Unable to parse config: %v", error)
		return nil, error
	}
	return &config, error
}

func LoadConfig(fileName string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(fileName)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	error := v.ReadInConfig()

	if error != nil {
		log.Printf("Unable to read config: %v", error)
		if _, ok := error.(viper.ConfigFileNotFoundError); ok{
			return nil, errors.New("config file not found")
		}
		return nil, error
	}
	return v, nil
}


func getConfigPath(env string) string {
	switch env {
	case "docker":
		return "config/config-docker"
	case "production":
		return "config/config-production"
	default:
		return "../config/config-development"
	}
}