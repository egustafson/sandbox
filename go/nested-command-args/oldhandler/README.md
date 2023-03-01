Original:  Nested Command(s) with Arguments
-------------------------------------------

> This is the original prototype work done to demonstrate how to combine the
> `flag` (or `pflag`) packages as part of an internal CLI processing engine.
> This work is superseded by the work(s) in its parent directory.

> (original README.md) below:

This example attempts to use the flag (or pflag) package in
combination with arguments that represent commands and sub-commands to
process flags at each level of the sub-command.

Structure of initial command line:

    rootCmd -rootArg1 -rootArg2 subCmd -subArg1 ... subSubCmd -subSubArg1

