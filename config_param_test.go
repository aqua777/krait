package krait

import (
	"testing"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

// unsupportedFlagValue is a pflag.Value whose Type() returns an unrecognised
// string so we can exercise the default error branch of setVarPtrValue.
type unsupportedFlagValue struct{}

func (u *unsupportedFlagValue) String() string { return "" }
func (u *unsupportedFlagValue) Set(_ string) error { return nil }
func (u *unsupportedFlagValue) Type() string   { return "unsupported_custom_type" }

const (
	testFlagName      = "test-flag"
	testFlagUsageDesc = "test flag"
)

type testData struct {
	name          string
	expectedValue any
	actualValue   any
	varType       ConfigParamType
}

type testConfigParam struct {
	suite.Suite
}

func TestConfigParamSuite(t *testing.T) {
	suite.Run(t, new(testConfigParam))
}

func (s *testConfigParam) TestString() {
	tests := []struct {
		name     string
		param    ConfigParam
		expected string
	}{
		{
			name: "all fields populated",
			param: ConfigParam{
				Name:               "test-param",
				Flag:               "--test-flag",
				ShortFlag:          "-t",
				EnvironmentVarName: "TEST_ENV_VAR",
				DefaultValue:       "default-value",
			},
			expected: "Name: test-param, Flag: --test-flag, ShortFlag: -t, EnvironmentVarName: TEST_ENV_VAR, DefaultValue: default-value",
		},
		{
			name: "empty fields",
			param: ConfigParam{
				Name: "empty-param",
			},
			expected: "Name: empty-param, Flag: , ShortFlag: , EnvironmentVarName: , DefaultValue: <nil>",
		},
		{
			name: "numeric default value",
			param: ConfigParam{
				Name:               "number-param",
				Flag:               "--number",
				ShortFlag:          "-n",
				EnvironmentVarName: "NUMBER_VAR",
				DefaultValue:       42,
			},
			expected: "Name: number-param, Flag: --number, ShortFlag: -n, EnvironmentVarName: NUMBER_VAR, DefaultValue: 42",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			got := tt.param.String()
			s.Equal(tt.expected, got)
		})
	}
}

func (s *testConfigParam) TestSetVarPtrValue() {
	for _, tt := range s.getTestData() {
		s.Run(tt.name, func() {
			flag := s.setFlagVar(tt.varType, tt.actualValue)

			// Create config param
			param := &ConfigParam{
				Name:   "test-param",
				Flag:   testFlagName,
				VarPtr: tt.actualValue,
			}

			// Create viper instance and set value
			v := viper.New()
			v.Set(testFlagName, tt.expectedValue)

			// Call setVarPtrValue
			err := param.setVarPtrValue(flag, v)
			s.NoError(err, "setVarPtrValue should not return an error for supported type")

			// Verify the value was set correctly
			s.verifyValue(tt.varType, tt.expectedValue, tt.actualValue)
		})
	}
}

// TestSetVarPtrValue_UnsupportedType verifies that setVarPtrValue returns an
// error when the flag has a type that is not handled by the switch.
func (s *testConfigParam) TestSetVarPtrValue_UnsupportedType() {
	fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
	fs.Var(&unsupportedFlagValue{}, testFlagName, testFlagUsageDesc)
	flag := fs.Lookup(testFlagName)

	param := &ConfigParam{
		Name:   "test-param",
		Flag:   testFlagName,
		VarPtr: new(string),
	}

	v := viper.New()
	err := param.setVarPtrValue(flag, v)
	s.Error(err, "setVarPtrValue should return an error for an unsupported flag type")
	s.Contains(err.Error(), "unsupported flag type", "error message should mention unsupported flag type")
}

func (s *testConfigParam) getTestData() []*testData {
	return []*testData{
		{
			name:          "string value",
			expectedValue: "test-value",
			actualValue:   new(string),
			varType:       typeString,
		},
		{
			name:          "int value",
			expectedValue: 42,
			actualValue:   new(int),
			varType:       typeInt,
		},
		{
			name:          "int8 value",
			expectedValue: int8(8),
			actualValue:   new(int8),
			varType:       typeInt8,
		},
		{
			name:          "int16 value",
			expectedValue: int16(16),
			actualValue:   new(int16),
			varType:       typeInt16,
		},
		{
			name:          "int32 value",
			expectedValue: int32(32),
			actualValue:   new(int32),
			varType:       typeInt32,
		},
		{
			name:          "int64 value",
			expectedValue: int64(64),
			actualValue:   new(int64),
			varType:       typeInt64,
		},
		{
			name:          "uint value",
			expectedValue: uint(42),
			actualValue:   new(uint),
			varType:       typeUint,
		},
		{
			name:          "uint8 value",
			expectedValue: uint8(8),
			actualValue:   new(uint8),
			varType:       typeUint8,
		},
		{
			name:          "uint16 value",
			expectedValue: uint16(16),
			actualValue:   new(uint16),
			varType:       typeUint16,
		},
		{
			name:          "uint32 value",
			expectedValue: uint32(32),
			actualValue:   new(uint32),
			varType:       typeUint32,
		},
		{
			name:          "uint64 value",
			expectedValue: uint64(64),
			actualValue:   new(uint64),
			varType:       typeUint64,
		},
		{
			name:          "bool value",
			expectedValue: true,
			actualValue:   new(bool),
			varType:       typeBool,
		},
		{
			name:          "float32 value",
			expectedValue: float32(3.14),
			actualValue:   new(float32),
			varType:       typeFloat32,
		},
		{
			name:          "float64 value",
			expectedValue: float64(3.14159),
			actualValue:   new(float64),
			varType:       typeFloat64,
		},
		{
			name:          "duration value",
			expectedValue: time.Hour,
			actualValue:   new(time.Duration),
			varType:       typeDuration,
		},
		{
			name:          "string slice value",
			expectedValue: []string{"one", "two", "three"},
			actualValue:   new([]string),
			varType:       typeStringSlice,
		},
		{
			name:          "string map value",
			expectedValue: map[string]string{"key1": "value1", "key2": "value2"},
			actualValue:   &map[string]string{},
			varType:       typeStringToString,
		},
	}
}

