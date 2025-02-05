package commands

import "fmt"

func Help() {
	fmt.Println("Duckshell - a very flexible shell.")
	fmt.Println("Available commands:")
	fmt.Println("  ld [options] [path] - list directory contents")
	fmt.Println("  dsh --ver - display Duckshell version")
	fmt.Println("  dsh --sysinfo - display system information")
	fmt.Println("  dsh -h|--help - display this help message")
	fmt.Println("  dsh --enable|--disable [function] - enable or disable a function")
	fmt.Println("  dsh --ctl [function] [setting:value] - control a function's settings")
	fmt.Println("  mkfile [path] - create an empty file")
}
