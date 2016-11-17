// +build windows

package win

import (
	"unsafe"
)

var (
	// Library
	libpdh uintptr

	// Functions
	pdhAddCounter                   uintptr
	pdhAddEnglishCounter            uintptr
	pdhBindInputDataSource          uintptr
	pdhCalculateCounterFromRawValue uintptr
	pdhCloseQuery                   uintptr
	pdhCollectQueryData             uintptr
	pdhCollectQueryDataEx           uintptr
	pdhCollectQueryDataWithTime     uintptr
	pdhEnumObjectItems              uintptr
	pdhExpandCounterPath            uintptr
	pdhExpandWildCardPath           uintptr
	pdhGetCounterInfo               uintptr
	pdhGetCounterTimeBase           uintptr
	pdhGetDllVersion                uintptr
	pdhGetFormattedCounterValue     uintptr
	pdhGetLogFileType               uintptr
	pdhGetRawCounterValue           uintptr
	pdhLookupPerfIndexByName        uintptr
	pdhLookupPerfNameByIndex        uintptr
	pdhMakeCounterPath              uintptr
	pdhOpenQuery                    uintptr
	pdhRemoveCounter                uintptr
	pdhSetCounterScaleFactor        uintptr
	pdhSetDefaultRealTimeDataSource uintptr
	pdhValidatePathEx               uintptr
	pdhValidatePath                 uintptr
)

func init() {
	// Library
	libpdh = doLoadLibrary("pdh.dll")

	// Functions
	pdhAddCounter = doGetProcAddress(libpdh, "PdhAddCounterW")
	pdhAddEnglishCounter = doGetProcAddress(libpdh, "PdhAddEnglishCounterW")
	pdhBindInputDataSource = doGetProcAddress(libpdh, "PdhBindInputDataSourceW")
	pdhCalculateCounterFromRawValue = doGetProcAddress(libpdh, "PdhCalculateCounterFromRawValue")
	pdhCloseQuery = doGetProcAddress(libpdh, "PdhCloseQuery")
	pdhCollectQueryData = doGetProcAddress(libpdh, "PdhCollectQueryData")
	pdhCollectQueryDataEx = doGetProcAddress(libpdh, "PdhCollectQueryDataEx")
	pdhCollectQueryDataWithTime = doGetProcAddress(libpdh, "PdhCollectQueryDataWithTime")
	pdhEnumObjectItems = doGetProcAddress(libpdh, "PdhEnumObjectItemsW")
	pdhExpandCounterPath = doGetProcAddress(libpdh, "PdhExpandCounterPathW")
	pdhExpandWildCardPath = doGetProcAddress(libpdh, "PdhExpandWildCardPathW")
	pdhGetCounterInfo = doGetProcAddress(libpdh, "PdhGetCounterInfoW")
	pdhGetCounterTimeBase = doGetProcAddress(libpdh, "PdhGetCounterTimeBase")
	pdhGetDllVersion = doGetProcAddress(libpdh, "PdhGetDllVersion")
	pdhGetFormattedCounterValue = doGetProcAddress(libpdh, "PdhGetFormattedCounterValue")
	pdhGetLogFileType = doGetProcAddress(libpdh, "PdhGetLogFileTypeW")
	pdhGetRawCounterValue = doGetProcAddress(libpdh, "PdhGetRawCounterValue")
	pdhLookupPerfIndexByName = doGetProcAddress(libpdh, "PdhLookupPerfIndexByNameW")
	pdhLookupPerfNameByIndex = doGetProcAddress(libpdh, "PdhLookupPerfNameByIndexW")
	pdhMakeCounterPath = doGetProcAddress(libpdh, "PdhMakeCounterPathW")
	pdhOpenQuery = doGetProcAddress(libpdh, "PdhOpenQueryW")
	pdhRemoveCounter = doGetProcAddress(libpdh, "PdhRemoveCounter")
	pdhSetCounterScaleFactor = doGetProcAddress(libpdh, "PdhSetCounterScaleFactor")
	pdhSetDefaultRealTimeDataSource = doGetProcAddress(libpdh, "PdhSetDefaultRealTimeDataSource")
	pdhValidatePathEx = doGetProcAddress(libpdh, "PdhValidatePathExW")
	pdhValidatePath = doGetProcAddress(libpdh, "PdhValidatePathW")
}

