package dhnetsdk

// #cgo CFLAGS: -Iinclude/
// #cgo amd64 386 CFLAGS: -DX86=1
// #cgo LDFLAGS: -Llib/ -ldhnetsdk -lpthread -Wl,-rpath ./lib/
// #include <stdio.h>
// #include <stdlib.h>
// #include <string.h>
// #include <stdbool.h>
// #include "dhnetsdk.h"
//
// extern void goDisconnect(LDWORD dwUser, char *pchDVRIP, LONG nDVRPort);
// extern void goReConnect(LDWORD dwUser, char *pchDVRIP, LONG nDVRPort);
// extern bool goDvrMessage(LDWORD dwUser,LONG lCommand,char *pBuf, DWORD dwBufLen,char *pchDVRIP, LONG nDVRPort);
//
// void CALLBACK cDisConnectFunc(LLONG lLoginID, char *pchDVRIP, LONG nDVRPort, LDWORD dwUser) {
//     if (0 != dwUser)
//     {
//         goDisconnect(dwUser, pchDVRIP, nDVRPort);
//     }
// }
//
// void CALLBACK cReConnectFunc(LLONG lLoginID, char *pchDVRIP, LONG nDVRPort, LDWORD dwUser)
// {
//     if (0 != dwUser)
//     {
//         goReConnect(dwUser, pchDVRIP, nDVRPort);
//     }
// }
//
// BOOL CALLBACK cMessCallBack(LONG lCommand, LLONG lLoginID, char *pBuf, DWORD dwBufLen, char *pchDVRIP, LONG nDVRPort, LDWORD dwUser)
// {
//     if(0 == dwUser)
//     {
//         return false;
//     }
// 	   return goDvrMessage(dwUser, lCommand, pBuf, dwBufLen, pchDVRIP, nDVRPort);
// }
import "C"

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/mattn/go-pointer"
	"github.com/yudai/pp"
)

type (
	ReconnectFunc  func(client *Client, ip string, port int)
	DisconnectFunc func(client *Client, ip string, port int)
	DVRMessageFunc func(client *Client, cmd DhAlarmType, buf []byte, ip string, port int) bool
)

type (
	ReconnectVisitor struct {
		Client   *Client
		Callback ReconnectFunc
	}

	DisconnectVisitor struct {
		Client   *Client
		Callback DisconnectFunc
	}

	DrvMessageVisitor struct {
		Client   *Client
		Callback DVRMessageFunc
	}
)

type Client struct {
	LoginID int

	DeviceInfo DeviceInfo
	// SerialNumber []byte
}

func ClientInit(callback DisconnectFunc) (*Client, error) {
	var (
		cli Client
		v   = DisconnectVisitor{
			Client:   &cli,
			Callback: callback,
		}
	)
	p := pointer.Save(v)

	ret := C.CLIENT_Init(C.fDisConnect(C.cDisConnectFunc), (C.long)(uintptr(p)))
	if ret > 0 {
		return &cli, nil
	}
	return nil, errors.New("init client error")
}

func ClientCleanup() error {

	C.CLIENT_Cleanup()
	return nil
}

func (client *Client) SetAutoReconnect(callback ReconnectFunc) error {
	var v = ReconnectVisitor{
		Client:   client,
		Callback: callback,
	}

	p := pointer.Save(v)
	// defer pointer.Unref(p)
	C.CLIENT_SetAutoReconnect(C.fHaveReConnect(C.cReConnectFunc), (C.long)(uintptr(p)))
	return nil
}

func (client *Client) SetDVRMessCallBack(callback DVRMessageFunc) error {
	var v = DrvMessageVisitor{
		Client:   client,
		Callback: callback,
	}

	p := pointer.Save(&v)

	C.CLIENT_SetDVRMessCallBack(C.fMessCallBack(C.cMessCallBack), (C.long)(uintptr(p)))
	return nil
}

func (client *Client) StartListen() bool {
	b := C.CLIENT_StartListenEx(C.LLONG(client.LoginID))
	return b > 0
}

func (client *Client) StopListen() bool {
	b := C.CLIENT_StopListen(C.LLONG(client.LoginID))
	return b > 0
}

// func (client *Client) AlarmReset() bool {

// }

// typedef struct tagNET_IN_LOGIN_WITH_HIGHLEVEL_SECURITY
// {
// 	DWORD dwSize				// 结构体大小
// 	char						szIP[64];			// IP
// 	int							nPort;				// 端口
// 	char						szUserName[64];		// 用户名
// 	char						szPassword[64];		// 密码
// 	EM_LOGIN_SPAC_CAP_TYPE		emSpecCap;			// 登录模式
// 	BYTE						byReserved[4];		// 字节对齐
// 	void*						pCapParam;			// 见 CLIENT_LoginEx 接口 pCapParam 与 nSpecCap 关系
// }NET_IN_LOGIN_WITH_HIGHLEVEL_SECURITY;

