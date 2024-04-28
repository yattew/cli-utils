package main

import (
	"bufio"
	"os"
)

func helpText() string {
	return "arg 1-> -echo; return the original lines with count attached at the end\n" +
		"arg 1-> -recho; return the original lines with count attached at the begining\n" +
		"arg 2-> string; string whose length is to be found\n" +
		"stdin -> string; length of lines sent to stdin"
}

func main() {
	fileInfo, _ := os.Stdin.Stat()
	args := os.Args
	if len(args) == 2 || len(args) == 3 {
		if args[1] == "-echo" {
			if fileInfo.Mode()&os.ModeCharDevice == 0 {
				scanner := bufio.NewScanner(os.Stdin)
				for scanner.Scan() {
					println(scanner.Text(), len(scanner.Text()))
				}
			} else {
				println(args[2], len(args[2]))
			}
		} else if args[1] == "-recho" {
			if fileInfo.Mode()&os.ModeCharDevice == 0 {
				scanner := bufio.NewScanner(os.Stdin)
				for scanner.Scan() {
					println(len(scanner.Text()), scanner.Text())
				}
			} else {
				println(len(args[2]), args[2])
			}
		} else {
			println(len(args[1]))
		}
	} else if fileInfo.Mode()&os.ModeCharDevice == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			println(len(scanner.Text()))
		}
	} else {
		println(helpText())
	}

}
