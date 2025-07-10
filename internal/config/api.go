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

type LoggerConfig struct {
	AddSource bool `yaml:"source"`
}

type Config struct {
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"db"`
	Logger LoggerConfig `yaml:"config"`
}

func New() (cfg *Config, err error) {
	configPath, err := fetchConfigPath()
	if err != nil && !errors.Is(err, models.ErrConfigPathNotProvided) {
		return nil, fmt.Errorf("read config error: %v", err)
	}

	cfg, err = readFileConfig(configPath)
	if err != nil {
		return nil, err
	}

	cfg = readEnvConfig(cfg)

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
func readEnvConfig(cfg *Config) *Config {
	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		return nil
	}

	if port != 0 {
		cfg.Server.Port = port
	}

	address := os.Getenv("SERVER_ADDRESS")
	if address != "" {
		cfg.Server.Addr = address
	}

	maxResponseTime, err := time.ParseDuration(os.Getenv("SERVER_RESPONSE_TIME"))
	if err != nil {
		return nil
	}

	if maxResponseTime != 0 {
		cfg.Server.MaxResponseTime = maxResponseTime
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword != "" {
		cfg.DB.Password = dbPassword
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser != "" {
		cfg.DB.User = dbUser
	}

	dbName := os.Getenv("DB_NAME")
	if dbName != "" {
		cfg.DB.Name = dbName
	}

	source, err := strconv.ParseBool(os.Getenv("LOGGER_SOURCE"))
	if err != nil {
		return nil
	}

	cfg.Logger.AddSource = source

	return cfg
}

// Читает конфиг из файла
func readFileConfig(configPath string) (*Config, error) {
	var cfg Config
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return &cfg, err
	}

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return &cfg, err
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

func GetConnStr(user, password, name string) string {
	return fmt.Sprintf("postgres://%s:%s@db:5432/%s?sslmode=disable", user, password, name)
}
