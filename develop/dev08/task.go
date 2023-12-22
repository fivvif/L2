package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Input \"exit\" to end")
	for {

		fmt.Printf("$ ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		command := scanner.Text()
		switch {
		case strings.HasPrefix(command, "cd"):
			args := strings.Fields(command)
			if len(args) != 2 {
				fmt.Println("Usage: cd <directory>")
				continue
			}
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Println(err)
				continue
			}
		case command == "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(dir)
		case strings.HasPrefix(command, "echo"):
			args := strings.Fields(command)[1:]
			fmt.Println(strings.Join(args, " "))

		case command == "exit":
			os.Exit(0)
		default:
			cmd := exec.Command("/bin/sh", "-c", command)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}
		}

	}

}
