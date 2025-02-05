package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/GiganticDoor136/duckshell/modules/dsh/func"
)

func Ctl(ctlStr string) {
	parts := strings.SplitN(ctlStr, " ", 2)
	if len(parts) != 2 {
		fmt.Fprintln(os.Stderr, "Invalid control string format. Use: dsh --ctl <function> <setting:value>")
		return
	}

	funcName := parts[0]
	settings := parts[1]

	switch funcName {
	case "customixe":
		dshfunc.CtlCustomixe(settings)
		err := dshfunc.SaveCustomixeSettings()
		if err != nil {
			fmt.Println("Error saving customixe settings:", err)
		}
	case "customcmds":
		dshfunc.CtlCustomCmds(settings)
	default:
		fmt.Printf("Controlling function: %s with settings: %s\n", funcName, settings)
	}
}
