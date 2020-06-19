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
	"time"
	"unsafe"

	"github.com/yudai/pp"
)

// DWORD							dwSize;                 // 结构体大小
// char*							szDirs;                 // 工作目录列表,一次可查询多个目录,为空表示查询所有目录。目录之间以分号分隔,如“/mnt/dvr/sda0;/mnt/dvr/sda1”,szDirs==null 或"" 表示查询所有
// int								nMediaType;             // 文件类型,0:查询任意类型,1:查询jpg图片,2:查询dav
// int								nChannelID;             // 通道号从0开始,-1表示查询所有通道
// NET_TIME						stuStartTime;           // 开始时间
// NET_TIME						stuEndTime;             // 结束时间
// int								nEventLists[MAX_IVS_EVENT_NUM]; // 事件类型列表,参见智能分析事件类型
// int								nEventCount;            // 事件总数
// BYTE							byVideoStream;          // 视频码流 0-未知 1-主码流 2-辅码流1 3-辅码流2 4-辅码流3  5-所有的辅码流类型
// BYTE							bReserved[3];           // 字节对齐
// EM_RECORD_SNAP_FLAG_TYPE		emFalgLists[FLAG_TYPE_MAX]; // 录像或抓图文件标志, 不设置标志表示查询所有文件
// int								nFalgCount;             // 标志总数
// NET_RECORD_CARD_INFO			stuCardInfo;			// 卡号录像信息, emFalgLists包含卡号录像时有效
// int								nUserCount;             // 用户名有效个数
// char							szUserName[MAX_QUERY_USER_NUM][DH_NEW_USER_NAME_LENGTH]; // 用户名
// EM_RESULT_ORDER_TYPE			emResultOrder;			// 查询结果排序方式
// BOOL							bTime;                  // 是否按时间查询
// NET_EM_COMBINATION_MODE			emCombination;			// 查询结果是否合并录像文件
// EVENT_INFO						stuEventInfo[16];		// 事件信息（定制），当查询为 DH_FILE_QUERY_FILE_EX 类型时有效
// int								nEventInfoCount;		// stuEventInfo 个数
type NetInMediaQueryFile struct {
	cptr *C.NET_IN_MEDIA_QUERY_FILE
}

func (netfile *NetInMediaQueryFile) init() {
	if netfile.cptr != nil {
		return
	}

	netfile.cptr = &C.NET_IN_MEDIA_QUERY_FILE{}
	netfile.cptr.dwSize = C.uint(unsafe.Sizeof(*netfile.cptr))
}

func (netfile *NetInMediaQueryFile) Dirs() string {
	netfile.init()

	return C.GoString(netfile.cptr.szDirs)
}

func (netfile *NetInMediaQueryFile) SetDirs(dir string) {
	netfile.init()
	netfile.cptr.szDirs = C.CString(dir)
}

func (netfile *NetInMediaQueryFile) MediaType() int {
	netfile.init()

	return int(netfile.cptr.nMediaType)
}

func (netfile *NetInMediaQueryFile) SetMediaType(typ int) {
	netfile.init()
	netfile.cptr.nMediaType = C.int(typ)
}

func (netfile *NetInMediaQueryFile) ChannelID() int {
	netfile.init()

	return int(netfile.cptr.nChannelID)
}

func (netfile *NetInMediaQueryFile) SetChannelID(chanId int) {
	netfile.init()
	netfile.cptr.nChannelID = C.int(chanId)
}

func (netfile *NetInMediaQueryFile) StartTime() time.Time {
	netfile.init()

	return nt2time(netfile.cptr.stuStartTime)
}

func (netfile *NetInMediaQueryFile) SetStartTime(t time.Time) {
	netfile.init()
	netfile.cptr.stuStartTime = time2nt(t)
}

func (netfile *NetInMediaQueryFile) EndTime() time.Time {
	netfile.init()

	return nt2time(netfile.cptr.stuEndTime)
}

func (netfile *NetInMediaQueryFile) SetEndTime(t time.Time) {
	netfile.init()
	netfile.cptr.stuEndTime = time2nt(t)
}

func (netfile *NetInMediaQueryFile) EventTypes() []EventIvs {
	netfile.init()
	var events = make([]EventIvs, int(netfile.cptr.nEventCount))

	for i := range events {
		events[i] = EventIvs(netfile.cptr.nEventLists[i])
	}

	return events
}

func (netfile *NetInMediaQueryFile) SetEventTypes(events []EventIvs) {
	netfile.init()
	if len(events) >= C.MAX_IVS_EVENT_NUM {
		events = events[:C.MAX_IVS_EVENT_NUM]
	}

	netfile.cptr.nEventCount = C.int(len(events))
	for i, evt := range events {
		netfile.cptr.nEventLists[i] = C.int(evt)
	}
}

func (netfile *NetInMediaQueryFile) VideoStream() int {
	netfile.init()

	return int(netfile.cptr.byVideoStream)
}

func (netfile *NetInMediaQueryFile) SetVideoStream(typ int) {
	netfile.init()
	netfile.cptr.byVideoStream = C.BYTE(typ)
}

func (netfile *NetInMediaQueryFile) FlagLists() []EmRecordSnapFlagType {
	netfile.init()
	var flags = make([]EmRecordSnapFlagType, int(netfile.cptr.nFalgCount))

	for i := range flags {
		flags[i] = EmRecordSnapFlagType(netfile.cptr.emFalgLists[i])
	}

	return flags
}

func (netfile *NetInMediaQueryFile) SetFlagLists(flags []EmRecordSnapFlagType) {

	if len(flags) >= C.FLAG_TYPE_MAX {
		flags = flags[:C.FLAG_TYPE_MAX]
	}

	netfile.cptr.nFalgCount = C.int(len(flags))
	for i, flag := range flags {
		netfile.cptr.emFalgLists[i] = C.EM_RECORD_SNAP_FLAG_TYPE(flag)
	}
}

func (netfile *NetInMediaQueryFile) Print() {
	pp.Print(netfile.cptr)
}
