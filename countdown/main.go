package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func secondsToHMS(s int) (int, int, int) {
	h := s / (60 * 60)
	s = s % (60 * 60)
	m := s / 60
	s = s % 60
	return h, m, s
}

func main() {
	args := os.Args
	secondsLeft := 0
	if len(args) < 2 {
		println("too few arguments")
		return
	}
	if len(args) > 4 {
		println("too many arguments")
		return
	}
	for i := 1; i < len(args); i++ {
		arg := args[i]
		switch arg[len(arg)-1] {
		case 'H':
			fallthrough
		case 'h':
			hours, err := strconv.Atoi(arg[:len(arg)-1])
			if err != nil {
				println("error in cmdline arg", arg)
			}
			secondsLeft += hours * 60 * 60
		case 'M':
			fallthrough
		case 'm':
			mins, err := strconv.Atoi(arg[:len(arg)-1])
			if err != nil {
				println("error in cmdline arg", arg)
			}
			secondsLeft += mins * 60
		case 'S':
			fallthrough
		case 's':
			secs, err := strconv.Atoi(arg[:len(arg)-1])
			if err != nil {
				println("error in cmdline arg", arg)
			}
			secondsLeft += secs
		}
	}
	for ; secondsLeft >= 0; secondsLeft-- {
		h, m, s := secondsToHMS(secondsLeft)
		fmt.Printf("\r            ")
		fmt.Printf("\r%dH %dM %dS", h, m, s)
		time.Sleep(time.Second)
	}
	fmt.Println()
}
