package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	directory := "D:\\Study\\Magistracy\\System Administration and Programming\\Labs\\1\\2 lab"
	fileExtension := "jpg"
	findFilesByExtension(directory, fileExtension) //1

	fileDir := "D:\\Study\\Magistracy\\System Administration and Programming\\Labs\\1\\2 lab\\test2.txt"
	wordLength := 8
	readWords(fileDir, wordLength) //5
}

func findFilesByExtension(directory string, fileExtension string) {
	pattern := fmt.Sprintf("%s/*.%s", directory, fileExtension)

	files, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatalf("Error finding files: %v", err)
	}

	for _, filePath := range files {
		fmt.Println(filePath)
	}
}

func readWords(filename string, wordLength int) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		for _, word := range words {
			word = removeSeparators(word)
			if len([]rune(word)) == wordLength {
				fmt.Println(word)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func removeSeparators(s string) string {
	re := regexp.MustCompile("[,.!?-]")

	return re.ReplaceAllString(s, "")
}
