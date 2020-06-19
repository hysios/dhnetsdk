package dhnetsdk

// #cgo CFLAGS: -Iinclude/
// #cgo amd64 386 CFLAGS: -DX86=1
// #cgo LDFLAGS: -Llib/ -ldhnetsdk -lpthread -Wl,-rpath=./lib/
// #include <stdio.h>
// #include <stdlib.h>
// #include <string.h>
// #include <stdbool.h>
// #include "dhnetsdk.h"
import "C"

import (
	"fmt"
	"reflect"
	"strings"
	"text/tabwriter"
	"unsafe"
)

type DeviceNetInfo struct {
	cptr *C.DEVICE_NET_INFO_EX
}

type DeviceNetInfoEx struct {
	DeviceNetInfo
	cptr *C.DEVICE_NET_INFO_EX2
}

func (devinfo *DeviceNetInfo) init() {
	if devinfo.cptr != nil {
		return
	}

	devinfo.cptr = &C.DEVICE_NET_INFO_EX{}
}

func (devinfoex *DeviceNetInfoEx) init() {
	if devinfoex.cptr != nil {
		return
	}

	devinfoex.cptr = &C.DEVICE_NET_INFO_EX2{}
}

func (devinfoex *DeviceNetInfoEx) LocalIP() string {
	devinfoex.init()
	return C.GoStringN((*C.char)(&devinfoex.cptr.szLocalIP[0]), 64)
}

func (devinfoex *DeviceNetInfoEx) SetLocalIP(ip string) {
	devinfoex.init()
	var n C.ulong
	if len(ip) < 64 {
		n = C.ulong(len(ip))
	} else {
		n = 64
	}
	C.strncpy((*C.char)(unsafe.Pointer(&devinfoex.cptr.szLocalIP)), C.CString(ip), n)
}

func (devinfo *DeviceNetInfo) IPVersion() int {
	devinfo.init()

	return int(devinfo.cptr.iIPVersion)
}

func (devinfo *DeviceNetInfo) SetIPVersion(ver int) {
	devinfo.init()

	devinfo.cptr.iIPVersion = C.int(ver)
}

func (devinfo *DeviceNetInfo) IP() string {
	devinfo.init()

	return C.GoStringN((*C.char)(&devinfo.cptr.szIP[0]), 64)
}

func (devinfo *DeviceNetInfo) SetIP(ip string) {
	devinfo.init()
	var n C.ulong
	if len(ip) < 64 {
		n = C.ulong(len(ip))
	} else {
		n = 64
	}

	C.strncpy((*C.char)(unsafe.Pointer(&devinfo.cptr.szIP)), C.CString(ip), n)
}

func (devinfo *DeviceNetInfo) Port() int {
	devinfo.init()

	return int(devinfo.cptr.nPort)
}

func (devinfo *DeviceNetInfo) SetPort(port int) {
	devinfo.init()

	devinfo.cptr.nPort = C.int(port)
}

func (devinfo *DeviceNetInfo) Submask() string {
	devinfo.init()

	return C.GoStringN((*C.char)(&devinfo.cptr.szSubmask[0]), 64)
}

func (devinfo *DeviceNetInfo) SetSubmask(mask string) {
	devinfo.init()
	var n C.ulong
	if C.ulong(len(mask)) < 64 {
		n = C.ulong(len(mask))
	} else {
		n = 64
	}

	C.strncpy((*C.char)(unsafe.Pointer(&devinfo.cptr.szSubmask)), C.CString(mask), n)
}

func (devinfo *DeviceNetInfo) Gateway() string {
	devinfo.init()

	return C.GoStringN((*C.char)(&devinfo.cptr.szGateway[0]), 64)
}

func (devinfo *DeviceNetInfo) SetGateway(gateway string) {
	devinfo.init()
	var n C.ulong
	if len(gateway) < 64 {
		n = C.ulong(len(gateway))
	} else {
		n = 64
	}

	C.strncpy((*C.char)(unsafe.Pointer(&devinfo.cptr.szGateway)), C.CString(gateway), n)
}

