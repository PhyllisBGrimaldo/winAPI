// +build windows

package win

import (
	"syscall"
	"unsafe"
)

var (
	// Library
	libws2_32 uintptr

	// Functions
	freeAddrInfoExW                 uintptr
	freeAddrInfoW                   uintptr
	getAddrInfoEx                   uintptr
	getAddrInfoW                    uintptr
	getNameInfoW                    uintptr
	inetNtopW                       uintptr
	inetPtonW                       uintptr
	wPUCompleteOverlappedRequest    uintptr
	wSAAccept                       uintptr
	wSAAddressToString              uintptr
	wSAAsyncGetHostByAddr           uintptr
	wSAAsyncGetHostByName           uintptr
	wSAAsyncGetProtoByName          uintptr
	wSAAsyncGetProtoByNumber        uintptr
	wSAAsyncGetServByName           uintptr
	wSAAsyncGetServByPort           uintptr
	wSAAsyncSelect                  uintptr
	wSACancelAsyncRequest           uintptr
	wSACancelBlockingCall           uintptr
	wSACleanup                      uintptr
	wSACloseEvent                   uintptr
	wSAConnect                      uintptr
	wSACreateEvent                  uintptr
	wSADuplicateSocket              uintptr
	wSAEnumNameSpaceProviders       uintptr
	wSAEnumNetworkEvents            uintptr
	wSAEnumProtocols                uintptr
	wSAEventSelect                  uintptr
	wSAGetLastError                 uintptr
	wSAGetOverlappedResult          uintptr
	wSAGetQOSByName                 uintptr
	wSAGetServiceClassInfo          uintptr
	wSAGetServiceClassNameByClassId uintptr
	wSAHtonl                        uintptr
	wSAHtons                        uintptr
	wSAInstallServiceClass          uintptr
	wSAIoctl                        uintptr
	wSAIsBlocking                   uintptr
	wSAJoinLeaf                     uintptr
	wSALookupServiceBegin           uintptr
	wSALookupServiceEnd             uintptr
	wSALookupServiceNext            uintptr
	wSANSPIoctl                     uintptr
	wSANtohl                        uintptr
	wSANtohs                        uintptr
	wSAPoll                         uintptr
	wSAProviderConfigChange         uintptr
	wSARecv                         uintptr
	wSARecvDisconnect               uintptr
	wSARecvFrom                     uintptr
	wSARemoveServiceClass           uintptr
	wSASend                         uintptr
	wSASendDisconnect               uintptr
	wSASendMsg                      uintptr
	wSASendTo                       uintptr
	wSASetBlockingHook              uintptr
	wSASetLastError                 uintptr
	wSASetService                   uintptr
	wSASocket                       uintptr
	wSAStartup                      uintptr
	wSAStringToAddress              uintptr
	wSAUnhookBlockingHook           uintptr
	wSApSetPostRoutine              uintptr
	wSCDeinstallProvider            uintptr
	wSCEnableNSProvider             uintptr
	wSCEnumProtocols                uintptr
	wSCGetProviderPath              uintptr
	wSCInstallNameSpace             uintptr
	wSCInstallProvider              uintptr
	wSCUnInstallNameSpace           uintptr
	wSCWriteProviderOrder           uintptr
	gethostname                     uintptr
)

