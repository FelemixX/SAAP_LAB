package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	directory := "D:\\Study\\Magistracy\\System Administration and Programming\\Labs\\1\\2 lab"
	fileExtension := "jpg"
	findFilesByExtension(directory, fileExtension) //1

	fileDir := "D:\\Study\\Magistracy\\System Administration and Programming\\Labs\\1\\2 lab\\test.txt"
	wordLength := 4
	readWords(fileDir, wordLength) //2
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
		fmt.Println(scanner.Text())
		words := strings.Fields(scanner.Text())
		for _, word := range words {
			if len(word) == wordLength {
				fmt.Print(word)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
