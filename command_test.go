package krait

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aqua777/krait/testing"
)

func TestCommand(t *testing.T) {
	testing.Run(t, new(testCommand))
}

type testCommand struct {
	testing.Suite
	expectedError error
}

func (suite *testCommand) SetupTest() {
	Reset() // Reset global state before each test
	suite.expectedError = fmt.Errorf("test error")
}

// verifyExpectedError checks that the error matches the expected error and includes the context of what failed
func (suite *testCommand) verifyExpectedError(err error, context string) {
	suite.Error(err, "Execute should return error when "+context)
	suite.Equal(suite.expectedError, err, "Error should match the "+context+" error")
}

func (suite *testCommand) TestNewCommand() {
	cmd := New("app", "test app", "test app long description")
	suite.NotNil(cmd.cmd)
	suite.NotNil(cmd.viper)
	suite.NotNil(cmd.argsViper)
	suite.Nil(cmd.SanityCheck)
	suite.Nil(cmd.BeforeRun)
	suite.Nil(cmd.AfterRun)
	suite.Nil(cmd.Run)
}

func (suite *testCommand) TestCommandWithoutRunnerShouldFail() {
	suite.WithCustom(nil, nil, func() {
		cmd := New("app", "test app", "test app long description").WithBeforeRun(func(args []string) error { return nil })
		err := cmd.Execute()
		suite.Error(err, "Execute should return error when runner is not defined")
		suite.Contains(err.Error(), "cannot be executed", "Error message should indicate command cannot be executed")
	})
}

func (suite *testCommand) TestCommandWithNilRunnerShouldFail() {
	suite.WithCustom(nil, nil, func() {
		cmd := New("app", "test app", "test app long description").WithBeforeRun(func(args []string) error { return nil }).WithRun(nil)
		err := cmd.Execute()
		suite.Error(err, "Execute should return error when runner is not defined")
		suite.Contains(err.Error(), "run function is nil", "Error message should indicate that run function is nil")
	})
}

func (suite *testCommand) TestCommandWithoutRunnerAndWithSubCommandsShouldDisplayUsage() {
	suite.WithCustom(nil, nil, func() {
		// Create a parent command without a runner but with subcommands
		parentCmd := New("parent", "parent command", "parent command description")

		// Add a subcommand
		subCmd := New("sub", "sub command", "sub command description").
			WithRun(func(args []string) error {
				return nil
			})
		parentCmd.WithCommand(subCmd)

		// Capture usage output
		var buf bytes.Buffer
		oldUsageOutput := usageOutput
		usageOutput = &buf
		defer func() { usageOutput = oldUsageOutput }()

		// Execute should not return error
		err := parentCmd.Execute()
		suite.NoError(err, "Execute should not return error when command has subcommands")

		// Verify usage was displayed
		output := buf.String()
		suite.Contains(output, "Usage:", "Usage output should be displayed")
		suite.Contains(output, "parent [command]", "Usage should show parent command syntax")
		suite.Contains(output, "sub", "Usage should list subcommand")
	})
}

func (suite *testCommand) TestSanityCheckError() {
	suite.WithCustom(nil, nil, func() {
		suite.verifyExpectedError(
			newTestCommand().WithSanityCheck(func() error { return suite.expectedError }).Execute(),
			"sanity check fails")
	})
}

func (suite *testCommand) TestBeforeRunError() {
	suite.WithCustom(nil, nil, func() {
		suite.verifyExpectedError(
			newTestCommand().WithBeforeRun(func(args []string) error { return suite.expectedError }).Execute(),
			"BeforeRun fails")
	})
}

func (suite *testCommand) TestAfterRunError() {
	suite.WithCustom(nil, nil, func() {
		suite.verifyExpectedError(
			newTestCommand().WithAfterRun(func(args []string) error { return suite.expectedError }).Execute(),
			"AfterRun fails")
	})
}

func (suite *testCommand) TestRunError() {
	suite.WithCustom(nil, nil, func() {
		suite.verifyExpectedError(
			New("app", "test app", "test app long description").WithRun(func(args []string) error { return suite.expectedError }).Execute(),
			"Run fails")
	})
}