func init() {
	// Library
	libws2_32 = doLoadLibrary("ws2_32.dll")

	// Functions
	freeAddrInfoExW = doGetProcAddress(libws2_32, "FreeAddrInfoExW")
	freeAddrInfoW = doGetProcAddress(libws2_32, "FreeAddrInfoW")
	getAddrInfoEx = doGetProcAddress(libws2_32, "GetAddrInfoExW")
	getAddrInfoW = doGetProcAddress(libws2_32, "GetAddrInfoW")
	getNameInfoW = doGetProcAddress(libws2_32, "GetNameInfoW")
	inetNtopW = doGetProcAddress(libws2_32, "InetNtopW")
	inetPtonW = doGetProcAddress(libws2_32, "InetPtonW")
	wPUCompleteOverlappedRequest = doGetProcAddress(libws2_32, "WPUCompleteOverlappedRequest")
	wSAAccept = doGetProcAddress(libws2_32, "WSAAccept")
	wSAAddressToString = doGetProcAddress(libws2_32, "WSAAddressToStringW")
	wSAAsyncGetHostByAddr = doGetProcAddress(libws2_32, "WSAAsyncGetHostByAddr")
	wSAAsyncGetHostByName = doGetProcAddress(libws2_32, "WSAAsyncGetHostByName")
	wSAAsyncGetProtoByName = doGetProcAddress(libws2_32, "WSAAsyncGetProtoByName")
	wSAAsyncGetProtoByNumber = doGetProcAddress(libws2_32, "WSAAsyncGetProtoByNumber")
	wSAAsyncGetServByName = doGetProcAddress(libws2_32, "WSAAsyncGetServByName")
	wSAAsyncGetServByPort = doGetProcAddress(libws2_32, "WSAAsyncGetServByPort")
	wSAAsyncSelect = doGetProcAddress(libws2_32, "WSAAsyncSelect")
	wSACancelAsyncRequest = doGetProcAddress(libws2_32, "WSACancelAsyncRequest")
	wSACancelBlockingCall = doGetProcAddress(libws2_32, "WSACancelBlockingCall")
	wSACleanup = doGetProcAddress(libws2_32, "WSACleanup")
	wSACloseEvent = doGetProcAddress(libws2_32, "WSACloseEvent")
	wSAConnect = doGetProcAddress(libws2_32, "WSAConnect")
	wSACreateEvent = doGetProcAddress(libws2_32, "WSACreateEvent")
	wSADuplicateSocket = doGetProcAddress(libws2_32, "WSADuplicateSocketW")
	wSAEnumNameSpaceProviders = doGetProcAddress(libws2_32, "WSAEnumNameSpaceProvidersW")
	wSAEnumNetworkEvents = doGetProcAddress(libws2_32, "WSAEnumNetworkEvents")
	wSAEnumProtocols = doGetProcAddress(libws2_32, "WSAEnumProtocolsW")
	wSAEventSelect = doGetProcAddress(libws2_32, "WSAEventSelect")
	wSAGetLastError = doGetProcAddress(libws2_32, "WSAGetLastError")
	wSAGetOverlappedResult = doGetProcAddress(libws2_32, "WSAGetOverlappedResult")
	wSAGetQOSByName = doGetProcAddress(libws2_32, "WSAGetQOSByName")
	wSAGetServiceClassInfo = doGetProcAddress(libws2_32, "WSAGetServiceClassInfoW")
	wSAGetServiceClassNameByClassId = doGetProcAddress(libws2_32, "WSAGetServiceClassNameByClassIdW")
	wSAHtonl = doGetProcAddress(libws2_32, "WSAHtonl")
	wSAHtons = doGetProcAddress(libws2_32, "WSAHtons")
	wSAInstallServiceClass = doGetProcAddress(libws2_32, "WSAInstallServiceClassW")
	wSAIoctl = doGetProcAddress(libws2_32, "WSAIoctl")
	wSAIsBlocking = doGetProcAddress(libws2_32, "WSAIsBlocking")
	wSAJoinLeaf = doGetProcAddress(libws2_32, "WSAJoinLeaf")
	wSALookupServiceBegin = doGetProcAddress(libws2_32, "WSALookupServiceBeginW")
	wSALookupServiceEnd = doGetProcAddress(libws2_32, "WSALookupServiceEnd")
	wSALookupServiceNext = doGetProcAddress(libws2_32, "WSALookupServiceNextW")
	wSANSPIoctl = doGetProcAddress(libws2_32, "WSANSPIoctl")
	wSANtohl = doGetProcAddress(libws2_32, "WSANtohl")
	wSANtohs = doGetProcAddress(libws2_32, "WSANtohs")
	wSAPoll = doGetProcAddress(libws2_32, "WSAPoll")
	wSAProviderConfigChange = doGetProcAddress(libws2_32, "WSAProviderConfigChange")
	wSARecv = doGetProcAddress(libws2_32, "WSARecv")
	wSARecvDisconnect = doGetProcAddress(libws2_32, "WSARecvDisconnect")
	wSARecvFrom = doGetProcAddress(libws2_32, "WSARecvFrom")
	wSARemoveServiceClass = doGetProcAddress(libws2_32, "WSARemoveServiceClass")
	wSASend = doGetProcAddress(libws2_32, "WSASend")
	wSASendDisconnect = doGetProcAddress(libws2_32, "WSASendDisconnect")
	wSASendMsg = doGetProcAddress(libws2_32, "WSASendMsg")
	wSASendTo = doGetProcAddress(libws2_32, "WSASendTo")
	wSASetBlockingHook = doGetProcAddress(libws2_32, "WSASetBlockingHook")
	wSASetLastError = doGetProcAddress(libws2_32, "WSASetLastError")
	wSASetService = doGetProcAddress(libws2_32, "WSASetServiceW")
	wSASocket = doGetProcAddress(libws2_32, "WSASocketW")
	wSAStartup = doGetProcAddress(libws2_32, "WSAStartup")
	wSAStringToAddress = doGetProcAddress(libws2_32, "WSAStringToAddressW")
	wSAUnhookBlockingHook = doGetProcAddress(libws2_32, "WSAUnhookBlockingHook")
	wSApSetPostRoutine = doGetProcAddress(libws2_32, "WSApSetPostRoutine")
	wSCDeinstallProvider = doGetProcAddress(libws2_32, "WSCDeinstallProvider")
	wSCEnableNSProvider = doGetProcAddress(libws2_32, "WSCEnableNSProvider")
	wSCEnumProtocols = doGetProcAddress(libws2_32, "WSCEnumProtocols")
	wSCGetProviderPath = doGetProcAddress(libws2_32, "WSCGetProviderPath")
	wSCInstallNameSpace = doGetProcAddress(libws2_32, "WSCInstallNameSpace")
	wSCInstallProvider = doGetProcAddress(libws2_32, "WSCInstallProvider")
	wSCUnInstallNameSpace = doGetProcAddress(libws2_32, "WSCUnInstallNameSpace")
	wSCWriteProviderOrder = doGetProcAddress(libws2_32, "WSCWriteProviderOrder")
	gethostname = doGetProcAddress(libws2_32, "gethostname")
}

func FreeAddrInfoExW(ai *ADDRINFOEX) {
	syscall3(freeAddrInfoExW, 1,
		uintptr(unsafe.Pointer(ai)),
		0,
		0)
}

