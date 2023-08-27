package dmenu

import (
	"fmt"
	"os"
	"os/exec"
)

func Run(args []string) string {
	cmd := exec.Command("dmenu", args...)
	data, err := cmd.CombinedOutput()
	out := string(data)
	if err != nil {
    if out != "" {
		  fmt.Println(out)
    }
    fmt.Fprintf(os.Stderr, "error: %s\n", err)
    return ""
	}
	return out
}
