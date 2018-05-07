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
	"github.com/ghts/xing_types"

	"fmt"
	"runtime"
	"time"
	"unsafe"
)

func Go소켓_C함수_호출(ch초기화 chan lib.T신호) (에러 error) {
	if TR소켓_중계_중.G값() {
		ch초기화 <- lib.P신호_초기화
		return nil
	} else if 에러 = TR소켓_중계_중.S값(true); 에러 != nil {
		ch초기화 <- lib.P신호_초기화
		return 에러
	}

	defer TR소켓_중계_중.S값(false)
	defer lib.S에러패닉_처리기{M에러_포인터: &에러}.S실행()

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

	lib.F메모("XingAPI는 싱글 스레드 전용이므로, 다중처리가 불가함.")
	lib.F메모("XingAPI 인스턴스를 여러 개 띄우서 다중처리 하는 방안을 알아볼 것.")

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

		i질의값 := 수신값.S해석기(xt.F바이트_변환값_해석).G해석값_단순형(0)
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
	defer lib.S에러패닉_처리기{M함수: func() { ch종료 <- lib.P신호_종료 }}.S실행()

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

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
	defer lib.S에러패닉_처리기{M함수with내역: func(r interface{}) { ch에러 <- lib.New에러(r) }}.S실행()

	switch 질의값.G_TR구분() {
	case xt.TR조회, xt.TR주문:
		식별번호 := 에러체크(f조회_및_주문_질의_처리(질의값)).(int)
		ch회신값 <- 식별번호
	case xt.TR실시간_정보_구독, xt.TR실시간_정보_해지:
		에러체크(f실시간_정보_구독_해지_처리(질의값))
		ch회신값 <- lib.P신호_OK
	case xt.TR실시간_정보_일괄_해지:
		에러체크(F실시간_정보_모두_해지())
		ch회신값 <- lib.P신호_OK
	case xt.TR접속:
		접속_처리_결과 := f접속_처리()
		ch회신값 <- 접속_처리_결과
	case xt.TR접속됨:
		ch회신값 <- F접속됨()
	case xt.TR서버_이름:
		ch회신값 <- F서버_이름()
	case xt.TR에러_코드:
		ch회신값 <- F에러_코드()
	case xt.TR에러_메시지:
		ch회신값 <- F에러_메시지(질의값.(*lib.S질의값_정수).M정수값)
	case xt.TR코드별_쿼터:
		ch회신값 <- F초당_TR쿼터(질의값.(*lib.S질의값_문자열).M문자열)
	case xt.TR계좌_수량:
		ch회신값 <- F계좌_수량()
	case xt.TR계좌_번호:
		ch회신값 <- F계좌_번호(질의값.(*lib.S질의값_정수).M정수값)
	case xt.TR계좌_이름:
		ch회신값 <- F계좌_이름(질의값.(*lib.S질의값_문자열).M문자열)
	case xt.TR계좌_상세명:
		ch회신값 <- F계좌_상세명(질의값.(*lib.S질의값_문자열).M문자열)
	case xt.TR압축_해제:
		바이트_모음 := 질의값.(*lib.S질의값_바이트_변환).M바이트_변환.G바이트_모음_단순형()
		ch회신값 <- F압축_해제(unsafe.Pointer(&바이트_모음), len(바이트_모음))
	case xt.TR소켓_테스트:
		ch회신값 <- lib.P신호_OK
	case xt.TR전일_당일:
		f전일_당일_설정(질의값)
		ch회신값 <- lib.P신호_OK
	case xt.TR종료:
		select {
		case Ch메인_종료 <- lib.P신호_종료:
		default:
		}

		ch회신값 <- lib.P신호_종료
	default:
		panic(lib.New에러("예상하지 못한 TR구분값 : '%v'", int(질의값.G_TR구분())))
	}
}

