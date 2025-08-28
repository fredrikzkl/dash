package storage

import "strings"

type Entry struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Command string `json:"command"`
}

func NewEntry(path string) Entry {
	split := strings.Split(path, "/")
	name := split[len(split)-1]
	return Entry{
		Name: name,
		Path: path,
	}
}

// Testing
func getMockEntries() []Entry {
	warlockEntry := NewEntry("/Users/fredrik/vippsnummer")
	smashedEntry := NewEntry("/Users/fredrik/shopping-basket")
	return []Entry{
		warlockEntry,
		smashedEntry,
	}
}
