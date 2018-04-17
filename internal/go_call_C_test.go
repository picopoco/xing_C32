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

import (
	"github.com/ghts/lib"
	"github.com/ghts/xing_types"
	"github.com/go-mangos/mangos"

	"testing"
)

var 소켓REP_테스트용_TR수신, 소켓SUB_테스트용_콜백, 소켓SUB_테스트용_실시간정보 mangos.Socket

//P함수_질의 T함수 = iota
//P함수_실시간_정보_구독
//P함수_실시간_정보_해지
//P함수_실시간_정보_모두_해지
//P함수_서버_이름
//P함수_에러_코드
//P함수_에러_메시지
//P함수_TR쿼터
//P함수_계좌_수량
//P함수_계좌_이름
//P함수_계좌_상세명
//P함수_압축_해제

func TestF자료형_크기_비교_확인(t *testing.T) {
	t.Parallel()
	lib.F테스트_에러없음(t, lib.F패닉2에러(자료형_크기_비교_확인))
}

func TestP접속됨(t *testing.T) {
	t.Parallel()
	if !lib.F인터넷에_접속됨() {
		t.SkipNow()
	}

	소켓_질의, 에러 := lib.New소켓_질의(lib.P주소_Xing_C함수_호출, lib.F임의_변환_형식(), lib.P10초)
	lib.F테스트_에러없음(t, 에러)

	응답 := 소켓_질의.S질의(xt.S호출_인수_기본형{M함수: xt.P함수_접속됨}).G응답()
	lib.F테스트_에러없음(t, 응답.G에러())
	lib.F테스트_같음(t, 응답.G수량(), 1)

	var 접속됨 bool
	lib.F테스트_에러없음(t, 응답.G값(0, &접속됨))
	lib.F테스트_같음(t, 접속됨, F접속됨())
}

// 초기화 중 접속이 되므로,개발 과정에서만 사용됨.
//func TestP접속(t *testing.T) {
//	소켓_질의, 에러 := lib.New소켓_질의(lib.P주소_Xing_C함수_호출, lib.F임의_변환_형식(), lib.P30초)
//	lib.F테스트_에러없음(t, 에러)
//
//	질의값 := xt.New호출_인수_기본형(xt.P함수_접속)
//	응답 := 소켓_질의.S질의(질의값).G응답()
//	lib.F테스트_에러없음(t, 응답.G에러())
//	lib.F테스트_같음(t, 응답.G수량(), 1)
//
//	해석값, 에러 := 응답.G해석값(0)
//	lib.F테스트_에러없음(t, 에러)
//
//	switch 해석값.(type) {
//	case bool:
//		var 로그인_됨 bool
//		lib.F테스트_에러없음(t, 응답.G값(0, &로그인_됨))
//		lib.F테스트_참임(t, 로그인_됨)
//	case error:
//		var 에러 error
//		lib.F테스트_에러없음(t, 응답.G값(0, &에러))
//		lib.F에러_출력(에러)
//		t.Fail()
//	}
//}
