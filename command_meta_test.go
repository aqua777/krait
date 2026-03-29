package krait

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	kraittest "github.com/aqua777/krait/testing"
)

func TestCommandMeta(t *testing.T) {
	kraittest.Run(t, new(testCommandMeta))
}

type testCommandMeta struct {
	kraittest.Suite
}

func (suite *testCommandMeta) SetupTest() {
	Reset()
}

// WithHidden

func (suite *testCommandMeta) TestWithHiddenCommandNotInUsage() {
	suite.WithCustom(nil, nil, func() {
		hidden := New("secret", "hidden command", "").
			WithHidden().
			WithRun(func(args []string) error { return nil })

		var buf bytes.Buffer
		oldUsageOutput := usageOutput
		usageOutput = &buf
		defer func() { usageOutput = oldUsageOutput }()

		parent := New("app", "app", "").WithCommand(hidden)
		parent.Usage()

		suite.NotContains(buf.String(), "secret", "Hidden command must not appear in usage output")
	})
}

func (suite *testCommandMeta) TestWithHiddenCommandStillExecutes() {
	suite.WithCustom(nil, []string{"secret"}, func() {
		ran := false
		hidden := New("secret", "hidden command", "").
			WithHidden().
			WithRun(func(args []string) error {
				ran = true
				return nil
			})

		App("app", "app", "").WithCommand(hidden)
		suite.NoError(Execute())
		suite.True(ran, "Hidden command must still execute when called directly")
	})
}

// WithDeprecated

func (suite *testCommandMeta) TestWithDeprecatedCommandStillExecutes() {
	suite.WithCustom(nil, []string{"old"}, func() {
		ran := false
		deprecated := New("old", "old command", "").
			WithDeprecated("use 'new' instead").
			WithRun(func(args []string) error {
				ran = true
				return nil
			})

		App("app", "app", "").WithCommand(deprecated)
		suite.NoError(Execute())
		suite.True(ran, "Deprecated command must still execute")
	})
}

func (suite *testCommandMeta) TestWithDeprecatedFieldStoredOnCommand() {
	suite.WithCustom(nil, nil, func() {
		deprecated := New("old", "old command", "").
			WithDeprecated("use 'new' instead").
			WithRun(func(args []string) error { return nil })

		suite.Equal("use 'new' instead", deprecated.cmd.Deprecated, "Deprecated message must be stored on the underlying Cobra command")
	})
}

// WithFlagHidden

func (suite *testCommandMeta) TestWithFlagHiddenRemovedFromUsage() {
	suite.WithCustom(nil, nil, func() {
		cmd := newTestCommand().
			WithString("param", "A secret param", "secret-flag", "SECRET_FLAG").
			WithFlagHidden("secret-flag")

		usage := cmd.UsageString()
		suite.NotContains(usage, "secret-flag", "Hidden flag must not appear in usage output")
	})
}

func (suite *testCommandMeta) TestWithFlagHiddenStillFunctional() {
	suite.WithCustom(nil, []string{"--secret-flag", "hidden-value"}, func() {
		cmd := newTestCommand().
			WithString("param", "A secret param", "secret-flag", "SECRET_FLAG").
			WithFlagHidden("secret-flag")

		suite.NoError(cmd.Execute())
		suite.Equal("hidden-value", GetString("param"), "Hidden flag must still accept values")
	})
}

func (suite *testCommandMeta) TestWithFlagHiddenPanicsOnUnknownFlag() {
	suite.WithCustom(nil, nil, func() {
		suite.Panics(func() {
			newTestCommand().WithFlagHidden("nonexistent-flag")
		}, "WithFlagHidden must panic for unregistered flag")
	})
}

// WithFlagDeprecated

func (suite *testCommandMeta) TestWithFlagDeprecatedFlagIsHiddenFromUsage() {
	suite.WithCustom(nil, nil, func() {
		// Cobra's MarkDeprecated also marks the flag hidden, so it won't appear in usage.
		cmd := newTestCommand().
			WithString("param", "An old param", "old-flag", "OLD_FLAG").
			WithFlagDeprecated("old-flag", "use --new-flag instead")

		usage := cmd.UsageString()
		suite.NotContains(usage, "old-flag", "Deprecated flag is hidden by Cobra and must not appear in usage")
	})
}

