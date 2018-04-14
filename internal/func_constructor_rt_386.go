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

	"unsafe"
)

func f실시간_데이터_해석(rt *REALTIME_DATA) (값 interface{}, 에러 error) {
	RT코드 := l.F2문자열(rt.TrCode)

	switch RT코드 {
	case xt.RT현물주문_접수:
		New현물주문_접수(rt)
	case xt.RT현물주문_체결:
		New현물주문_체결(rt)
	case xt.RT현물주문_정정:
		New현물주문_정정(rt)
	case xt.RT현물주문_취소:
		New현물주문_취소(rt)
	case xt.RT현물주문_거부:
		New현물주문_거부(rt)
	case xt.RT코스피_호가_잔량:
		return New코스피_호가_잔량(rt)
	case xt.RT코스피_시간외_호가_잔량:
		return New코스피_시간외_호가_잔량(rt)
	case xt.RT코스피_체결:
		return New코스피_체결(rt)
	case xt.RT코스피_예상_체결:
		return New코스피_예상_체결(rt)
	case xt.RT코스피_ETF_NAV:
		return New코스피_ETF_NAV(rt)
	case xt.RT주식_VI발동해제:
		return New주식_VI발동해제(rt)
	case xt.RT시간외_단일가VI발동해제:
		return New시간외_단일가VI발동해제(rt)
	case xt.RT장_운영정보:
		return New장_운영정보(rt)
	case xt.RT코스닥_체결, xt.RT코스닥_예상_체결,
		xt.RT코스피_거래원, xt.RT코스닥_거래원,
		xt.RT코스피_기세, xt.RT코스닥_LP호가,
		xt.RT코스닥_호가잔량, xt.RT코스닥_시간외_호가잔량,
		xt.RT지수, xt.RT예상지수,
		xt.RT실시간_뉴스_제목_패킷,
		xt.RT업종별_투자자별_매매_현황:
		return nil, l.New에러("미구현 RT코드 : '%v'", RT코드)
	}

	return nil, l.New에러("예상하지 못한 RT코드 : '%v'", RT코드)
}

// SC0
func New현물주문_접수(rt *REALTIME_DATA) (값 *xt.S주문_응답, 에러 error) {
	g := (*SC0_OutBlock)(unsafe.Pointer(rt.Data))

	시각_문자열 := l.F2문자열(g.Ordtm)
	시각_문자열 = 시각_문자열[:6] + "." + 시각_문자열[7:]
	시각 := l.F2금일_시각_단순형("150405.999", 시각_문자열)

	종목코드 := l.F2문자열_공백제거(g.Shtcode)
	종목코드 = 종목코드[1:] // 맨 앞의 'A' 제거

	s := new(xt.S주문_응답)
	s.M주문번호 = l.F2정수64_단순형(g.Ordno)
	s.M원_주문번호 = l.F2정수64_단순형_공백은_0(g.Orgordno)
	s.RT코드 = l.F2문자열(rt.TrCode)
	s.M종목코드 = 종목코드
	s.M수량 = l.F2정수64_단순형(g.Ordqty)
	s.M가격 = l.F2정수64_단순형(g.Ordprice)
	s.M잔량 = 0
	s.M시각 = 시각

	return s, nil
}

// SC1
func New현물주문_체결(rt *REALTIME_DATA) (값 *xt.S주문_응답, 에러 error) {
	g := (*SC1_OutBlock)(unsafe.Pointer(rt.Data))

	시각_문자열 := l.F2문자열(g.Exectime)
	시각_문자열 = 시각_문자열[:6] + "." + 시각_문자열[7:]
	시각 := l.F2금일_시각_단순형("150405.999", 시각_문자열)

	종목코드 := l.F2문자열_공백제거(g.ShtnIsuno)
	종목코드 = 종목코드[1:] // 맨 앞의 'A' 제거

	s := new(xt.S주문_응답)
	s.M주문번호 = l.F2정수64_단순형(g.Ordno)
	s.M원_주문번호 = l.F2정수64_단순형_공백은_0(g.Orgordno)
	s.RT코드 = l.F2문자열(rt.TrCode)
	s.M종목코드 = 종목코드
	s.M수량 = l.F2정수64_단순형(g.Execqty)
	s.M가격 = l.F2정수64_단순형(g.Execprc)
	s.M잔량 = l.F2정수64_단순형(g.Unercqty)
	s.M시각 = 시각

	return s, nil
}

