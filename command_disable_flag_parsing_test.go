package krait

import (
	"testing"

	kraittest "github.com/aqua777/krait/testing"
)

func TestCommandDisableFlagParsing(t *testing.T) {
	kraittest.Run(t, new(testCommandDisableFlagParsing))
}

type testCommandDisableFlagParsing struct {
	kraittest.Suite
}

func (suite *testCommandDisableFlagParsing) SetupTest() {
	Reset()
}

func (suite *testCommandDisableFlagParsing) TestWithDisableFlagParsingDeliversVerbatimArgs() {
	suite.WithCustom(nil, []string{"exec", "--foo", "bar"}, func() {
		var received []string
		exec := New("exec", "Execute pass-through command", "").
			WithDisableFlagParsing().
			WithRun(func(args []string) error {
				received = args
				return nil
			})

		err := App("app", "app", "").WithCommand(exec).Execute()
		suite.NoError(err)
		suite.Equal([]string{"--foo", "bar"}, received)
	})
}

func (suite *testCommandDisableFlagParsing) TestWithoutDisableFlagParsingUnknownFlagCausesError() {
	suite.WithCustom(nil, []string{"exec", "--foo", "bar"}, func() {
		exec := New("exec", "Execute command", "").
			WithRun(func(args []string) error { return nil })

		err := App("app", "app", "").WithCommand(exec).Execute()
		suite.Error(err)
	})
}

func (suite *testCommandDisableFlagParsing) TestWithDisableFlagParsingChaining() {
	cmd := New("exec", "Execute pass-through command", "")
	result := cmd.WithDisableFlagParsing()
	suite.Same(cmd, result)
}

func (suite *testCommandDisableFlagParsing) TestWithDisableFlagParsingSetsCobraField() {
	cmd := New("exec", "Execute pass-through command", "")
	suite.False(cmd.cmd.DisableFlagParsing)
	cmd.WithDisableFlagParsing()
	suite.True(cmd.cmd.DisableFlagParsing)
}

func (suite *testCommandDisableFlagParsing) TestDefaultDisableFlagParsingIsFalse() {
	cmd := New("exec", "Execute command", "")
	suite.False(cmd.disableFlagParsing)
	suite.False(cmd.cmd.DisableFlagParsing)
}
