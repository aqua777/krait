package krait

import (
	"os"
	"path/filepath"
	stdtesting "testing"

	kraittesting "github.com/aqua777/krait/testing"
	"github.com/stretchr/testify/require"
)

type testCommandConfigWrite struct {
	kraittesting.Suite
}

func TestCommandConfigWrite(t *stdtesting.T) {
	kraittesting.Run(t, new(testCommandConfigWrite))
}

func TestValidateCommandWriteContext(t *stdtesting.T) {
	t.Helper()
	Reset()
	t.Cleanup(Reset)
	cmd := New("t", "d", "")
	require.ErrorIs(t, validateCommandWriteContext(cmd), errWriteConfigNotInLifecycle)
}

func TestWithExecutingCommandNoCurrent(t *stdtesting.T) {
	t.Helper()
	Reset()
	t.Cleanup(Reset)
	err := withExecutingCommand((*Command).WriteConfig)
	require.ErrorIs(t, err, errWriteConfigNotInLifecycle)
}

func (s *testCommandConfigWrite) SetupTest() {
	Reset()
}

func (s *testCommandConfigWrite) TestWriteConfigAsWritesYAML() {
	dir := s.T().TempDir()
	src := filepath.Join(dir, "in.yaml")
	out := filepath.Join(dir, "out.yaml")
	s.Require().NoError(os.WriteFile(src, []byte("app:\n  name: FromFile\n"), 0o600))

	cmd := New("t", "d", "")
	cmd.WithConfig(src, "config", "", "")
	cmd.WithString("app.name", "app name", "name", "", "default")
	cmd.WithRun(func([]string) error {
		s.NoError(cmd.WriteConfigAs(out))
		return nil
	})
	s.NoError(cmd.Execute())

	data, err := os.ReadFile(out)
	s.NoError(err)
	s.Contains(string(data), "app:")
	s.Contains(string(data), "FromFile")
}

func (s *testCommandConfigWrite) TestSafeWriteConfigAsExistsError() {
	dir := s.T().TempDir()
	src := filepath.Join(dir, "in.yaml")
	out := filepath.Join(dir, "out.yaml")
	s.Require().NoError(os.WriteFile(src, []byte("k: v\n"), 0o600))
	s.Require().NoError(os.WriteFile(out, []byte("existing\n"), 0o600))

	cmd := New("t", "d", "")
	cmd.WithConfig(src, "config", "", "")
	cmd.WithString("k", "k", "k", "", "x")
	cmd.WithRun(func([]string) error {
		s.Error(cmd.SafeWriteConfigAs(out))
		return nil
	})
	s.NoError(cmd.Execute())
}

func (s *testCommandConfigWrite) TestSafeWriteConfigAsSucceedsWhenMissing() {
	dir := s.T().TempDir()
	src := filepath.Join(dir, "in.yaml")
	out := filepath.Join(dir, "out.yaml")
	s.Require().NoError(os.WriteFile(src, []byte("k: v\n"), 0o600))

	cmd := New("t", "d", "")
	cmd.WithConfig(src, "config", "", "")
	cmd.WithString("k", "k", "k", "", "x")
	cmd.WithRun(func([]string) error {
		s.NoError(cmd.SafeWriteConfigAs(out))
		return nil
	})
	s.NoError(cmd.Execute())
	s.FileExists(out)
}

func (s *testCommandConfigWrite) TestWriteConfigWithWithConfig() {
	dir := s.T().TempDir()
	cfg := filepath.Join(dir, "app.yaml")
	s.Require().NoError(os.WriteFile(cfg, []byte("app:\n  name: A\n"), 0o600))

	cmd := New("t", "d", "")
	cmd.WithConfig(cfg, "config", "", "")
	cmd.WithString("app.name", "name", "name", "", "d")
	cmd.WithRun(func([]string) error {
		s.NoError(cmd.WriteConfig())
		return nil
	})
	s.NoError(cmd.Execute())
}

func (s *testCommandConfigWrite) TestSafeWriteConfigNoPathKnown() {
	cmd := New("t", "d", "")
	cmd.WithString("app.name", "name", "name", "", "only-default")
	cmd.WithRun(func([]string) error {
		s.Error(cmd.SafeWriteConfig())
		return nil
	})
	s.NoError(cmd.Execute())
}

func (s *testCommandConfigWrite) TestPackageWriteConfigAsDelegates() {
	dir := s.T().TempDir()
	src := filepath.Join(dir, "in.yaml")
	out := filepath.Join(dir, "out.yaml")
	s.Require().NoError(os.WriteFile(src, []byte("x: y\n"), 0o600))

	cmd := New("t", "d", "")
	cmd.WithConfig(src, "config", "", "")
	cmd.WithString("x", "x", "x", "", "z")
	cmd.WithRun(func([]string) error {
		s.NoError(WriteConfigAs(out))
		return nil
	})
	s.NoError(cmd.Execute())
	got, err := os.ReadFile(out)
	s.NoError(err)
	s.Contains(string(got), "y")
}

func (s *testCommandConfigWrite) TestPackageSafeWriteConfigAsDelegates() {
	dir := s.T().TempDir()
	src := filepath.Join(dir, "in.yaml")
	out := filepath.Join(dir, "new.yaml")
	s.Require().NoError(os.WriteFile(src, []byte("p: q\n"), 0o600))

	cmd := New("t", "d", "")
	cmd.WithConfig(src, "config", "", "")
	cmd.WithString("p", "p", "p", "", "z")
	cmd.WithRun(func([]string) error {
		s.NoError(SafeWriteConfigAs(out))
		return nil
	})
	s.NoError(cmd.Execute())
	data, err := os.ReadFile(out)
	s.NoError(err)
	s.Contains(string(data), "q")
}

func (s *testCommandConfigWrite) TestAllWriteMethodsBeforeExecute() {
	cmd := New("t", "d", "")
	cmd.WithString("k", "k", "k", "", "v")

	s.ErrorIs(cmd.WriteConfig(), errWriteConfigNotInLifecycle)
	s.ErrorIs(cmd.SafeWriteConfig(), errWriteConfigNotInLifecycle)
	s.ErrorIs(cmd.WriteConfigAs(s.T().TempDir()+"/x.yaml"), errWriteConfigNotInLifecycle)
	s.ErrorIs(cmd.SafeWriteConfigAs(s.T().TempDir()+"/y.yaml"), errWriteConfigNotInLifecycle)

	s.ErrorIs(WriteConfig(), errWriteConfigNotInLifecycle)
	s.ErrorIs(SafeWriteConfig(), errWriteConfigNotInLifecycle)
	s.ErrorIs(WriteConfigAs("/tmp/krait-write-test.yaml"), errWriteConfigNotInLifecycle)
	s.ErrorIs(SafeWriteConfigAs("/tmp/krait-safe-write-test.yaml"), errWriteConfigNotInLifecycle)
}
