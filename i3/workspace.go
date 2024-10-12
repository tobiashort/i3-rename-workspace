package i3

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/tobiashort/i3-rename-workspace/must"
)

type Workspace struct {
	Num     int
	Name    string
	Focused bool
}

func GetWorkspaces() map[int]Workspace {
	cmd := exec.Command("i3-msg", "-t", "get_workspaces")
	data, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(data))
		panic(err)
	}

	var workspaces []Workspace
	must.Do(json.Unmarshal(data, &workspaces))

	workspacesMap := make(map[int]Workspace)
	for _, workspace := range workspaces {
		workspacesMap[workspace.Num] = workspace
	}

	return workspacesMap
}

func OnWorkspaceChange(init func(workspace Workspace), focus func(workspace Workspace)) {
	cmd := exec.Command("i3-msg", "-t", "subscribe", "-m", "[\"workspace\"]")
	reader := must.Do2(cmd.StdoutPipe())
	buffered := bufio.NewReader(reader)

	go func() {
		for {
			data, err := buffered.ReadBytes('\n')
			if err != nil {
				panic(err)
			}
			var event struct {
				Change  string
				Current Workspace
			}
			must.Do(json.Unmarshal(data, &event))
			if event.Change == "init" {
				init(event.Current)
			} else if event.Change == "focus" {
				focus(event.Current)
			}
		}
	}()

	go func() {
		err := cmd.Run()
		panic(fmt.Sprintf("Stream closed. Error: %s", err.Error()))
	}()
}

func RenameWorkspace(fromName string, toName string) {
	cmd := exec.Command("i3-msg", "rename workspace", fromName, "to", toName)
	data, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(data))
		panic(err)
	}
}
