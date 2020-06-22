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
	"image/color"
	"time"
	"unsafe"
)

type DhObject struct {
	cptr *C.DH_MSG_OBJECT
}

func (obj *DhObject) init() {
	if obj.cptr != nil {
		return
	}

	obj.cptr = &C.DH_MSG_OBJECT{}
}

func (obj *DhObject) ObjectID() int {
	obj.init()

	return int(obj.cptr.nObjectID)
}

func (obj *DhObject) SetObjectID(id int) {
	obj.init()

	obj.cptr.nObjectID = C.int(id)
}

func (obj *DhObject) ObjectType() string {
	obj.init()
	return GoStr((*C.char)(&obj.cptr.szObjectType[0]), 128)
}

func (obj *DhObject) SetObjectType(objtyp string) {
	obj.init()
	var n C.ulong
	if len(objtyp) < 128 {
		n = C.ulong(len(objtyp))
	} else {
		n = 128
	}

	C.strncpy((*C.char)(unsafe.Pointer(&obj.cptr.szObjectType)), C.CString(objtyp), n)
}

func (obj *DhObject) Confidence() int {
	obj.init()

	return int(obj.cptr.nConfidence)
}

func (obj *DhObject) SetConfidence(conf int) {
	obj.init()

	obj.cptr.nConfidence = C.int(conf)
}

func (obj *DhObject) Action() int {
	obj.init()

	return int(obj.cptr.nAction)
}

func (obj *DhObject) SetAction(act int) {
	obj.init()

	obj.cptr.nAction = C.int(act)
}

func (obj *DhObject) BoundingBox() DhRect {
	obj.init()
	boundBox := uintptr(unsafe.Pointer(obj.cptr)) + 140
	box := (*C.DH_RECT)(unsafe.Pointer(boundBox))
	return DhRect{cptr: box}
}

// func (obj *DhObject) SetBoundingBox(rect DhRect) {
// 	obj.init()

// 	rect.init()

// 	rboundBox := uintptr(unsafe.Pointer(rect.cptr)) + 140
// 	rbox := (*C.BoundingBox)(unsafe.Pointer(rboundBox))

// 	boundBox := uintptr(unsafe.Pointer(obj.cptr)) + 140
// 	box := (*C.BoundingBox)(unsafe.Pointer(boundBox))

// 	// 	if rect.cptr != nil {
// 	// 		obj.cptr.BoundingBox = *rect.cptr
// 	// 	}
// }

func (obj *DhObject) Center() DhPoint {
	obj.init()
	return DhPoint{cptr: &obj.cptr.Center}
}

func (obj *DhObject) SetCenter(point DhPoint) {
	obj.init()

	if point.cptr != nil {
		obj.cptr.Center = *point.cptr
	}
}

func (obj *DhObject) PolygonNum() int {
	obj.init()

	return int(obj.cptr.nPolygonNum)
}

func (obj *DhObject) SetPolygonNum(n int) {
	obj.init()

	obj.cptr.nPolygonNum = C.int(n)
}

func (obj *DhObject) Contours() []DhPoint {
	obj.init()
	var n = 0
	if obj.PolygonNum() < C.DH_MAX_POLYGON_NUM {
		n = obj.PolygonNum()
	} else {
		n = int(C.DH_MAX_POLYGON_NUM)
	}

	var points []DhPoint
	for i := 0; i < n; i++ {
		points = append(points, DhPoint{cptr: &obj.cptr.Contour[i]})
	}

	return points
}

func (obj *DhObject) SetContours(points []DhPoint) {
	obj.init()
	var n = 0

	if obj.PolygonNum() < C.DH_MAX_POLYGON_NUM {
		n = obj.PolygonNum()
	} else {
		n = int(C.DH_MAX_POLYGON_NUM)
	}

	for i, point := range points[:n] {
		point.init()
		obj.cptr.Contour[i] = *point.cptr
	}
}

func (obj *DhObject) MainColor() color.Color {
	obj.init()

	return hex2rgb(int(obj.cptr.rgbaMainColor))
}

func (obj *DhObject) SetMainColor(clr color.RGBA) {
	obj.init()

	obj.cptr.rgbaMainColor = C.DWORD(rgb2hex(clr))
}

func (obj *DhObject) Text() string {
	obj.init()

	return GoStr((*C.char)(&obj.cptr.szText[0]), 128)
}

func (obj *DhObject) SetText(txt string) {
	obj.init()
	var n C.ulong
	if len(txt) < 128 {
		n = C.ulong(len(txt))
	} else {
		n = 128
	}

	C.strncpy((*C.char)(unsafe.Pointer(&obj.cptr.szText)), C.CString(txt), n)
}