// CLIENT_LoginWithHighLevelSecurity 输出参数
// typedef struct tagNET_OUT_LOGIN_WITH_HIGHLEVEL_SECURITY
// {
// 	DWORD						dwSize;				// 结构体大小
// 	NET_DEVICEINFO_Ex			stuDeviceInfo;		// 设备信息
// 	int							nError;				// 错误码，见 CLIENT_Login 接口错误码
// 	BYTE						byReserved[132];	// 预留字段
// }NET_OUT_LOGIN_WITH_HIGHLEVEL_SECURITY;

// typedef struct
// {
//     BYTE                sSerialNumber[DH_SERIALNO_LEN];     // 序列号
//     int                 nAlarmInPortNum;                    // DVR报警输入个数
//     int                 nAlarmOutPortNum;                   // DVR报警输出个数
//     int                 nDiskNum;                           // DVR硬盘个数
//     int                 nDVRType;                           // DVR类型,见枚举 NET_DEVICE_TYPE
//     int                 nChanNum;                           // DVR通道个数
//     BYTE                byLimitLoginTime;                   // 在线超时时间,为0表示不限制登陆,非0表示限制的分钟数
//     BYTE                byLeftLogTimes;                     // 当登陆失败原因为密码错误时,通过此参数通知用户,剩余登陆次数,为0时表示此参数无效
//     BYTE                bReserved[2];                       // 保留字节,字节对齐
//     int                 nLockLeftTime;                      // 当登陆失败,用户解锁剩余时间（秒数）, -1表示设备未设置该参数
//     char                Reserved[24];                       // 保留
// } NET_DEVICEINFO_Ex, *LPNET_DEVICEINFO_Ex;

type DeviceInfo struct {
	cptr *C.NET_DEVICEINFO_Ex
}

type LoginSecurity struct {
	cptr *C.NET_IN_LOGIN_WITH_HIGHLEVEL_SECURITY
}

type LoginSecurityOut struct {
	cptr *C.NET_OUT_LOGIN_WITH_HIGHLEVEL_SECURITY
}

func (login *LoginSecurity) init() {
	if login.cptr != nil {
		return
	}

	login.cptr = &C.NET_IN_LOGIN_WITH_HIGHLEVEL_SECURITY{}
	login.cptr.dwSize = C.uint(unsafe.Sizeof(*login.cptr))
	// login.cptr.szIP = C.strncpy(unsafe.Pointer(login.ctrp.szIP), C.CString(login.IP), len(login.IP))
}

func (login *LoginSecurity) IP() string {
	login.init()
	return C.GoStringN(&login.cptr.szIP[0], 64)
}

func (login *LoginSecurity) SetIP(ip string) {
	login.init()
	C.strncpy((*C.char)(unsafe.Pointer(&login.cptr.szIP)), C.CString(ip), C.ulong(len(ip)))
}

func (login *LoginSecurity) Port() int {
	login.init()
	return int(login.cptr.nPort)
}

func (login *LoginSecurity) SetPort(port int) {
	login.init()
	login.cptr.nPort = C.int(port)
}

func (login *LoginSecurity) UserName() string {
	login.init()
	return C.GoStringN(&login.cptr.szUserName[0], 64)
}

func (login *LoginSecurity) SetUserName(username string) {
	login.init()
	C.strncpy((*C.char)(unsafe.Pointer(&login.cptr.szUserName)), C.CString(username), C.ulong(len(username)))
}

func (login *LoginSecurity) Password() string {
	login.init()
	return C.GoStringN(&login.cptr.szPassword[0], 64)
}

func (login *LoginSecurity) SetPassword(password string) {
	login.init()
	C.strncpy((*C.char)(unsafe.Pointer(&login.cptr.szPassword)), C.CString(password), C.ulong(len(password)))
}

func (login *LoginSecurity) SpecCap() EmLoginSpacCapType {
	login.init()
	return EmLoginSpacCapType(login.cptr.emSpecCap)
}

func (login *LoginSecurity) SetSpecCap(specCap EmLoginSpacCapType) {
	login.init()

	login.cptr.emSpecCap = C.EM_LOGIN_SPAC_CAP_TYPE(C.int(specCap))
}

func (login *LoginSecurity) Print() {
	pp.Print(login.cptr)
}

func (logout *LoginSecurityOut) init() {
	if logout.cptr != nil {
		return
	}

	logout.cptr = &C.NET_OUT_LOGIN_WITH_HIGHLEVEL_SECURITY{}
	logout.cptr.dwSize = C.uint(unsafe.Sizeof(*logout.cptr))
}

