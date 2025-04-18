//go:build darwin

package ioctl

// These values are placeholders to support compilation on MacOS.
const (
	TRANSMISSION_IOCTL_DISCONNECT = 43784
	TRANSMISSION_IOCTL_CLEAR_SOCK = 43780
	TRANSMISSION_IOCTL_CLEAR_QUE  = 43781
)
