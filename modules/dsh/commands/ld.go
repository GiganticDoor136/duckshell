package commands

import (
	"fmt"
	"os"
	"os/exec"
//	"path/filepath"
	"strings"
)

func Ld(args []string) {
	longFormat := false
	path := "."

	for i := 0; i < len(args); i++ {
		arg := args[i]
		if arg == "-l" {
			longFormat = true
		} else if !strings.HasPrefix(arg, "-") {
			path = arg
		}
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Directory '%s' not found\n", path)
		return
	}

	var cmd *exec.Cmd
	if longFormat {
		cmd = exec.Command("ls", "-l", path)
	} else {
		cmd = exec.Command("ls", path)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing ls: %v\n", err)
		return
	}

	fmt.Print(string(output))
}
