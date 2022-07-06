// Shelby Simpson 6/30/22
// Coffee Shop Assessment
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var size, coffeeType, flavor, flavorType string
	var cost float32

SIZE:
	fmt.Print("Do you want small, medium, or large? ")
	fmt.Scan(&size)
	size = strings.ToLower(size)
	switch size {
	case "small":
		cost += 2
	case "medium":
		cost += 3
	case "large":
		cost += 4
	default:
		fmt.Println("You entered an incorrect size.  Please try again.")
		goto SIZE
	}
COFFEETYPE:
	fmt.Print("\nDo you want brewed, espresso, or cold brew? ")
	scanner := bufio.NewScanner(os.Stdin)
	if ok := scanner.Scan(); ok {
		coffeeType = scanner.Text()
	}
	coffeeType = strings.ToLower(coffeeType)
	switch coffeeType {
	case "brewed":
	case "espresso":
		cost += .5
	case "cold brew":
		cost += 1
	default:
		fmt.Println("You entered an incorrect type.  Please try again.")
		goto COFFEETYPE
	}
FLAVOR:
	fmt.Print("\nDo you want a flavored syrup? ")
	fmt.Scan(&flavor)
	flavor = strings.ToLower(flavor)
FLAVORTYPE:
	if flavor == "yes" {
		fmt.Print("Do you want hazelnut, vanilla, or caramel? ")
		fmt.Scan(&flavorType)
		flavorType = strings.ToLower(flavorType)
		switch flavorType {
		case "hazelnut":
			cost += .5
		case "vanilla":
			cost += .5
		case "caramel":
			cost += .5
		default:
			fmt.Println("You entered an incorrect type of flavor.  Please try again.")
			goto FLAVORTYPE
		}
	} else if flavor != "no" {
		fmt.Println("You did not indicate whether you wanted a flavor.  Please try again.")
		goto FLAVOR
	}

	order := fmt.Sprintf("You asked for a %v cup of %v coffee with", size, coffeeType)
	if flavor == "yes" {
		order += " " + flavorType + " syrup.\n"
	} else {
		order += "out syrup.\n"
	}
	fmt.Printf(order)
	fmt.Printf("Your cup of coffee costs %.2f\n", cost)
	fmt.Printf("The price with a tip is %.2f\n", cost+(cost*.15))
}
