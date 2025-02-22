package models

type Asset struct {
    URL         string `json:"url"`
    Type        string `json:"type"`  // web, api, mobile, etc.
    Description string `json:"description"`
    Status      string `json:"status"` // active, archived, testing
    Notes       string `json:"notes"`
}

func NewAsset() *Asset {
    return &Asset{
        Status: "active",
    }
}

