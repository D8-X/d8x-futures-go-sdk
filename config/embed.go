package config

import (
	"embed"
	"encoding/json"
	"os"
)

//go:embed embedded/*
var EmbededConfigs embed.FS

func GetSymbolList() (map[string]string, error) {
	jsonData, err := os.ReadFile("embedded/symbolList.json")
	if err != nil {
		return nil, err
	}
	// Define a map to store the data
	data := make(map[string]string)

	// Unmarshal the JSON data into the map
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
