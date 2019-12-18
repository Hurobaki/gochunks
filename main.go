package main

import (
	"flag"
	"github.com/Hurobaki/gochunks/config"
	"github.com/Hurobaki/gochunks/directories"
	"github.com/Hurobaki/gochunks/errors"
	"github.com/Hurobaki/gochunks/flags"
	"github.com/Hurobaki/gochunks/format"
	"github.com/Hurobaki/gochunks/utils"
	"github.com/Hurobaki/gochunks/zip"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func createChunks(directoryName string, files []string) error {
	var chunkNumber = 0
	var newChunk string

	for index, file := range files {
		if index % *flags.ChunkSize == 0 {
			newChunk = format.Concatenate("", config.SubDirectory, strconv.Itoa(chunkNumber))

			if err := directories.Create(utils.FullPath([]string{*flags.Output, newChunk}...)); err != nil {
				return errors.CreateError("Something went wrong while creating sub directories", err)
			}

			chunkNumber++
		}

		if err := os.Rename(utils.FullPath([]string{directoryName, file}...), utils.FullPath([]string{*flags.Output, newChunk, file}...)); err != nil {
			return errors.CreateError("Something went wrong with Rename() method", err)
		}
	}
	
	return nil
}

func createZip() error {
	files, _ := ioutil.ReadDir(*flags.Output)

	for _, file := range files {
		if file.IsDir() {
			files, err := directories.GetFiles(utils.FullPath(*flags.Output, file.Name()))

			if err != nil {
				log.Fatal(err)
			}

			err = zip.ZipFiles(utils.FullPath(*flags.Output, format.Concatenate("", file.Name(), ".zip")), files, utils.FullPath(*flags.Output, file.Name()))

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return nil
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

	files, err := directories.GetFiles(flags.DirectoryParameterName)

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
		err := createZip()
		if err != nil {
			log.Fatal(err)
		}
	}

	if !*flags.Keep {
		err = directories.CleanDirectory(*flags.Output, directories.IsDirectory)

		if err != nil {
			log.Fatal(errors.CreateError("Something went wrong with CleanDirectory() method", err))
		}
	}
}

