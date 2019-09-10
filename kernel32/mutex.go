package kernel32

import (
	"syscall"
	"unsafe"
)

var (
	kernel32        = syscall.NewLazyDLL("kernel32.dll")
	procCreateMutex = kernel32.NewProc("CreateMutexW")
	procOpenMutex   = kernel32.NewProc("OpenMutexW")
)

func CreateMutex(name string) (uintptr, error) {
	paramMutexName := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name)))

	ret, _, err := procCreateMutex.Call(0, 0, paramMutexName)

	switch int(err.(syscall.Errno)) {
	case 0:
		return ret, nil
	default:
		return ret, err
	}
}

// 2031617 SYNCRONIZE
func OpenMutex(dwDesiredAccess uintptr, bInheritHandle uintptr, lpName string) (uintptr, error) {
	paramMutexName := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpName)))
	ret, _, err := procOpenMutex.Call(
		dwDesiredAccess, //0x00100000, // SYNCRONIZE
		bInheritHandle,  // Not inheritable
		paramMutexName)

	switch int(err.(syscall.Errno)) {
	case 0:
		return ret, nil
	default:
		return ret, err
	}
}

// 防止进程多开，返回 true 表示进程已经开启
func ProcessMutex(name string) bool {
	r, err := openMutex(name)
	if err == nil || r != 0 {
		return true
	}
	r, err = createMutex(name)
	if err != nil || r == 0 {
		return true
	}
	return false
}
