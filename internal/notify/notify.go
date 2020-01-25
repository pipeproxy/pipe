package notify

import (
	"os"
	"syscall"

	"github.com/wzshiming/notify"
)

type Signal int

const (
	_ Signal = iota
	Reload
	Reopen
	Stop
)

var signalMapping = map[Signal][]os.Signal{
	Reload: {syscall.SIGHUP},
	Reopen: {syscall.SIGUSR1},
	Stop:   {syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM},
}

// On signal callback.
func On(sig Signal, f func()) func() {
	sigs, ok := signalMapping[sig]
	if !ok || len(sigs) == 0 {
		return func() {}
	}
	return notify.OnSlice(sigs, f)
}

// Once signal callback.
func Once(sig Signal, f func()) {
	sigs, ok := signalMapping[sig]
	if !ok || len(sigs) == 0 {
		return
	}
	notify.OnceSlice(sigs, f)
	return
}

// Kill send signal to pid
func Kill(pid int, sig Signal) error {
	sigs, ok := signalMapping[sig]
	if !ok || len(sigs) == 0 {
		return nil
	}
	return syscall.Kill(pid, sigs[0].(syscall.Signal))
}
