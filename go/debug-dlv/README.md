Demo Delve Debugger
===================

Install.  Don't do this in a project directory that has a `go.mod`
unless you want Delve added to your dependencies.

```
go get -u github.com/go-delve/delve/cmd/dlv
```

Debug:

```
> dlv debug
(dlv) break main.main
Breakpoint 1 set at 0x4aa43b for main.main() ./main.go:7
(dlv) c
> main.main() ./main.go:7 (hits goroutine(1):1 total:1) (PC: 0x4aa43b)
     2:
     3: import "fmt"
     4:
     5: var i, j int = 1, 2
     6:
=>   7: func main() {
     8:         var c, python, java = true, false, "no!"
     9:         fmt.Println(i, j, c, python, java)
    10:
    11:         fmt.Println("done.")
    12: }
```

Find a reference (like):
* https://www.jamessturtevant.com/posts/Using-the-Go-Delve-Debugger-from-the-command-line/ 

