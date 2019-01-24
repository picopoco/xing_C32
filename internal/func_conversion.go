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
)

func f2시장구분(값 interface{}) lib.T시장구분 {
	문자열 := lib.F2문자열_EUC_KR_공백제거(값)

	switch 문자열 {
	case "KOSPI", "KOSPI200":
		return lib.P시장구분_코스피
	case "KOSDAQ":
		return lib.P시장구분_코스닥
	default:
		panic(lib.New에러("예상하지 못한 값 : '%v'", 문자열))
	}
}

func f2Xing매수매도(매수_매도 lib.T매수_매도) xing.T매수_매도 {
	switch 매수_매도 {
	case lib.P매도:
		return xing.P매도
	case lib.P매수:
		return xing.P매수
	default:
		panic(lib.New에러("예상하지 못한 매수 매도 구분값. %v", 매수_매도))
	}
}

func f2매수매도(매수_매도 xing.T매수_매도) (lib.T매수_매도, error) {
	switch 매수_매도 {
	case xing.P매도:
		return lib.P매도, nil
	case xing.P매수:
		return lib.P매수, nil
	default:
		return lib.T매수_매도(0), lib.New에러("예상하지 못한 매수 매도 구분값. %v", 매수_매도)
	}
}

func f2Xing주문조건(주문_조건 lib.T주문조건) xing.T주문조건 {
	switch 주문_조건 {
	case lib.P주문조건_없음:
		return xing.P주문조건_없음
	case lib.P주문조건_IOC:
		return xing.P주문조건_IOC
	case lib.P주문조건_FOK:
		return xing.P주문조건_FOK
	default:
		panic(lib.New에러("예상하지 못한 신용거래_구분 값. %v", 주문_조건))
	}
}

func f2주문조건(주문_조건 xing.T주문조건) lib.T주문조건 {
	switch 주문_조건 {
	case xing.P주문조건_없음:
		return lib.P주문조건_없음
	case xing.P주문조건_IOC:
		return lib.P주문조건_IOC
	case xing.P주문조건_FOK:
		return lib.P주문조건_FOK
	default:
		panic(lib.New에러("예상하지 못한 주문_조건 값. %v", 주문_조건))
	}
}

func f2Xing신용거래_구분(신용거래_구분 lib.T신용거래_구분) xing.T신용거래_구분 {
	switch 신용거래_구분 {
	case lib.P신용거래_해당없음:
		return xing.P신용거래_아님
	case lib.P신용거래_유통융자신규:
		return xing.P유통융자신규
	case lib.P신용거래_자기융자신규:
		return xing.P자기융자신규
	case lib.P신용거래_유통대주신규:
		return xing.P유통대주신규
	case lib.P신용거래_자기대주신규:
		return xing.P자기대주신규
	case lib.P신용거래_유통융자상환:
		return xing.P유통융자상환
	case lib.P신용거래_자기융자상환:
		return xing.P자기융자상환
	case lib.P신용거래_유통대주상환:
		return xing.P유통대주상환
	case lib.P신용거래_자기대주상환:
		return xing.P자기대주상환
	case lib.P신용거래_예탁담보대출상환:
		return xing.P예탁담보대출상환
	default:
		panic(lib.New에러("예상하지 못한 신용거래_구분 값. %v", 신용거래_구분))
	}
}

func f2신용거래_구분(신용거래_구분 xing.T신용거래_구분) lib.T신용거래_구분 {
	switch 신용거래_구분 {
	case xing.P유통융자신규:
		return lib.P신용거래_유통융자신규
	case xing.P자기융자신규:
		return lib.P신용거래_자기융자신규
	case xing.P유통대주신규:
		return lib.P신용거래_유통대주신규
	case xing.P자기대주신규:
		return lib.P신용거래_자기대주신규
	case xing.P유통융자상환:
		return lib.P신용거래_유통융자상환
	case xing.P자기융자상환:
		return lib.P신용거래_자기융자상환
	case xing.P유통대주상환:
		return lib.P신용거래_유통대주상환
	}

	return lib.P신용거래_해당없음
}

func f2Xing호가유형(호가_유형 lib.T호가유형) xing.T호가유형 {
	switch 호가_유형 {
	case lib.P호가유형_지정가:
		return xing.P지정가
	case lib.P호가유형_시장가:
		return xing.P시장가
	case lib.P호가유형_조건부_지정가:
		return xing.P조건부_지정가
	case lib.P호가유형_최유리_지정가:
		return xing.P최유리_지정가
	case lib.P호가유형_최우선_지정가:
		return xing.P최우선_지정가
	case lib.P호가유형_시간외종가_장개시전:
		return xing.P시간외종가_장개시전
	case lib.P호가유형_시간외종가:
		return xing.P시간외종가
	case lib.P호가유형_시간외단일가:
		return xing.P시간외단일가
	default:
		panic(lib.New에러("예상하지 못한 호가_유형 값. %v", 호가_유형))
	}
}

func f2호가유형(호가_유형 xing.T호가유형) lib.T호가유형 {
	switch 호가_유형 {
	case xing.P지정가:
		return lib.P호가유형_지정가
	case xing.P시장가:
		return lib.P호가유형_시장가
	case xing.P조건부_지정가:
		return lib.P호가유형_조건부_지정가
	case xing.P최유리_지정가:
		return lib.P호가유형_최유리_지정가
	case xing.P최우선_지정가:
		return lib.P호가유형_최우선_지정가
	case xing.P시간외종가_장개시전:
		return lib.P호가유형_시간외종가_장개시전
	case xing.P시간외종가:
		return lib.P호가유형_시간외종가
	case xing.P시간외단일가:
		return lib.P호가유형_시간외단일가
	default:
		panic(lib.New에러("예상하지 못한 호가_유형 값. '%v'", 호가_유형))
	}
}

func f2Xing수정구분(값 int64) []xing.T수정구분 {
	if 값 == 0 {
		return []xing.T수정구분{xing.P수정구분_없음}
	}

	수정구분_ALL := []xing.T수정구분{
		xing.P수정구분_불성실공시종목,
		xing.P수정구분_수정주가,
		xing.P수정구분_뮤추얼펀드,
		xing.P수정구분_정리매매종목,
		xing.P수정구분_ETF종목,
		xing.P수정구분_증거금100퍼센트,
		xing.P수정구분_종가범위연장,
		xing.P수정구분_시가범위연장,
		xing.P수정구분_권리중간배당락,
		xing.P수정구분_중간배당락,
		xing.P수정구분_CB발동예고,
		xing.P수정구분_우선주,
		xing.P수정구분_기준가조정,
		xing.P수정구분_거래정지,
		xing.P수정구분_투자경고,
		xing.P수정구분_관리종목,
		xing.P수정구분_기업분할,
		xing.P수정구분_주식병합,
		xing.P수정구분_액면병합,
		xing.P수정구분_액면분할,
		xing.P수정구분_배당락,
		xing.P수정구분_권리락}

	수정구분_모음 := make([]xing.T수정구분, 0)
	잔여값 := uint32(값)

	for _, 수정구분 := range 수정구분_ALL {
		if 잔여값 >= 수정구분.G정수값() {
			잔여값 -= 수정구분.G정수값()
			수정구분_모음 = append(수정구분_모음, 수정구분)
		}
	}

	if 잔여값 > 0 {
		panic(lib.New에러with출력("예상하지 못한 값 : '%v'", 값))
	}

	return 수정구분_모음
}
