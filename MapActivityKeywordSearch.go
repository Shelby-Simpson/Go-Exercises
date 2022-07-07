// Shelby Simpson
// Map Activity: Keyword Search
package main

import "fmt"

func main() {
	var searchTerm string
	var inMap bool
	produce := map[string]string{
		"apple":      "fruit",
		"banana":     "fruit",
		"kiwi":       "fruit",
		"cherry":     "fruit",
		"strawberry": "fruit",
		"onion":      "vegetable",
		"celery":     "vegetable",
		"spinach":    "vegetable",
		"lettuce":    "vegetable",
		"beet":       "vegetable",
	}
SEARCH:
	fmt.Print("Please enter a search term: ")
	fmt.Scan(&searchTerm)
	if _, ok := produce[searchTerm]; ok {
		inMap = true
		fmt.Println(searchTerm, ":", produce[searchTerm])
	}
	for k, v := range produce {
		if v == searchTerm {
			inMap = true
			fmt.Println(k, ":", v)
		}
	}
	if !inMap {
		fmt.Println("We could not find your search term in the map.  Please try again.")
		goto SEARCH
	}
}
