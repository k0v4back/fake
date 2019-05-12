package faker

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

const (
	FIRSTNAMES 	= "FirstNames"
	LASTNAMES 	= "LastNames"
	DIR 		= "./data/"
	MALE 		= "male"
	FEMALE 		= "female"
)

func RandSex() string {
	sex := MALE
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 1 {
		sex = FEMALE
	}
	return sex
}

//Name return the first name and last name of the random sex
func Name(language string) string {
	sex := RandSex()
	return SexFirstName(language, sex) + " " + SexFirstName(language, sex)
}

//SexName return the first name and last name with specific sex
func SexName(language string, sex string) string {
	switch sex {
	case MALE:
		return SexLastName(language, MALE) + " " + SexLastName(language, MALE)
	case FEMALE:
		return SexLastName(language, FEMALE) + " " + SexLastName(language, FEMALE)
	default:
		return "Wrong sex"
	}
}

//FirstName return the first name of random sex
func FirstName(language string) string {
	sex := RandSex()
	file := ReadFile(language, sex + FIRSTNAMES)
	fileLines := Split(file)
	return RandInt(fileLines)
}

//LastName return the last name of random sex
func LastName(language string) string {
	sex := RandSex()
	file := ReadFile(language, sex + LASTNAMES)
	fileLines := Split(file)
	return RandInt(fileLines)
}

//FirstName return the first name of specific sex
func SexFirstName(language, sex string) string {
	if CheckExistsSex(sex) {
		file := ReadFile(language, sex + FIRSTNAMES)
		fileLines := Split(file)
		return RandInt(fileLines)
	}
	return "Wrong sex"
}

//SexLastName return the last name of specific sex
func SexLastName(language, sex string) string {
	if CheckExistsSex(sex) {
		file := ReadFile(language, sex + LASTNAMES)
		fileLines := Split(file)
		return RandInt(fileLines)
	}
	return "Wrong sex"
}

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