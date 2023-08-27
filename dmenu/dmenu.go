package dmenu

import (
	"fmt"
	"os/exec"
  "strings"
)

func Prompt(args string) string {
	cmd := exec.Command("dmenu", strings.Split(args, " ")...)
	data, err := cmd.CombinedOutput()
	out := string(data)
	if err != nil {
		fmt.Println(out)
		panic(err)
	}
	return out
}
