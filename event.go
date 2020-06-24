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
// extern int goAnalyzerDataCallBack(LDWORD dwUser, DWORD dwAlarmType, void* pAlarmInfo, BYTE *pBuffer, DWORD dwBufSize, int nSequence);
//
// int CALLBACK cAnalyzerDataCallBack(LLONG lAnalyzerHandle, DWORD dwAlarmType, void* pAlarmInfo, BYTE *pBuffer, DWORD dwBufSize, LDWORD dwUser, int nSequence, void *reserved) {
// 	if (dwUser != 0) {
//    return goAnalyzerDataCallBack(dwUser, dwAlarmType, pAlarmInfo, pBuffer, dwBufSize, nSequence);
//  }
// 	return 0;
// }
import "C"

import (
	"time"
	"unsafe"

	"github.com/mattn/go-pointer"
)

type PictureExFunc func(client *Client, AlarmType EventIvs, alarmInfo interface{}, frame []byte, seq int) int
type PictureExVistor struct {
	Client   *Client
	Callback PictureExFunc
}

func (client *Client) RealLoadPictureEx(channel int, evt EventIvs, callback PictureExFunc) {
	var (
		v = PictureExVistor{
			Client:   client,
			Callback: callback,
		}
	)

	client.StopLoadPic()

	p := pointer.Save(v)
	client.pictureVisitorp = p

	client.realpicHandle = C.CLIENT_RealLoadPictureEx(C.LLONG(client.LoginID), C.int(channel), C.uint(EventIvsAll), 1, C.fAnalyzerDataCallBack(C.cAnalyzerDataCallBack), C.LDWORD(uintptr(p)), nil)
}

func (client *Client) StopLoadPic() bool {
	if client.realpicHandle > 0 {
		b := C.CLIENT_StopLoadPic(C.LLONG(client.realpicHandle))
		return b > 0
	}
	return false
}

type TrafficParkingInfo struct {
	cptr *C.DEV_EVENT_TRAFFIC_PARKING_INFO
}

func (tpark *TrafficParkingInfo) init() {
	if tpark.cptr != nil {
		return
	}

	tpark.cptr = &C.DEV_EVENT_TRAFFIC_PARKING_INFO{}
}

func (tpark *TrafficParkingInfo) ChannelID() int {
	tpark.init()

	return int(tpark.cptr.nChannelID)
}

func (tpark *TrafficParkingInfo) SetChannelID(ver int) {
	tpark.init()

	tpark.cptr.nChannelID = C.int(ver)
}

func (tpark *TrafficParkingInfo) Name() string {
	tpark.init()

	return GoStr((*C.char)(&tpark.cptr.szName[0]), 128)
}

func (tpark *TrafficParkingInfo) SetName(name string) {
	tpark.init()
	var n C.ulong
	if len(name) < 128 {
		n = C.ulong(len(name))
	} else {
		n = 128
	}

	C.strncpy((*C.char)(unsafe.Pointer(&tpark.cptr.szName)), C.CString(name), n)
}

func (tpark *TrafficParkingInfo) PTS() float64 {
	tpark.init()

	return float64(tpark.cptr.PTS)
}

func (tpark *TrafficParkingInfo) SetPTS(pts float64) {
	tpark.init()

	tpark.cptr.PTS = C.double(pts)
}

func (tpark *TrafficParkingInfo) UTC() time.Time {
	tpark.init()

	return ntex2time(tpark.cptr.UTC)
}

func (tpark *TrafficParkingInfo) SetUTC(utc time.Time) {
	tpark.init()

	tpark.cptr.UTC = time2ntex(utc)
}

func (tpark *TrafficParkingInfo) EventID() int {
	tpark.init()

	return int(tpark.cptr.nEventID)
}

func (tpark *TrafficParkingInfo) SetEventID(evt int) {
	tpark.init()

	tpark.cptr.nEventID = C.int(evt)
}

func (tpark *TrafficParkingInfo) Object() DhObject {
	tpark.init()

	return DhObject{cptr: &tpark.cptr.stuObject}
}

func (tpark *TrafficParkingInfo) SetObject(obj DhObject) {
	tpark.init()
	tpark.cptr.stuObject = *obj.cptr
}

