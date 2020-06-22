package dhnetsdk

// #cgo CFLAGS: -Iinclude/
// #cgo amd64 386 CFLAGS: -DX86=1
// #cgo LDFLAGS: -Llib/ -ldhnetsdk -lpthread -Wl,-rpath ./lib/
// #include <stdio.h>
// #include <stdlib.h>
// #include <string.h>
// #include <stdbool.h>
// #include "dhnetsdk.h"
import "C"

import (
	"unsafe"

	"github.com/mattn/go-pointer"
)

//export goDisconnect
func goDisconnect(user_data unsafe.Pointer, ip *C.char, port C.int) {
	if v, ok := pointer.Restore(user_data).(DisconnectVisitor); ok {
		defer pointer.Unref(user_data)
		if v.Callback != nil {
			v.Callback(v.Client, C.GoString(ip), int(port))
		}

	}
}

//export goReConnect
func goReConnect(user_data unsafe.Pointer, ip *C.char, port C.int) {
	if v, ok := pointer.Restore(user_data).(ReconnectVisitor); ok {
		defer pointer.Unref(user_data)
		if v.Callback != nil {
			v.Callback(v.Client, C.GoString(ip), int(port))
		}
	}

}

//export goDvrMessage
func goDvrMessage(user_data unsafe.Pointer, cmd C.int, buf *C.char, l C.int, ip *C.char, port C.int) {
	if v, ok := pointer.Restore(user_data).(DrvMessageVisitor); ok {
		defer pointer.Unref(user_data)
		if v.Callback != nil {
			v.Callback(v.Client, DhAlarmType(cmd), C.GoBytes(unsafe.Pointer(buf), l), C.GoString(ip), int(port))
		}
	}
}

//export goSearchDevicesEx
func goSearchDevicesEx(user_data unsafe.Pointer, devinfoex *C.DEVICE_NET_INFO_EX2) {
	if v, ok := pointer.Restore(user_data).(*SearchVisitor); ok {
		if v.Callback != nil {
			devinfo := DeviceNetInfo{cptr: &devinfoex.stuDevInfo}
			v.Callback(v.Search, &DeviceNetInfoEx{cptr: devinfoex, DeviceNetInfo: devinfo})
		}
	}
}

//export goAnalyzerDataCallBack
func goAnalyzerDataCallBack(user_data unsafe.Pointer, dwAlarmType C.DWORD, pAlarmInfo uintptr, pBuffer *C.char, dwBufSize C.DWORD, nSequence C.int) C.int {
	if v, ok := pointer.Restore(user_data).(PictureExVistor); ok {
		defer pointer.Unref(user_data)

		if v.Callback != nil {
			alarmType := EventIvs(dwAlarmType)
			switch alarmType {
			case EventIvsTrafficParking:
				var alarmInfo TrafficParkingInfo
				alarmInfo.cptr = (*C.DEV_EVENT_TRAFFIC_PARKING_INFO)(unsafe.Pointer(pAlarmInfo))
				buf := C.GoBytes(unsafe.Pointer(pBuffer), C.int(dwBufSize))
				return C.int(v.Callback(v.Client, alarmType, &alarmInfo, buf, int(nSequence)))
			default:
				buf := C.GoBytes(unsafe.Pointer(pBuffer), C.int(dwBufSize))
				return C.int(v.Callback(v.Client, alarmType, nil, buf, int(nSequence)))
			}
		}
	}

	return 0
}
