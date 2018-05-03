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
// #include <stdlib.h>
// #include <windef.h>
// #include "./func.h"
import "C"

import (
	"github.com/ghts/lib"
	"github.com/ghts/xing_types"
	"gopkg.in/ini.v1"

	"bytes"
	"time"
	"unsafe"
)

func F접속(서버_구분 xt.T서버_구분) bool {
	if F접속됨() {
		return true
	}

	var c서버_이름 *C.char
	var c포트_번호 C.int

	switch 서버_구분 {
	case xt.P서버_실거래:
		if lib.F테스트_모드_실행_중() {
			panic("테스트 모드에서 실서버 접속 시도.")
		}

		c서버_이름 = C.CString("hts.ebestsec.co.kr")
		c포트_번호 = C.int(20001)
	case xt.P서버_모의투자:
		if !lib.F테스트_모드_실행_중() {
			panic("실제 운용 모드에서 모의투자서버 접속 시도.")
		}

		c서버_이름 = C.CString("demo.ebestsec.co.kr")
		c포트_번호 = C.int(20001)
	case xt.P서버_XingACE:
		if !lib.F테스트_모드_실행_중() {
			panic("실제 운용 모드에서 XingACE 가상거래소 접속 시도.")
		}

		c서버_이름 = C.CString("127.0.0.1")
		c포트_번호 = C.int(0)
	}

	defer C.free(unsafe.Pointer(c서버_이름))

	접속_결과 := bool(C.etkConnect(c서버_이름, c포트_번호))

	return 접속_결과
}

func F접속됨() bool {
	return bool(C.etkIsConnected())
}

func F로그인() (로그인_결과 bool) {
	defer lib.S에러패닉_처리기{M함수: func() { 로그인_결과 = false }}.S실행_No출력()

	if lib.F파일_없음(설정화일_경로) {
		버퍼 := new(bytes.Buffer)
		버퍼.WriteString("Xing 설정화일 없음\n")
		버퍼.WriteString("%v가 존재하지 않습니다.\n")
		버퍼.WriteString("sample_config.ini를 참조하여 새로 생성하십시오.")
		panic(lib.New에러(버퍼.String(), 설정화일_경로))
	}

	cfg파일 := 에러체크(ini.Load(설정화일_경로)).(*ini.File)
	섹션 := 에러체크(cfg파일.GetSection("XingAPI_LogIn_Info")).(*ini.Section)

	키_ID := 에러체크(섹션.GetKey("ID")).(*ini.Key)
	c아이디 := C.CString(키_ID.String())

	키_PWD := 에러체크(섹션.GetKey("PWD")).(*ini.Key)
	c암호 := C.CString(키_PWD.String())

	키_CertPWD := 에러체크(섹션.GetKey("CertPWD")).(*ini.Key)
	공인인증서_암호 := lib.F조건부_값(lib.F테스트_모드_실행_중(), "", 키_CertPWD.String()).(string)
	c공인인증서_암호 := C.CString(공인인증서_암호)

	로그인_결과 = bool(C.etkLogin(c아이디, c암호, c공인인증서_암호))

	C.free(unsafe.Pointer(c아이디))
	C.free(unsafe.Pointer(c암호))
	C.free(unsafe.Pointer(c공인인증서_암호))

	return 로그인_결과
}

func F로그아웃_및_접속해제() error {
	if !bool(C.etkLogout()) {
		return lib.New에러("로그아웃 실패.")
	}

	if !bool(C.etkDisconnect()) {
		return lib.New에러("접속 해제 실패.")
	}

	for F접속됨() {
		lib.F대기(lib.P500밀리초)
	}

	return nil
}

