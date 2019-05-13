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

func RandomIntForFiles(fileLines []string) string {
	return fileLines[RandomInt(0, len(fileLines))]
}

func RandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min) + min
}

func RandomLanguage() string {
	languages := []string{"ru", "en", "de"}
	random := RandomInt(0, len(languages))
	return languages[random]
}