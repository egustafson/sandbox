Nested Commands with Args
=========================

This example attempts to use the flag (or pflag) package in
combination with arguments that represent commands and sub-commands to
process flags at each level of the sub-command.

Structure of initial command line:

    rootCmd -rootArg1 -rootArg2 subCmd -subArg1 ... subSubCmd -subSubArg1

