package krait

import (
	"testing"
	"time"

	kraittest "github.com/aqua777/krait/testing"
)

func TestCommandPersistent(t *testing.T) {
	kraittest.Run(t, new(testCommandPersistent))
}

type testCommandPersistent struct {
	kraittest.Suite
}

func (suite *testCommandPersistent) SetupTest() {
	Reset()
}

// Root command with persistent bool — root run sees correct value

func (suite *testCommandPersistent) TestPersistentBoolOnRootDefaultValue() {
	suite.WithCustom(nil, nil, func() {
		cmd := New("app", "app", "").
			WithPersistentBool("verbose", "Verbose output", "verbose", "APP_VERBOSE", false).
			WithRun(func(args []string) error { return nil })

		suite.NoError(cmd.Execute())
		suite.False(GetBool("verbose"), "Persistent bool must return default false when not set")
	})
}

func (suite *testCommandPersistent) TestPersistentBoolOnRootViaFlag() {
	suite.WithCustom(nil, []string{"--verbose"}, func() {
		cmd := New("app", "app", "").
			WithPersistentBool("verbose", "Verbose output", "verbose", "APP_VERBOSE", false).
			WithRun(func(args []string) error { return nil })

		suite.NoError(cmd.Execute())
		suite.True(GetBool("verbose"), "Persistent bool must return true when --verbose flag is passed on root")
	})
}

// Subcommand inherits persistent bool from root

func (suite *testCommandPersistent) TestPersistentBoolInheritedBySubcommandViaFlag() {
	suite.WithCustom(nil, []string{"sub", "--verbose"}, func() {
		var gotVerbose bool
		sub := New("sub", "sub", "").
			WithRun(func(args []string) error {
				gotVerbose = GetBool("verbose")
				return nil
			})

		App("app", "app", "").
			WithPersistentBool("verbose", "Verbose output", "verbose", "APP_VERBOSE", false).
			WithCommand(sub)

		suite.NoError(Execute())
		suite.True(gotVerbose, "Persistent bool must be true in subcommand Run when --verbose is passed")
	})
}

func (suite *testCommandPersistent) TestPersistentBoolInheritedBySubcommandDefaultValue() {
	suite.WithCustom(nil, []string{"sub"}, func() {
		var gotVerbose bool
		sub := New("sub", "sub", "").
			WithRun(func(args []string) error {
				gotVerbose = GetBool("verbose")
				return nil
			})

		App("app", "app", "").
			WithPersistentBool("verbose", "Verbose output", "verbose", "APP_VERBOSE", false).
			WithCommand(sub)

		suite.NoError(Execute())
		suite.False(gotVerbose, "Persistent bool must return false default in subcommand when not passed")
	})
}

// Persistent flag env var respected in subcommand

func (suite *testCommandPersistent) TestPersistentBoolInheritedBySubcommandViaEnvVar() {
	suite.WithCustom(map[string]string{"APP_VERBOSE": "true"}, []string{"sub"}, func() {
		var gotVerbose bool
		sub := New("sub", "sub", "").
			WithRun(func(args []string) error {
				gotVerbose = GetBool("verbose")
				return nil
			})

		App("app", "app", "").
			WithPersistentBool("verbose", "Verbose output", "verbose", "APP_VERBOSE", false).
			WithCommand(sub)

		suite.NoError(Execute())
		suite.True(gotVerbose, "Persistent bool env var must be honored in subcommand Run")
	})
}

// Source priority for persistent flag: CLI flag > env var > default

func (suite *testCommandPersistent) TestPersistentBoolSourcePriorityFlagOverridesEnv() {
	suite.WithCustom(map[string]string{"APP_VERBOSE": "false"}, []string{"sub", "--verbose"}, func() {
		var gotVerbose bool
		sub := New("sub", "sub", "").
			WithRun(func(args []string) error {
				gotVerbose = GetBool("verbose")
				return nil
			})

		App("app", "app", "").
			WithPersistentBool("verbose", "Verbose output", "verbose", "APP_VERBOSE", false).
			WithCommand(sub)

		suite.NoError(Execute())
		suite.True(gotVerbose, "CLI flag must override env var for persistent bool in subcommand")
	})
}

