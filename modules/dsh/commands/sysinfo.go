package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"gopkg.in/yaml.v3"
)

type SysinfoConfig struct {
	CPU struct {
		Model string `yaml:"model"`
		Cores int    `yaml:"cores"`
	} `yaml:"cpu"`
	Memory struct {
		Total string `yaml:"total"`
		Free  string `yaml:"free"`
	} `yaml:"memory"`
	OS struct {
		Name     string `yaml:"name"`
		Kernel   string `yaml:"kernel"`
		Uptime   string `yaml:"uptime"`
		Hostname string `yaml:"hostname"`
	} `yaml:"os"`
	Customixe struct {
		PromptStyle string `yaml:"prompt_style"`
		ShowTime    bool   `yaml:"show_time"`
	} `yaml:"customixe"`
	Colors struct {
		PromptUser string `yaml:"prompt_user"`
		PromptHost string `yaml:"prompt_host"`
		PromptPath string `yaml:"prompt_path"`
		PromptChar string `yaml:"prompt_char"`
		Reset      string `yaml:"reset"`
	} `yaml:"colors"`
}

func Sysinfo() {
	config, err := loadSysinfoConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error loading sysinfo config:", err)
		return
	}

	printSysinfo(config)
}

func loadSysinfoConfig() (*SysinfoConfig, error) {
	configPath := filepath.Join("configs", "sysinfo", "sysi.conf")
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("error opening sysinfo config file: %w", err)
	}
	defer file.Close()

	var config SysinfoConfig
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("error decoding sysinfo config file: %w", err)
	}

	return &config, nil
}

func printSysinfo(config *SysinfoConfig) {
	printCPUInfo(config)
	printMemoryInfo(config)
	printOSInfo(config)
	printCustomixeInfo(config)
}

func printCPUInfo(config *SysinfoConfig) {
	fmt.Println("CPU Info:")
	fmt.Printf("  Model: %s\n", getCPUInfo(config.CPU.Model))
	coresStr := strconv.Itoa(config.CPU.Cores) // Преобразование int в string
	fmt.Printf("  Cores: %s\n", getCPUInfo(coresStr))
}

func printMemoryInfo(config *SysinfoConfig) {
	fmt.Println("\nMemory Info:")
	fmt.Printf("  Total: %s\n", getMemoryInfo(config.Memory.Total))
	fmt.Printf("  Free: %s\n", getMemoryInfo(config.Memory.Free))
}

func printOSInfo(config *SysinfoConfig) {
	fmt.Println("\nOS Info:")
	fmt.Printf("  Name: %s\n", getOSInfo("name"))
	fmt.Printf("  Kernel: %s\n", getOSInfo("kernel"))
	fmt.Printf("  Uptime: %s\n", getOSInfo("uptime"))
	fmt.Printf("  Hostname: %s\n", getOSInfo("hostname"))
}

func printCustomixeInfo(config *SysinfoConfig) {
	fmt.Println("\nCustomixe Info:")
	fmt.Printf("  Prompt Style: %s\n", config.Customixe.PromptStyle)
	fmt.Printf("  Show Time: %t\n", config.Customixe.ShowTime)

	printPrompt(config)
}

func printPrompt(config *SysinfoConfig) {
	reset := config.Colors.Reset

	userColor := config.Colors.PromptUser
	hostColor := config.Colors.PromptHost
	pathColor := config.Colors.PromptPath
	charColor := config.Colors.PromptChar

	fmt.Printf("%s%s%s@%s%s:%s%s$%s ", userColor, os.Getenv("USER"), reset, hostColor, getHostname(), pathColor, getCurrentPath(), charColor)

	if config.Customixe.PromptStyle == "lean" {
		fmt.Print("❯ ")
	} else if config.Customixe.PromptStyle == "classic" {
		fmt.Print("› ")
	} else if config.Customixe.PromptStyle == "rainbow" {
		fmt.Print("› ")
	} else if config.Customixe.PromptStyle == "pure" {
		fmt.Print("> ")
	}
	fmt.Print(reset)
}

func getCPUInfo(key string) string {
	switch key {
	case "model":
		// Linux
		if runtime.GOOS == "linux" {
			cmd := exec.Command("lscpu", "-p")
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error executing lscpu:", err)
				return "Unknown"
			}

			lines := strings.Split(string(output), "\n")
			for _, line := range lines {
				if strings.HasPrefix(line, "Model name:") {
					parts := strings.Split(line, ":")
					if len(parts) > 1 {
						return strings.TrimSpace(parts[1])
					}
				}
			}
		}
		// macOS
		if runtime.GOOS == "darwin" {
			cmd := exec.Command("sysctl", "-n", "machdep.cpu.brand_string")
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error executing sysctl:", err)
				return "Unknown"
			}
			return strings.TrimSpace(string(output))
		}
		// FreeBSD
		if runtime.GOOS == "freebsd" {
			cmd := exec.Command("sysctl", "-n", "hw.model")
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error executing sysctl:", err)
				return "Unknown"
			}
			return strings.TrimSpace(string(output))
		}
	case "cores":
		// Linux
		if runtime.GOOS == "linux" {
			cmd := exec.Command("lscpu", "-p")
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error executing lscpu:", err)
				return "Unknown"
			}

			lines := strings.Split(string(output), "\n")
			cores := 0
			for _, line := range lines {
				if strings.HasPrefix(line, "Core(s) per socket:") {
					parts := strings.Split(line, ":")
					if len(parts) > 1 {
						coresPerSocket, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
						cores += coresPerSocket
					}
				}
			}
			return strconv.Itoa(cores)
		}
		// macOS
		if runtime.GOOS == "darwin" {
			cmd := exec.Command("sysctl", "-n", "hw.ncpu")
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error executing sysctl:", err)
				return "Unknown"
			}
			return strings.TrimSpace(string(output))
		}
		// FreeBSD
		if runtime.GOOS == "freebsd" {
			cmd := exec.Command("sysctl", "-n", "hw.ncpu")
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error executing sysctl:", err)
				return "Unknown"
			}
			return strings.TrimSpace(string(output))
		}
	}
	return "Unknown"
}

