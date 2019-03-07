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
(자유 소프트웨어 재단 : Free Software Foundation, In,
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

import "C"
import (
	"github.com/ghts/lib"
	"github.com/ghts/xing"
	"runtime"
	"unsafe"
)

func Go소켓_C함수_호출(ch초기화 chan lib.T신호) (에러 error) {
	if TR_수신_중.G값() {
		ch초기화 <- lib.P신호_초기화
		return nil
	} else if 에러 = TR_수신_중.S값(true); 에러 != nil {
		ch초기화 <- lib.P신호_초기화
		return 에러
	}

	defer TR_수신_중.S값(false)
	defer lib.S예외처리{M에러: &에러}.S실행()

	var 수신값 *lib.S바이트_변환_모음
	var 질의값 lib.I질의값
	var ok bool

	ch도우미_초기화 := make(chan lib.T신호)
	ch도우미_종료 := make(chan lib.T신호)
	ch질의값 := make(chan lib.I질의값, 1)
	ch회신값 := make(chan interface{})
	ch에러 := make(chan error)

	go go소켓_C함수_호출_도우미(ch도우미_초기화, ch도우미_종료, ch질의값, ch회신값, ch에러)
	<-ch도우미_초기화

	ch종료 := lib.F공통_종료_채널()
	ch초기화 <- lib.P신호_초기화 // 초기화 완료.

	for {
		if 수신값, 에러 = 소켓REP_TR수신.G수신(); 에러 != nil {
			select {
			case <-ch종료:
				return nil
			default:
				lib.F에러_출력(에러)
			}

			continue
		}

		lib.F조건부_패닉(수신값.G수량() != 1,
			"잘못된 메시지 길이 : 예상값 1, 실제값 %v.", 수신값.G수량())

		i질의값 := 수신값.S해석기(xing.F바이트_변환값_해석).G해석값_단순형(0)
		if 질의값, ok = i질의값.(lib.I질의값); !ok {
			에러 := lib.New에러with출력("'I질의값'형이 아님 : '%T'", i질의값)
			소켓REP_TR수신.S송신(lib.JSON, 에러)
			continue
		}

		ch질의값 <- 질의값

		select {
		case 회신값 := <-ch회신값:
			소켓REP_TR수신.S송신(수신값.G변환_형식(0), 회신값)
		case 에러 := <-ch에러:
			lib.F에러_출력(에러)
			소켓REP_TR수신.S송신(lib.JSON, 에러)
		case <-ch도우미_종료:
			go go소켓_C함수_호출_도우미(ch도우미_초기화, ch도우미_종료, ch질의값, ch회신값, ch에러)
			<-ch도우미_초기화
		case <-ch종료:
			return nil
		}

		lib.F실행권한_양보()
	}
}

// 단일 스레드에서 API처리와 윈도우 메시지 처리를 하기 위한 중간 매개 함수
func go소켓_C함수_호출_도우미(ch초기화, ch종료 chan lib.T신호,
	ch질의값 chan lib.I질의값, ch회신값 chan interface{}, ch에러 chan error) {
	defer lib.S예외처리{M함수: func() { ch종료 <- lib.P신호_종료 }}.S실행()

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	f초기화_XingAPI() // 모든 API 액세스를 단일 스레드에서 하기 위해서 여기에서 API 초기화를 실행함.
	ch공통_종료 := lib.F공통_종료_채널()
	ch초기화 <- lib.P신호_초기화

	for {
		select {
		case 질의값 := <-ch질의값:
			f질의값_처리(질의값, ch회신값, ch에러)
		case <-ch공통_종료:
			return
		default: // PASS
		}

		lib.F윈도우_메시지_처리()
		lib.F실행권한_양보()
	}
}

func f질의값_처리(질의값 lib.I질의값, ch회신값 chan interface{}, ch에러 chan error) {
	var 에러 error

	defer lib.S예외처리{M에러: &에러, M함수: func() { ch에러 <- 에러 }}.S실행()

	switch 질의값.TR구분() {
	case xing.TR조회, xing.TR주문:
		식별번호 := lib.F확인(f조회_및_주문_질의_처리(질의값)).(int)
		ch회신값 <- 식별번호
	case xing.TR실시간_정보_구독, xing.TR실시간_정보_해지:
		lib.F확인(f실시간_정보_구독_해지_처리(질의값))
		ch회신값 <- lib.P신호_OK
	case xing.TR실시간_정보_일괄_해지:
		lib.F확인(F실시간_정보_모두_해지())
		ch회신값 <- lib.P신호_OK
	case xing.TR접속:
		ch회신값 <- f접속_처리()
	case xing.TR접속됨:
		ch회신값 <- F접속됨()
	case xing.TR서버_이름:
		ch회신값 <- F서버_이름()
	case xing.TR에러_코드:
		ch회신값 <- F에러_코드()
	case xing.TR에러_메시지:
		ch회신값 <- F에러_메시지(질의값.(*lib.S질의값_정수).M정수값)
	case xing.TR계좌_수량:
		ch회신값 <- F계좌_수량()
	case xing.TR계좌번호_모음:
		ch회신값 <- F계좌번호_모음()
	case xing.TR계좌_이름:
		ch회신값 <- F계좌_이름(질의값.(*lib.S질의값_문자열).M문자열)
	case xing.TR계좌_상세명:
		ch회신값 <- F계좌_상세명(질의값.(*lib.S질의값_문자열).M문자열)
	case xing.TR계좌_별명:
		ch회신값 <- F계좌_별명(질의값.(*lib.S질의값_문자열).M문자열)
	case xing.TR코드별_전송_제한:
		ch회신값 <- TR코드별_전송_제한(질의값.(*lib.S질의값_문자열_모음).M문자열_모음)
	case xing.TR소켓_테스트:
		ch회신값 <- lib.P신호_OK
	case xing.TR종료:
		F리소스_정리()
		ch회신값 <- lib.P신호_종료
		F회신_중단_종료()
		Ch메인_종료 <- lib.P신호_종료
	default:
		panic(lib.New에러("예상하지 못한 TR구분값 : '%v'", int(질의값.TR구분())))
	}
}

func f조회_및_주문_질의_처리(질의값 lib.I질의값) (식별번호 int, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 식별번호 = 0 }}.S실행()

	lib.F조건부_패닉(!F접속됨(), "XingAPI에 접속되어 있지 않습니다.")

	var c데이터 unsafe.Pointer
	defer lib.F조건부_실행(c데이터 != nil, F메모리_해제, c데이터)

	var 길이 int
	연속_조회_여부 := false
	연속_조회_키 := ""
	TR코드 := 질의값.(lib.I질의값).TR코드()

	switch TR코드 {
	case xing.TR현물_정상_주문_CSPAT00600:
		c데이터 = unsafe.Pointer(xing.NewCSPAT00600InBlock(질의값.(*xing.S질의값_정상_주문_CSPAT00600)))
		길이 = xing.SizeCSPAT00600InBlock1
	case xing.TR현물_정정_주문_CSPAT00700:
		c데이터 = unsafe.Pointer(xing.NewCSPAT00700InBlock(질의값.(*xing.S질의값_정정_주문_CSPAT00700)))
		길이 = xing.SizeCSPAT00700InBlock1
	case xing.TR현물_취소_주문_CSPAT00800:
		c데이터 = unsafe.Pointer(xing.NewCSPAT00800InBlock(질의값.(*xing.S질의값_취소_주문_CSPAT00800)))
		길이 = xing.SizeCSPAT00800InBlock1
	case xing.TR시간_조회_t0167:
		c데이터 = unsafe.Pointer(C.CString(""))
		defer F메모리_해제(c데이터)
		길이 = 0
	case xing.TR체결_미체결_조회_t0425:
		c데이터 = unsafe.Pointer(xing.NewT0425InBlock(질의값.(*xing.S질의값_체결_미체결_조회_t0425)))
		길이 = xing.SizeT0425InBlock
	case xing.TR현물_호가_조회_t1101:
		c데이터 = unsafe.Pointer(xing.NewT1101InBlock(질의값.(*lib.S질의값_단일_종목)))
		길이 = xing.SizeT1101InBlock
	case xing.TR현물_시세_조회_t1102:
		c데이터 = unsafe.Pointer(xing.NewT1102InBlock(질의값.(*lib.S질의값_단일_종목)))
		길이 = xing.SizeT1102InBlock
	case xing.TR현물_기간별_조회_t1305:
		연속키 := lib.F2문자열_공백제거(질의값.(*xing.S질의값_현물_기간별_조회_t1305).M연속키)
		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(xing.NewT1305InBlock(질의값.(*xing.S질의값_현물_기간별_조회_t1305)))
		길이 = xing.SizeT1305InBlock
	case xing.TR현물_당일_전일_분틱_조회_t1310:
		연속키 := lib.F2문자열_공백제거(질의값.(*xing.S질의값_현물_전일당일_분틱_조회_t1310).M연속키)
		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(xing.NewT1310InBlock(질의값.(*xing.S질의값_현물_전일당일_분틱_조회_t1310)))
		길이 = xing.SizeT1310InBlock
	case xing.TR관리_불성실_투자유의_조회_t1404:
		연속키 := lib.F2문자열_공백제거(질의값.(*xing.S질의값_관리종목_조회_t1404).M연속키)
		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(xing.NewT1404InBlock(질의값.(*xing.S질의값_관리종목_조회_t1404)))
		길이 = xing.SizeT1404InBlock
	case xing.TR투자경고_매매정지_정리매매_조회_t1405:
		연속키 := lib.F2문자열_공백제거(질의값.(*xing.S질의값_투자경고_조회_t1405).M연속키)
		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(xing.NewT1405InBlock(질의값.(*xing.S질의값_투자경고_조회_t1405)))
		길이 = xing.SizeT1405InBlock
	case xing.TR_ETF_시간별_추이_t1902:
		연속키 := lib.F2문자열_공백제거(질의값.(*lib.S질의값_단일종목_연속키).M연속키)
		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}
		c데이터 = unsafe.Pointer(xing.NewT1902InBlock(질의값.(*lib.S질의값_단일종목_연속키)))
		길이 = xing.SizeT1902InBlock
	case xing.TR기업정보_요약_t3320:
		c데이터 = unsafe.Pointer(xing.NewT3320InBlock(질의값.(*lib.S질의값_단일_종목)))
		길이 = xing.SizeT3320InBlock
	case xing.TR재무순위_종합_t3341:
		c데이터 = unsafe.Pointer(xing.NewT3341InBlock(질의값.(*xing.S질의값_재무순위_t3341)))
		길이 = xing.SizeT3341InBlock
	case xing.TR현물_차트_틱_t8411:
		연속키 := lib.F2문자열_공백제거(질의값.(*xing.S질의값_현물_차트_틱_t8411).M연속일자) +
			lib.F2문자열_공백제거(질의값.(*xing.S질의값_현물_차트_틱_t8411).M연속시간)

		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(xing.NewT8411InBlock(질의값.(*xing.S질의값_현물_차트_틱_t8411)))
		길이 = xing.SizeT8411InBlock
	case xing.TR현물_차트_분_t8412:
		연속키 := lib.F2문자열_공백제거(질의값.(*xing.S질의값_현물_차트_분_t8412).M연속일자) +
			lib.F2문자열_공백제거(질의값.(*xing.S질의값_현물_차트_분_t8412).M연속시간)

		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(xing.NewT8412InBlock(질의값.(*xing.S질의값_현물_차트_분_t8412)))
		길이 = xing.SizeT8412InBlock
	case xing.TR현물_차트_일주월_t8413:
		연속키 := lib.F2문자열_공백제거(질의값.(*xing.S질의값_현물_차트_일주월_t8413).M연속일자)

		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(xing.NewT8413InBlock(질의값.(*xing.S질의값_현물_차트_일주월_t8413)))
		길이 = xing.SizeT8413InBlock
	case xing.TR증시_주변_자금_추이_t8428:
		연속키 := lib.F2문자열_공백제거(질의값.(*xing.S질의값_증시주변자금추이_t8428).M연속키)
		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(xing.NewT8428InBlock(질의값.(*xing.S질의값_증시주변자금추이_t8428)))
		길이 = xing.SizeT8428InBlock
	case xing.TR현물_종목_조회_t8436:
		c데이터 = unsafe.Pointer(xing.NewT8436InBlock(질의값.(*lib.S질의값_문자열)))
		길이 = xing.SizeT8436InBlock
	default:
		panic(lib.New에러("미구현 : '%v'", TR코드))
	}

	lib.F조건부_패닉(c데이터 == nil, "c데이터 설정 실패.")

	식별번호 = F질의(TR코드, c데이터, 길이, 연속_조회_여부, 연속_조회_키, lib.P30초)

	if 식별번호 < 0 {
		return 0, lib.New에러("TR호출 실패. 반환된 식별번호가 음수임. '%v' '%v'", 식별번호, TR코드)
	}

	switch TR코드 {
	case xing.TR현물_당일_전일_분틱_조회_t1310: // 전일/당일 구분을 저장해야 함.
		당일전일_구분 := 질의값.(*xing.S질의값_현물_전일당일_분틱_조회_t1310).M당일전일구분
		대기_항목 := New콜백_대기_항목(식별번호, TR코드, 당일전일_구분)
		콜백_대기_저장소.S추가(식별번호, 대기_항목)
	}

	return 식별번호, nil
}

