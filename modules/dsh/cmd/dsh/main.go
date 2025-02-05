package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/GiganticDoor136/duckshell/modules/dsh/commands"
	"github.com/GiganticDoor136/duckshell/modules/dsh/func"
)

func main() {
	ver := flag.Bool("ver", false, "Display Duckshell version")
	sysinfoFlag := flag.Bool("sysinfo", false, "Display system information")
	help := flag.Bool("h", false, "Display help message")
	helpLong := flag.Bool("help", false, "Display help message")
	enable := flag.String("enable", "", "Enable a function")
	disable := flag.String("disable", "", "Disable a function")
	ctl := flag.String("ctl", "", "Control a function")

	flag.Parse()

	err := dshfunc.LoadCustomCmds()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error loading custom commands:", err)
	}

	if *ver {
		fmt.Println("Duckshell version: 1.0.0")
		return
	}

	if *sysinfoFlag {
		commands.Sysinfo()
		return
	}

	if *help || *helpLong {
		commands.Help()
		return
	}

	if *enable != "" {
		commands.Enable(*enable)
		return
	}

	if *disable != "" {
		commands.Disable(*disable)
		return
	}

	if *ctl != "" {
		commands.Ctl(*ctl)
		return
	}

	args := flag.Args()
	if len(args) > 0 {
		command := args[0]

		if newCmd, ok := dshfunc.CustomCmds()[command]; ok {
			args[0] = newCmd 
			command = newCmd
			fmt.Println("Executing custom command:", command, args)
		}

		switch command {
		case "ld":
			commands.Ld(args[1:])
		case "mkfile":
			commands.Mkfile(args[1:])
		default:
			fmt.Println("Unknown command:", command)
		}
	}
}