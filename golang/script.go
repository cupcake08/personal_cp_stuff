package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	codeforces "github.com/cupcake08/cp_scripts/codeforces"
)

// path to template file
const TEMPLATE string = "/home/ankit/CP/template.cpp"

func impStuff(params ...string) {
	var folderName string
	var filesCount int
	folderName = params[0]
	filesCount, err := strconv.Atoi(params[1])
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	err = os.Mkdir(folderName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Chdir(folderName)
	if err != nil {
		log.Fatal(err)
	}
	// filesCount, err := strconv.Atoi(os.Args[2])

	if err != nil {
		log.Fatal(err)
	}

	oldFile, err := os.Open(TEMPLATE)
	if err != nil {
		log.Fatal(err)
	}
	defer oldFile.Close()

	bytes, err := io.ReadAll(oldFile)
	if err != nil {
		log.Fatal(err.Error())
	}
	currentTime := time.Now()

	s := fmt.Sprintf(`/**
*
* author: Ankit Bhankharia
* created at: %d-%d-%d %d:%d:%d
*/
`,
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Minute(),
		currentTime.Second(),
	)
	idx := len(s)
	idx++

	for i := 0; i < filesCount; i++ {
		fileName := fmt.Sprintf("%s.cpp", string(rune(65+i)))
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}

		file.WriteString(s + "\n")
		n, err := file.WriteAt(bytes, int64(idx))
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("%d bytes written\n", n)

		err = file.Sync()
		if err != nil {
			log.Fatal(err)
		}

		file.Close()
	}

	log.Println("Files Created Sucsessfully.")
}

func main() {
	platform := os.Args[1]

	os.Chdir("..")
	if platform == "codeforces" {
		if len(os.Args) < 3 {
			log.Fatal("Not enough arguments. Aborting.")
			os.Exit(1)
		}

		contestId := os.Args[2]
		_, problems, err := codeforces.CodeforcesStandings(contestId)
		if err != nil {
			log.Fatal(err)
		}
		impStuff("codeforces_contest_"+contestId, fmt.Sprint(problems))
	} else {
		if len(os.Args) < 4 {
			log.Fatal("Not enough arguments. Aborting.")
			os.Exit(1)
		}
		impStuff(os.Args[2], os.Args[3])
	}
}