func New현물주문_정정(rt *REALTIME_DATA) (값 *xt.S주문_응답, 에러 error) {
	g := (*SC2_OutBlock)(unsafe.Pointer(rt.Data))

	시각_문자열 := l.F2문자열(g.Exectime)
	시각_문자열 = 시각_문자열[:6] + "." + 시각_문자열[7:]
	시각 := l.F2금일_시각_단순형("150405.999", 시각_문자열)

	종목코드 := l.F2문자열_공백제거(g.ShtnIsuno)
	종목코드 = 종목코드[1:] // 맨 앞의 'A' 제거

	s := new(xt.S주문_응답)
	s.M주문번호 = l.F2정수64_단순형(g.Ordno)
	s.M원_주문번호 = l.F2정수64_단순형(g.Orgordno)
	s.RT코드 = l.F2문자열(rt.TrCode)
	s.M종목코드 = 종목코드
	s.M수량 = l.F2정수64_단순형(g.Mdfycnfqty)
	s.M가격 = l.F2정수64_단순형(g.Mdfycnfprc)
	s.M잔량 = l.F2정수64_단순형(g.Unercqty)
	s.M시각 = 시각

	return s, nil
}

func New현물주문_취소(rt *REALTIME_DATA) (값 *xt.S주문_응답, 에러 error) {
	g := (*SC3_OutBlock)(unsafe.Pointer(rt.Data))

	시각_문자열 := l.F2문자열(g.Exectime)
	시각_문자열 = 시각_문자열[:6] + "." + 시각_문자열[7:]
	시각 := l.F2금일_시각_단순형("150405.999", 시각_문자열)

	종목코드 := l.F2문자열_공백제거(g.ShtnIsuno)
	종목코드 = 종목코드[1:] // 맨 앞의 'A' 제거

	s := new(xt.S주문_응답)
	s.M주문번호 = l.F2정수64_단순형(g.Ordno)
	s.M원_주문번호 = l.F2정수64_단순형(g.Orgordno)
	s.RT코드 = l.F2문자열(rt.TrCode)
	s.M종목코드 = 종목코드
	s.M수량 = l.F2정수64_단순형(g.Canccnfqty)
	s.M잔량 = l.F2정수64_단순형(g.Orgordunercqty)
	s.M시각 = 시각

	return s, nil
}

func New현물주문_거부(rt *REALTIME_DATA) (값 *xt.S주문_응답, 에러 error) {
	g := (*SC4_OutBlock)(unsafe.Pointer(rt.Data))

	시각_문자열 := l.F2문자열(g.Exectime)
	시각_문자열 = 시각_문자열[:6] + "." + 시각_문자열[7:]
	시각 := l.F2금일_시각_단순형("150405.999", 시각_문자열)

	종목코드 := l.F2문자열_공백제거(g.ShtnIsuno)
	종목코드 = 종목코드[1:] // 맨 앞의 'A' 제거

	s := new(xt.S주문_응답)
	s.M주문번호 = l.F2정수64_단순형(g.Ordno)
	s.M원_주문번호 = l.F2정수64_단순형(g.Orgordno)
	s.RT코드 = l.F2문자열(rt.TrCode)
	s.M종목코드 = 종목코드
	s.M수량 = l.F2정수64_단순형(g.Rjtqty)
	s.M잔량 = l.F2정수64_단순형(g.Unercqty)
	l.F문자열_출력("%v", l.F2문자열(g.Exectime))
	s.M시각 = 시각

	return s, nil
}

