package main

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var iconizeTests = []struct {
	name          string
	args          []string
	inputFixture  string
	goldenFixture string
}{
	{"Nerd", []string{"-f=nerd"}, "nerd/iconize.input", "nerd/iconize.golden"},
	{"NerdColor", []string{"-f=nerd", "-c"}, "nerd/iconize.input", "nerd/iconize_color.golden"},
	{"NerdDir", []string{"-f=nerd", "-d"}, "nerd/iconize.input", "nerd/iconize_dir.golden"},
	{"NerdDirColor", []string{"-f=nerd", "-c", "-d"}, "nerd/iconize.input", "nerd/iconize_dir_color.golden"},
}

func TestIconizeStdin(t *testing.T) {
	binary, err := getBinary()
	require.NoError(t, err, "Cannot get binary path")

	projectPath, err := filepath.Abs("..")
	require.NoError(t, err, "Cannot get project path")

	for _, tt := range iconizeTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			input, err := loadFixture(tt.inputFixture)
			require.NoErrorf(t, err, "Cannot load fixture %s", tt.inputFixture)

			golden, err := loadFixture(tt.goldenFixture)
			require.NoErrorf(t, err, "Cannot load fixture %s", tt.goldenFixture)

			cmdArgs := append([]string{"iconize"}, tt.args...)
			cmd := exec.Command(binary, cmdArgs...)
			cmd.Stdin = bytes.NewReader(input)
			cmd.Dir = projectPath

			result, err := cmd.Output()
			require.NoErrorf(t, err, "Cannot execute binary %s", binary)

			assert.Equal(t, string(golden), string(result))
		})
	}
}

func TestIconizeArgs(t *testing.T) {
	binary, err := getBinary()
	require.NoError(t, err, "Cannot get binary path")

	projectPath, err := filepath.Abs("..")
	require.NoError(t, err, "Cannot get project path")

	for _, tt := range iconizeTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			input, err := loadFixture(tt.inputFixture)
			require.NoErrorf(t, err, "Cannot load fixture %s", tt.inputFixture)

			golden, err := loadFixture(tt.goldenFixture)
			require.NoErrorf(t, err, "Cannot load fixture %s", tt.goldenFixture)

			inputLines := strings.Split(string(input), "\n")
			inputLines = inputLines[0 : len(inputLines)-1]

			cmdArgs := append([]string{"iconize"}, append(tt.args, inputLines...)...)
			cmd := exec.Command(binary, cmdArgs...)
			cmd.Dir = projectPath

			result, err := cmd.Output()
			require.NoErrorf(t, err, "Cannot execute binary %s", binary)

			assert.Equal(t, string(golden), string(result))
		})
	}
}
