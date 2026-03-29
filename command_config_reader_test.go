package krait

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	stdtesting "testing"

	kraittesting "github.com/aqua777/krait/testing"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

type readerConfigError struct{}

func (readerConfigError) Read([]byte) (int, error) {
	return 0, fmt.Errorf("read failed")
}

type testCommandConfigReader struct {
	kraittesting.Suite
}

func TestCommandConfigReader(t *stdtesting.T) {
	kraittesting.Run(t, new(testCommandConfigReader))
}

func TestReadNamedVipersFromReader(t *stdtesting.T) {
	t.Helper()
	v1, v2 := viper.New(), viper.New()
	err := readNamedVipersFromReader(v1, v2, strings.NewReader("greeting: hello\n"), "yaml")
	require.NoError(t, err)
	require.Equal(t, "hello", v1.GetString("greeting"))
	require.Equal(t, "hello", v2.GetString("greeting"))
}

func TestReadNamedVipersFromReaderReadError(t *stdtesting.T) {
	t.Helper()
	v1, v2 := viper.New(), viper.New()
	err := readNamedVipersFromReader(v1, v2, readerConfigError{}, "yaml")
	require.Error(t, err)
}

func TestReadNamedVipersFromReaderParseError(t *stdtesting.T) {
	t.Helper()
	v1, v2 := viper.New(), viper.New()
	err := readNamedVipersFromReader(v1, v2, strings.NewReader(":\n"), "yaml")
	require.Error(t, err)
}

func (s *testCommandConfigReader) SetupTest() {
	Reset()
}

func (s *testCommandConfigReader) TestWithConfigReaderValidYAML() {
	cmd := New("t", "d", "")
	cmd.WithConfigReader(strings.NewReader("app:\n  name: FromReader\n"), "yaml")
	cmd.WithString("app.name", "n", "name", "", "default")
	cmd.WithRun(func([]string) error {
		s.Equal("FromReader", GetString("app.name"))
		return nil
	})
	s.NoError(cmd.Execute())
}

func (s *testCommandConfigReader) TestWithConfigReaderInvalidConfigType() {
	ran := false
	cmd := New("t", "d", "")
	cmd.WithConfigReader(strings.NewReader("k: v\n"), "__not_supported_type__")
	cmd.WithString("k", "k", "k", "", "")
	cmd.WithRun(func([]string) error {
		ran = true
		return nil
	})
	err := cmd.Execute()
	s.Error(err)
	s.False(ran)
	var unsupported viper.UnsupportedConfigError
	s.True(errors.As(err, &unsupported), "got %#v", err)
}

func (s *testCommandConfigReader) TestWithConfigReaderEnvOverrides() {
	env := map[string]string{"APP_NAME": "FromEnv"}
	s.WithCustom(env, nil, func() {
		cmd := New("t", "d", "")
		cmd.WithConfigReader(strings.NewReader("app:\n  name: FromReader\n"), "yaml")
		cmd.WithString("app.name", "n", "name", "APP_NAME", "default")
		cmd.WithRun(func([]string) error {
			s.Equal("FromEnv", GetString("app.name"))
			return nil
		})
		s.NoError(cmd.Execute())
	})
}

func (s *testCommandConfigReader) TestWithConfigReaderFlagOverrides() {
	args := []string{"--name", "FromFlag"}
	s.WithCustom(nil, args, func() {
		cmd := New("t", "d", "")
		cmd.WithConfigReader(strings.NewReader("app:\n  name: FromReader\n"), "yaml")
		cmd.WithString("app.name", "n", "name", "APP_NAME", "default")
		cmd.WithRun(func([]string) error {
			s.Equal("FromFlag", GetString("app.name"))
			return nil
		})
		s.NoError(cmd.Execute())
	})
}

func (s *testCommandConfigReader) TestWithConfigOverridesReader() {
	dir := s.T().TempDir()
	cfg := filepath.Join(dir, "c.yaml")
	s.Require().NoError(os.WriteFile(cfg, []byte("app:\n  name: FromFile\n"), 0o600))

	cmd := New("t", "d", "")
	cmd.WithConfig(cfg, "config", "", "")
	cmd.WithConfigReader(strings.NewReader("app:\n  name: FromReader\n"), "yaml")
	cmd.WithString("app.name", "n", "name", "", "default")
	cmd.WithRun(func([]string) error {
		s.Equal("FromFile", GetString("app.name"))
		return nil
	})
	s.NoError(cmd.Execute())
}
