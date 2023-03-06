Embedded CLI Processing with nested 'Commands'
==============================================
<!-- VSCode:  ctrl-shift-v to preview markdown -->

> Original example archived in directory `oldhandler`

## Goals

An updated example -- this example attempts to prototype the following:

* Input is a single string with space separated commands, sub-commands, and
  arguments as part of each "command stage".
* Argument flags are prefixed with '+' instead of '--' and the '+' is translated
  into '--', or '-' for simple flags so that the Golang `flag` or `pflag`
  packages can be utilized internally
* "Handler" functions can register with a central command registry.  (This
  registry is then used to dispatch the execution of the handler when a CLI
  string arrives.)
* Nested commands register as "parent-cmd:child-cmd" using ':' as the separator.
  This allows sub-command command names to be reused across multiple parent
  commands.
* The entire framework is encapsulated as a golang package.
* Built-in Help and support for Help when coding with this package.
* Handle the corner case of a CLI with "cmd-a cmd-b cmd-c flag-1 ..." with out
  requiring intermediate command handlers for 'cmd-a' and 'cmd-b'.

Note:  heavy inspiration was taken from the spf13/cobra package.  Initially
Cobra was considered as a solution.  

## Usage and Examples

see: example_test.go