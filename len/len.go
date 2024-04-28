package main

import (
	"bufio"
	"os"
)

func helpText() string {
	return "arg -> string; string whose length is to be found\n" +
		"stdin -> string; length of lines sent to stdin"
}

func main() {
	fileInfo, _ := os.Stdin.Stat()
	if fileInfo.Mode()&os.ModeCharDevice == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			println(len(scanner.Text()))
		}
	} else {
		args := os.Args
		if len(args) == 2 {
			println(len(args[1]))
		} else {

			println(helpText())
		}
	}
}
