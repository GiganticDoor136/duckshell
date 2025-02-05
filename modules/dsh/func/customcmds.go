package dshfunc

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var customCmds = make(map[string]string)

func LoadCustomCmds() error {
	configPath := filepath.Join("configs", "ccmds", "ccmds.conf")
	file, err := os.Open(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Custom commands config file not found, using defaults.")
			return nil
		}
		return fmt.Errorf("error opening custom commands config file: %w", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error reading custom commands config file: %w", err)
		}

		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Split(line, " >> ")
		if len(parts) != 2 {
			continue
		}

		existingCmdPart := parts[0]
		newCmdPart := parts[1]

		existingCmd := strings.TrimSpace(strings.Trim(existingCmdPart, "(cmd):"))
		newCmd := strings.TrimSpace(strings.Trim(newCmdPart, "(cmd):"))

		customCmds[newCmd] = existingCmd
	}

	return nil
}
