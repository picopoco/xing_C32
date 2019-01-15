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

	"bytes"
	"errors"
	"math"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

const 크기CSPAT00600OutBlock1 = int(unsafe.Sizeof(CSPAT00600OutBlock1{}))
const 크기CSPAT00600OutBlock2 = int(unsafe.Sizeof(CSPAT00600OutBlock2{}))
const 크기CSPAT00700OutBlock1 = int(unsafe.Sizeof(CSPAT00700OutBlock1{}))
const 크기CSPAT00700OutBlock2 = int(unsafe.Sizeof(CSPAT00700OutBlock2{}))
const 크기CSPAT00800OutBlock1 = int(unsafe.Sizeof(CSPAT00800OutBlock1{}))
const 크기CSPAT00800OutBlock2 = int(unsafe.Sizeof(CSPAT00800OutBlock2{}))
const 크기T1301OutBlock = int(unsafe.Sizeof(T1301OutBlock{}))
const 크기T1301OutBlock1 = int(unsafe.Sizeof(T1301OutBlock1{}))
const 크기T1305OutBlock = int(unsafe.Sizeof(T1305OutBlock{}))
const 크기T1305OutBlock1 = int(unsafe.Sizeof(T1305OutBlock1{}))
const 크기T1310OutBlock = int(unsafe.Sizeof(T1310OutBlock{}))
const 크기T1310OutBlock1 = int(unsafe.Sizeof(T1310OutBlock1{}))
const 크기T1902OutBlock = int(unsafe.Sizeof(T1902OutBlock{}))
const 크기T1902OutBlock1 = int(unsafe.Sizeof(T1902OutBlock1{}))
const 크기T8428OutBlock = int(unsafe.Sizeof(T8428OutBlock{}))
const 크기T8428OutBlock1 = int(unsafe.Sizeof(T8428OutBlock1{}))
const 크기T8436OutBlock = int(unsafe.Sizeof(T8436OutBlock{}))

func tr데이터_해석(tr *TR_DATA) (값 interface{}, 에러 error) {
	lib.S예외처리{M에러: &에러, M함수: func() {
		체크("에러패닉 처리")
		값 = nil
	}}.S실행()

	TR코드 := lib.F2문자열(tr.TrCode)
	데이터_길이 := int(tr.DataLength)

	switch TR코드 {
	default:
		return nil, lib.New에러("구현되지 않은 TR코드. %v", TR코드)
	case xing.TR현물_정상_주문:
		switch 데이터_길이 {
		case 크기CSPAT00600OutBlock1:
			체크()
			return New현물_정상주문_응답1(tr)
		case 크기CSPAT00600OutBlock2:
			체크()
			return New현물_정상주문_응답2(tr)
		default:
			s := new(xing.S현물_정상_주문_응답)
			s.M응답1 = 에러체크(New현물_정상주문_응답1(tr)).(*xing.S현물_정상_주문_응답1)
			s.M응답2, 에러 = New현물_정상주문_응답2(tr)
			return s, 에러
		}
	case xing.TR현물_정정_주문:
		switch 데이터_길이 {
		case 크기CSPAT00700OutBlock1:
			체크()
			return New현물_정정주문_응답1(tr)
		case 크기CSPAT00700OutBlock2:
			체크()
			return New현물_정정주문_응답2(tr)
		default:
			s := new(xing.S현물_정정_주문_응답)
			s.M응답1 = 에러체크(New현물_정정주문_응답1(tr)).(*xing.S현물_정정_주문_응답1)
			s.M응답2, 에러 = New현물_정정주문_응답2(tr)
			return s, 에러
		}
	case xing.TR현물_취소_주문:
		switch 데이터_길이 {
		case 크기CSPAT00800OutBlock1:
			체크()
			return New현물_취소주문_응답1(tr)
		case 크기CSPAT00800OutBlock2:
			체크()
			return New현물_취소주문_응답2(tr)
		default:
			s := new(xing.S현물_취소_주문_응답)
			s.M응답1 = 에러체크(New현물_취소주문_응답1(tr)).(*xing.S현물_취소_주문_응답1)
			s.M응답2, 에러 = New현물_취소주문_응답2(tr)

			return s, 에러
		}
	case xing.TR시간_조회:
		g := (*T0167OutBlock)(unsafe.Pointer(tr.Data))
		날짜_문자열 := lib.F2문자열(g.Date)
		시간_문자열 := lib.F2문자열(g.Time)

		return lib.F2포맷된_시각("20060102150405.99999999", 날짜_문자열+시간_문자열[:6]+"."+시간_문자열[7:])
	case xing.TR현물_호가_조회:
		return New현물호가조회(tr)
	case xing.TR현물_시세_조회:
		return New현물시세조회_응답(tr)
	//case xing.TR현물_시간대별_체결_조회:
	//	switch 데이터_길이 {
	//	case 크기T1301OutBlock:
	//		return New현물_시간대별_체결_조회_응답_헤더(tr)
	//	default:
	//		return New현물_시간대별_체결_조회_응답_반복값_모음(tr)
	//	}
	case xing.TR현물_기간별_조회:
		switch 데이터_길이 {
		case 크기T1305OutBlock:
			return New현물_기간별_조회_응답_헤더(tr)
		default:
			return New현물_기간별_조회_응답_반복값_모음(tr)
		}
	case xing.TR현물_당일_전일_분틱_조회:
		switch 데이터_길이 {
		case 크기T1310OutBlock:
			return New현물_당일전일분틱조회_응답_헤더(tr)
		default:
			return New현물_당일전일분틱조회_응답_반복값_모음(tr)
		}
	case xing.TR_ETF_시간별_추이:
		switch 데이터_길이 {
		case 크기T1902OutBlock:
			return NewETF시간별_추이_응답_헤더(tr)
		default:
			return NewETF시간별_추이_응답_반복값_모음(tr)
		}
	case xing.TR증시_주변_자금_추이:
		switch 데이터_길이 {
		case 크기T8428OutBlock:
			return New증시주변자금추이_응답_헤더(tr)
		default:
			return New증시주변자금추이_응답_반복값_모음(tr)
		}
	case xing.TR현물_종목_조회:
		return New주식종목조회_응답_반복값_모음(tr)
	}
}

