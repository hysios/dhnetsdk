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

type MediafileIvsEventParam struct {
	cptr *C.MEDIAFILE_IVS_EVENT_PARAM
}

func (ivs *MediafileIvsEventParam) init() {
	if ivs.cptr != nil {
		return
	}

	ivs.cptr = &C.MEDIAFILE_IVS_EVENT_PARAM{}
	ivs.cptr.dwSize = C.uint(unsafe.Sizeof(*ivs.cptr))
}

func (ivs *MediafileIvsEventParam) ChannelID() int {
	ivs.init()

	return int(ivs.cptr.nChannelID)
}

func (ivs *MediafileIvsEventParam) SetChannelID(chanId int) {
	ivs.init()
	ivs.cptr.nChannelID = C.int(chanId)
}

func (ivs *MediafileIvsEventParam) StartTime() time.Time {
	ivs.init()

	return nt2time(ivs.cptr.stuStartTime)
}

func (ivs *MediafileIvsEventParam) SetStartTime(t time.Time) {
	ivs.init()
	ivs.cptr.stuStartTime = time2nt(t)
}

func (ivs *MediafileIvsEventParam) EndTime() time.Time {
	ivs.init()

	return nt2time(ivs.cptr.stuEndTime)
}

func (ivs *MediafileIvsEventParam) SetEndTime(t time.Time) {
	ivs.init()
	ivs.cptr.stuEndTime = time2nt(t)
}

func (ivs *MediafileIvsEventParam) MediaType() int {
	ivs.init()

	return int(ivs.cptr.nMediaType)
}

func (ivs *MediafileIvsEventParam) SetMediaType(typ int) {
	ivs.init()
	ivs.cptr.nMediaType = C.int(typ)
}

func (ivs *MediafileIvsEventParam) VideoStream() int {
	ivs.init()

	return int(ivs.cptr.nVideoStream)
}

func (ivs *MediafileIvsEventParam) SetVideoStream(typ int) {
	ivs.init()
	ivs.cptr.nVideoStream = C.int(typ)
}

func (ivs *MediafileIvsEventParam) EventTypes() []EventIvs {
	ivs.init()
	var events = make([]EventIvs, int(ivs.cptr.nEventCount))

	for i := range events {
		events[i] = EventIvs(ivs.cptr.nEventLists[i])
	}

	return events
}

func (ivs *MediafileIvsEventParam) SetEventTypes(events []EventIvs) {
	ivs.init()

	if len(events) >= C.MAX_IVS_EVENT_NUM {
		events = events[:C.MAX_IVS_EVENT_NUM]
	}

	ivs.cptr.nEventCount = C.int(len(events))
	for i, evt := range events {
		ivs.cptr.nEventLists[i] = C.int(evt)
	}

}

func (ivs *MediafileIvsEventParam) FlagLists() []EmRecordSnapFlagType {
	ivs.init()
	var flags = make([]EmRecordSnapFlagType, int(ivs.cptr.nFalgCount))

	for i := range flags {
		flags[i] = EmRecordSnapFlagType(ivs.cptr.emFalgLists[i])
	}

	return flags
}

func (ivs *MediafileIvsEventParam) SetFlagLists(flags []EmRecordSnapFlagType) {

	if len(flags) >= C.FLAG_TYPE_MAX {
		flags = flags[:C.FLAG_TYPE_MAX]
	}

	ivs.cptr.nFalgCount = C.int(len(flags))
	for i, flag := range flags {
		ivs.cptr.emFalgLists[i] = C.EM_RECORD_SNAP_FLAG_TYPE(flag)
	}
}

func (ivs *MediafileIvsEventParam) RuleType() int {
	ivs.init()

	return int(ivs.cptr.nRuleType)
}

func (ivs *MediafileIvsEventParam) SetRuleType(typ int) {
	ivs.init()
	ivs.cptr.nRuleType = C.int(typ)
}

func (ivs *MediafileIvsEventParam) Action() NetCrossregionActionInfo {
	ivs.init()

	return NetCrossregionActionInfo(ivs.cptr.emAction)
}

func (ivs *MediafileIvsEventParam) SetAction(act NetCrossregionActionInfo) {
	ivs.init()
	ivs.cptr.emAction = C.NET_CROSSREGION_ACTION_INFO(act)
}

func (ivs *MediafileIvsEventParam) Print() {
	pp.Print(ivs.cptr)
}
