package dmenu

import (
	"fmt"
	"os/exec"
)

func Prompt(prompt string) string {
	cmd := exec.Command("dmenu", "-p", prompt)
	data, err := cmd.CombinedOutput()
	out := string(data)
	if err != nil {
		fmt.Println(out)
		panic(err)
	}
	return out
}
