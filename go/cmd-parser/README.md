Command Line Parser
===================

_see: Partially Implemented_

This snippet implements a simple CLI parser that supports a tree of sub-commands.

Input:  
* a space separated sequence of tokens on a single line
* commands and sub-commands start with an alpha [a-zA-z]
* flags (tokens starting with a non-alpha, i.e. '-') are allowed.
* Registered "command" object that states it's "cli" as a single usage line.
  See detail below explaining how this usage line is parsed when registering a
  command.
* sub-command tokens take precedence over arguments or parameters.  Note that
  arguments may appear in a command line as a token that starts with an alpha
  character.

Result (i.e. output): when a command line is passed into the 'parser', the
'parser' returns the proper command object.

Usage Syntax
------------

The following is an informal, and likely incomplete description of how the
`usage` string is formed.

    command sub-command sub-sub-command <arg> [optional-arg] -flag <flag-arg>

* commands and sub-commands are always present FIRST in the command line
* all commands and sub-commands must be present to match
* longest sequence of sub-commands match.
  * a sub-command takes precedence over an argument value matching a shorter
    sub-command.
* arguments are matched in order and assigned to a symbol matching the string in
  the argument.  Argument names must start with an alpha
  * `<argfoo>` will be parsed into symbol `argfoo`
* optional arguments must occur *after* (required) arguments.
* flags may be interspersed and begin with one or two dashes
* optional flag-args (`-f [flag-arg]`) are populated in a greedy manner, except
  if the `flag-arg` begins with a dash.
  * optional flag-args where the command line token begins with a `-` will be
    interpreted as the beginning of a new flag and a nul optional value.

Partially Implemented
---------------------

The specification above is reasonably complete.  The implementation as of this
commit is only _partially complete_.

* Command matching works even with parameters that could be in conflict.
* Argument (parameter) and Flag population is *NOT* implemented.
* No structures for holding args or flags exist
* The return value of Lookup() should be something that has parsed args, flags,
  and other details of the cmd-line parsing: TBD

