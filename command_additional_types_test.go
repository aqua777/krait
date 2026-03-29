package krait

import (
	"os"
	"testing"
	"time"

	kraittest "github.com/aqua777/krait/testing"
)

func TestAdditionalTypes(t *testing.T) {
	kraittest.Run(t, new(testAdditionalTypes))
}

type testAdditionalTypes struct {
	kraittest.Suite
}

func (suite *testAdditionalTypes) SetupTest() {
	Reset()
}

// WithIntSlice — named param

func (suite *testAdditionalTypes) TestWithIntSliceFromFlag() {
	suite.WithCustom(nil, []string{"--ids", "1,2,3"}, func() {
		cmd := newTestCommand().
			WithIntSlice("ids", "IDs", "ids", "APP_IDS")

		suite.NoError(cmd.Execute())
		suite.Equal([]int{1, 2, 3}, GetIntSlice("ids"), "WithIntSlice must parse comma-separated integers from flag")
	})
}

func (suite *testAdditionalTypes) TestWithIntSliceDefaultValue() {
	suite.WithCustom(nil, nil, func() {
		cmd := newTestCommand().
			WithIntSlice("ids", "IDs", "ids", "APP_IDS", []int{10, 20})

		suite.NoError(cmd.Execute())
		suite.Equal([]int{10, 20}, GetIntSlice("ids"), "WithIntSlice must use default value when flag is not set")
	})
}

func (suite *testAdditionalTypes) TestWithIntSlicePShortFlag() {
	suite.WithCustom(nil, []string{"-i", "7,8,9"}, func() {
		cmd := newTestCommand().
			WithIntSliceP("ids", "IDs", "ids", "i", "APP_IDS")

		suite.NoError(cmd.Execute())
		suite.Equal([]int{7, 8, 9}, GetIntSlice("ids"), "WithIntSliceP must accept short flag")
	})
}

// WithIntSliceVar — var-bound

func (suite *testAdditionalTypes) TestWithIntSliceVarFromFlag() {
	suite.WithCustom(nil, []string{"--ids", "4,5,6"}, func() {
		var ids []int
		cmd := New("app", "app", "").
			WithIntSliceVar(&ids, "IDs", "ids", "APP_IDS").
			WithRun(func(args []string) error { return nil })

		suite.NoError(cmd.Execute())
		suite.Equal([]int{4, 5, 6}, ids, "WithIntSliceVar must populate pointer from CLI flag")
	})
}

func (suite *testAdditionalTypes) TestWithIntSliceVarDefaultValue() {
	suite.WithCustom(nil, nil, func() {
		var ids []int
		cmd := New("app", "app", "").
			WithIntSliceVar(&ids, "IDs", "ids", "APP_IDS", []int{99, 100}).
			WithRun(func(args []string) error { return nil })

		suite.NoError(cmd.Execute())
		suite.Equal([]int{99, 100}, ids, "WithIntSliceVar must use default value when flag is not set")
	})
}

func (suite *testAdditionalTypes) TestWithIntSliceVarPShortFlag() {
	suite.WithCustom(nil, []string{"-i", "1,2"}, func() {
		var ids []int
		cmd := New("app", "app", "").
			WithIntSliceVarP(&ids, "IDs", "ids", "i", "APP_IDS").
			WithRun(func(args []string) error { return nil })

		suite.NoError(cmd.Execute())
		suite.Equal([]int{1, 2}, ids, "WithIntSliceVarP short flag must populate pointer")
	})
}

// GetIntSlice — package-level getter

func (suite *testAdditionalTypes) TestGetIntSliceBeforeExecuteReturnsNil() {
	suite.WithCustom(nil, nil, func() {
		result := GetIntSlice("nonexistent")
		suite.Nil(result, "GetIntSlice before Execute must return nil for unknown key")
	})
}

// GetTime — config file source

func (suite *testAdditionalTypes) TestGetTimeFromConfigFile() {
	yaml := []byte("created_at: \"2024-01-15T10:30:00Z\"\n")
	f, err := os.CreateTemp("", "krait-test-*.yaml")
	suite.Require().NoError(err)
	defer os.Remove(f.Name())
	_, err = f.Write(yaml)
	suite.Require().NoError(err)
	suite.Require().NoError(f.Close())

	suite.WithCustom(nil, nil, func() {
		cmd := New("app", "app", "").
			WithConfig(f.Name(), "config", "", "").
			WithRun(func(args []string) error { return nil })

		suite.NoError(cmd.Execute())
		got := GetTime("created_at")
		expected := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
		suite.Equal(expected, got, "GetTime must parse RFC3339 timestamp from config file")
	})
}

// GetSizeInBytes — config file source

func (suite *testAdditionalTypes) TestGetSizeInBytesFromConfigFile() {
	yaml := []byte("max_size: \"10mb\"\n")
	f, err := os.CreateTemp("", "krait-test-*.yaml")
	suite.Require().NoError(err)
	defer os.Remove(f.Name())
	_, err = f.Write(yaml)
	suite.Require().NoError(err)
	suite.Require().NoError(f.Close())

	suite.WithCustom(nil, nil, func() {
		cmd := New("app", "app", "").
			WithConfig(f.Name(), "config", "", "").
			WithRun(func(args []string) error { return nil })

		suite.NoError(cmd.Execute())
		got := GetSizeInBytes("max_size")
		suite.Equal(uint(10*1024*1024), got, "GetSizeInBytes must parse '10mb' to bytes from config file")
	})
}

// GetStringMapStringSlice — config file source

func (suite *testAdditionalTypes) TestGetStringMapStringSliceFromConfigFile() {
	yaml := []byte("headers:\n  accept:\n    - application/json\n    - text/html\n  content-type:\n    - application/json\n")
	f, err := os.CreateTemp("", "krait-test-*.yaml")
	suite.Require().NoError(err)
	defer os.Remove(f.Name())
	_, err = f.Write(yaml)
	suite.Require().NoError(err)
	suite.Require().NoError(f.Close())

	suite.WithCustom(nil, nil, func() {
		cmd := New("app", "app", "").
			WithConfig(f.Name(), "config", "", "").
			WithRun(func(args []string) error { return nil })

		suite.NoError(cmd.Execute())
		got := GetStringMapStringSlice("headers")
		suite.Equal([]string{"application/json", "text/html"}, got["accept"], "GetStringMapStringSlice must return correct slice for 'accept'")
		suite.Equal([]string{"application/json"}, got["content-type"], "GetStringMapStringSlice must return correct slice for 'content-type'")
	})
}