func FreeAddrInfoW(ai PADDRINFO) {
	syscall3(freeAddrInfoW, 1,
		uintptr(unsafe.Pointer(ai)),
		0,
		0)
}

func GetAddrInfoEx(name /*const*/ *WCHAR, servname /*const*/ *WCHAR, namespace uint32, namespace_id *GUID, hints /*const*/ *ADDRINFOEX, result **ADDRINFOEX, timeout *Timeval, overlapped *OVERLAPPED, completion_routine LPLOOKUPSERVICE_COMPLETION_ROUTINE, handle *HANDLE) int32 {
	completion_routineCallback := syscall.NewCallback(completion_routine)
	ret1 := syscall12(getAddrInfoEx, 10,
		uintptr(unsafe.Pointer(name)),
		uintptr(unsafe.Pointer(servname)),
		uintptr(namespace),
		uintptr(unsafe.Pointer(namespace_id)),
		uintptr(unsafe.Pointer(hints)),
		uintptr(unsafe.Pointer(result)),
		uintptr(unsafe.Pointer(timeout)),
		uintptr(unsafe.Pointer(overlapped)),
		completion_routineCallback,
		uintptr(unsafe.Pointer(handle)),
		0,
		0)
	return int32(ret1)
}

func GetAddrInfoW(nodename string, servname string, hints /*const*/ *ADDRINFO, res *PADDRINFO) int32 {
	nodenameStr := unicode16FromString(nodename)
	servnameStr := unicode16FromString(servname)
	ret1 := syscall6(getAddrInfoW, 4,
		uintptr(unsafe.Pointer(&nodenameStr[0])),
		uintptr(unsafe.Pointer(&servnameStr[0])),
		uintptr(unsafe.Pointer(hints)),
		uintptr(unsafe.Pointer(res)),
		0,
		0)
	return int32(ret1)
}

func GetNameInfoW(sa /*const*/ *SOCKADDR, salen Socklen_t, host PWCHAR, hostlen uint32, serv PWCHAR, servlen uint32, flags int32) int32 {
	ret1 := syscall9(getNameInfoW, 7,
		uintptr(unsafe.Pointer(sa)),
		uintptr(salen),
		uintptr(unsafe.Pointer(host)),
		uintptr(hostlen),
		uintptr(unsafe.Pointer(serv)),
		uintptr(servlen),
		uintptr(flags),
		0,
		0)
	return int32(ret1)
}

func InetNtopW(family int32, addr uintptr, buffer PWSTR, aLen SIZE_T) string {
	ret1 := syscall6(inetNtopW, 4,
		uintptr(family),
		addr,
		uintptr(unsafe.Pointer(buffer)),
		uintptr(aLen),
		0,
		0)
	return stringFromUnicode16((*uint16)(unsafe.Pointer(ret1)))
}

func InetPtonW(family int32, addr string, buffer uintptr) int32 {
	addrStr := unicode16FromString(addr)
	ret1 := syscall3(inetPtonW, 3,
		uintptr(family),
		uintptr(unsafe.Pointer(&addrStr[0])),
		buffer)
	return int32(ret1)
}

func WPUCompleteOverlappedRequest(s SOCKET, overlapped LPWSAOVERLAPPED, error uint32, transferred uint32, errcode *int32) WSAEVENT {
	ret1 := syscall6(wPUCompleteOverlappedRequest, 5,
		uintptr(s),
		uintptr(unsafe.Pointer(overlapped)),
		uintptr(error),
		uintptr(transferred),
		uintptr(unsafe.Pointer(errcode)),
		0)
	return WSAEVENT(ret1)
}

func WSAAccept(s SOCKET, addr *Sockaddr, addrlen *int32, lpfnCondition LPCONDITIONPROC, dwCallbackData *uint32) SOCKET {
	lpfnConditionCallback := syscall.NewCallback(func(lpCallerIdRawArg LPWSABUF, lpCallerDataRawArg LPWSABUF, lpSQOSRawArg LPQOS, lpGQOSRawArg LPQOS, lpCalleeIdRawArg LPWSABUF, lpCalleeDataRawArg LPWSABUF, gRawArg *GROUP, dwCallbackDataRawArg DWORD_PTR) uintptr {
		ret := lpfnCondition(lpCallerIdRawArg, lpCallerDataRawArg, lpSQOSRawArg, lpGQOSRawArg, lpCalleeIdRawArg, lpCalleeDataRawArg, gRawArg, dwCallbackDataRawArg)
		return uintptr(ret)
	})
	ret1 := syscall6(wSAAccept, 5,
		uintptr(s),
		uintptr(unsafe.Pointer(addr)),
		uintptr(unsafe.Pointer(addrlen)),
		lpfnConditionCallback,
		uintptr(unsafe.Pointer(dwCallbackData)),
		0)
	return SOCKET(ret1)
}

func WSAAddressToString(sockaddr *SOCKADDR, aLen uint32, info LPWSAPROTOCOL_INFO, string LPWSTR, lenstr *uint32) int32 {
	ret1 := syscall6(wSAAddressToString, 5,
		uintptr(unsafe.Pointer(sockaddr)),
		uintptr(aLen),
		uintptr(unsafe.Pointer(info)),
		uintptr(unsafe.Pointer(string)),
		uintptr(unsafe.Pointer(lenstr)),
		0)
	return int32(ret1)
}

