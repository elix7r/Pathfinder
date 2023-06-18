package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	dirPath string
	homeDir = os.Getenv("HOME")
)

func showUsage() {
	fmt.Println("Usage: Pathfinder [options]")
	fmt.Println("Options:")
	flag.PrintDefaults()
}

func main() {
	flag.StringVar(&dirPath, "path", homeDir, "Path to your project directory")
	showHelp := flag.Bool("help", false, "Show help")

	flag.Parse()

	if *showHelp {
		showUsage()
		os.Exit(0)
	}

	dir, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Printf("You inputted wrong directory path: %v\n", err)
		return
	}

	fmt.Printf("+----------------------+--------------------------------+\n")
	fmt.Printf("| %-20s | %-30s |\n", "IDE", "Project")
	fmt.Printf("+----------------------+--------------------------------+\n")

	for _, entry := range dir {
		if entry.IsDir() {
			dirName := entry.Name()
			projectName := strings.TrimPrefix(dirName, "IDE")

			baseDir := filepath.Base(dirPath)
			fmt.Printf("| %-20s | %-30s |\n", baseDir, projectName)
		}
	}

	fmt.Printf("+----------------------+--------------------------------+\n")
}
