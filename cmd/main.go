package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/tobiashort/clap-go"

	"github.com/tobiashort/i3-rename-workspace/dmenu"
	"github.com/tobiashort/i3-rename-workspace/i3"
)

type Args struct {
	DmenuArgs string `clap:"short=,long=dmenu,default-value='-p Rename:',description='Arguments passed to dmenu'"`
}

var workspaces map[int]i3.Workspace

func main() {
	args := Args{}
	clap.Parse(&args)
	dmenuArgs := strings.Split(args.DmenuArgs, " ")

	workspaces = i3.GetWorkspaces()

	i3.OnWorkspaceChange(
		// init
		func(w i3.Workspace) {
			if workspace, ok := workspaces[w.Num]; ok {
				i3.RenameWorkspace(w.Name, workspace.Name)
			} else {
				workspaces[w.Num] = w
			}
		},
		// focus
		func(w i3.Workspace) {
			for num, workspace := range workspaces {
				if num == w.Num {
					workspace.Focused = true
				} else {
					workspace.Focused = false
				}
				workspaces[num] = workspace
			}
		},
	)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGUSR1)

	for range signals {
		name := dmenu.Run(dmenuArgs)
		name = strings.TrimSpace(name)
		for _, workspace := range workspaces {
			if workspace.Focused {
				if name == "" {
					name = fmt.Sprintf("%d", workspace.Num)
				} else {
					name = fmt.Sprintf("%d:%s", workspace.Num, name)
				}
				i3.RenameWorkspace(workspace.Name, name)
				workspace.Name = name
				workspaces[workspace.Num] = workspace
			}
		}
	}
}
