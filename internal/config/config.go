package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	ScanInterval int      `json:"scan_interval"`
	MaxThreads   int      `json:"max_threads"`
	IPRange      []string `json:"ip_range"`
}

// LoadConfig считывает JSON-файл и преобразует его в структуру Config
func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err // Если ошибка - возвращаем её
	}
	defer file.Close()

	var cfg Config

	decoder := json.NewDecoder(file) // Создание декодера JSON
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
