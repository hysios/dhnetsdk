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
)

// typedef struct
// {
//     int                 nChannelID;                     // 通道号从0开始,-1表示查询所有通道
//     NET_TIME            StartTime;                      // 开始时间
//     NET_TIME            EndTime;                        // 结束时间
//     int                 nMediaType;                     // 文件类型,0:任意类型, 1:jpg图片, 2:dav文件
//     int                 nEventType;                     // 事件类型,详见"智能分析事件类型", 0:表示查询任意事件,此参数废弃,请使用pEventTypes
//     char                szPlateNumber[32];              // 车牌号, "\0"则表示查询任意车牌号
//     int                 nSpeedUpperLimit;               // 查询的车速范围; 速度上限 单位: km/h
//     int                 nSpeedLowerLimit;               // 查询的车速范围; 速度下限 单位: km/h
//     BOOL                bSpeedLimit;                    // 是否按速度查询; TRUE:按速度查询,nSpeedUpperLimit和nSpeedLowerLimit有效。
//     DWORD               dwBreakingRule;                 // 违章类型：
//                                                         // 当事件类型为 EVENT_IVS_TRAFFICGATE时
//                                                         //        第一位:逆行;  第二位:压线行驶; 第三位:超速行驶;
//                                                         //        第四位：欠速行驶; 第五位:闯红灯;
//                                                         // 当事件类型为 EVENT_IVS_TRAFFICJUNCTION
//                                                         //        第一位:闯红灯;  第二位:不按规定车道行驶;
//                                                         //        第三位:逆行; 第四位：违章掉头;
//                                                         //        第五位:压线行驶;
//     char                szPlateType[32];                // 车牌类型,"Unknown" 未知,"Normal" 蓝牌黑牌,"Yellow" 黄牌,"DoubleYellow" 双层黄尾牌,"Police" 警牌"Armed" 武警牌,
//                                                         // "Military" 部队号牌,"DoubleMilitary" 部队双层,"SAR" 港澳特区号牌,"Trainning" 教练车号牌
//                                                         // "Personal" 个性号牌,"Agri" 农用牌,"Embassy" 使馆号牌,"Moto" 摩托车号牌,"Tractor" 拖拉机号牌,"Other" 其他号牌
// 														// "Civilaviation"民航号牌,"Black"黑牌
// 														// "PureNewEnergyMicroCar"纯电动新能源小车,"MixedNewEnergyMicroCar,"混合新能源小车,"PureNewEnergyLargeCar",纯电动新能源大车
// 														// "MixedNewEnergyLargeCar"混合新能源大车
//     char                szPlateColor[16];               // 车牌颜色, "Blue"蓝色,"Yellow"黄色, "White"白色,"Black"黑色
//     char                szVehicleColor[16];             // 车身颜色:"White"白色, "Black"黑色, "Red"红色, "Yellow"黄色, "Gray"灰色, "Blue"蓝色,"Green"绿色
//     char                szVehicleSize[16];              // 车辆大小类型:"Light-duty":小型车;"Medium":中型车; "Oversize":大型车; "Unknown": 未知
//     int                 nGroupID;                       // 事件组编号(此值>=0时有效)
//     short               byLane;                         // 车道号(此值>=0时表示具体车道,-1表示所有车道,即不下发此字段)
//     BYTE                byFileFlag;                     // 文件标志, 0xFF-使用nFileFlagEx, 0-表示所有录像, 1-定时文件, 2-手动文件, 3-事件文件, 4-重要文件, 5-合成文件
//     BYTE                byRandomAccess;                 // 是否需要在查询过程中随意跳转,0-不需要,1-需要
//     int                 nFileFlagEx;                    // 文件标志, 按位表示: bit0-定时文件, bit1-手动文件, bit2-事件文件, bit3-重要文件, bit4-合成文件, bit5-黑名单图片 0xFFFFFFFF-所有录像
//     int                 nDirection;                     // 车道方向（车开往的方向）    0-北 1-东北 2-东 3-东南 4-南 5-西南 6-西 7-西北 8-未知 -1-所有方向
//     char*               szDirs;                         // 工作目录列表,一次可查询多个目录,为空表示查询所有目录。目录之间以分号分隔,如“/mnt/dvr/sda0;/mnt/dvr/sda1”,szDirs==null 或"" 表示查询所有
//     int*                pEventTypes;                    // 待查询的事件类型数组指针,事件类型,详见"智能分析事件类型",若为NULL则认为查询所有事件（缓冲需由用户申请）
//     int                 nEventTypeNum;                  // 事件类型数组大小
//     char*               pszDeviceAddress;               // 设备地址, NULL表示该字段不起作用
//     char*               pszMachineAddress;              // 机器部署地点, NULL表示该字段不起作用
//     char*               pszVehicleSign;                 // 车辆标识, 例如 "Unknown"-未知, "Audi"-奥迪, "Honda"-本田... NULL表示该字段不起作用
// 	WORD                wVehicleSubBrand;               // 车辆子品牌 需要通过映射表得到真正的子品牌 映射表详见开发手册
//     WORD                wVehicleYearModel;              // 车辆品牌年款 需要通过映射表得到真正的年款 映射表详见开发手册
// 	EM_SAFE_BELT_STATE	emSafeBeltState;				// 安全带状态
// 	EM_CALLING_STATE	emCallingState;					// 打电话状态
// 	EM_ATTACHMENT_TYPE	emAttachMentType;				// 车内饰品类型
// 	EM_CATEGORY_TYPE	emCarType;						// 车辆类型
// 	int                 bReserved[12];                  // 保留字段
// } MEDIA_QUERY_TRAFFICCAR_PARAM;

