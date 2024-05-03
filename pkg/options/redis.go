package options

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Redis struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       int    `yaml:"db"`
}

type RedisConfig struct {
	Redis `yaml:"redis"`
}

func LoadRedisConfig(filepath string) (*RedisConfig, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var cfg RedisConfig
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
