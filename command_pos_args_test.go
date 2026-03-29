package krait

import (
	"fmt"

	"github.com/aqua777/krait/testing"
)

type testCommandPosArgs struct {
	testing.Suite
}

func TestCommandPosArgs(t *testing.T) {
	testing.Run(t, new(testCommandPosArgs))
}

func (suite *testCommandPosArgs) TestNoArgsSuccess() {
	suite.WithCustom(map[string]string{}, []string{}, func() {
		cmd := newTestCommand().WithNoArgs()
		err := cmd.Execute()
		suite.NoError(err)
	})
}

func (suite *testCommandPosArgs) TestNoArgsFailure() {
	suite.WithCustom(map[string]string{}, []string{"arg1"}, func() {
		cmd := newTestCommand().WithNoArgs()
		err := cmd.Execute()
		suite.Error(err)
	})
}

func (suite *testCommandPosArgs) TestArbitraryArgsSuccess() {
	suite.WithCustom(map[string]string{}, []string{"arg1", "arg2"}, func() {
		cmd := newTestCommand().WithArbitraryArgs()
		err := cmd.Execute()
		suite.NoError(err)
	})
}

func (suite *testCommandPosArgs) TestMinimumNArgsSuccess() {
	suite.WithCustom(map[string]string{}, []string{"arg1", "arg2"}, func() {
		cmd := newTestCommand().WithMinimumNArgs(2)
		err := cmd.Execute()
		suite.NoError(err)
	})
}

func (suite *testCommandPosArgs) TestMinimumNArgsFailure() {
	suite.WithCustom(map[string]string{}, []string{"arg1"}, func() {
		cmd := newTestCommand().WithMinimumNArgs(2)
		err := cmd.Execute()
		suite.Error(err)
	})
}

func (suite *testCommandPosArgs) TestMaximumNArgsSuccess() {
	suite.WithCustom(map[string]string{}, []string{"arg1", "arg2"}, func() {
		cmd := newTestCommand().WithMaximumNArgs(2)
		err := cmd.Execute()
		suite.NoError(err)
	})
}

func (suite *testCommandPosArgs) TestMaximumNArgsFailure() {
	suite.WithCustom(map[string]string{}, []string{"arg1", "arg2", "arg3"}, func() {
		cmd := newTestCommand().WithMaximumNArgs(2)
		err := cmd.Execute()
		suite.Error(err)
	})
}

func (suite *testCommandPosArgs) TestRangeArgsSuccess() {
	for _, testCase := range [][]string{
		{"arg1", "arg2"},
		{"arg1", "arg2", "arg3"},
		{"arg1", "arg2", "arg3", "arg4"},
	} {
		suite.WithCustom(map[string]string{}, testCase, func() {
			cmd := newTestCommand().WithRangeArgs(2, 4)
			err := cmd.Execute()
			suite.NoError(err)
		})
	}
}

func (suite *testCommandPosArgs) TestRangeArgsFailure() {
	for _, testCase := range [][]string{
		{"arg1"},
		{"arg1", "arg2", "arg3", "arg4", "arg5"},
	} {
		suite.WithCustom(map[string]string{}, testCase, func() {
			cmd := newTestCommand().WithRangeArgs(2, 4)
			err := cmd.Execute()
			suite.Error(err)
		})
	}
}

func (suite *testCommandPosArgs) TestExactArgsSuccess() {
	suite.WithCustom(map[string]string{}, []string{"arg1", "arg2"}, func() {
		cmd := newTestCommand().WithExactArgs(2)
		err := cmd.Execute()
		suite.NoError(err)
	})
}

func (suite *testCommandPosArgs) TestExactArgsFailure() {
	for _, testCase := range [][]string{
		{"arg1"},
		{"arg1", "arg2", "arg3"},
	} {
		suite.WithCustom(map[string]string{}, testCase, func() {
			cmd := newTestCommand().WithExactArgs(2)
			err := cmd.Execute()
			suite.Error(err)
		})
	}
}

func (suite *testCommandPosArgs) TestCustomArgsSuccess() {
	suite.WithCustom(map[string]string{}, []string{"arg1", "arg2"}, func() {
		cmd := newTestCommand().WithCustomArgs(func(args []string) error {
			if len(args) != 2 {
				return fmt.Errorf("expected 2 arguments, got %d", len(args))
			}
			return nil
		})
		err := cmd.Execute()
		suite.NoError(err)
	})
}

func (suite *testCommandPosArgs) TestCustomArgsFailure() {
	suite.WithCustom(map[string]string{}, []string{"arg1"}, func() {
		cmd := newTestCommand().WithCustomArgs(func(args []string) error {
			if len(args) != 2 {
				return fmt.Errorf("expected 2 arguments, got %d", len(args))
			}
			return nil
		})
		err := cmd.Execute()
		suite.Error(err)
	})
}
