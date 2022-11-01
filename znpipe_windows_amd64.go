// +build windows
// go build mksyscall_windows.go && ./mksyscall_windows npipe_windows.go
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

package npipe

import "unsafe"
import "syscall"
import "golang.org/x/sys/windows"

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")

	procCreateNamedPipeW    = modkernel32.NewProc("CreateNamedPipeW")
	procConnectNamedPipe    = modkernel32.NewProc("ConnectNamedPipe")
	procDisconnectNamedPipe = modkernel32.NewProc("DisconnectNamedPipe")
	procWaitNamedPipeW      = modkernel32.NewProc("WaitNamedPipeW")
	procCreateEventW        = modkernel32.NewProc("CreateEventW")
	procGetOverlappedResult = modkernel32.NewProc("GetOverlappedResult")
	procCancelIoEx          = modkernel32.NewProc("CancelIoEx")
)

func createNamedPipe(name *uint16, openMode uint32, pipeMode uint32, maxInstances uint32, outBufSize uint32, inBufSize uint32, defaultTimeout uint32, sid *windows.SECURITY_DESCRIPTOR) (handle syscall.Handle, err error) {
	var sa *windows.SecurityAttributes
	if sid != nil {
		sa = &windows.SecurityAttributes{
			Length: uint32(unsafe.Sizeof(*sid)),
			SecurityDescriptor: sid,
		}
	}
	
	r0, _, e1 := syscall.Syscall9(procCreateNamedPipeW.Addr(), 8, uintptr(unsafe.Pointer(name)), uintptr(openMode), uintptr(pipeMode), uintptr(maxInstances), uintptr(outBufSize), uintptr(inBufSize), uintptr(defaultTimeout), uintptr(unsafe.Pointer(sa)), 0)
	handle = syscall.Handle(r0)
	if handle == syscall.InvalidHandle {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func cancelIoEx(handle syscall.Handle, overlapped *syscall.Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall(procCancelIoEx.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(overlapped)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func connectNamedPipe(handle syscall.Handle, overlapped *syscall.Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall(procConnectNamedPipe.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(overlapped)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func disconnectNamedPipe(handle syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procDisconnectNamedPipe.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func waitNamedPipe(name *uint16, timeout uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procWaitNamedPipeW.Addr(), 2, uintptr(unsafe.Pointer(name)), uintptr(timeout), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func createEvent(sa *syscall.SecurityAttributes, manualReset bool, initialState bool, name *uint16) (handle syscall.Handle, err error) {
	var _p0 uint32
	if manualReset {
		_p0 = 1
	} else {
		_p0 = 0
	}
	var _p1 uint32
	if initialState {
		_p1 = 1
	} else {
		_p1 = 0
	}
	r0, _, e1 := syscall.Syscall6(procCreateEventW.Addr(), 4, uintptr(unsafe.Pointer(sa)), uintptr(_p0), uintptr(_p1), uintptr(unsafe.Pointer(name)), 0, 0)
	handle = syscall.Handle(r0)
	if handle == syscall.InvalidHandle {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func getOverlappedResult(handle syscall.Handle, overlapped *syscall.Overlapped, transferred *uint32, wait bool) (err error) {
	var _p0 uint32
	if wait {
		_p0 = 1
	} else {
		_p0 = 0
	}
	r1, _, e1 := syscall.Syscall6(procGetOverlappedResult.Addr(), 4, uintptr(handle), uintptr(unsafe.Pointer(overlapped)), uintptr(unsafe.Pointer(transferred)), uintptr(_p0), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
