package conf

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	HTTPServer httpServer `yaml:"http_server"`
	Postgresql postgresql `yaml:"postgresql"`
}

type httpServer struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

func (h *httpServer) GetListenAddress() string {
	return fmt.Sprintf("%s:%d", h.Address, h.Port)
}

type postgresql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	SSLMode  string `yaml:"ssl_mode"`
}

func (p *postgresql) GetDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", p.User, p.Password, p.Host, p.Port, p.Database, p.SSLMode)
}

func Read(confPath string) (*Config, error) {
	confFile, err := os.Open(confPath)
	if err != nil {
		return nil, err
	}
	defer confFile.Close()

	decoder := yaml.NewDecoder(confFile)
	var config Config

	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