func f전일_당일_설정(질의값 lib.I질의값) (에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러}.S실행()

	전일_당일_설정_잠금.Lock()
	defer 전일_당일_설정_잠금.Unlock()

	바이트_변환_모음 := 질의값.(*lib.S질의값_바이트_변환_모음)
	전일_당일_설정_일자 = lib.New안전한_시각(바이트_변환_모음.M바이트_변환_모음.G해석값_단순형(0).(time.Time))
	전일 = lib.New안전한_시각(바이트_변환_모음.M바이트_변환_모음.G해석값_단순형(1).(time.Time))
	당일 = lib.New안전한_시각(바이트_변환_모음.M바이트_변환_모음.G해석값_단순형(2).(time.Time))

	fmt.Println("**************************")
	fmt.Println("* C32 전일 당일 설정 완료 *")
	fmt.Println("**************************")

	return nil
}

func f실시간_정보_구독_해지_처리(질의값 lib.I질의값) (에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러}.S실행()

	var 구독_해지_함수 func(string, string, int) error

	switch 질의값.G_TR구분() {
	case xt.TR실시간_정보_구독:
		구독_해지_함수 = F실시간_정보_구독
	case xt.TR실시간_정보_해지:
		구독_해지_함수 = F실시간_정보_해지
	}

	switch 질의값.G_TR코드() {
	case xt.RT현물주문_접수, xt.RT현물주문_체결, xt.RT현물주문_정정,
		xt.RT현물주문_거부, xt.RT현물주문_취소, xt.RT장_운영정보:
		// 단순 TR. 종목코드 및 단위길이가 필요없음.
		return 구독_해지_함수(질의값.G_TR코드(), "", 0)
	case xt.RT코스피_호가_잔량, xt.RT코스피_시간외_호가_잔량,
		xt.RT코스피_체결, xt.RT코스피_예상_체결,
		xt.RT코스피_ETF_NAV, xt.RT업종별_투자자별_매매_현황,
		xt.RT주식_VI발동해제, xt.RT시간외_단일가VI발동해제: // 복수 종목
		전체_종목코드 := 질의값.(*lib.S질의값_복수종목).G전체_종목코드()
		단위_길이 := len(질의값.(*lib.S질의값_복수종목).M종목코드_모음[0])
		return 구독_해지_함수(질의값.G_TR코드(), 전체_종목코드, 단위_길이)
	default:
		return lib.New에러("예상하지 못한 RT코드 : '%v'", 질의값.G_TR코드())
	}
}

