// File: main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"squeezlogs/parser"
	"squeezlogs/utils"
	"io/ioutil"
)

func main() {
	inputPath := flag.String("input", "sample/logs.json", "Path to input log JSON file")
	outputPath := flag.String("output", "output/filtered.json", "Path to output filtered log file")
	flag.Parse()

	logs, err := utils.ReadLogs(*inputPath)
	if err != nil {
		log.Fatalf("Error reading logs: %v", err)
	}

	deduped := parser.Deduplicate(logs)

	err = utils.WriteLogs(*outputPath, deduped)
	if err != nil {
		log.Fatalf("Error writing filtered logs: %v", err)
	}

	fmt.Printf("Processed %d unique log messages.\n", len(deduped))
	data, _ := ioutil.ReadFile(*outputPath)
	fmt.Println("logs that are ready to be ingested: ", string(data))
}
