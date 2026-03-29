package krait

import (
	"testing"

	kraittest "github.com/aqua777/krait/testing"
)

func TestCommandGroups(t *testing.T) {
	kraittest.Run(t, new(testCommandGroups))
}

type testCommandGroups struct {
	kraittest.Suite
}

func (suite *testCommandGroups) SetupTest() {
	Reset()
}

// WithGroup

func (suite *testCommandGroups) TestWithGroupAddsGroupToCobraCommand() {
	cmd := New("app", "app", "")
	cmd.WithGroup("mgmt", "Management Commands")
	suite.Len(cmd.groups, 1)
	suite.Equal("mgmt", cmd.groups[0].ID)
	suite.Equal("Management Commands", cmd.groups[0].Title)
}

func (suite *testCommandGroups) TestWithGroupPanicsOnDuplicateID() {
	cmd := New("app", "app", "")
	cmd.WithGroup("mgmt", "Management Commands")
	suite.Panics(func() {
		cmd.WithGroup("mgmt", "Duplicate")
	})
}

func (suite *testCommandGroups) TestWithGroupChaining() {
	cmd := New("app", "app", "")
	result := cmd.WithGroup("mgmt", "Management Commands")
	suite.Same(cmd, result)
}

func (suite *testCommandGroups) TestWithGroupMultipleDistinctGroupsRegistered() {
	cmd := New("app", "app", "")
	cmd.WithGroup("mgmt", "Management Commands").WithGroup("dev", "Developer Commands")
	suite.Len(cmd.groups, 2)
}

// InGroup

func (suite *testCommandGroups) TestInGroupSetsGroupIDOnCobraCommand() {
	sub := New("serve", "Serve", "")
	sub.InGroup("mgmt")
	suite.Equal("mgmt", sub.cmd.GroupID)
	suite.Equal("mgmt", sub.groupID)
}

func (suite *testCommandGroups) TestInGroupPanicsOnEmptyID() {
	sub := New("serve", "Serve", "")
	suite.Panics(func() {
		sub.InGroup("")
	})
}

func (suite *testCommandGroups) TestInGroupChaining() {
	sub := New("serve", "Serve", "")
	result := sub.InGroup("mgmt")
	suite.Same(sub, result)
}

// Integration: WithGroup + InGroup + WithCommand

func (suite *testCommandGroups) TestSubcommandGroupIDSetAfterWithCommand() {
	suite.WithCustom(nil, nil, func() {
		sub := New("serve", "Serve", "").
			InGroup("mgmt").
			WithRun(func(args []string) error { return nil })

		App("app", "app", "").
			WithGroup("mgmt", "Management Commands").
			WithCommand(sub)

		suite.Equal("mgmt", sub.cmd.GroupID)
	})
}

func (suite *testCommandGroups) TestSubcommandWithoutInGroupHasEmptyGroupID() {
	suite.WithCustom(nil, nil, func() {
		sub := New("serve", "Serve", "").
			WithRun(func(args []string) error { return nil })

		App("app", "app", "").
			WithGroup("mgmt", "Management Commands").
			WithCommand(sub)

		suite.Equal("", sub.cmd.GroupID)
	})
}
