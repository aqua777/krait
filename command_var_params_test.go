package krait

import (
	"fmt"
	"time"

	"github.com/aqua777/krait/testing"
)

// testCommandVarParams represents a test suite for testing variable parameter handling
type testCommandVarParams struct {
	testing.Suite
	expectedError error
}

func TestCommandVarParams(t *testing.T) {
	testing.Run(t, new(testCommandVarParams))
}

func (suite *testCommandVarParams) SetupTest() {
	Reset() // Reset global state before each test
	suite.expectedError = fmt.Errorf("test error")
}

func (suite *testCommandVarParams) TearDownTest() {
}

// addVar adds the appropriate variable parameter type to the command
func (suite *testCommandVarParams) addVar(cmd *Command, paramType ParamType, varPtr interface{}, defaultValue any) *Command {
	switch paramType {
	case StringParam:
		return cmd.WithStringVar(varPtr.(*string), testParamDesc, testParamFlag, testParamEnv, defaultValue.(string))
	case IntParam:
		return cmd.WithIntVar(varPtr.(*int), testParamDesc, testParamFlag, testParamEnv, defaultValue.(int))
	case BoolParam:
		return cmd.WithBoolVar(varPtr.(*bool), testParamDesc, testParamFlag, testParamEnv, defaultValue.(bool))
	case UintParam:
		return cmd.WithUintVar(varPtr.(*uint), testParamDesc, testParamFlag, testParamEnv, defaultValue.(uint))
	case Float32Param:
		return cmd.WithFloat32Var(varPtr.(*float32), testParamDesc, testParamFlag, testParamEnv, defaultValue.(float32))
	case Float64Param:
		return cmd.WithFloat64Var(varPtr.(*float64), testParamDesc, testParamFlag, testParamEnv, defaultValue.(float64))
	case DurationParam:
		return cmd.WithDurationVar(varPtr.(*time.Duration), testParamDesc, testParamFlag, testParamEnv, defaultValue.(time.Duration))
	case StringSliceParam:
		return cmd.WithStringSliceVar(varPtr.(*[]string), testParamDesc, testParamFlag, testParamEnv, defaultValue.([]string))
	case StringToStringParam:
		return cmd.WithStringToStringVar(varPtr.(*map[string]string), testParamDesc, testParamFlag, testParamEnv, defaultValue.(map[string]string))
	case Int8Param:
		return cmd.WithInt8Var(varPtr.(*int8), testParamDesc, testParamFlag, testParamEnv, defaultValue.(int8))
	case Int16Param:
		return cmd.WithInt16Var(varPtr.(*int16), testParamDesc, testParamFlag, testParamEnv, defaultValue.(int16))
	case Int32Param:
		return cmd.WithInt32Var(varPtr.(*int32), testParamDesc, testParamFlag, testParamEnv, defaultValue.(int32))
	case Int64Param:
		return cmd.WithInt64Var(varPtr.(*int64), testParamDesc, testParamFlag, testParamEnv, defaultValue.(int64))
	case Uint8Param:
		return cmd.WithUint8Var(varPtr.(*uint8), testParamDesc, testParamFlag, testParamEnv, defaultValue.(uint8))
	case Uint16Param:
		return cmd.WithUint16Var(varPtr.(*uint16), testParamDesc, testParamFlag, testParamEnv, defaultValue.(uint16))
	case Uint32Param:
		return cmd.WithUint32Var(varPtr.(*uint32), testParamDesc, testParamFlag, testParamEnv, defaultValue.(uint32))
	case Uint64Param:
		return cmd.WithUint64Var(varPtr.(*uint64), testParamDesc, testParamFlag, testParamEnv, defaultValue.(uint64))
	default:
		suite.FailNow("Unsupported parameter type")
		return nil
	}
}

