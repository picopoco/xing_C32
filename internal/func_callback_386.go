/* Copyright (C) 2015-2018 김운하(UnHa Kim)  unha.kim@kuh.pe.kr

이 파일은 GHTS의 일부입니다.

이 프로그램은 자유 소프트웨어입니다.
소프트웨어의 피양도자는 자유 소프트웨어 재단이 공표한 GNU LGPL 2.1판
규정에 따라 프로그램을 개작하거나 재배포할 수 있습니다.

이 프로그램은 유용하게 사용될 수 있으리라는 희망에서 배포되고 있지만,
특정한 목적에 적합하다거나, 이익을 안겨줄 수 있다는 묵시적인 보증을 포함한
어떠한 형태의 보증도 제공하지 않습니다.
보다 자세한 사항에 대해서는 GNU LGPL 2.1판을 참고하시기 바랍니다.
GNU LGPL 2.1판은 이 프로그램과 함께 제공됩니다.
만약, 이 문서가 누락되어 있다면 자유 소프트웨어 재단으로 문의하시기 바랍니다.
(자유 소프트웨어 재단 : Free Software Foundation, Inc.,
59 Temple Place - Suite 330, Boston, MA 02111-1307, USA)

Copyright (C) 2015-2018년 UnHa Kim (unha.kim@kuh.pe.kr)

This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, version 2.1 of the License.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>. */

package xing_C32

// #cgo CFLAGS: -Wall
// #include <windows.h>
// #include "./type_c.h"
import "C"

import (
	"github.com/ghts/lib"
	"github.com/ghts/xing_types"
	"unsafe"
)

func f콜백(콜백값 interface{}) (에러 error) {
	defer lib.S에러패닉_처리기{ M에러_포인터:&에러 }.S실행()

	return lib.New소켓_메시지_단순형(lib.P변환형식_기본값, 콜백값).S소켓_송신_기본형(소켓PUB_콜백)
}

//export OnDisconnected_Go
func OnDisconnected_Go() {
	응답값 := new(xt.S콜백_단순형)
	응답값.M콜백 = xt.P콜백_접속해제

	f콜백(응답값)
}

//export OnTrData_Go
func OnTrData_Go(c *C.TR_DATA_UNPACKED) {
	defer func() {
		F메모리_해제(unsafe.Pointer(c))
		lib.S에러패닉_처리기{}.S실행() }()

	g := (*TR_DATA)(unsafe.Pointer(c))

	데이터 := 확인(tr데이터_해석(g))
	바이트_변환값 := 확인(lib.New바이트_변환_매개체(lib.P변환형식_기본값, 데이터)).(*lib.S바이트_변환_매개체)
	콜백값 := xt.New콜백_TR데이터(int(g.RequestID), 바이트_변환값)

	f콜백(콜백값)
}

//export OnMessageAndError_Go
func OnMessageAndError_Go(c *C.MSG_DATA_UNPACKED, pointer *C.MSG_DATA) {
	defer func() {
		if c != nil {
			F메모리_해제(unsafe.Pointer(c))
		}
	}()

	g := (*MSG_DATA)(unsafe.Pointer(c))

	var 에러여부 bool
	switch g.SystemError {
	case 0: // 일반 메시지
		에러여부 = false
	case 1: // 에러 메시지
		에러여부 = true
	default:
		panic(lib.F2문자열("예상하지 못한 구분값. '%v'", g.SystemError))
	}

	콜백값 := xt.New콜백_메시지_및_에러()
	콜백값.M식별번호 = int(g.RequestID)
	콜백값.M코드 = lib.F2문자열_공백제거(g.MsgCode)
	콜백값.M내용 = lib.F2문자열_공백제거(g.MsgData)
	콜백값.M에러여부 = 에러여부

	// f메시지_해제() 에서 포인터가 필요함.
	메시지_저장소.S추가(콜백값.M식별번호, unsafe.Pointer(pointer))

	f콜백(콜백값)
}

//export OnReleaseData_Go
func OnReleaseData_Go(c C.int) {
	식별번호 := int(c)

	f데이터_해제(식별번호)

	메시지_모음 := 메시지_저장소.G값(식별번호)

	if 메시지_모음 != nil {
		for _, 메시지 := range 메시지_모음 {
			f메시지_해제(메시지)
		}
	}

	메시지_저장소.S삭제(식별번호)
}

//export OnRealtimeData_Go
func OnRealtimeData_Go(c *C.REALTIME_DATA_UNPACKED) {
	defer func() {
		F메모리_해제(unsafe.Pointer(c))
		lib.S에러패닉_처리기{}.S실행() }()

	g := (*REALTIME_DATA)(unsafe.Pointer(c))

	데이터 := 확인(f실시간_데이터_해석(g))
	바이트_변환값 := 확인(lib.New바이트_변환_매개체(lib.P변환형식_기본값, 데이터)).(*lib.S바이트_변환_매개체)
	값 := xt.New콜백_실시간_데이터(바이트_변환값)
	소켓_메시지 := 확인(lib.New소켓_메시지(lib.MsgPack, 값)).(lib.I소켓_메시지)
	확인(소켓_메시지.S소켓_송신_기본형(소켓PUB_실시간_정보))
}

//export OnLogin_Go
func OnLogin_Go(wParam *C.char, lParam *C.char) {
	코드 := C.GoString(wParam)
	//메시지 := C.GoString(lParam)

	// '접속_처리_잠금'에 의해서 접속 처리는 1개씩만 처리되므로, 'ch접속_처리'는 비어있어야 함.
	lib.F조건부_패닉(len(ch접속_처리) > 0, "'ch접속_처리'는 비어있어야 함.")

	if 정수, 에러 := lib.F2정수(코드); 에러 == nil && 정수 == 0 {
		ch접속_처리 <- true
	} else {
		lib.F체크포인트()
		ch접속_처리 <- false
	}
}

//export OnLogout_Go
func OnLogout_Go() {
	응답값 := new(xt.S콜백_단순형)
	응답값.M콜백 = xt.P콜백_로그아웃

	f콜백(응답값)
}

//export OnTimeout_Go
func OnTimeout_Go(c C.int) {
	응답값 := new(xt.S콜백_정수값)
	응답값.M콜백 = xt.P콜백_타임아웃
	응답값.M정수값 = int(c)

	f콜백(응답값)
}

//export OnLinkData_Go
func OnLinkData_Go() { // TODO
	응답값 := new(xt.S콜백_단순형)
	응답값.M콜백 = xt.P콜백_링크_데이터

	f콜백(응답값)
}

//export OnRealtimeDataChart_Go
func OnRealtimeDataChart_Go() { // TODO
	응답값 := new(xt.S콜백_단순형)
	응답값.M콜백 = xt.P콜백_실시간_차트_데이터

	f콜백(응답값)
}