func (tpark *TrafficParkingInfo) Vehicle() DhObject {
	tpark.init()
	vehic := (*C.DH_MSG_OBJECT)(unsafe.Pointer(uintptr(unsafe.Pointer(tpark.cptr)) + unsafe.Offsetof(tpark.cptr.stuObject) + unsafe.Sizeof(C.DH_MSG_OBJECT{})))

	return DhObject{cptr: vehic}
}

func (tpark *TrafficParkingInfo) SetVehicle(obj DhObject) {
	tpark.init()
	obj.init()

	vehic := (*C.DH_MSG_OBJECT)(unsafe.Pointer(uintptr(unsafe.Pointer(tpark.cptr)) + unsafe.Offsetof(tpark.cptr.stuObject) + unsafe.Sizeof(C.DH_MSG_OBJECT{})))
	*vehic = *obj.cptr
}

func (tpark *TrafficParkingInfo) Lane() int {
	tpark.init()

	return int(tpark.cptr.nLane)
}

func (tpark *TrafficParkingInfo) SetLane(lane int) {
	tpark.init()

	tpark.cptr.nLane = C.int(lane)
}

// func (tpark *TrafficParkingInfo) FileInfo() float64 {
// 	tpark.init()

// 	return int(tpark.cptr.PTS)
// }

// func (tpark *TrafficParkingInfo) stuFileInfo(pts float64) {
// 	tpark.init()

// 	tpark.cptr.PTS = C.double(pts)
// }

func (tpark *TrafficParkingInfo) EventAction() int {
	tpark.init()

	return int(tpark.cptr.bEventAction)
}

func (tpark *TrafficParkingInfo) SetEventAction(evt int) {
	tpark.init()

	tpark.cptr.bEventAction = C.BYTE(evt)
}

func (tpark *TrafficParkingInfo) ImageIndex() int {
	tpark.init()

	return int(tpark.cptr.byImageIndex)
}

func (tpark *TrafficParkingInfo) SetImageIndex(idx int) {
	tpark.init()

	tpark.cptr.byImageIndex = C.BYTE(idx)
}

func (tpark *TrafficParkingInfo) StartParkingTime() time.Time {
	tpark.init()

	return ntex2time(tpark.cptr.stuStartParkingTime)
}

func (tpark *TrafficParkingInfo) SetStartParkingTime(t time.Time) {
	tpark.init()

	tpark.cptr.stuStartParkingTime = time2ntex(t)
}

func (tpark *TrafficParkingInfo) Sequence() int {
	tpark.init()

	return int(tpark.cptr.nSequence)
}

func (tpark *TrafficParkingInfo) SetSequence(idx int) {
	tpark.init()

	tpark.cptr.nSequence = C.int(idx)
}

func (tpark *TrafficParkingInfo) AlarmIntervalTime() int {
	tpark.init()

	return int(tpark.cptr.nAlarmIntervalTime)
}

func (tpark *TrafficParkingInfo) SetAlarmIntervalTime(intval int) {
	tpark.init()

	tpark.cptr.nAlarmIntervalTime = C.int(intval)
}

func (tpark *TrafficParkingInfo) ParkingAllowedTime() int {
	tpark.init()

	return int(tpark.cptr.nParkingAllowedTime)
}

func (tpark *TrafficParkingInfo) SetParkingAllowedTime(tim int) {
	tpark.init()

	tpark.cptr.nParkingAllowedTime = C.int(tim)
}

func (tpark *TrafficParkingInfo) DetectRegionNum() int {
	tpark.init()

	return int(tpark.cptr.nDetectRegionNum)
}

func (tpark *TrafficParkingInfo) SetDetectRegionNum(n int) {
	tpark.init()

	tpark.cptr.nDetectRegionNum = C.int(n)
}

func (tpark *TrafficParkingInfo) DetectRegions() []DhPoint {
	tpark.init()
	var n = int(C.DH_MAX_DETECT_REGION_NUM)
	if int(tpark.cptr.nDetectRegionNum) < n {
		n = int(tpark.cptr.nDetectRegionNum)
	}

	var points []DhPoint
	for i := 0; i < n; i++ {
		// tpark.cptr.DetectRegion[i]
		points = append(points, DhPoint{cptr: &tpark.cptr.DetectRegion[i]})
	}

	return points
}

// func (tpark *TrafficParkingInfo) SetDetectRegion(points []DhPoint) {
// 	tpark.init()

