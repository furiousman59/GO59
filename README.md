
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
// Optionally: var execShell = es.execShell
es.execShell("command", bool, "args", "arg2")
```
If bool is false, print command output, otherwise, don't.

### Thanks readme.so for this readme!!
