package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {

    var fileFlag, repFlag, dir, placeholder, replacement string

    pflag.StringVarP(&fileFlag, "file", "f", "", "file to read from")
    pflag.StringVarP(&repFlag, "rep", "r", "", "replace")
    pflag.StringVarP(&dir, "dir", "d", "", "directory")
    pflag.StringVarP(&placeholder, "placeholder", "p", "", "placeholder")

    pflag.Parse()
    viper.BindPFlags(pflag.CommandLine) // Bind pflag to viper

    // Decide on the replacement based on flags
    if repFlag != "" {
        replacement = repFlag
    } else {
		inf, err := os.ReadFile(fileFlag)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
        replacement = string(inf)
    }
    // Output the values
    fmt.Printf("fileFlag: %s\n", fileFlag)
    fmt.Printf("repFlag: %s\n", repFlag)
    fmt.Printf("dir: %s\n", dir)
    fmt.Printf("placeholder: %s\n", placeholder)
    fmt.Printf("replacement: %s\n", replacement)
    // Check if at least one flag is set
    if replacement == "" {
        fmt.Println("Please specify either -f or -r flag.")
        os.Exit(1)
    }







	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			err = processFile(path, placeholder, replacement)
			if err != nil {
				fmt.Printf("Error walking the path %q: %v\n", dir, err)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", dir, err)
	}

	fmt.Println("Replacement complete.")
}

func processFile(filePath, placeholder, replacement string) error {
	input, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	
	output := strings.ReplaceAll(string(input), placeholder, replacement)

	err = os.WriteFile(filePath, []byte(output), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Processed file: %s\n", filePath)
	return nil
}