func F질의(TR코드 string, c데이터 unsafe.Pointer, 길이 int,
	연속_조회_여부 bool, 연속키 string, 타임아웃 time.Duration) int {
	cTR코드 := C.CString(TR코드)
	c길이 := C.int(길이)
	c연속_조회_여부 := C.bool(연속_조회_여부)
	c연속_조회_키 := C.CString(연속키)
	c타임아웃 := C.int(타임아웃 / time.Second)

	defer func() {
		C.free(unsafe.Pointer(cTR코드))
		C.free(unsafe.Pointer(c연속_조회_키))
	}()

	c식별번호 := C.etkRequest(cTR코드, c데이터, c길이, c연속_조회_여부, c연속_조회_키, c타임아웃)
	return int(c식별번호)
}

func F실시간_정보_구독(TR코드 string, 전체_종목코드 string, 단위_길이 int) error {
	cTR코드 := C.CString(TR코드)
	c전체_종목코드 := C.CString(전체_종목코드)
	c단위_길이 := C.int(단위_길이)

	defer func() {
		C.free(unsafe.Pointer(cTR코드))
		C.free(unsafe.Pointer(c전체_종목코드))
	}()

	if !bool(C.etkAdviseRealData(cTR코드, c전체_종목코드, c단위_길이)) {
		return lib.New에러("실시간 정보 신청 실패. %v", 전체_종목코드)
	}

	return nil
}

func F실시간_정보_해지(TR코드 string, 전체_종목코드 string, 단위_길이 int) error {
	cTR코드 := C.CString(TR코드)
	c전체_종목코드 := C.CString(전체_종목코드)
	c단위_길이 := C.int(단위_길이)

	defer func() {
		C.free(unsafe.Pointer(cTR코드))
		C.free(unsafe.Pointer(c전체_종목코드))
	}()

	if !bool(C.etkUnadviseRealData(cTR코드, c전체_종목코드, c단위_길이)) {
		return lib.New에러("실시간 정보 해지 실패. %v", 전체_종목코드)
	}

	return nil
}

func F실시간_정보_모두_해지() error {
	if !bool(C.etkUnadviseWindow()) {
		return lib.New에러("실시간 정보 모두 해지 실패. %v")
	}

	return nil
}

func F계좌_수량() int { return int(C.etkGetAccountListCount()) }

func F계좌_번호(인덱스 int) string {
	버퍼_초기값 := "            " // 12자리 공백문자열
	버퍼_크기 := C.int(len(버퍼_초기값))
	c버퍼 := C.CString(버퍼_초기값)
	defer C.free(unsafe.Pointer(c버퍼))

	C.etkGetAccountNo(C.int(인덱스), c버퍼, 버퍼_크기)

	바이트_모음 := C.GoBytes(unsafe.Pointer(c버퍼), 버퍼_크기)
	return lib.F2문자열_공백제거(바이트_모음)
}

func F계좌_이름(계좌_번호 string) string {
	버퍼_초기값 := "                                         "
	버퍼_크기 := C.int(len(버퍼_초기값))
	c버퍼 := C.CString(버퍼_초기값)
	c계좌번호 := C.CString(계좌_번호)

	defer func() {
		C.free(unsafe.Pointer(c버퍼))
		C.free(unsafe.Pointer(c계좌번호))
	}()

	C.etkGetAccountName(c계좌번호, c버퍼, 버퍼_크기)

	바이트_모음 := C.GoBytes(unsafe.Pointer(c버퍼), 버퍼_크기)
	return lib.F2문자열_EUC_KR(바이트_모음)

	//return C.GoString(c버퍼)
}

func F계좌_상세명(계좌_번호 string) string {
	버퍼_초기값 := "                                         "
	버퍼_크기 := C.int(len(버퍼_초기값))
	c버퍼 := C.CString(버퍼_초기값)
	c계좌번호 := C.CString(계좌_번호)

	defer func() {
		C.free(unsafe.Pointer(c버퍼))
		C.free(unsafe.Pointer(c계좌번호))
	}()

	C.etkGetAccountDetailName(c계좌번호, c버퍼, 버퍼_크기)

	바이트_모음 := C.GoBytes(unsafe.Pointer(c버퍼), 버퍼_크기)
	return lib.F2문자열_EUC_KR(바이트_모음)

	//return C.GoString(c버퍼)
}

