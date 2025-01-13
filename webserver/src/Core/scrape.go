package core

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var currentCmd *exec.Cmd

// Set the PYTHONPATH environment variable
func Init() {
	os.Setenv("PYTHONPATH", "/home/server/.local/lib/python3.10/site-packages")
}

func Search_Series(dbname string, input string) [][]string {
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

	cmd := exec.Command("python3", "../scripts/searchSeries.py", dbname, input)
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
func Preview_Series(key string) (data []string) {
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
		if len(result) == 1 {
			data = append(data, result[0])
		} else {
			data = append(data, strings.Join(result, ", "))
		}
	}

	return data
}

// scrape the full information for a series
func Create_Series(key string) Series {
	fmt.Printf("create input: %s\n", key)

	// get all the remaining data for the series
	cmd := exec.Command("python3", "../scripts/scrapeSeries.py", key)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Failed to run command: %v\n", err)
	}

	// Parse the JSON output
	var results []string
	err = json.Unmarshal(output, &results)
	if err != nil {
		log.Printf("Failed to unmarshal JSON: %v\n", err)
	}

	series := Series{}
	series.Title = results[0]
	// series.Subtitle = results[1]
	series.Image = results[2]
	// series.Description = results[3]
	// series.Genres = results[4]
	// series.Rating = results[5]
	// series.Year = results[6]
	// series.Seasons = results[7]
	// series.Episodes = results[8]

	//return an array of strings of the data
	return series
}
