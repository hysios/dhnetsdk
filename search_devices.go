package dhnetsdk

// #cgo CFLAGS: -Iinclude/
// #cgo amd64 386 CFLAGS: -DX86=1
// #cgo LDFLAGS: -Llib/ -ldhnetsdk -lpthread -Wl,-rpath=./lib/
// #include <stdio.h>
// #include <stdlib.h>
// #include <string.h>
// #include <stdbool.h>
// #include "dhnetsdk.h"
//
// extern void goSearchDevicesEx(void* dwUser, DEVICE_NET_INFO_EX2* devinfoex);
//
// void CALLBACK cSearchDevicesCBEx(LLONG lSearchHandle, DEVICE_NET_INFO_EX2 *pDevNetInfo, void* pUserData) {
//     printf("%p", pUserData);
//     if (0 != pUserData)
//     {
//         goSearchDevicesEx(pUserData, pDevNetInfo);
//     }
// }
import "C"

import (
	"unsafe"

	"github.com/mattn/go-pointer"
)

type SearchDeviceFunc func(search *NetSearchDevice, deviceinfoEx *DeviceNetInfoEx)
type SearchVisitor struct {
	Search   *NetSearchDevice
	Callback SearchDeviceFunc
}

type SendSearchType int

const (
	StMulticastAndBroadcast SendSearchType = C.EM_SEND_SEARCH_TYPE_MULTICAST_AND_BROADCAST // 组播和广播搜索
	StSearchTypeMulticast   SendSearchType = C.EM_SEND_SEARCH_TYPE_MULTICAST               // 组播搜索
	StSearchTypeBroadcast   SendSearchType = C.EM_SEND_SEARCH_TYPE_BROADCAST               // 广播搜索
)

type NetSearchDevice struct {
	incptr  *C.NET_IN_STARTSERACH_DEVICE
	outcptr *C.NET_OUT_STARTSERACH_DEVICE

	visitor *SearchVisitor
	Handle  int
}

func (nsd *NetSearchDevice) init() {
	if nsd.incptr != nil {
		return
	}

	if nsd.outcptr != nil {
		return
	}

	nsd.incptr = &C.NET_IN_STARTSERACH_DEVICE{}
	nsd.incptr.dwSize = C.uint(unsafe.Sizeof(*nsd.incptr))

	nsd.outcptr = &C.NET_OUT_STARTSERACH_DEVICE{}
	nsd.outcptr.dwSize = C.uint(unsafe.Sizeof(*nsd.outcptr))
}

func (nsd *NetSearchDevice) Start(callback SearchDeviceFunc, sendType SendSearchType) {
	nsd.init()

	if nsd.visitor != nil {
		pointer.Unref(unsafe.Pointer(nsd.visitor))
		nsd.visitor = &SearchVisitor{
			Search:   nsd,
			Callback: callback,
		}
	} else {
		nsd.visitor = &SearchVisitor{
			Search:   nsd,
			Callback: callback,
		}
	}
	p := pointer.Save(nsd.visitor)
	nsd.incptr.cbSearchDevices = C.fSearchDevicesCBEx(C.cSearchDevicesCBEx)
	nsd.incptr.pUserData = p
	// nsd.incptr.emSendType = C.EM_SEND_SEARCH_TYPE(sendType)

	nsd.Handle = int(C.CLIENT_StartSearchDevicesEx(nsd.incptr, nsd.outcptr))
}

func (nsd *NetSearchDevice) LocalIP() string {
	nsd.init()

	return GoStr((*C.char)(&nsd.incptr.szLocalIp[0]), C.MAX_LOCAL_IP_LEN)
}

func (nsd *NetSearchDevice) SetLocalIP(ip string) {
	nsd.init()
	var n C.ulong
	if len(ip) < C.MAX_LOCAL_IP_LEN {
		n = C.ulong(len(ip))
	} else {
		n = 88
	}

	C.strncpy((*C.char)(&nsd.incptr.szLocalIp[0]), C.CString(ip), n)
}

func (nsd *NetSearchDevice) Stop() bool {
	if nsd.visitor != nil {
		pointer.Unref(unsafe.Pointer(nsd.visitor))
		nsd.visitor = nil
	}

	b := C.CLIENT_StopSearchDevices(C.LLONG(nsd.Handle))

	return b > 0
}
