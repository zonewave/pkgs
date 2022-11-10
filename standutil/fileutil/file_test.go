package fileutil

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestAbsPath(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		ret, err := AbsPath("file_test.go")
		require.NoError(t, err)
		expectFileSuffix := "pkgs/standutil/fileutil/file_test.go"
		suffixLen := len(expectFileSuffix)
		require.GreaterOrEqual(t, len(ret), suffixLen)
		require.Equal(t, expectFileSuffix, ret[len(ret)-suffixLen:])

	})
	t.Run("use env", func(t *testing.T) {
		envKey := "TestAbsPath"
		err := os.Setenv(envKey, "standutil/fileutil")
		require.NoError(t, err)
		ret, err := AbsPath("$" + envKey + "/file_test")
		require.NoError(t, err)
		require.GreaterOrEqual(t, len(ret), len("standutil/fileutil/file_test"))
	})

}

func TestFileInfo(t *testing.T) {
	name, ext := FileNameAndExt("go/test.go")
	require.Equal(t, "go/test", name)
	require.Equal(t, "go", ext)
}

func (s *Suite) TestSearchInPaths() {
	s.afero.EXPECT().Exists("path1/test").Return(true, nil).Times(1)
	file, err := SearchInPaths(s.afero, []string{"path1", "path2"}, "test")
	s.Require().NoError(err)
	s.Require().Equal("path1/test", file)

	s.afero.EXPECT().Exists(gomock.Any()).Return(false, os.ErrNotExist).Times(2)
	file, err = SearchInPaths(s.afero, []string{"path1", "path2"}, "test")
	s.Require().ErrorIs(err, ErrNotFound)
	s.Require().Empty(file)

}

func TestFileExtNoDot(t *testing.T) {

	tests := []struct {
		name    string
		file    string
		wantExt string
	}{
		// TODO: Add test cases.
		{
			name:    "normal",
			file:    "test.go",
			wantExt: "go",
		},
		{
			name:    "no ext",
			file:    "test",
			wantExt: "",
		},
		{
			name:    "multi ext",
			file:    "test.go.tar.gz",
			wantExt: "gz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotExt := FileExtNoDot(tt.file); gotExt != tt.wantExt {
				t.Errorf("FileExtNoDot() = %v, want %v", gotExt, tt.wantExt)
			}
		})
	}
}
