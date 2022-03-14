package shellz_test

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ibrt/golang-bites/filez"
	"github.com/ibrt/golang-fixtures/fixturez"
	"github.com/stretchr/testify/require"

	"github.com/ibrt/golang-shell/shellz"
)

func TestCommand_Output(t *testing.T) {
	c := fixturez.CaptureOutput()
	defer c.Close()

	out, err := shellz.NewCommand("cat", "-b").
		SetStdin(strings.NewReader("input")).
		Output()
	fixturez.RequireNoError(t, err)
	require.Contains(t, out, "1\tinput")
	require.Contains(t, c.GetOutString(), "cat")
	require.Contains(t, c.GetOutString(), "-b")
}

func TestCommand_Run(t *testing.T) {
	c := fixturez.CaptureOutput()
	defer c.Close()

	err := shellz.NewCommand("cat").
		AddParams("-b").
		SetStdin(strings.NewReader("input")).
		Run()
	fixturez.RequireNoError(t, err)
	require.Contains(t, c.GetOutString(), "1\tinput")
	require.Contains(t, c.GetOutString(), "cat")
	require.Contains(t, c.GetOutString(), "-b")
}

func TestCommand_MustRun(t *testing.T) {
	c := fixturez.CaptureOutput()
	defer c.Close()

	require.Panics(t, func() {
		shellz.NewCommand("cat", "unknown.txt").MustRun()
	})
}

func TestCommand_MustOutput(t *testing.T) {
	c := fixturez.CaptureOutput()
	defer c.Close()

	require.Panics(t, func() {
		shellz.NewCommand("cat", "unknown.text").MustOutput()
	})
}

func TestCommand_HideCmd(t *testing.T) {
	c := fixturez.CaptureOutput()
	defer c.Close()

	err := shellz.NewCommand("cat").
		AddParamsString("-b").
		SetStdin(strings.NewReader("input")).
		SetLogf(nil).
		Run()
	fixturez.RequireNoError(t, err)
	require.Contains(t, c.GetOutString(), "1\tinput")
	require.NotContains(t, c.GetOutString(), "cat")
	require.NotContains(t, c.GetOutString(), "-b")
}

func TestCommand_Stdout(t *testing.T) {
	c := fixturez.CaptureOutput()
	defer c.Close()
	out := &bytes.Buffer{}

	err := shellz.NewCommand("cat", "-b").
		SetStdin(strings.NewReader("input")).
		SetStdout(out).
		Run()
	fixturez.RequireNoError(t, err)
	require.Contains(t, out.String(), "1\tinput")
	require.Contains(t, c.GetOutString(), "cat")
	require.Contains(t, c.GetOutString(), "-b")
}

func TestCommand_Stderr(t *testing.T) {
	c := fixturez.CaptureOutput()
	defer c.Close()
	out := &bytes.Buffer{}

	err := shellz.NewCommand("cat", "unknown.txt").
		SetStderr(out).
		Run()
	require.Error(t, err)
	require.Contains(t, out.String(), "No such file or directory")
	require.Contains(t, c.GetOutString(), "cat")
	require.Contains(t, c.GetOutString(), "unknown.txt")
}

func TestCommand_SetDir(t *testing.T) {
	c := fixturez.CaptureOutput()
	defer c.Close()
	dir := filepath.Dir(filez.MustAbs("."))

	err := shellz.NewCommand("pwd").
		SetDir(dir).
		Run()
	fixturez.RequireNoError(t, err)
	require.Contains(t, c.GetOutString(), dir+"\n")
	require.Contains(t, c.GetOutString(), "pwd")
}

func TestCommand_SetEnv(t *testing.T) {
	c := fixturez.CaptureOutput()
	defer c.Close()

	err := shellz.NewCommand("env").
		SetEnv("TEST_ENV_893892", "true").
		Run()
	fixturez.RequireNoError(t, err)
	require.Contains(t, c.GetOutString(), "TEST_ENV_893892=true")
	require.Contains(t, c.GetOutString(), "env")
}

func TestCommand_SetEnvMap(t *testing.T) {
	c := fixturez.CaptureOutput()
	defer c.Close()

	err := shellz.NewCommand("env").
		SetEnvMap(map[string]string{
			"TEST_ENV_893892": "true",
			"TEST_ENV_893893": "true",
		}).
		Run()
	fixturez.RequireNoError(t, err)
	require.Contains(t, c.GetOutString(), "TEST_ENV_893892=true")
	require.Contains(t, c.GetOutString(), "TEST_ENV_893893=true")
	require.Contains(t, c.GetOutString(), "env")
}