func WSAAsyncGetHostByAddr(hWnd HWND, uMsg uint32, addr /*const*/ LPCSTR, aLen int32, aType int32, sbuf LPSTR, buflen int32) HANDLE {
	ret1 := syscall9(wSAAsyncGetHostByAddr, 7,
		uintptr(hWnd),
		uintptr(uMsg),
		uintptr(unsafe.Pointer(addr)),
		uintptr(aLen),
		uintptr(aType),
		uintptr(unsafe.Pointer(sbuf)),
		uintptr(buflen),
		0,
		0)
	return HANDLE(ret1)
}

func WSAAsyncGetHostByName(hWnd HWND, uMsg uint32, name /*const*/ LPCSTR, sbuf LPSTR, buflen int32) HANDLE {
	ret1 := syscall6(wSAAsyncGetHostByName, 5,
		uintptr(hWnd),
		uintptr(uMsg),
		uintptr(unsafe.Pointer(name)),
		uintptr(unsafe.Pointer(sbuf)),
		uintptr(buflen),
		0)
	return HANDLE(ret1)
}

func WSAAsyncGetProtoByName(hWnd HWND, uMsg uint32, name /*const*/ LPCSTR, sbuf LPSTR, buflen int32) HANDLE {
	ret1 := syscall6(wSAAsyncGetProtoByName, 5,
		uintptr(hWnd),
		uintptr(uMsg),
		uintptr(unsafe.Pointer(name)),
		uintptr(unsafe.Pointer(sbuf)),
		uintptr(buflen),
		0)
	return HANDLE(ret1)
}

func WSAAsyncGetProtoByNumber(hWnd HWND, uMsg uint32, number int32, sbuf LPSTR, buflen int32) HANDLE {
	ret1 := syscall6(wSAAsyncGetProtoByNumber, 5,
		uintptr(hWnd),
		uintptr(uMsg),
		uintptr(number),
		uintptr(unsafe.Pointer(sbuf)),
		uintptr(buflen),
		0)
	return HANDLE(ret1)
}

func WSAAsyncGetServByName(hWnd HWND, uMsg uint32, name /*const*/ LPCSTR, proto /*const*/ LPCSTR, sbuf LPSTR, buflen int32) HANDLE {
	ret1 := syscall6(wSAAsyncGetServByName, 6,
		uintptr(hWnd),
		uintptr(uMsg),
		uintptr(unsafe.Pointer(name)),
		uintptr(unsafe.Pointer(proto)),
		uintptr(unsafe.Pointer(sbuf)),
		uintptr(buflen))
	return HANDLE(ret1)
}

func WSAAsyncGetServByPort(hWnd HWND, uMsg uint32, port int32, proto /*const*/ LPCSTR, sbuf LPSTR, buflen int32) HANDLE {
	ret1 := syscall6(wSAAsyncGetServByPort, 6,
		uintptr(hWnd),
		uintptr(uMsg),
		uintptr(port),
		uintptr(unsafe.Pointer(proto)),
		uintptr(unsafe.Pointer(sbuf)),
		uintptr(buflen))
	return HANDLE(ret1)
}

func WSAAsyncSelect(s SOCKET, hWnd HWND, uMsg uint32, lEvent LONG) int32 {
	ret1 := syscall6(wSAAsyncSelect, 4,
		uintptr(s),
		uintptr(hWnd),
		uintptr(uMsg),
		uintptr(lEvent),
		0,
		0)
	return int32(ret1)
}

func WSACancelAsyncRequest(hAsyncTaskHandle HANDLE) int32 {
	ret1 := syscall3(wSACancelAsyncRequest, 1,
		uintptr(hAsyncTaskHandle),
		0,
		0)
	return int32(ret1)
}

func WSACancelBlockingCall() int32 {
	ret1 := syscall3(wSACancelBlockingCall, 0,
		0,
		0,
		0)
	return int32(ret1)
}

func WSACleanup() int32 {
	ret1 := syscall3(wSACleanup, 0,
		0,
		0,
		0)
	return int32(ret1)
}

func WSACloseEvent(event WSAEVENT) bool {
	ret1 := syscall3(wSACloseEvent, 1,
		uintptr(event),
		0,
		0)
	return ret1 != 0
}

func WSAConnect(s SOCKET, name /*const*/ *Sockaddr, namelen int32, lpCallerData LPWSABUF, lpCalleeData LPWSABUF, lpSQOS *QOS, lpGQOS *QOS) int32 {
	ret1 := syscall9(wSAConnect, 7,
		uintptr(s),
		uintptr(unsafe.Pointer(name)),
		uintptr(namelen),
		uintptr(unsafe.Pointer(lpCallerData)),
		uintptr(unsafe.Pointer(lpCalleeData)),
		uintptr(unsafe.Pointer(lpSQOS)),
		uintptr(unsafe.Pointer(lpGQOS)),
		0,
		0)
	return int32(ret1)
}

func WSACreateEvent() WSAEVENT {
	ret1 := syscall3(wSACreateEvent, 0,
		0,
		0,
		0)
	return WSAEVENT(ret1)
}

