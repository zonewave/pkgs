package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test_run(t *testing.T) {
	err := os.Chdir("./")
	require.NoError(t, err)
	run([]string{"run", "./"})

}

func Test_r(t *testing.T) {
	err := os.Chdir("./")
	require.NoError(t, err)
	_, err = os.Create("tmp1.py")
	require.NoError(t, err)

	err = renameFiles("./")
	require.NoError(t, err)
	_, err = os.Stat("_1_tmp.py")
	require.NoError(t, err)
	_, err = os.Stat("tmp1.py")
	require.Error(t, err)
	err = os.Remove("_1_tmp.py")
	require.NoError(t, err)

}
