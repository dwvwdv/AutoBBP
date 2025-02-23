package config

import (
    "os"
    "path/filepath"
    "gopkg.in/yaml.v2"
)

type Config struct {
    AssetPath string `yaml:"asset_path"`
}

func LoadConfig() (*Config, error) {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        return nil, err
    }

    configPath := filepath.Join(homeDir, "config", "autobbp.yaml")
    data, err := os.ReadFile(configPath)
    if err != nil {
        return nil, err
    }

    var config Config
    if err := yaml.Unmarshal(data, &config); err != nil {
        return nil, err
    }

    return &config, nil
}

