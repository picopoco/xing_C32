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
	"sync"
	"time"
	"unsafe"
)

// _UNPACKED를 자동 변환하면 자꾸 문제가 생김.
// 원인을 알 수 없어서 자동변환되지 않도록 따로 선언함.

// C.TR_DATA_UNPACKED
type TR_DATA struct {
	RequestID           int32
	DataLength          int32
	TotalDataBufferSize int32
	ElapsedTime         int32
	DataMode            int32
	TrCode              [10]byte
	X_TrCode            [1]byte
	Cont                [1]byte
	ContKey             [18]byte
	X_ContKey           [1]byte
	None                [31]byte
	BlockName           [16]byte
	X_BlockName         [1]byte
	Pad_cgo_0           [1]byte
	Data                *byte
}

type TR_DATA_PACKED struct {
	RequestID           int32
	DataLength          int32
	TotalDataBufferSize int32
	ElapsedTime         int32
	DataMode            int32
	TrCode              [10]byte
	X_TrCode            [1]byte
	Cont                [1]byte
	ContKey             [18]byte
	X_ContKey           [1]byte
	None                [31]byte
	BlockName           [16]byte
	X_BlockName         [1]byte
	Pad_cgo_0           [4]byte
}

// C.REALTIME_DATA_UNPACKED
type REALTIME_DATA struct {
	TrCode     [3]byte
	X_TrCode   [1]byte
	KeyLength  int32
	KeyData    [32]byte
	X_KeyData  [1]byte
	RegKey     [32]byte
	X_RegKey   [1]byte
	Pad_cgo_0  [2]byte
	DataLength int32
	Data       *byte
}

type REALTIME_DATA_PACKED struct {
	TrCode    [3]byte
	X_TrCode  [1]byte
	KeyLength int32
	KeyData   [32]byte
	X_KeyData [1]byte
	RegKey    [32]byte
	X_RegKey  [1]byte
	Pad_cgo_0 [4]byte
	Pad_cgo_1 [4]byte
}

// C.MSG_DATA_UNPACKED
type MSG_DATA struct {
	RequestID   int32
	SystemError int32
	MsgCode     [5]byte
	X_MsgCode   [1]byte
	Pad_cgo_0   [2]byte
	MsgLength   int32
	MsgData     *byte
}

type MSG_DATA_PACKED struct {
	RequestID   int32
	SystemError int32
	MsgCode     [5]byte
	X_MsgCode   [1]byte
	Pad_cgo_0   [4]byte
	Pad_cgo_1   [4]byte
}

type S메시지_저장소 struct {
	sync.Mutex
	저장소 map[int][]unsafe.Pointer
}

func New메시지_저장소() *S메시지_저장소 {
	s := new(S메시지_저장소)
	s.저장소 = make(map[int][]unsafe.Pointer)

	return s
}

func (s *S메시지_저장소) G값(식별번호 int) []unsafe.Pointer {
	s.Lock()
	defer s.Unlock()

	기존_내용 := s.저장소[식별번호]

	return 기존_내용
}

func (s *S메시지_저장소) S추가(식별번호 int, 메시지 unsafe.Pointer) {
	s.Lock()
	defer s.Unlock()

	기존_내용, ok := s.저장소[식별번호]

	if !ok {
		s.저장소[식별번호] = []unsafe.Pointer{메시지}
	} else {
		s.저장소[식별번호] = append(기존_내용, 메시지)
	}
}

func (s *S메시지_저장소) S삭제(식별번호 int) {
	s.Lock()
	defer s.Unlock()

	delete(s.저장소, 식별번호)
}

func New콜백_대기_항목(식별번호 int, TR코드 string, 값 interface{}) *S콜백_대기_항목 {
	s := new(S콜백_대기_항목)
	s.M식별번호 = 식별번호
	s.M생성_시각 = time.Now()
	s.M값 = 값

	return s
}

type S콜백_대기_항목 struct {
	M식별번호  int
	M생성_시각 time.Time
	TR코드   string
	M값     interface{}
}

type S콜백_대기_저장소 struct {
	sync.Mutex
	저장소 map[int]*S콜백_대기_항목
}

func (s *S콜백_대기_저장소) G대기_항목(식별번호 int) *S콜백_대기_항목 {
	s.Lock()
	defer s.Unlock()

	return s.저장소[식별번호]
}

func (s *S콜백_대기_저장소) S추가(식별번호 int, 대기_항목 *S콜백_대기_항목) {
	s.Lock()
	defer s.Unlock()

	s.저장소[식별번호] = 대기_항목
}

func (s *S콜백_대기_저장소) S삭제(식별번호 int) {
	s.Lock()
	defer s.Unlock()

	delete(s.저장소, 식별번호)
}
