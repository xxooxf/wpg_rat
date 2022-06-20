package cmd

import (
	"fmt"
	"os/exec"
	"strings"
)

func ExecCommand(command string) string {
	result, err := exec.Command("cmd", "/c", command).Output()
	if err != nil {
		fmt.Println("执行cmd错误？", err)
		return "exec error"
	}
	return strings.TrimSpace(string(result))
}