func (suite *testCommand) TestUsage() {
	suite.WithCustom(nil, nil, func() {
		// Create a command with some parameters
		cmd := newTestCommand().
			WithStringP("config", "Config file path", "config", "c", "CONFIG_FILE", "default-config").
			WithBoolP("verbose", "Enable verbose output", "verbose", "v", "VERBOSE", false)

		// Temporarily replace usageOutput to capture the output
		var buf bytes.Buffer
		oldUsageOutput := usageOutput
		usageOutput = &buf
		defer func() { usageOutput = oldUsageOutput }()

		// Call Usage() function
		cmd.Usage()

		// Get the captured output
		output := buf.String()

		// Verify the output contains expected information
		suite.Contains(output, "Usage:", "Usage output should contain usage header")
		suite.Contains(output, "app [flags]", "Usage output should contain command usage")
		suite.Contains(output, "--config", "Usage output should contain config flag")
		suite.Contains(output, "-c", "Usage output should contain config shorthand")
		suite.Contains(output, "--verbose", "Usage output should contain verbose flag")
		suite.Contains(output, "-v", "Usage output should contain verbose shorthand")
	})
}

func (suite *testCommand) TestCurrentCommandEqualsRootCommand() {
	suite.WithCustom(nil, nil, func() {
		App("my-app", "", "").WithRun(func(args []string) error {
			return nil
		})
		if suite.NoError(Execute()) {
			root := Root()
			current := Current()
			suite.Same(root, current, "Current command should be the same as root command")
		}
	})
}

func (suite *testCommand) TestCurrentCommandEqualsSubCommand() {
	suite.WithCustom(nil, []string{"sub2"}, func() {
		sub1 := New("sub1", "", "").WithRun(func(args []string) error {
			fmt.Println("running: sub1")
			return nil
		})
		sub2 := New("sub2", "", "").WithRun(func(args []string) error {
			fmt.Println("running: sub2")
			return nil
		})
		sub3 := New("sub3", "", "").WithRun(func(args []string) error {
			fmt.Println("running: sub3")
			return nil
		})
		root := App("my-app", "", "").WithRun(func(args []string) error {
			fmt.Println("running: my-app")
			return nil
		}).WithCommand(sub1).WithCommand(sub2).WithCommand(sub3)

		if suite.NoError(Execute()) {
			suite.Same(root, Root(), "Root command should be the same as root command")
			suite.Same(sub2, Current(), "Current command should be the same as sub2 command")
			suite.NotSame(Root(), Current(), "Current command should not be the same as root command")
		}
	})
}

func (suite *testCommand) TestIsDebugNegative() {
	suite.WithCustom(nil, nil, func() {
		cmd := newTestCommand()
		if suite.NoError(cmd.Execute()) {
			suite.False(cmd.IsDebug(), "IsDebug should return false when debug is not enabled")
		}
	})
}

func (suite *testCommand) TestIsDebugPositiveShort() {
	suite.WithCustom(nil, []string{"-d"}, func() {
		cmd := newTestCommand().WithDebug("debug", "d", "APP_DEBUG")
		if suite.NoError(cmd.Execute()) {
			suite.True(cmd.IsDebug(), "IsDebug should return true when debug is enabled")
		}
	})
}

func (suite *testCommand) TestIsDebugPositiveLong() {
	suite.WithCustom(nil, []string{"--debug"}, func() {
		cmd := newTestCommand().WithDebug("debug", "", "APP_DEBUG")
		if suite.NoError(cmd.Execute()) {
			suite.True(cmd.IsDebug(), "IsDebug should return true when debug is enabled")
		}
	})
}

