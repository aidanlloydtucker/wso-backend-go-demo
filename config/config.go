package config

import (
	"flag"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Env          string `yaml:"env"`
	DatabaseType string `yaml:"database_type"`
	DatabaseArgs string `yaml:"database_args"`
	JWTRealm     string `yaml:"jwt_realm"`
	JWTSecretKey string `yaml:"jwt_secret_key"`
	GinMode      string `yaml:"gin_mode"`
}

func (c *Config) IsEnv(env string) bool {
	return c.Env == env
}

func (c *Config) IsDevelopment() bool {
	return c.IsEnv("development")
}

func (c *Config) IsTest() bool {
	return c.IsEnv("test")
}

func (c *Config) IsProduction() bool {
	return c.IsEnv("production")
}

func GetConfig() (*Config, error) {
	var env string
	var configPath string

	flag.StringVar(&env, "env", "development", "environment of server")
	flag.StringVar(&configPath, "config", "", "path to config file")

	flag.Parse()

	if configPath == "" {
		configPath += "config/environment/" + env + ".yml"
	}

	configPath, err := filepath.Abs(configPath)
	if err != nil {
		return nil, err
	}

	configFile, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	cfg := new(Config)

	dec := yaml.NewDecoder(configFile)
	err = dec.Decode(cfg)
	if err != nil {
		return nil, err
	}

	// Config Defaults
	if cfg.GinMode == "" {
		switch cfg.Env {
		case "development":
			cfg.GinMode = "development"
		case "test":
			cfg.GinMode = "test"
		case "production":
			cfg.GinMode = "production"
		default:
			cfg.GinMode = "development"
		}
	}

	return cfg, nil
}