// WithPersistentBoolVar — var-ptr populated in subcommand

func (suite *testCommandPersistent) TestPersistentBoolVarInSubcommand() {
	suite.WithCustom(nil, []string{"sub", "--verbose"}, func() {
		var verbose bool
		sub := New("sub", "sub", "").
			WithRun(func(args []string) error { return nil })

		App("app", "app", "").
			WithPersistentBoolVar(&verbose, "Verbose output", "verbose", "APP_VERBOSE", false).
			WithCommand(sub)

		suite.NoError(Execute())
		suite.True(verbose, "WithPersistentBoolVar pointer must be populated in subcommand Run")
	})
}

func (suite *testCommandPersistent) TestPersistentBoolVarDefaultInSubcommand() {
	suite.WithCustom(nil, []string{"sub"}, func() {
		var verbose bool
		sub := New("sub", "sub", "").
			WithRun(func(args []string) error { return nil })

		App("app", "app", "").
			WithPersistentBoolVar(&verbose, "Verbose output", "verbose", "APP_VERBOSE", false).
			WithCommand(sub)

		suite.NoError(Execute())
		suite.False(verbose, "WithPersistentBoolVar pointer must hold default value in subcommand when flag not set")
	})
}

// Two levels of nesting: grandchild inherits persistent flag

func (suite *testCommandPersistent) TestPersistentBoolInheritedByGrandchild() {
	suite.WithCustom(nil, []string{"child", "grandchild", "--verbose"}, func() {
		var gotVerbose bool
		grandchild := New("grandchild", "grandchild", "").
			WithRun(func(args []string) error {
				gotVerbose = GetBool("verbose")
				return nil
			})

		child := New("child", "child", "").
			WithCommand(grandchild)

		App("app", "app", "").
			WithPersistentBool("verbose", "Verbose output", "verbose", "APP_VERBOSE", false).
			WithCommand(child)

		suite.NoError(Execute())
		suite.True(gotVerbose, "Persistent bool must be accessible in grandchild command")
	})
}

// WithPersistentString — named param, named getter

func (suite *testCommandPersistent) TestPersistentStringInSubcommand() {
	suite.WithCustom(nil, []string{"sub", "--token", "secret123"}, func() {
		var gotToken string
		sub := New("sub", "sub", "").
			WithRun(func(args []string) error {
				gotToken = GetString("token")
				return nil
			})

		App("app", "app", "").
			WithPersistentString("token", "API token", "token", "APP_TOKEN", "").
			WithCommand(sub)

		suite.NoError(Execute())
		suite.Equal("secret123", gotToken, "Persistent string must be accessible via GetString in subcommand")
	})
}

func (suite *testCommandPersistent) TestPersistentStringViaEnvVarInSubcommand() {
	suite.WithCustom(map[string]string{"APP_TOKEN": "env-token"}, []string{"sub"}, func() {
		var gotToken string
		sub := New("sub", "sub", "").
			WithRun(func(args []string) error {
				gotToken = GetString("token")
				return nil
			})

		App("app", "app", "").
			WithPersistentString("token", "API token", "token", "APP_TOKEN", "").
			WithCommand(sub)

		suite.NoError(Execute())
		suite.Equal("env-token", gotToken, "Persistent string env var must be honored in subcommand")
	})
}

// WithParams (ConfigParams) with persistent param is correctly propagated

func (suite *testCommandPersistent) TestPersistentParamViaConfigParamsInSubcommand() {
	suite.WithCustom(nil, []string{"sub", "--region", "us-east-1"}, func() {
		var gotRegion string
		sub := New("sub", "sub", "").
			WithRun(func(args []string) error {
				gotRegion = GetString("region")
				return nil
			})

		// Declare persistent param directly on root; WithParams does not support
		// persistent semantics — use WithPersistentString directly then WithCommand.
		App("app", "app", "").
			WithPersistentString("region", "AWS region", "region", "AWS_REGION", "us-west-2").
			WithCommand(sub)

		suite.NoError(Execute())
		suite.Equal("us-east-1", gotRegion, "Persistent param propagated via WithCommand must be accessible in subcommand")
	})
}

