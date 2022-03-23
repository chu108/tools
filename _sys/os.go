package _sys

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

type OS struct {
}

func NewOS() *OS {
	return &OS{}
}

// IsWindows
func (*OS) IsWindows() bool {
	return "windows" == runtime.GOOS
}

// IsLinux
func (*OS) IsLinux() bool {
	return "linux" == runtime.GOOS
}

// IsDarwin
func (*OS) IsDarwin() bool {
	return "darwin" == runtime.GOOS
}

// Pwd
func (*OS) Pwd() string {
	file, _ := exec.LookPath(os.Args[0])
	pwd, _ := filepath.Abs(file)

	return filepath.Dir(pwd)
}

func (obj *OS) Home() (string, error) {
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}
	// cross compile support
	if obj.IsWindows() {
		return obj.homeWindows()
	}
	// Unix-like system, so just assume Unix
	return obj.homeUnix()
}

func (*OS) homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}
	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}
	return result, nil
}

func (*OS) homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}