func New현물호가조회(tr *TR_DATA) (s *xing.S현물_호가조회_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { s = nil }}.S실행()

	g := (*T1101OutBlock)(unsafe.Pointer(tr.Data))

	s = new(xing.S현물_호가조회_응답)
	s.M한글명 = lib.F2문자열_EUC_KR(g.Hname)
	s.M현재가 = lib.F2정수64_단순형(g.Price)
	s.M전일대비구분 = xing.T전일대비_구분(lib.F2정수64_단순형(g.Sign))
	s.M전일대비등락폭 = lib.F2정수64_단순형(g.Change)
	s.M등락율 = lib.F2실수_소숫점_추가(g.Diff, 2)
	s.M누적거래량 = lib.F2정수64_단순형(g.Volume)
	s.M전일종가 = lib.F2정수64_단순형(g.Jnilclose)
	s.M매도호가_모음 = make([]int64, 10)
	s.M매수호가_모음 = make([]int64, 10)
	s.M매도호가수량_모음 = make([]int64, 10)
	s.M매수호가수량_모음 = make([]int64, 10)
	s.M직전매도대비수량_모음 = make([]int64, 10)
	s.M직전매수대비수량_모음 = make([]int64, 10)
	s.M매도호가_모음[0] = lib.F2정수64_단순형(g.Offerho1)
	s.M매수호가_모음[0] = lib.F2정수64_단순형(g.Bidho1)
	s.M매도호가수량_모음[0] = lib.F2정수64_단순형(g.Offerrem1)
	s.M매수호가수량_모음[0] = lib.F2정수64_단순형(g.Bidrem1)
	s.M직전매도대비수량_모음[0] = lib.F2정수64_단순형(g.Preoffercha1)
	s.M직전매수대비수량_모음[0] = lib.F2정수64_단순형(g.Prebidcha1)
	s.M매도호가_모음[1] = lib.F2정수64_단순형(g.Offerho2)
	s.M매수호가_모음[1] = lib.F2정수64_단순형(g.Bidho2)
	s.M매도호가수량_모음[1] = lib.F2정수64_단순형(g.Offerrem2)
	s.M매수호가수량_모음[1] = lib.F2정수64_단순형(g.Bidrem2)
	s.M직전매도대비수량_모음[1] = lib.F2정수64_단순형(g.Preoffercha2)
	s.M직전매수대비수량_모음[1] = lib.F2정수64_단순형(g.Prebidcha2)
	s.M매도호가_모음[2] = lib.F2정수64_단순형(g.Offerho3)
	s.M매수호가_모음[2] = lib.F2정수64_단순형(g.Bidho3)
	s.M매도호가수량_모음[2] = lib.F2정수64_단순형(g.Offerrem3)
	s.M매수호가수량_모음[2] = lib.F2정수64_단순형(g.Bidrem3)
	s.M직전매도대비수량_모음[2] = lib.F2정수64_단순형(g.Preoffercha3)
	s.M직전매수대비수량_모음[2] = lib.F2정수64_단순형(g.Prebidcha3)
	s.M매도호가_모음[3] = lib.F2정수64_단순형(g.Offerho4)
	s.M매수호가_모음[3] = lib.F2정수64_단순형(g.Bidho4)
	s.M매도호가수량_모음[3] = lib.F2정수64_단순형(g.Offerrem4)
	s.M매수호가수량_모음[3] = lib.F2정수64_단순형(g.Bidrem4)
	s.M직전매도대비수량_모음[3] = lib.F2정수64_단순형(g.Preoffercha4)
	s.M직전매수대비수량_모음[3] = lib.F2정수64_단순형(g.Prebidcha4)
	s.M매도호가_모음[4] = lib.F2정수64_단순형(g.Offerho5)
	s.M매수호가_모음[4] = lib.F2정수64_단순형(g.Bidho5)
	s.M매도호가수량_모음[4] = lib.F2정수64_단순형(g.Offerrem5)
	s.M매수호가수량_모음[4] = lib.F2정수64_단순형(g.Bidrem5)
	s.M직전매도대비수량_모음[4] = lib.F2정수64_단순형(g.Preoffercha5)
	s.M직전매수대비수량_모음[4] = lib.F2정수64_단순형(g.Prebidcha5)
	s.M매도호가_모음[5] = lib.F2정수64_단순형(g.Offerho6)
	s.M매수호가_모음[5] = lib.F2정수64_단순형(g.Bidho6)
	s.M매도호가수량_모음[5] = lib.F2정수64_단순형(g.Offerrem6)
	s.M매수호가수량_모음[5] = lib.F2정수64_단순형(g.Bidrem6)
	s.M직전매도대비수량_모음[5] = lib.F2정수64_단순형(g.Preoffercha6)
	s.M직전매수대비수량_모음[5] = lib.F2정수64_단순형(g.Prebidcha6)
	s.M매도호가_모음[6] = lib.F2정수64_단순형(g.Offerho7)
	s.M매수호가_모음[6] = lib.F2정수64_단순형(g.Bidho7)
	s.M매도호가수량_모음[6] = lib.F2정수64_단순형(g.Offerrem7)
	s.M매수호가수량_모음[6] = lib.F2정수64_단순형(g.Bidrem7)
	s.M직전매도대비수량_모음[6] = lib.F2정수64_단순형(g.Preoffercha7)
	s.M직전매수대비수량_모음[6] = lib.F2정수64_단순형(g.Prebidcha7)
	s.M매도호가_모음[7] = lib.F2정수64_단순형(g.Offerho8)
	s.M매수호가_모음[7] = lib.F2정수64_단순형(g.Bidho8)
	s.M매도호가수량_모음[7] = lib.F2정수64_단순형(g.Offerrem8)
	s.M매수호가수량_모음[7] = lib.F2정수64_단순형(g.Bidrem8)
	s.M직전매도대비수량_모음[7] = lib.F2정수64_단순형(g.Preoffercha8)
	s.M직전매수대비수량_모음[7] = lib.F2정수64_단순형(g.Prebidcha8)
	s.M매도호가_모음[8] = lib.F2정수64_단순형(g.Offerho9)
	s.M매수호가_모음[8] = lib.F2정수64_단순형(g.Bidho9)
	s.M매도호가수량_모음[8] = lib.F2정수64_단순형(g.Offerrem9)
	s.M매수호가수량_모음[8] = lib.F2정수64_단순형(g.Bidrem9)
	s.M직전매도대비수량_모음[8] = lib.F2정수64_단순형(g.Preoffercha9)
	s.M직전매수대비수량_모음[8] = lib.F2정수64_단순형(g.Prebidcha9)
	s.M매도호가_모음[9] = lib.F2정수64_단순형(g.Offerho10)
	s.M매수호가_모음[9] = lib.F2정수64_단순형(g.Bidho10)
	s.M매도호가수량_모음[9] = lib.F2정수64_단순형(g.Offerrem10)
	s.M매수호가수량_모음[9] = lib.F2정수64_단순형(g.Bidrem10)
	s.M직전매도대비수량_모음[9] = lib.F2정수64_단순형(g.Preoffercha10)
	s.M직전매수대비수량_모음[9] = lib.F2정수64_단순형(g.Prebidcha10)
	s.M매도호가수량합 = lib.F2정수64_단순형(g.Offer)
	s.M매수호가수량합 = lib.F2정수64_단순형(g.Bid)
	s.M직전매도대비수량합 = lib.F2정수64_단순형(g.Preoffercha)
	s.M직전매수대비수량합 = lib.F2정수64_단순형(g.Prebidcha)

	if 시각_문자열 := lib.F2문자열_공백제거(g.Hotime); 시각_문자열 != "" {
		s.M수신시간 = lib.F2일자별_시각_단순형(당일.G값(), "150405.999", 시각_문자열[:6]+"."+시각_문자열[6:])
	}

	s.M예상체결가격 = lib.F2정수64_단순형(g.Yeprice)
	s.M예상체결수량 = lib.F2정수64_단순형(g.Yevolume)
	s.M예상체결전일구분 = xing.T전일대비_구분(lib.F2정수64_단순형(g.Yesign))
	s.M예상체결전일대비 = lib.F2정수64_단순형(g.Yechange)
	s.M예상체결등락율 = lib.F2실수_소숫점_추가(g.Yediff, 2)
	s.M시간외매도잔량 = lib.F2정수64_단순형(g.Tmoffer)
	s.M시간외매수잔량 = lib.F2정수64_단순형(g.Tmbid)
	s.M동시호가_구분 = xing.T동시호가_구분(lib.F2정수64_단순형(g.Status))
	s.M종목코드 = lib.F2문자열_공백제거(g.Shcode)
	s.M상한가 = lib.F2정수64_단순형(g.Uplmtprice)
	s.M하한가 = lib.F2정수64_단순형(g.Dnlmtprice)
	s.M시가 = lib.F2정수64_단순형(g.Open)
	s.M고가 = lib.F2정수64_단순형(g.High)
	s.M저가 = lib.F2정수64_단순형(g.Low)

	return s, nil
}

