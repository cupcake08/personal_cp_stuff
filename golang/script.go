package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// path to template file
const TEMPLATE string = "/home/ankit/CP/template.cpp"

func impStuff() {
	folderName := os.Args[1]
	err := os.Mkdir(folderName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Chdir(folderName)
	if err != nil {
		log.Fatal(err)
	}

	filesCount, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	oldFile, err := os.Open(TEMPLATE)
	if err != nil {
		log.Fatal(err)
	}
	defer oldFile.Close()

	for i := 0; i < filesCount; i++ {
		fileName := fmt.Sprintf("%s.cpp", string(rune(65+i)))
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(file, oldFile)
		if err != nil {
			log.Fatal(err)
		}

		err = file.Sync()
		if err != nil {
			log.Fatal(err)
		}

		file.Close()
	}

	log.Println("Files Created Sucsessfully.")
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Not enough argument1. Aborting.")
		os.Exit(1)
	}
	impStuff()
}
