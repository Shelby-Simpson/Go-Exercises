// Shelby Simpson
// Sort Map Based on Key Length
package main

import (
	"fmt"
	"sort"
)

func main() {
	keys := make([]string, 0, 4)
	produce := map[string]int{
		"orange":     5,
		"apple":      7,
		"mango":      3,
		"strawberry": 9,
	}
	for k := range produce {
		keys = append(keys, k)
	}
	// sort by key length
	sort.Slice(keys, func(i, j int) bool {
		return len(keys[i]) < len(keys[j])
	})
	fmt.Println(keys)
}