func New현물시세조회_응답(tr *TR_DATA) (s *xing.S현물_시세조회_응답, 에러 error) {
	defer lib.S예외처리{
		M에러: &에러,
		M함수: func() {
			lib.F문자열_출력("'%v'(%v)", s.M한글명, s.M종목코드)
			s = nil
		}}.S실행()

	당일값 := 당일.G값()

	g := (*T1102OutBlock)(unsafe.Pointer(tr.Data))

	s = new(xing.S현물_시세조회_응답)
	s.M종목코드 = lib.F2문자열_공백제거(g.Shcode)
	s.M한글명 = lib.F2문자열_EUC_KR(g.Hname)
	s.M현재가 = lib.F2정수64_단순형(g.Price)
	s.M전일대비구분 = xing.T전일대비_구분(lib.F2정수64_단순형(g.Sign))
	s.M전일대비등락폭 = lib.F2정수64_단순형(g.Change)
	s.M등락율 = lib.F2실수_소숫점_추가(g.Diff, 2)
	s.M누적거래량 = lib.F2정수64_단순형(g.Volume)
	s.M기준가 = lib.F2정수64_단순형(g.Recprice)
	s.M가중평균 = lib.F2정수64_단순형(g.Avg)
	s.M상한가 = lib.F2정수64_단순형(g.Uplmtprice)
	s.M하한가 = lib.F2정수64_단순형(g.Dnlmtprice)
	s.M전일거래량 = lib.F2정수64_단순형(g.Jnilvolume)
	s.M거래량차 = lib.F2정수64_단순형(g.Volumediff)
	s.M시가 = lib.F2정수64_단순형(g.Open)
	s.M시가시간 = lib.F2일자별_시각_단순형_공백은_초기값(당일값, "150405", g.Opentime)
	s.M고가 = lib.F2정수64_단순형(g.High)
	s.M고가시간 = lib.F2일자별_시각_단순형_공백은_초기값(당일값, "150405", g.Hightime)
	s.M저가 = lib.F2정수64_단순형(g.Low)
	s.M저가시간 = lib.F2일자별_시각_단순형_공백은_초기값(당일값, "150405", g.Lowtime)
	s.M52주_최고가 = lib.F2정수64_단순형(g.High52w)
	s.M52주_최고가일 = lib.F2포맷된_시각_단순형_공백은_초기값("20060102", g.High52wdate)
	s.M52주_최저가 = lib.F2정수64_단순형(g.Low52w)
	s.M52주_최저가일 = lib.F2포맷된_시각_단순형_공백은_초기값("20060102", g.Low52wdate)
	s.M소진율 = lib.F2실수_소숫점_추가(g.Exhratio, 2)
	s.PER = lib.F2실수_소숫점_추가(g.Per, 2)
	s.PBRX = lib.F2실수_소숫점_추가(g.Pbrx, 2)
	s.M상장주식수_천 = lib.F2정수64_단순형(g.Listing)
	s.M증거금율 = lib.F2정수64_단순형(g.Jkrate)
	s.M수량단위 = lib.F2정수64_단순형(g.Memedan)
	s.M매도증권사코드_모음 = make([]string, 5)
	s.M매수증권사코드_모음 = make([]string, 5)
	s.M매도증권사명_모음 = make([]string, 5)
	s.M매수증권사명_모음 = make([]string, 5)
	s.M총매도수량_모음 = make([]int64, 5)
	s.M총매수수량_모음 = make([]int64, 5)
	s.M매도증감_모음 = make([]int64, 5)
	s.M매수증감_모음 = make([]int64, 5)
	s.M매도비율_모음 = make([]float64, 5)
	s.M매수비율_모음 = make([]float64, 5)
	s.M매도증권사코드_모음[0] = lib.F2문자열_공백제거(g.Offernocd1)
	s.M매수증권사코드_모음[0] = lib.F2문자열_공백제거(g.Bidnocd1)
	s.M매도증권사명_모음[0] = lib.F2문자열_EUC_KR(g.Offerno1)
	s.M매수증권사명_모음[0] = lib.F2문자열_EUC_KR(g.Bidno1)
	s.M총매도수량_모음[0] = lib.F2정수64_단순형(g.Dvol1)
	s.M총매수수량_모음[0] = lib.F2정수64_단순형(g.Svol1)
	s.M매도증감_모음[0] = lib.F2정수64_단순형(g.Dcha1)
	s.M매수증감_모음[0] = lib.F2정수64_단순형(g.Scha1)
	s.M매도비율_모음[0] = lib.F2실수_소숫점_추가(g.Ddiff1, 2)
	s.M매수비율_모음[0] = lib.F2실수_소숫점_추가(g.Sdiff1, 2)
	s.M매도증권사코드_모음[1] = lib.F2문자열_공백제거(g.Offernocd2)
	s.M매수증권사코드_모음[1] = lib.F2문자열_공백제거(g.Bidnocd2)
	s.M매도증권사명_모음[1] = lib.F2문자열_EUC_KR(g.Offerno2)
	s.M매수증권사명_모음[1] = lib.F2문자열_EUC_KR(g.Bidno2)
	s.M총매도수량_모음[1] = lib.F2정수64_단순형(g.Dvol2)
	s.M총매수수량_모음[1] = lib.F2정수64_단순형(g.Svol2)
	s.M매도증감_모음[1] = lib.F2정수64_단순형(g.Dcha2)
	s.M매수증감_모음[1] = lib.F2정수64_단순형(g.Scha2)
	s.M매도비율_모음[1] = lib.F2실수_소숫점_추가(g.Ddiff2, 2)
	s.M매수비율_모음[1] = lib.F2실수_소숫점_추가(g.Sdiff2, 2)
	s.M매도증권사코드_모음[2] = lib.F2문자열_공백제거(g.Offernocd3)
	s.M매수증권사코드_모음[2] = lib.F2문자열_공백제거(g.Bidnocd3)
	s.M매도증권사명_모음[2] = lib.F2문자열_EUC_KR(g.Offerno3)
	s.M매수증권사명_모음[2] = lib.F2문자열_EUC_KR(g.Bidno3)
	s.M총매도수량_모음[2] = lib.F2정수64_단순형(g.Dvol3)
	s.M총매수수량_모음[2] = lib.F2정수64_단순형(g.Svol3)
	s.M매도증감_모음[2] = lib.F2정수64_단순형(g.Dcha3)
	s.M매수증감_모음[2] = lib.F2정수64_단순형(g.Scha3)
	s.M매도비율_모음[2] = lib.F2실수_소숫점_추가(g.Ddiff3, 2)
	s.M매수비율_모음[2] = lib.F2실수_소숫점_추가(g.Sdiff3, 2)
	s.M매도증권사코드_모음[3] = lib.F2문자열_공백제거(g.Offernocd4)
	s.M매수증권사코드_모음[3] = lib.F2문자열_공백제거(g.Bidnocd4)
	s.M매도증권사명_모음[3] = lib.F2문자열_EUC_KR(g.Offerno4)
	s.M매수증권사명_모음[3] = lib.F2문자열_EUC_KR(g.Bidno4)
	s.M총매도수량_모음[3] = lib.F2정수64_단순형(g.Dvol4)
	s.M총매수수량_모음[3] = lib.F2정수64_단순형(g.Svol4)
	s.M매도증감_모음[3] = lib.F2정수64_단순형(g.Dcha4)
	s.M매수증감_모음[3] = lib.F2정수64_단순형(g.Scha4)
	s.M매도비율_모음[3] = lib.F2실수_소숫점_추가(g.Ddiff4, 2)
	s.M매수비율_모음[3] = lib.F2실수_소숫점_추가(g.Sdiff4, 2)
	s.M매도증권사코드_모음[4] = lib.F2문자열_공백제거(g.Offernocd5)
	s.M매수증권사코드_모음[4] = lib.F2문자열_공백제거(g.Bidnocd5)
	s.M매도증권사명_모음[4] = lib.F2문자열_EUC_KR(g.Offerno5)
	s.M매수증권사명_모음[4] = lib.F2문자열_EUC_KR(g.Bidno5)
	s.M총매도수량_모음[4] = lib.F2정수64_단순형(g.Dvol5)
	s.M총매수수량_모음[4] = lib.F2정수64_단순형(g.Svol5)
	s.M매도증감_모음[4] = lib.F2정수64_단순형(g.Dcha5)
	s.M매수증감_모음[4] = lib.F2정수64_단순형(g.Scha5)
	s.M매도비율_모음[4] = lib.F2실수_소숫점_추가(g.Ddiff5, 2)
	s.M매수비율_모음[4] = lib.F2실수_소숫점_추가(g.Sdiff5, 2)
	s.M외국계_매도_합계수량 = lib.F2정수64_단순형(g.Fwdvl)
	s.M외국계_매도_직전대비 = lib.F2정수64_단순형(g.Ftradmdcha)
	s.M외국계_매도_비율 = lib.F2실수_소숫점_추가(g.Ftradmddiff, 2)
	s.M외국계_매수_합계수량 = lib.F2정수64_단순형(g.Fwsvl)
	s.M외국계_매수_직전대비 = lib.F2정수64_단순형(g.Ftradmscha)
	s.M외국계_매수_비율 = lib.F2실수_소숫점_추가(g.Ftradmsdiff, 2)
	s.M회전율 = lib.F2실수_소숫점_추가(g.Vol, 2)
	s.M누적거래대금 = lib.F2정수64_단순형(g.Value)
	s.M전일동시간거래량 = lib.F2정수64_단순형(g.Jvolume)
	s.M연중_최고가 = lib.F2정수64_단순형(g.Highyear)
	s.M연중_최고가_일자 = lib.F2포맷된_시각_단순형("20060102", g.Highyeardate)
	s.M연중_최저가 = lib.F2정수64_단순형(g.Lowyear)
	s.M연중_최저가_일자 = lib.F2포맷된_시각_단순형("20060102", g.Lowyeardate)
	s.M목표가 = lib.F2정수64_단순형(g.Target)
	s.M자본금 = lib.F2정수64_단순형(g.Capital)
	s.M유동주식수 = lib.F2정수64_단순형(g.Abscnt)
	s.M액면가 = lib.F2정수64_단순형(g.Parprice)
	s.M결산월 = uint8(lib.F2정수64_단순형_공백은_0(g.Gsmm))
	s.M대용가 = lib.F2정수64_단순형(g.Subprice)
	s.M시가총액_억 = lib.F2정수64_단순형(g.Total)
	s.M상장일 = lib.F2포맷된_시각_단순형("20060102", g.Listdate)
	s.M전분기명 = lib.F2문자열_EUC_KR_공백제거(g.Name)
	s.M전분기_매출액 = lib.F2정수64_단순형(g.Bfsales)
	s.M전분기_영업이익 = lib.F2정수64_단순형(g.Bfoperatingincome)
	s.M전분기_경상이익 = lib.F2정수64_단순형(g.Bfordinaryincome)
	s.M전분기_순이익 = lib.F2정수64_단순형(g.Bfnetincome)
	s.M전분기EPS = lib.F2실수_소숫점_추가(g.Bfeps, 2)
	s.M전전분기명 = lib.F2문자열_EUC_KR_공백제거(g.Name2)
	s.M전전분기_매출액 = lib.F2정수64_단순형(g.Bfsales2)
	s.M전전분기_영업이익 = lib.F2정수64_단순형(g.Bfoperatingincome2)
	s.M전전분기_경상이익 = lib.F2정수64_단순형(g.Bfordinaryincome2)
	s.M전전분기_순이익 = lib.F2정수64_단순형(g.Bfnetincome2)
	s.M전전분기EPS = lib.F2실수_소숫점_추가(g.Bfeps2, 2)
	s.M전년대비매출액 = lib.F2실수_소숫점_추가(g.Salert, 2)
	s.M전년대비영업이익 = lib.F2실수_소숫점_추가(g.Opert, 2)
	s.M전년대비경상이익 = lib.F2실수_소숫점_추가(g.Ordrt, 2)
	s.M전년대비순이익 = lib.F2실수_소숫점_추가(g.Netrt, 2)
	s.M전년대비EPS = lib.F2실수_소숫점_추가(g.Epsrt, 2)
	s.M락구분 = lib.F2문자열_EUC_KR(g.Info1)
	s.M관리_급등구분 = lib.F2문자열_EUC_KR(g.Info2)
	s.M정지_연장구분 = lib.F2문자열_EUC_KR(g.Info3)
	s.M투자_불성실구분 = lib.F2문자열_EUC_KR(g.Info4)
	s.M시장구분 = f2시장구분(g.Janginfo)
	s.T_PER = lib.F2실수_소숫점_추가(g.T_per, 2)
	s.M통화ISO코드 = lib.F2문자열_공백제거(g.Tonghwa)
	s.M총매도대금_모음 = make([]int64, 5)
	s.M총매수대금_모음 = make([]int64, 5)
	s.M총매도대금_모음[0] = lib.F2정수64_단순형_공백은_0(g.Dval1)
	s.M총매수대금_모음[0] = lib.F2정수64_단순형_공백은_0(g.Sval1)
	s.M총매도대금_모음[1] = lib.F2정수64_단순형_공백은_0(g.Dval2)
	s.M총매수대금_모음[1] = lib.F2정수64_단순형_공백은_0(g.Sval2)
	s.M총매도대금_모음[2] = lib.F2정수64_단순형_공백은_0(g.Dval3)
	s.M총매수대금_모음[2] = lib.F2정수64_단순형_공백은_0(g.Sval3)
	s.M총매도대금_모음[3] = lib.F2정수64_단순형_공백은_0(g.Dval4)
	s.M총매수대금_모음[3] = lib.F2정수64_단순형_공백은_0(g.Sval4)
	s.M총매도대금_모음[4] = lib.F2정수64_단순형_공백은_0(g.Dval5)
	s.M총매수대금_모음[4] = lib.F2정수64_단순형_공백은_0(g.Sval5)
	s.M총매도평단가_모음 = make([]int64, 5)
	s.M총매수평단가_모음 = make([]int64, 5)
	s.M총매도평단가_모음[0] = lib.F2정수64_단순형_공백은_0(g.Davg1)
	s.M총매수평단가_모음[0] = lib.F2정수64_단순형_공백은_0(g.Savg1)
	s.M총매도평단가_모음[1] = lib.F2정수64_단순형_공백은_0(g.Davg2)
	s.M총매수평단가_모음[1] = lib.F2정수64_단순형_공백은_0(g.Savg2)
	s.M총매도평단가_모음[2] = lib.F2정수64_단순형_공백은_0(g.Davg3)
	s.M총매수평단가_모음[2] = lib.F2정수64_단순형_공백은_0(g.Savg3)
	s.M총매도평단가_모음[3] = lib.F2정수64_단순형_공백은_0(g.Davg4)
	s.M총매수평단가_모음[3] = lib.F2정수64_단순형_공백은_0(g.Savg4)
	s.M총매도평단가_모음[4] = lib.F2정수64_단순형_공백은_0(g.Davg5)
	s.M총매수평단가_모음[4] = lib.F2정수64_단순형_공백은_0(g.Savg5)
	s.M외국계매도대금 = lib.F2정수64_단순형(g.Ftradmdval)
	s.M외국계매수대금 = lib.F2정수64_단순형_공백은_0(g.Ftradmsval)
	s.M외국계매도평단가 = lib.F2정수64_단순형_공백은_0(g.Ftradmdavg)
	s.M외국계매도평단가 = lib.F2정수64_단순형_공백은_0(g.Ftradmdavg)
	s.M외국계매수평단가 = lib.F2정수64_단순형_공백은_0(g.Ftradmsavg)
	s.M투자주의환기 = lib.F2문자열_EUC_KR_공백제거(g.Info5)
	s.M기업인수목적회사여부 = lib.F2참거짓(g.Spac_gubun, "N", false)
	s.M발행가격 = lib.F2정수64_단순형(g.Issueprice)
	s.M배분적용구분코드 = lib.F2문자열_EUC_KR(g.Alloc_gubun)
	s.M배분적용구분 = lib.F2문자열_EUC_KR(g.Alloc_text)
	s.M단기과열_VI발동 = lib.F2문자열_EUC_KR(g.Shterm_text)
	s.M정적VI상한가 = lib.F2정수64_단순형(g.Svi_uplmtprice)
	s.M정적VI하한가 = lib.F2정수64_단순형(g.Svi_dnlmtprice)
	s.M저유동성종목여부 = lib.F2참거짓(g.Low_lqdt_gu, 1, true)
	s.M이상급등종목여부 = lib.F2참거짓(g.Abnormal_rise_gu, 1, true)

	대차불가표시_문자열 := lib.F2문자열_EUC_KR_공백제거(g.Lend_text)
	switch 대차불가표시_문자열 {
	case "":
		s.M대차불가여부 = false
	case "대차불가":
		s.M대차불가여부 = true
	default:
		panic(lib.New에러("%v '대차불가표시_문자열' 예상하지 못한 값 : '%v'", s.M종목코드, 대차불가표시_문자열))
	}

	return s, nil
}

