package storage

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

func SaveEntries(entries []Entry) error {
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

	return os.WriteFile(entriesPath, data, 0600)
}

func LoadEntries() ([]Entry, error) {
	entriesPath, err := getEntriesFilePath()
	if err != nil {
		return nil, err
	}

	var entries []Entry

	if _, err := os.Stat(entriesPath); os.IsNotExist(err) {
		if err := SaveEntries(entries); err != nil {
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

func DeleteEntry(entryToDelete Entry) ([]Entry, error) {
	entries, err := LoadEntries()
	if err != nil {
		return entries, err
	}

	newEntries := make([]Entry, 0, len(entries))
	for _, e := range entries {
		if e != entryToDelete {
			newEntries = append(newEntries, e)
		}
	}

	err = SaveEntries(newEntries)
	if err != nil {
		return entries, err
	}

	return newEntries, nil
}