func PdhAddCounter(hquery PDH_HQUERY, path string, userdata *uint32, hcounter *PDH_HCOUNTER) PDH_STATUS {
	pathStr := unicode16FromString(path)
	ret1 := syscall6(pdhAddCounter, 4,
		uintptr(hquery),
		uintptr(unsafe.Pointer(&pathStr[0])),
		uintptr(unsafe.Pointer(userdata)),
		uintptr(unsafe.Pointer(hcounter)),
		0,
		0)
	return PDH_STATUS(ret1)
}

func PdhAddEnglishCounter(query PDH_HQUERY, path string, userdata *uint32, counter *PDH_HCOUNTER) PDH_STATUS {
	pathStr := unicode16FromString(path)
	ret1 := syscall6(pdhAddEnglishCounter, 4,
		uintptr(query),
		uintptr(unsafe.Pointer(&pathStr[0])),
		uintptr(unsafe.Pointer(userdata)),
		uintptr(unsafe.Pointer(counter)),
		0,
		0)
	return PDH_STATUS(ret1)
}

func PdhBindInputDataSource(source *PDH_HLOG, filenamelist /*const*/ *WCHAR) PDH_STATUS {
	ret1 := syscall3(pdhBindInputDataSource, 2,
		uintptr(unsafe.Pointer(source)),
		uintptr(unsafe.Pointer(filenamelist)),
		0)
	return PDH_STATUS(ret1)
}

func PdhCalculateCounterFromRawValue(handle PDH_HCOUNTER, format uint32, raw1 *PDH_RAW_COUNTER, raw2 *PDH_RAW_COUNTER, value *PDH_FMT_COUNTERVALUE) PDH_STATUS {
	ret1 := syscall6(pdhCalculateCounterFromRawValue, 5,
		uintptr(handle),
		uintptr(format),
		uintptr(unsafe.Pointer(raw1)),
		uintptr(unsafe.Pointer(raw2)),
		uintptr(unsafe.Pointer(value)),
		0)
	return PDH_STATUS(ret1)
}

func PdhCloseQuery(handle PDH_HQUERY) PDH_STATUS {
	ret1 := syscall3(pdhCloseQuery, 1,
		uintptr(handle),
		0,
		0)
	return PDH_STATUS(ret1)
}

func PdhCollectQueryData(handle PDH_HQUERY) PDH_STATUS {
	ret1 := syscall3(pdhCollectQueryData, 1,
		uintptr(handle),
		0,
		0)
	return PDH_STATUS(ret1)
}

func PdhCollectQueryDataEx(handle PDH_HQUERY, interval uint32, event HANDLE) PDH_STATUS {
	ret1 := syscall3(pdhCollectQueryDataEx, 3,
		uintptr(handle),
		uintptr(interval),
		uintptr(event))
	return PDH_STATUS(ret1)
}

func PdhCollectQueryDataWithTime(handle PDH_HQUERY, timestamp *LONGLONG) PDH_STATUS {
	ret1 := syscall3(pdhCollectQueryDataWithTime, 2,
		uintptr(handle),
		uintptr(unsafe.Pointer(timestamp)),
		0)
	return PDH_STATUS(ret1)
}