func NewCSPAT00600InBlock(질의값 *xing.S질의값_정상_주문) (g *CSPAT00600InBlock1) {
	g = new(CSPAT00600InBlock1)
	lib.F바이트_복사_문자열(g.AcntNo[:], 질의값.M계좌번호)
	lib.F바이트_복사_문자열(g.InptPwd[:], 질의값.M계좌_비밀번호)
	lib.F바이트_복사_문자열(g.IsuNo[:], 질의값.M종목코드)
	lib.F바이트_복사_정수(g.OrdQty[:], 질의값.M주문수량)
	lib.F바이트_복사_정수(g.OrdPrc[:], 질의값.M주문단가)
	lib.F바이트_복사_문자열(g.BnsTpCode[:], string(f2Xing매수매도(질의값.M매수_매도)))
	lib.F바이트_복사_문자열(g.OrdprcPtnCode[:], string(f2Xing호가유형(질의값.M호가유형)))
	lib.F바이트_복사_문자열(g.MgntrnCode[:], string(f2Xing신용거래_구분(질의값.M신용거래_구분)))

	// 대출일 : YYYYMMDD, 신용주문이 아닐 경우는 SPACE
	switch 질의값.M신용거래_구분 {
	case lib.P신용거래_해당없음:
		lib.F바이트_복사_문자열(g.LoanDt[:], "        ")
	default:
		lib.F바이트_복사_문자열(g.LoanDt[:], 질의값.M대출일.Format("20000102"))
	}

	lib.F바이트_복사_문자열(g.OrdCndiTpCode[:], string(f2Xing주문조건(질의값.M주문조건)))

	return g
}

func New현물_정상주문_응답1(tr *TR_DATA) (s *xing.S현물_정상_주문_응답1, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { s = nil }}.S실행()

	g := (*CSPAT00600OutBlockAll)(unsafe.Pointer(tr.Data)).OutBlock1

	if lib.F2문자열(g.LoanDt) == "00000000" {
		lib.F바이트_복사_문자열(g.LoanDt[:], "")
	}

	s = new(xing.S현물_정상_주문_응답1)
	s.M레코드_수량 = lib.F2정수_단순형(g.RecCnt)
	s.M계좌번호 = lib.F2문자열_공백제거(g.AcntNo)
	s.M계좌_비밀번호 = lib.F2문자열_공백제거(g.InptPwd)
	s.M종목코드 = lib.F2문자열_공백제거(g.IsuNo)
	s.M주문수량 = lib.F2정수64_단순형(g.OrdQty)
	s.M주문가격 = lib.F2정수64_단순형(g.OrdPrc)
	s.M매매구분 = 에러체크(f2매수매도(xing.T매수_매도(lib.F2문자열_공백제거(g.BnsTpCode)))).(lib.T매수_매도)
	s.M호가유형 = f2호가유형(xing.T호가유형(lib.F2문자열_공백제거(g.OrdprcPtnCode)))
	s.M프로그램_호가유형 = lib.F2문자열_공백제거(g.PrgmOrdprcPtnCode)
	s.M공매도_가능 = lib.F문자열_비교(g.StslAbleYn, "Y", true)
	s.M공매도_호가구분 = lib.F2문자열_공백제거(g.StslOrdprcTpCode)
	s.M통신매체_코드 = lib.F2문자열_공백제거(g.CommdaCode)
	s.M신용거래_구분 = f2신용거래_구분(xing.T신용거래_구분(lib.F2문자열_공백제거(g.MgntrnCode)))
	s.M대출일 = lib.F2포맷된_일자_단순형_공백은_초기값("20060102", g.LoanDt)
	s.M회원번호 = lib.F2문자열_공백제거(g.MbrNo)
	s.M주문조건_구분 = f2주문조건(xing.T주문조건(lib.F2문자열_공백제거(g.OrdCndiTpCode)))
	s.M전략코드 = lib.F2문자열_공백제거(g.StrtgCode)
	s.M그룹ID = lib.F2문자열_공백제거(g.GrpId)
	s.M주문회차 = lib.F2정수64_단순형(g.OrdSeqNo)
	s.M포트폴리오_번호 = lib.F2정수64_단순형(g.PtflNo)
	s.M트렌치_번호 = lib.F2정수64_단순형(g.TrchNo)
	s.M아이템_번호 = lib.F2정수64_단순형(g.ItemNo)
	s.M운용지시_번호 = lib.F2문자열_공백제거(g.OpDrtnNo)
	s.M유동성_공급자_여부 = lib.F문자열_비교(g.LpYn, "Y", true)
	s.M반대매매_구분 = lib.F2문자열_공백제거(g.CvrgTpCode)

	return s, nil
}

