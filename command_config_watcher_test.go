package krait

// Testing the WithConfigWatcher method.
//
// Full concurrency behaviour (concurrent reads during callback invocation, races
// between the watcher goroutine and the main run goroutine) cannot be unit tested
// without introducing significant timing dependencies. The demo script
// docs/sprints/current/demos/phase4_watcher.go covers this scenario end-to-end.
//
// The integration tests below use a temp file and a brief sleep to confirm that
// the callback fires after a write, but do not make assertions about concurrent
// state access.

import (
	"os"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	kraittest "github.com/aqua777/krait/testing"
)

func TestCommandConfigWatcher(t *testing.T) {
	kraittest.Run(t, new(testCommandConfigWatcher))
}

type testCommandConfigWatcher struct {
	kraittest.Suite
}

func (suite *testCommandConfigWatcher) SetupTest() {
	Reset()
}

func (suite *testCommandConfigWatcher) TestWithConfigWatcherChaining() {
	cmd := New("app", "app", "")
	result := cmd.WithConfigWatcher(func() {})
	suite.Same(cmd, result)
}

func (suite *testCommandConfigWatcher) TestWithConfigWatcherWithoutConfigFileIsIgnored() {
	suite.WithCustom(nil, nil, func() {
		ran := false
		err := App("app", "app", "").
			WithConfigWatcher(func() {}).
			WithRun(func(args []string) error {
				ran = true
				return nil
			}).
			Execute()
		suite.NoError(err)
		suite.True(ran)
	})
}

func (suite *testCommandConfigWatcher) TestWithConfigWatcherWithConfigReaderIsIgnored() {
	suite.WithCustom(nil, nil, func() {
		ran := false
		err := App("app", "app", "").
			WithConfigReader(strings.NewReader("host: localhost\n"), "yaml").
			WithConfigWatcher(func() {}).
			WithString("host", "Hostname", "host", "", "default").
			WithRun(func(args []string) error {
				ran = true
				return nil
			}).
			Execute()
		suite.NoError(err)
		suite.True(ran)
	})
}

func (suite *testCommandConfigWatcher) TestWithConfigWatcherFiresOnFileChange() {
	f, err := os.CreateTemp("", "krait-watcher-*.yaml")
	suite.Require().NoError(err)
	defer os.Remove(f.Name())

	_, err = f.WriteString("val: original\n")
	suite.Require().NoError(err)
	f.Close()

	var callCount atomic.Int32

	suite.WithCustom(nil, nil, func() {
		err := App("app", "app", "").
			WithConfig(f.Name(), "config", "c", "").
			WithString("val", "Value", "val", "", "default").
			WithConfigWatcher(func() {
				callCount.Add(1)
			}).
			WithRun(func(args []string) error {
				// Write a new value to the config file to trigger the watcher.
				time.Sleep(50 * time.Millisecond)
				os.WriteFile(f.Name(), []byte("val: updated\n"), 0644)
				// Give fsnotify time to fire the callback.
				time.Sleep(500 * time.Millisecond)
				return nil
			}).
			Execute()
		suite.NoError(err)
	})

	suite.GreaterOrEqual(callCount.Load(), int32(1))
}

func (suite *testCommandConfigWatcher) TestWithConfigWatcherFiresOnSearchPathFileChange() {
	dir, err := os.MkdirTemp("", "krait-watcher-dir-*")
	suite.Require().NoError(err)
	defer os.RemoveAll(dir)

	configPath := dir + "/myapp.yaml"
	err = os.WriteFile(configPath, []byte("val: original\n"), 0644)
	suite.Require().NoError(err)

	var callCount atomic.Int32

	suite.WithCustom(nil, nil, func() {
		err := App("app", "app", "").
			WithConfigName("myapp").
			WithConfigPath(dir).
			WithString("val", "Value", "val", "", "default").
			WithConfigWatcher(func() {
				callCount.Add(1)
			}).
			WithRun(func(args []string) error {
				time.Sleep(50 * time.Millisecond)
				os.WriteFile(configPath, []byte("val: updated\n"), 0644)
				time.Sleep(500 * time.Millisecond)
				return nil
			}).
			Execute()
		suite.NoError(err)
	})

	suite.GreaterOrEqual(callCount.Load(), int32(1))
}