func (devinfo *DeviceNetInfo) Mac() string {
	devinfo.init()

	return C.GoString((*C.char)(&devinfo.cptr.szMac[0]))
}

func (devinfo *DeviceNetInfo) SetMac(mac string) {
	devinfo.init()
	var n C.ulong
	if len(mac) < C.DH_MACADDR_LEN {
		n = C.ulong(len(mac))
	} else {
		n = C.DH_MACADDR_LEN
	}

	C.strncpy((*C.char)(unsafe.Pointer(&devinfo.cptr.szMac)), C.CString(mac), n)
}

func (devinfo *DeviceNetInfo) DeviceType() string {
	devinfo.init()

	return C.GoStringN((*C.char)(&devinfo.cptr.szDeviceType[0]), C.DH_DEV_TYPE_LEN)
}

func (devinfo *DeviceNetInfo) SetDeviceType(devType string) {
	devinfo.init()
	var n C.ulong
	if len(devType) < C.DH_DEV_TYPE_LEN {
		n = C.ulong(len(devType))
	} else {
		n = C.DH_DEV_TYPE_LEN
	}

	C.strncpy((*C.char)(unsafe.Pointer(&devinfo.cptr.szDeviceType)), C.CString(devType), n)
}

func (devinfo *DeviceNetInfo) ManuFactory() int {
	devinfo.init()

	return int(devinfo.cptr.byManuFactory)
}

func (devinfo *DeviceNetInfo) SetManuFactory(factory int) {
	devinfo.init()

	devinfo.cptr.byManuFactory = C.BYTE(factory)
}

func (devinfo *DeviceNetInfo) Definition() int {
	devinfo.init()

	return int(devinfo.cptr.byDefinition)
}

func (devinfo *DeviceNetInfo) SetDefinition(define int) {
	devinfo.init()

	devinfo.cptr.byDefinition = C.BYTE(define)
}

func (devinfo *DeviceNetInfo) EnableDhcp() bool {
	devinfo.init()

	return bool(devinfo.cptr.bDhcpEn)
}

func (devinfo *DeviceNetInfo) SetEnableDhcp(en bool) {
	devinfo.init()

	devinfo.cptr.bDhcpEn = C.bool(en)
}

func (devinfo *DeviceNetInfo) VerifyData() string {
	devinfo.init()

	return C.GoStringN((*C.char)(unsafe.Pointer(&devinfo.cptr.verifyData[0])), 88)
}

func (devinfo *DeviceNetInfo) SetVerifyData(verify string) {
	devinfo.init()
	var n C.ulong
	if len(verify) < 88 {
		n = C.ulong(len(verify))
	} else {
		n = 88
	}

	C.strncpy((*C.char)(&devinfo.cptr.verifyData[0]), C.CString(verify), n)
}

func (devinfo DeviceNetInfo) String() string {
	var (
		sb strings.Builder
		v  = reflect.ValueOf(&devinfo)
	)
	// v = reflect.Indirect(v)\
	w := tabwriter.NewWriter(&sb, 0, 0, 4, ' ', tabwriter.AlignRight|tabwriter.Debug)
	t := v.Type()
	for i := 0; i < t.NumMethod(); i++ {
		meth := t.Method(i)
		if strings.HasPrefix(meth.Name, "Set") {
			continue
		}

		if meth.Name == "String" {
			continue
		}

		realmeth := v.Method(i)
		// vals := realmeth.Call([]reflect.Value{reflect.ValueOf(devinfo)})

		vals := realmeth.Call([]reflect.Value{})
		if len(vals) == 1 {
			fmt.Fprintf(w, "%s\t\t%v\n", meth.Name, vals[0].Interface())
		}
	}
	w.Flush()
	return sb.String()
}

