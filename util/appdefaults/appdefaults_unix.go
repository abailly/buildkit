// +build !windows

package appdefaults

const (
	Address = "unix:///run/buildkit/buildd.sock"
	Root    = ".buildstate"
)
