package krait

import (
	"fmt"
	"time"

	"github.com/aqua777/krait/testing"
)

// testKrait defines the test suite for krait package
type testKrait struct {
	testing.Suite
}

// SetupSuite runs before all tests in the suite
func (s *testKrait) SetupSuite() {
}

// TearDownSuite runs after all tests in the suite
func (s *testKrait) TearDownSuite() {
	Reset()
}

// SetupTest runs before each test
func (s *testKrait) SetupTest() {
	Reset()
}

// TearDownTest runs after each test
func (s *testKrait) TearDownTest() {
}

// TestKrait runs the test suite
func TestKrait(t *testing.T) {
	testing.Run(t, new(testKrait))
}

// TestGetViperInstance_NoCommand tests getViperInstance when no command is set
func (s *testKrait) TestGetViperInstance_NoCommand() {
	// When no currentViper is set, should return a non-nil bare viper instance
	v1 := getViperInstance()
	s.NotNil(v1, "Should return a non-nil viper instance")

	// Test default values for all Get* functions
	s.Equal("", GetString("nonexistent"), "GetString should return empty string for nonexistent key")
	s.Equal(false, GetBool("nonexistent"), "GetBool should return false for nonexistent key")
	s.Equal(0, GetInt("nonexistent"), "GetInt should return 0 for nonexistent key")
	s.Equal(int32(0), GetInt32("nonexistent"), "GetInt32 should return 0 for nonexistent key")
	s.Equal(int64(0), GetInt64("nonexistent"), "GetInt64 should return 0 for nonexistent key")
	s.Equal(uint(0), GetUint("nonexistent"), "GetUint should return 0 for nonexistent key")
	s.Equal(uint16(0), GetUint16("nonexistent"), "GetUint16 should return 0 for nonexistent key")
	s.Equal(uint32(0), GetUint32("nonexistent"), "GetUint32 should return 0 for nonexistent key")
	s.Equal(uint64(0), GetUint64("nonexistent"), "GetUint64 should return 0 for nonexistent key")
	s.Equal(float64(0), GetFloat64("nonexistent"), "GetFloat64 should return 0 for nonexistent key")
	s.Nil(Get("nonexistent"), "Get should return nil for nonexistent key")
}

// TestGetViperInstance_WithCurrentViper tests getViperInstance when currentViper is set
func (s *testKrait) TestGetViperInstance_WithCurrentViper() {
	cmd := New("test", "test description", "test long description")
	currentViper = cmd.viper

	v := getViperInstance()
	s.NotNil(v, "Should return a non-nil viper instance")
	s.Same(cmd.viper, v, "Should return the pinned currentViper")
}

// TestGetViperInstance_AfterReset tests getViperInstance behavior after Reset
func (s *testKrait) TestGetViperInstance_AfterReset() {
	cmd := New("test", "test description", "test long description")
	currentViper = cmd.viper

	v1 := getViperInstance()
	s.Same(cmd.viper, v1, "Should return command viper before reset")

	Reset()

	v2 := getViperInstance()
	s.NotNil(v2, "Should return a non-nil viper instance after reset")
	s.NotSame(v1, v2, "Should return a different viper instance after reset:\nv1=%p;\nv2=%p;", v1, v2)
}

func (s *testKrait) TestReset() {
	s.Nil(currentViper, "Should be nil")
	currentViper = newViper()
	s.NotNil(currentViper, "Should not be nil")
	Reset()
	s.Nil(currentViper, "Should be nil")
}

// TestAllSettings tests various parameter types and their values
func (s *testKrait) TestAllSettings() {
	cmd := newTestCommand().
		WithString("string-param", "test description", "test-string", "TEST_STRING", "default-string").
		WithInt("int-param", "test description", "test-int", "TEST_INT", 42).
		WithBool("bool-param", "test description", "test-bool", "TEST_BOOL", true).
		WithDuration("duration-param", "test description", "test-duration", "TEST_DURATION", 5*time.Minute)

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
		s.WithCustom(d.env, d.flags, func() {
			err := cmd.Execute()
			s.NoError(err, "Execute should not return error")

			allSettings := AllSettings()
			s.Contains(allSettings, "string-param", "String parameter should be in AllSettings")
			s.Contains(allSettings, "int-param", "Int parameter should be in AllSettings")
			s.Contains(allSettings, "bool-param", "Bool parameter should be in AllSettings")
			s.Contains(allSettings, "duration-param", "Duration parameter should be in AllSettings")

			json := AsJson()
			s.Contains(json, "\"string-param\":", "String parameter should be in JSON")
			s.Contains(json, "\"int-param\":", "Int parameter should be in JSON")
			s.Contains(json, "\"bool-param\":", "Bool parameter should be in JSON")
			s.Contains(json, "\"duration-param\":", "Duration parameter should be in JSON")
		})
	}
}

// TestApp tests the App function for creating a root command
func (s *testKrait) TestApp() {
	// Test creating a new root command
	app := App("test-app", "test description", "test long description")
	s.NotNil(app, "App should return a non-nil command")
	s.Equal("test-app", app.Name, "App name should match")
	s.Equal("test description", app.Description, "App description should match")
	s.Equal("test long description", app.LongDescription, "App long description should match")
	s.Same(app, rootCommand, "App should be set as root command")

	// Test that creating another root command panics
	s.Panics(func() {
		App("another-app", "desc", "long desc")
	}, "Creating another root command should panic")
}

