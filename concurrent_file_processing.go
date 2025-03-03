package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func processFile(filename string, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filename, err)
		results <- 0 
		return
	}
	defer file.Close()

	wordCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text()) 
		wordCount += len(words)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		results <- 0
		return
	}

	results <- wordCount
}

func main() {
	files := []string{"twilio_2FA_recovery_code.txt", "ifrs-9-financial-instruments.txt", "ngrok_recovery_codes.txt"} // Replace with actual file paths

	results := make(chan int, len(files)) 
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go processFile(file, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	totalWords := 0
	for count := range results {
		totalWords += count
	}

	fmt.Printf("Total words across all files: %d\n", totalWords)
}