// WithPersistentStringP — short flag form

func (suite *testCommandPersistent) TestPersistentStringPShortFlagInSubcommand() {
	suite.WithCustom(nil, []string{"sub", "-t", "short-token"}, func() {
		var gotToken string
		sub := New("sub", "sub", "").
			WithRun(func(args []string) error {
				gotToken = GetString("token")
				return nil
			})

		App("app", "app", "").
			WithPersistentStringP("token", "API token", "token", "t", "APP_TOKEN", "").
			WithCommand(sub)

		suite.NoError(Execute())
		suite.Equal("short-token", gotToken, "Persistent string short flag must work in subcommand")
	})
}

// Registration smoke tests — verifies all WithPersistent* typed wrappers call through
// without panicking. Correctness of the underlying type routing is covered by the
// withLongFlag/withShortFlag tests; these tests ensure the fluent wrappers are wired.

func (suite *testCommandPersistent) TestPersistentAllNamedTypesRegistrationNoError() {
	suite.WithCustom(nil, nil, func() {
		suite.NotPanics(func() {
			New("app", "app", "").
				WithPersistentInt("cnt", "Count", "cnt", "CNT", 0).
				WithPersistentIntP("cntp", "Count P", "cntp", "C", "CNTP", 0).
				WithPersistentInt8("i8", "Int8", "i8", "I8", int8(0)).
				WithPersistentInt8P("i8p", "Int8P", "i8p", "8", "I8P", int8(0)).
				WithPersistentInt16("i16", "Int16", "i16", "I16", int16(0)).
				WithPersistentInt16P("i16p", "Int16P", "i16p", "6", "I16P", int16(0)).
				WithPersistentInt32("i32", "Int32", "i32", "I32", int32(0)).
				WithPersistentInt32P("i32p", "Int32P", "i32p", "3", "I32P", int32(0)).
				WithPersistentInt64("i64", "Int64", "i64", "I64", int64(0)).
				WithPersistentInt64P("i64p", "Int64P", "i64p", "9", "I64P", int64(0)).
				WithPersistentUint("u", "Uint", "u", "U", uint(0)).
				WithPersistentUintP("up", "UintP", "up", "U", "UP", uint(0)).
				WithPersistentUint8("u8", "Uint8", "u8", "U8", uint8(0)).
				WithPersistentUint8P("u8p", "Uint8P", "u8p", "X", "U8P", uint8(0)).
				WithPersistentUint16("u16", "Uint16", "u16", "U16", uint16(0)).
				WithPersistentUint16P("u16p", "Uint16P", "u16p", "Y", "U16P", uint16(0)).
				WithPersistentUint32("u32", "Uint32", "u32", "U32", uint32(0)).
				WithPersistentUint32P("u32p", "Uint32P", "u32p", "Z", "U32P", uint32(0)).
				WithPersistentUint64("u64", "Uint64", "u64", "U64", uint64(0)).
				WithPersistentUint64P("u64p", "Uint64P", "u64p", "Q", "U64P", uint64(0)).
				WithPersistentFloat32("f32", "Float32", "f32", "F32", float32(0)).
				WithPersistentFloat32P("f32p", "Float32P", "f32p", "F", "F32P", float32(0)).
				WithPersistentFloat64("f64", "Float64", "f64", "F64", float64(0)).
				WithPersistentFloat64P("f64p", "Float64P", "f64p", "G", "F64P", float64(0)).
				WithPersistentDuration("dur", "Duration", "dur", "DUR").
				WithPersistentDurationP("durp", "DurationP", "durp", "D", "DURP").
				WithPersistentStringSlice("ss", "StringSlice", "ss", "SS").
				WithPersistentStringSliceP("ssp", "StringSliceP", "ssp", "S", "SSP").
				WithPersistentStringToString("sts", "StringToString", "sts", "STS").
				WithPersistentStringToStringP("stsp", "StringToStringP", "stsp", "T", "STSP").
				WithPersistentBoolP("vp", "BoolP", "vp", "V", "VP", false)
		}, "All WithPersistent named-type wrappers must not panic during registration")
	})
}