// TestRoot tests the Root function for getting the root command
func (s *testKrait) TestRoot() {
	// When no root command is set
	s.Nil(Root(), "Root should return nil when no root command is set")

	// When root command is set
	app := App("test-app", "test description", "test long description")
	s.Same(app, Root(), "Root should return the root command")
}

// TestCurrent tests the Current function for getting the current command
func (s *testKrait) TestCurrent() {
	// When no command is set
	s.Nil(Current(), "Current should return nil when no command is set")

	// When a command is set
	cmd := New("test", "test description", "test long description")
	currentCommand = cmd
	s.Same(cmd, Current(), "Current should return the current command")

	// After reset
	Reset()
	s.Nil(Current(), "Current should return nil after reset")
}

// TestWithUsageOnError tests that usage is displayed when an error occurs if WithUsageOnError is set
func (s *testKrait) TestWithUsageOnError() {
	cmd := newTestCommand()
	s.True(cmd.cmd.SilenceUsage, "SilenceUsage should be true")

	cmd = newTestCommand().WithUsageOnError()
	s.False(cmd.cmd.SilenceUsage, "SilenceUsage should be false")
}

// TestExecute_Success tests the Execute function with a valid root command
func (s *testKrait) TestExecute_Success() {
	s.WithCustom(nil, nil, func() {
		// Create a root command that succeeds
		App("test-app", "test description", "test long description").
			WithRun(func(args []string) error {
				return nil
			})

		// Execute should not panic or return error since we catch it
		s.NotPanics(func() {
			Execute()
		}, "Execute should not panic with valid root command")
	},
	)
}

// TestExecute_NoRoot tests the Execute function when no root command is set
func (s *testKrait) TestExecute_NoRoot() {
	s.WithCustom(nil, nil, func() {
		// Execute should panic since there's no root command
		s.Panics(func() {
			Execute()
		}, "Execute should panic when no root command is set")
	},
	)
}

// TestExecute_RootError tests the Execute function when root command returns error
func (s *testKrait) TestExecute_RootError() {
	s.WithCustom(nil, nil, func() {
		// Create a root command that returns an error
		expectedErr := fmt.Errorf("test error")
		App("test-app", "test description", "test long description").
			WithRun(func(args []string) error {
				return expectedErr
			})

		s.Error(Execute(), "Execute should return error when root command returns error")
	},
	)
}

// TestGetBeforeExecuteDoesNotPoison verifies that calling a Get* function before
// Execute() does not lock the global viper to a bare instance; after Execute()
// runs, Get* must return values from the active command's viper.
func (s *testKrait) TestGetBeforeExecuteDoesNotPoison() {
	s.WithCustom(nil, nil, func() {
		// Call GetString before Execute — previously this would pin a bare viper
		// via sync.Once, making subsequent calls return empty values.
		earlyResult := GetString("string-param")
		s.Equal("", earlyResult, "Should return empty string before Execute")

		App("test-app", "test description", "test long description").
			WithString("string-param", "a string", "string-param", "TEST_STRING", "hello-world").
			WithRun(func(args []string) error { return nil })

		s.NoError(Execute())

		// After Execute the global getter must see the command's configured value.
		s.Equal("hello-world", GetString("string-param"), "GetString should return command default after Execute")
	})
}

func (s *testKrait) TestIsDebugNegativeWithoutCurrentCommand() {
	s.WithCustom(nil, nil, func() {
		s.False(IsDebug(), "IsDebug should return false when debug is not enabled")
	})
}

func (s *testKrait) TestIsDebugNegativeWithCurrentCommand() {
	s.WithCustom(nil, nil, func() {
		App("test-app", "test description", "test long description").
			WithRun(func(args []string) error {
				return nil
			}).
			WithDebug("debug", "", "APP_DEBUG")
		if s.NoError(Execute()) {
			s.False(IsDebug(), "IsDebug should return false when debug is not enabled")
		}
	})
}

func (s *testKrait) TestIsDebugPositiveWithCurrentCommand() {
	s.WithCustom(nil, []string{"--debug"}, func() {
		App("test-app", "test description", "test long description").
			WithRun(func(args []string) error {
				return nil
			}).
			WithDebug("debug", "", "APP_DEBUG")
		if s.NoError(Execute()) {
			s.True(IsDebug(), "IsDebug should return true when debug is enabled")
		}
	})
}

// Unmarshal (package-level) — additional coverage beyond command_constraints_test.go

func (s *testKrait) TestUnmarshalDecodesDefaultValue() {
	s.WithCustom(nil, nil, func() {
		type Config struct {
			App struct {
				Port int `mapstructure:"port"`
			} `mapstructure:"app"`
		}

		var got Config
		cmd := New("app", "app", "").
			WithInt("app.port", "Port", "port", "APP_PORT", 8080).
			WithRun(func(args []string) error {
				return Unmarshal(&got)
			})

		s.NoError(cmd.Execute())
		s.Equal(8080, got.App.Port, "Package-level Unmarshal must decode default value into struct")
	})
}
