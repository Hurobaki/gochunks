package directories

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func RemoveContents(dirName string) error {
	directory, err := ioutil.ReadDir(dirName)

	if err != nil {
		return errors.New(fmt.Sprintf("Something went wrong with directory RemoveContents() method : %s", err.Error()))
	}

	for _, d := range directory {
		os.RemoveAll(path.Join([]string{dirName, d.Name()}...))
	}

	return nil
}

// refacto GetDirectoryFiles => GetFiles
func GetDirectoryFiles(dirName string) ([]string, error) {
	var files []string
	f, err := os.Open(dirName)

	if err != nil {
		return files, err
	}

	fileInfo, err := f.Readdir(-1)
	f.Close()

	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}

	return files, nil
}

func Create(dirName string) error {
	err := os.Mkdir(dirName, os.ModePerm)

	if err != nil {
		return errors.New(fmt.Sprintf("Something went wrong with directory Create() method : %s", err.Error()))
	}

	return nil
}

func Exists(dirName string) (bool, error) {
	dir, err := os.Getwd()
	if err != nil {
		return true, err
	}

	if _, err := os.Stat(dir + "/" + dirName); os.IsNotExist(err) {
		return false, nil
	}

	return true, err
}

func IsDirectory(dirName string) (bool, error) {
	return true, nil
}