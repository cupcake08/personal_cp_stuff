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

func impStuff(result *codeforces.Result, params ...string) {
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

	oldFile, err := os.Open(TEMPLATE)
	if err != nil {
		log.Fatal(err)
	}
	defer oldFile.Close()

	bytes, err := io.ReadAll(oldFile)
	if err != nil {
		log.Fatal(err.Error())
	}
	if len(params) > 2 {
		folderName = params[2]
	}
	currentTime := time.Now()

	for i := 0; i < filesCount; i++ {
		fileName := fmt.Sprintf("%s.cpp", string(rune(65+i)))

		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}

		if result != nil {
			fileName = result.Problems[i].Name
		} else {
			fileName = "NA"
		}
		s := fmt.Sprintf(`/**
*
* Author: Ankit Bhankharia
* Created At: %d-%d-%d %d:%d:%d
* Contest: %s
* Problem: %s
*/
`,
			currentTime.Year(),
			currentTime.Month(),
			currentTime.Day(),
			currentTime.Hour(),
			currentTime.Minute(),
			currentTime.Second(),
			folderName,
			fileName,
		)
		idx := len(s)
		idx++

		file.WriteString(s + "\n")

		_, err = file.WriteAt(bytes, int64(idx))

		if err != nil {
			log.Fatal(err.Error())
		}

		err = file.Sync()

		if err != nil {
			log.Fatal(err)
		}

		file.Close()
	}
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
		result, err := codeforces.CodeforcesStandings(contestId)
		if err != nil {
			log.Fatal(err)
		}
		impStuff(result, "codeforces_contest_"+contestId, fmt.Sprint(len(result.Problems)), result.Contest.Name)
	} else {
		if len(os.Args) < 4 {
			log.Fatal("Not enough arguments. Aborting.")
			os.Exit(1)
		}
		impStuff(nil, os.Args[2], os.Args[3])
	}
}
