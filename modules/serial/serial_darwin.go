// +build darwin dragonfly freebsd netbsd openbsd
package serial

import (
	"syscall"
	"unsafe"
	"golang.org/x/sys/unix"
)

const ioctlReadTermios = unix.TIOCGETA
const ioctlWriteTermios = unix.TIOCSETA

const (
	FREAD  = 0x0001
	FWRITE = 0x0002


	CRTSCTS = 0x80000000
)


func Tcflush(fd, which uintptr) error {
	var com int
	switch which {
	case syscall.TCIFLUSH:
		com = FREAD
	case syscall.TCOFLUSH:
		com = FWRITE
	case syscall.TCIOFLUSH:
		com = FREAD | FWRITE
	default:
		return syscall.EINVAL
	}
	return ioctl(fd, syscall.TIOCFLUSH, uintptr(unsafe.Pointer(&com)))
}


func setSpeed(s *unix.Termios, baud int) error {
	var rate uint64

	switch baud {
	case 50:
		rate = syscall.B50
	case 75:
		rate = syscall.B75
	case 110:
		rate = syscall.B110
	case 134:
		rate = syscall.B134
	case 150:
		rate = syscall.B150
	case 200:
		rate = syscall.B200
	case 300:
		rate = syscall.B300
	case 600:
		rate = syscall.B600
	case 1200:
		rate = syscall.B1200
	case 1800:
		rate = syscall.B1800
	case 2400:
		rate = syscall.B2400
	case 4800:
		rate = syscall.B4800
	case 9600:
		rate = syscall.B9600
	case 19200:
		rate = syscall.B19200
	case 38400:
		rate = syscall.B38400
	case 57600:
		rate = syscall.B57600
	case 115200:
		rate = syscall.B115200
	case 230400:
		rate = syscall.B230400
	default:
		return syscall.EINVAL
	}
	s.Cflag |= rate
	s.Ispeed = rate
	s.Ospeed = rate
	return nil
}