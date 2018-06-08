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
	"github.com/ghts/xing"

	"fmt"
	"os"
	"time"
)

func init() {
	lib.TR구분_String = xing.TR구분_String

	// t1310은 전일/당일 구분이 필요함.
	콜백_대기_저장소 = new(S콜백_대기_저장소)
	콜백_대기_저장소.저장소 = make(map[int]*S콜백_대기_항목)

	lib.F메모("의도하지 않게 로그아웃된 경우 재접속 하는 기능 구현할 것.")
	lib.F메모("F계좌_별명(): 원인미상의 메모리 에러가 발생함.")
	lib.F메모("콜백_대기_항목 을 주기적으로 정리 해야함.")
	lib.F메모("XingAPI 인스턴스를 여러 개 생성해서 다중처리 하는 방안을 알아볼 것.")
}

func F초기화() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	f초기화_TR전송_제한()
	f초기화_Go루틴()
	f초기화_서버_접속()

	return nil
}

func f초기화_XingAPI() {
	if API_초기화_완료.G값() {
		return
	} else {
		API_초기화_완료.S값(true)
	}

	f의존성_확인()

	lib.F조건부_패닉(lib.F환경변수("GOARCH") != "386", "C32 모듈은 32비트 전용입니다.")

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
	go Go소켓_C함수_호출(ch초기화)
	<-ch초기화
}

func f초기화_서버_접속() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	lib.F조건부_패닉(!lib.F인터넷에_접속됨(), "서버 접속이 불가 : 인터넷 접속을 확인하십시오.")

	const 타임아웃 = lib.P30초
	ch타임아웃 := time.After(타임아웃)

	질의값 := lib.New질의값_기본형(lib.TR접속, "")
	소켓REQ := lib.NewNano소켓REQ_단순형(lib.P주소_Xing_C함수_호출, lib.P30초, 타임아웃)
	defer 소켓REQ.Close()

	for i := 0; i < 100; i++ {
		if 응답, 에러 := 소켓REQ.G질의_응답(lib.P변환형식_기본값, 질의값); 에러 != nil {
			lib.F에러_출력(에러)
			continue
		} else if 응답.G에러() != nil {
			lib.F에러_출력(응답.G에러())
			continue
		} else if !응답.G해석값_단순형(0).(bool) {
			lib.F문자열_출력("접속 처리 실행 실패 후 재시도.")
			continue
		}

		var 접속_성공_여부 = false

		select {
		case 접속_성공_여부 = <-ch로그인:
		case <-ch타임아웃:
		}

		if !접속_성공_여부 {
			lib.F문자열_출력("접속 실패 후 재시도.")
			continue
		}

		break
	}

	F콜백(xing.New콜백_신호(xing.P신호_C32_READY))

	fmt.Println("***************************")
	fmt.Println("*       C32 READY         *")
	fmt.Println("***************************")
	return nil
}

func f초기화_TR전송_제한() {
	코드별_10분당_TR전송_제한 := make(map[string]int)
	코드별_10분당_TR전송_제한[xing.TR현물_기간별_조회] = 200
	코드별_10분당_TR전송_제한[xing.TR_ETF_시간별_추이] = 200
	코드별_10분당_TR전송_제한[xing.TR현물계좌_예수금_주문가능금액_총평가] = 200
	코드별_10분당_TR전송_제한[xing.TR현물계좌_잔고내역] = 200
	코드별_10분당_TR전송_제한[xing.TR현물계좌_주문체결내역] = 200
	코드별_10분당_TR전송_제한[xing.TR주식계좌_기간별_수익률_상세] = 200
	코드별_10분당_TR전송_제한[xing.TR계좌별_신용한도] = 200
	코드별_10분당_TR전송_제한[xing.TR현물계좌_증거금률별_주문가능수량] = 200
	코드별_10분당_TR전송_제한[xing.TR종목별_증시_일정] = 200
	코드별_10분당_TR전송_제한[xing.TR해외_실시간_지수] = 200
	코드별_10분당_TR전송_제한[xing.TR해외_지수_조회] = 200
	코드별_10분당_TR전송_제한[xing.TR증시_주변_자금_추이] = 200

	for TR코드, 초당_제한_횟수 := range 코드별_10분당_TR전송_제한 {
		tr전송_코드별_10분당_제한[TR코드] = lib.New전송_권한_TR코드별(TR코드, 초당_제한_횟수, lib.P10분)
	}

	코드별_초당_TR전송_제한 := make(map[string]int)
	코드별_초당_TR전송_제한[xing.TR현물_정상_주문] = 30
	코드별_초당_TR전송_제한[xing.TR현물_정정_주문] = 30
	코드별_초당_TR전송_제한[xing.TR현물_취소_주문] = 30
	코드별_초당_TR전송_제한[xing.TR현물_호가_조회] = 5
	코드별_초당_TR전송_제한[xing.TR현물_시세_조회] = 5
	코드별_초당_TR전송_제한[xing.TR현물_기간별_조회] = 1
	코드별_초당_TR전송_제한[xing.TR현물_당일_전일_분틱_조회] = 1
	코드별_초당_TR전송_제한[xing.TR_ETF_시간별_추이] = 1
	코드별_초당_TR전송_제한[xing.TR현물_종목_조회] = 2
	코드별_초당_TR전송_제한[xing.TR계좌_거래_내역] = 1
	코드별_초당_TR전송_제한[xing.TR현물계좌_예수금_주문가능금액_총평가] = 1
	코드별_초당_TR전송_제한[xing.TR현물계좌_잔고내역] = 1
	코드별_초당_TR전송_제한[xing.TR현물계좌_주문체결내역] = 1
	코드별_초당_TR전송_제한[xing.TR주식_체결_미체결] = 1
	코드별_초당_TR전송_제한[xing.TR주식_매매일지_수수료_금일] = 2
	코드별_초당_TR전송_제한[xing.TR주식_매매일지_수수료_날짜_지정] = 2
	코드별_초당_TR전송_제한[xing.TR주식_잔고_2] = 1
	코드별_초당_TR전송_제한[xing.TR시간_조회] = 5
	코드별_초당_TR전송_제한[xing.TR주식계좌_기간별_수익률_상세] = 1
	코드별_초당_TR전송_제한[xing.TR계좌별_신용한도] = 1
	코드별_초당_TR전송_제한[xing.TR현물계좌_증거금률별_주문가능수량] = 1
	코드별_초당_TR전송_제한[xing.TR종목별_증시_일정] = 1
	코드별_초당_TR전송_제한[xing.TR해외_실시간_지수] = 1
	코드별_초당_TR전송_제한[xing.TR해외_지수_조회] = 1
	코드별_초당_TR전송_제한[xing.TR증시_주변_자금_추이] = 1

	for TR코드, 초당_제한_횟수 := range 코드별_초당_TR전송_제한 {
		tr전송_코드별_초당_제한[TR코드] = lib.New전송_권한_TR코드별(TR코드, 초당_제한_횟수, lib.P1초)
	}
}

func F리소스_정리() {
	F실시간_정보_모두_해지()
	F로그아웃_및_접속해제()
	F자원_해제()

	lib.F공통_종료_채널_닫기()
	lib.F패닉억제_호출(소켓REP_TR수신.Close)
	lib.F패닉억제_호출(소켓PUB_실시간_정보.Close)
}