package krait

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type testCommandHelpers struct {
	suite.Suite
}

func TestCommandHelpers(t *testing.T) {
	suite.Run(t, new(testCommandHelpers))
}

// TestWithLongFlag_UnsupportedType verifies that withLongFlag panics when
// defaultValue has a type that is not supported by pflag.
func (s *testCommandHelpers) TestWithLongFlag_UnsupportedType() {
	cmd := New("test", "test", "test")
	s.Panics(func() {
		cmd.withLongFlag(nil, "test-flag", "test flag description", struct{ value int }{value: 1})
	}, "withLongFlag should panic for an unsupported default value type")
}

// TestWithShortFlag_UnsupportedType verifies that withShortFlag panics when
// defaultValue has a type that is not supported by pflag.
func (s *testCommandHelpers) TestWithShortFlag_UnsupportedType() {
	cmd := New("test", "test", "test")
	s.Panics(func() {
		cmd.withShortFlag(nil, "test-flag", "t", "test flag description", struct{ value int }{value: 1})
	}, "withShortFlag should panic for an unsupported default value type")
}

// TestWithParam_ConfigOnlyVarBoundPanics verifies that withParam panics when
// flag is empty and a var pointer is provided — config-only mode is not
// supported for var-bound parameters.
func (s *testCommandHelpers) TestWithParam_ConfigOnlyVarBoundPanics() {
	cmd := New("test", "test", "test")
	var target string
	s.Panics(func() {
		cmd.withParam("some.key", &target, "", "", "", "description", "default")
	}, "withParam should panic when flag is empty and varPtr is non-nil")
}
