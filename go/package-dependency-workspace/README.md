# Golang Package Dependency Introspection - Workspace


This directory houses two, separate golang packages and serves two
purposes:

1. Showcase Golang Workspace functionality
2. Showcase package inclusion introspection w/ version info.

## Workspace

This directory is a Workspace with two child packages beneath it.
Note that the dependent packages are not _required_ to be in
sub-directories of the workspace, but happen to be in this example.

## Package version introspection

<ins>Problem</ins>:  Upon including a golang package be able to
identify the version of the included package.

The `pkgverlib` and `pkgvermain` packages constitute an example where
`pkgverlib` is a separate golang package and included by `pkgvermain`.
The example shows how the version of the library package included in
the main program can be determined at run-time.
