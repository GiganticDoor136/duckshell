package dshfunc

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

type CustomixeSettings struct { 
	ASCII            int    `json:"ascii"`
	Unicode          int    `json:"unicode"`
	Font             string `json:"font,omitempty"`
	Diamond          int    `json:"diamond"`
	UpwardsArrow     int    `json:"upwards_arrow"`
	PromptStyleLean  int    `json:"prompt_style_lean"`
	PromptStyleClassic int    `json:"prompt_style_classic"`
	PromptStyleRainbow int    `json:"prompt_style_rainbow"`
	PromptStylePure  int    `json:"prompt_style_pure"`
}

var customizeSettings = CustomixeSettings{ 
	ASCII:            1,
	Unicode:          0,
	Font:             "",
	Diamond:          0,
	UpwardsArrow:     0,
	PromptStyleLean:  1,
	PromptStyleClassic: 0,
	PromptStyleRainbow: 0,
	PromptStylePure:  0,
}

func LoadCustomixeSettings() error {
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current directory: %w", err)
	}
	configPath := filepath.Join(currentDir, "configs", "customixe", "customixe.conf")
	// fontPath := filepath.Join(currentDir, "modules", "dsh", "func", "data", "fonts", "Raleway_Regular.ttf")
	file, err := os.Open(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Customixe config file not found, using defaults.") 
			customizeSettings = CustomixeSettings{
				ASCII:            1,
				Unicode:          0,
				Font:             "", 
				Diamond:          0,
				UpwardsArrow:     0,
				PromptStyleLean:  1,
				PromptStyleClassic: 0,
				PromptStyleRainbow: 0,
				PromptStylePure:  0,
			}
			return nil 
		}
		return fmt.Errorf("error opening customixe config file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&customizeSettings)
	if err != nil {
		return fmt.Errorf("error decoding customixe config file: %w", err)
	}

	return nil
}

func SaveCustomixeSettings() error {
	configPath := filepath.Join("configs", "customixe", "customixe.conf")
	file, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("error creating customixe config file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(customizeSettings)
	if err != nil {
		return fmt.Errorf("error encoding customixe config file: %w", err)
	}

	return nil
}

func EnableCustomixe() {
	fmt.Println("Enabling customize")
	customizeSettings.PromptStyleLean = 1 
	customizeSettings.PromptStyleClassic = 0
	customizeSettings.PromptStyleRainbow = 0
	customizeSettings.PromptStylePure = 0

	if err := SaveCustomixeSettings(); err != nil {
		fmt.Println("Error saving customixe settings:", err)
	}
}

func DisableCustomixe() {
	fmt.Println("Disabling customize")
	customizeSettings = CustomixeSettings{ 
		ASCII:            1,
		Unicode:          0,
		Font:             "",
		Diamond:          0,
		UpwardsArrow:     0,
		PromptStyleLean:  1,
		PromptStyleClassic: 0,
		PromptStyleRainbow: 0,
		PromptStylePure:  0,
	}

	if err := SaveCustomixeSettings(); err != nil {
		fmt.Println("Error saving customixe settings:", err)
	}
}

func CtlCustomixe(settingsStr string) {
    settingsParts := strings.Split(settingsStr, ",")
    for _, setting := range settingsParts {
        keyVal := strings.SplitN(setting, ":", 2)
        if len(keyVal) == 2 {
            key := keyVal[0]
            valueStr := keyVal[1]

            switch key {
            case "ASCII", "Unicode", "Diamond", "UpwardsArrow", "PromptStyleLean", "PromptStyleClassic", "PromptStyleRainbow", "PromptStylePure":
                value, err := strconv.Atoi(valueStr)
                if err != nil {
                    fmt.Println("Invalid value for setting", key, ":", valueStr, "Error:", err)
                    continue
                }
                reflect.ValueOf(&customizeSettings).Elem().FieldByName(strings.ToTitle(key)).SetInt(int64(value))
            case "Font":
                customizeSettings.Font = valueStr // Присваиваем строку напрямую, без Atoi
            default:
                fmt.Println("Unknown setting:", key)
            }

            fmt.Printf("Setting %s to %v\n", key, valueStr)
        }
    }

    if err := SaveCustomixeSettings(); err != nil {
        fmt.Println("Error saving customixe settings:", err)
    }
}