func (s *testConfigParam) setFlagVar(varType ConfigParamType, actualValue any) *pflag.Flag {
	fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
	switch varType {
	case typeString:
		fs.StringVar(actualValue.(*string), testFlagName, "", testFlagUsageDesc)
	case typeInt:
		fs.IntVar(actualValue.(*int), testFlagName, 0, testFlagUsageDesc)
	case typeInt8:
		fs.Int8Var(actualValue.(*int8), testFlagName, 0, testFlagUsageDesc)
	case typeInt16:
		fs.Int16Var(actualValue.(*int16), testFlagName, 0, testFlagUsageDesc)
	case typeInt32:
		fs.Int32Var(actualValue.(*int32), testFlagName, 0, testFlagUsageDesc)
	case typeInt64:
		fs.Int64Var(actualValue.(*int64), testFlagName, 0, testFlagUsageDesc)
	case typeUint:
		fs.UintVar(actualValue.(*uint), testFlagName, 0, testFlagUsageDesc)
	case typeUint8:
		fs.Uint8Var(actualValue.(*uint8), testFlagName, 0, testFlagUsageDesc)
	case typeUint16:
		fs.Uint16Var(actualValue.(*uint16), testFlagName, 0, testFlagUsageDesc)
	case typeUint32:
		fs.Uint32Var(actualValue.(*uint32), testFlagName, 0, testFlagUsageDesc)
	case typeUint64:
		fs.Uint64Var(actualValue.(*uint64), testFlagName, 0, testFlagUsageDesc)
	case typeBool:
		fs.BoolVar(actualValue.(*bool), testFlagName, false, testFlagUsageDesc)
	case typeFloat32:
		fs.Float32Var(actualValue.(*float32), testFlagName, 0, testFlagUsageDesc)
	case typeFloat64:
		fs.Float64Var(actualValue.(*float64), testFlagName, 0, testFlagUsageDesc)
	case typeDuration:
		fs.DurationVar(actualValue.(*time.Duration), testFlagName, 0, testFlagUsageDesc)
	case typeStringSlice:
		fs.StringSliceVar(actualValue.(*[]string), testFlagName, nil, testFlagUsageDesc)
	case typeStringToString:
		fs.StringToStringVar(actualValue.(*map[string]string), testFlagName, nil, testFlagUsageDesc)
	}
	flag := fs.Lookup(testFlagName)
	return flag
}

func (s *testConfigParam) verifyValue(varType ConfigParamType, expectedValue, actualValue any) {
	switch varType {
	case typeString:
		s.Equal(expectedValue, *actualValue.(*string))
	case typeInt:
		s.Equal(expectedValue, *actualValue.(*int))
	case typeInt8:
		s.Equal(expectedValue, *actualValue.(*int8))
	case typeInt16:
		s.Equal(expectedValue, *actualValue.(*int16))
	case typeInt32:
		s.Equal(expectedValue, *actualValue.(*int32))
	case typeInt64:
		s.Equal(expectedValue, *actualValue.(*int64))
	case typeUint:
		s.Equal(expectedValue, *actualValue.(*uint))
	case typeUint8:
		s.Equal(expectedValue, *actualValue.(*uint8))
	case typeUint16:
		s.Equal(expectedValue, *actualValue.(*uint16))
	case typeUint32:
		s.Equal(expectedValue, *actualValue.(*uint32))
	case typeUint64:
		s.Equal(expectedValue, *actualValue.(*uint64))
	case typeBool:
		s.Equal(expectedValue, *actualValue.(*bool))
	case typeFloat32:
		s.Equal(expectedValue, *actualValue.(*float32))
	case typeFloat64:
		s.Equal(expectedValue, *actualValue.(*float64))
	case typeDuration:
		s.Equal(expectedValue, *actualValue.(*time.Duration))
	case typeStringSlice:
		s.Equal(expectedValue, *actualValue.(*[]string))
	case typeStringToString:
		s.Equal(expectedValue, *actualValue.(*map[string]string))
	}
}

func (s *testConfigParam) TestIsConfigOnlyBothEmpty() {
	param := &ConfigParam{Flag: "", EnvironmentVarName: ""}
	s.True(param.IsConfigOnly(), "Both empty: should be config-only")
}

func (s *testConfigParam) TestIsConfigOnlyFlagNonEmpty() {
	param := &ConfigParam{Flag: "my-flag", EnvironmentVarName: ""}
	s.False(param.IsConfigOnly(), "Non-empty flag: should not be config-only")
}

func (s *testConfigParam) TestIsConfigOnlyEnvVarNonEmpty() {
	param := &ConfigParam{Flag: "", EnvironmentVarName: "MY_VAR"}
	s.False(param.IsConfigOnly(), "Non-empty env var: should not be config-only")
}

func (s *testConfigParam) TestIsConfigOnlyBothNonEmpty() {
	param := &ConfigParam{Flag: "my-flag", EnvironmentVarName: "MY_VAR"}
	s.False(param.IsConfigOnly(), "Both non-empty: should not be config-only")
}
