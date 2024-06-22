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

	fileoverrep := false

    var fileFlag, repFlag, dir, placeholder string

    pflag.StringP("file", "f", "", "file to read from")
    pflag.StringP("rep", "r", "", "replace")
    pflag.StringP("dir", "d", "", "directory")
    pflag.StringP("placeholder", "p", "", "placeholder")

    pflag.Parse()

    if viper.IsSet("file") {
        fileFlag = viper.GetString("file")
		fileoverrep = true
    }
    if viper.IsSet("rep") {
        repFlag = viper.GetString("rep")
		if fileoverrep {
			fmt.Println("Please specify -f OR -r")
			os.Exit(1)
		}
		fileoverrep = false
    }
    if viper.IsSet("dir") {
        dir = viper.GetString("dir")
    }
    if viper.IsSet("placeholder") {
        placeholder = viper.GetString("placeholder")
    }

    // Check if at least one flag is set
    if fileFlag == "" && repFlag == "" && dir == ""  && placeholder == ""   {
        fmt.Println("Please specify (-f or -r) and -d and -p flag.")
        os.Exit(1)
    }

	// Access flag values
	var replacement string

	if fileoverrep {
		inf, err := os.ReadFile(fileFlag)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		replacement = string(inf)
	} else {
		replacement = repFlag
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