func (devinfo *DeviceNetInfo) SerialNo() string {
	devinfo.init()

	return C.GoStringN((*C.char)(unsafe.Pointer(&devinfo.cptr.szSerialNo[0])), C.DH_DEV_SERIALNO_LEN)
}

func (devinfo *DeviceNetInfo) SetSerialNo(serialno string) {
	devinfo.init()
	var n C.ulong
	if len(serialno) < C.DH_DEV_SERIALNO_LEN {
		n = C.ulong(len(serialno))
	} else {
		n = 88
	}

	C.strncpy((*C.char)(&devinfo.cptr.szSerialNo[0]), C.CString(serialno), n)
}

func (devinfo *DeviceNetInfo) DevSoftVersion() string {
	devinfo.init()

	return C.GoStringN((*C.char)(unsafe.Pointer(&devinfo.cptr.szDevSoftVersion[0])), C.DH_MAX_URL_LEN)
}

func (devinfo *DeviceNetInfo) SetDevSoftVersion(softver string) {
	devinfo.init()
	var n C.ulong
	if len(softver) < C.DH_MAX_URL_LEN {
		n = C.ulong(len(softver))
	} else {
		n = 88
	}

	C.strncpy((*C.char)(&devinfo.cptr.szDevSoftVersion[0]), C.CString(softver), n)
}

func (devinfo *DeviceNetInfo) DetailType() string {
	devinfo.init()

	return C.GoStringN((*C.char)(unsafe.Pointer(&devinfo.cptr.szDetailType[0])), C.DH_DEV_TYPE_LEN)
}

func (devinfo *DeviceNetInfo) SetDetailType(detail string) {
	devinfo.init()
	var n C.ulong
	if len(detail) < C.DH_DEV_TYPE_LEN {
		n = C.ulong(len(detail))
	} else {
		n = 88
	}

	C.strncpy((*C.char)(&devinfo.cptr.szDetailType[0]), C.CString(detail), n)
}

func (devinfo *DeviceNetInfo) Vendor() string {
	devinfo.init()

	return C.GoStringN((*C.char)(unsafe.Pointer(&devinfo.cptr.szVendor[0])), C.DH_MAX_STRING_LEN)
}

func (devinfo *DeviceNetInfo) SetVendor(vendor string) {
	devinfo.init()
	var n C.ulong
	if len(vendor) < C.DH_MAX_STRING_LEN {
		n = C.ulong(len(vendor))
	} else {
		n = 88
	}

	C.strncpy((*C.char)(&devinfo.cptr.szVendor[0]), C.CString(vendor), n)
}

func (devinfo *DeviceNetInfo) DevName() string {
	devinfo.init()

	return C.GoStringN((*C.char)(unsafe.Pointer(&devinfo.cptr.szDevName[0])), C.DH_MACHINE_NAME_NUM)
}

func (devinfo *DeviceNetInfo) SetDevName(devname string) {
	devinfo.init()
	var n C.ulong
	if len(devname) < C.DH_MACHINE_NAME_NUM {
		n = C.ulong(len(devname))
	} else {
		n = 88
	}

	C.strncpy((*C.char)(&devinfo.cptr.szDevName[0]), C.CString(devname), n)
}

func (devinfo *DeviceNetInfo) UserName() string {
	devinfo.init()

	return C.GoStringN((*C.char)(unsafe.Pointer(&devinfo.cptr.szUserName[0])), C.DH_USER_NAME_LENGTH_EX)
}

func (devinfo *DeviceNetInfo) SetUserName(username string) {
	devinfo.init()
	var n C.ulong
	if len(username) < C.DH_USER_NAME_LENGTH_EX {
		n = C.ulong(len(username))
	} else {
		n = 88
	}

	C.strncpy((*C.char)(&devinfo.cptr.szUserName[0]), C.CString(username), n)
}

