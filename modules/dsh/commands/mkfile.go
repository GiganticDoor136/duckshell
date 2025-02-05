package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

func Mkfile(args []string) {
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Usage: mkfile <path>")
		return
	}

	path := args[0]
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating directories:", err)
		return
	}

	file, err := os.Create(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Println("File created:", path)
}
