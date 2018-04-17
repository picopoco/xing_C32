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
// #include "./func.h"
import "C"

import (
	"github.com/ghts/lib"
	"github.com/ghts/xing_types"

	"os"
)

func F초기화() (에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러}.S실행()

	f초기화_설정화일()
	f초기화_XingAPI()
	f초기화_Go루틴()

	lib.F메모("초기 접속에서 자주 에러가 발생함. 이유를 모르겠음.")
	f초기화_서버_접속()
	f초기화_TR전송_제한()

	lib.F메모("초기화 일부 과정 임시 보류..")
	//f초기화_소켓()
	//f초기화_완료_통보()

	return nil
}

func f초기화_설정화일() {
	lib.F조건부_패닉(lib.F파일_없음(설정화일_경로), "설정화일 config.ini를 찾을 수 없습니다.")
}

func f초기화_XingAPI() {
	// DLL파일이 있는 디렉토리로 이동. (빼먹으면 안 됨)
	원래_디렉토리, 에러 := os.Getwd()
	에러체크(에러)

	xing디렉토리, 에러 := XingAPI디렉토리()
	에러체크(에러)

	에러체크(os.Chdir(xing디렉토리))

	// XingAPI 초기화 ('반드시' DLL파일이 있는 디렉토리에서 실행해야 함.)
	C.initXingApi(0)

	// 원래 디렉토리로 이동
	에러체크(os.Chdir(원래_디렉토리))
}

func f초기화_Go루틴() {
	ch초기화 := make(chan lib.T신호)
	go Go루틴_소켓_C함수_호출(ch초기화)
	<-ch초기화
}

func f초기화_서버_접속() {
	lib.F조건부_패닉(!lib.F인터넷에_접속됨(), "서버 접속이 불가 : 인터넷 접속을 확인하십시오.")

	질의값 := xt.New호출_인수_기본형(xt.P함수_접속)
	소켓_질의 := lib.New소켓_질의_단순형(lib.P주소_Xing_C함수_호출, lib.F임의_변환_형식(), lib.P30초)
	응답 := 소켓_질의.S질의(질의값).G응답_검사()
	값 := 에러체크(응답.G해석값(0))

	switch 변환값 := 값.(type) {
	case bool:
		로그인_성공 := 응답.G해석값_단순형(0).(bool)

		if !로그인_성공 && lib.F테스트_모드_실행_중() {
			lib.F문자열_출력("모의투자 기간이 종료되었는 지 확인하십시오.\n%v")
		} else if !로그인_성공 {
			panic("접속 실패.")
		}
	case error:
		if lib.F테스트_모드_실행_중() {
			lib.F문자열_출력("모의투자 기간이 종료되었는 지 확인하십시오.\n%v")
		}

		panic(변환값)
	}
}

func f초기화_소켓() {
	var 에러 error
	소켓PUB_콜백, 에러 = lib.New소켓PUB(lib.P주소_Xing_C함수_콜백)
	에러체크(에러)

	소켓PUB_실시간_정보, 에러 = lib.New소켓PUB(lib.P주소_Xing_C함수_실시간)
	에러체크(에러)

	lib.F대기(lib.P1초)
}

func f초기화_완료_통보() {
	질의값 := lib.New질의값_기본형()
	질의값.TR구분 = lib.TR초기화
	질의값.TR코드 = "C함수 모듈"

	lib.New소켓_질의_단순형(lib.P주소_Xing_TR, lib.P변환형식_기본값, lib.P30초).S질의(질의값).G응답_검사()
}

