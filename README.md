Binds $mon+n in order to rename a current workspace.

```
~/.config/i3/config

set $mod Mod1
exec --no-startup-id i3-rename-workspace
bindsym $mod+n exec --no-startup-id killall -SIGUSR1 i3-rename-workspace
```
