package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

const (
	cfgDirname  = "config"
	cfgFilename = "config.yaml"
	envFilename = ".env"

	pgConnStrEnv = "PG_CONN_STR"
)

// Config - структура содержащая параметры настройки требуемые для запуска программы.
type Config struct {
	Pg struct {
		ConnStr string
	}
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

// New - конструктор структуры Config.
// На вход дается путь к рабочей директории
func New(wd string) (*Config, error) {
	cfgDirpath := filepath.Join(wd, cfgDirname)
	f, err := os.ReadFile(filepath.Join(cfgDirpath, cfgFilename))
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	if err := yaml.Unmarshal(f, cfg); err != nil {
		return nil, err
	}
	if err := godotenv.Load(filepath.Join(wd, envFilename)); err != nil {
		return nil, err
	}
	pgConnStr := os.Getenv(pgConnStrEnv)
	if pgConnStr == "" {
		return nil, fmt.Errorf("environment variable %s not set", pgConnStrEnv)
	}
	cfg.Pg.ConnStr = pgConnStr
	return cfg, nil
}