func (obj *DhObject) ObjectSubType() string {
	obj.init()

	return GoStr((*C.char)(&obj.cptr.szObjectSubType[0]), 128)
}

func (obj *DhObject) SetObjectSubType(typ string) {
	obj.init()
	var n C.ulong
	if len(typ) < 128 {
		n = C.ulong(len(typ))
	} else {
		n = 128
	}

	C.strncpy((*C.char)(unsafe.Pointer(&obj.cptr.szObjectSubType)), C.CString(typ), n)
}

func (obj *DhObject) ColorLogoIndex() int {
	obj.init()

	return int(obj.cptr.wColorLogoIndex)
}

func (obj *DhObject) SetColorLogoIndex(idx int) {
	obj.init()

	obj.cptr.wColorLogoIndex = C.WORD(idx)
}

func (obj *DhObject) SubBrand() int {
	obj.init()

	return int(obj.cptr.wSubBrand)
}

func (obj *DhObject) SetSubBrand(idx int) {
	obj.init()

	obj.cptr.wColorLogoIndex = C.WORD(idx)
}

func (obj *DhObject) PicEnble() bool {
	obj.init()

	return bool(obj.cptr.bPicEnble)
}

func (obj *DhObject) SetPicEnble(enb bool) {
	obj.init()

	obj.cptr.bPicEnble = C.bool(enb)
}

func (obj *DhObject) PicInfo() DhPicInfo {
	obj.init()

	picinfo := (*C.DH_PIC_INFO)(unsafe.Pointer(uintptr(unsafe.Pointer(obj.cptr)) + unsafe.Offsetof(obj.cptr.bPicEnble) + 1))
	return DhPicInfo{cptr: picinfo}
}

func (obj *DhObject) SetPicInfo(info DhPicInfo) {
	obj.init()

	picinfo := (*C.DH_PIC_INFO)(unsafe.Pointer(uintptr(unsafe.Pointer(obj.cptr)) + unsafe.Offsetof(obj.cptr.bPicEnble) + 1))
	*picinfo = *info.cptr
}

func (obj *DhObject) ShotFrame() bool {
	obj.init()

	return bool(obj.cptr.bShotFrame)
}

func (obj *DhObject) SetShotFrame(enb bool) {
	obj.init()

	obj.cptr.bShotFrame = C.bool(enb)
}

func (obj *DhObject) Color() bool {
	obj.init()

	return bool(obj.cptr.bColor)
}

func (obj *DhObject) SetColor(enb bool) {
	obj.init()

	obj.cptr.bColor = C.bool(enb)
}

func (obj *DhObject) TimeType() int {
	obj.init()

	return int(obj.cptr.byTimeType)
}

func (obj *DhObject) SetTimeType(n int) {
	obj.init()

	obj.cptr.byTimeType = C.BYTE(n)
}

func (obj *DhObject) CurrentTime() time.Time {
	obj.init()

	return ntex2time(obj.cptr.stuCurrentTime)
}

func (obj *DhObject) SetCurrentTime(t time.Time) {
	obj.init()

	obj.cptr.stuCurrentTime = time2ntex(t)
}

func (obj *DhObject) StartTime() time.Time {
	obj.init()

	return ntex2time(obj.cptr.stuStartTime)
}

func (obj *DhObject) SetStartTime(t time.Time) {
	obj.init()

	obj.cptr.stuStartTime = time2ntex(t)
}

func (obj *DhObject) EndTime() time.Time {
	obj.init()

	return ntex2time(obj.cptr.stuEndTime)
}

func (obj *DhObject) SetEndTime(t time.Time) {
	obj.init()

	obj.cptr.stuEndTime = time2ntex(t)
}

func (obj *DhObject) OriginalBoundingBox() DhRect {
	obj.init()

	boundingBox := (*C.DH_RECT)(unsafe.Pointer(uintptr(unsafe.Pointer(obj.cptr)) + unsafe.Offsetof(obj.cptr.stuEndTime) + unsafe.Sizeof(C.NET_TIME_EX{})))
	return DhRect{cptr: boundingBox}
}

func (obj *DhObject) SetOriginalBoundingBox(rect DhRect) {
	obj.init()
	box := (*C.DH_RECT)(unsafe.Pointer(uintptr(unsafe.Pointer(obj.cptr)) + unsafe.Offsetof(obj.cptr.stuEndTime) + unsafe.Sizeof(C.NET_TIME_EX{})))
	*box = *rect.cptr
}

