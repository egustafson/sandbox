package command

import (
	"fmt"
	"strings"
	"unicode"
)

type CommandRegistry interface {
	Register(ch *CommandHandler) error
	Lookup(cmdLine string) (ch *CommandHandler, ok bool)
}

var _ CommandRegistry = &CommandRepo{}

type CommandHandler struct {
	Use     string
	Handler string // a placeholder for a handler function
}

type CommandRepo struct {
	CommandMap map[string]*CommandHandler
}

func NewCommandRegistry() *CommandRepo {
	return &CommandRepo{
		CommandMap: make(map[string]*CommandHandler),
	}
}

func (repo *CommandRepo) Register(ch *CommandHandler) error {
	cmdSeq, _ := parseCmdLine(ch.Use)
	cmd := strings.Join(cmdSeq, " ") // rejoin the command sequence with a single [space] rune
	if _, ok := repo.CommandMap[cmd]; ok {
		return duplicateCommandHandlerError(cmd)
	}
	repo.CommandMap[cmd] = ch
	return nil
}

func (repo *CommandRepo) Lookup(cmdLine string) (ch *CommandHandler, ok bool) {
	cmd, args := parseCmdLine(cmdLine)
	for len(cmd) > 0 {
		cmdStr := strings.Join(cmd, " ") // rejoin with a single [space] rune
		ch, ok := repo.CommandMap[cmdStr]
		if ok {
			return ch, true // <-- eventually more sophistocated, return args, parse args, ....
		}
		args = append( // push
			[]string{cmd[len(cmd)-1]}, // last element of cmd (as a slice)
			args...,                   // all of args
		)
		cmd = cmd[:len(cmd)-1] // remove the last element, it went to args
	}
	// nothing found
	return nil, false
}

func parseCmdLine(cmdLine string) (cmd, args []string) {
	argList := strings.Fields(cmdLine) // split on whitespace
	split := len(argList)
	for i, a := range argList {
		r := []rune(a)
		if !unicode.IsLetter(r[0]) {
			split = i
			break
		}
	}
	cmd = argList[0:split]
	args = argList[split:]
	return cmd, args
}

// ----------  Errors  ----------

type DuplicateCommandHandlerError interface {
	error
}

func duplicateCommandHandlerError(command string) DuplicateCommandHandlerError {
	e := fmt.Errorf("duplicate command handler: %s", command)
	return DuplicateCommandHandlerError(e)
}
