package krait

import (
	"testing"

	"github.com/spf13/cobra"
	kraittest "github.com/aqua777/krait/testing"
)

func TestCommandCompletion(t *testing.T) {
	kraittest.Run(t, new(testCommandCompletion))
}

type testCommandCompletion struct {
	kraittest.Suite
}

func (suite *testCommandCompletion) SetupTest() {
	Reset()
}

// WithValidArgs

func (suite *testCommandCompletion) TestWithValidArgsSetsCobraValidArgs() {
	cmd := New("app", "app", "")
	cmd.WithValidArgs("start", "stop")
	suite.Equal([]string{"start", "stop"}, cmd.cmd.ValidArgs)
}

func (suite *testCommandCompletion) TestWithValidArgsChaining() {
	cmd := New("app", "app", "")
	result := cmd.WithValidArgs("start", "stop")
	suite.Same(cmd, result)
}

// WithValidArgsFunction

func (suite *testCommandCompletion) TestWithValidArgsFunctionSetsCobraValidArgsFunction() {
	cmd := New("app", "app", "")
	fn := func(args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"dynamic"}, cobra.ShellCompDirectiveDefault
	}
	cmd.WithValidArgsFunction(fn)
	suite.NotNil(cmd.cmd.ValidArgsFunction)
}

func (suite *testCommandCompletion) TestWithValidArgsFunctionChaining() {
	cmd := New("app", "app", "")
	fn := func(args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return nil, cobra.ShellCompDirectiveDefault
	}
	result := cmd.WithValidArgsFunction(fn)
	suite.Same(cmd, result)
}

func (suite *testCommandCompletion) TestWithValidArgsFunctionDelegatesToCallback() {
	cmd := New("app", "app", "")
	called := false
	fn := func(args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		called = true
		suite.Equal([]string{"a"}, args)
		suite.Equal("p", toComplete)
		return []string{"ping"}, cobra.ShellCompDirectiveDefault
	}
	cmd.WithValidArgsFunction(fn)
	results, directive := cmd.cmd.ValidArgsFunction(cmd.cmd, []string{"a"}, "p")
	suite.True(called)
	suite.Equal([]string{"ping"}, results)
	suite.Equal(cobra.ShellCompDirectiveDefault, directive)
}

// Both WithValidArgs and WithValidArgsFunction — ValidArgsFunction takes precedence

func (suite *testCommandCompletion) TestWithValidArgsFunctionTakesPrecedenceOverValidArgs() {
	cmd := New("app", "app", "")
	cmd.WithValidArgs("static1", "static2")
	fn := func(args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"dynamic"}, cobra.ShellCompDirectiveDefault
	}
	cmd.WithValidArgsFunction(fn)
	// Cobra uses ValidArgsFunction when set, ignoring ValidArgs — verify both are wired
	suite.NotNil(cmd.cmd.ValidArgsFunction)
	suite.Equal([]string{"static1", "static2"}, cmd.cmd.ValidArgs)
}

// WithFlagCompletion

func (suite *testCommandCompletion) TestWithFlagCompletionRegisteredFlagExecutesNormally() {
	suite.WithCustom(nil, nil, func() {
		ran := false
		cmd := App("app", "app", "").
			WithString("format", "Output format", "format", "", "text").
			WithFlagCompletion("format", func(args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
				return []string{"text", "json"}, cobra.ShellCompDirectiveDefault
			}).
			WithRun(func(args []string) error {
				ran = true
				return nil
			})
		err := cmd.Execute()
		suite.NoError(err)
		suite.True(ran)
	})
}

func (suite *testCommandCompletion) TestWithFlagCompletionUnregisteredFlagReturnsError() {
	suite.WithCustom(nil, nil, func() {
		cmd := App("app", "app", "").
			WithFlagCompletion("nonexistent", func(args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
				return nil, cobra.ShellCompDirectiveDefault
			}).
			WithRun(func(args []string) error { return nil })
		err := cmd.Execute()
		suite.Error(err)
		suite.Contains(err.Error(), "WithFlagCompletion")
	})
}

func (suite *testCommandCompletion) TestWithFlagCompletionChaining() {
	cmd := New("app", "app", "").
		WithString("format", "Output format", "format", "", "text")
	fn := func(args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return nil, cobra.ShellCompDirectiveDefault
	}
	result := cmd.WithFlagCompletion("format", fn)
	suite.Same(cmd, result)
}
