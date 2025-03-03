package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"cty.sh/wc/pkg/buffered"
	"cty.sh/wc/pkg/common"
	"cty.sh/wc/pkg/fileoutput"
	"cty.sh/wc/pkg/semaphore"
	"cty.sh/wc/pkg/shared"
)

var (
	outputDir string
	mode      string
)

var rootCmd = &cobra.Command{
	Use:   "wc [file]",
	Short: "A word counter application with multiple concurrent implementations",
	Long: `A word counter application that demonstrates different concurrent programming patterns in Go.
Available modes:
  - sequential: Uses sequential processing
  - buffered:  Uses buffered channels
  - shared:    Uses shared memory with atomic operations
  - semaphore: Uses semaphores for concurrency control`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		file := args[0]

		// Verify file exists
		if _, err := os.Stat(file); err != nil {
			return fmt.Errorf("file not found: %s", file)
		}

		var stats common.FileStats
		var err error

		// Execute counting based on selected mode
		switch mode {
		case "sequential":
			stats, err = common.CountAll(file)
		case "buffered":
			stats, err = buffered.CountAllConcurrent(file)
		case "shared":
			stats, err = shared.CountAllConcurrent(file)
		case "semaphore":
			stats, err = semaphore.CountAllConcurrent(file)
		default:
			return fmt.Errorf("invalid mode: %s", mode)
		}

		if err != nil {
			return err
		}

		// Print results to stdout
		fmt.Printf("\t%d\t%d\t%d\t%s\n", stats.Lines, stats.Words, stats.Chars, stats.Path)

		// Save to file if output directory is specified
		if outputDir != "" {
			if err := fileoutput.CountAndSave(file, stats, outputDir); err != nil {
				return fmt.Errorf("failed to save output: %w", err)
			}
			fmt.Printf("Results saved to %s\n", filepath.Join(outputDir,
				fmt.Sprintf("%s_%s.json",
					filepath.Base(file),
					"latest")))
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outputDir, "output", "o", "", "Directory to save results (optional)")
	rootCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "sequential", "Processing mode (sequential|buffered|shared|semaphore)")
}
