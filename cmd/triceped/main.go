package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jessevdk/go-flags"
	"github.com/ru84/triceped/pkg/converter"
)

var opts struct {
	Verbose    []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	InputFile  string `short:"i" long:"input" description:"Input Terraform file" required:"true"`
	OutputFile string `short:"o" long:"output" description:"Output Bicep file (optional, defaults to <input>.bicep)"`
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	parser := flags.NewParser(&opts, flags.Default)
	parser.ShortDescription = "helps you keep your API well described"
	if _, err := parser.Parse(); err != nil {
		return err
	}

	// Set default output file if not provided
	if opts.OutputFile == "" {
		opts.OutputFile = defaultOutputFile(opts.InputFile)
	}

	// Convert Terraform to Bicep
	output, err := converter.TerraformToBicep(opts.InputFile)
	if err != nil {
		return fmt.Errorf("converting terraform to bicep: %w", err)
	}

	f, err := os.Create(opts.OutputFile)
	if err != nil {
		return fmt.Errorf("creating output file %q: %w", opts.OutputFile, err)
	}
	defer f.Close()

	if _, err := f.WriteString(output); err != nil {
		return fmt.Errorf("writing to output file: %w", err)
	}

	fmt.Printf("Conversion complete, bicep file saved as %q\n", opts.OutputFile)
	return nil
}

// defaultOutputFile generates the default output file name based on the input file
func defaultOutputFile(inputFile string) string {
	baseName := filepath.Base(inputFile)
	return baseName[:len(baseName)-len(filepath.Ext(baseName))] + ".bicep"
}
