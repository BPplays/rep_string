package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
    if len(os.Args) < 4 {
        fmt.Println("Usage: go run main.go <dir> <placeholder> <replacement>")
        return
    }

    dir := os.Args[1]
    placeholder := os.Args[2]
    replacement := os.Args[3]

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
