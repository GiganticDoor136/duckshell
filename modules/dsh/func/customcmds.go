package dshfunc

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/GiganticDoor136/duckshell/modules/dsh/func"
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

func GetCustomCmds() map[string]string {
	cmdsCopy := make(map[string]string)
	for k, v := range customCmds {
		cmdsCopy[k] = v
	}
	return cmdsCopy
}

func RunCustomCmd(cmdName string) error {
	cmd, ok := customCmds[cmdName]
	if !ok {
		return fmt.Errorf("custom command '%s' not found", cmdName)
	}

	parts := strings.Fields(cmd)
	if len(parts) == 0 {
		return fmt.Errorf("invalid custom command definition")
	}
	commandName := parts[0]
	var args []string
	if len(parts) > 1 {
		args = parts[1:]
	}

	c := exec.Command(commandName, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	err := c.Run()

	if err != nil {
		return fmt.Errorf("error running custom command '%s': %w", cmdName, err)
	}

	return nil
}

func CtlCustomCmds(settingsStr string) {
	settingsParts := strings.Split(settingsStr, ",")
	for _, setting := range settingsParts {
		keyVal := strings.SplitN(setting, ":", 2)
		if len(keyVal) == 2 {
			key := keyVal[0]
			valueStr := keyVal[1]

			customCmds[key] = valueStr

			fmt.Printf("Setting custom command %s to %v\n", key, valueStr)
		}
	}

	if err := SaveCustomCmds(); err != nil {
		fmt.Println("Error saving custom commands:", err)
	}
}

func SaveCustomCmds() error {
	configPath := filepath.Join("configs", "ccmds", "ccmds.conf")
	file, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("error creating custom commands config file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for cmd, expanded := range customCmds {
		line := fmt.Sprintf("(cmd):%s >> (cmd):%s\n", cmd, expanded)
		_, err := writer.WriteString(line)
		if err != nil {
			return fmt.Errorf("error writing custom commands config file: %w", err)
		}
	}

	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("error flushing custom commands config file: %w", err)
	}

	return nil
}

func DisableCustomCmds() {
	customCmds = make(map[string]string)

	if err := SaveCustomCmds(); err != nil {
		fmt.Println("Error saving custom commands:", err)
	}
}

func CustomCmds() { // Оставляем ТОЛЬКО ЭТО объявление функции
	err := dshfunc.LoadCustomCmds()
	if err != nil {
			fmt.Fprintln(os.Stderr, "Error loading custom commands:", err)
			return
	}

	cmds := dshfunc.GetCustomCmds()

	if len(cmds) == 0 {
			fmt.Println("No custom commands available.")
			return
	}

	fmt.Println("Available custom commands:")
	for name := range cmds {
			fmt.Println("-", name)
	}

	fmt.Print("Enter the name of the command to run: ")

	var cmdName string
	_, err = fmt.Scanln(&cmdName)
	if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			return
	}

	err = dshfunc.RunCustomCmd(cmdName)
	if err != nil {
			fmt.Fprintln(os.Stderr, "Error running command:", err)
			return
	}
}