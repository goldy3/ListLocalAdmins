package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <input_file_path> <output_file_path>")
		os.Exit(1)
	}

	inputFilePath := os.Args[1]
	outputFilePath := os.Args[2]

	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("Failed to open input file: %s", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatalf("Failed to create output file: %s", err)
	}
	defer outputFile.Close()

	csvWriter := csv.NewWriter(outputFile)
	defer csvWriter.Flush()

	// Write headers to the CSV file
	headers := []string{"Server", "UserName", "OtherDetail"}
	if err := csvWriter.Write(headers); err != nil {
		log.Fatalf("Failed to write headers to CSV file: %s", err)
	}

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		serverNames := strings.Split(line, "\n")
		for _, serverName := range serverNames {
			serverName = strings.TrimSpace(serverName)
			if serverName == "" {
				continue
			}

			fmt.Printf("Executing on server: %s\n", serverName)
			results := executePowerShellCommand(serverName)
			for _, result := range results {
				if err := csvWriter.Write([]string{serverName, result[0], result[1]}); err != nil {
					log.Printf("Failed to write to CSV file: %s", err)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to read input file: %s", err)
	}
}

func executePowerShellCommand(serverName string) [][]string {
	cmd := exec.Command("powershell", "Invoke-Command", "-ComputerName", serverName, "-ScriptBlock", `{Get-LocalGroupMember -Group "Administrators" | Select-Object Name, PrincipalSource | Format-List}`)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing command on %s: %s\n", serverName, err)
		return [][]string{{"Error", err.Error()}}
	}

	return parsePowerShellOutput(string(output))
}

// parsePowerShellOutput parses the structured output into slice of slice string for CSV writing
func parsePowerShellOutput(output string) [][]string {
	lines := strings.Split(output, "\n")
	var results [][]string
	var current []string

	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			current = append(current, strings.TrimSpace(parts[1]))
		}
		if len(current) == 2 { // Assuming each entry has 2 elements (Name, PrincipalSource)
			results = append(results, current)
			current = nil
		}
	}

	return results
}
