package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Open file for read
	filerc, err := os.Open("002-longest-sentence/data.txt")
	check(err)
	defer filerc.Close()
	// read file
	buf := new(bytes.Buffer)
	buf.ReadFrom(filerc)
	contents := buf.String()
	// Replace newline to space
	re := regexp.MustCompile(`\r?\n`)
	contents = re.ReplaceAllLiteralString(contents, " ")
	// Get sentences
	arrayContents := strings.Split(contents, ".")
	biggest := 0
	for _, s := range arrayContents {
		// Get words in sentence
		words := strings.Fields(s)
		if len(words) > biggest {
			biggest = len(words)
		}
	}
	fmt.Println(biggest)
}
