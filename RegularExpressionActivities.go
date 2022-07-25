// Shelby Simpson
// Regular Expression Activities
package main

import (
	"fmt"
	"regexp"
)

// Activity 1
func lettersAndNumbers(s string) (bool, error) {
	return regexp.MatchString(`^[a-zA-Z0-9]+$`, s)
}

// Activity 2
func iAndn(s string) []string {
	re := regexp.MustCompile(`in*`)
	return re.FindAllString(s, -1)
}

// Activity 3
func iAndn1(s string) []string {
	re := regexp.MustCompile(`in+`)
	return re.FindAllString(s, -1)
}

// Activity 4
func iAndn2(s string) []string {
	re := regexp.MustCompile(`i(nn|n)`)
	return re.FindAllString(s, -1)
}

// Activity 5
func iAndn3(s string) []string {
	re := regexp.MustCompile(`in{3}`)
	return re.FindAllString(s, -1)
}

func lowercaseWithUnderscore(s string) []string {
	re := regexp.MustCompile(`[a-z]+_[a-z]+`)
	return re.FindAllString(s, -1)
}

func startaEndb(s string) []string {
	re := regexp.MustCompile(`a.*b`)
	return re.FindAllString(s, -1)
}

func main() {
	str1 := "Letters1numbers"
	fmt.Println(lettersAndNumbers(str1))
	str2 := "inniniasdasinnnn"
	fmt.Println(iAndn(str2))
	fmt.Println(iAndn1(str2))
	fmt.Println(iAndn2(str2))
	fmt.Println(iAndn3(str2))
	str3 := "A_a_asd_S_a_b"
	fmt.Println(lowercaseWithUnderscore(str3))
	str4 := "sdacbasdva"
	fmt.Println(startaEndb(str4))
}
