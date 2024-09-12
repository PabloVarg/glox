package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/PabloVarg/glox/internal/sysexits"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		runPrompt()
		return
	}

	if len(args) == 2 {
		runFile(args[1])
		return
	}

	printHelp()
}

func printHelp() {
	fmt.Println("Usage: glox [file]")

	os.Exit(sysexits.EX_USAGE)
}

func runFile(file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}

	contents, err := io.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}

	run(string(contents))
}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(">> ")
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		run(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}

func run(source string) {
	log.Println(source)
}
