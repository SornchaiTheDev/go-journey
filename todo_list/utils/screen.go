package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {

	// Clearing the screen using ANSI escape codes
	// Thanks Chat-GPT

	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
