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

import (
	"github.com/ghts/lib"

	"testing"
)

func TestMain(m *testing.M) {
	defer lib.S예외처리{M함수: func() { lib.F체크포인트() }}.S실행()

	f테스트_준비()
	defer f테스트_정리()

	m.Run()
}

func f테스트_준비() {
	defer lib.S예외처리{}.S실행()

	lib.F테스트_모드_시작()

	ch초기화 := make(chan lib.T신호)
	go go테스트용_TR콜백_수신(ch초기화)
	<-ch초기화

	F초기화()
	f초기화_작동_확인()
}

func f테스트_정리() {
	F리소스_정리()
	lib.F테스트_모드_종료()
}

func go테스트용_TR콜백_수신(ch초기화 chan lib.T신호) {
	소켓REP_TR콜백 := lib.NewNano소켓REP_단순형(lib.P주소_Xing_C함수_콜백)

	ch초기화 <- lib.P신호_초기화

	for {
		값, 에러 := 소켓REP_TR콜백.G수신()
		if 에러 != nil {
			lib.F에러_출력(에러)
			continue
		}

		소켓REP_TR콜백.S송신(값.G변환_형식(0), lib.P신호_OK)
	}
}

func f초기화_작동_확인() {
	소켓REQ := lib.NewNano소켓REQ_단순형(lib.P주소_Xing_C함수_호출, lib.P10초)
	defer 소켓REQ.Close()

	질의값 := lib.New질의값_기본형(lib.TR접속됨, "")

	if 응답 := 소켓REQ.G질의_응답_검사(lib.P변환형식_기본값, 질의값); 응답.G에러() != nil {
		lib.F에러_출력(응답.G에러())
		f초기화_작동_확인() // 재귀 호출로 재시도
	} else if 접속됨, ok := 응답.G해석값_단순형(0).(bool); !ok {
		panic(lib.New에러("예상하지 못한 자료형 : '%T'", 응답.G해석값_단순형(0)))
	} else if !접속됨 {
		panic(lib.New에러("이 시점에 접속되어 있어야 함."))
	}
}