//go:build !windows
// +build !windows

package system

import (
	"os"
	"strconv"
	"syscall"
)

// StatT type contains status of a file. It contains metadata
// like permission, owner, group, size, etc about a file.
type StatT struct {
	mode uint32
	uid  uint32
	gid  uint32
	rdev uint64
	size int64
	mtim syscall.Timespec
	dev  uint64
	platformStatT
}

// Mode returns file's permission mode.
func (s StatT) Mode() uint32 {
	return s.mode
}

// UID returns file's user id of owner.
func (s StatT) UID() uint32 {
	return s.uid
}

// GID returns file's group id of owner.
func (s StatT) GID() uint32 {
	return s.gid
}

// Rdev returns file's device ID (if it's special file).
func (s StatT) Rdev() uint64 {
	return s.rdev
}

// Size returns file's size.
func (s StatT) Size() int64 {
	return s.size
}

// Mtim returns file's last modification time.
func (s StatT) Mtim() syscall.Timespec {
	return s.mtim
}

// Dev returns a unique identifier for owning filesystem
func (s StatT) Dev() uint64 {
	return s.dev
}

// Stat takes a path to a file and returns
// a system.StatT type pertaining to that file.
//
// Throws an error if the file does not exist
func Stat(path string) (*StatT, error) {
	s := &syscall.Stat_t{}
	if err := syscall.Stat(path, s); err != nil {
		return nil, &os.PathError{Op: "Stat", Path: path, Err: err}
	}
	return fromStatT(s)
}

// Fstat takes an open file descriptor and returns
// a system.StatT type pertaining to that file.
//
// Throws an error if the file descriptor is invalid
func Fstat(fd int) (*StatT, error) {
	s := &syscall.Stat_t{}
	if err := syscall.Fstat(fd, s); err != nil {
		return nil, &os.PathError{Op: "Fstat", Path: strconv.Itoa(fd), Err: err}
	}
	return fromStatT(s)
}