func getMemoryInfo(key string) string {
	switch key {
	case "total":
		switch runtime.GOOS {
		case "linux":
			cmd := exec.Command("free", "-m")
			output, err := cmd.Output()
			if err != nil {
				return "Error: " + err.Error()
			}
			lines := strings.Split(string(output), "\n")
			for _, line := range lines {
				if strings.HasPrefix(line, "Mem:") {
					parts := strings.Fields(line)
					if len(parts) > 1 {
						return parts[1]
					}
				}
			}
		case "darwin":
			cmd := exec.Command("vm_stat")
			output, err := cmd.Output()
			if err != nil {
				return "Error: " + err.Error()
			}
			lines := strings.Split(string(output), "\n")
			for _, line := range lines {
				if strings.Contains(line, "total_bytes") {
					parts := strings.Fields(line)
					if len(parts) > 2 {
						totalMemory := parts[2]
						totalMemoryMB := convertBytesToMegabytes(totalMemory)
						return totalMemoryMB
					}
				}
			}
		case "freebsd":
			cmd := exec.Command("sysctl", "-n", "hw.physmem")
			output, err := cmd.Output()
			if err != nil {
				return "Error: " + err.Error()
			}
			totalMemoryBytes := strings.TrimSpace(string(output))
			totalMemoryMB := convertBytesToMegabytes(totalMemoryBytes)
			return totalMemoryMB
		default:
			return "Unsupported operating system"
		}
	case "free":
		switch runtime.GOOS {
		case "linux":
			cmd := exec.Command("free", "-m")
			output, err := cmd.Output()
			if err != nil {
				return "Error: " + err.Error()
			}
			lines := strings.Split(string(output), "\n")
			for _, line := range lines {
				if strings.HasPrefix(line, "Mem:") {
					parts := strings.Fields(line)
					if len(parts) > 3 {
						return parts[3]
					}
				}
			}
		case "darwin":
			cmd := exec.Command("vm_stat")
			output, err := cmd.Output()
			if err != nil {
				return "Error: " + err.Error()
			}
			lines := strings.Split(string(output), "\n")
			for _, line := range lines {
				if strings.Contains(line, "free_bytes") {
					parts := strings.Fields(line)
					if len(parts) > 2 {
						freeMemory := parts[2]
						freeMemoryMB := convertBytesToMegabytes(freeMemory)
						return freeMemoryMB
					}
				}
			}
		case "freebsd":
			cmd := exec.Command("sysctl", "-n", "vm.avail_pages")
			output, err := cmd.Output()
			if err != nil {
				return "Error: " + err.Error()
			}
			availPages := strings.TrimSpace(string(output))
			availPagesUint, err := strconv.ParseUint(availPages, 10, 64) // Corrected: handle error
			if err != nil {
				return "Error: " + err.Error() // Corrected: return error string
			}
			availMemoryBytes := availPagesUint * 4096
			availMemoryMB := convertBytesToMegabytes(strconv.FormatUint(availMemoryBytes, 10))
			return availMemoryMB
		default:
			return "Unsupported operating system"
		}
	default:
		return ""
	}
	return ""
}

func convertBytesToMegabytes(bytes string) string {
	bytesInt, err := strconv.ParseUint(bytes, 10, 64)
	if err != nil {
		return "Error: " + err.Error()
	}
	megabytes := bytesInt / (1024 * 1024)
	return strconv.FormatUint(megabytes, 10)
}

func getOSInfo(key string) string {
	switch key {
	case "name":
		return runtime.GOOS
	case "kernel":
		// Implementation for kernel depends on the OS.  This is a placeholder.
		return "Unknown"
	case "uptime":
		// Implementation for uptime depends on the OS. This is a placeholder.
		return "Unknown"
	case "hostname":
		hostname, err := os.Hostname()
		if err != nil {
			return "Unknown"
		}
		return hostname
	default:
		return "Unknown"
	}
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "Unknown"
	}
	return hostname
}

func getCurrentPath() string {
	path, err := os.Getwd()
	if err != nil {
		return "Unknown"
	}
	return path
}
