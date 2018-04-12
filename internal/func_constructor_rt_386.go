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
	"github.com/ghts/api_bridge_xing_C"
	"github.com/ghts/lib"
	"github.com/ghts/types_xing"

	"sync"
	"unsafe"
)

var once_H1_OutBlock sync.Once

func New코스피_호가_잔량_실시간정보(데이터 unsafe.Pointer, 길이 int) (값 *xing.S코스피_호가_잔량_실시간_정보, 에러 error) {
	defer lib.F에러패닉_처리(lib.S에러패닉_처리{
		M에러: &에러,
		M함수: func() { 값 = nil }})

	once_H1_OutBlock.Do(func() {
		lib.F조건부_패닉(unsafe.Sizeof(c.H1_OutBlock{}) != uintptr(길이),
			"H1_OutBlock 길이 불일치. '%v' '%v'", unsafe.Sizeof(c.H1_OutBlock{}), 길이)
	})

	g := (*c.H1_OutBlock)(데이터)
	s := new(xing.S코스피_호가_잔량_실시간_정보)
	s.M종목코드 = lib.F2문자열(g.Shcode)
	s.M시각 = lib.F2금일_시각_단순형("150405", g.Hotime)
	s.M동시호가_구분 = xing.T동시호가_구분(lib.F2정수64_단순형(g.Donsigubun))
	s.M배분적용_구분 = lib.F2참거짓(g.Gubun, " ", false)

	매도호가_모음 := []int64{
		lib.F2정수64_단순형(g.Offerho1), lib.F2정수64_단순형(g.Offerho2), lib.F2정수64_단순형(g.Offerho3),
		lib.F2정수64_단순형(g.Offerho4), lib.F2정수64_단순형(g.Offerho5), lib.F2정수64_단순형(g.Offerho6),
		lib.F2정수64_단순형(g.Offerho7), lib.F2정수64_단순형(g.Offerho8), lib.F2정수64_단순형(g.Offerho9),
		lib.F2정수64_단순형(g.Offerho10)}

	매도잔량_모음 := []int64{
		lib.F2정수64_단순형(g.Offerrem1), lib.F2정수64_단순형(g.Offerrem2), lib.F2정수64_단순형(g.Offerrem3),
		lib.F2정수64_단순형(g.Offerrem4), lib.F2정수64_단순형(g.Offerrem5), lib.F2정수64_단순형(g.Offerrem6),
		lib.F2정수64_단순형(g.Offerrem7), lib.F2정수64_단순형(g.Offerrem8), lib.F2정수64_단순형(g.Offerrem9),
		lib.F2정수64_단순형(g.Offerrem10)}

	매수호가_모음 := []int64{
		lib.F2정수64_단순형(g.Bidho1), lib.F2정수64_단순형(g.Bidho2), lib.F2정수64_단순형(g.Bidho3),
		lib.F2정수64_단순형(g.Bidho4), lib.F2정수64_단순형(g.Bidho5), lib.F2정수64_단순형(g.Bidho6),
		lib.F2정수64_단순형(g.Bidho7), lib.F2정수64_단순형(g.Bidho8), lib.F2정수64_단순형(g.Bidho9),
		lib.F2정수64_단순형(g.Bidho10)}

	매수잔량_모음 := []int64{
		lib.F2정수64_단순형(g.Bidrem1), lib.F2정수64_단순형(g.Bidrem2), lib.F2정수64_단순형(g.Bidrem3),
		lib.F2정수64_단순형(g.Bidrem4), lib.F2정수64_단순형(g.Bidrem5), lib.F2정수64_단순형(g.Bidrem6),
		lib.F2정수64_단순형(g.Bidrem7), lib.F2정수64_단순형(g.Bidrem8), lib.F2정수64_단순형(g.Bidrem9),
		lib.F2정수64_단순형(g.Bidrem10)}

	if len(매도호가_모음) != len(매도잔량_모음) {
		return nil, lib.New에러("매도호가, 매도잔량 수량이 서로 다름. %v %v",
			len(매도호가_모음), len(매도잔량_모음))
	}

	if len(매수호가_모음) != len(매수잔량_모음) {
		return nil, lib.New에러("매수호가, 매수잔량 수량이 서로 다름. %v %v",
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

	s.M매도_총잔량 = lib.F2정수64_단순형(g.Totofferrem)
	s.M매수_총잔량 = lib.F2정수64_단순형(g.Totbidrem)

	return s, nil
}

var once_H2_OutBlock sync.Once

func New코스피_시간외_호가_잔량_실시간_정보(데이터 unsafe.Pointer, 길이 int) (값 *xing.S코스피_시간외_호가_잔량_실시간_정보, 에러 error) {
	defer lib.F에러패닉_처리(lib.S에러패닉_처리{
		M에러: &에러,
		M함수: func() { 값 = nil }})

	once_H2_OutBlock.Do(func() {
		lib.F조건부_패닉(unsafe.Sizeof(c.H2_OutBlock{}) != uintptr(길이),
			"H2_OutBlock 길이 불일치. '%v' '%v'", unsafe.Sizeof(c.H2_OutBlock{}), 길이)
	})

	g := (*c.H2_OutBlock)(데이터)
	s := new(xing.S코스피_시간외_호가_잔량_실시간_정보)
	s.M종목코드 = lib.F2문자열(g.Shcode)
	s.M시각 = lib.F2금일_시각_단순형("150405", g.Hotime)
	s.M매도잔량 = lib.F2정수64_단순형(g.Tmofferrem)
	s.M매수잔량 = lib.F2정수64_단순형(g.Tmbidrem)
	s.M매도수량_직전대비 = lib.F2정수64_단순형(g.Pretmoffercha)
	s.M매수수량_직전대비 = lib.F2정수64_단순형(g.Pretmbidcha)

	return s, nil
}

func New코스피_체결(데이터 unsafe.Pointer, 길이 int) (값 *xing.S코스피_체결, 에러 error) {
	defer lib.F에러패닉_처리(lib.S에러패닉_처리{
		M에러: &에러,
		M함수: func() { 값 = nil }})

	g := (*c.S3_OutBlock)(데이터)
	s := new(xing.S코스피_체결)
	s.M종목코드 = lib.F2문자열(g.Shcode)
	s.M시각 = lib.F2금일_시각_단순형("150405", g.Chetime)
	s.M전일대비구분 = xing.T전일대비_구분(lib.F2정수64_단순형(g.Sign))
	s.M전일대비등락폭 = lib.F2정수64_단순형(g.Change)
	s.M전일대비등락율 = lib.F2실수_단순형(g.Drate)
	s.M현재가 = lib.F2정수64_단순형(g.Price)
	s.M시가시각 = lib.F2금일_시각_단순형("150405", g.Opentime)
	s.M시가 = lib.F2정수64_단순형(g.Open)
	s.M고가시각 = lib.F2금일_시각_단순형("150405", g.Hightime)
	s.M고가 = lib.F2정수64_단순형(g.High)
	s.M저가시각 = lib.F2금일_시각_단순형("150405", g.Lowtime)
	s.M저가 = lib.F2정수64_단순형(g.Low)

	switch lib.F2문자열(g.Cgubun) {
	case "+":
		s.M체결구분 = lib.P매수
	case "-":
		s.M체결구분 = lib.P매도
	default:
		lib.F패닉("예상하지 못한 체결구분 값 : '%v'", lib.F2문자열(g.Cgubun))
	}

	s.M체결량 = lib.F2정수64_단순형(g.Cvolume)
	s.M누적거래량 = lib.F2정수64_단순형(g.Volume)
	s.M누적거래대금 = lib.F2정수64_단순형(g.Value)
	s.M매도누적체결량 = lib.F2정수64_단순형(g.Mdvolume)
	s.M매도누적체결건수 = lib.F2정수64_단순형(g.Mdchecnt)
	s.M매수누적체결량 = lib.F2정수64_단순형(g.Msvolume)
	s.M매수누적체결건수 = lib.F2정수64_단순형(g.Mschecnt)
	s.M체결강도 = lib.F2실수_단순형(g.Cpower)
	s.M가중평균가 = lib.F2정수64_단순형(g.WAvrg)
	s.M매도호가 = lib.F2정수64_단순형(g.Offerho)
	s.M매수호가 = lib.F2정수64_단순형(g.Bidho)

	switch lib.F2문자열_공백제거(g.Status) {
	case "0", "00":
		s.M장_정보 = lib.P장_중
	case "4", "04":
		s.M장_정보 = lib.P장_후_시간외
	case "10":
		s.M장_정보 = lib.P장_전_시간외
	default:
		lib.F패닉("예상하지 못한 장 정보 값 : '%v'", lib.F2문자열_공백제거(g.Status))
	}

	s.M전일동시간대거래량 = lib.F2정수64_단순형(g.Jnilvolume)

	return s, nil
}

func New코스피_예상_체결(데이터 unsafe.Pointer, 길이 int) (값 *xing.S코스피_예상_체결, 에러 error) {
	defer lib.F에러패닉_처리(lib.S에러패닉_처리{
		M에러: &에러,
		M함수: func() { 값 = nil }})

	g := (*c.YS3OutBlock)(데이터)
	s := new(xing.S코스피_예상_체결)
	s.M종목코드 = lib.F2문자열(g.Shcode)
	s.M시각 = lib.F2금일_시각_단순형("150405", g.Hotime)
	s.M예상체결가격 = lib.F2정수64_단순형(g.Yeprice)
	s.M예상체결수량 = lib.F2정수64_단순형(g.Yevolume)
	s.M예상체결가전일종가대비구분 = xing.T전일대비_구분(lib.F2정수64_단순형(g.Jnilysign))
	s.M예상체결가전일종가대비등락폭 = lib.F2정수64_단순형(g.Preychange)
	s.M예상체결가전일종가대비등락율 = lib.F2실수_단순형(g.Jnilydrate)
	s.M예상매도호가 = lib.F2정수64_단순형(g.Yofferho0)
	s.M예상매수호가 = lib.F2정수64_단순형(g.Ybidho0)
	s.M예상매도호가수량 = lib.F2정수64_단순형(g.Yofferrem0)
	s.M예상매수호가수량 = lib.F2정수64_단순형(g.Ybidrem0)

	return s, nil
}

func New코스피_ETF_NAV(데이터 unsafe.Pointer, 길이 int) (값 *xing.S코스피_ETF_NAV, 에러 error) {
	defer lib.F에러패닉_처리(lib.S에러패닉_처리{
		M에러: &에러,
		M함수: func() { 값 = nil }})

	g := (*c.I5_OutBlock)(데이터)
	s := new(xing.S코스피_ETF_NAV)
	s.M종목코드 = lib.F2문자열(g.Shcode)
	s.M시각 = lib.F2금일_시각_단순형("15:04:05", g.Time)
	s.M현재가 = lib.F2정수64_단순형(g.Price)
	s.M전일대비구분 = xing.T전일대비_구분(lib.F2정수64_단순형(g.Sign))
	s.M전일대비등락폭 = lib.F2정수64_단순형(g.Change)
	s.M누적거래량 = lib.F2실수_단순형(g.Volume)
	s.M현재가NAV차이 = lib.F2실수_단순형(g.Navdiff)
	s.NAV = lib.F2실수_단순형(g.Nav)
	s.NAV전일대비 = lib.F2실수_단순형(g.Navdiff)
	s.M추적오차 = lib.F2실수_단순형_공백은_0(g.Crate)
	s.M괴리 = lib.F2실수_단순형_공백은_0(g.Grate)
	s.M지수 = lib.F2실수_단순형_공백은_0(g.Jisu)
	s.M지수전일대비등락폭 = lib.F2실수_단순형_공백은_0(g.Jichange)
	s.M지수전일대비등락율 = lib.F2실수_단순형_공백은_0(g.Jirate)

	return s, nil
}

func New주식_VI발동해제(데이터 unsafe.Pointer, 길이 int) (값 *xing.S주식_VI발동해제, 에러 error) {
	defer lib.F에러패닉_처리(lib.S에러패닉_처리{
		M에러: &에러,
		M함수: func() {
			값 = nil
		}})

	g := (*c.VI_OutBlock)(데이터)
	s := new(xing.S주식_VI발동해제)
	s.M종목코드 = lib.F2문자열(g.Shcode)
	s.M참조코드 = lib.F2문자열(g.Ref_shcode)
	s.M시각 = lib.F2금일_시각_단순형("150405", g.Time)
	s.M구분 = xing.VI발동해제(lib.F2정수64_단순형(g.Vi_gubun))
	s.M정적VI발동_기준가격 = lib.F2정수64_단순형(g.Svi_recprice)
	s.M동적VI발동_기준가격 = lib.F2정수64_단순형(g.Dvi_recprice)
	s.VI발동가격 = lib.F2정수64_단순형(g.Vi_trgprice)

	return s, nil
}

func New시간외_단일가VI발동해제(데이터 unsafe.Pointer, 길이 int) (값 *xing.S시간외_단일가VI발동해제, 에러 error) {
	defer lib.F에러패닉_처리(lib.S에러패닉_처리{
		M에러: &에러,
		M함수: func() {
			값 = nil
		}})

	g := (*c.DVIOutBlock)(데이터)
	s := new(xing.S시간외_단일가VI발동해제)
	s.M종목코드 = lib.F2문자열(g.Shcode)
	s.M참조코드 = lib.F2문자열(g.Ref_shcode)
	s.M시각 = lib.F2금일_시각_단순형("150405", g.Time)
	s.M구분 = xing.VI발동해제(lib.F2정수64_단순형(g.Vi_gubun))
	s.M정적VI발동_기준가격 = lib.F2정수64_단순형(g.Svi_recprice)
	s.M동적VI발동_기준가격 = lib.F2정수64_단순형(g.Dvi_recprice)
	s.VI발동가격 = lib.F2정수64_단순형(g.Vi_trgprice)

	return s, nil
}

func New장_운영정보(데이터 unsafe.Pointer, 길이 int) (값 *xing.S장_운영정보, 에러 error) {
	defer lib.F에러패닉_처리(lib.S에러패닉_처리{
		M에러: &에러,
		M함수: func() {
			값 = nil
		}})

	g := (*c.JIFOutBlock)(데이터)
	s := new(xing.S장_운영정보)
	s.M장_구분 = xing.T시장구분(lib.F2문자열(g.Jangubun))
	s.M장_상태 = xing.T시장상태(lib.F2정수_단순형(g.Jstatus))

	return s, nil
}