func (suite *testCommand) TestAllAndArgsSettings() {
	for _, d := range []struct {
		env   map[string]string
		flags []string
	}{
		{
			env: map[string]string{
				"TEST_STRING":   "string",
				"TEST_INT":      "123",
				"TEST_BOOL":     "false",
				"TEST_DURATION": "11h22m33s"},
			flags: []string{},
		}, {
			env: map[string]string{},
			flags: []string{
				"--test-string", "string",
				"--test-int", "123",
				"--test-bool=false",
				"--test-duration", "11h22m33s",
			},
		},
	} {
		suite.WithCustom(d.env, d.flags, func() {
			cmd := newTestCommand().
				WithString("string-param", "test description", "test-string", "TEST_STRING", "default-string").
				WithInt("int-param", "test description", "test-int", "TEST_INT", 42).
				WithBool("bool-param", "test description", "test-bool", "TEST_BOOL", true).
				WithDuration("duration-param", "test description", "test-duration", "TEST_DURATION", 5*time.Minute)
			err := cmd.Execute()
			suite.NoError(err, "Execute should not return error")

			allSettings := cmd.AllSettings()
			fmt.Println("AllSettings:", allSettings)
			suite.Contains(allSettings, "string-param", "String parameter should be in AllSettings")
			suite.Contains(allSettings, "int-param", "Int parameter should be in AllSettings")
			suite.Contains(allSettings, "bool-param", "Bool parameter should be in AllSettings")
			suite.Contains(allSettings, "duration-param", "Duration parameter should be in AllSettings")

			json := cmd.AllSettingsAsJson()
			fmt.Println("JSON:", json)
			suite.Contains(json, "\"string-param\":", "String parameter should be in JSON")
			suite.Contains(json, "\"int-param\":", "Int parameter should be in JSON")
			suite.Contains(json, "\"bool-param\":", "Bool parameter should be in JSON")
			suite.Contains(json, "\"duration-param\":", "Duration parameter should be in JSON")
		})

		suite.WithCustom(d.env, d.flags, func() {
			var s string
			var i int
			var b bool
			var duration time.Duration

			cmd := newTestCommand().
				WithStringVar(&s, "test description", "test-string", "TEST_STRING", "default-string").
				WithIntVar(&i, "test description", "test-int", "TEST_INT", 42).
				WithBoolVar(&b, "test description", "test-bool", "TEST_BOOL", true).
				WithDurationVar(&duration, "test description", "test-duration", "TEST_DURATION", 5*time.Minute)
			err := cmd.Execute()
			suite.NoError(err, "Execute should not return error")

			argsSettings := cmd.ArgsSettings()
			fmt.Println("ArgsSettings:", argsSettings)
			suite.Contains(argsSettings, "test-string", "String parameter should be in ArgsSettings")
			suite.Contains(argsSettings, "test-int", "Int parameter should be in ArgsSettings")
			suite.Contains(argsSettings, "test-bool", "Bool parameter should be in ArgsSettings")
			suite.Contains(argsSettings, "test-duration", "Duration parameter should be in ArgsSettings")

			json := cmd.ArgsSettingsAsJson()
			fmt.Println("JSON:", json)
			suite.Contains(json, "\"test-string\":", "String parameter should be in JSON")
			suite.Contains(json, "\"test-int\":", "Int parameter should be in JSON")
			suite.Contains(json, "\"test-bool\":", "Bool parameter should be in JSON")
			suite.Contains(json, "\"test-duration\":", "Duration parameter should be in JSON")
		})
	}
}

func (suite *testCommand) TestWithParamsWhenNil() {
	suite.WithCustom(nil, nil, func() {
		cmd := newTestCommand().WithParams(nil)
		suite.NoError(cmd.Execute(), "Execute should not return error when params is nil")
	})
}

func (suite *testCommand) TestWithParamsWhenNotNil() {
	suite.WithCustom(nil, []string{"--test-string", "test-value"}, func() {
		params := NewConfigParams().With("test.string", "test-string", "t", "TEST_STRING", "This is test string", "default-value", nil)
		cmd := newTestCommand().WithParams(params)
		suite.NoError(cmd.Execute(), "Execute should not return error when params is not nil")
		suite.Equal("test-value", GetString("test.string"), "String parameter should be set to test-value")
	})
}