// DWORD               dwSize;                         // 结构体大小
// // 查询过滤条件
// int                 nChannelID;                     // 通道号
// NET_TIME            stuStartTime;                   // 起始时间
// NET_TIME            stuEndTime;                     // 结束时间
// int                 nMediaType;                     // 文件类型,0:任意类型, 1:jpg图片, 2:dav文件

// int                 nVideoStream;                    // 视频码流 0-未知 1-主码流 2-辅码流1 3-辅码流2 4-辅码流3
// int                 nEventLists[MAX_IVS_EVENT_NUM];  // 事件类型列表,参见智能分析事件类型
// int                 nEventCount;                     // 事件总数
// EM_RECORD_SNAP_FLAG_TYPE emFalgLists[FLAG_TYPE_MAX]; // 录像或抓图文件标志, 不设置标志表示查询所有文件
// int                 nFalgCount;                      // 标志总数

// int                 nRuleType;                      // 智能分析事件名, 事件类型,详见"智能分析事件类型"
// NET_CROSSREGION_ACTION_INFO emAction;               // 事件动作
// int					nIvsObjectNum;					// 对象类型个数
// EM_MEDIAFILE_IVS_OBJECT		emIvsObject[DH_MAX_OBJECT_LIST];	// 规则触发的对象类型

type MediaQueryTrafficcar struct {
	cptr *C.MEDIA_QUERY_TRAFFICCAR_PARAM
}

func (media *MediaQueryTrafficcar) init() {
	if media.cptr != nil {
		return
	}

	media.cptr = &C.MEDIA_QUERY_TRAFFICCAR_PARAM{}

}

func (media *MediaQueryTrafficcar) ChannelID() int {
	media.init()
	return int(media.cptr.nChannelID)
}

func (media *MediaQueryTrafficcar) SetChannelID(chanId int) {
	media.init()

	media.cptr.nChannelID = C.int(chanId)
}

func (media *MediaQueryTrafficcar) StartTime() time.Time {
	media.init()

	return nt2time(media.cptr.StartTime)
}

func (media *MediaQueryTrafficcar) SetStartTime(t time.Time) {
	media.init()

	media.cptr.StartTime = time2nt(t)
}

func (media *MediaQueryTrafficcar) EndTime() time.Time {
	media.init()

	return nt2time(media.cptr.EndTime)
}

func (media *MediaQueryTrafficcar) SetEndTime(t time.Time) {
	media.init()

	media.cptr.EndTime = time2nt(t)
}

func (media *MediaQueryTrafficcar) MediaType() int {
	media.init()

	return int(media.cptr.nMediaType)
}

func (media *MediaQueryTrafficcar) SetMediaType(typ int) {
	media.init()

	media.cptr.nMediaType = C.int(typ)
}

func (media *MediaQueryTrafficcar) EventType() int {
	media.init()

	return int(media.cptr.nEventType)
}

func (media *MediaQueryTrafficcar) SetEventType(typ int) {
	media.init()

	media.cptr.nEventType = C.int(typ)
}

//     int*                pEventTypes;                    // 待查询的事件类型数组指针,事件类型,详见"智能分析事件类型",若为NULL则认为查询所有事件（缓冲需由用户申请）
//     int                 nEventTypeNum;
func (media *MediaQueryTrafficcar) EventTypes() []EventIvs {
	media.init()
	var events = make([]EventIvs, int(media.cptr.nEventTypeNum))

	for i := range events {
		base := uintptr(unsafe.Pointer(media.cptr.pEventTypes))
		offset := uintptr(i) * unsafe.Sizeof(*media.cptr.pEventTypes)

		events[i] = EventIvs(*(*C.int)(unsafe.Pointer(base + offset)))
	}

	return events
}

func (media *MediaQueryTrafficcar) SetEventTypes(events []EventIvs) {
	media.init()
	var base = C.malloc(C.ulong(uintptr(media.cptr.nEventTypeNum) * unsafe.Sizeof(*media.cptr.pEventTypes)))

	media.cptr.pEventTypes = (*C.int)(base)
	media.cptr.nEventTypeNum = C.int(len(events))
	for i, evt := range events {
		offset := (*C.int)(unsafe.Pointer(uintptr(base) + uintptr(i)*unsafe.Sizeof(*media.cptr.pEventTypes)))
		*offset = C.int(evt)
	}
}

func (media *MediaQueryTrafficcar) PlateNumber() string {
	media.init()

	return C.GoStringN(&media.cptr.szPlateNumber[0], 32)
}

func (media *MediaQueryTrafficcar) SetPlateNumber(plate string) {
	media.init()

	C.strncpy((*C.char)(unsafe.Pointer(&media.cptr.szPlateNumber)), C.CString(plate), 32)
}

func (media *MediaQueryTrafficcar) FileFlag() int {
	media.init()

	return int(media.cptr.byFileFlag)
}

func (media *MediaQueryTrafficcar) SetFileFlag(typ int) {
	media.init()

	media.cptr.byFileFlag = C.BYTE(typ)
}

func (media *MediaQueryTrafficcar) FileFlagEx() int {
	media.init()

	return int(media.cptr.nFileFlagEx)
}

func (media *MediaQueryTrafficcar) SetFileFlagEx(typ int) {
	media.init()

	media.cptr.nFileFlagEx = C.int(typ)
}
