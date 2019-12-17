package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/Hurobaki/gochunks/directories"
	"github.com/Hurobaki/gochunks/flags"
	"github.com/Hurobaki/gochunks/zip"
	"io/ioutil"
	"log"
	"os"
)

const DirectoryResultName = "prismic_output"
const SubDirectoryName = "chunk"
const ChunkSize = 200

func createChunks(directoryName string, files []string) error {
	var chunkNumber = 0
	var newChunk string

	for index, file := range files {
		if index % ChunkSize == 0 {
			newChunk = fmt.Sprintf("%s%d", SubDirectoryName,chunkNumber)
			err := directories.Create(fmt.Sprintf("%s/%s", DirectoryResultName, newChunk))

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

func createZip() {
	files, _ := ioutil.ReadDir(DirectoryResultName)

	for _, file := range files {
		if file.IsDir() {
			fmt.Println("c'est un dossier !!!")
			files, err := directories.GetDirectoryFiles(fmt.Sprintf("%s/%s", DirectoryResultName, file.Name()))

			if err != nil {
				log.Fatal(err)
			}

			err = zip.ZipFiles(fmt.Sprintf("%s/%s.zip", DirectoryResultName, file.Name()), files, fmt.Sprintf("%s/%s/", DirectoryResultName, file.Name()))

			if err != nil {
				log.Fatal(err)
			}
		}
	}
}


func main() {
	flags.Zip = flag.Bool("zip", false, "create zip files")
	flag.Parse()

	var directoryName string

	if len(flag.Args()) > 0 {
		directoryName = flag.Args()[0]
	}

	files, err := directories.GetDirectoryFiles(directoryName)

	if err != nil {
		log.Fatal("Error getting files", err)
	}

	fmt.Println(len(files))

	dirExists, err := directories.Exists(DirectoryResultName)

	if err != nil {
		log.Fatal("problème vérification existence dossier", err)
	}

	if !dirExists {
		err := directories.Create(DirectoryResultName)

		if err != nil {
			log.Fatal("problème création dossier", err)
		}
	} else {
		err := directories.RemoveContents(DirectoryResultName)

		if err != nil {
			log.Fatal("pwet", err)
		}
	}

	err = createChunks(directoryName, files)

	if *flags.Zip {
		createZip()
	}


	if err != nil {
		log.Fatal("", err)
	}

}

