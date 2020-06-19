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
	"errors"
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

// MEDIAFILE_FACE_DETECTION_PARAM

type EmFileQueryType int

type cobject interface {
	init()
}

const (
	DhFileQueryTrafficcar        EmFileQueryType = C.DH_FILE_QUERY_TRAFFICCAR         // 交通车辆信息,对应结构体为MEDIA_QUERY_TRAFFICCAR_PARAM
	DhFileQueryAtm               EmFileQueryType = C.DH_FILE_QUERY_ATM                // ATM信息
	DhFileQueryAtmtxn            EmFileQueryType = C.DH_FILE_QUERY_ATMTXN             // ATM交易信息
	DhFileQueryFace              EmFileQueryType = C.DH_FILE_QUERY_FACE               // 人脸信息 MEDIAFILE_FACERECOGNITION_PARAM 和 MEDIAFILE_FACERECOGNITION_INFO
	DhFileQueryFile              EmFileQueryType = C.DH_FILE_QUERY_FILE               // 文件信息对应 NET_IN_MEDIA_QUERY_FILE 和 NET_OUT_MEDIA_QUERY_FILE
	DhFileQueryTrafficcarEx      EmFileQueryType = C.DH_FILE_QUERY_TRAFFICCAR_EX      // 交通车辆信息, 扩展DH_FILE_QUERY_TRAFFICCAR, 支持更多的字段，对应结构体为MEDIA_QUERY_TRAFFICCAR_PARAM_EX
	DhFileQueryFaceDetection     EmFileQueryType = C.DH_FILE_QUERY_FACE_DETECTION     // 人脸检测事件信息 MEDIAFILE_FACE_DETECTION_PARAM 和 MEDIAFILE_FACE_DETECTION_INFO
	DhFileQueryIvsEvent          EmFileQueryType = C.DH_FILE_QUERY_IVS_EVENT          // 智能事件信息 MEDIAFILE_IVS_EVENT_PARAM 和 MEDIAFILE_IVS_EVENT_INFO
	DhFileQueryAnalyseObject     EmFileQueryType = C.DH_FILE_QUERY_ANALYSE_OBJECT     // 智能分析其他物体(人和车除外) MEDIAFILE_ANALYSE_OBJECT_PARAM 和 MEDIAFILE_ANALYSE_OBJECT_INFO
	DhFileQueryMptRecordFile     EmFileQueryType = C.DH_FILE_QUERY_MPT_RECORD_FILE    // MPT设备的录像文件 MEDIAFILE_MPT_RECORD_FILE_PARAM 和 MEDIAFILE_MPT_RECORD_FILE_INFO
	DhFileQueryXrayDetection     EmFileQueryType = C.DH_FILE_QUERY_XRAY_DETECTION     // X光检包裹信息对应 MEDIAFILE_XRAY_DETECTION_PARAM 和 MEDIAFILE_XRAY_DETECTION_INFO
	DhFileQueryHumanTrait        EmFileQueryType = C.DH_FILE_QUERY_HUMAN_TRAIT        // 人体检测 MEDIAFILE_HUMAN_TRAIT_PARAM 和 MEDIAFILE_HUMAN_TRAIT_INFO
	DhFileQueryNonmotor          EmFileQueryType = C.DH_FILE_QUERY_NONMOTOR           // 非机动车查询,  MEDIAFILE_NONMOTOR_PARAM 和 MEDIAFILE_NONMOTOR_INFO
	DhFileQueryDoorcontrolRecord EmFileQueryType = C.DH_FILE_QUERY_DOORCONTROL_RECORD // 门打开事件查询, MEDIAFILE_DOORCONTROL_RECORD_PARAM 和 MEDIAFILE_DOORCONTROL_RECORD_INFO
	DhFileQueryFacebodyDetect    EmFileQueryType = C.DH_FILE_QUERY_FACEBODY_DETECT    // 人像检测查询，MEDIAFILE_FACEBODY_DETECT_PARAM 和 MEDIAFILE_FACEBODY_DETECT_INFO
	DhFileQueryFacebodyAnalyse   EmFileQueryType = C.DH_FILE_QUERY_FACEBODY_ANALYSE   // 人像识别查询，MEDIAFILE_FACEBODY_ANALYSE_PARAM 和 MEDIAFILE_FACEBODY_ANALYSE_INFO
	DhFileQueryFileEx            EmFileQueryType = C.DH_FILE_QUERY_FILE_EX            // 文件信息扩展（定制），对应 NET_IN_MEDIA_QUERY_FILE 和 NET_OUT_MEDIA_QUERY_FILE
	// 此时 NET_IN_MEDIA_QUERY_FILE 中的 stuEventInfo 字段有效, nEventLists 及 nEventCount字段无效
)

