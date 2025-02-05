package sysinfo

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/GiganticDoor136/duckshell/modules/dsh/func"
)

type SysinfoConfig struct {
	CPUModel      string
	CPUCoors      string
	MemTotal      string
	MemFree       string
	OSName        string
	KernelVersion string
}

func LoadSysinfoConfig() (*SysinfoConfig, error) {
	return &SysinfoConfig{}, nil 
}

func PrintSysinfo(config *SysinfoConfig) {
	printCPUInfo(config)
	printMemoryInfo(config)
	printOSInfo(config)

	customixeSettings := GetCustomixeSettings()
	fmt.Println("Current Customixe Settings:", customixeSettings) 
}

func printCPUInfo(config *SysinfoConfig) {
	fmt.Print("CPU Info:\n")
	// Linux
	if runtime.GOOS == "linux" {
		printLinuxCPUInfo(config)
	}
	// macOS
	if runtime.GOOS == "darwin" {
		printMacOSCPUInfo(config)
	}
	// FreeBSD
	if runtime.GOOS == "freebsd" {
		printFreeBSDCUInfo(config)
	}
}

func printLinuxCPUInfo(config *SysinfoConfig) {
	cpuModel := getCPUModel()
	config.CPUModel = cpuModel
	fmt.Printf("CPU Model: %s\n", cpuModel)

	cpuCores := getCPUCoors()
	config.CPUCoors = cpuCores
	fmt.Printf("CPU Cores: %s\n", cpuCores)
}

func printMacOSCPUInfo(config *SysinfoConfig) {
	cmd := exec.Command("sysctl", "-n", "machdep.cpu.brand_string")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error executing sysctl:", err)
		return
	}
	fmt.Print(string(output))
}

func printFreeBSDCUInfo(config *SysinfoConfig) {
	cmd := exec.Command("sysctl", "-n", "hw.model")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error executing sysctl:", err)
		return
	}
	fmt.Print(string(output))
}

func getCPUModel() string {
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

	return "Unknown"
}

func getCPUCoors() string {
	cmd := exec.Command("lscpu", "-p")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error executing lscpu:", err)
		return "Unknown"
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Core(s) per socket:") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				return strings.TrimSpace(parts[1])
			}
		}
	}

	return "Unknown"
}

func printMemoryInfo(config *SysinfoConfig) {
	fmt.Print("\nMemory Info:\n")
	if runtime.GOOS == "linux" {
		printLinuxMemoryInfo(config)
	}
	if runtime.GOOS == "darwin" {
		printMacOSMemoryInfo(config)
	}
	if runtime.GOOS == "freebsd" {
		printFreeBSDMemoryInfo(config)
	}
}

func printLinuxMemoryInfo(config *SysinfoConfig) {
	memTotal, memFree := getLinuxMemoryInfo()
	config.MemTotal = memTotal
	config.MemFree = memFree
	fmt.Printf("Mem Total: %s\n", memTotal)
	fmt.Printf("Mem Free: %s\n", memFree)
}

func getLinuxMemoryInfo() (string, string) {
	cmd := exec.Command("free", "-m")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error executing free:", err)
		return "Unknown", "Unknown"
	}
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Mem:") {
			parts := strings.Fields(line)
			if len(parts) > 3 {
				return parts[1], parts[3]
			}
		}
	}
	return "Unknown", "Unknown"
}

func printMacOSMemoryInfo(config *SysinfoConfig) {
	cmd := exec.Command("vm_stat")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error executing vm_stat:", err)
		return
	}
	fmt.Print(string(output))
}

func printFreeBSDMemoryInfo(config *SysinfoConfig) {
	cmd := exec.Command("vmstat", "-n")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error executing vmstat:", err)
		return
	}
	fmt.Print(string(output))
}

func printOSInfo(config *SysinfoConfig) {
	fmt.Print("\nOS Info:\n")
	if runtime.GOOS == "linux" {
		printLinuxOSInfo(config)
	}
	if runtime.GOOS == "darwin" {
		printMacOSInfo(config)
	}
	if runtime.GOOS == "freebsd" {
		printFreeBSDOSInfo(config)
	}
}

func printLinuxOSInfo(config *SysinfoConfig) {
	osName := getLinuxOSName()
	kernelVersion := getLinuxKernelVersion()
	config.OSName = osName
	config.KernelVersion = kernelVersion
	fmt.Printf("OS Name: %s\n", osName)
	fmt.Printf("Kernel Version: %s\n", kernelVersion)
}

func getLinuxOSName() string {
	cmd := exec.Command("lsb_release", "-a")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error executing lsb_release:", err)
		return "Unknown"
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Distributor ID:") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				return strings.TrimSpace(parts[1])
			}
		}
	}

	return "Unknown"
}

func getLinuxKernelVersion() string {
	cmd := exec.Command("uname", "-r")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error executing uname:", err)
		return "Unknown"
	}
	return strings.TrimSpace(string(output))
}

func printMacOSInfo(config *SysinfoConfig) {
	cmd := exec.Command("sw_vers")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error executing sw_vers:", err)
		return
	}
	fmt.Print(string(output))
}

func printFreeBSDOSInfo(config *SysinfoConfig) {