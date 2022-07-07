// Shelby Simpson
// Map of Structs
package main

import (
	"fmt"
	"sort"
)

type acc struct {
	name   string
	number int
}

func main() {
	temp := map[string]acc{"Me": {name: "Me", number: 1},
		"You": {name: "You", number: 2}}
	key := make([]string, 0, len(temp))
	for i := range temp {
		key = append(key, i)
	}
	sort.Strings(key)
	fmt.Println(key)
}
