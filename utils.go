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
	"fmt"
	"image/color"
	"io"
	"log"
	"reflect"
	"strings"
	"text/tabwriter"
	"time"
)

func nt2time(nt C.NET_TIME) time.Time {

	return time.Date(int(nt.dwYear), time.Month(nt.dwMonth), int(nt.dwDay), int(nt.dwHour), int(nt.dwMinute), int(nt.dwSecond), 0, time.UTC)
}

func time2nt(t time.Time) C.NET_TIME {
	return C.NET_TIME{
		dwYear:   C.uint(t.Year()),   // 年
		dwMonth:  C.uint(t.Month()),  // 月
		dwDay:    C.uint(t.Day()),    // 日
		dwHour:   C.uint(t.Hour()),   // 时
		dwMinute: C.uint(t.Minute()), // 分
		dwSecond: C.uint(t.Second()), // 秒
	}
}

func ntex2time(nt C.NET_TIME_EX) time.Time {
	log.Printf("nt %#v", nt)
	return time.Date(
		int(nt.dwYear),
		time.Month(nt.dwMonth),
		int(nt.dwDay),
		int(nt.dwHour),
		int(nt.dwMinute),
		int(nt.dwSecond),
		int(nt.dwMillisecond*1e6),
		time.Local)
}

func time2ntex(t time.Time) C.NET_TIME_EX {
	nano := t.UnixNano() % 1e6

	return C.NET_TIME_EX{
		dwYear:        C.uint(t.Year()),   // 年
		dwMonth:       C.uint(t.Month()),  // 月
		dwDay:         C.uint(t.Day()),    // 日
		dwHour:        C.uint(t.Hour()),   // 时
		dwMinute:      C.uint(t.Minute()), // 分
		dwSecond:      C.uint(t.Second()), // 秒
		dwMillisecond: C.uint(nano),
		dwUTC:         0,
	}
}

func FormatString(val interface{}) string {
	var (
		sb strings.Builder
		v  = reflect.ValueOf(val)
	)
	// v = reflect.Indirect(v)\
	w := tabwriter.NewWriter(&sb, 0, 0, 4, ' ', tabwriter.AlignRight|tabwriter.Debug)
	formatStringIndent(v, 0, w)
	w.Flush()
	return sb.String()
}

var timType = reflect.TypeOf(time.Time{})

func formatStringIndent(v reflect.Value, depth int, w io.Writer) {
	t := v.Type()
	for i := 0; i < t.NumMethod(); i++ {
		meth := t.Method(i)
		if strings.HasPrefix(meth.Name, "Set") {
			continue
		}

		if meth.Name == "String" {
			continue
		}

		realmeth := v.Method(i)
		// vals := realmeth.Call([]reflect.Value{reflect.ValueOf(devinfo)})
		inarg := meth.Type.NumIn()
		outarg := meth.Type.NumOut()
		// log.Printf("%s in %d out %d\n", meth.Name, inarg, outarg)

		if !(inarg == 1 && outarg == 1) {
			continue
		}

		vals := realmeth.Call([]reflect.Value{})
		if len(vals) == 1 {
			ident := strings.Repeat("\t", depth)

			rv := vals[0]
			switch rv.Kind() {
			case reflect.Struct:
				if rv.Type() == timType {
					if rv.CanInterface() {
						// log.Printf("time %s", meth.Name)
						fmt.Fprintf(w, "%s\t%s\t%v\n", meth.Name, ident, rv.Interface())
					}
				} else {
					formatStringIndent(rv, depth+1, w)
				}
			default:
				if rv.CanInterface() {
					// log.Printf("other %s", meth.Name)

					fmt.Fprintf(w, "%s\t%s\t%v\n", meth.Name, ident, rv.Interface())
				}
			}
		}
	}
}

func hex2rgb(hclr int) color.Color {
	return color.RGBA{R: uint8(hclr >> 24), G: uint8((hclr & 0xff00) >> 16), B: uint8((hclr & 0xff00) >> 8)}
}

func rgb2hex(clr color.RGBA) int {
	return int(clr.R)<<24 | int(clr.G)<<16 | int(clr.B)<<8
}