func New현물_정상주문_응답2(tr *TR_DATA) (s *xing.S현물_정상_주문_응답2, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { s = nil }}.S실행()

	g := (*CSPAT00600OutBlockAll)(unsafe.Pointer(tr.Data)).OutBlock2

	if lib.F2문자열_공백제거(g.OrdNo) == "" {	// 주문 에러발생시 공백 문자열이 수신됨.
		return nil, lib.New에러("New현물_정상주문_응답2() : 주문번호 생성 에러.")
	}

	s = new(xing.S현물_정상_주문_응답2)
	s.M레코드_수량 = lib.F2정수_단순형(g.RecCnt)
	s.M주문번호 = lib.F2정수64_단순형(g.OrdNo)

	if 시각_문자열 := lib.F2문자열_공백제거(g.OrdTime); 시각_문자열 != "" {
		시각_문자열 = lib.F문자열_삽입(lib.F2문자열_공백제거(g.OrdTime), ".", 6)
		s.M주문시각 = lib.F2금일_시각_단순형("150405.999999", 시각_문자열)
	} else {
		s.M주문시각 = time.Time{}
	}

	s.M주문시장_코드 = xing.T주문_시장구분(lib.F2정수_단순형(g.OrdMktCode))
	s.M주문유형_코드 = lib.F2문자열_공백제거(g.OrdPtnCode)
	s.M종목코드 = lib.F2문자열_공백제거(g.ShtnIsuNo)
	s.M관리사원_번호 = lib.F2문자열_공백제거(g.MgempNo)
	s.M주문금액 = lib.F2정수64_단순형(g.OrdAmt)
	s.M예비_주문번호 = lib.F2정수64_단순형_공백은_0(g.SpareOrdNo)
	s.M반대매매_일련번호 = lib.F2정수64_단순형_공백은_0(g.CvrgSeqno)
	s.M예약_주문번호 = lib.F2정수64_단순형_공백은_0(g.RsvOrdNo)
	s.M재사용_주문수량 = lib.F2정수64_단순형_공백은_0(g.RuseOrdQty)
	s.M현금_주문금액 = lib.F2정수64_단순형(g.MnyOrdAmt)
	s.M대용_주문금액 = lib.F2정수64_단순형(g.SubstOrdAmt)
	s.M재사용_주문금액 = lib.F2정수64_단순형(g.RuseOrdAmt)
	s.M계좌명 = lib.F2문자열_공백제거(g.AcntNm)
	s.M종목명 = lib.F2문자열_공백제거(g.IsuNm)

	if strings.HasPrefix(s.M종목코드, "A") {
		s.M종목코드 = s.M종목코드[1:]
	}

	return s, nil
}

func NewCSPAT00700InBlock(질의값 *xing.S질의값_정정_주문) (g *CSPAT00700InBlock1) {
	g = new(CSPAT00700InBlock1)
	lib.F바이트_복사_정수(g.OrgOrdNo[:], 질의값.M원주문번호)
	lib.F바이트_복사_문자열(g.AcntNo[:], 질의값.M계좌번호)
	lib.F바이트_복사_문자열(g.InptPwd[:], 질의값.M계좌_비밀번호)
	lib.F바이트_복사_문자열(g.IsuNo[:], 질의값.M종목코드)
	lib.F바이트_복사_정수(g.OrdQty[:], 질의값.M주문수량)
	lib.F바이트_복사_문자열(g.OrdprcPtnCode[:], string(f2Xing호가유형(질의값.M호가유형)))
	lib.F바이트_복사_문자열(g.OrdCndiTpCode[:], string(f2Xing주문조건(질의값.M주문조건)))
	lib.F바이트_복사_정수(g.OrdPrc[:], 질의값.M주문단가)

	return g
}

func New현물_정정주문_응답1(tr *TR_DATA) (s *xing.S현물_정정_주문_응답1, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { s = nil }}.S실행()

	g := (*CSPAT00700OutBlockAll)(unsafe.Pointer(tr.Data)).OutBlock1

	s = new(xing.S현물_정정_주문_응답1)
	s.M레코드_수량 = lib.F2정수_단순형(g.RecCnt)
	s.M원_주문번호 = lib.F2정수64_단순형(g.OrgOrdNo)
	s.M계좌번호 = lib.F2문자열_공백제거(g.AcntNo)
	s.M계좌_비밀번호 = lib.F2문자열_공백제거(g.InptPwd)
	s.M종목코드 = lib.F2문자열_공백제거(g.IsuNo)
	s.M주문수량 = lib.F2정수64_단순형(g.OrdQty)
	s.M호가유형 = f2호가유형(xing.T호가유형(lib.F2문자열_공백제거(g.OrdprcPtnCode)))
	s.M주문조건 = f2주문조건(xing.T주문조건(lib.F2문자열_공백제거(g.OrdCndiTpCode)))
	s.M주문가격 = lib.F2정수64_단순형(g.OrdPrc)
	s.M통신매체_코드 = lib.F2문자열_공백제거(g.CommdaCode)
	s.M전략코드 = lib.F2문자열_공백제거(g.StrtgCode)
	s.M그룹ID = lib.F2문자열_공백제거(g.GrpId)
	s.M주문회차 = lib.F2정수64_단순형(g.OrdSeqNo)
	s.M포트폴리오_번호 = lib.F2정수64_단순형(g.PtflNo)
	s.M바스켓_번호 = lib.F2정수64_단순형(g.BskNo)
	s.M트렌치_번호 = lib.F2정수64_단순형(g.TrchNo)
	s.M아이템_번호 = lib.F2정수64_단순형(g.ItemNo)

	return s, nil
}

func New현물_정정주문_응답2(tr *TR_DATA) (s *xing.S현물_정정_주문_응답2, 에러 error) {
	defer func() {
		if r := recover(); r != nil {
			s = nil
			에러 = lib.New에러(r)
		}}()

	g := (*CSPAT00700OutBlockAll)(unsafe.Pointer(tr.Data)).OutBlock2

	if lib.F2문자열_공백제거(g.OrdNo) == "" {	// 주문 에러발생시 공백 문자열이 수신됨.
		// 에러가 너무 장황해서 lib.New에러() 대신에 errors.New()로 대체함.
		return nil, errors.New("New현물_정정주문_응답2() : 주문번호 생성 에러.")
	}

	시각_문자열 := lib.F2문자열_공백제거(g.OrdTime)
	if 시각_문자열 != "" {
		시각_문자열 = lib.F문자열_삽입(lib.F2문자열_공백제거(g.OrdTime), ".", 6)
	}

	if lib.F2문자열(g.LoanDt) == "00000000" {
		lib.F바이트_복사_문자열(g.LoanDt[:], "")
	}

	s = new(xing.S현물_정정_주문_응답2)
	s.M레코드_수량 = lib.F2정수_단순형(g.RecCnt)
	s.M주문번호 = lib.F2정수64_단순형(g.OrdNo)
	s.M모_주문번호 = lib.F2정수64_단순형_공백은_0(g.PrntOrdNo)
	s.M주문시각 = lib.F2금일_시각_단순형_공백은_초기값("150405.999999", 시각_문자열)
	s.M주문시장_코드 = xing.T주문_시장구분(lib.F2정수64_단순형_공백은_0(g.OrdMktCode))
	s.M주문유형_코드 = lib.F2문자열_공백제거(g.OrdPtnCode)
	s.M종목코드 = lib.F2문자열_공백제거(g.ShtnIsuNo) // 단축종목번호
	s.M공매도_호가구분 = lib.F2문자열_공백제거(g.StslOrdprcTpCode)
	s.M공매도_가능 = lib.F문자열_비교(g.StslAbleYn, "Y", true)
	s.M신용거래_구분 = f2신용거래_구분(xing.T신용거래_구분(lib.F2문자열_공백제거(g.MgntrnCode)))
	s.M대출일 = lib.F2포맷된_일자_단순형_공백은_초기값("20060102", g.LoanDt)
	s.M반대매매주문_구분 = lib.F2문자열_공백제거(g.CvrgOrdTp)
	s.M관리사원_번호 = lib.F2문자열_공백제거(g.MgempNo)
	s.M주문금액 = lib.F2정수64_단순형_공백은_0(g.OrdAmt)
	s.M매매구분 = 에러체크(f2매수매도(xing.T매수_매도(lib.F2문자열_공백제거(g.BnsTpCode)))).(lib.T매수_매도)
	s.M예비_주문번호 = lib.F2정수64_단순형_공백은_0(g.SpareOrdNo)
	s.M반대매매_일련번호 = lib.F2정수64_단순형_공백은_0(g.CvrgSeqno)
	s.M예약_주문번호 = lib.F2정수64_단순형_공백은_0(g.RsvOrdNo)
	s.M현금_주문금액 = lib.F2정수64_단순형(g.MnyOrdAmt)
	s.M대용_주문금액 = lib.F2정수64_단순형(g.SubstOrdAmt)
	s.M재사용_주문금액 = lib.F2정수64_단순형(g.RuseOrdAmt)
	s.M계좌명 = lib.F2문자열_공백제거(g.AcntNm)
	s.M종목명 = lib.F2문자열_공백제거(g.IsuNm)

	if strings.HasPrefix(s.M종목코드, "A") {
		s.M종목코드 = s.M종목코드[1:]
	}

	return s, nil
}

func NewCSPAT00800InBlock(질의값 *xing.S질의값_취소_주문) (g *CSPAT00800InBlock1) {
	g = new(CSPAT00800InBlock1)
	lib.F바이트_복사_정수(g.OrgOrdNo[:], 질의값.M원주문번호)
	lib.F바이트_복사_문자열(g.AcntNo[:], 질의값.M계좌번호)
	lib.F바이트_복사_문자열(g.InptPwd[:], 질의값.M계좌_비밀번호)
	lib.F바이트_복사_문자열(g.IsuNo[:], 질의값.M종목코드)
	lib.F바이트_복사_정수(g.OrdQty[:], 질의값.M주문수량)

	return g
}

