// Shelby Simpson
// Data Collections Activity 4
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var paragraph string
	fmt.Println("Please enter a paragraph of text:")
	scanner := bufio.NewScanner(os.Stdin)
	if ok := scanner.Scan(); ok {
		paragraph = scanner.Text()
	}
	paragraph = strings.ToLower(paragraph)
	wordArray := strings.Fields(paragraph)
	fmt.Println("Number of words:", len(wordArray))

	var wordMap = make(map[string]int)
	for _, val := range wordArray {
		if _, ok := wordMap[val]; ok {
			wordMap[val]++
		} else {
			wordMap[val] = 1
		}
	}
	fmt.Println("Frequencies:")
	for k, v := range wordMap {
		fmt.Println(k, ":", v)
	}
	// Remove punctuation from paragraph
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	var newParagraph string
	if err == nil {
		newParagraph = reg.ReplaceAllString(paragraph, "")
	}
	newWordArray := strings.Fields(newParagraph)
	fmt.Println("Number of words without punctuation:", len(newWordArray))

	var newWordMap = make(map[string]int)
	for _, val := range newWordArray {
		if _, ok := newWordMap[val]; ok {
			newWordMap[val]++
		} else {
			newWordMap[val] = 1
		}
	}
	fmt.Println("Frequencies without punctuation:")
	for k, v := range newWordMap {
		fmt.Println(k, ":", v)
	}
}
