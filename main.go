package main

import (
	"bufio"
	"fmt"
	"minishell/cmd"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// 컨텍스트 인식 프롬프트
		fmt.Print(cmd.GetPrompt())

		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		command := parts[0]
		args := parts[1:]

		// 내장 명령어 먼저 확인
		if cmd.RunBuiltin(command, args) {
			continue
		}

		// 외부 명령어 실행
		c := exec.Command(command, args...)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		if err := c.Run(); err != nil {
			fmt.Printf("mysh: %s: 명령어를 찾을 수 없습니다\n", command)
		}
	}
}
