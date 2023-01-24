package dmenu

import (
	"os"
	"os/exec"
	"strings"

	"github.com/t-hg/i3-rename-workspace/must"
)

func GetString(prompt string) string {
	cmd := exec.Command("dmenu", "-p", prompt)
	devNull := must.Do2(os.Open("/dev/null"))
	defer devNull.Close()
	cmd.Stdin = devNull
	data, err := cmd.CombinedOutput()
	out := strings.TrimSpace(string(data))
	if err != nil {
		panic(out)
	}
	return out
}
