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
(자유 소프트웨어 재단 : Free Software Foundation, In,
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

import (
	"github.com/ghts/lib"
	"github.com/ghts/types_xing"
	"github.com/go-mangos/mangos"

	"sync"
	"unsafe"
)

var TR소켓_중계_중 = lib.New안전한_bool(false)
var ch도우미_종료 chan error
var 소켓REP mangos.Socket = nil

func Go루틴_소켓_C함수_호출(ch초기화 chan lib.T신호) (에러 error) {
	if TR소켓_중계_중.G값() {
		ch초기화 <- lib.P신호_초기화
		return nil
	} else if 에러 = TR소켓_중계_중.S값(true); 에러 != nil {
		ch초기화 <- lib.P신호_초기화
		return 에러
	}

	defer TR소켓_중계_중.S값(false)
	defer lib.F에러패닉_처리(lib.S에러패닉_처리{M에러: &에러})

	// 병렬처리를 위해서 raw모드를 사용함. raw모드 사용법은 go-mangos 예제를 참고.
	소켓REP, 에러 = lib.New소켓REP_raw(lib.P주소_Xing_C함수_호출) // 종료할 때 닫음.
	lib.F에러체크(에러)

	도우미_수량 := 10

	if lib.F테스트_모드_실행_중() {
		도우미_수량 = 1
	}

	ch도우미_초기화 := make(chan lib.T신호, 도우미_수량)
	ch도우미_종료 = make(chan error, 도우미_수량)

	for i := 0; i < 도우미_수량; i++ {
		go c함수_호출_도우미(ch도우미_초기화)
	}

	for i := 0; i < 도우미_수량; i++ {
		신호 := <-ch도우미_초기화
		lib.F조건부_패닉(신호 != lib.P신호_초기화, "%v번째 c함수_호출_도우미 초기화 실패.", i)
	}

	ch종료 := lib.F공통_종료_채널()
	ch초기화 <- lib.P신호_초기화 // 초기화 완료.

	for {
		select {
		default:
			lib.F윈도우_메시지_처리()
			lib.F실행권한_양보()
		case <-ch종료:
			return nil
		case 에러 = <-ch도우미_종료:
			select {
			case <-ch종료:
				return nil
			default:
				lib.F에러_출력(에러)

				go c함수_호출_도우미(ch도우미_초기화)

				신호 := <-ch도우미_초기화
				lib.F조건부_패닉(신호 != lib.P신호_초기화, "c함수_호출_도우미 초기화 실패.")
			}
		}
	}
}

// 소켓으로 받은 c함수 호출 요청을 실제로 처리하는 도우미 함수
// 여러 개의 도우미 함수 루틴이 동시에 존재하므로, 처리시간으로 인한 지연에 신경쓰지 않아도 됨.
// 에러나 패닉이 발생하지 않는 한 종료할 필요없음.
func c함수_호출_도우미(ch초기화 chan<- lib.T신호) (에러 error) {
	var raw메시지 *mangos.Message // 최대한 재활용 해야 성능 문제를 걱정할 필요가 없어진다.

	defer func() { ch도우미_종료 <- 에러 }()
	defer lib.F에러패닉_처리(lib.S에러패닉_처리{
		M에러: &에러,
		M함수with패닉내역: func(r interface{}) {
			lib.F체크포인트(r)

			if raw메시지 != nil {
				lib.F체크포인트(r)

				if 메시지, 에러 := lib.New소켓_메시지_에러(r); 에러 == nil {
					메시지.S소켓_회신(소켓REP, lib.P10초, raw메시지)
				}
			}
		}})

	ch종료 := lib.F공통_종료_채널()
	ch초기화 <- lib.P신호_초기화

	var 변환_형식 lib.T변환
	var 함수 xing.T함수
	var 회신_메시지 lib.I소켓_메시지

	for {
		select {
		case <-ch종료:
			return nil
		default:
			raw메시지, 에러 = 소켓REP.RecvMsg()
			lib.F에러체크(에러)

			수신_메시지 := lib.New소켓_메시지from바이트_모음(raw메시지.Body)
			lib.F에러체크(수신_메시지.G에러())

			lib.F조건부_패닉(수신_메시지.G길이() != 1,
				"잘못된 메시지 길이. 예상값 1, 실제값 %v. %v",
				수신_메시지.G길이(), 수신_메시지)

			질의값, 에러 := 수신_메시지.G해석값(0, xing.F바이트_변환값_해석)
			lib.F에러체크(에러)

			변환_형식 = 수신_메시지.G변환_형식(0)
			함수 = 질의값.(xing.I호출_인수).G함수()

			switch 함수 {
			case xing.P함수_접속:
				go f접속_처리(변환_형식, raw메시지, ch종료)
				continue
			case xing.P함수_질의:
				lib.F체크포인트(회신_메시지)

				TR질의값, 에러 := 질의값.(*xing.S호출_인수_질의).M질의값.G해석값(xing.F바이트_변환값_해석)
				lib.F에러체크(에러)

				식별번호, 에러 := f조회_및_주문_질의_처리(TR질의값.(lib.I질의값))

				if 에러 != nil {
					회신_메시지, 에러 = lib.New소켓_메시지_에러(에러)
				} else {
					회신_메시지, 에러 = lib.New소켓_메시지(변환_형식, 식별번호)
				}
			case xing.P함수_접속됨:
				회신_메시지, 에러 = lib.New소켓_메시지(변환_형식, F접속됨())
			case xing.P함수_로그아웃_및_접속해제:
				회신_메시지, 에러 = lib.New소켓_메시지_에러(F로그아웃_및_접속해제())
			case xing.P함수_실시간_정보_구독:
				s := 질의값.(xing.S호출_인수_실시간_정보)
				에러 := F실시간_정보_구독(s.TR코드, s.M전체_종목코드, s.M단위_길이)
				회신_메시지, 에러 = lib.New소켓_메시지_에러(에러)
			case xing.P함수_실시간_정보_해지:
				s := 질의값.(xing.S호출_인수_실시간_정보)
				에러 := F실시간_정보_해지(s.TR코드, s.M전체_종목코드, s.M단위_길이)
				회신_메시지, 에러 = lib.New소켓_메시지_에러(에러)
			case xing.P함수_실시간_정보_모두_해지:
				회신_메시지, 에러 = lib.New소켓_메시지_에러(F실시간_정보_모두_해지())
			case xing.P함수_서버_이름:
				회신_메시지, 에러 = lib.New소켓_메시지(변환_형식, F서버_이름())
			case xing.P함수_에러_코드:
				회신_메시지, 에러 = lib.New소켓_메시지(변환_형식, F에러_코드())
			case xing.P함수_에러_메시지:
				회신_메시지, 에러 = lib.New소켓_메시지(변환_형식, F에러_메시지(질의값.(xing.S호출_인수_정수값).M정수값))
			case xing.P함수_TR쿼터:
				회신_메시지, 에러 = lib.New소켓_메시지(변환_형식, F초당_TR쿼터(질의값.(xing.S호출_인수_문자열).M문자열))
			case xing.P함수_계좌_수량:
				회신_메시지, 에러 = lib.New소켓_메시지(변환_형식, F계좌_수량())
			case xing.P함수_계좌_이름:
				회신_메시지, 에러 = lib.New소켓_메시지(변환_형식, F계좌_이름(질의값.(xing.S호출_인수_문자열).M문자열))
			case xing.P함수_계좌_상세명:
				회신_메시지, 에러 = lib.New소켓_메시지(변환_형식, F계좌_상세명(질의값.(xing.S호출_인수_문자열).M문자열))
			case xing.P함수_압축_해제:
				s := 질의값.(xing.S호출_인수_압축_해제)
				회신_메시지, 에러 = lib.New소켓_메시지(변환_형식, F압축_해제(unsafe.Pointer(&s.M바이트_모음), s.M길이))
			case xing.P함수_종료:
				에러 = F로그아웃_및_접속해제()
				f소켓_질의_회신(lib.New소켓_메시지_내용없음(), raw메시지, ch종료)
				lib.F공통_종료_채널_닫기()
				return 에러
			default:
				lib.F패닉("예상하지 못한 함수 : '%v'", 함수)
			}

			if 에러 != nil {
				회신_메시지, _ = lib.New소켓_메시지_에러(에러)
			}

			f소켓_질의_회신(회신_메시지, raw메시지, ch종료)
		}
	}

	return nil
}

func f소켓_질의_회신(회신_메시지 lib.I소켓_메시지, raw메시지 *mangos.Message, ch종료 chan lib.T신호) {
	if lib.F포맷된_문자열("%v", 회신_메시지) == "<nil>" {
		lib.F에러_출력("nil 회신 메시지")
	}

	select {
	case <-ch종료: // f종료TR_처리()에서 회신 및 Free() 처리함.
	default:
		lib.F에러체크(회신_메시지.S소켓_회신(소켓REP, lib.P30초, raw메시지))

		raw메시지.Free() // GC부담을 덜어서 성능 향상을 꾀함.				}
	}
}

var 접속_처리_잠금 sync.Mutex
var ch접속_처리 = make(chan bool, 1)

func f접속_처리(변환_형식 lib.T변환, raw메시지 *mangos.Message, ch종료 chan lib.T신호) {
	접속_처리_잠금.Lock()
	defer 접속_처리_잠금.Unlock()

	defer lib.F에러패닉_처리(lib.S에러패닉_처리{
		M함수with패닉내역: func(r interface{}) { lib.F에러_출력(r) }})

	// '접속_처리_잠금'에 의해서 접속 처리는 1개씩만 처리되므로, 'ch접속_처리'는 비어있어야 함.
	lib.F조건부_패닉(len(ch접속_처리) > 0, "'ch접속_처리'는 비어있어야 함.")

	for len(ch접속_처리) > 0 {
		lib.F체크포인트("예상과 다른 상황 : 'ch접속_처리'가 비어있지 않음.")
		<-ch접속_처리
	}

	서버_구분 := lib.F조건부_값(lib.F테스트_모드_실행_중(), xing.P서버_모의투자, xing.P서버_실거래).(xing.T서버_구분)

	var 접속_여부 bool

	switch {
	case F접속됨():
		접속_여부 = true
	case !F접속(서버_구분):
		접속_여부 = false
	case !F로그인():
		접속_여부 = false
	default:
		접속_여부 = <-ch접속_처리 // 로그인 콜백 함수가 실행될 때까지 기다림.
	}

	회신_메시지, 에러 := lib.New소켓_메시지(변환_형식, 접속_여부)
	lib.F에러체크(에러)

	f소켓_질의_회신(회신_메시지, raw메시지, ch종료)
}

func f조회_및_주문_질의_처리(질의값 lib.I질의값) (식별번호 int, 에러 error) {
	defer lib.F에러패닉_처리(lib.S에러패닉_처리{
		M에러: &에러,
		M함수: func() { 식별번호 = 0 }})

	lib.F조건부_패닉(!F접속됨(), "XingAPI에 접속되어 있지 않습니다.")
	f질의값_종목코드_검사(질의값)

	var c데이터 unsafe.Pointer
	defer lib.F조건부_실행(c데이터 != nil, F메모리_해제, c데이터)

	var 길이 int
	연속_조회_여부 := false
	연속_조회_키 := ""
	TR코드 := 질의값.(lib.I질의값).G_TR코드()

	fTR전송권한획득(TR코드)

	switch TR코드 {
	case xing.TR현물_정상주문:
		c데이터 = unsafe.Pointer(NewCSPAT00600InBlock(질의값.(*xing.S질의값_정상주문)))
		길이 = int(unsafe.Sizeof(CSPAT00600InBlock1{}))
	case xing.TR현물_정정주문:
		c데이터 = unsafe.Pointer(NewCSPAT00700InBlock(질의값.(*xing.S질의값_정정주문)))
		길이 = int(unsafe.Sizeof(CSPAT00700InBlock1{}))
	case xing.TR현물_취소주문:
		c데이터 = unsafe.Pointer(NewCSPAT00800InBlock(질의값.(*xing.S질의값_취소주문)))
		길이 = int(unsafe.Sizeof(CSPAT00800InBlock1{}))
	case xing.TR시간_조회:
		c데이터 = unsafe.Pointer(C문자열(""))
		길이 = 0
	case xing.TR현물_호가_조회:
		g := new(T1101InBlock)
		lib.F바이트_복사_문자열(g.Shcode[:], 질의값.(*lib.S질의값_단일종목).M종목코드)
		c데이터 = unsafe.Pointer(g)
		길이 = int(unsafe.Sizeof(T1101InBlock{}))
	case xing.TR현물_시세_조회:
		g := new(T1102InBlock)
		lib.F바이트_복사_문자열(g.Shcode[:], 질의값.(*lib.S질의값_단일종목).M종목코드)
		c데이터 = unsafe.Pointer(g)
		길이 = int(unsafe.Sizeof(T1102InBlock{}))
	case xing.TR현물_기간별_조회:
		연속키 := lib.F2문자열_공백제거(질의값.(*xing.S질의값_기간별_주가_조회).M연속키)
		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(NewT1305InBlock(질의값.(*xing.S질의값_기간별_주가_조회)))
		길이 = int(unsafe.Sizeof(T1310InBlock{}))
	case xing.TR현물_당일_전일_분틱_조회:
		연속키 := lib.F2문자열_공백제거(질의값.(*xing.S질의값_현물_전일당일_분틱_조회).M연속키)
		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(NewT1310InBlock(질의값.(*xing.S질의값_현물_전일당일_분틱_조회)))
		길이 = int(unsafe.Sizeof(T1310InBlock{}))
	case xing.TR_ETF_시간별_추이:
		연속키 := lib.F2문자열_공백제거(질의값.(*xing.S질의값_단일종목_연속키).M연속키)
		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(NewT1902InBlock(질의값.(*xing.S질의값_단일종목_연속키)))
		길이 = int(unsafe.Sizeof(T1902InBlock{}))
	case xing.TR증시_주변_자금_추이:
		연속키 := lib.F2문자열_공백제거(질의값.(*xing.S질의값_증시주변자금추이).M연속키)
		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(NewT8428InBlock(질의값.(*xing.S질의값_증시주변자금추이)))
		길이 = int(unsafe.Sizeof(T8428InBlock{}))
	case xing.TR현물_종목_조회:
		c데이터 = unsafe.Pointer(NewT8436InBlock(질의값.(*lib.S질의값_문자열)))
		길이 = int(unsafe.Sizeof(T8436InBlock{}))
	case xing.TR계좌_번호:
		// 자체적으로 생성한 편의성을 위한 가상 TR임.
		// 여기서 응답까지 다 처리함.

		계좌_수량 := F계좌_수량()
		계좌번호_모음 := make([]string, 계좌_수량)

		for i := 0; i < 계좌_수량; i++ {
			계좌번호_모음[i], 에러 = F계좌_번호(i)
			lib.F에러체크(에러)
		}

		변환값, 에러 := lib.New바이트_변환_매개체(lib.JSON, 계좌번호_모음)
		lib.F에러체크(에러)

		바이트_모음, 에러 := 변환값.G바이트_모음()
		lib.F에러체크(에러)

		응답값 := new(xing.S콜백_TR_데이터)
		응답값.M콜백 = xing.P콜백_데이터
		응답값.TR코드 = xing.TR계좌_번호
		응답값.M데이터 = 바이트_모음

		f콜백(응답값)
		return 999999, nil
	case xing.TR계좌_거래_내역,
		xing.TR현물계좌_예수금_주문가능금액_총평가,
		xing.TR현물계좌_잔고내역,
		xing.TR현물계좌_주문체결내역,
		xing.TR주식_체결_미체결,
		xing.TR주식_매매일지_수수료_금일,
		xing.TR주식_매매일지_수수료_날짜_지정,
		xing.TR주식_잔고_2,
		xing.TR주식계좌_기간별_수익률_상세,
		xing.TR계좌별_신용한도,
		xing.TR현물계좌_증거금률별_주문가능수량,
		xing.TR종목별_증시_일정,
		xing.TR해외_실시간_지수,
		xing.TR해외_지수_조회:
		fallthrough
	default:
		lib.F패닉("미구현")
	}

	lib.F조건부_패닉(c데이터 == nil, "c데이터 설정 실패.")

	식별번호 = F질의(TR코드, c데이터, 길이, 연속_조회_여부, 연속_조회_키, lib.P30초)

	if 식별번호 < 0 {
		return 0, lib.New에러("TR호출 실패. 반환된 식별번호가 음수임. %v", 식별번호)
	}

	return 식별번호, nil
}