func PdhEnumObjectItems(szDataSource string, szMachineName string, szObjectName string, mszCounterList LPWSTR, pcchCounterListLength *uint32, mszInstanceList LPWSTR, pcchInstanceListLength *uint32, dwDetailLevel uint32, dwFlags uint32) PDH_STATUS {
	szDataSourceStr := unicode16FromString(szDataSource)
	szMachineNameStr := unicode16FromString(szMachineName)
	szObjectNameStr := unicode16FromString(szObjectName)
	ret1 := syscall9(pdhEnumObjectItems, 9,
		uintptr(unsafe.Pointer(&szDataSourceStr[0])),
		uintptr(unsafe.Pointer(&szMachineNameStr[0])),
		uintptr(unsafe.Pointer(&szObjectNameStr[0])),
		uintptr(unsafe.Pointer(mszCounterList)),
		uintptr(unsafe.Pointer(pcchCounterListLength)),
		uintptr(unsafe.Pointer(mszInstanceList)),
		uintptr(unsafe.Pointer(pcchInstanceListLength)),
		uintptr(dwDetailLevel),
		uintptr(dwFlags))
	return PDH_STATUS(ret1)
}

func PdhExpandCounterPath(szWildCardPath string, mszExpandedPathList LPWSTR, pcchPathListLength *uint32) PDH_STATUS {
	szWildCardPathStr := unicode16FromString(szWildCardPath)
	ret1 := syscall3(pdhExpandCounterPath, 3,
		uintptr(unsafe.Pointer(&szWildCardPathStr[0])),
		uintptr(unsafe.Pointer(mszExpandedPathList)),
		uintptr(unsafe.Pointer(pcchPathListLength)))
	return PDH_STATUS(ret1)
}

func PdhExpandWildCardPath(szDataSource string, szWildCardPath string, mszExpandedPathList LPWSTR, pcchPathListLength *uint32, dwFlags uint32) PDH_STATUS {
	szDataSourceStr := unicode16FromString(szDataSource)
	szWildCardPathStr := unicode16FromString(szWildCardPath)
	ret1 := syscall6(pdhExpandWildCardPath, 5,
		uintptr(unsafe.Pointer(&szDataSourceStr[0])),
		uintptr(unsafe.Pointer(&szWildCardPathStr[0])),
		uintptr(unsafe.Pointer(mszExpandedPathList)),
		uintptr(unsafe.Pointer(pcchPathListLength)),
		uintptr(dwFlags),
		0)
	return PDH_STATUS(ret1)
}

func PdhGetCounterInfo(handle PDH_HCOUNTER, text BOOLEAN, size *uint32, info *PDH_COUNTER_INFO) PDH_STATUS {
	ret1 := syscall6(pdhGetCounterInfo, 4,
		uintptr(handle),
		uintptr(text),
		uintptr(unsafe.Pointer(size)),
		uintptr(unsafe.Pointer(info)),
		0,
		0)
	return PDH_STATUS(ret1)
}

func PdhGetCounterTimeBase(handle PDH_HCOUNTER, base *LONGLONG) PDH_STATUS {
	ret1 := syscall3(pdhGetCounterTimeBase, 2,
		uintptr(handle),
		uintptr(unsafe.Pointer(base)),
		0)
	return PDH_STATUS(ret1)
}

func PdhGetDllVersion(version *uint32) PDH_STATUS {
	ret1 := syscall3(pdhGetDllVersion, 1,
		uintptr(unsafe.Pointer(version)),
		0,
		0)
	return PDH_STATUS(ret1)
}

func PdhGetFormattedCounterValue(handle PDH_HCOUNTER, format uint32, aType *uint32, value *PDH_FMT_COUNTERVALUE) PDH_STATUS {
	ret1 := syscall6(pdhGetFormattedCounterValue, 4,
		uintptr(handle),
		uintptr(format),
		uintptr(unsafe.Pointer(aType)),
		uintptr(unsafe.Pointer(value)),
		0,
		0)
	return PDH_STATUS(ret1)
}

func PdhGetLogFileType(log /*const*/ *WCHAR, aType *uint32) PDH_STATUS {
	ret1 := syscall3(pdhGetLogFileType, 2,
		uintptr(unsafe.Pointer(log)),
		uintptr(unsafe.Pointer(aType)),
		0)
	return PDH_STATUS(ret1)
}

