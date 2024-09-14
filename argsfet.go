package main

import (
	"syscall"
	"unsafe"
)

func ObtainProcessCommandArgs(processName string) string {
	// 将进程名转换为UTF-16指针
	processNameUTF16, _ := syscall.UTF16PtrFromString(processName)

	// 加载DLL以及函数
	dll := syscall.MustLoadDLL("argsfet.dll")
	GetProcessIdByName := dll.MustFindProc("GetProcessIdByName")
	GetCommandLineByProcessId := dll.MustFindProc("GetCommandLineByProcessId")

	// 获取PID
	ret, _, _ := GetProcessIdByName.Call(uintptr(unsafe.Pointer(processNameUTF16)))

	// 获取命令行参数
	ret, _, _ = GetCommandLineByProcessId.Call(ret)
	commandLine := syscall.UTF16ToString((*[1 << 16]uint16)(unsafe.Pointer(ret))[:])

	// 释放DLL
	_ = dll.Release()
	return commandLine
}
