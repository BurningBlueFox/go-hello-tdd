package iteration

// Repeat repeats the character 5 times and returns it in a new string
func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += character
	}
	return repeated
}
