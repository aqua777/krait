package krait

import (
	"time"

	"github.com/aqua777/krait/testing"
)

type testCommandConfigFile struct {
	testing.Suite
}

func TestCommandConfigFile(t *testing.T) {
	testing.Run(t, new(testCommandConfigFile))
}

func (suite *testCommandConfigFile) SetupTest() {
	Reset() // Reset global state before each test
}

// Helper function to verify config values
func (suite *testCommandConfigFile) verifyConfigValues() {
	// Test string value
	suite.Equal("MyApplication", GetString("app.name"), "String value should match")

	// Test integer value
	suite.Equal(8080, GetInt("app.port"), "Integer value should match")

	// Test unsigned integer value
	suite.Equal(uint(1000), GetUint("app.max_connections"), "Unsigned integer value should match")

	// Test boolean value
	suite.Equal(true, GetBool("app.debug_mode"), "Boolean value should match")

	// Test duration values
	suite.Equal(30*time.Second, GetDuration("app.timeout"), "Timeout duration should match")
	suite.Equal(5*time.Minute, GetDuration("app.refresh_interval"), "Refresh interval should match")

	// Test string slice
	expectedOrigins := []string{
		"http://localhost:3000",
		"https://example.com",
		"https://api.example.com",
	}
	suite.ElementsMatch(expectedOrigins, GetStringSlice("app.allowed_origins"), "String slice should match")

	// Test string map
	expectedDB := map[string]string{
		"host": "localhost",
		"port": "5432",
		"name": "mydb",
		"user": "admin",
	}
	suite.Equal(expectedDB, GetStringToString("app.database"), "String map should match")

	// Test nested configuration
	suite.Equal("info", GetString("app.logging.level"), "Nested string value should match")
	suite.Equal("json", GetString("app.logging.format"), "Nested string value should match")
	suite.Equal("stdout", GetString("app.logging.output"), "Nested string value should match")

	// Test retry configuration
	suite.Equal(3, GetInt("app.retry.attempts"), "Retry attempts should match")
	suite.Equal(time.Second, GetDuration("app.retry.delay"), "Retry delay should match")
	suite.Equal(5*time.Second, GetDuration("app.retry.max_delay"), "Retry max delay should match")
}

// Helper function to setup command with common parameters
func (suite *testCommandConfigFile) setupTestCommand(shortFlag, defaultConfigPath string) *Command {
	cmd := New("test", "Test command", "Test command long description")
	cmd.WithConfig(defaultConfigPath, "config", shortFlag, "TEST_CONFIG")

	// Add parameters that match our config file
	cmd.WithString("app.name", "Application name", "name", "APP_NAME")
	cmd.WithInt("app.port", "Application port", "port", "APP_PORT")
	cmd.WithUint("app.max_connections", "Max connections", "max-connections", "APP_MAX_CONNECTIONS")
	cmd.WithBool("app.debug_mode", "Debug mode", "debug", "APP_DEBUG")
	cmd.WithDuration("app.timeout", "Timeout duration", "timeout", "APP_TIMEOUT")
	cmd.WithDuration("app.refresh_interval", "Refresh interval", "refresh", "APP_REFRESH")
	cmd.WithStringSlice("app.allowed_origins", "Allowed origins", "origins", "APP_ORIGINS")
	cmd.WithStringToString("app.database", "Database configuration", "db", "APP_DB")

	// Add a simple run function
	cmd.WithRun(func(args []string) error {
		return nil
	})

	return cmd
}

func (suite *testCommandConfigFile) TestCommandWithDefaultConfigFile() {
	suite.WithCustom(nil, nil, func() {
		cmd := suite.setupTestCommand("", "examples/app.yaml")
		err := cmd.Execute()
		suite.NoError(err, "Execute should not return an error")
		suite.verifyConfigValues()
	})
}

func (suite *testCommandConfigFile) TestCommandWithConfigFileFromArgsLong() {
	args := []string{
		"--config", "examples/app.yaml",
	}

	suite.WithCustom(nil, args, func() {
		cmd := suite.setupTestCommand("", "") // Empty default config path
		err := cmd.Execute()
		suite.NoError(err, "Execute should not return an error")
		suite.verifyConfigValues()
	})
}