func (devinfo *DeviceNetInfo) Password() string {
	devinfo.init()

	return C.GoStringN((*C.char)(unsafe.Pointer(&devinfo.cptr.szPassWord[0])), C.DH_USER_NAME_LENGTH_EX)
}

func (devinfo *DeviceNetInfo) SetPassword(password string) {
	devinfo.init()
	var n C.ulong
	if len(password) < C.DH_USER_NAME_LENGTH_EX {
		n = C.ulong(len(password))
	} else {
		n = 88
	}

	C.strncpy((*C.char)(&devinfo.cptr.szPassWord[0]), C.CString(password), n)
}

func (devinfo *DeviceNetInfo) HttpPort() int {
	devinfo.init()

	return int(devinfo.cptr.nHttpPort)
}

func (devinfo *DeviceNetInfo) SetHttpPort(port int) {
	devinfo.init()

	devinfo.cptr.nHttpPort = C.ushort(port)
}

func (devinfo *DeviceNetInfo) VideoInput() int {
	devinfo.init()

	return int(devinfo.cptr.wVideoInputCh)
}

func (devinfo *DeviceNetInfo) SetVideoInput(input int) {
	devinfo.init()

	devinfo.cptr.wVideoInputCh = C.WORD(input)
}

func (devinfo *DeviceNetInfo) RemoteVideoInput() int {
	devinfo.init()

	return int(devinfo.cptr.wRemoteVideoInputCh)
}

func (devinfo *DeviceNetInfo) SetRemoteVideoInput(input int) {
	devinfo.init()

	devinfo.cptr.wRemoteVideoInputCh = C.WORD(input)
}

func (devinfo *DeviceNetInfo) VideoOutput() int {
	devinfo.init()

	return int(devinfo.cptr.wVideoOutputCh)
}

func (devinfo *DeviceNetInfo) SetVideoOutput(output int) {
	devinfo.init()

	devinfo.cptr.wVideoOutputCh = C.WORD(output)
}

func (devinfo *DeviceNetInfo) AlarmInput() int {
	devinfo.init()

	return int(devinfo.cptr.wAlarmInputCh)
}

func (devinfo *DeviceNetInfo) SetAlarmInput(alarm int) {
	devinfo.init()

	devinfo.cptr.wAlarmInputCh = C.WORD(alarm)
}

func (devinfo *DeviceNetInfo) AlarmOutput() int {
	devinfo.init()

	return int(devinfo.cptr.wAlarmOutputCh)
}

func (devinfo *DeviceNetInfo) SetAlarmOutput(alarm int) {
	devinfo.init()

	devinfo.cptr.wAlarmOutputCh = C.WORD(alarm)
}

func (devinfo *DeviceNetInfo) NewWordLen() bool {
	devinfo.init()

	return devinfo.cptr.bNewWordLen > 0
}

func (devinfo *DeviceNetInfo) SetNewWordLen(n bool) {
	devinfo.init()

	if n {
		devinfo.cptr.bNewWordLen = 1
	} else {
		devinfo.cptr.bNewWordLen = 0
	}
}

func (devinfo *DeviceNetInfo) NewPassWord() string {
	devinfo.init()

	return C.GoStringN((*C.char)(unsafe.Pointer(&devinfo.cptr.szNewPassWord[0])), C.DH_COMMON_STRING_64)
}

func (devinfo *DeviceNetInfo) SetNewPassWord(password string) {
	devinfo.init()
	var n C.ulong
	if len(password) < C.DH_USER_NAME_LENGTH_EX {
		n = C.ulong(len(password))
	} else {
		n = 88
	}

	C.strncpy((*C.char)(&devinfo.cptr.szNewPassWord[0]), C.CString(password), n)
}

func (devinfo *DeviceNetInfo) InitStatus() int {
	devinfo.init()

	return int(devinfo.cptr.byInitStatus)
}

