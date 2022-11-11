package config

import (
	"errors"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func (suite *Suite) TestJSON() {
	cfg := &AppConfig{}
	gotErr := Initialize(suite.JSONFile, cfg)
	suite.NoError(gotErr)
	got, ok := Get().(*AppConfig)
	suite.True(ok)
	suite.Equal("test-app", got.AppName)
	suite.Equal("test", got.AppEnv)
	suite.Equal(true, got.Debug)
	suite.Equal("root@tcp(localhost:3306)/test", got.Database.DSN)
}

func (suite *Suite) TestENV() {
	_ = os.Setenv("APP_ENV", "development")
	_ = os.Setenv("APP_DEBUG", "false")
	_ = os.Setenv("DB_DSN", "test@tcp(localhost:3306)")

	cfg := &AppConfig{}
	gotErr := Initialize(suite.EnvFile, cfg)
	suite.NoError(gotErr)
	got, ok := Get().(*AppConfig)
	suite.True(ok)
	suite.Equal("test-app", got.AppName)
	suite.Equal("development", got.AppEnv)
	suite.Equal(false, got.Debug)
	suite.Equal("test@tcp(localhost:3306)", got.Database.DSN)
}

func (suite *Suite) TestYAML() {
	cfg := &AppConfig{}
	gotErr := Initialize(suite.YAMLFile, cfg)
	suite.NoError(gotErr)
	got, ok := Get().(*AppConfig)
	suite.True(ok)
	suite.Equal("test-app", got.AppName)
	suite.Equal("test", got.AppEnv)
	suite.Equal(true, got.Debug)
	suite.Equal("root@tcp(localhost:3306)/test", got.Database.DSN)
}

func (suite *Suite) TestTOML() {
	cfg := &AppConfig{}
	gotErr := Initialize(suite.TOMLFile, cfg)
	suite.NoError(gotErr)
	got, ok := Get().(*AppConfig)
	suite.True(ok)
	suite.Equal("test-app", got.AppName)
	suite.Equal("test", got.AppEnv)
	suite.Equal(true, got.Debug)
	suite.Equal("root@tcp(localhost:3306)/test", got.Database.DSN)
}

func (suite *Suite) TestCSV() {
	cfg := &CSVConfig{}
	gotErr := Initialize(suite.JSONFile, cfg)
	suite.Require().NoError(gotErr)
	got, ok := Get().(*CSVConfig)
	suite.True(ok)
	suite.Equal("test-app", got.AppName)
	suite.NotEmpty(got.LessonExtensions)
}

func Test_set(t *testing.T) {
	type args struct {
		configFile   string
		cfgStructPtr interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		// TODO: Add test cases.
		{
			"ok",
			args{
				"1.json",
				&AppConfig{},
			},
			nil,
		},
		{
			"InvalidConfigTypeError",
			args{
				"1.jsont",
				&AppConfig{},
			},
			InvalidConfigTypeError("jsont"),
		},
		{
			"ok",
			args{
				"1.json",
				3,
			},
			InvalidConfigTypeError("should be a pointer to a struct"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := set(tt.args.configFile, tt.args.cfgStructPtr); !errors.Is(err, tt.wantErr) {
				t.Errorf("set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoader(t *testing.T) {
	l := Loader()
	require.NotEmpty(t, l)
	l.SetDebug(true)
	require.Equal(t, true, l.debug)
}