func (suite *testCommandConfigFile) TestCommandWithConfigFileFromArgsShort() {
	args := []string{
		"-c", "examples/app.yaml",
	}

	suite.WithCustom(nil, args, func() {
		cmd := suite.setupTestCommand("c", "") // Empty default config path
		err := cmd.Execute()
		suite.NoError(err, "Execute should not return an error")
		suite.verifyConfigValues()
	})
}

func (suite *testCommandConfigFile) TestCommandWithInvalidConfigFile() {
	args := []string{
		"--config", "example/non_existent.yaml",
	}

	suite.WithCustom(nil, args, func() {
		// Create a test command
		cmd := suite.setupTestCommand("c", "") // Empty default config path
		err := cmd.Execute()
		suite.Error(err, "Execute should return an error with invalid config file")
	})
}

// TestCommandWithConfigFileFromEnvVar verifies that the config file path can be
// provided via environment variable when no CLI flag is given.
func (suite *testCommandConfigFile) TestCommandWithConfigFileFromEnvVar() {
	env := map[string]string{
		"TEST_CONFIG": "examples/app.yaml",
	}

	suite.WithCustom(env, nil, func() {
		cmd := suite.setupTestCommand("c", "") // Empty default config path
		err := cmd.Execute()
		suite.NoError(err, "Execute should not return an error when config path is set via env var")
		suite.verifyConfigValues()
	})
}

// TestCommandWithConfigFileCliOverridesEnvVar verifies that an explicit CLI flag
// takes precedence over the environment variable for the config file path.
func (suite *testCommandConfigFile) TestConfigOnlyParamResolvedFromConfigFile() {
	suite.WithCustom(nil, nil, func() {
		cmd := New("test", "Test command", "")
		cmd.WithConfig("examples/app.yaml", "config", "", "TEST_CONFIG")
		cmd.WithString("app.logging.level", "Log level", "", "", "default-level")
		cmd.WithRun(func(args []string) error { return nil })

		err := cmd.Execute()
		suite.NoError(err)
		suite.Equal("info", GetString("app.logging.level"), "Config file value should override default for config-only param")
	})
}

func (suite *testCommandConfigFile) TestConfigOnlyParamFallsBackToDefault() {
	suite.WithCustom(nil, nil, func() {
		cmd := New("test", "Test command", "")
		cmd.WithString("app.internal.key", "Internal key", "", "", "my-default")
		cmd.WithRun(func(args []string) error { return nil })

		err := cmd.Execute()
		suite.NoError(err)
		suite.Equal("my-default", GetString("app.internal.key"), "Config-only param should fall back to default when no config file")
	})
}

func (suite *testCommandConfigFile) TestConfigOnlyParamNotOverriddenByEnvVar() {
	env := map[string]string{
		"APP_LOGGING_LEVEL": "warn",
	}
	suite.WithCustom(env, nil, func() {
		cmd := New("test", "Test command", "")
		cmd.WithConfig("examples/app.yaml", "config", "", "TEST_CONFIG")
		cmd.WithString("app.logging.level", "Log level", "", "", "default-level")
		cmd.WithRun(func(args []string) error { return nil })

		err := cmd.Execute()
		suite.NoError(err)
		suite.Equal("info", GetString("app.logging.level"), "Config-only param should not be overridden by env var")
	})
}

func (suite *testCommandConfigFile) TestCommandWithConfigFileCliOverridesEnvVar() {
	env := map[string]string{
		"TEST_CONFIG": "example/non_existent.yaml", // would fail if used
	}
	args := []string{
		"--config", "examples/app.yaml",
	}

	suite.WithCustom(env, args, func() {
		cmd := suite.setupTestCommand("c", "")
		err := cmd.Execute()
		suite.NoError(err, "CLI flag should override env var for config file path")
		suite.verifyConfigValues()
	})
}
