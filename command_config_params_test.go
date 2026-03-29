package krait

import (
	"fmt"
	"time"

	"github.com/aqua777/krait/testing"
)

// Untestable coverage notes for command.go internalProcessParams:
//
// - argsViper.BindPFlag error return (Loop 1, var-bound path): Viper's BindPFlag only
//   errors when the flag argument is nil. The call site guards with Lookup() != nil
//   before calling BindPFlag, making the nil-flag case unreachable.
//
// - viper.BindPFlag error return (Loop 3, named-param path): Same reason as above.
//
// - setVarPtrValue error return (Loops 2 and 4): setVarPtrValue only errors when the
//   flag type is unsupported. Registration of an unsupported type panics in
//   withLongFlag/withShortFlag before any Viper state is modified, so by the time
//   internalProcessParams runs, all registered flags are of supported types.

type testCommandConfigParams struct {
	testing.Suite
	expectedError error
}

func TestCommandConfigParams(t *testing.T) {
	testing.Run(t, new(testCommandConfigParams))
}

func (suite *testCommandConfigParams) SetupTest() {
	Reset() // Reset global state before each test
	suite.expectedError = fmt.Errorf("test error")
}

func (suite *testCommandConfigParams) TearDownTest() {
}

// addParam adds the appropriate parameter type to the command
func (suite *testCommandConfigParams) addParam(cmd *Command, paramType ParamType, defaultVal interface{}) *Command {
	switch paramType {
	case StringParam:
		return cmd.WithString(testParamName, testParamDesc, testParamFlag, testParamEnv, defaultVal.(string))
	case IntParam:
		return cmd.WithInt(testParamName, testParamDesc, testParamFlag, testParamEnv, defaultVal.(int))
	case BoolParam:
		return cmd.WithBool(testParamName, testParamDesc, testParamFlag, testParamEnv, defaultVal.(bool))
	case UintParam:
		return cmd.WithUint(testParamName, testParamDesc, testParamFlag, testParamEnv, defaultVal.(uint))
	case Float32Param:
		return cmd.WithFloat32(testParamName, testParamDesc, testParamFlag, testParamEnv, defaultVal.(float32))
	case Float64Param:
		return cmd.WithFloat64(testParamName, testParamDesc, testParamFlag, testParamEnv, defaultVal.(float64))
	case DurationParam:
		return cmd.WithDuration(testParamName, testParamDesc, testParamFlag, testParamEnv, defaultVal.(time.Duration))
	case StringSliceParam:
		return cmd.WithStringSlice(testParamName, testParamDesc, testParamFlag, testParamEnv, defaultVal.([]string))
	case StringToStringParam:
		return cmd.WithStringToString(testParamName, testParamDesc, testParamFlag, testParamEnv, defaultVal.(map[string]string))
	case Int8Param:
		return cmd.WithInt8(testParamName, testParamDesc, testParamFlag, testParamEnv, defaultVal.(int8))
	case Int16Param:
		return cmd.WithInt16(testParamName, testParamDesc, testParamFlag, testParamEnv, defaultVal.(int16))
	case Int32Param:
		return cmd.WithInt32(testParamName, testParamDesc, testParamFlag, testParamEnv, defaultVal.(int32))
	case Int64Param:
		return cmd.WithInt64(testParamName, testParamDesc, testParamFlag, testParamEnv, defaultVal.(int64))
	case Uint8Param:
		return cmd.WithUint8(testParamName, testParamDesc, testParamFlag, testParamEnv, defaultVal.(uint8))
	case Uint16Param:
		return cmd.WithUint16(testParamName, testParamDesc, testParamFlag, testParamEnv, defaultVal.(uint16))
	case Uint32Param:
		return cmd.WithUint32(testParamName, testParamDesc, testParamFlag, testParamEnv, defaultVal.(uint32))
	case Uint64Param:
		return cmd.WithUint64(testParamName, testParamDesc, testParamFlag, testParamEnv, defaultVal.(uint64))
	default:
		suite.FailNow("Unsupported parameter type")
		return nil
	}
}

