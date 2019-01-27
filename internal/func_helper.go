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
// #include <./types_c.h>
import "C"

import (
	"github.com/ghts/lib"

	"os"
	"unsafe"
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
	case 전송_권한.G코드() != TR코드:
		panic("예상하지 못한 경우.")
	}

	전송_권한.G전송_권한_획득()
}

func TR전송_코드별_초당_제한_확인(TR코드 string) {
	전송_권한, 존재함 := tr전송_코드별_초당_제한[TR코드]

	switch {
	case !존재함:
		panic(lib.New에러("전송제한을 찾을 수 없음 : '%v'", TR코드))
	case 전송_권한.G코드() != TR코드:
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

func f자료형_크기_비교_확인() (에러 error) {
	lib.S예외처리{M에러: &에러}.S실행()

	lib.F조건부_패닉(unsafe.Sizeof(TR_DATA{}) != unsafe.Sizeof(C.TR_DATA_UNPACKED{}), "TR_DATA_UNPACKED 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(REALTIME_DATA{}) != unsafe.Sizeof(C.REALTIME_DATA_UNPACKED{}), "REALTIME_DATA_UNPACKED 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(MSG_DATA{}) != unsafe.Sizeof(C.MSG_DATA_UNPACKED{}), "MSG_DATA_UNPACKED 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(TR_DATA_PACKED{}) != unsafe.Sizeof(C.TR_DATA{}), "TR_DATA 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(REALTIME_DATA_PACKED{}) != unsafe.Sizeof(C.REALTIME_DATA{}), "REALTIME_DATA 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(MSG_DATA_PACKED{}) != unsafe.Sizeof(C.MSG_DATA{}), "MSG_DATA 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(LINK_DATA{}) != unsafe.Sizeof(C.LINK_DATA{}), "LINK_DATA 크기 불일치")

	lib.F조건부_패닉(unsafe.Sizeof(CSPAT00600InBlock1{}) != unsafe.Sizeof(C.CSPAT00600InBlock1{}), "CSPAT00600InBlock1 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(CSPAT00600OutBlock1{}) != unsafe.Sizeof(C.CSPAT00600OutBlock1{}), "CSPAT00600OutBlock1 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(CSPAT00600OutBlock2{}) != unsafe.Sizeof(C.CSPAT00600OutBlock2{}), "CSPAT00600OutBlock2 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(CSPAT00600OutBlockAll{}) != unsafe.Sizeof(C.CSPAT00600OutBlockAll{}), "CSPAT00600OutBlockAll 크기 불일치")

	lib.F조건부_패닉(unsafe.Sizeof(CSPAT00700InBlock1{}) != unsafe.Sizeof(C.CSPAT00700InBlock1{}), "CSPAT00700InBlock1 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(CSPAT00700OutBlock1{}) != unsafe.Sizeof(C.CSPAT00700OutBlock1{}), "CSPAT00700OutBlock1 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(CSPAT00700OutBlock2{}) != unsafe.Sizeof(C.CSPAT00700OutBlock2{}), "CSPAT00700OutBlock2 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(CSPAT00700OutBlockAll{}) != unsafe.Sizeof(C.CSPAT00700OutBlockAll{}), "CSPAT00700OutBlockAll 크기 불일치")

	lib.F조건부_패닉(unsafe.Sizeof(CSPAT00800InBlock1{}) != unsafe.Sizeof(C.CSPAT00800InBlock1{}), "CSPAT00800InBlock1 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(CSPAT00800OutBlock1{}) != unsafe.Sizeof(C.CSPAT00800OutBlock1{}), "CSPAT00800OutBlock1 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(CSPAT00800OutBlock2{}) != unsafe.Sizeof(C.CSPAT00800OutBlock2{}), "CSPAT00800OutBlock2 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(CSPAT00800OutBlockAll{}) != unsafe.Sizeof(C.CSPAT00800OutBlockAll{}), "CSPAT00800OutBlockAll 크기 불일치")

	lib.F조건부_패닉(unsafe.Sizeof(SC0_OutBlock{}) != unsafe.Sizeof(C.SC0_OutBlock{}), "SC0_OutBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(SC1_OutBlock{}) != unsafe.Sizeof(C.SC1_OutBlock{}), "SC1_OutBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(SC2_OutBlock{}) != unsafe.Sizeof(C.SC2_OutBlock{}), "SC2_OutBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(SC3_OutBlock{}) != unsafe.Sizeof(C.SC3_OutBlock{}), "SC3_OutBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(SC4_OutBlock{}) != unsafe.Sizeof(C.SC4_OutBlock{}), "SC4_OutBlock 크기 불일치")

	lib.F조건부_패닉(unsafe.Sizeof(T0167OutBlock{}) != unsafe.Sizeof(C.T0167OutBlock{}), "T0167OutBlock 크기 불일치")

	lib.F조건부_패닉(unsafe.Sizeof(T1101InBlock{}) != unsafe.Sizeof(C.T1101InBlock{}), "T1101InBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(T1101OutBlock{}) != unsafe.Sizeof(C.T1101OutBlock{}), "T1101OutBlock 크기 불일치")

	lib.F조건부_패닉(unsafe.Sizeof(T1102InBlock{}) != unsafe.Sizeof(C.T1102InBlock{}), "T1102InBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(T1102OutBlock{}) != unsafe.Sizeof(C.T1102OutBlock{}), "T1102OutBlock 크기 불일치")

	lib.F조건부_패닉(unsafe.Sizeof(T1305InBlock{}) != unsafe.Sizeof(C.T1305InBlock{}), "T1305InBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(T1305OutBlock{}) != unsafe.Sizeof(C.T1305OutBlock{}), "T1305OutBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(T1305OutBlock1{}) != unsafe.Sizeof(C.T1305OutBlock1{}), "T1305OutBlock1 크기 불일치")

	lib.F조건부_패닉(unsafe.Sizeof(T1310InBlock{}) != unsafe.Sizeof(C.T1310InBlock{}), "T1310InBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(T1310OutBlock{}) != unsafe.Sizeof(C.T1310OutBlock{}), "T1310OutBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(T1310OutBlock1{}) != unsafe.Sizeof(C.T1310OutBlock1{}), "T1310OutBlock1 크기 불일치")

	lib.F조건부_패닉(unsafe.Sizeof(T1901InBlock{}) != unsafe.Sizeof(C.T1901InBlock{}), "T1901InBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(T1901OutBlock{}) != unsafe.Sizeof(C.T1901OutBlock{}), "T1901OutBlock 크기 불일치")

	lib.F조건부_패닉(unsafe.Sizeof(T1902InBlock{}) != unsafe.Sizeof(C.T1902InBlock{}), "T1902InBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(T1902OutBlock{}) != unsafe.Sizeof(C.T1902OutBlock{}), "T1902OutBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(T1902OutBlock1{}) != unsafe.Sizeof(C.T1902OutBlock1{}), "T1902OutBlock1 크기 불일치")

	lib.F조건부_패닉(unsafe.Sizeof(T8411InBlock{}) != unsafe.Sizeof(C.T8411InBlock{}), "T8411InBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(T8411OutBlock{}) != unsafe.Sizeof(C.T8411OutBlock{}), "T8411OutBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(T8411OutBlock1{}) != unsafe.Sizeof(C.T8411OutBlock1{}), "T8411OutBlock1 크기 불일치")

	lib.F조건부_패닉(unsafe.Sizeof(T8412InBlock{}) != unsafe.Sizeof(C.T8412InBlock{}), "T8412InBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(T8412OutBlock{}) != unsafe.Sizeof(C.T8412OutBlock{}), "T8412OutBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(T8412OutBlock1{}) != unsafe.Sizeof(C.T8412OutBlock1{}), "T8412OutBlock1 크기 불일치")

	lib.F조건부_패닉(unsafe.Sizeof(T8413InBlock{}) != unsafe.Sizeof(C.T8413InBlock{}), "T8413InBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(T8413OutBlock{}) != unsafe.Sizeof(C.T8413OutBlock{}), "T8413OutBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(T8413OutBlock1{}) != unsafe.Sizeof(C.T8413OutBlock1{}), "T8413OutBlock1 크기 불일치")

	lib.F조건부_패닉(unsafe.Sizeof(T8428InBlock{}) != unsafe.Sizeof(C.T8428InBlock{}), "T8428InBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(T8428OutBlock{}) != unsafe.Sizeof(C.T8428OutBlock{}), "T8428OutBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(T8428OutBlock1{}) != unsafe.Sizeof(C.T8428OutBlock1{}), "T8428OutBlock1 크기 불일치")

	lib.F조건부_패닉(unsafe.Sizeof(T8436InBlock{}) != unsafe.Sizeof(C.T8436InBlock{}), "T8436InBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(T8436OutBlock{}) != unsafe.Sizeof(C.T8436OutBlock{}), "T8436OutBlock 크기 불일치")

	lib.F조건부_패닉(unsafe.Sizeof(H1_InBlock{}) != unsafe.Sizeof(C.H1_InBlock{}), "H1_InBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(H1_OutBlock{}) != unsafe.Sizeof(C.H1_OutBlock{}), "H1_OutBlock 크기 불일치")

	lib.F조건부_패닉(unsafe.Sizeof(H2_InBlock{}) != unsafe.Sizeof(C.H2_InBlock{}), "H2_InBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(H2_OutBlock{}) != unsafe.Sizeof(C.H2_OutBlock{}), "H2_OutBlock 크기 불일치")

	lib.F조건부_패닉(unsafe.Sizeof(S3_InBlock{}) != unsafe.Sizeof(C.S3_InBlock{}), "S3_InBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(S3_OutBlock{}) != unsafe.Sizeof(C.S3_OutBlock{}), "S3_OutBlock 크기 불일치")

	lib.F조건부_패닉(unsafe.Sizeof(YS3InBlock{}) != unsafe.Sizeof(C.YS3InBlock{}), "YS3InBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(YS3OutBlock{}) != unsafe.Sizeof(C.YS3OutBlock{}), "YS3OutBlock 크기 불일치")

	lib.F조건부_패닉(unsafe.Sizeof(I5_InBlock{}) != unsafe.Sizeof(C.I5_InBlock{}), "I5_InBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(I5_OutBlock{}) != unsafe.Sizeof(C.I5_OutBlock{}), "I5_OutBlock 크기 불일치")

	lib.F조건부_패닉(unsafe.Sizeof(VI_InBlock{}) != unsafe.Sizeof(C.VI_InBlock{}), "VI_InBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(VI_OutBlock{}) != unsafe.Sizeof(C.VI_OutBlock{}), "VI_OutBlock 크기 불일치")

	lib.F조건부_패닉(unsafe.Sizeof(DVIInBlock{}) != unsafe.Sizeof(C.DVIInBlock{}), "DVIInBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(DVIOutBlock{}) != unsafe.Sizeof(C.DVIOutBlock{}), "DVIOutBlock 크기 불일치")

	lib.F조건부_패닉(unsafe.Sizeof(JIFInBlock{}) != unsafe.Sizeof(C.JIFInBlock{}), "JIFInBlock 크기 불일치")
	lib.F조건부_패닉(unsafe.Sizeof(JIFOutBlock{}) != unsafe.Sizeof(C.JIFOutBlock{}), "JIFOutBlock 크기 불일치")

	return nil
}