func PdhGetRawCounterValue(handle PDH_HCOUNTER, aType *uint32, value *PDH_RAW_COUNTER) PDH_STATUS {
	ret1 := syscall3(pdhGetRawCounterValue, 3,
		uintptr(handle),
		uintptr(unsafe.Pointer(aType)),
		uintptr(unsafe.Pointer(value)))
	return PDH_STATUS(ret1)
}

func PdhLookupPerfIndexByName(machine string, name string, index *uint32) PDH_STATUS {
	machineStr := unicode16FromString(machine)
	nameStr := unicode16FromString(name)
	ret1 := syscall3(pdhLookupPerfIndexByName, 3,
		uintptr(unsafe.Pointer(&machineStr[0])),
		uintptr(unsafe.Pointer(&nameStr[0])),
		uintptr(unsafe.Pointer(index)))
	return PDH_STATUS(ret1)
}

func PdhLookupPerfNameByIndex(machine string, index uint32, buffer LPWSTR, size *uint32) PDH_STATUS {
	machineStr := unicode16FromString(machine)
	ret1 := syscall6(pdhLookupPerfNameByIndex, 4,
		uintptr(unsafe.Pointer(&machineStr[0])),
		uintptr(index),
		uintptr(unsafe.Pointer(buffer)),
		uintptr(unsafe.Pointer(size)),
		0,
		0)
	return PDH_STATUS(ret1)
}

func PdhMakeCounterPath(e *PDH_COUNTER_PATH_ELEMENTS, buffer LPWSTR, buflen *uint32, flags uint32) PDH_STATUS {
	ret1 := syscall6(pdhMakeCounterPath, 4,
		uintptr(unsafe.Pointer(e)),
		uintptr(unsafe.Pointer(buffer)),
		uintptr(unsafe.Pointer(buflen)),
		uintptr(flags),
		0,
		0)
	return PDH_STATUS(ret1)
}

func PdhOpenQuery(source string, userdata *uint32, handle *PDH_HQUERY) PDH_STATUS {
	sourceStr := unicode16FromString(source)
	ret1 := syscall3(pdhOpenQuery, 3,
		uintptr(unsafe.Pointer(&sourceStr[0])),
		uintptr(unsafe.Pointer(userdata)),
		uintptr(unsafe.Pointer(handle)))
	return PDH_STATUS(ret1)
}

func PdhRemoveCounter(handle PDH_HCOUNTER) PDH_STATUS {
	ret1 := syscall3(pdhRemoveCounter, 1,
		uintptr(handle),
		0,
		0)
	return PDH_STATUS(ret1)
}

func PdhSetCounterScaleFactor(handle PDH_HCOUNTER, factor LONG) PDH_STATUS {
	ret1 := syscall3(pdhSetCounterScaleFactor, 2,
		uintptr(handle),
		uintptr(factor),
		0)
	return PDH_STATUS(ret1)
}

func PdhSetDefaultRealTimeDataSource(source uint32) PDH_STATUS {
	ret1 := syscall3(pdhSetDefaultRealTimeDataSource, 1,
		uintptr(source),
		0,
		0)
	return PDH_STATUS(ret1)
}

func PdhValidatePathEx(source PDH_HLOG, path string) PDH_STATUS {
	pathStr := unicode16FromString(path)
	ret1 := syscall3(pdhValidatePathEx, 2,
		uintptr(source),
		uintptr(unsafe.Pointer(&pathStr[0])),
		0)
	return PDH_STATUS(ret1)
}

func PdhValidatePath(path string) PDH_STATUS {
	pathStr := unicode16FromString(path)
	ret1 := syscall3(pdhValidatePath, 1,
		uintptr(unsafe.Pointer(&pathStr[0])),
		0,
		0)
	return PDH_STATUS(ret1)
}