func f실시간_정보_구독_해지_처리(질의값 lib.I질의값) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	var 함수 func(string, string, int) error
	var 전체_종목코드 string
	var 단위_길이 int

	switch 질의값.TR구분() {
	case xing.TR실시간_정보_구독:
		함수 = F실시간_정보_구독
	case xing.TR실시간_정보_해지:
		함수 = F실시간_정보_해지
	}

	switch 변환값 := 질의값.(type) {
	case lib.I종목코드_모음:
		전체_종목코드 = 변환값.G전체_종목코드()
		단위_길이 = len(변환값.G종목코드_모음()[0])
	case lib.I종목코드:
		전체_종목코드 = 변환값.G종목코드()
		단위_길이 = len(변환값.G종목코드())
	default:
		전체_종목코드 = ""
		단위_길이 = 0
	}

	return 함수(질의값.TR코드(), 전체_종목코드, 단위_길이)
}

func f접속_처리() bool {
	defer lib.S예외처리{}.S실행()

	접속_처리_잠금.Lock()
	defer 접속_처리_잠금.Unlock()

	서버_구분 := lib.F조건부_값(lib.F테스트_모드_실행_중(), xing.P서버_모의투자, xing.P서버_실거래).(xing.T서버_구분)

	switch {
	case !F접속(서버_구분):
		return false
	case !F로그인():
		return false
	}

	return true // 로그인 콜백 함수가 실행될 때까지 기다림.
}
