package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func parse(line string) (string, []string) {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return "", []string{}
	}
	return parts[0], parts[1:]
}

func execute(cmd string, args []string) {
	switch cmd {
	case "cd":
		if len(args) == 0 {
			os.Chdir(os.Getenv("HOME"))
		} else {
			err := os.Chdir(args[0])
			if err != nil {
				fmt.Println("cd:", err)
			}
		}
	case "pwd":
		dir, _ := os.Getwd()
		fmt.Println(dir)
	case "exit":
		os.Exit(0)
	default:
		// 외부 명령어 실행 (PATH 자동 탐색)
		c := exec.Command(cmd, args...)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		err := c.Run()
		if err != nil {
			fmt.Printf("mysh: %s: 명령어를 찾을 수 없습니다\n", cmd)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("mysh> ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		cmd, args := parse(line)
		execute(cmd, args)
	}
}
