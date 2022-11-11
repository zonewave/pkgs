package fileutil

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestAbsPath(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		ret := AbsPath("file_test.go")
		expectFileSuffix := "pkgs/util/fileutil/file_test.go"
		suffixLen := len(expectFileSuffix)
		require.GreaterOrEqual(t, len(ret), suffixLen)
		require.Equal(t, expectFileSuffix, ret[len(ret)-suffixLen:])

	})
	t.Run("use env", func(t *testing.T) {
		envKey := "TestAbsPath"
		err := os.Setenv(envKey, "util/fileutil")
		require.NoError(t, err)
		ret := AbsPath("$" + envKey + "/file_test")
		require.GreaterOrEqual(t, len(ret), len("util/fileutil/file_test"))
	})

}

func TestFileInfo(t *testing.T) {
	name, ext := FileInfo("test.go")
	require.Equal(t, "test", name)
	require.Equal(t, "go", ext)
}