func (logout *LoginSecurityOut) ErrorCode() int {
	logout.init()
	return int(logout.cptr.nError)
}

func (logout *LoginSecurityOut) SetErrorCode(err int) {
	logout.init()
	logout.cptr.nError = C.int(err)
}

func (logout *LoginSecurityOut) DeviceInfo() *DeviceInfo {
	var dev = DeviceInfo{cptr: &logout.cptr.stuDeviceInfo}
	log.Printf("%#v", logout.cptr.stuDeviceInfo)
	return &dev
}

func (di *DeviceInfo) init() {
	if di.cptr != nil {
		return
	}

	di.cptr = &C.NET_DEVICEINFO_Ex{}
	// di.cptr.dwSize = C.uint(unsafe.Sizeof(di.cptr))
}

//     BYTE                sSerialNumber[DH_SERIALNO_LEN];     // 序列号
//     int                 nAlarmInPortNum;                    // DVR报警输入个数
//     int                 nAlarmOutPortNum;                   // DVR报警输出个数
//     int                 nDiskNum;                           // DVR硬盘个数
//     int                 nDVRType;                           // DVR类型,见枚举 NET_DEVICE_TYPE
//     int                 nChanNum;                           // DVR通道个数
//     BYTE                byLimitLoginTime;                   // 在线超时时间,为0表示不限制登陆,非0表示限制的分钟数
//     BYTE                byLeftLogTimes;                     // 当登陆失败原因为密码错误时,通过此参数通知用户,剩余登陆次数,为0时表示此参数无效
//     BYTE                bReserved[2];                       // 保留字节,字节对齐
//     int                 nLockLeftTime;                      // 当登陆失败,用户解锁剩余时间（秒数）, -1表示设备未设置该参数
//     char                Reserved[24];                       // 保留

const DhSerialnoLen = 48 // 设备序列号字符长度

func (di *DeviceInfo) SerialNumber() string {
	di.init()

	return C.GoStringN((*C.char)(unsafe.Pointer(&di.cptr.sSerialNumber[0])), DhSerialnoLen)
}

func (di *DeviceInfo) AlarmInPort() int {
	di.init()

	return int(di.cptr.nAlarmInPortNum)
}

func (di *DeviceInfo) AlarmOutPort() int {
	di.init()

	return int(di.cptr.nAlarmOutPortNum)
}

func (di *DeviceInfo) DiskNum() int {
	di.init()

	return int(di.cptr.nDiskNum)
}

func (di *DeviceInfo) DVRType() int {
	di.init()

	return int(di.cptr.nDVRType)
}

func (di *DeviceInfo) ChanNum() int {
	di.init()

	return int(di.cptr.nChanNum)
}

func (di *DeviceInfo) LimitLoginTime() time.Duration {
	di.init()

	if di.cptr.byLimitLoginTime == 0 {
		return time.Duration(0xffffffff)
	}

	return time.Duration(di.cptr.byLimitLoginTime) * time.Minute
}

func (di *DeviceInfo) LeftLogTimes() int {
	di.init()
	return int(di.cptr.byLeftLogTimes)
}

func (di *DeviceInfo) LockLeftTime() time.Duration {
	return time.Duration(di.cptr.nLockLeftTime) * time.Second
}

func (di DeviceInfo) String() string {
	return FormatString(&di)
}

func LoginWithLevelSecurity(security *LoginSecurity, out *LoginSecurityOut) (int, error) {
	out.init()

	m_lLoginId := C.CLIENT_LoginWithHighLevelSecurity(security.cptr, out.cptr)
	if m_lLoginId == 0 {
		return 0, ErrLoginFailed
	}
	return int(m_lLoginId), nil
}

type LoginOptFunc func(*LoginSecurity)

func (cli *Client) Login(addr string, user, pass string, opts ...LoginOptFunc) (err error) {
	var (
		security LoginSecurity
		port     int
		out      LoginSecurityOut
	)

	for _, opt := range opts {
		opt(&security)
	}

	addrs := strings.SplitN(addr, ":", 2)
	if len(addrs) == 2 {
		addr = addrs[0]
		if port, err = strconv.Atoi(addrs[1]); err != nil {
			return err
		}
	} else {
		return ErrInvalidAddress
	}
	security.SetIP(addr)
	security.SetPort(port)
	security.SetUserName(user)
	security.SetPassword(pass)

	id, err := LoginWithLevelSecurity(&security, &out)
	if err != nil {
		return err
	}
	cli.LoginID = id
	cli.DeviceInfo = *out.DeviceInfo()
	return nil
}
