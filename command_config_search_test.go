package krait

import (
	"errors"
	"os"
	"path/filepath"
	stdtesting "testing"

	kraittesting "github.com/aqua777/krait/testing"
	"github.com/spf13/viper"
)

type testCommandConfigSearch struct {
	kraittesting.Suite
}

func TestCommandConfigSearch(t *stdtesting.T) {
	kraittesting.Run(t, new(testCommandConfigSearch))
}

func (s *testCommandConfigSearch) SetupTest() {
	Reset()
}

func TestReadViperConfigFromSearchPaths(t *stdtesting.T) {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "demo.yaml")
	if err := os.WriteFile(path, []byte("greeting: hello\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	v := viper.New()
	if err := readViperConfigFromSearchPaths(v, "demo", []string{dir}); err != nil {
		t.Fatal(err)
	}
	if got := v.GetString("greeting"); got != "hello" {
		t.Fatalf("greeting: got %q want hello", got)
	}
}

func (s *testCommandConfigSearch) TestSearchPathSinglePathLoadsValues() {
	dir := s.T().TempDir()
	cfgPath := filepath.Join(dir, "myapp.yaml")
	s.Require().NoError(os.WriteFile(cfgPath, []byte("app:\n  name: FromSearch\n"), 0o600))

	cmd := New("t", "d", "")
	cmd.WithConfigName("myapp").WithConfigPath(dir)
	cmd.WithString("app.name", "n", "name", "", "default")
	cmd.WithRun(func([]string) error { return nil })

	s.NoError(cmd.Execute())
	s.Equal("FromSearch", GetString("app.name"))
}

func (s *testCommandConfigSearch) TestSearchPathFirstMatchingPathWins() {
	dir1 := s.T().TempDir()
	dir2 := s.T().TempDir()
	s.Require().NoError(os.WriteFile(filepath.Join(dir1, "myapp.yaml"), []byte("app:\n  name: First\n"), 0o600))
	s.Require().NoError(os.WriteFile(filepath.Join(dir2, "myapp.yaml"), []byte("app:\n  name: Second\n"), 0o600))

	cmd := New("t", "d", "")
	cmd.WithConfigName("myapp").WithConfigPath(dir1).WithConfigPath(dir2)
	cmd.WithString("app.name", "n", "name", "", "default")
	cmd.WithRun(func([]string) error { return nil })

	s.NoError(cmd.Execute())
	s.Equal("First", GetString("app.name"))
}

func (s *testCommandConfigSearch) TestSearchPathNoFileReturnsBeforeRun() {
	dir := s.T().TempDir()
	ran := false
	cmd := New("t", "d", "")
	cmd.WithConfigName("missing").WithConfigPath(dir)
	cmd.WithString("app.name", "n", "name", "", "default")
	cmd.WithRun(func([]string) error {
		ran = true
		return nil
	})

	err := cmd.Execute()
	s.Error(err)
	s.False(ran)
	var notFound viper.ConfigFileNotFoundError
	s.True(errors.As(err, &notFound), "want ConfigFileNotFoundError, got %v", err)
}

func (s *testCommandConfigSearch) TestWithConfigOverridesSearchPath() {
	dir := s.T().TempDir()
	s.Require().NoError(os.WriteFile(filepath.Join(dir, "myapp.yaml"), []byte("app:\n  name: FromSearch\n"), 0o600))

	cmd := New("t", "d", "")
	cmd.WithConfig("examples/app.yaml", "config", "", "TEST_CFG_SEARCH")
	cmd.WithConfigName("myapp").WithConfigPath(dir)
	cmd.WithString("app.name", "n", "name", "APP_NAME", "default")
	cmd.WithRun(func([]string) error { return nil })

	s.NoError(cmd.Execute())
	s.Equal("MyApplication", GetString("app.name"))
}

func (s *testCommandConfigSearch) TestEnvOverridesDiscoveredConfig() {
	dir := s.T().TempDir()
	s.Require().NoError(os.WriteFile(filepath.Join(dir, "myapp.yaml"), []byte("app:\n  name: FromFile\n"), 0o600))

	env := map[string]string{"APP_NAME": "FromEnv"}
	s.WithCustom(env, nil, func() {
		cmd := New("t", "d", "")
		cmd.WithConfigName("myapp").WithConfigPath(dir)
		cmd.WithString("app.name", "n", "name", "APP_NAME", "default")
		cmd.WithRun(func([]string) error { return nil })
		s.NoError(cmd.Execute())
		s.Equal("FromEnv", GetString("app.name"))
	})
}

func (s *testCommandConfigSearch) TestFlagOverridesDiscoveredConfig() {
	dir := s.T().TempDir()
	s.Require().NoError(os.WriteFile(filepath.Join(dir, "myapp.yaml"), []byte("app:\n  name: FromFile\n"), 0o600))

	args := []string{"--name", "FromFlag"}
	s.WithCustom(nil, args, func() {
		cmd := New("t", "d", "")
		cmd.WithConfigName("myapp").WithConfigPath(dir)
		cmd.WithString("app.name", "n", "name", "APP_NAME", "default")
		cmd.WithRun(func([]string) error { return nil })
		s.NoError(cmd.Execute())
		s.Equal("FromFlag", GetString("app.name"))
	})
}
