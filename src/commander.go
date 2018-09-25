package command

import (
	"strings"
	"fmt"
  "log"
)

func Parse(cmd string) {
	if strings.Contains(cmd, "!echo") {
		fmt.Println(strings.Split(cmd, "!echo ")[1])
	} else if strings.Contains(cmd, "!quit") {
    log.Fatal("account has sent !quit command! Shutting down GoAT...")
	} else {
		fmt.Println("Unknown command!")
	}
}
