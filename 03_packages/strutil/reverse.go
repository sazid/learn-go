package strutil

// Reverse : Reverses the given string and returns a new string
// Any name starting with Capital (Uppercase) letter is exported/public in Go
func Reverse(s string) string {

	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
