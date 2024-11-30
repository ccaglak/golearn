package iteration

//import "strings"

func Repeat(char string, repeatCount int) (repeat string) {
	for i := 0; i < repeatCount; i++ {
		repeat += char
	}
	//repeat = strings.Repeat(char, repeatCount)
	return
}