// DH_RECT             stuSignBoundingBox;

func (obj *DhObject) CurrentSequence() int {
	obj.init()

	return int(obj.cptr.dwCurrentSequence)
}

func (obj *DhObject) SetCurrentSequence(n int) {
	obj.init()

	obj.cptr.dwCurrentSequence = C.DWORD(n)
}

func (obj *DhObject) BeginSequence() int {
	obj.init()

	return int(obj.cptr.dwBeginSequence)
}

func (obj *DhObject) SetBeginSequence(n int) {
	obj.init()

	obj.cptr.dwBeginSequence = C.DWORD(n)
}

func (obj *DhObject) EndSequence() int {
	obj.init()

	return int(obj.cptr.dwEndSequence)
}

func (obj *DhObject) SetEndSequence(n int) {
	obj.init()

	obj.cptr.dwEndSequence = C.DWORD(n)
}

func (obj *DhObject) BeginFileOffset() int64 {
	obj.init()

	return int64(obj.cptr.nBeginFileOffset)
}

func (obj *DhObject) SetBeginFileOffset(n int64) {
	obj.init()

	obj.cptr.nBeginFileOffset = C.INT64(n)
}

func (obj *DhObject) EndFileOffset() int64 {
	obj.init()

	return int64(obj.cptr.nEndFileOffset)
}

func (obj *DhObject) SetEndFileOffset(n int64) {
	obj.init()

	obj.cptr.nEndFileOffset = C.INT64(n)
}

func (obj *DhObject) ColorSimilar() []byte {
	obj.init()

	return C.GoBytes(unsafe.Pointer(&obj.cptr.byColorSimilar[0]), C.NET_COLOR_TYPE_MAX)
}

func (obj *DhObject) SetColorSimilar(similar []byte) {
	obj.init()
	var n C.ulong
	if C.ulong(len(similar)) < C.NET_COLOR_TYPE_MAX {
		n = C.ulong(len(similar))
	} else {
		n = C.NET_COLOR_TYPE_MAX
	}

	C.strncpy((*C.char)(unsafe.Pointer(&obj.cptr.byColorSimilar)), (*C.char)(unsafe.Pointer(&similar[0])), n)
}

func (obj *DhObject) UpperBodyColorSimilar() []byte {
	obj.init()

	return C.GoBytes(unsafe.Pointer(&obj.cptr.byUpperBodyColorSimilar[0]), C.NET_COLOR_TYPE_MAX)
}

func (obj *DhObject) SetUpperBodyColorSimilar(similar []byte) {
	obj.init()
	var n C.ulong
	if C.ulong(len(similar)) < C.NET_COLOR_TYPE_MAX {
		n = C.ulong(len(similar))
	} else {
		n = C.NET_COLOR_TYPE_MAX
	}

	C.strncpy((*C.char)(unsafe.Pointer(&obj.cptr.byUpperBodyColorSimilar)), (*C.char)(unsafe.Pointer(&similar[0])), n)
}

func (obj *DhObject) LowerBodyColorSimilar() []byte {
	obj.init()

	return C.GoBytes(unsafe.Pointer(&obj.cptr.byLowerBodyColorSimilar[0]), C.NET_COLOR_TYPE_MAX)
}

func (obj *DhObject) SetLowerBodyColorSimilar(similar []byte) {
	obj.init()
	var n C.ulong
	if C.ulong(len(similar)) < C.NET_COLOR_TYPE_MAX {
		n = C.ulong(len(similar))
	} else {
		n = C.NET_COLOR_TYPE_MAX
	}

	C.strncpy((*C.char)(unsafe.Pointer(&obj.cptr.byLowerBodyColorSimilar)), (*C.char)(unsafe.Pointer(&similar[0])), n)
}

func (obj *DhObject) RelativeID() int {
	obj.init()

	return int(obj.cptr.nRelativeID)
}

func (obj *DhObject) SetRelativeID(n int) {
	obj.init()

	obj.cptr.nRelativeID = C.int(n)
}

func (obj *DhObject) SubText() string {
	obj.init()

	return GoStr((*C.char)(&obj.cptr.szSubText[0]), 20)
}

func (obj *DhObject) SetSubText(txt string) {
	obj.init()
	var n C.ulong
	if len(txt) < 20 {
		n = C.ulong(len(txt))
	} else {
		n = 20
	}

	C.strncpy((*C.char)(unsafe.Pointer(&obj.cptr.szSubText)), C.CString(txt), n)
}

func (obj *DhObject) BrandYear() int {
	obj.init()

	return int(obj.cptr.wBrandYear)
}