func (suite *testCommandMeta) TestWithFlagDeprecatedPanicsOnUnknownFlag() {
	suite.WithCustom(nil, nil, func() {
		suite.Panics(func() {
			newTestCommand().WithFlagDeprecated("nonexistent-flag", "msg")
		}, "WithFlagDeprecated must panic for unregistered flag")
	})
}

// WithSilenceErrors

func (suite *testCommandMeta) TestWithSilenceErrorsErrorStillReturned() {
	suite.WithCustom(nil, nil, func() {
		expectedErr := fmt.Errorf("boom")
		cmd := New("app", "app", "").
			WithSilenceErrors().
			WithRun(func(args []string) error { return expectedErr })

		err := cmd.Execute()
		suite.Equal(expectedErr, err, "Error must still be returned from Execute even when silenced")
	})
}

func (suite *testCommandMeta) TestWithSilenceErrorsDoesNotWriteToStderr() {
	suite.WithCustom(nil, nil, func() {
		cmd := New("app", "app", "").
			WithSilenceErrors().
			WithRun(func(args []string) error { return fmt.Errorf("boom") })

		// Cobra writes to the command's err writer; capture it.
		var buf bytes.Buffer
		cmd.cmd.SetErr(&buf)
		_ = cmd.Execute()

		suite.Empty(buf.String(), "Silenced errors must not be written to stderr writer")
	})
}

// WithAliases

func (suite *testCommandMeta) TestWithAliasesRespondToAliasA() {
	suite.WithCustom(nil, []string{"a"}, func() {
		ran := false
		sub := New("full-name", "sub", "").
			WithAliases("a", "b").
			WithRun(func(args []string) error {
				ran = true
				return nil
			})

		App("app", "app", "").WithCommand(sub)
		suite.NoError(Execute(), "Execute via alias 'a' must succeed")
		suite.True(ran, "Command must run when called via alias 'a'")
	})
}

func (suite *testCommandMeta) TestWithAliasesRespondToAliasB() {
	suite.WithCustom(nil, []string{"b"}, func() {
		ran := false
		sub := New("full-name", "sub", "").
			WithAliases("a", "b").
			WithRun(func(args []string) error {
				ran = true
				return nil
			})

		App("app", "app", "").WithCommand(sub)
		suite.NoError(Execute(), "Execute via alias 'b' must succeed")
		suite.True(ran, "Command must run when called via alias 'b'")
	})
}

func (suite *testCommandMeta) TestWithAliasesChaining() {
	suite.WithCustom(nil, nil, func() {
		cmd := newTestCommand().WithAliases("x", "y")
		suite.Equal([]string{"x", "y"}, cmd.cmd.Aliases, "Aliases must be stored on the underlying cobra command")
	})
}

// WithVersion

func (suite *testCommandMeta) TestWithVersionPrintsVersion() {
	suite.WithCustom(nil, []string{"--version"}, func() {
		var buf strings.Builder
		cmd := New("app", "app", "").
			WithVersion("1.2.3").
			WithRun(func(args []string) error { return nil })
		cmd.cmd.SetOut(&buf)

		_ = cmd.Execute()
		suite.Contains(buf.String(), "1.2.3", "--version flag must output the version string")
	})
}

func (suite *testCommandMeta) TestWithVersionChaining() {
	suite.WithCustom(nil, nil, func() {
		cmd := newTestCommand().WithVersion("2.0.0")
		suite.Equal("2.0.0", cmd.cmd.Version, "Version must be stored on the underlying cobra command")
	})
}

// IsSet (command-level)

func (suite *testCommandMeta) TestIsSetReturnsTrueForKeyWithDefault() {
	suite.WithCustom(nil, nil, func() {
		// Viper's IsSet returns true for any key that has a value, including from SetDefault.
		cmd := newTestCommand().
			WithString("app.host", "Hostname", "host", "APP_HOST", "localhost")

		suite.NoError(cmd.Execute())
		suite.True(cmd.IsSet("app.host"), "IsSet must return true for a registered key with a default value")
	})
}

