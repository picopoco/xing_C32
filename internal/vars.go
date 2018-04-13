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
	"github.com/go-mangos/mangos"
	"time"
)

// 전역 변수는 항상 동시 액세스로 인한 오류의 위험이 있어서 한 군데 몰아서 관리함.

// 다중 사용에 안전한 값들.
var (
	소켓REP_TR수신 mangos.Socket = nil
	소켓PUB_콜백   mangos.Socket = nil // PUB소켓은 수신 기능이 없으며, 지연없는 'non-blocking'방식으로 동작.

	ch호출_도우미_종료 chan error

	메시지_저장소 = New메시지_저장소()
)

// 초기화 이후에는 사실상 읽기 전용이어서, 다중 사용에 문제가 없는 값들.
var (
	tr전송_코드별_10분당_제한 = make(map[string]lib.I전송_권한_TR코드별)
	tr전송_코드별_초당_제한   = make(map[string]lib.I전송_권한_TR코드별)

	설정화일_경로 = lib.F_GOPATH() + `/src/github.com/ghts/api_bridge_xing/internal/` + "config.ini"

	전일_금일_초기값            = time.Time{}
	영업일_기준_전일, 영업일_기준_당일 time.Time
)