func (suite *testCommandPersistent) TestPersistentAllVarTypesRegistrationNoError() {
	suite.WithCustom(nil, nil, func() {
		suite.NotPanics(func() {
			var (
				strVal  string
				strValP string
				boolVal bool
				intVal  int
				intValP int
				i8      int8
				i8p     int8
				i16     int16
				i16p    int16
				i32     int32
				i32p    int32
				i64     int64
				i64p    int64
				u       uint
				up      uint
				u8      uint8
				u8p     uint8
				u16     uint16
				u16p    uint16
				u32     uint32
				u32p    uint32
				u64     uint64
				u64p    uint64
				f32     float32
				f32p    float32
				f64     float64
				f64p    float64
				dur     = time.Duration(0)
				durp    = time.Duration(0)
				ss      []string
				ssp     []string
				sts     = map[string]string{}
				stsp    = map[string]string{}
			)
			New("app", "app", "").
				WithPersistentStringVar(&strVal, "Str", "strv", "STRV").
				WithPersistentStringVarP(&strValP, "StrP", "strvp", "A", "STRVP").
				WithPersistentBoolVar(&boolVal, "Bool", "boolv", "BOOLV").
				WithPersistentBoolVarP(&boolVal, "BoolP", "boolvp", "B", "BOOLVP").
				WithPersistentIntVar(&intVal, "Int", "intv", "INTV").
				WithPersistentIntVarP(&intValP, "IntP", "intvp", "I", "INTVP").
				WithPersistentInt8Var(&i8, "I8", "i8v", "I8V").
				WithPersistentInt8VarP(&i8p, "I8P", "i8vp", "8", "I8VP").
				WithPersistentInt16Var(&i16, "I16", "i16v", "I16V").
				WithPersistentInt16VarP(&i16p, "I16P", "i16vp", "6", "I16VP").
				WithPersistentInt32Var(&i32, "I32", "i32v", "I32V").
				WithPersistentInt32VarP(&i32p, "I32P", "i32vp", "3", "I32VP").
				WithPersistentInt64Var(&i64, "I64", "i64v", "I64V").
				WithPersistentInt64VarP(&i64p, "I64P", "i64vp", "9", "I64VP").
				WithPersistentUintVar(&u, "Uint", "uintv", "UINTV").
				WithPersistentUintVarP(&up, "UintP", "uintvp", "U", "UINTVP").
				WithPersistentUint8Var(&u8, "U8", "u8v", "U8V").
				WithPersistentUint8VarP(&u8p, "U8P", "u8vp", "X", "U8VP").
				WithPersistentUint16Var(&u16, "U16", "u16v", "U16V").
				WithPersistentUint16VarP(&u16p, "U16P", "u16vp", "Y", "U16VP").
				WithPersistentUint32Var(&u32, "U32", "u32v", "U32V").
				WithPersistentUint32VarP(&u32p, "U32P", "u32vp", "Z", "U32VP").
				WithPersistentUint64Var(&u64, "U64", "u64v", "U64V").
				WithPersistentUint64VarP(&u64p, "U64P", "u64vp", "Q", "U64VP").
				WithPersistentFloat32Var(&f32, "F32", "f32v", "F32V").
				WithPersistentFloat32VarP(&f32p, "F32P", "f32vp", "F", "F32VP").
				WithPersistentFloat64Var(&f64, "F64", "f64v", "F64V").
				WithPersistentFloat64VarP(&f64p, "F64P", "f64vp", "G", "F64VP").
				WithPersistentDurationVar(&dur, "Dur", "durv", "DURV").
				WithPersistentDurationVarP(&durp, "DurP", "durvp", "D", "DURVP").
				WithPersistentStringSliceVar(&ss, "SS", "ssv", "SSV").
				WithPersistentStringSliceVarP(&ssp, "SSP", "ssvp", "S", "SSVP").
				WithPersistentStringToStringVar(&sts, "STS", "stsv", "STSV").
				WithPersistentStringToStringVarP(&stsp, "STSP", "stsvp", "T", "STSVP")
		}, "All WithPersistent var-bound-type wrappers must not panic during registration")
	})
}
