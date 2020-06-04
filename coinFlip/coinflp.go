//Simple use of switch statements
package main

import "fmt"

func sigNum(x int) int {
	switch {
	case x > 0:
		return +1
	case x < 0:
		return -1
	default:
		return 0
	}
}

func main() {
	x := -50
	y := sigNum(x)
	switch {
	case y == 1:
		fmt.Printf("number %d is greater than %d ", x, y)
	case y == -1:
		fmt.Printf("number %d is less than %d ", x, y)
	case y == 0:
		fmt.Printf("number is %d ", y)
	}

}
