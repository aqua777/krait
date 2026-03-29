package krait

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

// Untestable coverage note for helpers.go asJson:
//
// - json.MarshalIndent error return: The function is only ever called with
//   map[string]interface{} values produced by Viper's AllSettings(). These maps
//   contain only JSON-serialisable types (strings, numbers, bools, slices, maps),
//   so MarshalIndent cannot return an error in practice. Forcing a marshal error
//   would require replacing the json package, which is not feasible without
//   introducing test-only indirection that violates Principle #5.

type testHelpers struct {
	suite.Suite
}

func TestHelpers(t *testing.T) {
	suite.Run(t, new(testHelpers))
}

func (s *testHelpers) TestIsDefaultString() {
	s.True(isDefault(""), "Empty string should be default")
	s.False(isDefault("hello"), "Non-empty string should not be default")
}

func (s *testHelpers) TestIsDefaultInt() {
	s.True(isDefault(0), "Zero int should be default")
	s.False(isDefault(42), "Non-zero int should not be default")
}

func (s *testHelpers) TestIsDefaultInt8() {
	s.True(isDefault(int8(0)), "Zero int8 should be default")
	s.False(isDefault(int8(8)), "Non-zero int8 should not be default")
}

func (s *testHelpers) TestIsDefaultInt16() {
	s.True(isDefault(int16(0)), "Zero int16 should be default")
	s.False(isDefault(int16(16)), "Non-zero int16 should not be default")
}

func (s *testHelpers) TestIsDefaultInt32() {
	s.True(isDefault(int32(0)), "Zero int32 should be default")
	s.False(isDefault(int32(32)), "Non-zero int32 should not be default")
}

func (s *testHelpers) TestIsDefaultInt64() {
	s.True(isDefault(int64(0)), "Zero int64 should be default")
	s.False(isDefault(int64(64)), "Non-zero int64 should not be default")
}

func (s *testHelpers) TestIsDefaultUint() {
	s.True(isDefault(uint(0)), "Zero uint should be default")
	s.False(isDefault(uint(1)), "Non-zero uint should not be default")
}

func (s *testHelpers) TestIsDefaultUint8() {
	s.True(isDefault(uint8(0)), "Zero uint8 should be default")
	s.False(isDefault(uint8(8)), "Non-zero uint8 should not be default")
}

func (s *testHelpers) TestIsDefaultUint16() {
	s.True(isDefault(uint16(0)), "Zero uint16 should be default")
	s.False(isDefault(uint16(16)), "Non-zero uint16 should not be default")
}

func (s *testHelpers) TestIsDefaultUint32() {
	s.True(isDefault(uint32(0)), "Zero uint32 should be default")
	s.False(isDefault(uint32(32)), "Non-zero uint32 should not be default")
}

func (s *testHelpers) TestIsDefaultUint64() {
	s.True(isDefault(uint64(0)), "Zero uint64 should be default")
	s.False(isDefault(uint64(64)), "Non-zero uint64 should not be default")
}

func (s *testHelpers) TestIsDefaultFloat32() {
	s.True(isDefault(float32(0)), "Zero float32 should be default")
	s.False(isDefault(float32(3.14)), "Non-zero float32 should not be default")
}

func (s *testHelpers) TestIsDefaultFloat64() {
	s.True(isDefault(float64(0)), "Zero float64 should be default")
	s.False(isDefault(float64(3.14)), "Non-zero float64 should not be default")
}

func (s *testHelpers) TestIsDefaultBool() {
	s.True(isDefault(false), "False bool should be default")
	s.False(isDefault(true), "True bool should not be default")
}

func (s *testHelpers) TestIsDefaultDuration() {
	s.True(isDefault(time.Duration(0)), "Zero duration should be default")
	s.False(isDefault(time.Hour), "Non-zero duration should not be default")
}

func (s *testHelpers) TestIsDefaultStringSlice() {
	s.True(isDefault([]string{}), "Empty string slice should be default")
	s.False(isDefault([]string{"hello"}), "Non-empty string slice should not be default")
}

func (s *testHelpers) TestIsDefaultStringMap() {
	s.True(isDefault(map[string]string{}), "Empty string map should be default")
	s.False(isDefault(map[string]string{"key": "value"}), "Non-empty string map should not be default")
}

func (s *testHelpers) TestIsDefaultCustomType() {
	type customStruct struct {
		_ string
	}
	s.False(isDefault(customStruct{}), "Custom type should return false as default")
}