func WSADuplicateSocket(s SOCKET, dwProcessId uint32, lpProtocolInfo LPWSAPROTOCOL_INFO) int32 {
	ret1 := syscall3(wSADuplicateSocket, 3,
		uintptr(s),
		uintptr(dwProcessId),
		uintptr(unsafe.Pointer(lpProtocolInfo)))
	return int32(ret1)
}

func WSAEnumNameSpaceProviders(aLen *uint32, buffer LPWSANAMESPACE_INFO) int32 {
	ret1 := syscall3(wSAEnumNameSpaceProviders, 2,
		uintptr(unsafe.Pointer(aLen)),
		uintptr(unsafe.Pointer(buffer)),
		0)
	return int32(ret1)
}

func WSAEnumNetworkEvents(s SOCKET, hEvent WSAEVENT, lpEvent *WSANETWORKEVENTS) int32 {
	ret1 := syscall3(wSAEnumNetworkEvents, 3,
		uintptr(s),
		uintptr(hEvent),
		uintptr(unsafe.Pointer(lpEvent)))
	return int32(ret1)
}

func WSAEnumProtocols(protocols *int32, buffer LPWSAPROTOCOL_INFO, aLen *uint32) int32 {
	ret1 := syscall3(wSAEnumProtocols, 3,
		uintptr(unsafe.Pointer(protocols)),
		uintptr(unsafe.Pointer(buffer)),
		uintptr(unsafe.Pointer(aLen)))
	return int32(ret1)
}

func WSAEventSelect(s SOCKET, hEvent WSAEVENT, lEvent LONG) int32 {
	ret1 := syscall3(wSAEventSelect, 3,
		uintptr(s),
		uintptr(hEvent),
		uintptr(lEvent))
	return int32(ret1)
}

func WSAGetLastError() int32 {
	ret1 := syscall3(wSAGetLastError, 0,
		0,
		0,
		0)
	return int32(ret1)
}

func WSAGetOverlappedResult(s SOCKET, lpOverlapped LPWSAOVERLAPPED, lpcbTransfer *uint32, fWait bool, lpdwFlags *uint32) bool {
	ret1 := syscall6(wSAGetOverlappedResult, 5,
		uintptr(s),
		uintptr(unsafe.Pointer(lpOverlapped)),
		uintptr(unsafe.Pointer(lpcbTransfer)),
		getUintptrFromBool(fWait),
		uintptr(unsafe.Pointer(lpdwFlags)),
		0)
	return ret1 != 0
}

func WSAGetQOSByName(s SOCKET, lpQOSName LPWSABUF, lpQOS *QOS) bool {
	ret1 := syscall3(wSAGetQOSByName, 3,
		uintptr(s),
		uintptr(unsafe.Pointer(lpQOSName)),
		uintptr(unsafe.Pointer(lpQOS)))
	return ret1 != 0
}

func WSAGetServiceClassInfo(provider *GUID, service *GUID, aLen *uint32, info LPWSASERVICECLASSINFO) int32 {
	ret1 := syscall6(wSAGetServiceClassInfo, 4,
		uintptr(unsafe.Pointer(provider)),
		uintptr(unsafe.Pointer(service)),
		uintptr(unsafe.Pointer(aLen)),
		uintptr(unsafe.Pointer(info)),
		0,
		0)
	return int32(ret1)
}

func WSAGetServiceClassNameByClassId(class *GUID, service LPWSTR, aLen *uint32) int32 {
	ret1 := syscall3(wSAGetServiceClassNameByClassId, 3,
		uintptr(unsafe.Pointer(class)),
		uintptr(unsafe.Pointer(service)),
		uintptr(unsafe.Pointer(aLen)))
	return int32(ret1)
}

func WSAHtonl(s SOCKET, hostlong ULONG, lpnetlong *ULONG) int32 {
	ret1 := syscall3(wSAHtonl, 3,
		uintptr(s),
		uintptr(hostlong),
		uintptr(unsafe.Pointer(lpnetlong)))
	return int32(ret1)
}

func WSAHtons(s SOCKET, hostshort uint16, lpnetshort *uint16) int32 {
	ret1 := syscall3(wSAHtons, 3,
		uintptr(s),
		uintptr(hostshort),
		uintptr(unsafe.Pointer(lpnetshort)))
	return int32(ret1)
}

func WSAInstallServiceClass(info LPWSASERVICECLASSINFO) int32 {
	ret1 := syscall3(wSAInstallServiceClass, 1,
		uintptr(unsafe.Pointer(info)),
		0,
		0)
	return int32(ret1)
}

func WSAIoctl(s SOCKET, code uint32, in_buff LPVOID, in_size uint32, out_buff LPVOID, out_size uint32, ret_size *uint32, overlapped LPWSAOVERLAPPED, completion LPWSAOVERLAPPED_COMPLETION_ROUTINE) int32 {
	completionCallback := syscall.NewCallback(completion)
	ret1 := syscall9(wSAIoctl, 9,
		uintptr(s),
		uintptr(code),
		uintptr(unsafe.Pointer(in_buff)),
		uintptr(in_size),
		uintptr(unsafe.Pointer(out_buff)),
		uintptr(out_size),
		uintptr(unsafe.Pointer(ret_size)),
		uintptr(unsafe.Pointer(overlapped)),
		completionCallback)
	return int32(ret1)
}

