package faker

func RandSex() string {
	sex := MALE
	if RandomInt(1, 2) == 1 {
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
	return RandomIntForFiles(fileLines)
}

//LastName return the last name of random sex
func LastName(language string) string {
	sex := RandSex()
	file := ReadFile(language, sex + LASTNAMES)
	fileLines := Split(file)
	return RandomIntForFiles(fileLines)
}

//FirstName return the first name of specific sex
func SexFirstName(language, sex string) string {
	if CheckExistsSex(sex) {
		file := ReadFile(language, sex + FIRSTNAMES)
		fileLines := Split(file)
		return RandomIntForFiles(fileLines)
	}
	return "Wrong sex"
}

//SexLastName return the last name of specific sex
func SexLastName(language, sex string) string {
	if CheckExistsSex(sex) {
		file := ReadFile(language, sex + LASTNAMES)
		fileLines := Split(file)
		return RandomIntForFiles(fileLines)
	}
	return "Wrong sex"
}

func CheckExistsSex(sex string) bool {
	if sex == MALE || sex == FEMALE {
		return true
	}
	return false
}