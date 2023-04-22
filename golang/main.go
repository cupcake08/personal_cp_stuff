package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	codeforcescrawler "github.com/cupcake08/codeforces_crawler"
	codeforces "github.com/cupcake08/cp_scripts/codeforces"
)

// path to template file
const TEMPLATE string = "/home/ankit/CP/template.cpp"

func impStuff(result *codeforces.Result, params ...string) {
	var folderName string
	var filesCount int

	if result != nil {
		folderName = "codeforces_contest_" + params[0]
		filesCount = len(result.Problems)
	} else {
		folderName = params[0]
		fc, err := strconv.Atoi(params[1])
		if err != nil {
			log.Fatal(err.Error())
			os.Exit(1)
		}
		filesCount = fc
	}

	err := os.Mkdir(folderName, os.ModePerm)

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

	currentTime := time.Now()
	var crawler *codeforcescrawler.Contest
	if result != nil {
		f, err := strconv.Atoi(params[0])
		if err != nil {
			log.Fatal(err.Error())
		}
		crawler = codeforcescrawler.NewContest(f)
	}

	for i := 0; i < filesCount; i++ {
		x := string(rune(65 + i))
		var fileName string
		if crawler != nil {
			err := os.Mkdir(x, 0777)
			if err != nil {
				log.Fatal(err.Error())
			}

			err = os.Chdir(x)
			if err != nil {
				log.Fatal(err.Error())
			}

			_, err = os.Getwd()
			if err != nil {
				log.Fatal(err.Error())
			}
			crawler.GetTestCases(result.Problems[i].Index)
		}

		fileName = fmt.Sprintf("%s.cpp", x)
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		if result != nil {
			folderName = result.Contest.Name
			fileName = result.Problems[i].Name
		} else {
			fileName = "NA"
		}

		var s strings.Builder
		s.WriteString("/**\n")
		s.WriteString("*\n")
		s.WriteString("*    Author      : Ankit Bhankharia\n")
		s.WriteString(fmt.Sprintf(
			"*    Created At  : %d-%d-%d %d:%d:%d\n",
			currentTime.Year(),
			currentTime.Month(),
			currentTime.Day(),
			currentTime.Hour(),
			currentTime.Minute(),
			currentTime.Second(),
		),
		)
		s.WriteString(fmt.Sprintf("*    Contest     : %s\n", folderName))
		s.WriteString(fmt.Sprintf("*    Problem     : %s\n", fileName))
		s.WriteString("*\n")
		s.WriteString("**/\n")

		idx := len(s.String())
		idx++

		file.WriteString(s.String() + "\n")

		_, err = file.WriteAt(bytes, int64(idx))

		if err != nil {
			log.Fatal(err.Error())
		}

		err = file.Sync()

		if err != nil {
			log.Fatal(err)
		}
		if crawler != nil {
			os.Chdir("..")
		}
	}
}

func main() {
	platform := flag.String("platform", "Codeforces", "Select From Which Platform You Need to use it for.")
	contestId := flag.String("contestId", "599", "Codeforces Contest Id for which you want to work")
	filesCount := flag.String("count", "0", "Number of files You want to create.")
	folderName := flag.String("name", "test", "Name Of Your Folder")
	flag.Parse()
	if *platform == "Codeforces" {
		if len(os.Args) < 3 {
			log.Fatal("Not enough arguments. Aborting.")
			os.Exit(1)
		}
		result, err := codeforces.CodeforcesStandings(*contestId)
		if err != nil {
			log.Fatal(err)
		}
		impStuff(result, *contestId)
	} else {
		if len(os.Args) < 4 {
			log.Fatal("Not enough arguments. Aborting.")
			os.Exit(1)
		}
		impStuff(nil, *folderName, *filesCount)
	}
}
