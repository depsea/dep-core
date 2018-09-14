package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// DBConfig -
type DBConfig struct {
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// RedisConfig -
type RedisConfig struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}

// ServerConfig -
type ServerConfig struct {
	Addr string `json:"addr"`
}

// Config -
type Config struct {
	DB     DBConfig
	Redis  RedisConfig
	Server ServerConfig
}

// GetConfig -
func GetConfig() (*Config, error) {
	data, err := ioutil.ReadFile("./config/config.json")

	config := Config{}

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// GetServerAddr -
func GetServerAddr() (string, error) {
	config, err := GetConfig()

	if err != nil {
		return "", err
	}

	return config.Server.Addr, nil
}

// GetDBInfo -
func GetDBInfo() (string, error) {
	config, err := GetConfig()

	if err != nil {
		return "", err
	}

	db := config.DB

	pgInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		db.Host, db.Port, db.User, db.Password, db.Name, "disable",
	)

	return pgInfo, nil
}
