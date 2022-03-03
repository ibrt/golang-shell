package shell

import (
	"fmt"
	"io"
	"strings"

	"github.com/codeskyblue/go-sh"
	"github.com/ibrt/golang-errors/errorz"
)

// Logf describes a function that echoes the details of the command being run.
type Logf func(string, ...interface{})

var (
	// DefaultLogf is the default implementation of Logf.
	DefaultLogf = func(cmd string, params ...interface{}) {
		fmt.Println(append([]interface{}{cmd}, params...))
	}
)

// Command describes a command to be spawned in a shell.
type Command struct {
	cmd    string
	params []interface{}
	logf   Logf
	env    map[string]string
	dir    string
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
}

// NewCommand creates a new Command.
func NewCommand(cmd string, params ...interface{}) *Command {
	return &Command{
		cmd:    cmd,
		params: params,
		logf:   DefaultLogf,
		env:    make(map[string]string),
	}
}

// AddParams appends the given params to the command.
func (c *Command) AddParams(params ...interface{}) *Command {
	c.params = append(c.params, params...)
	return c
}

// AddParamsString appends the given params to the command.
func (c *Command) AddParamsString(params ...string) *Command {
	for _, param := range params {
		c.params = append(c.params, param)
	}
	return c
}

// SetLogf overrides DefaultLogf for this command. Set to nil to skip logging.
func (c *Command) SetLogf(logf Logf) *Command {
	c.logf = logf
	return c
}

// SetDir sets the command working directory.
func (c *Command) SetDir(dir string) *Command {
	c.dir = dir
	return c
}

// SetStdin sets the command standard input.
func (c *Command) SetStdin(stdin io.Reader) *Command {
	c.stdin = stdin
	return c
}

// SetStdout sets the command standard output.
func (c *Command) SetStdout(stdout io.Writer) *Command {
	c.stdout = stdout
	return c
}

// SetStderr sets the command standard error.
func (c *Command) SetStderr(stderr io.Writer) *Command {
	c.stderr = stderr
	return c
}

// SetEnv sets an environment variable.
func (c *Command) SetEnv(key, value string) *Command {
	c.env[key] = value
	return c
}

// SetEnvMap sets a map of environment variable.
func (c *Command) SetEnvMap(m map[string]string) *Command {
	for k, v := range m {
		c.env[k] = v
	}
	return c
}

// Run runs the command.
func (c *Command) Run() error {
	return errorz.MaybeWrap(c.toSH().Run(), errorz.Skip())
}

// MustRun runs the commands, panics on error.
func (c *Command) MustRun() {
	errorz.MaybeMustWrap(c.Run(), errorz.Skip())
}

// Output runs the command and returns its combined output as string.
func (c *Command) Output() (string, error) {
	rawOutput, err := c.toSH().Output()
	if err != nil {
		return "", errorz.Wrap(err, errorz.Skip())
	}
	return strings.TrimSpace(string(rawOutput)), nil
}

// MustOutput is like output but panics on error.
func (c *Command) MustOutput() string {
	output, err := c.Output()
	errorz.MaybeMustWrap(err, errorz.Skip())
	return output
}

func (c *Command) toSH() *sh.Session {
	shl := sh.NewSession()
	shl.ShowCMD = false

	if c.dir != "" {
		shl.SetDir(c.dir)
	}

	if c.stdin != nil {
		shl.Stdin = c.stdin
	}

	if c.stdout != nil {
		shl.Stdout = c.stdout
	}

	if c.stderr != nil {
		shl.Stderr = c.stderr
	}

	for k, v := range c.env {
		shl.SetEnv(k, v)
	}

	if c.logf != nil {
		c.logf(c.cmd, c.params...)
	}

	return shl.Command(c.cmd, c.params...)
}
