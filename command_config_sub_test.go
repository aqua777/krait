package krait

import (
	"os"
	"path/filepath"
	stdtesting "testing"

	kraittesting "github.com/aqua777/krait/testing"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

type testCommandConfigSub struct {
	kraittesting.Suite
}

func TestCommandConfigSub(t *stdtesting.T) {
	kraittesting.Run(t, new(testCommandConfigSub))
}

func TestViperNamedSubMap(t *stdtesting.T) {
	t.Helper()
	v := viper.New()
	require.Nil(t, viperNamedSubMap(v, ""))
	require.Nil(t, viperNamedSubMap(v, "missing"))
	v.Set("db.host", "localhost")
	v.Set("db.port", 5432)
	m := viperNamedSubMap(v, "db")
	require.NotNil(t, m)
	require.Equal(t, "localhost", m["host"])
	require.Equal(t, 5432, m["port"])
	v2 := viper.New()
	v2.Set("db", map[string]any{})
	require.Nil(t, viperNamedSubMap(v2, "db"))
	v3 := viper.New()
	v3.Set("db", "scalar")
	require.Nil(t, viperNamedSubMap(v3, "db"))
	v4 := viper.New()
	v4.Set("db.host", "h")
	v4.Set("db.tls.enabled", true)
	m4 := viperNamedSubMap(v4, "db")
	require.NotNil(t, m4)
	tls, ok := m4["tls"].(map[string]any)
	require.True(t, ok)
	require.Equal(t, true, tls["enabled"])
}

func (s *testCommandConfigSub) SetupTest() {
	Reset()
}

func (s *testCommandConfigSub) TestSubDbHostAndPort() {
	dir := s.T().TempDir()
	cfg := filepath.Join(dir, "c.yaml")
	s.Require().NoError(os.WriteFile(cfg, []byte("db:\n  host: fromfile\n  port: 5432\n"), 0o600))

	cmd := New("t", "d", "")
	cmd.WithConfig(cfg, "config", "", "")
	cmd.WithString("db.host", "h", "db-host", "", "")
	cmd.WithString("db.port", "p", "db-port", "", "")
	cmd.WithString("cache.host", "c", "cache-host", "", "")
	cmd.WithRun(func([]string) error {
		sub := cmd.Sub("db")
		s.Require().NotNil(sub)
		s.Equal("fromfile", sub["host"])
		s.Equal(5432, sub["port"])
		return nil
	})
	s.NoError(cmd.Execute())
}

func (s *testCommandConfigSub) TestSubNonexistentNil() {
	cmd := New("t", "d", "")
	cmd.WithString("x", "x", "x", "", "")
	cmd.WithRun(func([]string) error {
		s.Nil(cmd.Sub("nonexistent"))
		return nil
	})
	s.NoError(cmd.Execute())
}

func (s *testCommandConfigSub) TestSubEmptyKeyNil() {
	cmd := New("t", "d", "")
	cmd.WithString("db.host", "h", "db-host", "", "")
	cmd.WithRun(func([]string) error {
		s.Nil(cmd.Sub(""))
		s.Nil(Sub(""))
		return nil
	})
	s.NoError(cmd.Execute())
}

func (s *testCommandConfigSub) TestPackageSubDelegates() {
	dir := s.T().TempDir()
	cfg := filepath.Join(dir, "c.yaml")
	s.Require().NoError(os.WriteFile(cfg, []byte("db:\n  host: h\n  port: 9\n"), 0o600))

	cmd := New("t", "d", "")
	cmd.WithConfig(cfg, "config", "", "")
	cmd.WithString("db.host", "h", "db-host", "", "")
	cmd.WithString("db.port", "p", "db-port", "", "")
	cmd.WithRun(func([]string) error {
		pkg := Sub("db")
		s.Require().NotNil(pkg)
		s.Equal(cmd.Sub("db"), pkg)
		return nil
	})
	s.NoError(cmd.Execute())
}

func (s *testCommandConfigSub) TestSubBeforeExecuteNil() {
	Reset()
	s.T().Cleanup(Reset)
	cmd := New("t", "d", "")
	cmd.WithString("db.host", "h", "db-host", "", "")
	s.Nil(cmd.Sub("db"))
	s.Nil(Sub("db"))
}

func (s *testCommandConfigSub) TestSubCLIPriorityOverFile() {
	oldArgs := os.Args
	s.T().Cleanup(func() { os.Args = oldArgs })
	os.Args = []string{"krait-sub-test", "--db-host", "fromcli"}

	dir := s.T().TempDir()
	cfg := filepath.Join(dir, "c.yaml")
	s.Require().NoError(os.WriteFile(cfg, []byte("db:\n  host: fromfile\n  port: 1\n"), 0o600))

	cmd := New("t", "d", "")
	cmd.WithConfig(cfg, "config", "", "")
	cmd.WithString("db.host", "h", "db-host", "", "")
	cmd.WithString("db.port", "p", "db-port", "", "")
	cmd.WithRun(func([]string) error {
		sub := Sub("db")
		s.Require().NotNil(sub)
		s.Equal("fromcli", sub["host"])
		s.EqualValues(1, sub["port"])
		return nil
	})
	s.NoError(cmd.Execute())
}
