package git

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

// Git Create Git manipulation command
type Git struct {
	Dir string
}

// NewGit create default Git object
func NewGit(dir string) *Git {
	return &Git{Dir: dir}
}

// Update is update command of repo
func (g *Git) Update() (cmd *exec.Cmd) {
	args := []string{"pull"}
	cmd = gitCmd(args)
	cmd.Dir = g.Dir

	return
}

// UpdateCurrent is update command of repo of current branch
func (g *Git) UpdateCurrent() (cmd *exec.Cmd) {
	args := []string{"pull", "origin", currentBranch(g.Dir)}
	cmd = gitCmd(args)
	cmd.Dir = g.Dir

	return
}

func (g *Git) Clone(r string) (cmd *exec.Cmd) {
	args := []string{"clone", r, g.Dir}
	cmd = gitCmd(args)

	return
}

func (g *Git) LogOneline() (cmd *exec.Cmd) {
	args := []string{"--no-pager", "log", "-1", "--oneline"}
	cmd = gitCmd(args)
	cmd.Dir = g.Dir

	return
}

func currentBranch(path string) string {
	args := []string{"rev-parse", "--abbrev-ref", "HEAD"}
	cmd := exec.Command("git", args...)
	output := new(bytes.Buffer)
	cmd.Stdout = output
	cmd.Stderr = output
	cmd.Dir = path

	err := cmd.Run()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(output.String())
}

func gitCmd(args []string) (cmd *exec.Cmd) {
	cmd = exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return
}
