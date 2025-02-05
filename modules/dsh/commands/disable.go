package commands

import (
	"fmt"

	"github.com/GiganticDoor136/duckshell/modules/dsh/func"
)

func Disable(funcName string) {
	switch funcName {
	case "customixe":
		dshfunc.DisableCustomixe()
	case "customcmds":
		dshfunc.DisableCustomcmds()
	default:
		fmt.Println("Disabling function:", funcName)
	}
}