// addVarP adds the appropriate variable parameter type to the command with a shorthand flag
func (suite *testCommandVarParams) addVarP(cmd *Command, paramType ParamType, varPtr interface{}, defaultValue any) *Command {
	switch paramType {
	case StringParam:
		return cmd.WithStringVarP(varPtr.(*string), testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultValue.(string))
	case IntParam:
		return cmd.WithIntVarP(varPtr.(*int), testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultValue.(int))
	case BoolParam:
		return cmd.WithBoolVarP(varPtr.(*bool), testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultValue.(bool))
	case UintParam:
		return cmd.WithUintVarP(varPtr.(*uint), testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultValue.(uint))
	case Float32Param:
		return cmd.WithFloat32VarP(varPtr.(*float32), testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultValue.(float32))
	case Float64Param:
		return cmd.WithFloat64VarP(varPtr.(*float64), testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultValue.(float64))
	case DurationParam:
		return cmd.WithDurationVarP(varPtr.(*time.Duration), testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultValue.(time.Duration))
	case StringSliceParam:
		return cmd.WithStringSliceVarP(varPtr.(*[]string), testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultValue.([]string))
	case StringToStringParam:
		return cmd.WithStringToStringVarP(varPtr.(*map[string]string), testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultValue.(map[string]string))
	case Int8Param:
		return cmd.WithInt8VarP(varPtr.(*int8), testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultValue.(int8))
	case Int16Param:
		return cmd.WithInt16VarP(varPtr.(*int16), testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultValue.(int16))
	case Int32Param:
		return cmd.WithInt32VarP(varPtr.(*int32), testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultValue.(int32))
	case Int64Param:
		return cmd.WithInt64VarP(varPtr.(*int64), testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultValue.(int64))
	case Uint8Param:
		return cmd.WithUint8VarP(varPtr.(*uint8), testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultValue.(uint8))
	case Uint16Param:
		return cmd.WithUint16VarP(varPtr.(*uint16), testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultValue.(uint16))
	case Uint32Param:
		return cmd.WithUint32VarP(varPtr.(*uint32), testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultValue.(uint32))
	case Uint64Param:
		return cmd.WithUint64VarP(varPtr.(*uint64), testParamDesc, testParamFlag, testParamShortFlag, testParamEnv, defaultValue.(uint64))
	default:
		suite.FailNow("Unsupported parameter type")
		return nil
	}
}

func (suite *testCommandVarParams) TestParameters() {
	for _, tc := range getTestCases() {
		// skip test cases that have env variables
		if len(tc.env) > 0 {
			continue
		}

		suite.Run(tc.name, func() {
			Reset()
			suite.WithCustom(nil, tc.args, func() {
				varPtr := suite.getVarPtr(tc.paramType, tc.paramDefaultVal)
				err := suite.addVar(newTestCommand(), tc.paramType, varPtr, tc.paramDefaultVal).Execute()

				if suite.NoError(err, "Execute should not return error") {
					// Verify the parameter value based on type
					suite.assertEqual(tc.paramType, tc.expectedValue, varPtr)
				}
			})

			Reset()
			suite.WithCustom(nil, tc.shortArgs, func() {
				varPtr := suite.getVarPtr(tc.paramType, tc.paramDefaultVal)
				err := suite.addVarP(newTestCommand(), tc.paramType, varPtr, tc.paramDefaultVal).Execute()

				if suite.NoError(err, "Execute should not return error") {
					// Verify the parameter value based on type
					suite.assertEqual(tc.paramType, tc.expectedValue, varPtr)
				}
			})
		})
	}
}

// TestParametersFromEnv tests that var-bound parameters are correctly populated
// from environment variables, and that CLI flags still override env vars.
func (suite *testCommandVarParams) TestParametersFromEnv() {
	for _, tc := range getTestCases() {
		if len(tc.env) == 0 {
			continue
		}

		suite.Run(tc.name, func() {
			Reset()
			suite.WithCustom(tc.env, tc.args, func() {
				varPtr := suite.getVarPtr(tc.paramType, tc.paramDefaultVal)
				err := suite.addVar(newTestCommand(), tc.paramType, varPtr, tc.paramDefaultVal).Execute()

				if suite.NoError(err, "Execute should not return error") {
					suite.assertEqual(tc.paramType, tc.expectedValue, varPtr)
				}
			})
		})
	}
}

