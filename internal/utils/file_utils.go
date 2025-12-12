package utils

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func ReplaceInFile(path, old, replacement string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	updated := strings.ReplaceAll(string(data), old, replacement)
	return os.WriteFile(path, []byte(updated), os.ModePerm)
}

func ReplaceInFileRegex(path, pattern, replacement string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	reg := regexp.MustCompile(pattern)
	updated := reg.ReplaceAllString(string(data), replacement)
	return os.WriteFile(path, []byte(updated), os.ModePerm)
}

func DeleteEmptyDirs(root string) error {
	var dirs []string

	filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			dirs = append(dirs, p)
		}
		return nil
	})

	// Reverse-delete so nested dirs get removed first
	for i := len(dirs) - 1; i >= 0; i-- {
		dir := dirs[i]
		entries, _ := os.ReadDir(dir)
		if len(entries) == 0 {
			_ = os.Remove(dir)
		}
	}

	return nil
}

func FileExists(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}
