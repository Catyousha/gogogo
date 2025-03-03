package fileoutput

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"cty.sh/wc/pkg/common"
)

// FileOutput represents the output format for file statistics
type FileOutput struct {
	Stats    common.FileStats `json:"stats"`
	Time     time.Time        `json:"time"`
	Duration time.Duration    `json:"duration"`
}

// CountAndSave saves the provided file statistics to an output file
func CountAndSave(inputFile string, stats common.FileStats, outputDir string) error {
	start := time.Now()

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Prepare output data
	output := FileOutput{
		Stats:    stats,
		Time:     time.Now(),
		Duration: time.Since(start),
	}

	// Create output file name based on input file and timestamp
	outputFile := filepath.Join(outputDir,
		fmt.Sprintf("%s_%s.json",
			filepath.Base(inputFile),
			time.Now().Format("20060102_150405")))

	// Marshal data to JSON
	data, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal output: %w", err)
	}

	// Write to file
	if err := os.WriteFile(outputFile, data, 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	return nil
}

// ReadResults reads the results from a previously saved output file
func ReadResults(outputFile string) (*FileOutput, error) {
	data, err := os.ReadFile(outputFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read output file: %w", err)
	}

	var output FileOutput
	if err := json.Unmarshal(data, &output); err != nil {
		return nil, fmt.Errorf("failed to unmarshal output: %w", err)
	}

	return &output, nil
}
