package commands

import (
	"fmt"

	"github.com/GiganticDoor136/duckshell/modules/dsh/func"
)

func Enable(funcName string) {
	switch funcName {
	case "customixe":
		err := dshfunc.LoadCustomixeSettings()
		if err != nil {
			fmt.Println("Error loading customixe settings:", err)
		}
		dshfunc.EnableCustomixe()
	default:
		fmt.Println("Enabling function:", funcName)
	}
}

func Disable(funcName string) {
	switch funcName {
	case "customixe":
		dshfunc.DisableCustomixe()
	default:
		fmt.Println("Disabling function:", funcName)
	}
}
