package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetPrompt() string {
	// 현재 디렉토리
	dir, _ := os.Getwd()
	home := os.Getenv("HOME")
	dir = strings.Replace(dir, home, "~", 1)

	// Git 브랜치 확인
	branch := getGitBranch()

	// 프롬프트 조합
	if branch != "" {
		return fmt.Sprintf("mysh %s (%s) > ", dir, branch)
	}
	return fmt.Sprintf("mysh %s > ", dir)
}

func getGitBranch() string {
	out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}
