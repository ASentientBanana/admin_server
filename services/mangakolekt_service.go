package services

import (
	"os"
	"path"
)

type VersionEntry struct {
	Name string
	Path string
}

func GetDirEntries(root string) ([]VersionEntry, error) {
	entries, err := os.ReadDir(root)
	if err != nil {
		return []VersionEntry{}, err
	}
	files := []VersionEntry{}
	for _, entry := range entries {
		f := path.Join(root, entry.Name())
		if entry.IsDir() {
			res, err := GetDirEntries(f)
			if err == nil {
				for _, re := range res {
					files = append(files, re)
				}
			}
		} else {
			files = append(files, VersionEntry{Name: entry.Name(), Path: f})
		}
	}
	return files, nil
}
