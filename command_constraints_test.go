package krait

import (
	"testing"

	kraittest "github.com/aqua777/krait/testing"
)

func TestCommandConstraints(t *testing.T) {
	kraittest.Run(t, new(testCommandConstraints))
}

type testCommandConstraints struct {
	kraittest.Suite
}

func (suite *testCommandConstraints) SetupTest() {
	Reset()
}

// WithRequired

func (suite *testCommandConstraints) TestWithRequiredErrorWhenFlagAbsent() {
	suite.WithCustom(nil, nil, func() {
		cmd := newTestCommand().
			WithString("app.host", "Hostname", "host", "APP_HOST").
			WithRequired("host")

		err := cmd.Execute()
		suite.Error(err, "Execute must return an error when a required flag is absent")
	})
}

func (suite *testCommandConstraints) TestWithRequiredSucceedsWhenFlagProvided() {
	suite.WithCustom(nil, []string{"--host", "example.com"}, func() {
		cmd := newTestCommand().
			WithString("app.host", "Hostname", "host", "APP_HOST").
			WithRequired("host")

		suite.NoError(cmd.Execute(), "Execute must succeed when a required flag is provided")
	})
}

func (suite *testCommandConstraints) TestWithRequiredPanicsOnUnknownFlag() {
	suite.WithCustom(nil, nil, func() {
		suite.Panics(func() {
			newTestCommand().WithRequired("nonexistent-flag")
		}, "WithRequired must panic for an unregistered flag")
	})
}

func (suite *testCommandConstraints) TestWithRequiredChaining() {
	suite.WithCustom(nil, nil, func() {
		cmd := newTestCommand().
			WithString("app.host", "Hostname", "host", "APP_HOST").
			WithString("app.port", "Port", "port", "APP_PORT")
		result := cmd.WithRequired("host")
		suite.Equal(cmd, result, "WithRequired must return the command for chaining")
	})
}

// WithFlagsRequiredTogether

func (suite *testCommandConstraints) TestWithFlagsRequiredTogetherNoneSucceeds() {
	suite.WithCustom(nil, nil, func() {
		cmd := newTestCommand().
			WithString("app.user", "User", "user", "APP_USER").
			WithString("app.pass", "Pass", "pass", "APP_PASS").
			WithFlagsRequiredTogether("user", "pass")

		suite.NoError(cmd.Execute(), "Execute must succeed when none of the required-together flags are provided")
	})
}

func (suite *testCommandConstraints) TestWithFlagsRequiredTogetherAllSucceeds() {
	suite.WithCustom(nil, []string{"--user", "alice", "--pass", "secret"}, func() {
		cmd := newTestCommand().
			WithString("app.user", "User", "user", "APP_USER").
			WithString("app.pass", "Pass", "pass", "APP_PASS").
			WithFlagsRequiredTogether("user", "pass")

		suite.NoError(cmd.Execute(), "Execute must succeed when all required-together flags are provided")
	})
}

func (suite *testCommandConstraints) TestWithFlagsRequiredTogetherPartialErrors() {
	suite.WithCustom(nil, []string{"--user", "alice"}, func() {
		cmd := newTestCommand().
			WithString("app.user", "User", "user", "APP_USER").
			WithString("app.pass", "Pass", "pass", "APP_PASS").
			WithFlagsRequiredTogether("user", "pass")

		err := cmd.Execute()
		suite.Error(err, "Execute must return an error when only some required-together flags are provided")
	})
}

func (suite *testCommandConstraints) TestWithFlagsRequiredTogetherPanicsOnUnknownFlag() {
	suite.WithCustom(nil, nil, func() {
		suite.Panics(func() {
			newTestCommand().
				WithString("app.user", "User", "user", "APP_USER").
				WithFlagsRequiredTogether("user", "nonexistent-flag")
		}, "WithFlagsRequiredTogether must panic when any flag is not registered")
	})
}

// WithFlagsOneRequired

func (suite *testCommandConstraints) TestWithFlagsOneRequiredNoneErrors() {
	suite.WithCustom(nil, nil, func() {
		cmd := newTestCommand().
			WithString("app.src", "Source", "src", "APP_SRC").
			WithString("app.dst", "Destination", "dst", "APP_DST").
			WithFlagsOneRequired("src", "dst")

		err := cmd.Execute()
		suite.Error(err, "Execute must return an error when none of the one-required flags are provided")
	})
}

func (suite *testCommandConstraints) TestWithFlagsOneRequiredOneSucceeds() {
	suite.WithCustom(nil, []string{"--src", "/tmp/a"}, func() {
		cmd := newTestCommand().
			WithString("app.src", "Source", "src", "APP_SRC").
			WithString("app.dst", "Destination", "dst", "APP_DST").
			WithFlagsOneRequired("src", "dst")

		suite.NoError(cmd.Execute(), "Execute must succeed when one of the one-required flags is provided")
	})
}

func (suite *testCommandConstraints) TestWithFlagsOneRequiredAllSucceeds() {
	suite.WithCustom(nil, []string{"--src", "/tmp/a", "--dst", "/tmp/b"}, func() {
		cmd := newTestCommand().
			WithString("app.src", "Source", "src", "APP_SRC").
			WithString("app.dst", "Destination", "dst", "APP_DST").
			WithFlagsOneRequired("src", "dst")

		suite.NoError(cmd.Execute(), "Execute must succeed when all one-required flags are provided")
	})
}

