package faker

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

func ReadFile(language, filename string) []byte {
	readFile, err := ioutil.ReadFile(DIR + language + "/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	return readFile
}

func Split(file []byte) []string {
	fileLines := strings.Split(string(file), "\n")
	rand.Seed(time.Now().UnixNano())
	return fileLines
}

func RandInt(fileLines []string) string {
	return fileLines[rand.Intn(len(fileLines))]
}

func CheckExistsSex(sex string) bool {
	if sex == MALE || sex == FEMALE {
		return true
	}
	return false
}