type EmRecordSnapFlagType int

const (
	FlagTypeTiming             EmRecordSnapFlagType = C.FLAG_TYPE_TIMING               // //定时文件
	FlagTypeManual             EmRecordSnapFlagType = C.FLAG_TYPE_MANUAL               // //手动文件
	FlagTypeMarked             EmRecordSnapFlagType = C.FLAG_TYPE_MARKED               // //重要文件
	FlagTypeEvent              EmRecordSnapFlagType = C.FLAG_TYPE_EVENT                // //事件文件
	FlagTypeMosaic             EmRecordSnapFlagType = C.FLAG_TYPE_MOSAIC               // //合成图片
	FlagTypeCutout             EmRecordSnapFlagType = C.FLAG_TYPE_CUTOUT               // //抠图图片
	FlagTypeLeaveWord          EmRecordSnapFlagType = C.FLAG_TYPE_LEAVE_WORD           // //留言文件
	FlagTypeTalkbackLocalSide  EmRecordSnapFlagType = C.FLAG_TYPE_TALKBACK_LOCAL_SIDE  // //对讲本地方文件
	FlagTypeTalkbackRemoteSide EmRecordSnapFlagType = C.FLAG_TYPE_TALKBACK_REMOTE_SIDE // //对讲远程方文件
	FlagTypeSynopsisVideo      EmRecordSnapFlagType = C.FLAG_TYPE_SYNOPSIS_VIDEO       // //浓缩视频
	FlagTypeOriginalVideo      EmRecordSnapFlagType = C.FLAG_TYPE_ORIGINAL_VIDEO       // //原始视频
	FlagTypePreOriginalVideo   EmRecordSnapFlagType = C.FLAG_TYPE_PRE_ORIGINAL_VIDEO   // //已经预处理的原始视频
	FlagTypeBlackPlate         EmRecordSnapFlagType = C.FLAG_TYPE_BLACK_PLATE          // //黑名单图片
	FlagTypeOriginalPic        EmRecordSnapFlagType = C.FLAG_TYPE_ORIGINAL_PIC         // //原始图片
	FlagTypeCard               EmRecordSnapFlagType = C.FLAG_TYPE_CARD                 // //卡号录像
	FlagTypeMax                EmRecordSnapFlagType = C.FLAG_TYPE_MAX
)

type NetCrossregionActionInfo int

const (
	EmCrossregionActionUnknow    NetCrossregionActionInfo = C.EM_CROSSREGION_ACTION_UNKNOW
	EmCrossregionActionInside    NetCrossregionActionInfo = C.EM_CROSSREGION_ACTION_INSIDE    //在区域内
	EmCrossregionActionCross     NetCrossregionActionInfo = C.EM_CROSSREGION_ACTION_CROSS     //穿越区域
	EmCrossregionActionAppear    NetCrossregionActionInfo = C.EM_CROSSREGION_ACTION_APPEAR    //出现
	EmCrossregionActionDisappear NetCrossregionActionInfo = C.EM_CROSSREGION_ACTION_DISAPPEAR //消失
)

func (cli *Client) FindFileEx(mediaType EmFileQueryType, queryCondition interface{}, waittime time.Duration) (findfile *FindFile, err error) {
	var (
		mediaQuery cobject
		ok         bool
	)

	if mediaQuery, ok = queryCondition.(cobject); !ok {
		return nil, errors.New("invalid query object")
	}
	mediaQuery.init()
	v := reflect.ValueOf(mediaQuery)
	v = reflect.Indirect(v)
	cval := v.FieldByName("cptr")

	ffHandle := C.CLIENT_FindFileEx(C.long(cli.LoginID),
		C.EM_FILE_QUERY_TYPE(DhFileQueryTrafficcar),
		unsafe.Pointer(cval.UnsafeAddr()),
		nil,
		C.int(waittime/time.Millisecond))
	if ffHandle == 0 {
		return nil, fmt.Errorf("FindFileEx failed %s", Err(LastErrorCode()))
	}

	return &FindFile{FindFileHandle: int(ffHandle)}, nil
}

