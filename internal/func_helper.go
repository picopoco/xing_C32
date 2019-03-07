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

// #cgo CFLAGS: -Wall
// #include <stdlib.h>
// #include <../../xing_common/types_c.h>
import "C"

import (
	"github.com/ghts/lib"
	"github.com/ghts/xing_common"
	"os"
)

func XingAPI디렉토리() (string, error) {
	파일경로, 에러 := lib.F실행파일_검색(xing_dll)
	if 에러 == nil {
		return lib.F디렉토리명(파일경로)
	}

	기본_위치 := `C:\eBEST\xingAPI\xingAPI.dll`
	if _, 에러 := os.Stat(기본_위치); 에러 == nil {
		lib.F실행경로_추가(기본_위치)

		if _, 에러 := lib.F실행파일_검색(xing_dll); 에러 != nil {
			return "", lib.New에러("실행경로에 추가시켰으나 여전히 찾을 수 없음.")
		}

		return lib.F디렉토리명(기본_위치)
	}

	파일경로, 에러 = lib.F파일_검색(`C:\`, xing_dll)
	if 에러 == nil {
		lib.F실행경로_추가(파일경로)

		if _, 에러 := lib.F실행파일_검색(xing_dll); 에러 != nil {
			return "", lib.New에러("실행경로에 추가시켰으나 여전히 찾을 수 없음.")
		}

		return lib.F디렉토리명(파일경로)
	}

	return "", lib.New에러("DLL파일을 찾을 수 없습니다.")
}

func f자료형_문자열_해석(g *xt.TR_DATA) (자료형_문자열 string, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 자료형_문자열 = "" }}.S실행()

	TR코드 := lib.F2문자열_공백제거(g.TrCode)
	길이 := lib.F2정수_단순형(g.DataLength)

	switch TR코드 {
	case xt.TR현물_정상_주문_CSPAT00600:
		switch 길이 {
		case xt.SizeCSPAT00600OutBlock:
			return xt.P자료형_CSPAT00600OutBlock, nil
		case xt.SizeCSPAT00600OutBlock1:
			return xt.P자료형_CSPAT00600OutBlock1, nil
		case xt.SizeCSPAT00600OutBlock2:
			return xt.P자료형_CSPAT00600OutBlock2, nil
		}
	case xt.TR현물_정정_주문_CSPAT00700:
		switch 길이 {
		case xt.SizeCSPAT00700OutBlock:
			return xt.P자료형_CSPAT00700OutBlock, nil
		case xt.SizeCSPAT00700OutBlock1:
			return xt.P자료형_CSPAT00700OutBlock1, nil
		case xt.SizeCSPAT00700OutBlock2:
			return xt.P자료형_CSPAT00700OutBlock2, nil
		}
	case xt.TR현물_취소_주문_CSPAT00800:
		switch 길이 {
		case xt.SizeCSPAT00800OutBlock:
			return xt.P자료형_CSPAT00800OutBlock, nil
		case xt.SizeCSPAT00800OutBlock1:
			return xt.P자료형_CSPAT00800OutBlock1, nil
		case xt.SizeCSPAT00800OutBlock2:
			return xt.P자료형_CSPAT00800OutBlock2, nil
		}
	case xt.TR시간_조회_t0167:
		return xt.P자료형_T0167OutBlock, nil
	case xt.TR체결_미체결_조회_t0425:
		// Non-block 모드는 Occurs데이터 수량을 나타내는 5바이트 추가됨.
		if 길이 < (xt.SizeT0425OutBlock+5) ||
			(길이-(xt.SizeT0425OutBlock+5))%xt.SizeT0425OutBlock1 != 0 {
			break
		}

		return xt.P자료형_T0425OutBlockAll, nil
	case xt.TR현물_호가_조회_t1101:
		return xt.P자료형_T1101OutBlock, nil
	case xt.TR현물_시세_조회_t1102:
		return xt.P자료형_T1102OutBlock, nil
	case xt.TR현물_기간별_조회_t1305:
		switch {
		case 길이 == xt.SizeT1305OutBlock:
			return xt.P자료형_T1305OutBlock, nil
		case 길이%xt.SizeT1305OutBlock1 == 0:
			return xt.P자료형_T1305OutBlock1, nil
		}
	case xt.TR현물_당일_전일_분틱_조회_t1310:
		switch {
		case 길이 == xt.SizeT1310OutBlock:
			return xt.P자료형_T1310OutBlock, nil
		case 길이%xt.SizeT1310OutBlock1 == 0:
			return xt.P자료형_T1310OutBlock1, nil
		}
	case xt.TR관리_불성실_투자유의_조회_t1404:
		switch {
		case 길이 == xt.SizeT1404OutBlock:
			return xt.P자료형_T1404OutBlock, nil
		case 길이%xt.SizeT1404OutBlock1 == 0:
			return xt.P자료형_T1404OutBlock1, nil
		}
	case xt.TR투자경고_매매정지_정리매매_조회_t1405:
		switch {
		case 길이 == xt.SizeT1405OutBlock:
			return xt.P자료형_T1405OutBlock, nil
		case 길이%xt.SizeT1405OutBlock1 == 0:
			return xt.P자료형_T1405OutBlock1, nil
		}
	case xt.TR_ETF_시간별_추이_t1902:
		switch {
		case 길이 == xt.SizeT1902OutBlock:
			return xt.P자료형_T1902OutBlock, nil
		case 길이%xt.SizeT1902OutBlock1 == 0:
			return xt.P자료형_T1902OutBlock1, nil
		}
	case xt.TR기업정보_요약_t3320:
		switch 길이 {
		case xt.SizeT3320OutBlock:
			return xt.P자료형_T3320OutBlock, nil
		case xt.SizeT3320OutBlock1:
			return xt.P자료형_T3320OutBlock1, nil
		}
	case xt.TR재무순위_종합_t3341:
		switch {
		case 길이 == xt.SizeT3341OutBlock:
			return xt.P자료형_T3341OutBlock, nil
		case 길이%xt.SizeT3341OutBlock1 == 0:
			return xt.P자료형_T3341OutBlock1, nil
		}
	case xt.TR현물_차트_틱_t8411:
		switch {
		case 길이 == xt.SizeT8411OutBlock:
			return xt.P자료형_T8411OutBlock, nil
		case 길이%xt.SizeT8411OutBlock1 == 0:
			return xt.P자료형_T8411OutBlock1, nil
		}
	case xt.TR현물_차트_분_t8412:
		switch {
		case 길이 == xt.SizeT8412OutBlock:
			return xt.P자료형_T8412OutBlock, nil
		case 길이%xt.SizeT8412OutBlock1 == 0:
			return xt.P자료형_T8412OutBlock1, nil
		}
	case xt.TR현물_차트_일주월_t8413:
		switch {
		case 길이 == xt.SizeT8413OutBlock:
			return xt.P자료형_T8413OutBlock, nil
		case 길이%xt.SizeT8413OutBlock1 == 0:
			return xt.P자료형_T8413OutBlock1, nil
		}
	case xt.TR증시_주변_자금_추이_t8428:
		switch {
		case 길이 == xt.SizeT8428OutBlock:
			return xt.P자료형_T8428OutBlock, nil
		case 길이%xt.SizeT8428OutBlock1 == 0:
			return xt.P자료형_T8428OutBlock1, nil
		}
	case xt.TR현물_종목_조회_t8436:
		if 길이%xt.SizeT8436OutBlock == 0 {
			return xt.P자료형_T8436OutBlock, nil
		}
	}

	panic(lib.New에러("예상하지 못한 TR코드 & 길이 : '%v' '%v'", TR코드, 길이))
}