func (suite *testCommandVarParams) getVarPtr(paramType ParamType, paramDefaultVal interface{}) interface{} {
	var varPtr interface{}
	switch paramType {
	case StringParam:
		v := paramDefaultVal.(string)
		varPtr = &v
	case IntParam:
		v := paramDefaultVal.(int)
		varPtr = &v
	case Int8Param:
		v := paramDefaultVal.(int8)
		varPtr = &v
	case Int16Param:
		v := paramDefaultVal.(int16)
		varPtr = &v
	case Int32Param:
		v := paramDefaultVal.(int32)
		varPtr = &v
	case Int64Param:
		v := paramDefaultVal.(int64)
		varPtr = &v
	case UintParam:
		v := paramDefaultVal.(uint)
		varPtr = &v
	case Uint8Param:
		v := paramDefaultVal.(uint8)
		varPtr = &v
	case Uint16Param:
		v := paramDefaultVal.(uint16)
		varPtr = &v
	case Uint32Param:
		v := paramDefaultVal.(uint32)
		varPtr = &v
	case Uint64Param:
		v := paramDefaultVal.(uint64)
		varPtr = &v
	case BoolParam:
		v := paramDefaultVal.(bool)
		varPtr = &v
	case Float32Param:
		v := paramDefaultVal.(float32)
		varPtr = &v
	case Float64Param:
		v := paramDefaultVal.(float64)
		varPtr = &v
	case DurationParam:
		v := paramDefaultVal.(time.Duration)
		varPtr = &v
	case StringSliceParam:
		v := paramDefaultVal.([]string)
		varPtr = &v
	case StringToStringParam:
		v := paramDefaultVal.(map[string]string)
		varPtr = &v
	}
	return varPtr
}

func (suite *testCommandVarParams) assertEqual(paramType, expectedValue interface{}, varPtr interface{}) {
	switch paramType {
	case StringParam:
		suite.Equal(expectedValue, *(varPtr.(*string)))
	case IntParam:
		suite.Equal(expectedValue, *(varPtr.(*int)))
	case Int8Param:
		suite.Equal(expectedValue, *(varPtr.(*int8)))
	case Int16Param:
		suite.Equal(expectedValue, *(varPtr.(*int16)))
	case Int32Param:
		suite.Equal(expectedValue, *(varPtr.(*int32)))
	case Int64Param:
		suite.Equal(expectedValue, *(varPtr.(*int64)))
	case UintParam:
		suite.Equal(expectedValue, *(varPtr.(*uint)))
	case Uint8Param:
		suite.Equal(expectedValue, *(varPtr.(*uint8)))
	case Uint16Param:
		suite.Equal(expectedValue, *(varPtr.(*uint16)))
	case Uint32Param:
		suite.Equal(expectedValue, *(varPtr.(*uint32)))
	case Uint64Param:
		suite.Equal(expectedValue, *(varPtr.(*uint64)))
	case BoolParam:
		suite.Equal(expectedValue, *(varPtr.(*bool)))
	case Float32Param:
		suite.Equal(expectedValue, *(varPtr.(*float32)))
	case Float64Param:
		suite.Equal(expectedValue, *(varPtr.(*float64)))
	case DurationParam:
		suite.Equal(expectedValue, *(varPtr.(*time.Duration)))
	case StringSliceParam:
		suite.Equal(expectedValue, *(varPtr.(*[]string)))
	case StringToStringParam:
		suite.Equal(expectedValue, *(varPtr.(*map[string]string)))
	}
}
