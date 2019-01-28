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
	"github.com/ghts/xing"

	"unsafe"
)

func tr데이터_해석(tr *xing.TR_DATA) (값 interface{}, 에러 error) {
	lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	const 큰수 = 1<<20
	TR코드 := lib.F2문자열(tr.TrCode)
	데이터_길이 := int(tr.DataLength)

	switch TR코드 {
	case xing.TR현물_정상_주문:
		switch 데이터_길이 {
		case xing.SizeCSPAT00600OutBlock1:
			return (*xing.CSPAT00600OutBlock1)(unsafe.Pointer(tr.Data)), nil
		case xing.SizeCSPAT00600OutBlock2:
			return (*xing.CSPAT00600OutBlock2)(unsafe.Pointer(tr.Data)), nil
		}
	case xing.TR현물_정정_주문:
		switch 데이터_길이 {
		case xing.SizeCSPAT00700OutBlock1:
			return (*xing.CSPAT00700OutBlock1)(unsafe.Pointer(tr.Data)), nil
		case xing.SizeCSPAT00700OutBlock2:
			return (*xing.CSPAT00700OutBlock2)(unsafe.Pointer(tr.Data)), nil
		}
	case xing.TR현물_취소_주문:
		switch 데이터_길이 {
		case xing.SizeCSPAT00800OutBlock1:
			return (*xing.CSPAT00800OutBlock1)(unsafe.Pointer(tr.Data)), nil
		case xing.SizeCSPAT00800OutBlock2:
			return (*xing.CSPAT00800OutBlock2)(unsafe.Pointer(tr.Data)), nil
		}
	case xing.TR시간_조회:
		g := (*xing.T0167OutBlock)(unsafe.Pointer(tr.Data))
		날짜_문자열 := lib.F2문자열(g.Date)
		시간_문자열 := lib.F2문자열(g.Time)

		return lib.F2포맷된_시각("20060102150405.99999999", 날짜_문자열+시간_문자열[:6]+"."+시간_문자열[7:])
	case xing.TR현물_호가_조회:
		return (*xing.T1101OutBlock)(unsafe.Pointer(tr.Data)), nil
	case xing.TR현물_시세_조회:
		return (*xing.T1102OutBlock)(unsafe.Pointer(tr.Data)), nil
	case xing.TR현물_기간별_조회:
		if 데이터_길이 == xing.SizeT1305OutBlock {
			return (*xing.T1305OutBlock)(unsafe.Pointer(tr.Data)), nil
		}

		수량 := int(tr.DataLength) / xing.SizeT1305OutBlock1
		g_모음 := (*[큰수]xing.T1305OutBlock1)(unsafe.Pointer(tr.Data))[:수량:수량]
		return g_모음, nil
	case xing.TR현물_당일_전일_분틱_조회:
		if 데이터_길이 == xing.SizeT1310OutBlock {
			return (*xing.T1310OutBlock)(unsafe.Pointer(tr.Data)), nil
		}

		수량 := int(tr.DataLength) / xing.SizeT1310OutBlock1
		g_모음 := (*[큰수]xing.T1310OutBlock1)(unsafe.Pointer(tr.Data))[:수량:수량]
		return g_모음, nil
	case xing.TR_ETF_시간별_추이:
		if 데이터_길이 == xing.SizeT1902OutBlock {
			return (*xing.T1902OutBlock)(unsafe.Pointer(tr.Data)), nil
		}

		수량 := int(tr.DataLength) / xing.SizeT1902OutBlock1
		g_모음 := (*[큰수]xing.T1310OutBlock1)(unsafe.Pointer(tr.Data))[:수량:수량]
		return g_모음, nil
	case xing.TR기업정보_요약:
		switch 데이터_길이 {
		case xing.SizeT3320OutBlock:
			return (*xing.T3320OutBlock)(unsafe.Pointer(tr.Data)), nil
		case xing.SizeT3320OutBlock1:
			return (*xing.T3320OutBlock1)(unsafe.Pointer(tr.Data)), nil
		}
	case xing.TR현물_차트_틱:
		if 데이터_길이 == xing.SizeT8411OutBlock {
			return (*xing.T8411OutBlock)(unsafe.Pointer(tr.Data)), nil
		}

		수량 := int(tr.DataLength) / xing.SizeT8411OutBlock1
		g_모음 := (*[큰수]xing.T8411OutBlock1)(unsafe.Pointer(tr.Data))[:수량:수량]
		return g_모음, nil
	case xing.TR현물_차트_분:
		if 데이터_길이 == xing.SizeT8412OutBlock {
			return (*xing.T8412OutBlock)(unsafe.Pointer(tr.Data)), nil
		}

		수량 := int(tr.DataLength) / xing.SizeT8412OutBlock1
		g_모음 := (*[큰수]xing.T8412OutBlock1)(unsafe.Pointer(tr.Data))[:수량:수량]
		return g_모음, nil
	case xing.TR현물_차트_일주월:
		if 데이터_길이 == xing.SizeT8413OutBlock {
			return (*xing.T8413OutBlock)(unsafe.Pointer(tr.Data)), nil
		}

		수량 := int(tr.DataLength) / xing.SizeT8413OutBlock1
		g_모음 := (*[큰수]xing.T8413OutBlock1)(unsafe.Pointer(tr.Data))[:수량:수량]
		return g_모음, nil
	case xing.TR증시_주변_자금_추이:
		if 데이터_길이 == xing.SizeT8428OutBlock {
			return (*xing.T8428OutBlock)(unsafe.Pointer(tr.Data)), nil
		}

		수량 := int(tr.DataLength) / xing.SizeT8428OutBlock1
		g_모음 := (*[큰수]xing.T8428OutBlock1)(unsafe.Pointer(tr.Data))[:수량:수량]
		return g_모음, nil
	case xing.TR현물_종목_조회:
		수량 := int(tr.DataLength) / xing.SizeT8436OutBlock
		g_모음 := (*[큰수]xing.T8436OutBlock)(unsafe.Pointer(tr.Data))[:수량:수량]
		return g_모음, nil
	default:
		return nil, lib.New에러("구현되지 않은 xing.TR코드. %v", TR코드)
	}

	panic(lib.New에러with출력("예상하지 못한 데이터 길이 : '%v' '%v'", TR코드, 데이터_길이))
}