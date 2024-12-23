package config

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	JWT      JwtConfig
	Otp      OtpConfig
}

type ServerConfig struct {
	Port    string
	RunMode string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SSLMode  string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

type RedisConfig struct {
	Host               string
	Port               string
	User               string
	Password           string
	Db                 string
	DialTimeout 			 time.Duration
	ReadTimeout 			 time.Duration
  WriteTimeout 			 time.Duration
	IdleCheckFrequency time.Duration
	PoolSize           int
	PoolTimeout  time.Duration
}

type JwtConfig struct {
	Secret string
  RefreshSecret string
  AccessTokenExpireDuration time.Duration
  RefreshTokenExpireDuration time.Duration
}

type OtpConfig struct{
	Digits int
	ExpireTime time.Duration
}

func GetConfig() *Config {
	cfgPath := getConfig(os.Getenv("APP_ENV"))
	v, err := LoadConfig(cfgPath, "yml")
	if err != nil {
		log.Fatalf("Error in load config, %v", err)
	}

	cfg, err := ParseConfig(v)
	if err != nil {
		log.Fatalf("Error in parse config, %v", err)
	}

	return cfg
}

func ParseConfig(v *viper.Viper) (*Config, error)  {
	var cfg Config
	err := v.Unmarshal(&cfg)
	
	if err != nil {
		log.Printf("Unable to parse config: %v", err)
		return nil, err
	}
	return &cfg, nil
}

func LoadConfig(fileName string, fileType string) (*viper.Viper ,error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(fileName)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	
	if err != nil {
		log.Printf("Unable to read config: %v", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("Config file not found")
		}
		return nil, err
	}
	return v, nil
}	

func getConfig(env string) string {
	if env == "docker" {
		return "config/config-docker"
	} else if env == "production" {
		return "config/config-production"
	} else {
		return "../src/config/config-development"
	}
}