func New현물_취소주문_응답1(tr *TR_DATA) (s *xing.S현물_취소_주문_응답1, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { s = nil }}.S실행()

	g := (*CSPAT00800OutBlockAll)(unsafe.Pointer(tr.Data)).OutBlock1

	s = new(xing.S현물_취소_주문_응답1)
	s.M레코드_수량 = lib.F2정수_단순형(g.RecCnt)
	s.M원_주문번호 = lib.F2정수64_단순형(g.OrgOrdNo)
	s.M계좌번호 = lib.F2문자열_공백제거(g.AcntNo)
	s.M계좌_비밀번호 = lib.F2문자열_공백제거(g.InptPwd)
	s.M종목코드 = lib.F2문자열_공백제거(g.IsuNo)
	s.M주문수량 = lib.F2정수64_단순형(g.OrdQty)
	s.M통신매체_코드 = lib.F2문자열_공백제거(g.CommdaCode)
	s.M그룹ID = lib.F2문자열_공백제거(g.GrpId)
	s.M전략코드 = lib.F2문자열_공백제거(g.StrtgCode)
	s.M주문회차 = lib.F2정수64_단순형(g.OrdSeqNo)
	s.M포트폴리오_번호 = lib.F2정수64_단순형(g.PtflNo)
	s.M바스켓_번호 = lib.F2정수64_단순형(g.BskNo)
	s.M트렌치_번호 = lib.F2정수64_단순형(g.TrchNo)
	s.M아이템_번호 = lib.F2정수64_단순형(g.ItemNo)

	return s, nil
}

func New현물_취소주문_응답2(tr *TR_DATA) (s *xing.S현물_취소_주문_응답2, 에러 error) {
	defer func() {
		if r := recover(); r != nil {
			s = nil
			에러 = lib.New에러(r)
		}}()
	
	g := (*CSPAT00800OutBlockAll)(unsafe.Pointer(tr.Data)).OutBlock2

	if lib.F2문자열_공백제거(g.OrdNo) == "" {	// 주문 에러발생시 공백 문자열이 수신됨.
		// 에러가 너무 장황해서 lib.New에러() 대신에 errors.New()로 대체함.
		return nil, errors.New("New현물_취소주문_응답2() : 주문번호 생성 에러.")
	}

	시각_문자열 := lib.F2문자열_공백제거(g.OrdTime)
	if 시각_문자열 != "" {
		시각_문자열 = lib.F문자열_삽입(lib.F2문자열_공백제거(g.OrdTime), ".", 6)
	}

	if lib.F2문자열(g.LoanDt) == "00000000" {
		lib.F바이트_복사_문자열(g.LoanDt[:], "")
	}

	s = new(xing.S현물_취소_주문_응답2)
	s.M레코드_수량 = lib.F2정수_단순형(g.RecCnt)
	s.M주문번호 = lib.F2정수64_단순형(g.OrdNo)
	s.M모_주문번호 = lib.F2정수64_단순형_공백은_0(g.PrntOrdNo)
	s.M주문시각 = lib.F2금일_시각_단순형("150405.999999", 시각_문자열)
	s.M주문시장_코드 = xing.T주문_시장구분(lib.F2정수_단순형(g.OrdMktCode))
	s.M주문유형_코드 = lib.F2문자열_공백제거(g.OrdPtnCode)
	s.M종목코드 = lib.F2문자열_공백제거(g.ShtnIsuNo)
	s.M공매도_호가구분 = lib.F2문자열_공백제거(g.StslOrdprcTpCode)
	s.M공매도_가능 = lib.F문자열_비교(g.StslAbleYn, "Y", true)
	s.M신용거래_코드 = f2신용거래_구분(xing.T신용거래_구분(lib.F2문자열_공백제거(g.MgntrnCode)))
	s.M대출일 = lib.F2포맷된_일자_단순형_공백은_초기값("20060102", g.LoanDt)
	s.M반대매매주문_구분 = lib.F2문자열_공백제거(g.CvrgOrdTp)
	s.M유동성공급자_여부 = lib.F문자열_비교(g.LpYn, "Y", true)
	s.M관리사원_번호 = lib.F2문자열_공백제거(g.MgempNo)
	s.M예비_주문번호 = lib.F2정수64_단순형_공백은_0(g.SpareOrdNo)
	s.M반대매매_일련번호 = lib.F2정수64_단순형_공백은_0(g.CvrgSeqno)
	s.M예약_주문번호 = lib.F2정수64_단순형_공백은_0(g.RsvOrdNo)
	s.M계좌명 = lib.F2문자열_공백제거(g.AcntNm)
	s.M종목명 = lib.F2문자열_공백제거(g.IsuNm)

	if strings.HasPrefix(s.M종목코드, "A") {
		s.M종목코드 = s.M종목코드[1:]
	}

	return s, nil
}

func NewT1902InBlock(질의값 *xing.S질의값_단일종목_연속키) (g *T1902InBlock) {
	g = new(T1902InBlock)
	lib.F바이트_복사_문자열(g.ShCode[:], 질의값.M종목코드)

	if lib.F2문자열_공백제거(질의값.M연속키) == "" {
		lib.F바이트_배열_공백문자열_채움(g.Time[:])
	} else {
		lib.F바이트_복사_문자열(g.Time[:], 질의값.M연속키)
	}

	return g
}

func NewETF시간별_추이_응답_헤더(tr *TR_DATA) (s *xing.S_ETF시간별_추이_응답_헤더, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { s = nil }}.S실행()

	g := (*T1902OutBlock)(unsafe.Pointer(tr.Data))

	s = new(xing.S_ETF시간별_추이_응답_헤더)
	s.M연속키 = lib.F2문자열_공백제거(g.Time)
	s.M종목명 = lib.F2문자열_EUC_KR_공백제거(g.HName)
	s.M업종지수명 = lib.F2문자열_EUC_KR_공백제거(g.UpName)

	return s, nil
}

func NewETF시간별_추이_응답_반복값_모음(tr *TR_DATA) (값 *xing.S_ETF시간별_추이_응답_반복값_모음, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	// C배열 -> Go슬라이스 : https://github.com/golang/go/wiki/cgo : Turning C arrays into Go slices
	배열_길이 := int(tr.DataLength) / 크기T1902OutBlock1
	lib.F조건부_패닉(배열_길이 >= (1<<20), "미리 확보하는 메모리가 부족함.")

	g_모음 := (*[1 << 20]T1902OutBlock1)(unsafe.Pointer(tr.Data))[:배열_길이:배열_길이]
	배열 := make([]*xing.S_ETF시간별_추이_응답_반복값, len(g_모음), len(g_모음))

	당일값 := 당일.G값()

	for i, g := range g_모음 {
		s := new(xing.S_ETF시간별_추이_응답_반복값)

		if lib.F2문자열_EUC_KR(g.Time) == "장:마:감" {
			s.M시각 = lib.F2일자별_시각_단순형(당일값, "15:04:05", g_모음[i+1].Time).Add(lib.P10초)
		} else {
			s.M시각 = lib.F2일자별_시각_단순형(당일값, "15:04:05", g.Time)
		}

		s.M현재가 = lib.F2정수64_단순형(g.Price)
		s.M전일대비구분 = xing.T전일대비_구분(lib.F2정수_단순형(g.Sign))
		s.M전일대비등락폭 = s.M전일대비구분.G부호보정_정수64(lib.F2정수64_단순형(g.Change))
		s.M누적_거래량 = lib.F2실수_단순형(g.Volume)
		s.M현재가_NAV_차이 = lib.F2실수_단순형(g.NavDiff)
		s.NAV = lib.F2실수_단순형(g.Nav)
		s.NAV전일대비등락폭 = lib.F2실수_단순형(g.NavChange)
		s.M추적오차 = lib.F2실수_단순형(g.Crate)
		s.M괴리율 = lib.F2실수_단순형(g.Grate)
		s.M지수 = lib.F2실수_단순형(g.Jisu)
		s.M지수_전일대비등락폭 = lib.F2실수_단순형(g.JiChange)
		s.M지수_전일대비등락율 = lib.F2실수_단순형(g.JiRate)

		if uint8(g.X_jichange) == 160 && s.M지수_전일대비등락폭 > 0 {
			s.M지수_전일대비등락폭 = -1 * s.M지수_전일대비등락폭
		}

		배열[i] = s
	}

	값 = new(xing.S_ETF시간별_추이_응답_반복값_모음)
	값.M배열 = 배열

	return 값, nil
}

func NewT1305InBlock(질의값 *xing.S질의값_현물_기간별_조회) (g *T1305InBlock) {
	g = new(T1305InBlock)
	lib.F바이트_복사_문자열(g.Shcode[:], 질의값.M종목코드)
	lib.F바이트_복사_문자열(g.Dwmcode[:], lib.F2문자열(uint8(질의값.M일주월_구분)))
	lib.F바이트_복사_문자열(g.Date[:], 질의값.M연속키)
	lib.F바이트_복사_문자열(g.Idx[:], "    ") // 정수형인데, 사용안함(Space)으로 표시됨.
	lib.F바이트_복사_정수(g.Cnt[:], 질의값.M수량)

	if lib.F2문자열_공백제거(질의값.M연속키) == "" {
		lib.F바이트_복사_문자열(g.Date[:], "       ")
	}

	return g
}

