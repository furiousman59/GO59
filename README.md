
# execShell

Easy-To-Use Golang "os/exec" wrapper


## Installation


```bash
go get github.com/furiousman59/execShell
```

```go
import es "github.com/furiousman59/execShell"
```
## Usage

```go
// Optionally: var execShell = es.ExecShell
es.ExecShell("command", bool, "args", "arg2", "arg etc")

// The module also has a function called Strip, used to strip output of blank lines, used by ExecShell function
fmt.Println(es.Strip(output-of-something))
```
If bool is false, print command output, otherwise, don't.

### Thanks readme.so for this readme!!
