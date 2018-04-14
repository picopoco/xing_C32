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

	"testing"
)

func TestF함수_존재함(t *testing.T) {
	t.Parallel()

	l.F테스트_참임(t, f함수_존재함(FuncConnect))
	l.F테스트_참임(t, f함수_존재함(FuncIsConnected))
	l.F테스트_참임(t, f함수_존재함(FuncDisconnect))
	l.F테스트_참임(t, f함수_존재함(FuncLogin))
	l.F테스트_참임(t, f함수_존재함(FuncRequest))
	l.F테스트_참임(t, f함수_존재함(FuncReleaseRequestData))
	l.F테스트_참임(t, f함수_존재함(FuncReleaseMessageData))
	l.F테스트_참임(t, f함수_존재함(FuncAdviseRealData))
	l.F테스트_참임(t, f함수_존재함(FuncUnadviseRealData))
	l.F테스트_참임(t, f함수_존재함(FuncUnadviseWindow))
	l.F테스트_참임(t, f함수_존재함(FuncGetAccountListCount))
	l.F테스트_참임(t, f함수_존재함(FuncGetAccountList))
	l.F테스트_참임(t, f함수_존재함(FuncGetAccountName))
	l.F테스트_참임(t, f함수_존재함(FuncGetAcctDetailName))
	l.F테스트_참임(t, f함수_존재함(FuncGetAcctNickName))
	l.F테스트_참임(t, f함수_존재함(FuncGetServerName))
	l.F테스트_참임(t, f함수_존재함(FuncGetLastError))
	l.F테스트_참임(t, f함수_존재함(FuncGetErrorMessage))
	l.F테스트_참임(t, f함수_존재함(FuncGetTRCountPerSec))
	l.F테스트_참임(t, f함수_존재함(FuncRequestService))
	l.F테스트_참임(t, f함수_존재함(FuncRemoveService))
	l.F테스트_참임(t, f함수_존재함(FuncRequestLinkToHTS))
	l.F테스트_참임(t, f함수_존재함(FuncAdviseLinkFromHTS))
	l.F테스트_참임(t, f함수_존재함(FuncUnadviseLinkFromHTS))
	l.F테스트_참임(t, f함수_존재함(FuncDecompress))
}