func New현물_기간별_조회_응답_헤더(tr *TR_DATA) (값 *xing.S현물_기간별_조회_응답_헤더, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	g := (*T1305OutBlock)(unsafe.Pointer(tr.Data))

	값 = new(xing.S현물_기간별_조회_응답_헤더)
	값.M수량 = lib.F2정수64_단순형(g.Cnt)
	값.M연속키 = lib.F2문자열_공백제거(g.Date)

	return 값, nil
}

func New현물_기간별_조회_응답_반복값_모음(tr *TR_DATA) (값 *xing.S현물_기간별_조회_응답_반복값_모음, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	// C배열 -> Go슬라이스 : https://github.com/golang/go/wiki/cgo : Turning C arrays into Go slices
	배열_길이 := int(tr.DataLength) / 크기T1305OutBlock1
	lib.F조건부_패닉(배열_길이 >= (1<<20), "미리 확보하는 메모리가 부족함.")

	g_모음 := (*[1 << 20]T1305OutBlock1)(unsafe.Pointer(tr.Data))[:배열_길이:배열_길이]

	값 = new(xing.S현물_기간별_조회_응답_반복값_모음)
	값.M배열 = make([]*xing.S현물_기간별_조회_응답_반복값, 배열_길이, 배열_길이)

	for i, g := range g_모음 {
		일자_문자열_원본 := lib.F2문자열(g.Date)
		버퍼 := new(bytes.Buffer)
		버퍼.WriteString(일자_문자열_원본[0:4])
		버퍼.WriteString("/")
		버퍼.WriteString(일자_문자열_원본[4:6])
		버퍼.WriteString("/")
		버퍼.WriteString(일자_문자열_원본[6:])
		일자_문자열 := 버퍼.String()

		s := new(xing.S현물_기간별_조회_응답_반복값)
		s.M종목코드 = lib.F2문자열(g.Shcode)
		s.M일자 = lib.F2포맷된_일자_단순형("2006/01/02", 일자_문자열)
		s.M시가 = lib.F2정수64_단순형(g.Open)
		s.M고가 = lib.F2정수64_단순형(g.High)
		s.M저가 = lib.F2정수64_단순형(g.Low)
		s.M종가 = lib.F2정수64_단순형(g.Close)

		if 전일대비_구분값, 에러 := lib.F2정수64(g.Sign); 에러 == nil {
			s.M전일대비구분 = xing.T전일대비_구분(전일대비_구분값)
		} else if lib.F2문자열_공백제거(g.Sign) == "" &&
			lib.F2정수64_단순형(g.Change) == 0 && lib.F2실수_단순형(g.Diff) == 0.0 {
			s.M전일대비구분 = xing.P구분_보합
		} else {
			lib.F문자열_출력("일자 : '%v', 잘못된 전일구분. '%v'", s.M일자, lib.F2문자열(g.Sign))
			s.M전일대비구분 = xing.T전일대비_구분(0)
		}

		s.M전일대비등락폭 = s.M전일대비구분.G부호보정_정수64(lib.F2정수64_단순형(g.Change))
		s.M전일대비등락율 = s.M전일대비구분.G부호보정_실수64(lib.F2실수_단순형(g.Diff))
		s.M시가대비구분 = xing.T전일대비_구분(lib.F2정수64_단순형(g.O_sign))
		s.M시가대비등락폭 = s.M시가대비구분.G부호보정_정수64(lib.F2정수64_단순형(g.O_change))
		s.M시가대비등락율 = s.M시가대비구분.G부호보정_실수64(lib.F2실수_단순형(g.O_diff))
		s.M고가대비구분 = xing.T전일대비_구분(lib.F2정수64_단순형(g.H_sign))
		s.M고가대비등락폭 = s.M고가대비구분.G부호보정_정수64(lib.F2정수64_단순형(g.H_change))
		s.M고가대비등락율 = s.M고가대비구분.G부호보정_실수64(lib.F2실수_단순형(g.H_diff))
		s.M저가대비구분 = xing.T전일대비_구분(lib.F2정수64_단순형(g.L_sign))
		s.M저가대비등락폭 = s.M저가대비구분.G부호보정_정수64(lib.F2정수64_단순형(g.L_change))
		s.M저가대비등락율 = s.M저가대비구분.G부호보정_실수64(lib.F2실수_단순형(g.L_diff))
		s.M누적거래량 = lib.F2정수64_단순형(g.Volume)
		s.M누적거래대금_백만 = lib.F2정수64_단순형(g.Value)
		s.M거래_증가율 = lib.F2실수_단순형(g.Diff_vol)
		s.M체결강도 = lib.F2실수_단순형(g.Chdegree)
		s.M소진율 = lib.F2실수_단순형_공백은_0(g.Sojinrate)
		s.M회전율 = lib.F2실수_단순형(g.Changerate)
		s.M외국인_순매수 = lib.F2정수64_단순형_공백은_0(g.Fpvolume)
		s.M기관_순매수 = lib.F2정수64_단순형_공백은_0(g.Covolume)
		s.M개인_순매수 = lib.F2정수64_단순형_공백은_0(g.Ppvolume)
		s.M시가총액_백만 = lib.F2정수64_단순형(g.Marketcap)

		값.M배열[i] = s
	}

	return 값, nil
}

func NewT1310InBlock(질의값 *xing.S질의값_현물_전일당일_분틱_조회) (g *T1310InBlock) {
	g = new(T1310InBlock)
	lib.F바이트_복사_문자열(g.Daygb[:], strconv.Itoa(int(질의값.M당일전일구분)))
	lib.F바이트_복사_문자열(g.Timegb[:], strconv.Itoa(int(질의값.M분틱구분)))
	lib.F바이트_복사_문자열(g.Shcode[:], 질의값.M종목코드)
	lib.F바이트_복사_문자열(g.Endtime[:], strings.Replace(질의값.M종료시각.Format("15:04"), ":", "", -1))
	lib.F바이트_복사_문자열(g.Time[:], 질의값.M연속키)

	if lib.F2문자열_공백제거(질의값.M연속키) == "" {
		lib.F바이트_복사_문자열(g.Time[:], "          ")
	}

	return g
}

func New현물_당일전일분틱조회_응답_헤더(tr *TR_DATA) (값 *xing.S현물_전일당일분틱조회_응답_헤더, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	g := (*T1310OutBlock)(unsafe.Pointer(tr.Data))

	값 = new(xing.S현물_전일당일분틱조회_응답_헤더)
	값.M연속키 = lib.F2문자열(g.Time)

	return 값, nil
}

func New현물_당일전일분틱조회_응답_반복값_모음(tr *TR_DATA) (값 *xing.S현물_전일당일분틱조회_응답_반복값_모음, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	대기_항목 := 콜백_대기_저장소.G대기_항목(int(tr.RequestID))
	lib.F조건부_패닉(대기_항목 == nil, "대기 항목이 존재하지 않음.")
	lib.F조건부_패닉(대기_항목.TR코드 != xing.TR현물_당일_전일_분틱_조회 && 대기_항목.TR코드 != "",
		"예상과 다른 TR코드 : '%v' '%v'", xing.TR현물_당일_전일_분틱_조회, 대기_항목.TR코드)

	당일전일_구분, ok := 대기_항목.M값.(xing.T당일전일_구분)
	lib.F조건부_패닉(!ok, "예상과 다른 자료형 : '%T'", 대기_항목.M값)
	lib.F조건부_패닉(당일전일_구분 != xing.P당일전일구분_전일 && 당일전일_구분 != xing.P당일전일구분_당일,
		"예상과 다른 당일전일 구분값 : '%v'", int(당일전일_구분))

	// C배열 -> Go슬라이스 : https://github.com/golang/go/wiki/cgo : Turning C arrays into Go slices
	배열_길이 := int(tr.DataLength) / 크기T1310OutBlock1
	lib.F조건부_패닉(배열_길이 >= (1<<20), "미리 확보하는 메모리가 부족함.")

	g_모음 := (*[1 << 20]T1310OutBlock1)(unsafe.Pointer(tr.Data))[:배열_길이:배열_길이]

	값 = new(xing.S현물_전일당일분틱조회_응답_반복값_모음)
	값.M배열 = make([]*xing.S현물_전일당일분틱조회_응답_반복값, 배열_길이, 배열_길이)

	var 일자 time.Time
	if 당일전일_구분 == xing.P당일전일구분_전일 {
		일자 = 전일.G값()
	} else {
		일자 = 당일.G값()
	}

	for i, g := range g_모음 {
		시각_문자열_원본 := lib.F2문자열(g.Chetime[:6])

		버퍼 := new(bytes.Buffer)
		버퍼.WriteString(시각_문자열_원본[0:2])
		버퍼.WriteString(":")
		버퍼.WriteString(시각_문자열_원본[2:4])
		버퍼.WriteString(":")
		버퍼.WriteString(시각_문자열_원본[4:])
		시각_문자열 := 버퍼.String()

		s := new(xing.S현물_전일당일분틱조회_응답_반복값)
		s.M시각 = lib.F2일자별_시각_단순형(일자, "15:04:05", 시각_문자열)
		s.M현재가 = lib.F2정수64_단순형(g.Price[:])
		s.M전일대비구분 = xing.T전일대비_구분(lib.F2정수64_단순형(g.Sign))
		s.M전일대비등락폭 = s.M전일대비구분.G부호보정_정수64(lib.F2정수64_단순형(g.Change))
		s.M전일대비등락율 = s.M전일대비구분.G부호보정_실수64(lib.F2실수_단순형(g.Diff))
		s.M체결수량 = lib.F2정수64_단순형(g.Cvolume)
		s.M체결강도 = lib.F2실수_단순형(g.Chdegree)
		s.M거래량 = lib.F2정수64_단순형(g.Volume)
		s.M매도체결수량 = lib.F2정수64_단순형(g.Mdvolume)
		s.M매도체결건수 = lib.F2정수64_단순형(g.Mdchecnt)
		s.M매수체결수량 = lib.F2정수64_단순형(g.Msvolume)
		s.M매수체결건수 = lib.F2정수64_단순형(g.Mschecnt)
		s.M순체결량 = lib.F2정수64_단순형(g.Revolume)
		s.M순체결건수 = lib.F2정수64_단순형(g.Rechecnt)

		값.M배열[i] = s
	}

	return 값, nil
}

