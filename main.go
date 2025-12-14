package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("go")
	cmd.Args = append(cmd.Args, "build")
	cmd.Args = append(cmd.Args, "-o", "./i3-rename-workspace")
	cmd.Args = append(cmd.Args, "./cmd")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
