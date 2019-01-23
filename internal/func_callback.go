/* Copyright (C) 2015-2019 김운하(UnHa Kim)  unha.kim@kuh.pe.kr

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

Copyright (C) 2015-2019년 UnHa Kim (unha.kim@kuh.pe.kr)

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
// #include "./types_c.h"
import "C"

import (
	"github.com/ghts/lib"
	"github.com/ghts/xing"

	"fmt"
	"unsafe"
)

func F콜백(콜백값 xing.I콜백) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행_No출력()

	소켓REQ := 소켓REQ_저장소.G소켓()
	defer 소켓REQ_저장소.S회수(소켓REQ)

	i값 := 소켓REQ.G질의_응답_검사(lib.P변환형식_기본값, 콜백값).G해석값_단순형(0)

	switch 값 := i값.(type) {
	case error:
		return 값
	case lib.T신호:
		lib.F조건부_패닉(값 != lib.P신호_OK, "예상하지 못한 신호값 : '%v'", 값)
	default:
		panic(lib.New에러("예상하지 못한 자료형 : '%T'", i값))
	}

	return nil
}

//export OnTrData_Go
func OnTrData_Go(c *C.TR_DATA_UNPACKED) {
	var 바이트_변환값 *lib.S바이트_변환

	g := (*TR_DATA)(unsafe.Pointer(c))

	if 데이터, 에러 := tr데이터_해석(g); 에러 != nil {
		바이트_변환값 = 에러체크(lib.New바이트_변환(lib.JSON, 에러)).(*lib.S바이트_변환)
	} else {
		바이트_변환값 = 에러체크(lib.New바이트_변환(lib.P변환형식_기본값, 데이터)).(*lib.S바이트_변환)
	}

	콜백값 := xing.New콜백_TR데이터(int(g.RequestID), 바이트_변환값, lib.F2문자열_공백제거(g.TrCode))

	//lib.F체크포인트("C32 콜백 시작", lib.F2문자열_공백제거(g.TrCode), 콜백값.G콜백(), lib.F2문자열(g.BlockName))
	F콜백(콜백값)
	//lib.F체크포인트("C32 콜백 완료", lib.F2문자열_공백제거(g.TrCode), 콜백값.G콜백(), lib.F2문자열(g.BlockName))
}

//export OnMessageAndError_Go
func OnMessageAndError_Go(c *C.MSG_DATA_UNPACKED, pointer *C.MSG_DATA) {
	g := (*MSG_DATA)(unsafe.Pointer(c))

	var 에러여부 bool
	switch g.SystemError {
	case 0: // 일반 메시지
		에러여부 = false
	case 1: // 에러 메시지
		에러여부 = true
	default:
		panic(lib.New에러("예상하지 못한 구분값. '%v'", g.SystemError))
	}

	콜백값 := new(xing.S콜백_메시지_및_에러)
	콜백값.S콜백_기본형 = xing.New콜백_기본형(xing.P콜백_메시지_및_에러)
	콜백값.M식별번호 = int(g.RequestID)
	콜백값.M코드 = lib.F2문자열_공백제거(g.MsgCode)
	콜백값.M내용 = lib.F2문자열_EUC_KR_공백제거(C.GoBytes(unsafe.Pointer(g.MsgData), C.int(g.MsgLength)))
	콜백값.M에러여부 = 에러여부

	// f메시지_해제() 에서 포인터가 필요함.
	메시지_저장소.S추가(콜백값.M식별번호, unsafe.Pointer(pointer))

	if 에러여부 {
		체크(콜백값)
	}

	F콜백(콜백값)
}

//export OnReleaseData_Go
func OnReleaseData_Go(c C.int) {
	식별번호 := int(c)

	f데이터_해제(식별번호)

	if 메시지_모음 := 메시지_저장소.G값(식별번호); 메시지_모음 != nil {
		for _, 메시지 := range 메시지_모음 {
			f메시지_해제(메시지)
		}
	}

	메시지_저장소.S삭제(식별번호)

	F콜백(xing.New콜백_TR완료(식별번호))
}

//export OnRealtimeData_Go
func OnRealtimeData_Go(c *C.REALTIME_DATA_UNPACKED) {
	defer lib.S예외처리{}.S실행()

	g := (*REALTIME_DATA)(unsafe.Pointer(c))
	실시간_데이터 := 에러체크(f실시간_데이터_해석(g))

	소켓PUB_실시간_정보.S송신_검사(lib.P변환형식_기본값, 실시간_데이터)
}

//export OnLogin_Go
func OnLogin_Go(wParam *C.char, lParam *C.char) {
	코드 := C.GoString(wParam)
	정수, 에러 := lib.F2정수(코드)
	로그인_성공_여부 := (에러 == nil && 정수 == 0)

	if !로그인_성공_여부 && lib.F테스트_모드_실행_중() {
		fmt.Println("********************************")
		fmt.Println("*  모의 투자 기간을 확인하세요. *")
		fmt.Println("********************************")
		lib.F문자열_출력("")
	}

	select {
	case ch로그인 <- 로그인_성공_여부:
	default:
	}
}

//export OnLogout_Go
func OnLogout_Go() {
	// XingAPI가 신호를 보내오지 않음.  여기에 기능을 구현해 봤자 소용없음.
}

//export OnDisconnected_Go
func OnDisconnected_Go() {
	// XingAPI가 신호를 보내오지 않음.  여기에 기능을 구현해 봤자 소용없음.
}

//export OnTimeout_Go
func OnTimeout_Go(c C.int) {
	F콜백(xing.New콜백_타임아웃(int(c)))
}

//export OnLinkData_Go
func OnLinkData_Go() {
	F콜백(xing.New콜백_기본형(xing.P콜백_링크_데이터)) // TODO
}

//export OnRealtimeDataChart_Go
func OnRealtimeDataChart_Go() {
	F콜백(xing.New콜백_기본형(xing.P콜백_실시간_차트_데이터)) // TODO
}
