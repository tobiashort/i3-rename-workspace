package dmenu

import (
	"fmt"
	"os"
	"os/exec"
)

func Prompt(args []string) string {
	cmd := exec.Command("dmenu", args...)
	data, err := cmd.CombinedOutput()
	out := string(data)
	if err != nil {
		fmt.Println(out)
    fmt.Fprintf(os.Stderr, "error: %s", err)
    return ""
	}
	return out
}
