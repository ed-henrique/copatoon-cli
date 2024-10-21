package views

import (
	"os"
	"os/exec"
	"runtime"
)

func clearTerminal() {
	var (
		cmd  = ""
		args = make([]string, 0)
	)
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "cls"}
	default:
		cmd = "clear"
	}
	run := exec.Command(cmd, args...)
	run.Stdout = os.Stdout
	run.Run()
}