func (cli *Client) FindFileTraffic(mediaType EmFileQueryType, mediaQuery *MediaQueryTrafficcar, waittime time.Duration) (findfile *FindFile, err error) {
	ffHandle := C.CLIENT_FindFileEx(C.long(cli.LoginID),
		C.EM_FILE_QUERY_TYPE(mediaType),
		unsafe.Pointer(mediaQuery.cptr),
		nil,
		C.int(waittime/time.Millisecond))
	if ffHandle == 0 {
		return nil, fmt.Errorf("FindFileEx failed %s", Err(LastErrorCode()))
	}

	return &FindFile{FindFileHandle: int(ffHandle)}, nil
}

func (cli *Client) FindFileCar(mediaType EmFileQueryType, query *MediafileIvsEventParam, waittime time.Duration) (findfile *FindFile, err error) {
	// mediaQuery.

	ffHandle := C.CLIENT_FindFileEx(C.long(cli.LoginID),
		C.EM_FILE_QUERY_TYPE(mediaType),
		unsafe.Pointer(query.cptr),
		nil,
		C.int(waittime/time.Millisecond))
	if ffHandle == 0 {
		return nil, fmt.Errorf("FindFileEx failed %s", Err(LastErrorCode()))
	}

	return &FindFile{FindFileHandle: int(ffHandle)}, nil
}

func (cli *Client) FindFileNetFiles(mediaType EmFileQueryType, query *NetInMediaQueryFile, waittime time.Duration) (findfile *FindFile, err error) {
	// mediaQuery.

	ffHandle := C.CLIENT_FindFileEx(C.long(cli.LoginID),
		C.EM_FILE_QUERY_TYPE(mediaType),
		unsafe.Pointer(query.cptr),
		nil,
		C.int(waittime/time.Millisecond))
	if ffHandle == 0 {
		return nil, fmt.Errorf("FindFileEx failed %s", Err(LastErrorCode()))
	}

	return &FindFile{FindFileHandle: int(ffHandle)}, nil
}

func (cli *Client) FindFileTrafficEx(mediaType EmFileQueryType, query *MediaQueryTrafficcarEx, waittime time.Duration) (findfile *FindFile, err error) {
	// mediaQuery.

	ffHandle := C.CLIENT_FindFileEx(C.long(cli.LoginID),
		C.EM_FILE_QUERY_TYPE(mediaType),
		unsafe.Pointer(query.cptr),
		nil,
		C.int(waittime/time.Millisecond))
	if ffHandle == 0 {
		return nil, fmt.Errorf("FindFileEx failed %s", Err(LastErrorCode()))
	}

	return &FindFile{FindFileHandle: int(ffHandle)}, nil
}

type MediaQueryTrafficcarArray []C.MEDIA_QUERY_TRAFFICCAR_PARAM

// CLIENT_NET_API int    CALL_METHOD CLIENT_FindNextFileEx(LLONG lFindHandle, int nFilecount, void* pMediaFileInfo, int maxlen, void *reserved, int waittime);
func (findfile *FindFile) Next(maxcount int, mediaInfo interface{}, waittime time.Duration) (err error) {
	var media = reflect.ValueOf(mediaInfo)
	if media.Kind() != reflect.Ptr {
		return errors.New("mediaInfo must a can assignment value")
	}

	media = reflect.Indirect(media)
	if !(media.Kind() == reflect.Slice || media.Kind() == reflect.Array) {
		return errors.New("must a slice or array object")
	}

	switch v := mediaInfo.(type) {
	case MediaQueryTrafficcarArray:
		{
			var (
				size = uintptr(maxcount) * unsafe.Sizeof(MediaQueryTrafficcar{})
				ret  = C.CLIENT_FindNextFileEx(
					C.long(findfile.FindFileHandle),
					C.int(maxcount),
					unsafe.Pointer(&v[0]),
					C.int(size), nil, C.int(waittime/time.Millisecond),
				)
			)
			return SuccessOrErr(ret)
		}
	default:
		panic("nonimplement")
	}

	return
}

type FindFile struct {
	FindFileHandle int
}

func (find *FindFile) TotalCount(waittime time.Duration) (total int, err error) {

	err = SuccessOrErr(C.CLIENT_GetTotalFileCount(
		C.long(find.FindFileHandle),
		(*C.int)(unsafe.Pointer(&total)), nil,
		C.int(waittime/time.Millisecond),
	))
	return
}

func (find *FindFile) Close() error {
	return SuccessOrErr(C.CLIENT_FindCloseEx(C.long(find.FindFileHandle)))
}
