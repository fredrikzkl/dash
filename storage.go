package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func getConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "dash"), nil
}

func getEntriesFilePath() (string, error) {
	configDir, err := getConfigPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "entries.json"), nil
}

func saveEntries(entries []entry) error {
	entriesPath, err := getEntriesFilePath()
	if err != nil {
		return err
	}

	// Create directory if doesn't exist
	configDir := filepath.Dir(entriesPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(entries, "", "")
	if err != nil {
		return err
	}

	return os.WriteFile(entriesPath, data, 0644)
}

func loadEntries() ([]entry, error) {
	entriesPath, err := getEntriesFilePath()
	if err != nil {
		return nil, err
	}

	var entries []entry

	if _, err := os.Stat(entriesPath); os.IsNotExist(err) {
		if err := saveEntries(entries); err != nil {
			return entries, err
		}
		return entries, nil
	}

	data, err := os.ReadFile(entriesPath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}