func f조회_및_주문_질의_처리(질의값 lib.I질의값) (식별번호 int, 에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러, M함수: func() { 식별번호 = 0 }}.S실행()

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
	case xt.TR현물_정상주문:
		c데이터 = unsafe.Pointer(NewCSPAT00600InBlock(질의값.(*xt.S질의값_정상주문)))
		길이 = int(unsafe.Sizeof(CSPAT00600InBlock1{}))
	case xt.TR현물_정정주문:
		c데이터 = unsafe.Pointer(NewCSPAT00700InBlock(질의값.(*xt.S질의값_정정주문)))
		길이 = int(unsafe.Sizeof(CSPAT00700InBlock1{}))
	case xt.TR현물_취소주문:
		c데이터 = unsafe.Pointer(NewCSPAT00800InBlock(질의값.(*xt.S질의값_취소주문)))
		길이 = int(unsafe.Sizeof(CSPAT00800InBlock1{}))
	case xt.TR시간_조회:
		c데이터 = unsafe.Pointer(C문자열(""))
		길이 = 0
	case xt.TR현물_호가_조회:
		g := new(T1101InBlock)
		lib.F바이트_복사_문자열(g.Shcode[:], 질의값.(*lib.S질의값_단일종목).M종목코드)
		c데이터 = unsafe.Pointer(g)
		길이 = int(unsafe.Sizeof(T1101InBlock{}))
	case xt.TR현물_시세_조회:
		g := new(T1102InBlock)
		lib.F바이트_복사_문자열(g.Shcode[:], 질의값.(*lib.S질의값_단일종목).M종목코드)
		c데이터 = unsafe.Pointer(g)
		길이 = int(unsafe.Sizeof(T1102InBlock{}))
	case xt.TR현물_기간별_조회:
		연속키 := lib.F2문자열_공백제거(질의값.(*xt.S질의값_현물_기간별_조회).M연속키)
		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(NewT1305InBlock(질의값.(*xt.S질의값_현물_기간별_조회)))
		길이 = int(unsafe.Sizeof(T1310InBlock{}))
	case xt.TR현물_당일_전일_분틱_조회:
		연속키 := lib.F2문자열_공백제거(질의값.(*xt.S질의값_현물_전일당일_분틱_조회).M연속키)
		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(NewT1310InBlock(질의값.(*xt.S질의값_현물_전일당일_분틱_조회)))
		길이 = int(unsafe.Sizeof(T1310InBlock{}))
	case xt.TR_ETF_시간별_추이:
		연속키 := lib.F2문자열_공백제거(질의값.(*xt.S질의값_단일종목_연속키).M연속키)
		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(NewT1902InBlock(질의값.(*xt.S질의값_단일종목_연속키)))
		길이 = int(unsafe.Sizeof(T1902InBlock{}))
	case xt.TR증시_주변_자금_추이:
		연속키 := lib.F2문자열_공백제거(질의값.(*xt.S질의값_증시주변자금추이).M연속키)
		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(NewT8428InBlock(질의값.(*xt.S질의값_증시주변자금추이)))
		길이 = int(unsafe.Sizeof(T8428InBlock{}))
	case xt.TR현물_종목_조회:
		c데이터 = unsafe.Pointer(NewT8436InBlock(질의값.(*lib.S질의값_문자열)))
		길이 = int(unsafe.Sizeof(T8436InBlock{}))
	case xt.TR계좌_거래_내역,
		xt.TR현물계좌_예수금_주문가능금액_총평가,
		xt.TR현물계좌_잔고내역,
		xt.TR현물계좌_주문체결내역,
		xt.TR주식_체결_미체결,
		xt.TR주식_매매일지_수수료_금일,
		xt.TR주식_매매일지_수수료_날짜_지정,
		xt.TR주식_잔고_2,
		xt.TR주식계좌_기간별_수익률_상세,
		xt.TR계좌별_신용한도,
		xt.TR현물계좌_증거금률별_주문가능수량,
		xt.TR종목별_증시_일정,
		xt.TR해외_실시간_지수,
		xt.TR해외_지수_조회:
		fallthrough
	default:
		panic("미구현")
	}

	lib.F조건부_패닉(c데이터 == nil, "c데이터 설정 실패.")

	식별번호 = F질의(TR코드, c데이터, 길이, 연속_조회_여부, 연속_조회_키, lib.P30초)

	if 식별번호 < 0 {
		return 0, lib.New에러("TR호출 실패. 반환된 식별번호가 음수임. %v", 식별번호)
	}

	switch TR코드 {
	case xt.TR현물_당일_전일_분틱_조회: // 전일/당일 구분을 저장해야 함.
		당일전일_구분 := 질의값.(*xt.S질의값_현물_전일당일_분틱_조회).M당일전일구분
		대기_항목 := New콜백_대기_항목(식별번호, TR코드, 당일전일_구분)
		콜백_대기_저장소.S추가(식별번호, 대기_항목)
	}

	return 식별번호, nil
}

func f접속_처리() bool {
	defer lib.S에러패닉_처리기{}.S실행()

	접속_처리_잠금.Lock()
	defer 접속_처리_잠금.Unlock()

	서버_구분 := lib.F조건부_값(lib.F테스트_모드_실행_중(), xt.P서버_모의투자, xt.P서버_실거래).(xt.T서버_구분)

	switch {
	case !F접속(서버_구분):
		lib.F체크포인트()
		return false
	case !F로그인():
		lib.F체크포인트()
		return false
	}

	return true // 로그인 콜백 함수가 실행될 때까지 기다림.
}