func (suite *testCommandMeta) TestIsSetReturnsTrueWhenSetViaFlag() {
	suite.WithCustom(nil, []string{"--host", "example.com"}, func() {
		cmd := newTestCommand().
			WithString("app.host", "Hostname", "host", "APP_HOST", "localhost")

		suite.NoError(cmd.Execute())
		suite.True(cmd.IsSet("app.host"), "IsSet must return true when key was set via CLI flag")
	})
}

func (suite *testCommandMeta) TestIsSetReturnsTrueWhenSetViaEnv() {
	suite.WithCustom(map[string]string{"APP_HOST": "env.example.com"}, nil, func() {
		cmd := newTestCommand().
			WithString("app.host", "Hostname", "host", "APP_HOST", "localhost")

		suite.NoError(cmd.Execute())
		suite.True(cmd.IsSet("app.host"), "IsSet must return true when key was set via environment variable")
	})
}

func (suite *testCommandMeta) TestIsSetReturnsFalseForUnknownKey() {
	suite.WithCustom(nil, nil, func() {
		cmd := newTestCommand()
		suite.NoError(cmd.Execute())
		suite.False(cmd.IsSet("nonexistent.key"), "IsSet must return false for an unknown key")
	})
}

// IsSet (package-level)

func (suite *testCommandMeta) TestPackageLevelIsSetReturnsFalseForUnknownKeyBeforeExecute() {
	suite.WithCustom(nil, nil, func() {
		// currentViper is nil before Execute; getViperInstance returns a bare Viper
		// with no keys set, so any lookup returns false.
		suite.False(IsSet("app.host"), "Package-level IsSet must return false for an unknown key before Execute()")
	})
}

func (suite *testCommandMeta) TestPackageLevelIsSetReturnsTrueAfterFlagSet() {
	suite.WithCustom(nil, []string{"--host", "example.com"}, func() {
		cmd := newTestCommand().
			WithString("app.host", "Hostname", "host", "APP_HOST", "localhost")

		suite.NoError(cmd.Execute())
		suite.True(IsSet("app.host"), "Package-level IsSet must return true when key was set via CLI flag")
	})
}

// WithEnvPrefix

func (suite *testCommandMeta) TestWithEnvPrefixMapsKeyViaAutomaticEnv() {
	suite.WithCustom(map[string]string{"MYAPP_HOST": "prefix.example.com"}, nil, func() {
		cmd := newTestCommand().
			WithEnvPrefix("myapp").
			WithString("host", "Hostname", "host", "")

		suite.NoError(cmd.Execute())
		suite.Equal("prefix.example.com", GetString("host"), "WithEnvPrefix must map MYAPP_HOST to key 'host' via AutomaticEnv")
	})
}

func (suite *testCommandMeta) TestWithEnvPrefixExplicitBindEnvStillResolves() {
	suite.WithCustom(map[string]string{"EXPLICIT_HOST": "explicit.example.com"}, nil, func() {
		cmd := newTestCommand().
			WithEnvPrefix("myapp").
			WithString("host", "Hostname", "host", "EXPLICIT_HOST")

		suite.NoError(cmd.Execute())
		suite.Equal("explicit.example.com", GetString("host"), "Explicit BindEnv env var must still resolve correctly alongside WithEnvPrefix")
	})
}

func (suite *testCommandMeta) TestWithEnvPrefixChaining() {
	suite.WithCustom(nil, nil, func() {
		cmd := newTestCommand()
		result := cmd.WithEnvPrefix("myapp")
		suite.Equal(cmd, result, "WithEnvPrefix must return the command for chaining")
	})
}

func (suite *testCommandMeta) TestWithEnvPrefixAppliedToBothViperInstances() {
	suite.WithCustom(map[string]string{"MYAPP_CONFIG": "/tmp/myapp.yaml"}, nil, func() {
		var configFile string
		cmd := New("app", "app", "").
			WithEnvPrefix("myapp").
			WithStringVar(&configFile, "Config file path", "config", "").
			WithRun(func(args []string) error { return nil })

		suite.NoError(cmd.Execute())
		suite.Equal("/tmp/myapp.yaml", configFile, "WithEnvPrefix must also apply to argsViper for var-bound params")
	})
}
