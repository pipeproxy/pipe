package listener

import (
	"errors"
	"fmt"
	"net"
	"os"
	"reflect"
	"runtime"
	"strings"
)

func buildKey(network, addr string) string {
	return fmt.Sprintf("%s://%s", network, addr)
}

func sameAddress(a1, a2 string) string {
	host1, port1, err := net.SplitHostPort(a1)
	if err != nil {
		return a1
	}

	switch port1 {
	case "0", "":
		_, port2, err := net.SplitHostPort(a2)
		if err != nil {
			return a1
		}
		port1 = port2
	}
	return net.JoinHostPort(host1, port1)
}

var ErrNetClosing = errors.New("use of closed network connection")

// IsClosedConnError reports whether err is an error from use of a closed network connection.
func IsClosedConnError(err error) bool {
	if err == nil {
		return false
	}

	if err == ErrNetClosing || strings.Contains(err.Error(), ErrNetClosing.Error()) {
		return true
	}

	if runtime.GOOS == "windows" {
		if oe, ok := err.(*net.OpError); ok && oe.Op == "read" {
			if se, ok := oe.Err.(*os.SyscallError); ok && se.Syscall == "wsarecv" {
				const WSAECONNABORTED = 10053
				const WSAECONNRESET = 10054
				if n := errno(se.Err); n == WSAECONNRESET || n == WSAECONNABORTED {
					return true
				}
			}
		}
	}
	return false
}

func errno(v error) uintptr {
	if rv := reflect.ValueOf(v); rv.Kind() == reflect.Uintptr {
		return uintptr(rv.Uint())
	}
	return 0
}

var ErrAcceptTimeout = errors.New("i/o timeout")

// IsAcceptTimeoutError reports whether err is an error from use of a accept timeout.
func IsAcceptTimeoutError(err error) bool {
	if err == nil {
		return false
	}

	if err == ErrAcceptTimeout || strings.Contains(err.Error(), ErrAcceptTimeout.Error()) {
		return true
	}

	if oe, ok := err.(*net.OpError); ok && oe.Op == "accept" {
		return IsAcceptTimeoutError(oe.Err)
	}

	return false
}