func WSAIsBlocking() bool {
	ret1 := syscall3(wSAIsBlocking, 0,
		0,
		0,
		0)
	return ret1 != 0
}

func WSAJoinLeaf(s SOCKET, addr /*const*/ *Sockaddr, addrlen int32, lpCallerData LPWSABUF, lpCalleeData LPWSABUF, lpSQOS *QOS, lpGQOS *QOS, dwFlags uint32) SOCKET {
	ret1 := syscall9(wSAJoinLeaf, 8,
		uintptr(s),
		uintptr(unsafe.Pointer(addr)),
		uintptr(addrlen),
		uintptr(unsafe.Pointer(lpCallerData)),
		uintptr(unsafe.Pointer(lpCalleeData)),
		uintptr(unsafe.Pointer(lpSQOS)),
		uintptr(unsafe.Pointer(lpGQOS)),
		uintptr(dwFlags),
		0)
	return SOCKET(ret1)
}

func WSALookupServiceBegin(lpqsRestrictions LPWSAQUERYSET, dwControlFlags uint32, lphLookup *HANDLE) int32 {
	ret1 := syscall3(wSALookupServiceBegin, 3,
		uintptr(unsafe.Pointer(lpqsRestrictions)),
		uintptr(dwControlFlags),
		uintptr(unsafe.Pointer(lphLookup)))
	return int32(ret1)
}

func WSALookupServiceEnd(lookup HANDLE) int32 {
	ret1 := syscall3(wSALookupServiceEnd, 1,
		uintptr(lookup),
		0,
		0)
	return int32(ret1)
}

func WSALookupServiceNext(lookup HANDLE, flags uint32, aLen *uint32, results LPWSAQUERYSET) int32 {
	ret1 := syscall6(wSALookupServiceNext, 4,
		uintptr(lookup),
		uintptr(flags),
		uintptr(unsafe.Pointer(aLen)),
		uintptr(unsafe.Pointer(results)),
		0,
		0)
	return int32(ret1)
}

func WSANSPIoctl(hLookup HANDLE, dwControlCode uint32, lpvInBuffer LPVOID, cbInBuffer uint32, lpvOutBuffer LPVOID, cbOutBuffer uint32, lpcbBytesReturned *uint32, lpCompletion *WSACOMPLETION) int32 {
	ret1 := syscall9(wSANSPIoctl, 8,
		uintptr(hLookup),
		uintptr(dwControlCode),
		uintptr(unsafe.Pointer(lpvInBuffer)),
		uintptr(cbInBuffer),
		uintptr(unsafe.Pointer(lpvOutBuffer)),
		uintptr(cbOutBuffer),
		uintptr(unsafe.Pointer(lpcbBytesReturned)),
		uintptr(unsafe.Pointer(lpCompletion)),
		0)
	return int32(ret1)
}

func WSANtohl(s SOCKET, netlong ULONG, lphostlong *ULONG) int32 {
	ret1 := syscall3(wSANtohl, 3,
		uintptr(s),
		uintptr(netlong),
		uintptr(unsafe.Pointer(lphostlong)))
	return int32(ret1)
}

func WSANtohs(s SOCKET, netshort uint16, lphostshort *uint16) int32 {
	ret1 := syscall3(wSANtohs, 3,
		uintptr(s),
		uintptr(netshort),
		uintptr(unsafe.Pointer(lphostshort)))
	return int32(ret1)
}

func WSAPoll(wfds *WSAPOLLFD, count ULONG, timeout int32) int32 {
	ret1 := syscall3(wSAPoll, 3,
		uintptr(unsafe.Pointer(wfds)),
		uintptr(count),
		uintptr(timeout))
	return int32(ret1)
}

func WSAProviderConfigChange(handle *HANDLE, overlapped LPWSAOVERLAPPED, completion LPWSAOVERLAPPED_COMPLETION_ROUTINE) int32 {
	completionCallback := syscall.NewCallback(completion)
	ret1 := syscall3(wSAProviderConfigChange, 3,
		uintptr(unsafe.Pointer(handle)),
		uintptr(unsafe.Pointer(overlapped)),
		completionCallback)
	return int32(ret1)
}

func WSARecv(s SOCKET, lpBuffers LPWSABUF, dwBufferCount uint32, numberOfBytesReceived *uint32, lpFlags *uint32, lpOverlapped LPWSAOVERLAPPED, lpCompletionRoutine LPWSAOVERLAPPED_COMPLETION_ROUTINE) int32 {
	lpCompletionRoutineCallback := syscall.NewCallback(lpCompletionRoutine)
	ret1 := syscall9(wSARecv, 7,
		uintptr(s),
		uintptr(unsafe.Pointer(lpBuffers)),
		uintptr(dwBufferCount),
		uintptr(unsafe.Pointer(numberOfBytesReceived)),
		uintptr(unsafe.Pointer(lpFlags)),
		uintptr(unsafe.Pointer(lpOverlapped)),
		lpCompletionRoutineCallback,
		0,
		0)
	return int32(ret1)
}

func WSARecvDisconnect(s SOCKET, disconnectdata LPWSABUF) int32 {
	ret1 := syscall3(wSARecvDisconnect, 2,
		uintptr(s),
		uintptr(unsafe.Pointer(disconnectdata)),
		0)
	return int32(ret1)
}

