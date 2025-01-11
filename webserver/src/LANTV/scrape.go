package lantv

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var currentCmd *exec.Cmd

func Search_Series(input string) [][]string {
	fmt.Printf("search input: %s\n", input)

	if currentCmd != nil && currentCmd.Process != nil {
		err := currentCmd.Process.Kill()
		if err != nil {
			log.Printf("Failed to kill previous command: %v\n", err)
		}
		fmt.Printf("------KILL PREV COMMAND----------\n")
	}

	// Create a new context with cancellation
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd := exec.Command("python3", "../scripts/searchSeries.py", input)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Failed to run command: %v\n", err)
	}

	fmt.Printf("raw Output: %s", output)
	var sb strings.Builder
	for _, code := range output {
		sb.WriteByte(byte(code))
	}

	fmt.Printf("parsing\n")
	// Parse the JSON output
	var data [][]interface{}
	err = json.Unmarshal(output, &data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("converting\n")
	// Convert the results to a string slice
	var results [][]string
	for _, result := range data {
		results = append(results, []string{result[0].(string), result[1].(string)})
		fmt.Printf("Title: %s, Key: %s\n", result[0], result[1])
	}

	fmt.Printf("Results: %v\n", results)
	return results
}

// scrape the bare minimum information for a series
func Preview_Series(key string) [][]string {
	fmt.Printf("create input: %s\n", key)

	cmd := exec.Command("python3", "../scripts/scrapeSeriesPreview.py", key)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Failed to run command: %v\n", err)
		return nil
	}

	// Parse the JSON output
	var results [][]string
	err = json.Unmarshal(output, &results)
	if err != nil {
		log.Printf("Failed to unmarshal JSON: %v\n", err)
		return nil
	}

	// Print the results
	for _, result := range results {
		for _, item := range result {
			fmt.Printf("Result: %s\n", item)
		}
	}

	return results
}

// scrape the full information for a series
func Create_Series(key string) {
	fmt.Printf("create input: %s\n", key)

	cmd := exec.Command("python3", "../scripts/scrapeSeries.py", key)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Failed to run command: %v\n", err)
	}

	// Convert the byte slice to a string
	outputStr := string(output)
	fmt.Printf("Output: %s\n", outputStr)
}
