package cmdopt

import (
	"os"
)

// CommandOpt command opt
type CommandOpt struct {
	Command string
	Args    []string
}

// SubCommands sub command
var SubCommands []string = []string{"current", "list", "down", "up"}

// ParseArgs parse args
func ParseArgs() (opt CommandOpt) {
	if len(os.Args) < 2 {
		opt.Command = "help"
		return
	}
	subcmd := os.Args[1]
	for _, sub := range SubCommands {
		if subcmd == sub {
			opt.Command = sub
			if len(os.Args) > 2 {
				opt.Args = append(opt.Args, os.Args[2:]...)
			}
			return
		}
	}
	opt.Command = "help"
	return
}

// Run run command
func (co *CommandOpt) Run() {

	if co.Command == "help" {
		ho := &HelpOpt{}
		ho.Run()
	} else if co.Command == "current" {
		o := &CurrentOpt{}
		o.Run()
	} else if co.Command == "list" {
		o := &ListOpt{}
		o.Run()
	} else if co.Command == "down" {
		o := &DownOpt{Args: co.Args}
		o.Run()
	} else if co.Command == "up" {
		o := &UpOpt{Args: co.Args}
		o.Run()
	}
}
