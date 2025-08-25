package main

import "strings"

type entry struct {
	name string
	path string
}

func newEntry(path string) entry {
	split := strings.Split(path, "/")
	name := split[len(split)-1]
	return entry{
		name: name,
		path: path,
	}
}

// Testing
func getMockEntries() []entry {
	warlockEntry := newEntry("/users/john/warlock")
	smashedEntry := newEntry("/users/john/smashed")
	return []entry{
		warlockEntry,
		smashedEntry,
	}
}
