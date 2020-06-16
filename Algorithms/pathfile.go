package main

import "fmt"

func main() {
	fmt.Printf("the answer is %s", name("23455854209350894"))
}

// recru
func name(s string) string {
	ls := len(s)
	if ls <= 3 {
		return s
	}
	return name(s[:ls-3]) + "," + s[ls-3:]

}