func WSARecvFrom(s SOCKET, lpBuffers LPWSABUF, dwBufferCount uint32, lpNumberOfBytesRecvd *uint32, lpFlags *uint32, lpFrom *Sockaddr, lpFromlen *int32, lpOverlapped LPWSAOVERLAPPED, lpCompletionRoutine LPWSAOVERLAPPED_COMPLETION_ROUTINE) int32 {
	lpCompletionRoutineCallback := syscall.NewCallback(lpCompletionRoutine)
	ret1 := syscall9(wSARecvFrom, 9,
		uintptr(s),
		uintptr(unsafe.Pointer(lpBuffers)),
		uintptr(dwBufferCount),
		uintptr(unsafe.Pointer(lpNumberOfBytesRecvd)),
		uintptr(unsafe.Pointer(lpFlags)),
		uintptr(unsafe.Pointer(lpFrom)),
		uintptr(unsafe.Pointer(lpFromlen)),
		uintptr(unsafe.Pointer(lpOverlapped)),
		lpCompletionRoutineCallback)
	return int32(ret1)
}

func WSARemoveServiceClass(info *GUID) int32 {
	ret1 := syscall3(wSARemoveServiceClass, 1,
		uintptr(unsafe.Pointer(info)),
		0,
		0)
	return int32(ret1)
}

func WSASend(s SOCKET, lpBuffers LPWSABUF, dwBufferCount uint32, lpNumberOfBytesSent *uint32, dwFlags uint32, lpOverlapped LPWSAOVERLAPPED, lpCompletionRoutine LPWSAOVERLAPPED_COMPLETION_ROUTINE) int32 {
	lpCompletionRoutineCallback := syscall.NewCallback(lpCompletionRoutine)
	ret1 := syscall9(wSASend, 7,
		uintptr(s),
		uintptr(unsafe.Pointer(lpBuffers)),
		uintptr(dwBufferCount),
		uintptr(unsafe.Pointer(lpNumberOfBytesSent)),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(lpOverlapped)),
		lpCompletionRoutineCallback,
		0,
		0)
	return int32(ret1)
}

func WSASendDisconnect(s SOCKET, lpBuffers LPWSABUF) int32 {
	ret1 := syscall3(wSASendDisconnect, 2,
		uintptr(s),
		uintptr(unsafe.Pointer(lpBuffers)),
		0)
	return int32(ret1)
}

func WSASendMsg(s SOCKET, msg *WSAMSG, dwFlags uint32, lpNumberOfBytesSent *uint32, lpOverlapped LPWSAOVERLAPPED, lpCompletionRoutine LPWSAOVERLAPPED_COMPLETION_ROUTINE) int32 {
	lpCompletionRoutineCallback := syscall.NewCallback(lpCompletionRoutine)
	ret1 := syscall6(wSASendMsg, 6,
		uintptr(s),
		uintptr(unsafe.Pointer(msg)),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(lpNumberOfBytesSent)),
		uintptr(unsafe.Pointer(lpOverlapped)),
		lpCompletionRoutineCallback)
	return int32(ret1)
}

func WSASendTo(s SOCKET, lpBuffers LPWSABUF, dwBufferCount uint32, lpNumberOfBytesSent *uint32, dwFlags uint32, to /*const*/ *Sockaddr, tolen int32, lpOverlapped LPWSAOVERLAPPED, lpCompletionRoutine LPWSAOVERLAPPED_COMPLETION_ROUTINE) int32 {
	lpCompletionRoutineCallback := syscall.NewCallback(lpCompletionRoutine)
	ret1 := syscall9(wSASendTo, 9,
		uintptr(s),
		uintptr(unsafe.Pointer(lpBuffers)),
		uintptr(dwBufferCount),
		uintptr(unsafe.Pointer(lpNumberOfBytesSent)),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(to)),
		uintptr(tolen),
		uintptr(unsafe.Pointer(lpOverlapped)),
		lpCompletionRoutineCallback)
	return int32(ret1)
}

func WSASetBlockingHook(lpBlockFunc FARPROC) FARPROC {
	lpBlockFuncCallback := syscall.NewCallback(func() uintptr {
		ret := lpBlockFunc()
		return uintptr(unsafe.Pointer(ret))
	})
	ret1 := syscall3(wSASetBlockingHook, 1,
		lpBlockFuncCallback,
		0,
		0)
	return func() INT_PTR {
		ret2 := syscall3(ret1, 0,
			0,
			0,
			0)
		return (INT_PTR)(unsafe.Pointer(ret2))
	}
}

func WSASetLastError(iError int32) {
	syscall3(wSASetLastError, 1,
		uintptr(iError),
		0,
		0)
}

func WSASetService(query LPWSAQUERYSET, operation WSAESETSERVICEOP, flags uint32) int32 {
	ret1 := syscall3(wSASetService, 3,
		uintptr(unsafe.Pointer(query)),
		uintptr(operation),
		uintptr(flags))
	return int32(ret1)
}

func WSASocket(af int32, aType int32, protocol int32, lpProtocolInfo LPWSAPROTOCOL_INFO, g GROUP, dwFlags uint32) SOCKET {
	ret1 := syscall6(wSASocket, 6,
		uintptr(af),
		uintptr(aType),
		uintptr(protocol),
		uintptr(unsafe.Pointer(lpProtocolInfo)),
		uintptr(g),
		uintptr(dwFlags))
	return SOCKET(ret1)
}

