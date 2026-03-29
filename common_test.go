package krait

import (
	"time"
)

const (
	testParamName      = "test.param"
	testParamDesc      = "test parameter"
	testParamFlag      = "test-param"
	testParamEnv       = "TEST_PARAM"
	testParamShortFlag = "t"
)

// ParamType represents the type of parameter being tested
type ParamType int

const (
	StringParam ParamType = iota
	IntParam
	Int8Param
	Int16Param
	Int32Param
	Int64Param
	UintParam
	Uint8Param
	Uint16Param
	Uint32Param
	Uint64Param
	BoolParam
	Float32Param
	Float64Param
	DurationParam
	StringSliceParam
	StringToStringParam
)

// ParamTestCase represents a test case for testing parameter handling of any type
type ParamTestCase struct {
	name            string
	paramType       ParamType
	env             map[string]string
	args            []string
	shortArgs       []string
	paramDefaultVal interface{}
	expectedValue   interface{}
}

var (
	stringTestCases = []*ParamTestCase{
		{
			name:            "string default value",
			paramType:       StringParam,
			env:             map[string]string{},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: "default",
			expectedValue:   "default",
		},
		{
			name:      "string from environment",
			paramType: StringParam,
			env: map[string]string{
				testParamEnv: "env-value",
			},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: "default",
			expectedValue:   "env-value",
		},
		{
			name:            "string from flag",
			paramType:       StringParam,
			env:             map[string]string{},
			args:            []string{"--" + testParamFlag, "flag-value"},
			shortArgs:       []string{"-" + testParamShortFlag, "flag-value"},
			paramDefaultVal: "default",
			expectedValue:   "flag-value",
		},
		{
			name:      "string flag overrides environment",
			paramType: StringParam,
			env: map[string]string{
				testParamEnv: "env-value",
			},
			args:            []string{"--" + testParamFlag, "flag-value"},
			shortArgs:       []string{"-" + testParamShortFlag, "flag-value"},
			paramDefaultVal: "default",
			expectedValue:   "flag-value",
		},
	}

	intTestCases = []*ParamTestCase{
		{
			name:            "int default value",
			paramType:       IntParam,
			env:             map[string]string{},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: 42,
			expectedValue:   42,
		},
		{
			name:      "int from environment",
			paramType: IntParam,
			env: map[string]string{
				testParamEnv: "123",
			},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: 42,
			expectedValue:   123,
		},
		{
			name:            "int from flag",
			paramType:       IntParam,
			env:             map[string]string{},
			args:            []string{"--" + testParamFlag, "456"},
			shortArgs:       []string{"-" + testParamShortFlag, "456"},
			paramDefaultVal: 42,
			expectedValue:   456,
		},
		{
			name:      "int flag overrides environment",
			paramType: IntParam,
			env: map[string]string{
				testParamEnv: "123",
			},
			args:            []string{"--" + testParamFlag, "456"},
			shortArgs:       []string{"-" + testParamShortFlag, "456"},
			paramDefaultVal: 42,
			expectedValue:   456,
		},
	}

	boolTestCases = []*ParamTestCase{
		{
			name:            "bool default value",
			paramType:       BoolParam,
			env:             map[string]string{},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: false,
			expectedValue:   false,
		},
		{
			name:      "bool true from environment",
			paramType: BoolParam,
			env: map[string]string{
				testParamEnv: "true",
			},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: false,
			expectedValue:   true,
		},
		{
			name:      "bool false from environment",
			paramType: BoolParam,
			env: map[string]string{
				testParamEnv: "false",
			},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: false,
			expectedValue:   false,
		},
		{
			name:            "bool true from flag",
			paramType:       BoolParam,
			env:             map[string]string{},
			args:            []string{"--" + testParamFlag, "true"},
			shortArgs:       []string{"-" + testParamShortFlag, "true"},
			paramDefaultVal: false,
			expectedValue:   true,
		},
		{
			name:            "bool false from flag",
			paramType:       BoolParam,
			env:             map[string]string{},
			args:            []string{"--" + testParamFlag + "=false"},
			shortArgs:       []string{"-" + testParamShortFlag + "=false"},
			paramDefaultVal: false,
			expectedValue:   false,
		},
		{
			name:      "bool flag overrides environment",
			paramType: BoolParam,
			env: map[string]string{
				testParamEnv: "false",
			},
			args:            []string{"--" + testParamFlag, "true"},
			shortArgs:       []string{"-" + testParamShortFlag, "true"},
			paramDefaultVal: false,
			expectedValue:   true,
		},
	}

	uintTestCases = []*ParamTestCase{
		{
			name:            "uint default value",
			paramType:       UintParam,
			env:             map[string]string{},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: uint(42),
			expectedValue:   uint(42),
		},
		{
			name:      "uint from environment",
			paramType: UintParam,
			env: map[string]string{
				testParamEnv: "123",
			},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: uint(42),
			expectedValue:   uint(123),
		},
		{
			name:            "uint from flag",
			paramType:       UintParam,
			env:             map[string]string{},
			args:            []string{"--" + testParamFlag, "456"},
			shortArgs:       []string{"-" + testParamShortFlag, "456"},
			paramDefaultVal: uint(42),
			expectedValue:   uint(456),
		},
		{
			name:      "uint flag overrides environment",
			paramType: UintParam,
			env: map[string]string{
				testParamEnv: "123",
			},
			args:            []string{"--" + testParamFlag, "456"},
			shortArgs:       []string{"-" + testParamShortFlag, "456"},
			paramDefaultVal: uint(42),
			expectedValue:   uint(456),
		},
	}

	float32TestCases = []*ParamTestCase{
		{
			name:            "float32 default value",
			paramType:       Float32Param,
			env:             map[string]string{},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: float32(42.5),
			expectedValue:   float32(42.5),
		},
		{
			name:      "float32 from environment",
			paramType: Float32Param,
			env: map[string]string{
				testParamEnv: "123.45",
			},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: float32(42.5),
			expectedValue:   float32(123.45),
		},
		{
			name:            "float32 from flag",
			paramType:       Float32Param,
			env:             map[string]string{},
			args:            []string{"--" + testParamFlag, "456.78"},
			shortArgs:       []string{"-" + testParamShortFlag, "456.78"},
			paramDefaultVal: float32(42.5),
			expectedValue:   float32(456.78),
		},
		{
			name:      "float32 flag overrides environment",
			paramType: Float32Param,
			env: map[string]string{
				testParamEnv: "123.45",
			},
			args:            []string{"--" + testParamFlag, "456.78"},
			shortArgs:       []string{"-" + testParamShortFlag, "456.78"},
			paramDefaultVal: float32(42.5),
			expectedValue:   float32(456.78),
		},
	}

	float64TestCases = []*ParamTestCase{
		{
			name:            "float64 default value",
			paramType:       Float64Param,
			env:             map[string]string{},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: float64(42.5),
			expectedValue:   float64(42.5),
		},
		{
			name:      "float64 from environment",
			paramType: Float64Param,
			env: map[string]string{
				testParamEnv: "123.45",
			},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: float64(42.5),
			expectedValue:   float64(123.45),
		},
		{
			name:            "float64 from flag",
			paramType:       Float64Param,
			env:             map[string]string{},
			args:            []string{"--" + testParamFlag, "456.78"},
			shortArgs:       []string{"-" + testParamShortFlag, "456.78"},
			paramDefaultVal: float64(42.5),
			expectedValue:   float64(456.78),
		},
		{
			name:      "float64 flag overrides environment",
			paramType: Float64Param,
			env: map[string]string{
				testParamEnv: "123.45",
			},
			args:            []string{"--" + testParamFlag, "456.78"},
			shortArgs:       []string{"-" + testParamShortFlag, "456.78"},
			paramDefaultVal: float64(42.5),
			expectedValue:   float64(456.78),
		},
	}

	durationTestCases = []*ParamTestCase{
		{
			name:            "duration default value",
			paramType:       DurationParam,
			env:             map[string]string{},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: time.Duration(42 * time.Second),
			expectedValue:   42 * time.Second,
		},
		{
			name:      "duration from environment",
			paramType: DurationParam,
			env: map[string]string{
				testParamEnv: "123s",
			},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: 42 * time.Second,
			expectedValue:   123 * time.Second,
		},
		{
			name:            "duration from flag",
			paramType:       DurationParam,
			env:             map[string]string{},
			args:            []string{"--" + testParamFlag, "456s"},
			shortArgs:       []string{"-" + testParamShortFlag, "456s"},
			paramDefaultVal: 42 * time.Second,
			expectedValue:   456 * time.Second,
		},
		{
			name:      "duration flag overrides environment",
			paramType: DurationParam,
			env: map[string]string{
				testParamEnv: "123s",
			},
			args:            []string{"--" + testParamFlag, "456s"},
			shortArgs:       []string{"-" + testParamShortFlag, "456s"},
			paramDefaultVal: 42 * time.Second,
			expectedValue:   456 * time.Second,
		},
	}

	// TODO: add support for string slice from environment
	stringSliceTestCases = []*ParamTestCase{
		{
			name:            "string slice default value",
			paramType:       StringSliceParam,
			env:             map[string]string{},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: []string{"default1", "default2"},
			expectedValue:   []string{"default1", "default2"},
		},
		{
			name:            "string slice from flag (comma separated)",
			paramType:       StringSliceParam,
			env:             map[string]string{},
			args:            []string{"--" + testParamFlag, "flag1,flag2,flag3"},
			shortArgs:       []string{"-" + testParamShortFlag, "flag1,flag2,flag3"},
			paramDefaultVal: []string{"default1", "default2"},
			expectedValue:   []string{"flag1", "flag2", "flag3"},
		},
		{
			name:            "string slice from flag (multiple flags)",
			paramType:       StringSliceParam,
			env:             map[string]string{},
			args:            []string{"--" + testParamFlag, "flag1", "--" + testParamFlag, "flag2", "--" + testParamFlag, "flag3"},
			shortArgs:       []string{"-" + testParamShortFlag, "flag1", "-" + testParamShortFlag, "flag2", "-" + testParamShortFlag, "flag3"},
			paramDefaultVal: []string{"default1", "default2"},
			expectedValue:   []string{"flag1", "flag2", "flag3"},
		},
	}

	// TODO: add support for string to string from environment
	stringToStringTestCases = []*ParamTestCase{
		{
			name:            "string to string default value",
			paramType:       StringToStringParam,
			env:             map[string]string{},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: map[string]string{"key1": "val1", "key2": "val2"},
			expectedValue:   map[string]string{"key1": "val1", "key2": "val2"},
		},
		{
			name:            "string to string from flag",
			paramType:       StringToStringParam,
			env:             map[string]string{},
			args:            []string{"--" + testParamFlag, "key1=val1,key2=val2,key3=val3"},
			shortArgs:       []string{"-" + testParamShortFlag, "key1=val1,key2=val2,key3=val3"},
			paramDefaultVal: map[string]string{"default": "value"},
			expectedValue:   map[string]string{"key1": "val1", "key2": "val2", "key3": "val3"},
		},
		{
			name:            "string to string multiple flag values",
			paramType:       StringToStringParam,
			env:             map[string]string{},
			args:            []string{"--" + testParamFlag, "key1=val1", "--" + testParamFlag, "key2=val2"},
			shortArgs:       []string{"-" + testParamShortFlag, "key1=val1", "-" + testParamShortFlag, "key2=val2"},
			paramDefaultVal: map[string]string{"default": "value"},
			expectedValue:   map[string]string{"key1": "val1", "key2": "val2"},
		},
	}

	int8TestCases = []*ParamTestCase{
		{
			name:            "int8 default value",
			paramType:       Int8Param,
			env:             map[string]string{},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: int8(42),
			expectedValue:   int8(42),
		},
		{
			name:      "int8 from environment",
			paramType: Int8Param,
			env: map[string]string{
				testParamEnv: "123",
			},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: int8(42),
			expectedValue:   int8(123),
		},
		{
			name:            "int8 from flag",
			paramType:       Int8Param,
			env:             map[string]string{},
			args:            []string{"--" + testParamFlag, "100"},
			shortArgs:       []string{"-" + testParamShortFlag, "100"},
			paramDefaultVal: int8(42),
			expectedValue:   int8(100),
		},
		{
			name:      "int8 flag overrides environment",
			paramType: Int8Param,
			env: map[string]string{
				testParamEnv: "123",
			},
			args:            []string{"--" + testParamFlag, "100"},
			shortArgs:       []string{"-" + testParamShortFlag, "100"},
			paramDefaultVal: int8(42),
			expectedValue:   int8(100),
		},
	}

	int16TestCases = []*ParamTestCase{
		{
			name:            "int16 default value",
			paramType:       Int16Param,
			env:             map[string]string{},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: int16(42),
			expectedValue:   int16(42),
		},
		{
			name:      "int16 from environment",
			paramType: Int16Param,
			env: map[string]string{
				testParamEnv: "12345",
			},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: int16(42),
			expectedValue:   int16(12345),
		},
		{
			name:            "int16 from flag",
			paramType:       Int16Param,
			env:             map[string]string{},
			args:            []string{"--" + testParamFlag, "30000"},
			shortArgs:       []string{"-" + testParamShortFlag, "30000"},
			paramDefaultVal: int16(42),
			expectedValue:   int16(30000),
		},
		{
			name:      "int16 flag overrides environment",
			paramType: Int16Param,
			env: map[string]string{
				testParamEnv: "12345",
			},
			args:            []string{"--" + testParamFlag, "30000"},
			shortArgs:       []string{"-" + testParamShortFlag, "30000"},
			paramDefaultVal: int16(42),
			expectedValue:   int16(30000),
		},
	}

	int32TestCases = []*ParamTestCase{
		{
			name:            "int32 default value",
			paramType:       Int32Param,
			env:             map[string]string{},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: int32(42),
			expectedValue:   int32(42),
		},
		{
			name:      "int32 from environment",
			paramType: Int32Param,
			env: map[string]string{
				testParamEnv: "1234567",
			},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: int32(42),
			expectedValue:   int32(1234567),
		},
		{
			name:            "int32 from flag",
			paramType:       Int32Param,
			env:             map[string]string{},
			args:            []string{"--" + testParamFlag, "2000000"},
			shortArgs:       []string{"-" + testParamShortFlag, "2000000"},
			paramDefaultVal: int32(42),
			expectedValue:   int32(2000000),
		},
		{
			name:      "int32 flag overrides environment",
			paramType: Int32Param,
			env: map[string]string{
				testParamEnv: "1234567",
			},
			args:            []string{"--" + testParamFlag, "2000000"},
			shortArgs:       []string{"-" + testParamShortFlag, "2000000"},
			paramDefaultVal: int32(42),
			expectedValue:   int32(2000000),
		},
	}

	int64TestCases = []*ParamTestCase{
		{
			name:            "int64 default value",
			paramType:       Int64Param,
			env:             map[string]string{},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: int64(42),
			expectedValue:   int64(42),
		},
		{
			name:      "int64 from environment",
			paramType: Int64Param,
			env: map[string]string{
				testParamEnv: "1234567890",
			},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: int64(42),
			expectedValue:   int64(1234567890),
		},
		{
			name:            "int64 from flag",
			paramType:       Int64Param,
			env:             map[string]string{},
			args:            []string{"--" + testParamFlag, "9876543210"},
			shortArgs:       []string{"-" + testParamShortFlag, "9876543210"},
			paramDefaultVal: int64(42),
			expectedValue:   int64(9876543210),
		},
		{
			name:      "int64 flag overrides environment",
			paramType: Int64Param,
			env: map[string]string{
				testParamEnv: "1234567890",
			},
			args:            []string{"--" + testParamFlag, "9876543210"},
			shortArgs:       []string{"-" + testParamShortFlag, "9876543210"},
			paramDefaultVal: int64(42),
			expectedValue:   int64(9876543210),
		},
	}

	uint8TestCases = []*ParamTestCase{
		{
			name:            "uint8 default value",
			paramType:       Uint8Param,
			env:             map[string]string{},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: uint8(42),
			expectedValue:   uint8(42),
		},
		{
			name:      "uint8 from environment",
			paramType: Uint8Param,
			env: map[string]string{
				testParamEnv: "200",
			},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: uint8(42),
			expectedValue:   uint8(200),
		},
		{
			name:            "uint8 from flag",
			paramType:       Uint8Param,
			env:             map[string]string{},
			args:            []string{"--" + testParamFlag, "250"},
			shortArgs:       []string{"-" + testParamShortFlag, "250"},
			paramDefaultVal: uint8(42),
			expectedValue:   uint8(250),
		},
		{
			name:      "uint8 flag overrides environment",
			paramType: Uint8Param,
			env: map[string]string{
				testParamEnv: "200",
			},
			args:            []string{"--" + testParamFlag, "250"},
			shortArgs:       []string{"-" + testParamShortFlag, "250"},
			paramDefaultVal: uint8(42),
			expectedValue:   uint8(250),
		},
	}

	uint16TestCases = []*ParamTestCase{
		{
			name:            "uint16 default value",
			paramType:       Uint16Param,
			env:             map[string]string{},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: uint16(42),
			expectedValue:   uint16(42),
		},
		{
			name:      "uint16 from environment",
			paramType: Uint16Param,
			env: map[string]string{
				testParamEnv: "45000",
			},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: uint16(42),
			expectedValue:   uint16(45000),
		},
		{
			name:            "uint16 from flag",
			paramType:       Uint16Param,
			env:             map[string]string{},
			args:            []string{"--" + testParamFlag, "60000"},
			shortArgs:       []string{"-" + testParamShortFlag, "60000"},
			paramDefaultVal: uint16(42),
			expectedValue:   uint16(60000),
		},
		{
			name:      "uint16 flag overrides environment",
			paramType: Uint16Param,
			env: map[string]string{
				testParamEnv: "45000",
			},
			args:            []string{"--" + testParamFlag, "60000"},
			shortArgs:       []string{"-" + testParamShortFlag, "60000"},
			paramDefaultVal: uint16(42),
			expectedValue:   uint16(60000),
		},
	}

	uint32TestCases = []*ParamTestCase{
		{
			name:            "uint32 default value",
			paramType:       Uint32Param,
			env:             map[string]string{},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: uint32(42),
			expectedValue:   uint32(42),
		},
		{
			name:      "uint32 from environment",
			paramType: Uint32Param,
			env: map[string]string{
				testParamEnv: "3000000000",
			},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: uint32(42),
			expectedValue:   uint32(3000000000),
		},
		{
			name:            "uint32 from flag",
			paramType:       Uint32Param,
			env:             map[string]string{},
			args:            []string{"--" + testParamFlag, "4000000000"},
			shortArgs:       []string{"-" + testParamShortFlag, "4000000000"},
			paramDefaultVal: uint32(42),
			expectedValue:   uint32(4000000000),
		},
		{
			name:      "uint32 flag overrides environment",
			paramType: Uint32Param,
			env: map[string]string{
				testParamEnv: "3000000000",
			},
			args:            []string{"--" + testParamFlag, "4000000000"},
			shortArgs:       []string{"-" + testParamShortFlag, "4000000000"},
			paramDefaultVal: uint32(42),
			expectedValue:   uint32(4000000000),
		},
	}

	uint64TestCases = []*ParamTestCase{
		{
			name:            "uint64 default value",
			paramType:       Uint64Param,
			env:             map[string]string{},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: uint64(42),
			expectedValue:   uint64(42),
		},
		{
			name:      "uint64 from environment",
			paramType: Uint64Param,
			env: map[string]string{
				testParamEnv: "18000000000000000000",
			},
			args:            []string{},
			shortArgs:       []string{},
			paramDefaultVal: uint64(42),
			expectedValue:   uint64(18000000000000000000),
		},
		{
			name:            "uint64 from flag",
			paramType:       Uint64Param,
			env:             map[string]string{},
			args:            []string{"--" + testParamFlag, "18446744073709551615"},
			shortArgs:       []string{"-" + testParamShortFlag, "18446744073709551615"},
			paramDefaultVal: uint64(42),
			expectedValue:   uint64(18446744073709551615),
		},
		{
			name:      "uint64 flag overrides environment",
			paramType: Uint64Param,
			env: map[string]string{
				testParamEnv: "18000000000000000000",
			},
			args:            []string{"--" + testParamFlag, "18446744073709551615"},
			shortArgs:       []string{"-" + testParamShortFlag, "18446744073709551615"},
			paramDefaultVal: uint64(42),
			expectedValue:   uint64(18446744073709551615),
		},
	}
)

func getTestCases() []*ParamTestCase {
	testCases := []*ParamTestCase{}
	testCases = append(testCases, stringTestCases...)
	testCases = append(testCases, intTestCases...)
	testCases = append(testCases, int8TestCases...)
	testCases = append(testCases, int16TestCases...)
	testCases = append(testCases, int32TestCases...)
	testCases = append(testCases, int64TestCases...)
	testCases = append(testCases, uintTestCases...)
	testCases = append(testCases, uint8TestCases...)
	testCases = append(testCases, uint16TestCases...)
	testCases = append(testCases, uint32TestCases...)
	testCases = append(testCases, uint64TestCases...)
	testCases = append(testCases, boolTestCases...)
	testCases = append(testCases, float32TestCases...)
	testCases = append(testCases, float64TestCases...)
	testCases = append(testCases, durationTestCases...)
	testCases = append(testCases, stringSliceTestCases...)
	testCases = append(testCases, stringToStringTestCases...)
	return testCases
}

// newTestCommand creates a new command with an empty runner function for testing
func newTestCommand() *Command {
	return New("app", "test app", "test app long description").
		WithRun(func(args []string) error {
			return nil
		})
}