// addParamP adds the appropriate parameter type to the command with a short flag
func (suite *testCommandConfigParams) addParamP(cmd *Command, paramType ParamType, defaultVal interface{}) *Command {
	switch paramType {
	case StringParam:
		return cmd.WithStringP(testParamName, testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultVal.(string))
	case IntParam:
		return cmd.WithIntP(testParamName, testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultVal.(int))
	case BoolParam:
		return cmd.WithBoolP(testParamName, testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultVal.(bool))
	case UintParam:
		return cmd.WithUintP(testParamName, testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultVal.(uint))
	case Float32Param:
		return cmd.WithFloat32P(testParamName, testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultVal.(float32))
	case Float64Param:
		return cmd.WithFloat64P(testParamName, testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultVal.(float64))
	case DurationParam:
		return cmd.WithDurationP(testParamName, testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultVal.(time.Duration))
	case StringSliceParam:
		return cmd.WithStringSliceP(testParamName, testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultVal.([]string))
	case StringToStringParam:
		return cmd.WithStringToStringP(testParamName, testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultVal.(map[string]string))
	case Int8Param:
		return cmd.WithInt8P(testParamName, testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultVal.(int8))
	case Int16Param:
		return cmd.WithInt16P(testParamName, testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultVal.(int16))
	case Int32Param:
		return cmd.WithInt32P(testParamName, testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultVal.(int32))
	case Int64Param:
		return cmd.WithInt64P(testParamName, testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultVal.(int64))
	case Uint8Param:
		return cmd.WithUint8P(testParamName, testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultVal.(uint8))
	case Uint16Param:
		return cmd.WithUint16P(testParamName, testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultVal.(uint16))
	case Uint32Param:
		return cmd.WithUint32P(testParamName, testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultVal.(uint32))
	case Uint64Param:
		return cmd.WithUint64P(testParamName, testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultVal.(uint64))
	default:
		suite.FailNow("Unsupported parameter type")
		return nil
	}
}

// getValue gets the parameter value based on its type
func (suite *testCommandConfigParams) getValue(paramType ParamType) interface{} {
	switch paramType {
	case StringParam:
		return GetString(testParamName)
	case IntParam:
		return GetInt(testParamName)
	case BoolParam:
		return GetBool(testParamName)
	case UintParam:
		return GetUint(testParamName)
	case Float32Param:
		return GetFloat32(testParamName)
	case Float64Param:
		return GetFloat64(testParamName)
	case DurationParam:
		return GetDuration(testParamName)
	case StringSliceParam:
		return GetStringSlice(testParamName)
	case StringToStringParam:
		return GetStringToString(testParamName)
	case Int8Param:
		return GetInt8(testParamName)
	case Int16Param:
		return GetInt16(testParamName)
	case Int32Param:
		return GetInt32(testParamName)
	case Int64Param:
		return GetInt64(testParamName)
	case Uint8Param:
		return GetUint8(testParamName)
	case Uint16Param:
		return GetUint16(testParamName)
	case Uint32Param:
		return GetUint32(testParamName)
	case Uint64Param:
		return GetUint64(testParamName)
	default:
		suite.FailNow("Unsupported parameter type")
		return nil
	}
}

func (suite *testCommandConfigParams) TestParameters() {
	for _, tc := range getTestCases() {
		Reset()
		suite.Run(tc.name, func() {
			suite.WithCustom(tc.env, tc.args, func() {
				err := suite.addParam(newTestCommand(), tc.paramType, tc.paramDefaultVal).
					Execute()

				if suite.NoError(err, "Execute should not return error") {
					// Verify the parameter value
					value := suite.getValue(tc.paramType)
					suite.Equal(tc.expectedValue, value, "Parameter value should match expected")
				}
			})
			suite.WithCustom(tc.env, tc.shortArgs, func() {
				err := suite.addParamP(newTestCommand(), tc.paramType, tc.paramDefaultVal).
					Execute()

				if suite.NoError(err, "Execute should not return error") {
					// Verify the parameter value
					value := suite.getValue(tc.paramType)
					suite.Equal(tc.expectedValue, value, "Parameter value should match expected")
				}
			})
		})
	}
}

func (suite *testCommandConfigParams) TestParametersNoFlags() {
	suite.WithCustom(map[string]string{testParamEnv: "string-from-env"}, nil, func() {
		cmd := newTestCommand().withParam(testParamName, nil, emptyStr, emptyStr, testParamEnv, testParamDesc, "string-default-value")
		if suite.NoError(cmd.Execute()) {
			suite.Equal("string-from-env", GetString(testParamName))
		}
	})
}