func (obj *DhObject) SetBrandYear(n int) {
	obj.init()

	obj.cptr.wBrandYear = C.WORD(n)
}

func (obj DhObject) String() string {
	return FormatString(&obj)
}

type DhRect struct {
	cptr *C.DH_RECT
}

func (rect *DhRect) init() {
	if rect.cptr != nil {
		return
	}

	rect.cptr = &C.DH_RECT{}
}

func (rect *DhRect) Left() int {
	rect.init()

	return int(rect.cptr.left)
}

func (rect *DhRect) SetLeft(l int) {
	rect.init()

	rect.cptr.left = C.long(l)
}

func (rect *DhRect) Top() int {
	rect.init()

	return int(rect.cptr.top)
}

func (rect *DhRect) SetTop(t int) {
	rect.init()

	rect.cptr.top = C.long(t)
}

func (rect *DhRect) Right() int {
	rect.init()

	return int(rect.cptr.right)
}

func (rect *DhRect) SetRight(r int) {
	rect.init()

	rect.cptr.right = C.long(r)
}

func (rect *DhRect) Bottom() int {
	rect.init()

	return int(rect.cptr.bottom)
}

func (rect *DhRect) SetBottom(b int) {
	rect.init()

	rect.cptr.bottom = C.long(b)
}

func (rect DhRect) String() string {
	return FormatString(&rect)
}

type DhPoint struct {
	cptr *C.DH_POINT
}

func (point *DhPoint) init() {
	if point.cptr != nil {
		return
	}

	point.cptr = &C.DH_POINT{}
}

func (point *DhPoint) X() int {
	point.init()

	return int(point.cptr.nx)
}

func (point *DhPoint) SetX(x int) {
	point.init()

	point.cptr.nx = C.short(x)
}

func (point *DhPoint) Y() int {
	point.init()

	return int(point.cptr.ny)
}

func (point *DhPoint) SetY(y int) {
	point.init()

	point.cptr.ny = C.short(y)
}

func (point DhPoint) String() string {
	return FormatString(&point)
}

type DhPicInfo struct {
	cptr *C.DH_PIC_INFO
}

func (pic *DhPicInfo) init() {
	if pic.cptr != nil {
		return
	}

	pic.cptr = &C.DH_PIC_INFO{}
}

func (pic *DhPicInfo) OffSet() int {
	pic.init()

	return int(pic.cptr.dwOffSet)
}

func (pic *DhPicInfo) SetOffSet(oft int) {
	pic.init()

	pic.cptr.dwOffSet = C.DWORD(oft)
}

func (pic *DhPicInfo) FileLenth() int {
	pic.init()

	return int(pic.cptr.dwFileLenth)
}

func (pic *DhPicInfo) SetFileLenth(length int) {
	pic.init()

	pic.cptr.dwFileLenth = C.DWORD(length)
}

func (pic *DhPicInfo) Width() int {
	pic.init()

	return int(pic.cptr.wWidth)
}

func (pic *DhPicInfo) SetWidth(wid int) {
	pic.init()

	pic.cptr.wWidth = C.WORD(wid)
}

func (pic *DhPicInfo) Height() int {
	pic.init()

	return int(pic.cptr.wHeight)
}

func (pic *DhPicInfo) SetHeight(high int) {
	pic.init()

	pic.cptr.wHeight = C.WORD(high)
}

func (pic *DhPicInfo) FilePath() string {
	pic.init()

	return GoStr((*C.char)(pic.cptr.pszFilePath), pic.cptr.nFilePathLen)
}

func (pic *DhPicInfo) SetFilePath(txt string) {
	pic.init()

	C.strncpy((*C.char)(unsafe.Pointer(pic.cptr.pszFilePath)), C.CString(txt), C.ulong(len(txt)))
	pic.cptr.nFilePathLen = C.int(len(txt))
}

func (pic *DhPicInfo) IsDetected() bool {
	pic.init()

	return int(pic.cptr.bIsDetected) > 0
}

func (pic *DhPicInfo) SetDetected(b bool) {
	pic.init()
	if b {
		pic.cptr.bIsDetected = 1
	} else {
		pic.cptr.bIsDetected = 0
	}
}

func (pic *DhPicInfo) Point() DhPoint {
	pic.init()
	return DhPoint{cptr: &pic.cptr.stuPoint}
}

func (pic *DhPicInfo) SetPoint(point DhPoint) {
	pic.init()

	if point.cptr != nil {
		pic.cptr.stuPoint = *point.cptr
	}
}

func (pic DhPicInfo) String() string {
	return FormatString(&pic)
}
