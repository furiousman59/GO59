
# GO59 Library 

My personal Golang Library with all sorts of features.


## Installation


```bash
go get github.com/furiousman59/GO59
```

```go
import es "github.com/furiousman59/GO59"
```
## Usage

```go
// Optionally: var execShell = es.ExecShell
es.ExecShell("command", bool, "args", "arg2", "arg etc")
// If bool is false, print command output, otherwise, don't.

// Strip a string of blank lines
es.Strip(string)
// You'd probally want to use this in conjuction with Println or a var, so 
// var output = es.Strip(String)
// fmt.Println(output) 
// A EXAMPLE

// Check if file exist, returns bool 
es.FileExists(string)
```

### Thanks readme.so for this readme!!
