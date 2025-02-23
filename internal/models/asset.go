package models

import (
    "encoding/json"
    "os"
    "path/filepath"
)

type Asset struct {
    ID          string   `json:"id"`
    Name        string   `json:"name"`
    IP          string   `json:"ip"`
    Ports       []int    `json:"ports"`
    Vulnerabilities []string `json:"vulnerabilities"`

    URL         string `json:"url"`
    Type        string `json:"type"`  // web, api, mobile, etc.
    Description string `json:"description"`
    Status      string `json:"status"` // active, archived, testing
    Notes       string `json:"notes"`
}

func LoadAssets(assetPath, companyName string) ([]*Asset, error) {
    jsonPath := filepath.Join(assetPath, companyName + ".json")
    data, err := os.ReadFile(jsonPath)
    if err != nil {
        return nil, err
    }

    var assets []*Asset
    if err := json.Unmarshal(data, &assets); err != nil {
        return nil, err
    }

    return assets, nil
}

func NewAsset() *Asset {
    return &Asset{
        Status: "active",
    }
}

