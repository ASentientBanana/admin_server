package services

import (
	"os"
	"path"
)

type VersionEntry struct {
	Name string `json:"name" xml:"name"`
	Path string `json:"path" xml:"path"`
	Os   string `json:"os" xml:"os"`
}

func GetDirEntries(root string, alias string) ([]VersionEntry, error) {
	entries, err := os.ReadDir(root)
	if err != nil {
		return []VersionEntry{}, err
	}
	files := []VersionEntry{}
	for _, entry := range entries {
		f := path.Join(alias, entry.Name())
		if entry.IsDir() {
			res, err := GetDirEntries(f, alias)
			if err == nil {
				files = append(files, res...)
			}
		} else {
			files = append(files, VersionEntry{Name: entry.Name(), Path: f, Os: path.Base(root)})
		}
	}
	return files, nil
}
