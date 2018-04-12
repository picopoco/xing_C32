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
	"github.com/ghts/types_xing"
	"unsafe"
)

func f콜백(응답값 interface{}) lib.I소켓_메시지 {
	return lib.New소켓_질의(lib.P주소_Xing_C함수_콜백, lib.P변환형식_기본값, lib.P30초).S질의(응답값).G응답()
}

//export OnDisconnected_Go
func OnDisconnected_Go() {
	응답값 := new(xing.S콜백_단순형)
	응답값.M콜백 = xing.P콜백_접속해제

	f콜백(응답값)
}

//export OnTrData_Go
func OnTrData_Go(c *C.TR_DATA_UNPACKED) {
	defer F메모리_해제(unsafe.Pointer(c))

	g := (*TR_DATA)(unsafe.Pointer(c))

	응답값 := new(xing.S콜백_TR_데이터)
	응답값.M콜백 = xing.P콜백_데이터
	응답값.M식별번호 = int(g.RequestID)
	응답값.TR코드 = lib.F2문자열(g.TrCode)
	응답값.M소요시간_ms = int(g.ElapsedTime)
	응답값.M데이터_모드 = int8(g.DataMode)

	switch lib.F2문자열(g.Cont) {
	case "0", "N":
		응답값.M연속조회_여부 = false
	case "1", "Y":
		응답값.M연속조회_여부 = true
	default:
		응답값.M연속조회_여부 = false
	}

	응답값.M연속키 = lib.F2문자열(g.ContKey)
	응답값.M블록_이름 = lib.F2문자열(g.BlockName)
	응답값.M데이터 = C.GoBytes(unsafe.Pointer(g.Data), c.DataLength)

	f콜백(응답값)
}

//export OnMessageAndError_Go
func OnMessageAndError_Go(c *C.MSG_DATA_UNPACKED, pointer *C.MSG_DATA) {
	defer F메모리_해제(unsafe.Pointer(c))

	g := (*MSG_DATA)(unsafe.Pointer(c))

	메시지_저장소.S추가(int(g.RequestID), unsafe.Pointer(pointer))

	응답값 := new(xing.S콜백_메시지_및_에러)
	응답값.M콜백 = xing.P콜백_메시지_및_에러
	응답값.M식별번호 = int(g.RequestID)

	switch g.SystemError {
	case 0: // 일반 메시지
		응답값.M에러여부 = false
	case 1: // 에러 메시지
		응답값.M에러여부 = true
	default:
		lib.F패닉("예상하지 못한 구분값. '%v'", g.SystemError)
	}

	f콜백(응답값)
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
	defer F메모리_해제(unsafe.Pointer(c))

	g := (*REALTIME_DATA)(unsafe.Pointer(c))

	응답값 := new(xing.S콜백_실시간_데이터)
	응답값.M콜백 = xing.P콜백_데이터
	응답값.TR코드 = lib.F2문자열(g.TrCode)
	응답값.M키_데이터 = C.GoBytes(unsafe.Pointer(&g.KeyData), c.KeyLength)
	응답값.M등록키 = C.GoBytes(unsafe.Pointer(&g.RegKey), C.int(len(g.RegKey)))
	응답값.M데이터 = C.GoBytes(unsafe.Pointer(&g.Data), c.DataLength)

	f콜백(응답값)
}

//export OnLogin_Go
func OnLogin_Go(wParam *C.char, lParam *C.char) {
	코드 := C.GoString(wParam)
	//메시지 := C.GoString(lParam)

	응답값 := new(xing.S콜백_로그인)
	응답값.M콜백 = xing.P콜백_로그인

	정수, 에러 := lib.F2정수(코드)
	if 에러 == nil && 정수 == 0 {
		응답값.M로그인_성공_여부 = true
	} else {
		응답값.M로그인_성공_여부 = false
	}

	// 계좌정보 설정
	계좌_수량 := F계좌_수량()
	응답값.M계좌번호_모음 = make([]string, 계좌_수량)

	for i := 0; i < 계좌_수량; i++ {
		응답값.M계좌번호_모음[i], 에러 = F계좌_번호(i)
		lib.F에러체크(에러)

		lib.F문자열_출력("계좌번호 %v : %v", i, 응답값.M계좌번호_모음[i])
	}

	f콜백(응답값)
}

//export OnLogout_Go
func OnLogout_Go() {
	응답값 := new(xing.S콜백_단순형)
	응답값.M콜백 = xing.P콜백_로그아웃

	f콜백(응답값)
}

//export OnTimeout_Go
func OnTimeout_Go(c C.int) {
	응답값 := new(xing.S콜백_정수값)
	응답값.M콜백 = xing.P콜백_타임아웃
	응답값.M정수값 = int(c)

	f콜백(응답값)
}

//export OnLinkData_Go
func OnLinkData_Go() { // TODO
	응답값 := new(xing.S콜백_단순형)
	응답값.M콜백 = xing.P콜백_링크_데이터

	f콜백(응답값)
}

//export OnRealtimeDataChart_Go
func OnRealtimeDataChart_Go() { // TODO
	응답값 := new(xing.S콜백_단순형)
	응답값.M콜백 = xing.P콜백_실시간_차트_데이터

	f콜백(응답값)
}
