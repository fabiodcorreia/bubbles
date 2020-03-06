package server

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// ResourceFile holds the relative and absolute path of the file
type ResourceFile struct {
	RelPath string
	AbsPath string
}

// GetResourceContent receives a file path and return the content in a []byte
func GetResourceContent(path string) ([]byte, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// SearchFiles lookup for all the files on the rootDir and sub directories
func SearchFiles(rootDir string) ([]ResourceFile, error) {
	var resourceFiles []ResourceFile
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && skipDir(info) {
			return filepath.SkipDir
		}
		if !info.IsDir() && !skipFile(info) {
			rel, _ := filepath.Rel(rootDir, path)
			abs, _ := filepath.Abs(path)
			resourceFiles = append(resourceFiles, ResourceFile{
				AbsPath: abs,
				RelPath: "/" + rel, //? Better way to do this?
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return resourceFiles, nil
}

// skipDir checks if the file is a directory and if its hidden folder
// if the name starts with . and it's not the . or .. then it's hidden
func skipDir(info os.FileInfo) bool {
	return skipFile(info) && info.Name() != "." && info.Name() != ".."
}

// skipFile checks if the file is hidden
// if the name starts with . it's hidden
func skipFile(info os.FileInfo) bool {
	return info.Name()[0] == 46
}
