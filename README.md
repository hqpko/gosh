# gosh

顺手自用，他人不推荐使用

### example

```go
package main

import (
	"os"
	"path/filepath"

	"github.com/hqpko/gosh"
)

func main() {
	gosh.Run("ls -l")

	s := gosh.NewSession()
	gopath := os.Getenv("GOPATH")
	s.SetDir(filepath.Join(gopath, "src"))
	s.SetEvn("GOPATH", gopath)
	s.Run("pwd", "cd ..", "pwd")
}
```
