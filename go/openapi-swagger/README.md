Example OpenAPI REST API Server
================================

Details:

* Date: January 2026
* OpenAPI version: 3.0.0
  * https://www.openapis.org/
* Code Generator: https://openapi-codegen 
  * Version: v2 (https://github.com/oapi-codegen/oapi-codegen)
* Editor: "OpenAPI (Swagger) Editor"
  * https://marketplace.visualstudio.com/items?itemName=42Crunch.vscode-openapi
* Web Framework: [Gin](https://gin-gonic.com/en/)
  * https://pkg.go.dev/github.com/gin-gonic/gin

This is an attempt at a canonical example using the above tools to generate a
REST API (server) in Go.  This specific tool combination (oapi-codegen + Gin) is
not unique, but rather simply the choice of this example.

Some items of note:

OpenAPI version 3.0
-------------------

The OpenAPI specification at the time of this writing has matured to v3.2.0.
There are schema elements present in v3.2 that are not supported in 3.0.  The
tooling chosen all supports v3.0, but not necessarily newer versions, at this
time.  -- _caveat emptor_.

"go tool" support
-----------------

Starting in Go 1.24, a new method of dependency management has been added to
`go.mod` and the language:  "Tool dependencies", (see:
https://tip.golang.org/doc/modules/managing-dependencies#tools).  The use of
`openapi-codegen` (v2) recommends using 'go tools' as a best practice.  The
reader should reference the above link, and links in the README of the
`oapi-codegen` project for more details.

Specifically, at the bottom of this project's `go.mod` file, you will see a line
specifying the `oapi-codegen` tool.  This is in conjunction with the
`//go:generate ...` lines at the beginning of the `api/api.go` file.

`openai-codegen/v2` and `github.com/deepmap/oapi-codegen`
---------------------------------------------------------

At the time this example was created, `oapi-codegen` v2 was released.
Specifically, the code generation tool was released as v2.  The generated code,
and other necessary dependencies do not actually import packages from
https://github.com/openapi-codegen/oapi-codegen, but rather from other packages
within the same organization.  -- _caveat emptor_.

Additionally, the `openapi-codegen` project was previously housed in the
`depmap` GitHub organization.  This lead to some confusion.  See:
https://github.com/oapi-codegen/oapi-codegen/discussions/1605 for further
context.  Be aware that older documentation and AI code generation may confuse
the two locations.

Generating Golang
-----------------

The `api/api.go` file contains `//go:generate ...` lines at the top of the file.
These lines are used by the `go generate` command to create the necessary Go
code from the OpenAPI specification file in the root of the project,
`/openapi.yaml`.

Based on the PetShop example from the `openapi-codegen` project, the generated
code is placed in two separate files; this is simply this developers preference.
As a consiquence, there are two `//go:generate ...` lines in the `api/api.go`
file.

A reminder as to how the golang tools work:  normal go compilation does not
trigger code generation.  Code generation may be triggered either by running the
`go generate` command, or by manually running the underlying tool, in this case
the `oapi-codegen` command.  The `//go:generate ...` lines in the source code
serve double duty as both documentation and as instructions to the `go generate`
command.

URL Path Prefixing
------------------

This example intentionally deployed the "API" to a path prefixed with `/api`.
This lead to further understanding of the OpenAPI specification.  Specifically,
the `servers:` element in the OpenAPI specification is used to specify the base
URL for the API.  If this is not specified, the middleware that validates the
request will reject correct paths because they do not match the anticipated path
based on the specification.  The specification covers this topic in section 4.2
"Server Object" (https://spec.openapis.org/oas/v3.0.3#server-object).  I'm sure
there are other ways to address this.

<br>
<em>...that's all folks!</em>
