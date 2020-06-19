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
import "unsafe"

type MediaQueryTrafficcarEx struct {
	cptr  *C.MEDIA_QUERY_TRAFFICCAR_PARAM_EX
	Param *MediaQueryTrafficcar
}

func (carex *MediaQueryTrafficcarEx) init() {
	if carex.cptr != nil {
		return
	}

	carex.cptr = &C.MEDIA_QUERY_TRAFFICCAR_PARAM_EX{}
	carex.cptr.dwSize = C.uint(unsafe.Sizeof(*carex.cptr))
	carex.Param = &MediaQueryTrafficcar{cptr: &carex.cptr.stuParam}
}

func (carex *MediaQueryTrafficcarEx) Init() {
	carex.init()
}
