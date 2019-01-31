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
// #include <../../xing/types_c.h>
import "C"

import (
	"github.com/ghts/dep"
	"github.com/ghts/lib"
	"github.com/ghts/xing"
	"os"
)

func TR전송권한획득(TR코드 string) {
	lib.F조건부_패닉(TR코드 == "", "TR코드 없음.")
	TR전송_코드별_10분당_제한_확인(TR코드)
	TR전송_코드별_초당_제한_확인(TR코드)
}

func TR전송_코드별_10분당_제한_확인(TR코드 string) {
	전송_권한, 존재함 := tr전송_코드별_10분당_제한[TR코드]

	switch {
	case !존재함:
		return // 해당 TR코드 관련 제한이 존재하지 않음.
	case 전송_권한.TR코드() != TR코드:
		panic("예상하지 못한 경우.")
	}

	전송_권한.G전송_권한_획득()
}

func TR전송_코드별_초당_제한_확인(TR코드 string) {
	전송_권한, 존재함 := tr전송_코드별_초당_제한[TR코드]

	switch {
	case !존재함:
		panic(lib.New에러("전송제한을 찾을 수 없음 : '%v'", TR코드))
	case 전송_권한.TR코드() != TR코드:
		panic("예상하지 못한 경우.")
	case 전송_권한.G남은_수량() > 100:
		panic("전송 한도가 너무 큼. 1초당 한도와 10분당 한도를 혼동한 듯함.")
	}

	전송_권한.G전송_권한_획득()
}

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

func f의존성_확인() {
	dep.F의존관계_설정용_내용없는_함수()
}

func f자료형_문자열_해석(g *xing.TR_DATA) (자료형_문자열 string, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 자료형_문자열 = "" }}.S실행()

	TR코드 := lib.F2문자열_공백제거(g.TrCode)
	길이 := lib.F2정수_단순형(g.DataLength)

	switch TR코드 {
	case xing.TR현물_정상_주문:
		switch 길이 {
		case xing.SizeCSPAT00600OutBlock:
			return xing.P자료형_CSPAT00600OutBlock, nil
		case xing.SizeCSPAT00600OutBlock1:
			return xing.P자료형_CSPAT00600OutBlock1, nil
		case xing.SizeCSPAT00600OutBlock2:
			return xing.P자료형_CSPAT00600OutBlock2, nil
		}
	case xing.TR현물_정정_주문:
		switch 길이 {
		case xing.SizeCSPAT00700OutBlock:
			return xing.P자료형_CSPAT00700OutBlock, nil
		case xing.SizeCSPAT00700OutBlock1:
			return xing.P자료형_CSPAT00700OutBlock1, nil
		case xing.SizeCSPAT00700OutBlock2:
			return xing.P자료형_CSPAT00700OutBlock2, nil
		}
	case xing.TR현물_취소_주문:
		switch 길이 {
		case xing.SizeCSPAT00800OutBlock:
			return xing.P자료형_CSPAT00800OutBlock, nil
		case xing.SizeCSPAT00800OutBlock1:
			return xing.P자료형_CSPAT00800OutBlock1, nil
		case xing.SizeCSPAT00800OutBlock2:
			return xing.P자료형_CSPAT00800OutBlock2, nil
		}
	case xing.TR시간_조회:
		return xing.P자료형_T0167OutBlock, nil
	case xing.TR현물_호가_조회:
		return xing.P자료형_T1101OutBlock, nil
	case xing.TR현물_시세_조회:
		return xing.P자료형_T1102OutBlock, nil
	case xing.TR현물_기간별_조회:
		switch {
		case 길이 == xing.SizeT1305OutBlock:
			return xing.P자료형_T1305OutBlock, nil
		case 길이%xing.SizeT1305OutBlock1 == 0:
			return xing.P자료형_T1305OutBlock1, nil
		}
	case xing.TR현물_당일_전일_분틱_조회:
		switch {
		case 길이 == xing.SizeT1310OutBlock:
			return xing.P자료형_T1310OutBlock, nil
		case 길이%xing.SizeT1310OutBlock1 == 0:
			return xing.P자료형_T1310OutBlock1, nil
		}
	case xing.TR_ETF_시간별_추이:
		switch {
		case 길이 == xing.SizeT1902OutBlock:
			return xing.P자료형_T1902OutBlock, nil
		case 길이%xing.SizeT1902OutBlock1 == 0:
			return xing.P자료형_T1902OutBlock1, nil
		}
	case xing.TR기업정보_요약:
		switch 길이 {
		case xing.SizeT3320OutBlock:
			return xing.P자료형_T3320OutBlock, nil
		case xing.SizeT3320OutBlock1:
			return xing.P자료형_T3320OutBlock1, nil
		}
	case xing.TR현물_차트_틱:
		switch {
		case 길이 == xing.SizeT8411OutBlock:
			return xing.P자료형_T8411OutBlock, nil
		case 길이%xing.SizeT8411OutBlock1 == 0:
			return xing.P자료형_T8411OutBlock1, nil
		}
	case xing.TR현물_차트_분:
		switch {
		case 길이 == xing.SizeT8412OutBlock:
			return xing.P자료형_T8412OutBlock, nil
		case 길이%xing.SizeT8412OutBlock1 == 0:
			return xing.P자료형_T8412OutBlock1, nil
		}
	case xing.TR현물_차트_일주월:
		switch {
		case 길이 == xing.SizeT8413OutBlock:
			return xing.P자료형_T8413OutBlock, nil
		case 길이%xing.SizeT8413OutBlock1 == 0:
			return xing.P자료형_T8413OutBlock1, nil
		}
	case xing.TR증시_주변_자금_추이:
		switch {
		case 길이 == xing.SizeT8428OutBlock:
			return xing.P자료형_T8428OutBlock, nil
		case 길이%xing.SizeT8428OutBlock1 == 0:
			return xing.P자료형_T8428OutBlock1, nil
		}
	case xing.TR현물_종목_조회:
		if 길이%xing.SizeT8436OutBlock == 0 {
			return xing.P자료형_T8436OutBlock, nil
		}
	}

	panic(lib.New에러("예상하지 못한 TR코드 & 길이 : '%v' '%v'", TR코드, 길이))
}
