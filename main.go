package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// Delete everything inside prismic_output directory

const DirectoryResultName = "prismic_output"
const SubDirectoryName = "chunk"
const ChunkSize = 200

func createDirectory(dirName string) error {

	fmt.Println(dirName)

	err := os.Mkdir(dirName, os.ModePerm)

	if err != nil {
		log.Fatal("Au secours !", err)
	}

	return nil
}

func getDirectoryFiles(dirName string) ([]string, error) {
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

func directoryExistsOrCreate(dirName string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(dir)
	if _, err := os.Stat(dir + "/" + dirName); os.IsNotExist(err) {
		fmt.Println("Création du dossier 'prismic_output'")
		os.Mkdir(DirectoryResultName, os.ModePerm)
	}
	return nil
}

func createChunks(directoryName string, files []string) error {
	var chunkNumber = 0
	var newChunk string

	for index, file := range files {
		if index % ChunkSize == 0 {
			newChunk = fmt.Sprintf("%s%d", SubDirectoryName,chunkNumber)
			err := createDirectory(fmt.Sprintf("%s/%s", DirectoryResultName, newChunk))

			if err != nil {
				fmt.Println("Error lors de la création des sous dossiers", err)
			}
			chunkNumber++
		}

		err := os.Rename(fmt.Sprintf("%s/%s", directoryName, file), fmt.Sprintf("./%s/%s/%s",DirectoryResultName, newChunk, file))

		if err != nil {
			return errors.New(fmt.Sprintf("Something went wrong with Rename() method : %s", err.Error()))
		}
	}
	
	return nil
}


func main() {
	directoryName := os.Args[1]

	fmt.Println(directoryName)

	files, err := getDirectoryFiles(directoryName)

	if err != nil {
		log.Fatal("Error getting files", err)
	}

	fmt.Println(len(files))

	err = directoryExistsOrCreate(DirectoryResultName)

	if err != nil {
		log.Fatal("Directory exists or create", err)
	}

	err = createChunks(directoryName, files)
	
	if err != nil {
		log.Fatal("", err)
	}
	
}

