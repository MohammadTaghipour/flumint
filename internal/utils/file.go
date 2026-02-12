package utils

import (
	"fmt"
	"io"
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

	err := filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			dirs = append(dirs, p)
		}
		return nil
	})
	if err != nil {
		return err
	}

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

func IsFile(p string) bool {
	return !IsDirectory(p)
}

func IsDirectory(p string) bool {
	info, err := os.Stat(p)
	return info.IsDir() && err == nil
}

func FileExists(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}

func DirectoryExists(p string) bool {
	info, err := os.Stat(p)
	if os.IsNotExist(err) || !info.IsDir() {
		return false
	}
	return err == nil
}

func CopyDirectory(src, dest string) error {
	if !DirectoryExists(src) {
		return fmt.Errorf("source directory does not exist: %s", src)
	}

	return filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(dest, relPath)

		if info.IsDir() {
			// TODO: check for FileModes in entire project
			return os.MkdirAll(destPath, os.ModePerm)
		}

		return CopyFile(path, destPath)
	})
}

func CopyFile(src, dest string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	if err := os.MkdirAll(filepath.Dir(dest), os.ModePerm); err != nil {
		return err
	}

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}