func New코스피_호가_잔량(rt *REALTIME_DATA) (값 *xt.S코스피_호가_잔량_실시간_정보, 에러 error) {
	g := (*H1_OutBlock)(unsafe.Pointer(rt.Data))

	s := new(xt.S코스피_호가_잔량_실시간_정보)
	s.M종목코드 = l.F2문자열(g.Shcode)
	s.M시각 = l.F2금일_시각_단순형("150405", g.Hotime)
	s.M동시호가_구분 = xt.T동시호가_구분(l.F2정수64_단순형(g.Donsigubun))
	s.M배분적용_구분 = l.F2참거짓(g.Gubun, " ", false)

	매도호가_모음 := []int64{
		l.F2정수64_단순형(g.Offerho1), l.F2정수64_단순형(g.Offerho2), l.F2정수64_단순형(g.Offerho3),
		l.F2정수64_단순형(g.Offerho4), l.F2정수64_단순형(g.Offerho5), l.F2정수64_단순형(g.Offerho6),
		l.F2정수64_단순형(g.Offerho7), l.F2정수64_단순형(g.Offerho8), l.F2정수64_단순형(g.Offerho9),
		l.F2정수64_단순형(g.Offerho10)}

	매도잔량_모음 := []int64{
		l.F2정수64_단순형(g.Offerrem1), l.F2정수64_단순형(g.Offerrem2), l.F2정수64_단순형(g.Offerrem3),
		l.F2정수64_단순형(g.Offerrem4), l.F2정수64_단순형(g.Offerrem5), l.F2정수64_단순형(g.Offerrem6),
		l.F2정수64_단순형(g.Offerrem7), l.F2정수64_단순형(g.Offerrem8), l.F2정수64_단순형(g.Offerrem9),
		l.F2정수64_단순형(g.Offerrem10)}

	매수호가_모음 := []int64{
		l.F2정수64_단순형(g.Bidho1), l.F2정수64_단순형(g.Bidho2), l.F2정수64_단순형(g.Bidho3),
		l.F2정수64_단순형(g.Bidho4), l.F2정수64_단순형(g.Bidho5), l.F2정수64_단순형(g.Bidho6),
		l.F2정수64_단순형(g.Bidho7), l.F2정수64_단순형(g.Bidho8), l.F2정수64_단순형(g.Bidho9),
		l.F2정수64_단순형(g.Bidho10)}

	매수잔량_모음 := []int64{
		l.F2정수64_단순형(g.Bidrem1), l.F2정수64_단순형(g.Bidrem2), l.F2정수64_단순형(g.Bidrem3),
		l.F2정수64_단순형(g.Bidrem4), l.F2정수64_단순형(g.Bidrem5), l.F2정수64_단순형(g.Bidrem6),
		l.F2정수64_단순형(g.Bidrem7), l.F2정수64_단순형(g.Bidrem8), l.F2정수64_단순형(g.Bidrem9),
		l.F2정수64_단순형(g.Bidrem10)}

	if len(매도호가_모음) != len(매도잔량_모음) {
		return nil, l.New에러("매도호가, 매도잔량 수량이 서로 다름. %v %v",
			len(매도호가_모음), len(매도잔량_모음))
	}

	if len(매수호가_모음) != len(매수잔량_모음) {
		return nil, l.New에러("매수호가, 매수잔량 수량이 서로 다름. %v %v",
			len(매수호가_모음), len(매수잔량_모음))
	}

	s.M매도호가_모음 = make([]int64, 0)
	s.M매도잔량_모음 = make([]int64, 0)
	for i := 0; i < len(매도잔량_모음); i++ {
		if 매도호가_모음[i] == 0 || 매도잔량_모음[i] == 0 {
			continue
		}

		s.M매도호가_모음 = append(s.M매도호가_모음, 매도호가_모음[i])
		s.M매도잔량_모음 = append(s.M매도잔량_모음, 매도잔량_모음[i])
	}

	s.M매수호가_모음 = make([]int64, 0)
	s.M매수잔량_모음 = make([]int64, 0)
	for i := 0; i < len(매수잔량_모음); i++ {
		if 매수호가_모음[i] == 0 || 매수잔량_모음[i] == 0 {
			continue
		}

		s.M매수호가_모음 = append(s.M매수호가_모음, 매수호가_모음[i])
		s.M매수잔량_모음 = append(s.M매수잔량_모음, 매수잔량_모음[i])
	}

	s.M매도_총잔량 = l.F2정수64_단순형(g.Totofferrem)
	s.M매수_총잔량 = l.F2정수64_단순형(g.Totbidrem)

	return s, nil
}

