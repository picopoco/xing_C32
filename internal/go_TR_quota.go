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
	"time"
)

func go전송_제한_재충전(ch초기화 chan lib.T신호) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	ch도우미_초기화 := make(chan lib.T신호)
	ch도우미_종료 := make(chan lib.T신호)

	go go전송_제한_재충전_도우미(ch도우미_초기화, ch도우미_종료)
	<-ch도우미_초기화

	ch종료 := lib.F공통_종료_채널()
	ch초기화 <- lib.P신호_초기화 // 초기화 완료.

	for {
		select {
		case <-ch도우미_종료:
			go go전송_제한_재충전_도우미(ch도우미_초기화, ch도우미_종료)
			<-ch도우미_초기화
		case <-ch종료:
			return nil
		default:
			lib.F대기(lib.P500밀리초)
		}
	}

}

func go전송_제한_재충전_도우미(ch초기화, ch종료 chan lib.T신호) (에러 error) {
	defer lib.S예외처리{M함수: func() { ch종료 <- lib.P신호_종료 }}.S실행()

	ch공통_종료 := lib.F공통_종료_채널()
	ch초기화 <- lib.P신호_초기화

	매초_충전 := time.NewTicker(lib.P1초 + 3*lib.P10밀리초)
	defer 매초_충전.Stop()

	매3초_충전 := time.NewTicker(lib.P3초 + 9*lib.P10밀리초)
	defer 매3초_충전.Stop()

	for {
		select {
		case <-매초_충전.C:
			for _, 전송_권한 := range tr전송_코드별_초당_제한 {
				전송_권한.S1회_충전()
			}
		case <-매3초_충전.C:
			for _, 전송_권한 := range tr전송_코드별_10분당_제한 {
				전송_권한.S1회_충전()
			}
		case <-ch공통_종료:
			return
		default:
			lib.F대기(lib.P10밀리초)
		}
	}
}