// 	tpark.cptr.PTS = C.double(pts)
// }

func (tpark *TrafficParkingInfo) SnapFlagMask() int {
	tpark.init()

	return int(tpark.cptr.dwSnapFlagMask)
}

func (tpark *TrafficParkingInfo) SetSnapFlagMask(n int) {
	tpark.init()

	tpark.cptr.dwSnapFlagMask = C.DWORD(n)
}

// func (tpark *TrafficParkingInfo) stuResolution() float64 {
// 	tpark.init()

// 	return int(tpark.cptr.PTS)
// }

// func (tpark *TrafficParkingInfo) stuResolution(pts float64) {
// 	tpark.init()

// 	tpark.cptr.PTS = C.double(pts)
// }

func (tpark *TrafficParkingInfo) IsExistAlarmRecord() bool {
	tpark.init()

	return tpark.cptr.bIsExistAlarmRecord > 0
}

func (tpark *TrafficParkingInfo) SetExistAlarmRecord(b bool) {
	tpark.init()

	if b {
		tpark.cptr.bIsExistAlarmRecord = 1
	} else {
		tpark.cptr.bIsExistAlarmRecord = 0
	}
}

func (tpark *TrafficParkingInfo) AlarmRecordSize() int {
	tpark.init()

	return int(tpark.cptr.dwAlarmRecordSize)
}

func (tpark *TrafficParkingInfo) SetAlarmRecordSize(size int) {
	tpark.init()

	tpark.cptr.dwAlarmRecordSize = C.DWORD(size)
}

func (tpark *TrafficParkingInfo) AlarmRecordPath() string {
	tpark.init()

	return GoStr((*C.char)(&tpark.cptr.szAlarmRecordPath[0]), C.DH_COMMON_STRING_256)
}

func (tpark *TrafficParkingInfo) SetAlarmRecordPath(name string) {
	tpark.init()
	var n C.ulong
	if len(name) < C.DH_COMMON_STRING_256 {
		n = C.ulong(len(name))
	} else {
		n = C.DH_COMMON_STRING_256
	}

	C.strncpy((*C.char)(unsafe.Pointer(&tpark.cptr.szAlarmRecordPath)), C.CString(name), n)
}

func (tpark *TrafficParkingInfo) FTPPath() string {
	tpark.init()

	return GoStr((*C.char)(&tpark.cptr.szFTPPath[0]), C.DH_COMMON_STRING_256)
}

func (tpark *TrafficParkingInfo) SetFTPPath(name string) {
	tpark.init()
	var n C.ulong
	if len(name) < C.DH_COMMON_STRING_256 {
		n = C.ulong(len(name))
	} else {
		n = C.DH_COMMON_STRING_256
	}

	C.strncpy((*C.char)(unsafe.Pointer(&tpark.cptr.szFTPPath)), C.CString(name), n)
}

// func (tpark *TrafficParkingInfo) stuIntelliCommInfo() float64 {
// 	tpark.init()

// 	return int(tpark.cptr.PTS)
// }

// func (tpark *TrafficParkingInfo) stuIntelliCommInfo(pts float64) {
// 	tpark.init()

// 	tpark.cptr.PTS = C.double(pts)
// }

// func (tpark *TrafficParkingInfo) stuGPSInfo() float64 {
// 	tpark.init()

// 	return int(tpark.cptr.PTS)
// }

// func (tpark *TrafficParkingInfo) stuGPSInfo(pts float64) {
// 	tpark.init()

// 	tpark.cptr.PTS = C.double(pts)
// }

// func (tpark *TrafficParkingInfo) stTrafficCar() float64 {
// 	tpark.init()

// 	return int(tpark.cptr.PTS)
// }

// func (tpark *TrafficParkingInfo) stTrafficCar(pts float64) {
// 	tpark.init()

// 	tpark.cptr.PTS = C.double(pts)
// }

// func (tpark *TrafficParkingInfo) stCommInfo() float64 {
// 	tpark.init()

// 	return int(tpark.cptr.PTS)
// }

// func (tpark *TrafficParkingInfo) stCommInfo(pts float64) {
// 	tpark.init()

// 	tpark.cptr.PTS = C.double(pts)
// }

func (tpark TrafficParkingInfo) String() string {
	return FormatString(&tpark)
}
