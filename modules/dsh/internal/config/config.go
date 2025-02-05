package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type SysinfoConfig struct {
	CPUModel      string `json:"cpu_model"`
	CPUCoreCount  int    `json:"cpu_core_count"` 
	MemoryTotal   string `json:"mem_total"`
	MemoryFree    string `json:"mem_free"`
	OSName        string `json:"os_name"`
	KernelVersion string `json:"kernel_version"`
}

func LoadSysinfoConfig() (*SysinfoConfig, error) {
	configPath := filepath.Join("configs", "sysinfo", "sysi.conf")
	file, err := os.Open(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Sysinfo config file not found, using defaults.") 
			return &SysinfoConfig{}, nil                              
		}
		return nil, fmt.Errorf("error opening sysinfo config file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var config SysinfoConfig
	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("error decoding sysinfo config file: %w", err)
	}

	return &config, nil
}

func SaveSysinfoConfig(config *SysinfoConfig) error {
	configPath := filepath.Join("configs", "sysinfo", "sysi.conf")
	file, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("error creating sysinfo config file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(config)
	if err != nil {
		return fmt.Errorf("error encoding sysinfo config file: %w", err)
	}

	return nil
}