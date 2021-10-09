package replace

import (
	"errors"
	"fmt"
	"path"
	"strings"
)

type ReplaceArguments struct {
	// private members
	args     []string
	commands map[string]string
	// public members
	SourceDir  string
	SourceFile string
	TargetDir  string
	TargetFile string
	SourceExp  string
	TargetExp  string
}

//	Create instance of type Arguments
//		args : set received application arguments
//		returns: pointer of an instance
func NewReplaceArguments(args []string) (*ReplaceArguments, error) {
	a := &ReplaceArguments{args: args}
	if err := a.Validate(); err != nil {
		return nil, err
	}
	a.initialize()
	return a, nil
}

// Initialize type Arguments, build the commands in a key/value map
func (a *ReplaceArguments) initialize() {
	a.commands = make(map[string]string)
	for i := 1; i < len(a.args); i = i + 2 {
		cmd := a.args[i]
		cmd = strings.TrimPrefix(cmd, "-")
		a.commands[cmd] = a.args[i+1]
		switch cmd {
		case "s", "source":
			sVal := strings.ReplaceAll(a.commands[cmd], "\\", "/")
			if strings.Contains(sVal, "*") {
				a.SourceDir, a.SourceFile = path.Split(sVal)
			} else {
				a.SourceDir = sVal
				a.SourceFile = "*.*"
			}

		case "t", "target":
			sVal := strings.ReplaceAll(a.commands[cmd], "\\", "/")
			if strings.Contains(sVal, "*") {
				a.TargetDir, _ = path.Split(sVal)
			} else {
				a.TargetDir = sVal
			}

		case "f", "find":
			a.SourceExp = a.commands[cmd]

		case "r", "replace":
			a.TargetExp = a.commands[cmd]
		}
	}
}

// 	Validate type Arguments, check the correctnes of received command line parameters
//		returns : error in case any validation rule problem is detected or nil
func (a *ReplaceArguments) Validate() error {
	var err error
	// number of parameters
	if len(a.args) <= 4 {
		err = errors.New("Missing required parameters")
	}

	for i := 1; i < len(a.args); i = i + 2 {
		cmd := a.args[i]
		cmd = strings.Trim(cmd, " ")
		if cmd[0] != '-' {
			err = errors.New(fmt.Sprintf("Command not recognized: %s", cmd))
		}
	}

	return err
}

// Serialize to printable text
func (a *ReplaceArguments) ToString() string {
	return fmt.Sprintf("%v", a.commands) //fmt.Sprintf("%s", a.commands)
}
