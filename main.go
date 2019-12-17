package main

import (
	"flag"
	"fmt"
	"github.com/Hurobaki/gochunks/config"
	"github.com/Hurobaki/gochunks/directories"
	"github.com/Hurobaki/gochunks/errors"
	"github.com/Hurobaki/gochunks/flags"
	"github.com/Hurobaki/gochunks/zip"
	"io/ioutil"
	"log"
	"os"
)

func createChunks(directoryName string, files []string) error {
	var chunkNumber = 0
	var newChunk string

	for index, file := range files {
		if index % *flags.ChunkSize == 0 {
			newChunk = fmt.Sprintf("%s%d", config.SubDirectory,chunkNumber)
			err := directories.Create(fmt.Sprintf("%s/%s", config.DirectoryName, newChunk))

			if err != nil {
				fmt.Println("Error lors de la création des sous dossiers", err)
			}
			chunkNumber++
		}

		err := os.Rename(fmt.Sprintf("%s/%s", directoryName, file), fmt.Sprintf("./%s/%s/%s", config.DirectoryName, newChunk, file))

		if err != nil {
			return errors.CreateError("Something went wrong with Rename() method", err)
		}
	}
	
	return nil
}

func createZip() {
	files, _ := ioutil.ReadDir(config.DirectoryName)

	for _, file := range files {
		if file.IsDir() {
			files, err := directories.GetDirectoryFiles(fmt.Sprintf("%s/%s", config.DirectoryName, file.Name()))

			if err != nil {
				log.Fatal(err)
			}

			err = zip.ZipFiles(fmt.Sprintf("%s/%s.zip", config.DirectoryName, file.Name()), files, fmt.Sprintf("%s/%s/", config.DirectoryName, file.Name()))

			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func createDirectoryOutput() error {
	dirExists, err := directories.Exists(config.DirectoryName)

	if err != nil {
		return errors.CreateError("Something went wrong with Exists() method", err)
	}

	if !dirExists {
		err := directories.Create(config.DirectoryName)

		if err != nil {
			return errors.CreateError("Something went wrong with Create() method", err)
		}
	} else {
		err := directories.RemoveContents(config.DirectoryName)

		if err != nil {
			return errors.CreateError("Something went wrong with RemoveContents() method", err)
		}
	}

	return nil
}


func main() {
	flags.Zip = flag.Bool("zip", config.Zip, "create zip files")
	flags.ChunkSize = flag.Int("size", config.ChunkSize, "chunks size")
	flag.Parse()

	if len(flag.Args()) > 0 {
		flags.DirectoryParameterName = flag.Args()[0]
	}

	files, err := directories.GetDirectoryFiles(flags.DirectoryParameterName)

	if err != nil {
		log.Fatal(err)
	}

	err = createDirectoryOutput()

	if err != nil {
		log.Fatal(err)
	}

	err = createChunks(flags.DirectoryParameterName, files)

	if err != nil {
		log.Fatal(err)
	}

	if *flags.Zip {
		createZip()
	}
}