func NewT8428InBlock(질의값 *xing.S질의값_증시주변자금추이) (g *T8428InBlock) {
	시장구분_문자열 := ""
	switch 질의값.M시장구분 {
	case lib.P시장구분_코스피:
		시장구분_문자열 = "001"
	case lib.P시장구분_코스닥:
		시장구분_문자열 = "301"
	default:
		panic(lib.New에러("예상하지 못한 시장구분 값 : '%v'", 질의값.M시장구분))
	}

	g = new(T8428InBlock)
	lib.F바이트_복사_문자열(g.KeyDate[:], 질의값.M연속키)
	lib.F바이트_복사_문자열(g.Upcode[:], 시장구분_문자열)
	lib.F바이트_복사_정수(g.Cnt[:], 질의값.M수량)

	if lib.F2문자열_공백제거(질의값.M연속키) == "" {
		lib.F바이트_복사_문자열(g.KeyDate[:], "        ")
	}

	return g
}

func New증시주변자금추이_응답_헤더(tr *TR_DATA) (값 *xing.S증시주변자금추이_응답_헤더, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	g := (*T8428OutBlock)(unsafe.Pointer(tr.Data))

	값 = new(xing.S증시주변자금추이_응답_헤더)
	값.M연속키 = lib.F2문자열(g.Date)
	값.M인덱스 = lib.F2정수64_단순형(g.Idx)

	return 값, nil
}

func New증시주변자금추이_응답_반복값_모음(tr *TR_DATA) (값 *xing.S증시주변자금추이_응답_반복값_모음, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	// C배열 -> Go슬라이스 : https://github.com/golang/go/wiki/cgo : Turning C arrays into Go slices
	배열_길이 := int(tr.DataLength) / 크기T8428OutBlock1
	lib.F조건부_패닉(배열_길이 >= (1<<20), "미리 확보하는 메모리가 부족함.")

	g_모음 := (*[1 << 20]T8428OutBlock1)(unsafe.Pointer(tr.Data))[:배열_길이:배열_길이]

	값 = new(xing.S증시주변자금추이_응답_반복값_모음)
	값.M배열 = make([]*xing.S증시주변자금추이_응답_반복값, 배열_길이, 배열_길이)

	for i, g := range g_모음 {
		시각_문자열_원본 := lib.F2문자열(g.Date)

		lib.F조건부_패닉(len(시각_문자열_원본) != 8, "예상과 다른 시각 문자열 길이 : '%v', '%v'",
			len(시각_문자열_원본), 시각_문자열_원본)

		버퍼 := new(bytes.Buffer)
		버퍼.WriteString(시각_문자열_원본[0:4])
		버퍼.WriteString("/")
		버퍼.WriteString(시각_문자열_원본[4:6])
		버퍼.WriteString("/")
		버퍼.WriteString(시각_문자열_원본[6:])
		시각_문자열 := 버퍼.String()

		s := new(xing.S증시주변자금추이_응답_반복값)
		s.M일자 = lib.F2포맷된_시각_단순형("2006/01/02", 시각_문자열)
		s.M지수 = lib.F2실수_단순형(g.Jisu)
		s.M전일대비_구분 = xing.T전일대비_구분(lib.F2정수64_단순형(g.Sign))
		s.M전일대비_등락폭 = s.M전일대비_구분.G부호보정_실수64(lib.F2실수_단순형(g.Change))
		s.M전일대비_등락율 = s.M전일대비_구분.G부호보정_실수64(lib.F2실수_단순형(g.Diff))
		s.M거래량 = lib.F2정수64_단순형(g.Volume)
		s.M고객예탁금_억 = lib.F2정수64_단순형(g.Custmoney)
		s.M예탁증감_억 = lib.F2정수64_단순형(g.Yecha)

		if strings.Contains(strings.ToLower(lib.F2문자열(g.Vol)), "inf") {
			s.M회전율 = math.Inf(1)
		} else {
			s.M회전율 = lib.F2실수_단순형(g.Vol)
		}

		s.M미수금_억 = lib.F2정수64_단순형(g.Outmoney)
		s.M신용잔고_억 = lib.F2정수64_단순형(g.Trjango)
		s.M선물예수금_억 = lib.F2정수64_단순형(g.Futymoney)
		s.M주식형_억 = lib.F2정수64_단순형(g.Stkmoney)
		s.M혼합형_주식_억 = lib.F2정수64_단순형(g.Mstkmoney)
		s.M혼합형_채권_억 = lib.F2정수64_단순형(g.Mbndmoney)
		s.M채권형_억 = lib.F2정수64_단순형(g.Bndmoney)
		s.MMF_억 = lib.F2정수64_단순형(g.Mmfmoney)

		값.M배열[i] = s
	}

	return 값, nil
}

func NewT8436InBlock(질의값 *lib.S질의값_문자열) (g *T8436InBlock) {
	lib.F조건부_패닉(질의값.M문자열 != "0" && 질의값.M문자열 != "1" && 질의값.M문자열 != "2",
		"예상하지 못한 구분값 : '%v'", 질의값.M문자열)

	g = new(T8436InBlock)
	lib.F바이트_복사_문자열(g.Gubun[:], 질의값.M문자열)

	return g
}

func New주식종목조회_응답_반복값_모음(tr *TR_DATA) (값 *xing.S현물_종목조회_응답_반복값_모음, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	// C배열 -> Go슬라이스 : https://github.com/golang/go/wiki/cgo : Turning C arrays into Go slices
	배열_길이 := int(tr.DataLength) / 크기T8436OutBlock
	lib.F조건부_패닉(배열_길이 >= (1<<20), "미리 확보하는 메모리가 부족함.")

	g_모음 := (*[1 << 20]T8436OutBlock)(unsafe.Pointer(tr.Data))[:배열_길이:배열_길이]

	값 = new(xing.S현물_종목조회_응답_반복값_모음)
	값.M배열 = make([]*xing.S현물_종목조회_응답_반복값, 배열_길이, 배열_길이)

	for i, g := range g_모음 {
		s := new(xing.S현물_종목조회_응답_반복값)
		s.M종목명 = lib.F2문자열_EUC_KR_공백제거(g.HName)
		s.M종목코드 = lib.F2문자열_공백제거(g.ShCode)
		s.M주문수량단위 = lib.F2정수_단순형(g.MeMeDan)
		s.M상한가 = lib.F2정수64_단순형(g.UpLmtPrice)
		s.M하한가 = lib.F2정수64_단순형(g.DnLmtPrice)
		s.M전일가 = lib.F2정수64_단순형(g.JnilClose)
		s.M기준가 = lib.F2정수64_단순형(g.RecPrice)
		s.M증권그룹 = xing.T증권그룹(lib.F2정수_단순형(g.Bu12Gubun))
		s.M기업인수목적회사여부 = lib.F2참거짓(lib.F2문자열(g.SpacGubun), "Y", true)

		ETF구분 := lib.F2문자열_공백제거(g.EtfGubun)
		시장구분 := lib.F2문자열_공백제거(g.Gubun)

		switch {
		case ETF구분 == "1":
			s.M시장구분 = lib.P시장구분_ETF
		case ETF구분 == "2":
			s.M시장구분 = lib.P시장구분_ETN
		case 시장구분 == "1":
			s.M시장구분 = lib.P시장구분_코스피
		case 시장구분 == "2":
			s.M시장구분 = lib.P시장구분_코스닥
		default:
			panic(lib.New에러("예상하지 못한 경우 : '%v', '%v'", ETF구분, 시장구분))
		}

		값.M배열[i] = s

		switch {
		case s.M증권그룹 == xing.P증권그룹_상장지수펀드_ETF && s.M시장구분 == lib.P시장구분_ETN,
			s.M증권그룹 == xing.P증권그룹_ETN && s.M시장구분 == lib.P시장구분_ETF:
			lib.F문자열_출력(
				"종목코드 : '%v', 증권그룹 : '%v', 시장구분 : '%v'",
				s.M종목코드, s.M증권그룹, s.M시장구분)
		}
	}

	return 값, nil
}