func New코스피_시간외_호가_잔량(rt *REALTIME_DATA) (값 *xt.S코스피_시간외_호가_잔량_실시간_정보, 에러 error) {
	g := (*H2_OutBlock)(unsafe.Pointer(rt.Data))

	s := new(xt.S코스피_시간외_호가_잔량_실시간_정보)
	s.M종목코드 = l.F2문자열(g.Shcode)
	s.M시각 = l.F2금일_시각_단순형("150405", g.Hotime)
	s.M매도잔량 = l.F2정수64_단순형(g.Tmofferrem)
	s.M매수잔량 = l.F2정수64_단순형(g.Tmbidrem)
	s.M매도수량_직전대비 = l.F2정수64_단순형(g.Pretmoffercha)
	s.M매수수량_직전대비 = l.F2정수64_단순형(g.Pretmbidcha)

	return s, nil
}

func New코스피_체결(rt *REALTIME_DATA) (값 *xt.S코스피_체결, 에러 error) {
	g := (*S3_OutBlock)(unsafe.Pointer(rt.Data))

	s := new(xt.S코스피_체결)
	s.M종목코드 = l.F2문자열(g.Shcode)
	s.M시각 = l.F2금일_시각_단순형("150405", g.Chetime)
	s.M전일대비구분 = xt.T전일대비_구분(l.F2정수64_단순형(g.Sign))
	s.M전일대비등락폭 = l.F2정수64_단순형(g.Change)
	s.M전일대비등락율 = l.F2실수_단순형(g.Drate)
	s.M현재가 = l.F2정수64_단순형(g.Price)
	s.M시가시각 = l.F2금일_시각_단순형("150405", g.Opentime)
	s.M시가 = l.F2정수64_단순형(g.Open)
	s.M고가시각 = l.F2금일_시각_단순형("150405", g.Hightime)
	s.M고가 = l.F2정수64_단순형(g.High)
	s.M저가시각 = l.F2금일_시각_단순형("150405", g.Lowtime)
	s.M저가 = l.F2정수64_단순형(g.Low)

	switch l.F2문자열(g.Cgubun) {
	case "+":
		s.M체결구분 = l.P매수
	case "-":
		s.M체결구분 = l.P매도
	default:
		panic(l.F2문자열("예상하지 못한 체결구분 값 : '%v'", l.F2문자열(g.Cgubun)))
	}

	s.M체결량 = l.F2정수64_단순형(g.Cvolume)
	s.M누적거래량 = l.F2정수64_단순형(g.Volume)
	s.M누적거래대금 = l.F2정수64_단순형(g.Value)
	s.M매도누적체결량 = l.F2정수64_단순형(g.Mdvolume)
	s.M매도누적체결건수 = l.F2정수64_단순형(g.Mdchecnt)
	s.M매수누적체결량 = l.F2정수64_단순형(g.Msvolume)
	s.M매수누적체결건수 = l.F2정수64_단순형(g.Mschecnt)
	s.M체결강도 = l.F2실수_단순형(g.Cpower)
	s.M가중평균가 = l.F2정수64_단순형(g.WAvrg)
	s.M매도호가 = l.F2정수64_단순형(g.Offerho)
	s.M매수호가 = l.F2정수64_단순형(g.Bidho)

	switch l.F2문자열_공백제거(g.Status) {
	case "0", "00":
		s.M장_정보 = l.P장_중
	case "4", "04":
		s.M장_정보 = l.P장_후_시간외
	case "10":
		s.M장_정보 = l.P장_전_시간외
	default:
		panic(l.F2문자열("예상하지 못한 장 정보 값 : '%v'", l.F2문자열_공백제거(g.Status)))
	}

	s.M전일동시간대거래량 = l.F2정수64_단순형(g.Jnilvolume)

	return s, nil
}

func New코스피_예상_체결(rt *REALTIME_DATA) (값 *xt.S코스피_예상_체결, 에러 error) {
	g := (*YS3OutBlock)(unsafe.Pointer(rt.Data))

	s := new(xt.S코스피_예상_체결)
	s.M종목코드 = l.F2문자열(g.Shcode)
	s.M시각 = l.F2금일_시각_단순형("150405", g.Hotime)
	s.M예상체결가격 = l.F2정수64_단순형(g.Yeprice)
	s.M예상체결수량 = l.F2정수64_단순형(g.Yevolume)
	s.M예상체결가전일종가대비구분 = xt.T전일대비_구분(l.F2정수64_단순형(g.Jnilysign))
	s.M예상체결가전일종가대비등락폭 = l.F2정수64_단순형(g.Preychange)
	s.M예상체결가전일종가대비등락율 = l.F2실수_단순형(g.Jnilydrate)
	s.M예상매도호가 = l.F2정수64_단순형(g.Yofferho0)
	s.M예상매수호가 = l.F2정수64_단순형(g.Ybidho0)
	s.M예상매도호가수량 = l.F2정수64_단순형(g.Yofferrem0)
	s.M예상매수호가수량 = l.F2정수64_단순형(g.Ybidrem0)

	return s, nil
}

