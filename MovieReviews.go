// Shelby Simpson
// Movie Reviews Activity
package main

import (
	"fmt"
	"os"
	"strings"
)

func checkWord(word string, positiveWords map[string]int, negativeWords map[string]int) {
	if _, ok := positiveWords[word]; ok {
		positiveWords[word]++
	} else if _, ok = negativeWords[word]; ok {
		negativeWords[word]++
	}
}

func main() {
	var positiveWords = map[string]int{
		"great": 0, "good": 0, "excellent": 0, "acceptable": 0,
		"exceptional": 0, "favorable": 0, "mavelous": 0,
		"positive": 0, "satisfactory": 0, "satisfying": 0,
		"superb": 0, "valuable": 0, "wonderful": 0,
	}

	var negativeWords = map[string]int{
		"bad": 0, "atrocious": 0, "awful": 0, "cheap": 0,
		"crummy": 0, "dreadful": 0, "lousy": 0, "poor": 0,
		"rough": 0, "sad": 0, "unacceptable": 0, "blah": 0,
		"bummer": 0, "diddly": 0, "downer": 0, "garbage": 0,
		"gross": 0,
	}

	var positiveWordCount, negativeWordCount int

	f, err := os.ReadFile("/Users/shelbysimpson/Documents/golang/fileio/review2.txt")
	if err != nil {
		panic(err)
	}
	wordArray := strings.Fields(string(f))
	for _, word := range wordArray {
		checkWord(word, positiveWords, negativeWords)
	}
	for _, count := range positiveWords {
		positiveWordCount += count
	}
	for _, count := range negativeWords {
		negativeWordCount += count
	}
	if positiveWordCount > negativeWordCount {
		fmt.Println("This review is positive.")
	} else if positiveWordCount < negativeWordCount {
		fmt.Println("This review is negative.")
	} else {
		fmt.Println("This review is neutral.")
	}
}