func f초기화_TR전송_제한() {
	코드별_10분당_TR전송_제한 := make(map[string]int)
	코드별_10분당_TR전송_제한[xt.TR현물_기간별_조회] = 200
	코드별_10분당_TR전송_제한[xt.TR_ETF_시간별_추이] = 200
	코드별_10분당_TR전송_제한[xt.TR현물계좌_예수금_주문가능금액_총평가] = 200
	코드별_10분당_TR전송_제한[xt.TR현물계좌_잔고내역] = 200
	코드별_10분당_TR전송_제한[xt.TR현물계좌_주문체결내역] = 200
	코드별_10분당_TR전송_제한[xt.TR주식계좌_기간별_수익률_상세] = 200
	코드별_10분당_TR전송_제한[xt.TR계좌별_신용한도] = 200
	코드별_10분당_TR전송_제한[xt.TR현물계좌_증거금률별_주문가능수량] = 200
	코드별_10분당_TR전송_제한[xt.TR종목별_증시_일정] = 200
	코드별_10분당_TR전송_제한[xt.TR해외_실시간_지수] = 200
	코드별_10분당_TR전송_제한[xt.TR해외_지수_조회] = 200
	코드별_10분당_TR전송_제한[xt.TR증시_주변_자금_추이] = 200

	for TR코드, 초당_제한_횟수 := range 코드별_10분당_TR전송_제한 {
		tr전송_코드별_10분당_제한[TR코드] = lib.New전송_권한_TR코드별(TR코드, 초당_제한_횟수, lib.P10분)
	}

	코드별_초당_TR전송_제한 := make(map[string]int)
	코드별_초당_TR전송_제한[xt.TR계좌_번호] = 100
	코드별_초당_TR전송_제한[xt.TR현물_정상주문] = 30
	코드별_초당_TR전송_제한[xt.TR현물_정정주문] = 30
	코드별_초당_TR전송_제한[xt.TR현물_취소주문] = 30
	코드별_초당_TR전송_제한[xt.TR현물_호가_조회] = 5
	코드별_초당_TR전송_제한[xt.TR현물_시세_조회] = 5
	코드별_초당_TR전송_제한[xt.TR현물_기간별_조회] = 1
	코드별_초당_TR전송_제한[xt.TR현물_당일_전일_분틱_조회] = 1
	코드별_초당_TR전송_제한[xt.TR_ETF_시간별_추이] = 1
	코드별_초당_TR전송_제한[xt.TR현물_종목_조회] = 2
	코드별_초당_TR전송_제한[xt.TR계좌_거래_내역] = 1
	코드별_초당_TR전송_제한[xt.TR현물계좌_예수금_주문가능금액_총평가] = 1
	코드별_초당_TR전송_제한[xt.TR현물계좌_잔고내역] = 1
	코드별_초당_TR전송_제한[xt.TR현물계좌_주문체결내역] = 1
	코드별_초당_TR전송_제한[xt.TR주식_체결_미체결] = 1
	코드별_초당_TR전송_제한[xt.TR주식_매매일지_수수료_금일] = 2
	코드별_초당_TR전송_제한[xt.TR주식_매매일지_수수료_날짜_지정] = 2
	코드별_초당_TR전송_제한[xt.TR주식_잔고_2] = 1
	코드별_초당_TR전송_제한[xt.TR시간_조회] = 5
	코드별_초당_TR전송_제한[xt.TR주식계좌_기간별_수익률_상세] = 1
	코드별_초당_TR전송_제한[xt.TR계좌별_신용한도] = 1
	코드별_초당_TR전송_제한[xt.TR현물계좌_증거금률별_주문가능수량] = 1
	코드별_초당_TR전송_제한[xt.TR종목별_증시_일정] = 1
	코드별_초당_TR전송_제한[xt.TR해외_실시간_지수] = 1
	코드별_초당_TR전송_제한[xt.TR해외_지수_조회] = 1
	코드별_초당_TR전송_제한[xt.TR증시_주변_자금_추이] = 1

	for TR코드, 초당_제한_횟수 := range 코드별_초당_TR전송_제한 {
		tr전송_코드별_초당_제한[TR코드] = lib.New전송_권한_TR코드별(TR코드, 초당_제한_횟수, lib.P1초)
	}
}
