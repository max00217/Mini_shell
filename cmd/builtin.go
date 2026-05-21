package cmd

import (
	"fmt"
	"os"
)

func RunBuiltin(cmd string, args []string) bool {
	switch cmd {
	case "cd":
		return builtinCd(args)
	case "pwd":
		return builtinPwd()
	case "exit":
		fmt.Println("종료합니다.")
		os.Exit(0)
	}
	return false
}

func builtinCd(args []string) bool {
	if len(args) == 0 {
		home := os.Getenv("HOME")
		os.Chdir(home)
	} else {
		err := os.Chdir(args[0])
		if err != nil {
			fmt.Printf("cd: %s: 그런 파일이나 디렉터리가 없습니다\n", args[0])
		}
	}
	return true
}

func builtinPwd() bool {
	dir, _ := os.Getwd()
	fmt.Println(dir)
	return true
}
