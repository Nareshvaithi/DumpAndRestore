package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern := "test"
	text := "dsjhdsbbuuhoiihtest//"

	matched, err := regexp.MatchString(pattern, text)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Matched:", matched)
}
