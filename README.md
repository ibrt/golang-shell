# golang-shell
[![Go Reference](https://pkg.go.dev/badge/github.com/ibrt/golang-shell.svg)](https://pkg.go.dev/github.com/ibrt/golang-shell)
![CI](https://github.com/ibrt/golang-shell/actions/workflows/ci.yml/badge.svg)
[![codecov](https://codecov.io/gh/ibrt/golang-shell/branch/main/graph/badge.svg?token=BQVP881F9Z)](https://codecov.io/gh/ibrt/golang-shell)

Fluent interface to spawn processes in Go.

### Basic Example

```go
package main

import (
    "fmt"

    "github.com/ibrt/golang-shell/shellz"
)

func main() {
    // Pipe "ls" output to terminal.
    shellz.NewCommand("ls", ".").MustRun()

    out := shellz.NewCommand("/usr/bin/env", "bash", "-c").
        AddParams("echo $MY_VAR").
        SetEnv("MY_VAR", "my-var").
        MustOutput()

    // Will output: "my-var".
    fmt.Println(out)
}
```

### Developers

Contributions are welcome, please check in on proposed implementation before sending a PR. You can validate your changes
using the `./test.sh` script.