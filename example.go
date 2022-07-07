package main

import "fmt"

func mapFromFn(item map[string]float64) map[string]float64 {
	db := make(map[string]float64)
	for k, v := range item {
		db[k] = v * 2
	}
	return db
}

func main() {
	p := make(map[string]float64)
	p["iPhone"] = 99256.50
	p["Android"] = 5000
	fmt.Println(mapFromFn(p))
}

// func main() {
// 	mapData := map[string]interface{}{
// 		"Name":    "noknown",
// 		"Age":     23,
// 		"Admin":   true,
// 		"Hobbies": []string{"IT", "Travel"},
// 		"Complex": struct {
// 			real int
// 			img  int
// 		}{12, -1},
// 	}
// 	fmt.Println(mapData)
// }