func New코스피_ETF_NAV(rt *REALTIME_DATA) (값 *xt.S코스피_ETF_NAV, 에러 error) {
	g := (*I5_OutBlock)(unsafe.Pointer(rt.Data))

	s := new(xt.S코스피_ETF_NAV)
	s.M종목코드 = l.F2문자열(g.Shcode)
	s.M시각 = l.F2금일_시각_단순형("15:04:05", g.Time)
	s.M현재가 = l.F2정수64_단순형(g.Price)
	s.M전일대비구분 = xt.T전일대비_구분(l.F2정수64_단순형(g.Sign))
	s.M전일대비등락폭 = l.F2정수64_단순형(g.Change)
	s.M누적거래량 = l.F2실수_단순형(g.Volume)
	s.M현재가NAV차이 = l.F2실수_단순형(g.Navdiff)
	s.NAV = l.F2실수_단순형(g.Nav)
	s.NAV전일대비 = l.F2실수_단순형(g.Navdiff)
	s.M추적오차 = l.F2실수_단순형_공백은_0(g.Crate)
	s.M괴리 = l.F2실수_단순형_공백은_0(g.Grate)
	s.M지수 = l.F2실수_단순형_공백은_0(g.Jisu)
	s.M지수전일대비등락폭 = l.F2실수_단순형_공백은_0(g.Jichange)
	s.M지수전일대비등락율 = l.F2실수_단순형_공백은_0(g.Jirate)

	return s, nil
}

func New주식_VI발동해제(rt *REALTIME_DATA) (값 *xt.S주식_VI발동해제, 에러 error) {
	g := (*VI_OutBlock)(unsafe.Pointer(rt.Data))

	s := new(xt.S주식_VI발동해제)
	s.M종목코드 = l.F2문자열(g.Shcode)
	s.M참조코드 = l.F2문자열(g.Ref_shcode)
	s.M시각 = l.F2금일_시각_단순형("150405", g.Time)
	s.M구분 = xt.VI발동해제(l.F2정수64_단순형(g.Vi_gubun))
	s.M정적VI발동_기준가격 = l.F2정수64_단순형(g.Svi_recprice)
	s.M동적VI발동_기준가격 = l.F2정수64_단순형(g.Dvi_recprice)
	s.VI발동가격 = l.F2정수64_단순형(g.Vi_trgprice)

	return s, nil
}

func New시간외_단일가VI발동해제(rt *REALTIME_DATA) (값 *xt.S시간외_단일가VI발동해제, 에러 error) {
	g := (*DVIOutBlock)(unsafe.Pointer(rt.Data))

	s := new(xt.S시간외_단일가VI발동해제)
	s.M종목코드 = l.F2문자열(g.Shcode)
	s.M참조코드 = l.F2문자열(g.Ref_shcode)
	s.M시각 = l.F2금일_시각_단순형("150405", g.Time)
	s.M구분 = xt.VI발동해제(l.F2정수64_단순형(g.Vi_gubun))
	s.M정적VI발동_기준가격 = l.F2정수64_단순형(g.Svi_recprice)
	s.M동적VI발동_기준가격 = l.F2정수64_단순형(g.Dvi_recprice)
	s.VI발동가격 = l.F2정수64_단순형(g.Vi_trgprice)

	return s, nil
}

func New장_운영정보(rt *REALTIME_DATA) (값 *xt.S장_운영정보, 에러 error) {
	g := (*JIFOutBlock)(unsafe.Pointer(rt.Data))

	s := new(xt.S장_운영정보)
	s.M장_구분 = xt.T시장구분(l.F2문자열(g.Jangubun))
	s.M장_상태 = xt.T시장상태(l.F2정수_단순형(g.Jstatus))

	return s, nil
}
