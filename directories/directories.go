package directories

import (
	"fmt"
	"github.com/Hurobaki/gochunks/errors"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func RemoveContents(dirName string) error {
	directory, err := ioutil.ReadDir(dirName)

	if err != nil {
		return errors.CreateError("Something went wrong with directory RemoveContents() method", err)
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
		return errors.CreateError("Something went wrong with directory Create() method ", err)
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
	file, err := os.Stat(dirName)

	if err != nil {
		return false, err
	}

	isDir := file.IsDir()

	if !isDir {
		return false, nil
	}

	return true, nil
}

func RemoveDirectory(dirName string, removeAll bool) error {
	var err error = nil

	if removeAll {
		err = os.RemoveAll(dirName)
	} else {
		err = os.Remove(dirName)
	}

	if err != nil {
		return err
	}

	return nil
}

func CleanDirectory(dirName string, predicate interface{}) {
	files, err := GetDirectoryFiles(dirName)

	if err != nil {
		log.Fatal(errors.CreateError("Pwet", nil))
	}

	for _, file := range files {
		filePath := fmt.Sprintf("%s/%s", dirName, file)
		var isDirectory = false

		switch v := predicate.(type) {
		case bool:
			isDirectory = predicate.(bool)
		case func(string) (bool, error):
			isDir, err := v(filePath)

			if err != nil {
				log.Fatal("pwet")
			}

			isDirectory = isDir
		}

		if isDirectory {
			err := RemoveDirectory(filePath, true)

			if err != nil {
				fmt.Println(errors.CreateError("Something went wrong with RemoveDirectory() method ", err))
			}
		}
	}
}