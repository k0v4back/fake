package faker

//RandomStreet return random street in random language
func RandomStreet() string {
	language := RandomLanguage()
	file := ReadFile(language, STREETS)
	fileLines := Split(file)
	return RandomIntForFiles(fileLines)
}

//RandomStreet return street with specific language
func Street(language string) string {
	file := ReadFile(language, STREETS)
	fileLines := Split(file)
	return RandomIntForFiles(fileLines)
}