func (devinfo *DeviceNetInfo) SetInitStatus(n int) {
	devinfo.init()

	devinfo.cptr.byInitStatus = C.BYTE(n)
}

func (devinfo *DeviceNetInfo) PwdResetWay() int {
	devinfo.init()

	return int(devinfo.cptr.byPwdResetWay)
}

func (devinfo *DeviceNetInfo) SetPwdResetWay(n int) {
	devinfo.init()

	devinfo.cptr.byPwdResetWay = C.BYTE(n)
}

func (devinfo *DeviceNetInfo) SpecialAbility() int {
	devinfo.init()

	return int(devinfo.cptr.bySpecialAbility)
}

func (devinfo *DeviceNetInfo) SetSpecialAbility(n int) {
	devinfo.init()

	devinfo.cptr.bySpecialAbility = C.BYTE(n)
}

func (devinfo *DeviceNetInfo) NewDetailType() string {
	devinfo.init()

	return C.GoStringN((*C.char)(unsafe.Pointer(&devinfo.cptr.szNewDetailType[0])), C.DH_COMMON_STRING_64)
}

func (devinfo *DeviceNetInfo) SetNewDetailType(detail string) {
	devinfo.init()
	var n C.ulong
	if len(detail) < C.DH_COMMON_STRING_64 {
		n = C.ulong(len(detail))
	} else {
		n = 88
	}

	C.strncpy((*C.char)(&devinfo.cptr.szNewDetailType[0]), C.CString(detail), n)
}

func (devinfo *DeviceNetInfo) HasNewUserName() bool {
	devinfo.init()

	return devinfo.cptr.bNewUserName > 0
}

func (devinfo *DeviceNetInfo) SetHasNewUserName(n bool) {
	devinfo.init()

	if n {
		devinfo.cptr.bNewUserName = 1
	} else {
		devinfo.cptr.bNewUserName = 0
	}

}

func (devinfo *DeviceNetInfo) NewUserName() string {
	devinfo.init()

	return C.GoStringN((*C.char)(unsafe.Pointer(&devinfo.cptr.szNewUserName[0])), C.DH_COMMON_STRING_64)
}

func (devinfo *DeviceNetInfo) SetNewUserName(username string) {
	devinfo.init()
	var n C.ulong
	if len(username) < C.DH_COMMON_STRING_64 {
		n = C.ulong(len(username))
	} else {
		n = 88
	}

	C.strncpy((*C.char)(&devinfo.cptr.szNewUserName[0]), C.CString(username), n)
}

func (devinfo *DeviceNetInfo) PwdFindVersion() int {
	devinfo.init()

	return int(devinfo.cptr.byPwdFindVersion)
}

func (devinfo *DeviceNetInfo) SetPwdFindVersion(n int) {
	devinfo.init()

	devinfo.cptr.byPwdFindVersion = C.BYTE(n)
}

func (devinfo *DeviceNetInfo) DeviceID() string {
	devinfo.init()

	return C.GoStringN((*C.char)(unsafe.Pointer(&devinfo.cptr.szDeviceID[0])), C.DH_DEV_CUSTOM_DEVICEID_LEN)
}

func (devinfo *DeviceNetInfo) SetDeviceID(deviceid string) {
	devinfo.init()
	var n C.ulong
	if len(deviceid) < C.DH_DEV_CUSTOM_DEVICEID_LEN {
		n = C.ulong(len(deviceid))
	} else {
		n = 88
	}

	C.strncpy((*C.char)(&devinfo.cptr.szDeviceID[0]), C.CString(deviceid), n)
}

func (devinfo *DeviceNetInfo) UnLoginFuncMask() int {
	devinfo.init()

	return int(devinfo.cptr.dwUnLoginFuncMask)
}

func (devinfo *DeviceNetInfo) SetUnLoginFuncMask(mask int) {
	devinfo.init()

	devinfo.cptr.dwUnLoginFuncMask = C.DWORD(mask)
}
