package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test_run(t *testing.T) {
	err := os.Chdir("./")
	require.NoError(t, err)
	run([]string{"run", "./", "1"})

}

func Test_RenameFiles1(t *testing.T) {
	err := os.Chdir("./")
	require.NoError(t, err)
	_, err = os.Create("tmp1.py")
	require.NoError(t, err)

	err = renameFiles("./", RenameMethod1)
	require.NoError(t, err)
	_, err = os.Stat("_1_tmp.py")
	require.NoError(t, err)
	_, err = os.Stat("tmp1.py")
	require.Error(t, err)
	err = os.Remove("_1_tmp.py")
	require.NoError(t, err)

}

func Test_RenameFiles2(t *testing.T) {
	err := os.Chdir("./")
	require.NoError(t, err)
	_, err = os.Create("剑指 tt.py")
	require.NoError(t, err)

	err = renameFiles("./", RenameMethod2)
	require.NoError(t, err)
	_, err = os.Stat("tt.py")
	require.NoError(t, err)
	_, err = os.Stat("剑指 tt.py")
	require.Error(t, err)
	err = os.Remove("tt.py")
	require.NoError(t, err)

}