// 원인미상의 메모리 에러가 발생함.
//func F계좌_별명(계좌_번호 string) string {
//	버퍼_초기값 := "                                                     "
//	버퍼_크기 := len(버퍼_초기값)
//
//	c버퍼 := C.CString(버퍼_초기값)
//	c계좌번호 := C.CString(계좌_번호)
//
//	defer func() {
//		if c버퍼 != nil {
//			C.free(unsafe.Pointer(c버퍼))
//		}
//
//		if c계좌번호 != nil {
//			C.free(unsafe.Pointer(c계좌번호))
//		}
//	}()
//
//	C.etkGetAccountNickName(c계좌번호, c버퍼, C.int(버퍼_크기))
//
//	바이트_모음 := C.GoBytes(unsafe.Pointer(c버퍼), C.int(버퍼_크기))
//
//	return lib.F2문자열_EUC_KR(바이트_모음)
//}

func F서버_이름() string {
	버퍼_초기값 := "                                                   "
	버퍼_길이 := C.int(len(버퍼_초기값))
	c버퍼 := C.CString(버퍼_초기값)
	defer C.free(unsafe.Pointer(c버퍼))

	C.etkGetServerName(c버퍼, 버퍼_길이)

	바이트_모음 := C.GoBytes(unsafe.Pointer(c버퍼), 버퍼_길이)
	return lib.F2문자열_EUC_KR_공백제거(바이트_모음)
}

func F에러_코드() int { return int(C.etkGetLastError(0)) }

func F에러_메시지(에러_코드 int) string {
	go버퍼 := new(bytes.Buffer)
	for i := 0; i < 512; i++ {
		go버퍼.WriteString(" ")
	}

	버퍼_초기값 := go버퍼.String()
	버퍼_길이 := C.int(len(버퍼_초기값))
	c버퍼 := C.CString(버퍼_초기값)
	defer C.free(unsafe.Pointer(c버퍼))

	에러_메시지_길이 := C.etkGetErrorMessage(C.int(에러_코드), c버퍼, 버퍼_길이)

	if 에러_메시지_길이 == 0 {
		lib.New에러("에러 메시지를 구할 수 없습니다.")
		return ""
	}

	바이트_모음 := C.GoBytes(unsafe.Pointer(c버퍼), 버퍼_길이)
	return lib.F2문자열_EUC_KR_공백제거(바이트_모음)
}

func F초당_TR쿼터(TR코드 string) int {
	cTR코드 := C.CString(TR코드)
	defer C.free(unsafe.Pointer(cTR코드))

	return int(C.etkGetTRCountPerSec(cTR코드))
}

func F압축_해제(데이터 unsafe.Pointer, 데이터_길이 int) []byte {
	defer C.free(데이터) // 이게 문제가 될까?

	const 버퍼_길이 = 2000
	바이트_모음 := make([]byte, 버퍼_길이, 버퍼_길이)
	버퍼 := unsafe.Pointer(&바이트_모음)

	길이 := C.etkDecompress((*C.char)(데이터), C.int(데이터_길이), (*C.char)(버퍼), 버퍼_길이)

	return C.GoBytes(버퍼, 길이)
}

func f함수_존재함(함수명 string) bool {
	c함수명 := C.CString(함수명)
	defer C.free(unsafe.Pointer(c함수명))

	return bool(C.etkFuncExist(c함수명))
}

func f데이터_해제(식별번호 int) {
	C.etkReleaseRequestData(C.int(식별번호))
}

func f메시지_해제(포인터 unsafe.Pointer) {
	c := (*C.MSG_DATA)(포인터)

	C.etkReleaseMessageData(c)
}

func F자원_해제() {
	C.freeResource(0)
}

func F메모리_해제(포인터 unsafe.Pointer) {
	C.free(포인터)
}

func C문자열(문자열 string) *C.char {
	return C.CString(문자열)
}
