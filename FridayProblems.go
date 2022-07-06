// Shelby Simpson
// Friday 7/1/22 Problems
package main

import "fmt"

// 1.
// Signature: func sum(slice []int) int
func sum(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return sum
}

// 2.
func half(num int) (int, bool) {
	var isEven bool
	var half int
	if num%2 == 0 {
		isEven = true
	} else {
		isEven = false
	}
	half = num / 2
	return half, isEven
}

// 3.
func findGreatest(numbers ...float64) float64 {
	max := numbers[0]
	for _, val := range numbers {
		if val > max {
			max = val
		}
	}
	return max
}

// 4.
func makeOddGenerator() func() int {
	num := -1
	return func() int {
		num += 2
		return num
	}
}

// 5. Updated with memoization
func fibMemo(n int) int {
	cache := make(map[int]int)
	return fibonacci(n, cache)
}

func fibonacci(n int, cache map[int]int) int {
	if n < 2 {
		return n
	} else if _, ok := cache[n]; ok {
		return cache[n]
	} else {
		return fibonacci(n-1, cache) + fibonacci(n-2, cache)
	}
}

func main() {
	half, isEven := half(1)
	fmt.Println("Half and isEven: ", half, isEven)
	fmt.Println("Greatest number: ", findGreatest(10.01, 2.01, 7, 10.001))
	fmt.Println("Fibonacci: ", fibMemo(5))
	nextOddNum := makeOddGenerator()
	fmt.Println("Next odd num: ", nextOddNum())
	fmt.Println("Next odd num: ", nextOddNum())
	fmt.Println("Next odd num: ", nextOddNum())
}
