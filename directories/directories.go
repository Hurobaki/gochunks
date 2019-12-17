package directories

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func CreateDirectory(dirName string) error {

	fmt.Println(dirName)

	err := os.Mkdir(dirName, os.ModePerm)

	if err != nil {
		log.Fatal("Au secours !", err)
	}

	return nil
}

func RemoveContents(dirName string) error {
	directory, err := ioutil.ReadDir(dirName)

	if err != nil {
		return err
	}

	for _, d := range directory {
		os.RemoveAll(path.Join([]string{dirName, d.Name()}...))
	}

	return nil
}

func GetDirectoryFiles(dirName string) ([]string, error) {
	fmt.Println(dirName)
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

func ExistsOrCreate(dirName string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(dir)
	if _, err := os.Stat(dir + "/" + dirName); os.IsNotExist(err) {
		fmt.Println("Création du dossier 'prismic_output'")
		os.Mkdir(dirName, os.ModePerm)
	}
	return nil
}

func Create(dirName string) error {
	err := os.Mkdir(dirName, os.ModePerm)

	if err != nil {
		return err
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