func WSAStartup(wVersionRequested uint16, lpWSAData *WSADATA) int32 {
	ret1 := syscall3(wSAStartup, 2,
		uintptr(wVersionRequested),
		uintptr(unsafe.Pointer(lpWSAData)),
		0)
	return int32(ret1)
}

func WSAStringToAddress(addressString LPWSTR, addressFamily int32, lpProtocolInfo LPWSAPROTOCOL_INFO, lpAddress *SOCKADDR, lpAddressLength *int32) int32 {
	ret1 := syscall6(wSAStringToAddress, 5,
		uintptr(unsafe.Pointer(addressString)),
		uintptr(addressFamily),
		uintptr(unsafe.Pointer(lpProtocolInfo)),
		uintptr(unsafe.Pointer(lpAddress)),
		uintptr(unsafe.Pointer(lpAddressLength)),
		0)
	return int32(ret1)
}

func WSAUnhookBlockingHook() int32 {
	ret1 := syscall3(wSAUnhookBlockingHook, 0,
		0,
		0,
		0)
	return int32(ret1)
}

func WSApSetPostRoutine(lpPostRoutine LPWPUPOSTMESSAGE) int32 {
	lpPostRoutineCallback := syscall.NewCallback(func(unnamed0RawArg HWND, unnamed1RawArg UINT, unnamed2RawArg WPARAM, unnamed3RawArg LPARAM) uintptr {
		ret := lpPostRoutine(unnamed0RawArg, unnamed1RawArg, unnamed2RawArg, unnamed3RawArg)
		return uintptr(ret)
	})
	ret1 := syscall3(wSApSetPostRoutine, 1,
		lpPostRoutineCallback,
		0,
		0)
	return int32(ret1)
}

func WSCDeinstallProvider(lpProviderId *GUID, lpErrno *int32) int32 {
	ret1 := syscall3(wSCDeinstallProvider, 2,
		uintptr(unsafe.Pointer(lpProviderId)),
		uintptr(unsafe.Pointer(lpErrno)),
		0)
	return int32(ret1)
}

func WSCEnableNSProvider(provider *GUID, enable bool) int32 {
	ret1 := syscall3(wSCEnableNSProvider, 2,
		uintptr(unsafe.Pointer(provider)),
		getUintptrFromBool(enable),
		0)
	return int32(ret1)
}

func WSCEnumProtocols(protocols *int32, buffer LPWSAPROTOCOL_INFO, aLen *uint32, err *int32) int32 {
	ret1 := syscall6(wSCEnumProtocols, 4,
		uintptr(unsafe.Pointer(protocols)),
		uintptr(unsafe.Pointer(buffer)),
		uintptr(unsafe.Pointer(aLen)),
		uintptr(unsafe.Pointer(err)),
		0,
		0)
	return int32(ret1)
}

func WSCGetProviderPath(provider *GUID, path LPWSTR, aLen *int32, errcode *int32) int32 {
	ret1 := syscall6(wSCGetProviderPath, 4,
		uintptr(unsafe.Pointer(provider)),
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(aLen)),
		uintptr(unsafe.Pointer(errcode)),
		0,
		0)
	return int32(ret1)
}

func WSCInstallNameSpace(identifier LPWSTR, path LPWSTR, namespace uint32, version uint32, provider *GUID) int32 {
	ret1 := syscall6(wSCInstallNameSpace, 5,
		uintptr(unsafe.Pointer(identifier)),
		uintptr(unsafe.Pointer(path)),
		uintptr(namespace),
		uintptr(version),
		uintptr(unsafe.Pointer(provider)),
		0)
	return int32(ret1)
}

func WSCInstallProvider(lpProviderId /*const*/ *GUID, lpszProviderDllPath string, lpProtocolInfoList /*const*/ LPWSAPROTOCOL_INFO, dwNumberOfEntries uint32, lpErrno *int32) int32 {
	lpszProviderDllPathStr := unicode16FromString(lpszProviderDllPath)
	ret1 := syscall6(wSCInstallProvider, 5,
		uintptr(unsafe.Pointer(lpProviderId)),
		uintptr(unsafe.Pointer(&lpszProviderDllPathStr[0])),
		uintptr(unsafe.Pointer(lpProtocolInfoList)),
		uintptr(dwNumberOfEntries),
		uintptr(unsafe.Pointer(lpErrno)),
		0)
	return int32(ret1)
}

func WSCUnInstallNameSpace(lpProviderId *GUID) int32 {
	ret1 := syscall3(wSCUnInstallNameSpace, 1,
		uintptr(unsafe.Pointer(lpProviderId)),
		0,
		0)
	return int32(ret1)
}

func WSCWriteProviderOrder(entry *uint32, number uint32) int32 {
	ret1 := syscall3(wSCWriteProviderOrder, 2,
		uintptr(unsafe.Pointer(entry)),
		uintptr(number),
		0)
	return int32(ret1)
}

func Gethostname(name *CHAR, namelen int32) int32 {
	ret1 := syscall3(gethostname, 2,
		uintptr(unsafe.Pointer(name)),
		uintptr(namelen),
		0)
	return int32(ret1)
}