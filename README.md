
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

// Check if file or path exist, returns bool 
es.FileExists(string)

// Generate random (int) long string
es.StringRNG(int)

// Join paths together
es.Join("path", "path2, "path_etc")
```

### Thanks readme.so for this readme!!
