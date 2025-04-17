//go:build darwin

package ioctl

// These values are placeholders to support compilation on MacOS.
const (
	NEGOTIATION_IOCTL_SET_SOCK        = 43776
	NEGOTIATION_IOCTL_SET_BLOCKSIZE   = 43777
	NEGOTIATION_IOCTL_SET_SIZE_BLOCKS = 43783
	NEGOTIATION_IOCTL_DO_IT           = 43779
	NEGOTIATION_IOCTL_SET_TIMEOUT     = 43785
)