func (suite *testCommandConstraints) TestWithFlagsOneRequiredPanicsOnUnknownFlag() {
	suite.WithCustom(nil, nil, func() {
		suite.Panics(func() {
			newTestCommand().
				WithString("app.src", "Source", "src", "APP_SRC").
				WithFlagsOneRequired("src", "nonexistent-flag")
		}, "WithFlagsOneRequired must panic when any flag is not registered")
	})
}

// WithFlagsMutuallyExclusive

func (suite *testCommandConstraints) TestWithFlagsMutuallyExclusiveNoneSucceeds() {
	suite.WithCustom(nil, nil, func() {
		cmd := newTestCommand().
			WithBool("app.verbose", "Verbose", "verbose", "APP_VERBOSE").
			WithBool("app.quiet", "Quiet", "quiet", "APP_QUIET").
			WithFlagsMutuallyExclusive("verbose", "quiet")

		suite.NoError(cmd.Execute(), "Execute must succeed when none of the mutually-exclusive flags are provided")
	})
}

func (suite *testCommandConstraints) TestWithFlagsMutuallyExclusiveOneSucceeds() {
	suite.WithCustom(nil, []string{"--verbose"}, func() {
		cmd := newTestCommand().
			WithBool("app.verbose", "Verbose", "verbose", "APP_VERBOSE").
			WithBool("app.quiet", "Quiet", "quiet", "APP_QUIET").
			WithFlagsMutuallyExclusive("verbose", "quiet")

		suite.NoError(cmd.Execute(), "Execute must succeed when one mutually-exclusive flag is provided")
	})
}

func (suite *testCommandConstraints) TestWithFlagsMutuallyExclusiveBothErrors() {
	suite.WithCustom(nil, []string{"--verbose", "--quiet"}, func() {
		cmd := newTestCommand().
			WithBool("app.verbose", "Verbose", "verbose", "APP_VERBOSE").
			WithBool("app.quiet", "Quiet", "quiet", "APP_QUIET").
			WithFlagsMutuallyExclusive("verbose", "quiet")

		err := cmd.Execute()
		suite.Error(err, "Execute must return an error when both mutually-exclusive flags are provided")
	})
}

func (suite *testCommandConstraints) TestWithFlagsMutuallyExclusivePanicsOnUnknownFlag() {
	suite.WithCustom(nil, nil, func() {
		suite.Panics(func() {
			newTestCommand().
				WithBool("app.verbose", "Verbose", "verbose", "APP_VERBOSE").
				WithFlagsMutuallyExclusive("verbose", "nonexistent-flag")
		}, "WithFlagsMutuallyExclusive must panic when any flag is not registered")
	})
}

// cmd.Unmarshal

func (suite *testCommandConstraints) TestCmdUnmarshalReceivesFlagValue() {
	suite.WithCustom(nil, []string{"--host", "example.com"}, func() {
		type Config struct {
			App struct {
				Host string `mapstructure:"host"`
			} `mapstructure:"app"`
		}

		var got Config
		cmd := newTestCommand().
			WithString("app.host", "Hostname", "host", "APP_HOST", "localhost").
			WithRun(func(args []string) error {
				return Current().Unmarshal(&got)
			})

		suite.NoError(cmd.Execute())
		suite.Equal("example.com", got.App.Host, "Unmarshal must decode CLI flag value into struct")
	})
}

func (suite *testCommandConstraints) TestCmdUnmarshalReceivesEnvValue() {
	suite.WithCustom(map[string]string{"APP_HOST": "env.example.com"}, nil, func() {
		type Config struct {
			App struct {
				Host string `mapstructure:"host"`
			} `mapstructure:"app"`
		}

		var got Config
		cmd := newTestCommand().
			WithString("app.host", "Hostname", "host", "APP_HOST", "localhost").
			WithRun(func(args []string) error {
				return Current().Unmarshal(&got)
			})

		suite.NoError(cmd.Execute())
		suite.Equal("env.example.com", got.App.Host, "Unmarshal must decode env var value into struct")
	})
}

func (suite *testCommandConstraints) TestCmdUnmarshalUnknownKeyDecodesToZero() {
	suite.WithCustom(nil, nil, func() {
		type Config struct {
			NoSuchKey string `mapstructure:"no_such_key"`
		}

		var got Config
		cmd := newTestCommand().
			WithRun(func(args []string) error {
				return Current().Unmarshal(&got)
			})

		suite.NoError(cmd.Execute())
		suite.Equal("", got.NoSuchKey, "Unmarshal must leave unknown keys at zero value")
	})
}

// krait.Unmarshal

func (suite *testCommandConstraints) TestPackageLevelUnmarshalAfterExecute() {
	suite.WithCustom(nil, []string{"--host", "pkg.example.com"}, func() {
		type Config struct {
			App struct {
				Host string `mapstructure:"host"`
			} `mapstructure:"app"`
		}

		var got Config
		cmd := newTestCommand().
			WithString("app.host", "Hostname", "host", "APP_HOST", "localhost").
			WithRun(func(args []string) error {
				return Unmarshal(&got)
			})

		suite.NoError(cmd.Execute())
		suite.Equal("pkg.example.com", got.App.Host, "Package-level Unmarshal must decode current command's values")
	})
}

func (suite *testCommandConstraints) TestPackageLevelUnmarshalBeforeExecuteNoError() {
	suite.WithCustom(nil, nil, func() {
		type Config struct {
			App struct {
				Host string `mapstructure:"host"`
			} `mapstructure:"app"`
		}

		var got Config
		err := Unmarshal(&got)
		suite.NoError(err, "Package-level Unmarshal must not panic or error when called before Execute()")
		suite.Equal("", got.App.Host, "Package-level Unmarshal before Execute must decode zero values")
	})
}
