package main

import "strings"

type entry struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Command string `json:"command"`
}

func newEntry(path string) entry {
	split := strings.Split(path, "/")
	name := split[len(split)-1]
	return entry{
		Name: name,
		Path: path,
	}
}

// Testing
func getMockEntries() []entry {
	warlockEntry := newEntry("/Users/fredrik/vippsnummer")
	smashedEntry := newEntry("/Users/fredrik/shopping-basket")
	return []entry{
		warlockEntry,
		smashedEntry,
	}
}
