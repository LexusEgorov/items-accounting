package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"

	"github.com/LexusEgorov/items-accounting/internal/models"
)

type ServerConfig struct {
	Port            int           `yaml:"port"`
	Addr            string        `yaml:"address"`
	MaxResponseTime time.Duration `yaml:"maxResponseTime"`
}

type DBConfig struct {
	User     string `yaml:"user"`
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
}

type Config struct {
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"db"`
}

func New() (cfg *Config, err error) {
	configPath, err := fetchConfigPath()
	if err != nil {
		if !errors.Is(err, models.ErrConfigPathNotProvided) {
			return nil, fmt.Errorf("read config error: %v", err)
		}

		cfg, err = readEnvConfig()
	} else {
		cfg, err = readFileConfig(configPath)
	}
	if err != nil {
		return nil, err
	}

	err = checkConfig(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// Валидирует конфиг
func checkConfig(cfg *Config) error {
	if err := checkServerConfig(&cfg.Server); err != nil {
		return err
	}

	return checkDBConfig(&cfg.DB)
}

// Валидирует конфиг базы данных
func checkDBConfig(cfg *DBConfig) error {
	if cfg.Name == "" {
		return models.ErrBadDBName
	}

	if cfg.Password == "" {
		return models.ErrBadPassword
	}

	if cfg.User == "" {
		return models.ErrBadUserName
	}

	return nil
}

// Валидирует серверный конфиг
func checkServerConfig(cfg *ServerConfig) error {
	if cfg.Port <= 0 {
		return models.ErrBadConfigPort
	}

	if cfg.MaxResponseTime.Milliseconds() <= 0 {
		return models.ErrBadResponseTime
	}

	return nil
}

// Читает конфиг из env
func readEnvConfig() (*Config, error) {
	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		return nil, err
	}

	return &Config{
		Server: ServerConfig{
			Port: port,
		},
		DB: DBConfig{},
	}, nil
}

// Читает конфиг из файла
func readFileConfig(configPath string) (*Config, error) {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, err
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// Получает путь до конфигурационного файла через флаг или env
func fetchConfigPath() (string, error) {
	var path string

	flag.StringVar(&path, "config", "", "path to config file")
	flag.Parse()

	if path == "" {
		err := godotenv.Load()
		if err != nil {
			return "", err
		}

		path = os.Getenv("CONFIG_PATH")

		if path == "" {
			return "", models.ErrConfigPathNotProvided
		}
	}

	return path, nil
}
