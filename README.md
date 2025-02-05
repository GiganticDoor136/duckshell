# Duckshell (or dsh)

This is a flexible shell written in Rust, with Golang modules available.

## Description

Duckshell is a project that aims to create a convenient, flexible and functional shell that allows you to modify it however you want, we focus on flexibility. It provides the ability to extend functionality with modules written in Golang.

## Features

* Support for basic shell commands.
* Ability to plug in Golang modules to extend functionality.
* Commands `ld`, `dsh --ver`, `dsh --sysinfo`, `dsh --help`, `dsh --enable/--disable`, `dsh --ctl` and `mkfile`.
* Customization of prompt with `dsh --ctl customixe`.
* Ability to add custom commands with `customcmds`.
* Ability to change the module to your own by simply replacing the dsh.go file with your own (don't forget to rename your module, to dsh.go).

## Assembly

1.  Install Rust and Cargo: [https://www.rust-lang.org/tools/install](https://www.rust-lang.org/tools/install)
2.  Install Go: [https://go.dev/doc/install](https://go.dev/doc/install)
3.  Clone the repository: `git clone <repository link>`.
4.  Navigate to the project directory: `cd duckshell`.
5.  Build the project: `make build`.

## Usage

Run the executable file: `./target/debug/duckshell`.

## Example

```bash
$ duckshell
> ld -l /path/to/directory
> dsh --sysinfo
> mkfile /path/to/new/file.txt
> dsh --enable customixe
> dsh --ctl customixe ASCII:1,prompt-style_lean:1
> mkdir /path/to/new/directory 