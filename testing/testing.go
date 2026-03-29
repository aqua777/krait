package testing

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

// T is an alias for testing.T
type T = testing.T

// Run runs the test suite. It's a proxy for testify/suite.Run using our custom Suite type
func Run(t *T, s suite.TestingSuite) {
	suite.Run(t, s)
}

// Suite is the base test suite for krait tests
type Suite struct {
	suite.Suite

	// OriginalEnvironment stores the original environment variables before modification
	originalEnvironment map[string]string
	// OriginalArgs stores the original command line arguments before modification
	originalArgs []string
}

// WithCustom executes the given test function with custom environment variables and command line arguments.
// It saves the original environment and args, applies the custom ones, runs the test function,
// and then restores the original environment and args.
// Note: The args parameter should only contain the arguments without the application name.
// The application name 'test-app' will be automatically prepended.
func (s *Suite) WithCustom(env map[string]string, args []string, testFn func()) {
	// Save original environment
	s.originalEnvironment = make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if len(pair) == 2 {
			s.originalEnvironment[pair[0]] = pair[1]
		}
	}

	// Save original args
	s.originalArgs = os.Args

	// Apply custom environment
	os.Clearenv()
	for k, v := range env {
		os.Setenv(k, v)
	}
	// Restore any original env vars that weren't overridden
	for k, v := range s.originalEnvironment {
		if _, exists := env[k]; !exists {
			os.Setenv(k, v)
		}
	}

	// Set test args with 'test-app' prepended
	os.Args = append([]string{"test-app"}, args...)

	// Execute test function
	testFn()

	// Restore original environment
	os.Clearenv()
	for k, v := range s.originalEnvironment {
		os.Setenv(k, v)
	}

	// Restore original args
	os.Args = s.originalArgs
}
