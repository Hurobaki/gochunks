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
			err := directories.Create(fmt.Sprintf("%s/%s", *flags.Output, newChunk))

			if err != nil {
				fmt.Println("Error lors de la création des sous dossiers", err)
			}
			chunkNumber++
		}

		err := os.Rename(fmt.Sprintf("%s/%s", directoryName, file), fmt.Sprintf("./%s/%s/%s", *flags.Output, newChunk, file))

		if err != nil {
			return errors.CreateError("Something went wrong with Rename() method", err)
		}
	}
	
	return nil
}

func createZip() {
	files, _ := ioutil.ReadDir(*flags.Output)

	for _, file := range files {
		if file.IsDir() {
			files, err := directories.GetDirectoryFiles(fmt.Sprintf("%s/%s", *flags.Output, file.Name()))

			if err != nil {
				log.Fatal(err)
			}

			err = zip.ZipFiles(fmt.Sprintf("%s/%s.zip", *flags.Output, file.Name()), files, fmt.Sprintf("%s/%s/", *flags.Output, file.Name()))

			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func createDirectoryOutput() error {
	dirExists, err := directories.Exists(*flags.Output)

	if err != nil {
		return errors.CreateError("Something went wrong with Exists() method", err)
	}

	if !dirExists {
		err := directories.Create(*flags.Output)

		if err != nil {
			return errors.CreateError("Something went wrong with Create() method", err)
		}
	} else {
		err := directories.RemoveContents(*flags.Output)

		if err != nil {
			return errors.CreateError("Something went wrong with RemoveContents() method", err)
		}
	}

	return nil
}


func main() {
	flags.Zip = flag.Bool("zip", config.Zip, "create zip files")
	flags.ChunkSize = flag.Int("size", config.ChunkSize, "chunks size")
	flags.Keep = flag.Bool("keep", config.Keep, "keep output directories")
	flags.Output = flag.String("o", config.Output, "output directory")
	flag.Parse()

	if len(flag.Args()) > 0 {
		inputDirectoryName := flag.Args()[0]
		isDir, err := directories.IsDirectory(inputDirectoryName)

		if err != nil {
			log.Fatal(errors.CreateError("The required directory doesn't exist", err))
		}

		if isDir {
			flags.DirectoryParameterName = flag.Args()[0]
		} else {
			log.Fatal(errors.CreateError("Please enter a valid directory", nil))
		}
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

	if !*flags.Keep {
		directories.CleanDirectory(*flags.Output, directories.IsDirectory)
	}
}

