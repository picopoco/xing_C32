/* 증권사 제공 DLL함수 레퍼런스와 예제소스 코드의 'IXingAPI.h'파일,
* DevCenter의 C++ 헤더를 참조해서 약간 수정.
* COPIED FROM API provider's reference and sample source code.
* MODIFIED by GHTS Authors.
* LICENSING TERM follows that of original code.
*
* 저작권 관련 규정은 레퍼런스 및 헤더 파일, 샘플 소스코드의 원래 저작권 규정을 따름.
*
* 변수명명규칙
* Go언어와 데이터를 주고 받는 구조체의 멤버 변수는 Go언어와의 호환성을 위해서,
* Go언어에서 public형은 대문자로 시작해야 함.
* C 헤더 파일은 'go tool cgo -godefs'로 바로 Go자료형으로 변환해서 사용하므로,
* 구조체 멤버 필드의 경우 Go언어의 변수명명 규칙을 여기에서도 적용해서 첫 글자를 대문자로 함.
* 그 외 최근 주류 언어인 Java, C#의 관례에 따라 CamelCase를 적용함. */

# include <windef.h>

//------------------------------------------------------------------------------
// XingAPI DLL 함수
//------------------------------------------------------------------------------
typedef BOOL (__stdcall *F_BOOL)(); // ETK_IsConnected, ETK_Disconnnect, ETK_Logout 대체
typedef int (__stdcall *F_INT)();   // ETK_GetAccountListCount, ETK_GetLastError 대체
typedef BOOL (__stdcall *ETK_Connect)(HWND  hWnd, const char* pszSvr, int nPort, int nStartMsgID, int nTimeOut, int nSendMaxPacketSize);
typedef BOOL (__stdcall *ETK_Login)(HWND  hWnd, const char* pszID, const char* pszPwd, const char* pszCertPwd, int nType, BOOL bShowCertErrDlg);
typedef BOOL (__stdcall *ETK_Logout)(HWND hWnd);
typedef int (__stdcall *ETK_Request)(HWND hParentWnd, const char* pszTrCode, void* lpData, int nDataSize, BOOL bNext, const char* pszContinueKey, int nTimeOut);
typedef void (__stdcall *ETK_ReleaseRequestData)(int nRequestID);
typedef void (__stdcall *ETK_ReleaseMessageData)(LPARAM  lp);
typedef BOOL (__stdcall *ETK_AdviseRealData)(HWND  hWnd, const char* pszTrCode, const char* pszData, int nDataUnitLen);
typedef BOOL (__stdcall *ETK_UnadviseRealData)(HWND  hWnd, const char* pszTrCode, const char* pszData, int nDataUnitLen);
typedef BOOL (__stdcall *ETK_UnadviseWindow)(HWND  hWnd);
typedef BOOL (__stdcall *ETK_GetAccountList)(int nIndex, char* pszAcc, int nAccSize);
typedef void (__stdcall *ETK_GetAccountName)(LPCTSTR pszAcc, LPSTR pszAccName, int nAccNameSize);
typedef void (__stdcall *ETK_GetAcctDetailName)(LPCTSTR pszAcc, LPSTR pszAccName, int nAccNameSize);
typedef void (__stdcall *ETK_GetAcctNickName)(LPCTSTR pszAcc, LPSTR pszAccName, int nAccNameSize);
typedef void (__stdcall *ETK_GetServerName)(char* pszName);
typedef int (__stdcall *ETK_GetErrorMessage)(int nErrorCode, char* pszMsg, int nMsgSize);
typedef int (__stdcall *ETK_GetTRCountPerSec)(LPCTSTR pszCode);
//typedef int (__stdcall *ETK_RequestService)(HWND hWnd, LPCTSTR pszCode, LPCTSTR pszData);
//typedef int (__stdcall *ETK_RemoveService)(HWND hWnd, LPCTSTR pszCode, LPCTSTR pszData);
//typedef int (__stdcall *ETK_RequestLinkToHTS)(HWND hWnd, LPCTSTR pszLinkKey, LPCTSTR pszData, LPCTSTR pszFiller);
//typedef void (__stdcall *ETK_AdviseLinkFromHTS)(HWND hWnd);
//typedef void (__stdcall *ETK_UnadviseLinkFromHTS)();
typedef int (__stdcall *ETK_Decompress)(LPCTSTR pszSrc, LPCTSTR pszDest, int nSrcLen);

// '#pragma pack(push, 1)'이 적용되면 Go언어에서 읽을 수 없는 구조체들을
// 원래 메모리 저장방식을 사용해서 Go언어에서 읽을 수 있게 변환한 구조체.
typedef struct {    // 조회TR 수신 패킷
	int					RequestID;                  // Request ID
	int					DataLength;				    // 받은 데이터 크기
	int					TotalDataBufferSize;		// lpData에 할당된 크기
	int					ElapsedTime;				// 전송에서 수신까지 걸린시간(1/1000초)
	int					DataMode;					// 1:BLOCK MODE, 2:NON-BLOCK MODE
	char TrCode[10]; char _TrCode[1];			    // TR Code
	char				Cont[1];       			    // 다음조회 없음 : '0', 'N', 다음조회 있음 : '1', '2'
	char ContKey[18]; char _ContKey[1];		        // 연속키. Header타입이 B인 경우 이 값이 다음 조회 시 Key가 됨.
	char				None[31];                   // 사용자 데이터 (사용 안 함) 62
	char BlockName[16]; char _BlockName[1];        // Block 명, Block Mode 일때만 사용  17
	unsigned char*		Data;                       // 수신된 TR 데이터
} TR_DATA_UNPACKED;

typedef struct {    // 메시지 수신 패킷.
	int					RequestID;						// Request ID
	int					SystemError;				// 0:일반메시지, 1:시스템 에러 메시지
	char MsgCode[5]; char _MsgCode[1];             // 메시지 코드
	int					MsgLength;					// 메시지 데이터 길이
	char*		        MsgData;			// 메시지 데이터
} MSG_DATA_UNPACKED;

typedef struct {    // 실시간TR 수신 패킷
	char TrCode[3]; char _TrCode[1];		    // TR Code
	int					KeyLength;                 // 뭐지??
	char KeyData[32]; char _KeyData[1];         // 뭐지??
	char RegKey[32]; char _RegKey[1];          // 뭐지??
	int					DataLength;                // 받은 데이터 크기
	char*				Data;                    // 실시간 데이터
} REALTIME_DATA_UNPACKED;

// C 구조체 메모리 저장방식을 1바이트 단위로 설정.
// 이후에 '#pragma pack(pop)'으로 원래대로 되돌려야 함.
// 이 경우 Go언어에서 읽을 수 없는 경우가 종종 발생함.
#pragma pack(push, 1)

//------------------------------------------------------------------------------
// 기본 구조체
//------------------------------------------------------------------------------
typedef struct {    // 조회TR 수신 패킷
	int					RequestID;					// Request ID
	int					DataLength;				    // 받은 데이터 크기
	int					TotalDataBufferSize;		// lpData에 할당된 크기
	int					ElapsedTime;				// 전송에서 수신까지 걸린시간(1/1000초)
	int					DataMode;					// 현재 의미 없음. (1:BLOCK MODE, 2:NON-BLOCK MODE) 20
	char TrCode[10]; char _TrCode[1];			    // TR Code
	char				Cont[1];       			    // 다음조회 없음 : '0', 'N', 다음조회 있음 : '1', '2'
	char ContKey[18]; char _ContKey[1];		        // 연속키, Data Header가 B 인 경우에만 사용
	char				None[31];                   //szUserData[31]; // 사용자 데이터 (사용 안 함) 62
	char BlockName[16]; char _BlockName[1];		    // Block 명, Block Mode 일때만 사용  17
	unsigned char*		Data;                        // 수신된 TR 데이터
} TR_DATA;

typedef struct {    // 실시간TR 수신 패킷
	char TrCode[3]; char _TrCode[1];		    // TR Code
	int					KeyLength;                 // 뭐지??
	char KeyData[32]; char _KeyData[1];         // 뭐지??
	char RegKey[32]; char _RegKey[1];          // 뭐지??
	int					DataLength;                // 받은 데이터 크기
	char*				Data;                    // 실시간 데이터
} REALTIME_DATA;

typedef struct {    // 메시지 수신 패킷
	int					RequestID;			        // Request ID
	int					SystemError;				// 0:일반메시지, 1:시스템 에러 메시지
	char MsgCode[5]; char _MsgCode[1];              // 메시지 코드
	int					MsgLength;					// 메시지 데이터 길이
	char*		        MsgData;                // 메시지 데이터
} MSG_DATA;

typedef	struct {    // HTS-> API로 연동데이터 수신 패킷
    char                LinkName[32];             // 연동 키 ex) 주식 종목코드 연동 시, &STOCK_CODE
    char                LinkData[32];             // 연동 값 ex) 주식 종목코드 연동 시, 종목코드
    char                None[64];                   // 사용 안 함
} LINK_DATA;

//------------------------------------------------------------------------------
// 현물 정상주문 (CSPAT00600,ENCRYPT,SIGNATURE,HEADTYPE=B)
//------------------------------------------------------------------------------
typedef struct {
    char    acntNo[20];    //[string,   20] 계좌번호   StartPos 0, Length 20
    char    inptPwd[8];    //[string,    8] 입력비밀번호   StartPos 20, Length 8
    char    isuNo[12];    //[string,   12] 종목번호   StartPos 28, Length 12
    char    ordQty[16];    //[long  ,   16] 주문수량   StartPos 40, Length 16
    char    ordPrc[13];    //[double, 13.2] 주문가   StartPos 56, Length 13
    char    bnsTpCode[1];    //[string,    1] 매매구분   StartPos 69, Length 1
    char    ordprcPtnCode[2];    //[string,    2] 호가유형코드   StartPos 70, Length 2
    char    mgntrnCode[3];    //[string,    3] 신용거래코드   StartPos 72, Length 3
    char    loanDt[8];    //[string,    8] 대출일   StartPos 75, Length 8
    char    ordCndiTpCode[1];    //[string,    1] 주문조건구분   StartPos 83, Length 1
} CSPAT00600InBlock1;

typedef struct {
    char    recCnt[5];    //[long  ,    5] 레코드갯수   StartPos 0, Length 5
    char    acntNo[20];    //[string,   20] 계좌번호   StartPos 5, Length 20
    char    inptPwd[8];    //[string,    8] 입력비밀번호   StartPos 25, Length 8
    char    isuNo[12];    //[string,   12] 종목번호   StartPos 33, Length 12
    char    ordQty[16];    //[long  ,   16] 주문수량   StartPos 45, Length 16
    char    ordPrc[13];    //[double, 13.2] 주문가   StartPos 61, Length 13
    char    bnsTpCode[1];    //[string,    1] 매매구분   StartPos 74, Length 1
    char    ordprcPtnCode[2];    //[string,    2] 호가유형코드   StartPos 75, Length 2
    char    prgmOrdprcPtnCode[2];    //[string,    2] 프로그램호가유형코드   StartPos 77, Length 2
    char    stslAbleYn[1];    //[string,    1] 공매도가능여부   StartPos 79, Length 1
    char    stslOrdprcTpCode[1];    //[string,    1] 공매도호가구분   StartPos 80, Length 1
    char    commdaCode[2];    //[string,    2] 통신매체코드   StartPos 81, Length 2
    char    mgntrnCode[3];    //[string,    3] 신용거래코드   StartPos 83, Length 3
    char    loanDt[8];    //[string,    8] 대출일   StartPos 86, Length 8
    char    mbrNo[3];    //[string,    3] 회원번호   StartPos 94, Length 3
    char    ordCndiTpCode[1];    //[string,    1] 주문조건구분   StartPos 97, Length 1
    char    strtgCode[6];    //[string,    6] 전략코드   StartPos 98, Length 6
    char    grpId[20];    //[string,   20] 그룹ID   StartPos 104, Length 20
    char    ordSeqNo[10];    //[long  ,   10] 주문회차   StartPos 124, Length 10
    char    ptflNo[10];    //[long  ,   10] 포트폴리오번호   StartPos 134, Length 10
    char    bskNo[10];    //[long  ,   10] 바스켓번호   StartPos 144, Length 10
    char    trchNo[10];    //[long  ,   10] 트렌치번호   StartPos 154, Length 10
    char    itemNo[10];    //[long  ,   10] 아이템번호   StartPos 164, Length 10
    char    opDrtnNo[12];    //[string,   12] 운용지시번호   StartPos 174, Length 12
    char    lpYn[1];    //[string,    1] 유동성공급자여부   StartPos 186, Length 1
    char    cvrgTpCode[1];    //[string,    1] 반대매매구분   StartPos 187, Length 1
} CSPAT00600OutBlock1;

typedef struct {
    char    recCnt[5];    //[long  ,    5] 레코드갯수   StartPos 0, Length 5
    char    ordNo[10];    //[long  ,   10] 주문번호   StartPos 5, Length 10
    char    ordTime[9];    //[string,    9] 주문시각   StartPos 15, Length 9
    char    ordMktCode[2];    //[string,    2] 주문시장코드   StartPos 24, Length 2
    char    ordPtnCode[2];    //[string,    2] 주문유형코드   StartPos 26, Length 2
    char    shtnIsuNo[9];    //[string,    9] 단축종목번호   StartPos 28, Length 9
    char    mgempNo[9];    //[string,    9] 관리사원번호   StartPos 37, Length 9
    char    ordAmt[16];    //[long  ,   16] 주문금액   StartPos 46, Length 16
    char    spareOrdNo[10];    //[long  ,   10] 예비주문번호   StartPos 62, Length 10
    char    cvrgSeqno[10];    //[long  ,   10] 반대매매일련번호   StartPos 72, Length 10
    char    rsvOrdNo[10];    //[long  ,   10] 예약주문번호   StartPos 82, Length 10
    char    spotOrdQty[16];    //[long  ,   16] 실물주문수량   StartPos 92, Length 16
    char    ruseOrdQty[16];    //[long  ,   16] 재사용주문수량   StartPos 108, Length 16
    char    mnyOrdAmt[16];    //[long  ,   16] 현금주문금액   StartPos 124, Length 16
    char    substOrdAmt[16];    //[long  ,   16] 대용주문금액   StartPos 140, Length 16
    char    ruseOrdAmt[16];    //[long  ,   16] 재사용주문금액   StartPos 156, Length 16
    char    acntNm[40];    //[string,   40] 계좌명   StartPos 172, Length 40
    char    isuNm[40];    //[string,   40] 종목명   StartPos 212, Length 40
} CSPAT00600OutBlock2;

typedef struct {
    CSPAT00600OutBlock1	outBlock1;
    CSPAT00600OutBlock2	outBlock2;
} CSPAT00600OutBlockAll;

//------------------------------------------------------------------------------
// 현물 정정주문 (CSPAT00700,ENCRYPT,SIGNATURE,HEADTYPE=B)
//------------------------------------------------------------------------------
typedef struct {
    char    orgOrdNo[10];    //[long  ,   10] 원주문번호   StartPos 0, Length 10
    char    acntNo[20];    //[string,   20] 계좌번호   StartPos 10, Length 20
    char    inptPwd[8];    //[string,    8] 입력비밀번호   StartPos 30, Length 8
    char    isuNo[12];    //[string,   12] 종목번호   StartPos 38, Length 12
    char    ordQty[16];    //[long  ,   16] 주문수량   StartPos 50, Length 16
    char    ordprcPtnCode[2];    //[string,    2] 호가유형코드   StartPos 66, Length 2
    char    ordCndiTpCode[1];    //[string,    1] 주문조건구분   StartPos 68, Length 1
    char    ordPrc[13];    //[double, 13.2] 주문가   StartPos 69, Length 13
} CSPAT00700InBlock1;

typedef struct {
    char    recCnt[5];    //[long  ,    5] 레코드갯수   StartPos 0, Length 5
    char    orgOrdNo[10];    //[long  ,   10] 원주문번호   StartPos 5, Length 10
    char    acntNo[20];    //[string,   20] 계좌번호   StartPos 15, Length 20
    char    inptPwd[8];    //[string,    8] 입력비밀번호   StartPos 35, Length 8
    char    isuNo[12];    //[string,   12] 종목번호   StartPos 43, Length 12
    char    ordQty[16];    //[long  ,   16] 주문수량   StartPos 55, Length 16
    char    ordprcPtnCode[2];    //[string,    2] 호가유형코드   StartPos 71, Length 2
    char    ordCndiTpCode[1];    //[string,    1] 주문조건구분   StartPos 73, Length 1
    char    ordPrc[13];    //[double, 13.2] 주문가   StartPos 74, Length 13
    char    commdaCode[2];    //[string,    2] 통신매체코드   StartPos 87, Length 2
    char    strtgCode[6];    //[string,    6] 전략코드   StartPos 89, Length 6
    char    grpId[20];    //[string,   20] 그룹ID   StartPos 95, Length 20
    char    ordSeqNo[10];    //[long  ,   10] 주문회차   StartPos 115, Length 10
    char    ptflNo[10];    //[long  ,   10] 포트폴리오번호   StartPos 125, Length 10
    char    bskNo[10];    //[long  ,   10] 바스켓번호   StartPos 135, Length 10
    char    trchNo[10];    //[long  ,   10] 트렌치번호   StartPos 145, Length 10
    char    itemNo[10];    //[long  ,   10] 아이템번호   StartPos 155, Length 10
} CSPAT00700OutBlock1;

typedef struct {
    char    recCnt[5];    //[long  ,    5] 레코드갯수   StartPos 0, Length 5
    char    ordNo[10];    //[long  ,   10] 주문번호   StartPos 5, Length 10
    char    prntOrdNo[10];    //[long  ,   10] 모주문번호   StartPos 15, Length 10
    char    ordTime[9];    //[string,    9] 주문시각   StartPos 25, Length 9
    char    ordMktCode[2];    //[string,    2] 주문시장코드   StartPos 34, Length 2
    char    ordPtnCode[2];    //[string,    2] 주문유형코드   StartPos 36, Length 2
    char    shtnIsuNo[9];    //[string,    9] 단축종목번호   StartPos 38, Length 9
    char    prgmOrdprcPtnCode[2];    //[string,    2] 프로그램호가유형코드   StartPos 47, Length 2
    char    stslOrdprcTpCode[1];    //[string,    1] 공매도호가구분   StartPos 49, Length 1
    char    stslAbleYn[1];    //[string,    1] 공매도가능여부   StartPos 50, Length 1
    char    mgntrnCode[3];    //[string,    3] 신용거래코드   StartPos 51, Length 3
    char    loanDt[8];    //[string,    8] 대출일   StartPos 54, Length 8
    char    cvrgOrdTp[1];    //[string,    1] 반대매매주문구분   StartPos 62, Length 1
    char    lpYn[1];    //[string,    1] 유동성공급자여부   StartPos 63, Length 1
    char    mgempNo[9];    //[string,    9] 관리사원번호   StartPos 64, Length 9
    char    ordAmt[16];    //[long  ,   16] 주문금액   StartPos 73, Length 16
    char    bnsTpCode[1];    //[string,    1] 매매구분   StartPos 89, Length 1
    char    spareOrdNo[10];    //[long  ,   10] 예비주문번호   StartPos 90, Length 10
    char    cvrgSeqno[10];    //[long  ,   10] 반대매매일련번호   StartPos 100, Length 10
    char    rsvOrdNo[10];    //[long  ,   10] 예약주문번호   StartPos 110, Length 10
    char    mnyOrdAmt[16];    //[long  ,   16] 현금주문금액   StartPos 120, Length 16
    char    substOrdAmt[16];    //[long  ,   16] 대용주문금액   StartPos 136, Length 16
    char    ruseOrdAmt[16];    //[long  ,   16] 재사용주문금액   StartPos 152, Length 16
    char    acntNm[40];    //[string,   40] 계좌명   StartPos 168, Length 40
    char    isuNm[40];    //[string,   40] 종목명   StartPos 208, Length 40
} CSPAT00700OutBlock2;

typedef struct {
    CSPAT00700OutBlock1	outBlock1;
    CSPAT00700OutBlock2	outBlock2;
} CSPAT00700OutBlockAll;

//------------------------------------------------------------------------------
// 현물 취소주문 (CSPAT00800,ENCRYPT,SIGNATURE,HEADTYPE=B)
//------------------------------------------------------------------------------
typedef struct {
    char    orgOrdNo[10];    //[long  ,   10] 원주문번호   StartPos 0, Length 10
    char    acntNo[20];    //[string,   20] 계좌번호   StartPos 10, Length 20
    char    inptPwd[8];    //[string,    8] 입력비밀번호   StartPos 30, Length 8
    char    isuNo[12];    //[string,   12] 종목번호   StartPos 38, Length 12
    char    ordQty[16];    //[long  ,   16] 주문수량   StartPos 50, Length 16
} CSPAT00800InBlock1;

typedef struct {
    char    recCnt[5];    //[long  ,    5] 레코드갯수   StartPos 0, Length 5
    char    orgOrdNo[10];    //[long  ,   10] 원주문번호   StartPos 5, Length 10
    char    acntNo[20];    //[string,   20] 계좌번호   StartPos 15, Length 20
    char    inptPwd[8];    //[string,    8] 입력비밀번호   StartPos 35, Length 8
    char    isuNo[12];    //[string,   12] 종목번호   StartPos 43, Length 12
    char    ordQty[16];    //[long  ,   16] 주문수량   StartPos 55, Length 16
    char    commdaCode[2];    //[string,    2] 통신매체코드   StartPos 71, Length 2
    char    grpId[20];    //[string,   20] 그룹ID   StartPos 73, Length 20
    char    strtgCode[6];    //[string,    6] 전략코드   StartPos 93, Length 6
    char    ordSeqNo[10];    //[long  ,   10] 주문회차   StartPos 99, Length 10
    char    ptflNo[10];    //[long  ,   10] 포트폴리오번호   StartPos 109, Length 10
    char    bskNo[10];    //[long  ,   10] 바스켓번호   StartPos 119, Length 10
    char    trchNo[10];    //[long  ,   10] 트렌치번호   StartPos 129, Length 10
    char    itemNo[10];    //[long  ,   10] 아이템번호   StartPos 139, Length 10
} CSPAT00800OutBlock1;

typedef struct {
    char    recCnt[5];    //[long  ,    5] 레코드갯수   StartPos 0, Length 5
    char    ordNo[10];    //[long  ,   10] 주문번호   StartPos 5, Length 10
    char    prntOrdNo[10];    //[long  ,   10] 모주문번호   StartPos 15, Length 10
    char    ordTime[9];    //[string,    9] 주문시각   StartPos 25, Length 9
    char    ordMktCode[2];    //[string,    2] 주문시장코드   StartPos 34, Length 2
    char    ordPtnCode[2];    //[string,    2] 주문유형코드   StartPos 36, Length 2
    char    shtnIsuNo[9];    //[string,    9] 단축종목번호   StartPos 38, Length 9
    char    prgmOrdprcPtnCode[2];    //[string,    2] 프로그램호가유형코드   StartPos 47, Length 2
    char    stslOrdprcTpCode[1];    //[string,    1] 공매도호가구분   StartPos 49, Length 1
    char    stslAbleYn[1];    //[string,    1] 공매도가능여부   StartPos 50, Length 1
    char    mgntrnCode[3];    //[string,    3] 신용거래코드   StartPos 51, Length 3
    char    loanDt[8];    //[string,    8] 대출일   StartPos 54, Length 8
    char    cvrgOrdTp[1];    //[string,    1] 반대매매주문구분   StartPos 62, Length 1
    char    lpYn[1];    //[string,    1] 유동성공급자여부   StartPos 63, Length 1
    char    mgempNo[9];    //[string,    9] 관리사원번호   StartPos 64, Length 9
    char    bnsTpCode[1];    //[string,    1] 매매구분   StartPos 73, Length 1
    char    spareOrdNo[10];    //[long  ,   10] 예비주문번호   StartPos 74, Length 10
    char    cvrgSeqno[10];    //[long  ,   10] 반대매매일련번호   StartPos 84, Length 10
    char    rsvOrdNo[10];    //[long  ,   10] 예약주문번호   StartPos 94, Length 10
    char    acntNm[40];    //[string,   40] 계좌명   StartPos 104, Length 40
    char    isuNm[40];    //[string,   40] 종목명   StartPos 144, Length 40
} CSPAT00800OutBlock2;

typedef struct {
    CSPAT00800OutBlock1	outBlock1;
    CSPAT00800OutBlock2	outBlock2;
} CSPAT00800OutBlockAll;

//------------------------------------------------------------------------------
// 주식 주문 접수 실시간 정보 (SC0)
//------------------------------------------------------------------------------
typedef struct {
    char    lineseq             [  10];    // [long  ,   10] 라인일련번호                   StartPos 0, Length 10
    char    accno               [  11];    // [string,   11] 계좌번호                       StartPos 10, Length 11
    char    user                [   8];    // [string,    8] 조작자ID                       StartPos 21, Length 8
    char    len                 [   6];    // [long  ,    6] 헤더길이                       StartPos 29, Length 6
    char    gubun               [   1];    // [string,    1] 헤더구분                       StartPos 35, Length 1
    char    compress            [   1];    // [string,    1] 압축구분                       StartPos 36, Length 1
    char    encrypt             [   1];    // [string,    1] 암호구분                       StartPos 37, Length 1
    char    offset              [   3];    // [long  ,    3] 공통시작지점                   StartPos 38, Length 3
    char    trcode              [   8];    // [string,    8] TRCODE                         StartPos 41, Length 8
    char    compid              [   3];    // [string,    3] 이용사번호                     StartPos 49, Length 3
    char    userid              [  16];    // [string,   16] 사용자ID                       StartPos 52, Length 16
    char    media               [   2];    // [string,    2] 접속매체                       StartPos 68, Length 2
    char    ifid                [   3];    // [string,    3] I/F일련번호                    StartPos 70, Length 3
    char    seq                 [   9];    // [string,    9] 전문일련번호                   StartPos 73, Length 9
    char    trid                [  16];    // [string,   16] TR추적ID                       StartPos 82, Length 16
    char    pubip               [  12];    // [string,   12] 공인IP                         StartPos 98, Length 12
    char    prvip               [  12];    // [string,   12] 사설IP                         StartPos 110, Length 12
    char    pcbpno              [   3];    // [string,    3] 처리지점번호                   StartPos 122, Length 3
    char    bpno                [   3];    // [string,    3] 지점번호                       StartPos 125, Length 3
    char    termno              [   8];    // [string,    8] 단말번호                       StartPos 128, Length 8
    char    lang                [   1];    // [string,    1] 언어구분                       StartPos 136, Length 1
    char    proctm              [   9];    // [long  ,    9] AP처리시간                     StartPos 137, Length 9
    char    msgcode             [   4];    // [string,    4] 메세지코드                     StartPos 146, Length 4
    char    outgu               [   1];    // [string,    1] 메세지출력구분                 StartPos 150, Length 1
    char    compreq             [   1];    // [string,    1] 압축요청구분                   StartPos 151, Length 1
    char    funckey             [   4];    // [string,    4] 기능키                         StartPos 152, Length 4
    char    reqcnt              [   4];    // [long  ,    4] 요청레코드개수                 StartPos 156, Length 4
    char    filler              [   6];    // [string,    6] 예비영역                       StartPos 160, Length 6
    char    cont                [   1];    // [string,    1] 연속구분                       StartPos 166, Length 1
    char    contkey             [  18];    // [string,   18] 연속키값                       StartPos 167, Length 18
    char    varlen              [   2];    // [long  ,    2] 가변시스템길이                 StartPos 185, Length 2
    char    varhdlen            [   2];    // [long  ,    2] 가변해더길이                   StartPos 187, Length 2
    char    varmsglen           [   2];    // [long  ,    2] 가변메시지길이                 StartPos 189, Length 2
    char    trsrc               [   1];    // [string,    1] 조회발원지                     StartPos 191, Length 1
    char    eventid             [   4];    // [string,    4] I/F이벤트ID                    StartPos 192, Length 4
    char    ifinfo              [   4];    // [string,    4] I/F정보                        StartPos 196, Length 4
    char    filler1             [  41];    // [string,   41] 예비영역                       StartPos 200, Length 41
    char    ordchegb            [   2];    // [string,    2] 주문체결구분                   StartPos 241, Length 2
    char    marketgb            [   2];    // [string,    2] 시장구분                       StartPos 243, Length 2
    char    ordgb               [   2];    // [string,    2] 주문구분                       StartPos 245, Length 2
    char    orgordno            [  10];    // [long  ,   10] 원주문번호                     StartPos 247, Length 10
    char    accno1              [  11];    // [string,   11] 계좌번호                       StartPos 257, Length 11
    char    accno2              [   9];    // [string,    9] 계좌번호                       StartPos 268, Length 9
    char    passwd              [   8];    // [string,    8] 비밀번호                       StartPos 277, Length 8
    char    expcode             [  12];    // [string,   12] 종목번호                       StartPos 285, Length 12
    char    shtcode             [   9];    // [string,    9] 단축종목번호                   StartPos 297, Length 9
    char    hname               [  40];    // [string,   40] 종목명                         StartPos 306, Length 40
    char    ordqty              [  16];    // [long  ,   16] 주문수량                       StartPos 346, Length 16
    char    ordprice            [  13];    // [long  ,   13] 주문가격                       StartPos 362, Length 13
    char    hogagb              [   1];    // [string,    1] 주문조건                       StartPos 375, Length 1
    char    etfhogagb           [   2];    // [string,    2] 호가유형코드                   StartPos 376, Length 2
    char    pgmtype             [   2];    // [long  ,    2] 프로그램호가구분               StartPos 378, Length 2
    char    gmhogagb            [   1];    // [long  ,    1] 공매도호가구분                 StartPos 380, Length 1
    char    gmhogayn            [   1];    // [long  ,    1] 공매도가능여부                 StartPos 381, Length 1
    char    singb               [   3];    // [string,    3] 신용구분                       StartPos 382, Length 3
    char    loandt              [   8];    // [string,    8] 대출일                         StartPos 385, Length 8
    char    cvrgordtp           [   1];    // [string,    1] 반대매매주문구분               StartPos 393, Length 1
    char    strtgcode           [   6];    // [string,    6] 전략코드                       StartPos 394, Length 6
    char    groupid             [  20];    // [string,   20] 그룹ID                         StartPos 400, Length 20
    char    ordseqno            [  10];    // [long  ,   10] 주문회차                       StartPos 420, Length 10
    char    prtno               [  10];    // [long  ,   10] 포트폴리오번호                 StartPos 430, Length 10
    char    basketno            [  10];    // [long  ,   10] 바스켓번호                     StartPos 440, Length 10
    char    trchno              [  10];    // [long  ,   10] 트렌치번호                     StartPos 450, Length 10
    char    itemno              [  10];    // [long  ,   10] 아아템번호                     StartPos 460, Length 10
    char    brwmgmyn            [   1];    // [long  ,    1] 차입구분                       StartPos 470, Length 1
    char    mbrno               [   3];    // [long  ,    3] 회원사번호                     StartPos 471, Length 3
    char    procgb              [   1];    // [string,    1] 처리구분                       StartPos 474, Length 1
    char    admbrchno           [   3];    // [string,    3] 관리지점번호                   StartPos 475, Length 3
    char    futaccno            [  20];    // [string,   20] 선물계좌번호                   StartPos 478, Length 20
    char    futmarketgb         [   1];    // [string,    1] 선물상품구분                   StartPos 498, Length 1
    char    tongsingb           [   2];    // [string,    2] 통신매체구분                   StartPos 499, Length 2
    char    lpgb                [   1];    // [string,    1] 유동성공급자구분               StartPos 501, Length 1
    char    dummy               [  20];    // [string,   20] DUMMY                          StartPos 502, Length 20
    char    ordno               [  10];    // [long  ,   10] 주문번호                       StartPos 522, Length 10
    char    ordtm               [   9];    // [string,    9] 주문시각                       StartPos 532, Length 9
    char    prntordno           [  10];    // [long  ,   10] 모주문번호                     StartPos 541, Length 10
    char    mgempno             [   9];    // [string,    9] 관리사원번호                   StartPos 551, Length 9
    char    orgordundrqty       [  16];    // [long  ,   16] 원주문미체결수량               StartPos 560, Length 16
    char    orgordmdfyqty       [  16];    // [long  ,   16] 원주문정정수량                 StartPos 576, Length 16
    char    ordordcancelqty     [  16];    // [long  ,   16] 원주문취소수량                 StartPos 592, Length 16
    char    nmcpysndno          [  10];    // [long  ,   10] 비회원사송신번호               StartPos 608, Length 10
    char    ordamt              [  16];    // [long  ,   16] 주문금액                       StartPos 618, Length 16
    char    bnstp               [   1];    // [string,    1] 매매구분                       StartPos 634, Length 1
    char    spareordno          [  10];    // [long  ,   10] 예비주문번호                   StartPos 635, Length 10
    char    cvrgseqno           [  10];    // [long  ,   10] 반대매매일련번호               StartPos 645, Length 10
    char    rsvordno            [  10];    // [long  ,   10] 예약주문번호                   StartPos 655, Length 10
    char    mtordseqno          [  10];    // [long  ,   10] 복수주문일련번호               StartPos 665, Length 10
    char    spareordqty         [  16];    // [long  ,   16] 예비주문수량                   StartPos 675, Length 16
    char    orduserid           [  16];    // [string,   16] 주문사원번호                   StartPos 691, Length 16
    char    spotordqty          [  16];    // [long  ,   16] 실물주문수량                   StartPos 707, Length 16
    char    ordruseqty          [  16];    // [long  ,   16] 재사용주문수량                 StartPos 723, Length 16
    char    mnyordamt           [  16];    // [long  ,   16] 현금주문금액                   StartPos 739, Length 16
    char    ordsubstamt         [  16];    // [long  ,   16] 주문대용금액                   StartPos 755, Length 16
    char    ruseordamt          [  16];    // [long  ,   16] 재사용주문금액                 StartPos 771, Length 16
    char    ordcmsnamt          [  16];    // [long  ,   16] 수수료주문금액                 StartPos 787, Length 16
    char    crdtuseamt          [  16];    // [long  ,   16] 사용신용담보재사용금           StartPos 803, Length 16
    char    secbalqty           [  16];    // [long  ,   16] 잔고수량                       StartPos 819, Length 16
    char    spotordableqty      [  16];    // [long  ,   16] 실물가능수량                   StartPos 835, Length 16
    char    ordableruseqty      [  16];    // [long  ,   16] 재사용가능수량(매도)           StartPos 851, Length 16
    char    flctqty             [  16];    // [long  ,   16] 변동수량                       StartPos 867, Length 16
    char    secbalqtyd2         [  16];    // [long  ,   16] 잔고수량(D2)                   StartPos 883, Length 16
    char    sellableqty         [  16];    // [long  ,   16] 매도주문가능수량               StartPos 899, Length 16
    char    unercsellordqty     [  16];    // [long  ,   16] 미체결매도주문수량             StartPos 915, Length 16
    char    avrpchsprc          [  13];    // [long  ,   13] 평균매입가                     StartPos 931, Length 13
    char    pchsamt             [  16];    // [long  ,   16] 매입금액                       StartPos 944, Length 16
    char    deposit             [  16];    // [long  ,   16] 예수금                         StartPos 960, Length 16
    char    substamt            [  16];    // [long  ,   16] 대용금                         StartPos 976, Length 16
    char    csgnmnymgn          [  16];    // [long  ,   16] 위탁증거금현금                 StartPos 992, Length 16
    char    csgnsubstmgn        [  16];    // [long  ,   16] 위탁증거금대용                 StartPos 1008, Length 16
    char    crdtpldgruseamt     [  16];    // [long  ,   16] 신용담보재사용금               StartPos 1024, Length 16
    char    ordablemny          [  16];    // [long  ,   16] 주문가능현금                   StartPos 1040, Length 16
    char    ordablesubstamt     [  16];    // [long  ,   16] 주문가능대용                   StartPos 1056, Length 16
    char    ruseableamt         [  16];    // [long  ,   16] 재사용가능금액                 StartPos 1072, Length 16
} SC0_OutBlock;

//------------------------------------------------------------------------------
// 주식 주문 체결 실시간 정보 (SC1)
//------------------------------------------------------------------------------
typedef struct {
    char    lineseq             [  10];    // [long  ,   10] 라인일련번호                   StartPos 0, Length 10
    char    accno               [  11];    // [string,   11] 계좌번호                       StartPos 10, Length 11
    char    user                [   8];    // [string,    8] 조작자ID                       StartPos 21, Length 8
    char    len                 [   6];    // [long  ,    6] 헤더길이                       StartPos 29, Length 6
    char    gubun               [   1];    // [string,    1] 헤더구분                       StartPos 35, Length 1
    char    compress            [   1];    // [string,    1] 압축구분                       StartPos 36, Length 1
    char    encrypt             [   1];    // [string,    1] 암호구분                       StartPos 37, Length 1
    char    offset              [   3];    // [long  ,    3] 공통시작지점                   StartPos 38, Length 3
    char    trcode              [   8];    // [string,    8] TRCODE                         StartPos 41, Length 8
    char    compid              [   3];    // [string,    3] 이용사번호                     StartPos 49, Length 3
    char    userid              [  16];    // [string,   16] 사용자ID                       StartPos 52, Length 16
    char    media               [   2];    // [string,    2] 접속매체                       StartPos 68, Length 2
    char    ifid                [   3];    // [string,    3] I/F일련번호                    StartPos 70, Length 3
    char    seq                 [   9];    // [string,    9] 전문일련번호                   StartPos 73, Length 9
    char    trid                [  16];    // [string,   16] TR추적ID                       StartPos 82, Length 16
    char    pubip               [  12];    // [string,   12] 공인IP                         StartPos 98, Length 12
    char    prvip               [  12];    // [string,   12] 사설IP                         StartPos 110, Length 12
    char    pcbpno              [   3];    // [string,    3] 처리지점번호                   StartPos 122, Length 3
    char    bpno                [   3];    // [string,    3] 지점번호                       StartPos 125, Length 3
    char    termno              [   8];    // [string,    8] 단말번호                       StartPos 128, Length 8
    char    lang                [   1];    // [string,    1] 언어구분                       StartPos 136, Length 1
    char    proctm              [   9];    // [long  ,    9] AP처리시간                     StartPos 137, Length 9
    char    msgcode             [   4];    // [string,    4] 메세지코드                     StartPos 146, Length 4
    char    outgu               [   1];    // [string,    1] 메세지출력구분                 StartPos 150, Length 1
    char    compreq             [   1];    // [string,    1] 압축요청구분                   StartPos 151, Length 1
    char    funckey             [   4];    // [string,    4] 기능키                         StartPos 152, Length 4
    char    reqcnt              [   4];    // [long  ,    4] 요청레코드개수                 StartPos 156, Length 4
    char    filler              [   6];    // [string,    6] 예비영역                       StartPos 160, Length 6
    char    cont                [   1];    // [string,    1] 연속구분                       StartPos 166, Length 1
    char    contkey             [  18];    // [string,   18] 연속키값                       StartPos 167, Length 18
    char    varlen              [   2];    // [long  ,    2] 가변시스템길이                 StartPos 185, Length 2
    char    varhdlen            [   2];    // [long  ,    2] 가변해더길이                   StartPos 187, Length 2
    char    varmsglen           [   2];    // [long  ,    2] 가변메시지길이                 StartPos 189, Length 2
    char    trsrc               [   1];    // [string,    1] 조회발원지                     StartPos 191, Length 1
    char    eventid             [   4];    // [string,    4] I/F이벤트ID                    StartPos 192, Length 4
    char    ifinfo              [   4];    // [string,    4] I/F정보                        StartPos 196, Length 4
    char    filler1             [  41];    // [string,   41] 예비영역                       StartPos 200, Length 41
    char    ordxctptncode       [   2];    // [string,    2] 주문체결유형코드               StartPos 241, Length 2
    char    ordmktcode          [   2];    // [string,    2] 주문시장코드                   StartPos 243, Length 2
    char    ordptncode          [   2];    // [string,    2] 주문유형코드                   StartPos 245, Length 2
    char    mgmtbrnno           [   3];    // [string,    3] 관리지점번호                   StartPos 247, Length 3
    char    accno1              [  11];    // [string,   11] 계좌번호                       StartPos 250, Length 11
    char    accno2              [   9];    // [string,    9] 계좌번호                       StartPos 261, Length 9
    char    acntnm              [  40];    // [string,   40] 계좌명                         StartPos 270, Length 40
    char    Isuno               [  12];    // [string,   12] 종목번호                       StartPos 310, Length 12
    char    Isunm               [  40];    // [string,   40] 종목명                         StartPos 322, Length 40
    char    ordno               [  10];    // [long  ,   10] 주문번호                       StartPos 362, Length 10
    char    orgordno            [  10];    // [long  ,   10] 원주문번호                     StartPos 372, Length 10
    char    execno              [  10];    // [long  ,   10] 체결번호                       StartPos 382, Length 10
    char    ordqty              [  16];    // [long  ,   16] 주문수량                       StartPos 392, Length 16
    char    ordprc              [  13];    // [long  ,   13] 주문가격                       StartPos 408, Length 13
    char    execqty             [  16];    // [long  ,   16] 체결수량                       StartPos 421, Length 16
    char    execprc             [  13];    // [long  ,   13] 체결가격                       StartPos 437, Length 13
    char    mdfycnfqty          [  16];    // [long  ,   16] 정정확인수량                   StartPos 450, Length 16
    char    mdfycnfprc          [  16];    // [long  ,   16] 정정확인가격                   StartPos 466, Length 16
    char    canccnfqty          [  16];    // [long  ,   16] 취소확인수량                   StartPos 482, Length 16
    char    rjtqty              [  16];    // [long  ,   16] 거부수량                       StartPos 498, Length 16
    char    ordtrxptncode       [   4];    // [long  ,    4] 주문처리유형코드               StartPos 514, Length 4
    char    mtiordseqno         [  10];    // [long  ,   10] 복수주문일련번호               StartPos 518, Length 10
    char    ordcndi             [   1];    // [string,    1] 주문조건                       StartPos 528, Length 1
    char    ordprcptncode       [   2];    // [string,    2] 호가유형코드                   StartPos 529, Length 2
    char    nsavtrdqty          [  16];    // [long  ,   16] 비저축체결수량                 StartPos 531, Length 16
    char    shtnIsuno           [   9];    // [string,    9] 단축종목번호                   StartPos 547, Length 9
    char    opdrtnno            [  12];    // [string,   12] 운용지시번호                   StartPos 556, Length 12
    char    cvrgordtp           [   1];    // [string,    1] 반대매매주문구분               StartPos 568, Length 1
    char    unercqty            [  16];    // [long  ,   16] 미체결수량(주문)               StartPos 569, Length 16
    char    orgordunercqty      [  16];    // [long  ,   16] 원주문미체결수량               StartPos 585, Length 16
    char    orgordmdfyqty       [  16];    // [long  ,   16] 원주문정정수량                 StartPos 601, Length 16
    char    orgordcancqty       [  16];    // [long  ,   16] 원주문취소수량                 StartPos 617, Length 16
    char    ordavrexecprc       [  13];    // [long  ,   13] 주문평균체결가격               StartPos 633, Length 13
    char    ordamt              [  16];    // [long  ,   16] 주문금액                       StartPos 646, Length 16
    char    stdIsuno            [  12];    // [string,   12] 표준종목번호                   StartPos 662, Length 12
    char    bfstdIsuno          [  12];    // [string,   12] 전표준종목번호                 StartPos 674, Length 12
    char    bnstp               [   1];    // [string,    1] 매매구분                       StartPos 686, Length 1
    char    ordtrdptncode       [   2];    // [string,    2] 주문거래유형코드               StartPos 687, Length 2
    char    mgntrncode          [   3];    // [string,    3] 신용거래코드                   StartPos 689, Length 3
    char    adduptp             [   2];    // [string,    2] 수수료합산코드                 StartPos 692, Length 2
    char    commdacode          [   2];    // [string,    2] 통신매체코드                   StartPos 694, Length 2
    char    Loandt              [   8];    // [string,    8] 대출일                         StartPos 696, Length 8
    char    mbrnmbrno           [   3];    // [long  ,    3] 회원/비회원사번호              StartPos 704, Length 3
    char    ordacntno           [  20];    // [string,   20] 주문계좌번호                   StartPos 707, Length 20
    char    agrgbrnno           [   3];    // [string,    3] 집계지점번호                   StartPos 727, Length 3
    char    mgempno             [   9];    // [string,    9] 관리사원번호                   StartPos 730, Length 9
    char    futsLnkbrnno        [   3];    // [string,    3] 선물연계지점번호               StartPos 739, Length 3
    char    futsLnkacntno       [  20];    // [string,   20] 선물연계계좌번호               StartPos 742, Length 20
    char    futsmkttp           [   1];    // [string,    1] 선물시장구분                   StartPos 762, Length 1
    char    regmktcode          [   2];    // [string,    2] 등록시장코드                   StartPos 763, Length 2
    char    mnymgnrat           [   7];    // [long  ,    7] 현금증거금률                   StartPos 765, Length 7
    char    substmgnrat         [   9];    // [long  ,    9] 대용증거금률                   StartPos 772, Length 9
    char    mnyexecamt          [  16];    // [long  ,   16] 현금체결금액                   StartPos 781, Length 16
    char    ubstexecamt         [  16];    // [long  ,   16] 대용체결금액                   StartPos 797, Length 16
    char    cmsnamtexecamt      [  16];    // [long  ,   16] 수수료체결금액                 StartPos 813, Length 16
    char    crdtpldgexecamt     [  16];    // [long  ,   16] 신용담보체결금액               StartPos 829, Length 16
    char    crdtexecamt         [  16];    // [long  ,   16] 신용체결금액                   StartPos 845, Length 16
    char    prdayruseexecval    [  16];    // [long  ,   16] 전일재사용체결금액             StartPos 861, Length 16
    char    crdayruseexecval    [  16];    // [long  ,   16] 금일재사용체결금액             StartPos 877, Length 16
    char    spotexecqty         [  16];    // [long  ,   16] 실물체결수량                   StartPos 893, Length 16
    char    stslexecqty         [  16];    // [long  ,   16] 공매도체결수량                 StartPos 909, Length 16
    char    strtgcode           [   6];    // [string,    6] 전략코드                       StartPos 925, Length 6
    char    grpId               [  20];    // [string,   20] 그룹Id                         StartPos 931, Length 20
    char    ordseqno            [  10];    // [long  ,   10] 주문회차                       StartPos 951, Length 10
    char    ptflno              [  10];    // [long  ,   10] 포트폴리오번호                 StartPos 961, Length 10
    char    bskno               [  10];    // [long  ,   10] 바스켓번호                     StartPos 971, Length 10
    char    trchno              [  10];    // [long  ,   10] 트렌치번호                     StartPos 981, Length 10
    char    itemno              [  10];    // [long  ,   10] 아이템번호                     StartPos 991, Length 10
    char    orduserId           [  16];    // [string,   16] 주문자Id                       StartPos 1001, Length 16
    char    brwmgmtYn           [   1];    // [long  ,    1] 차입관리여부                   StartPos 1017, Length 1
    char    frgrunqno           [   6];    // [string,    6] 외국인고유번호                 StartPos 1018, Length 6
    char    trtzxLevytp         [   1];    // [string,    1] 거래세징수구분                 StartPos 1024, Length 1
    char    lptp                [   1];    // [string,    1] 유동성공급자구분               StartPos 1025, Length 1
    char    exectime            [   9];    // [string,    9] 체결시각                       StartPos 1026, Length 9
    char    rcptexectime        [   9];    // [string,    9] 거래소수신체결시각             StartPos 1035, Length 9
    char    rmndLoanamt         [  16];    // [long  ,   16] 잔여대출금액                   StartPos 1044, Length 16
    char    secbalqty           [  16];    // [long  ,   16] 잔고수량                       StartPos 1060, Length 16
    char    spotordableqty      [  16];    // [long  ,   16] 실물가능수량                   StartPos 1076, Length 16
    char    ordableruseqty      [  16];    // [long  ,   16] 재사용가능수량(매도)           StartPos 1092, Length 16
    char    flctqty             [  16];    // [long  ,   16] 변동수량                       StartPos 1108, Length 16
    char    secbalqtyd2         [  16];    // [long  ,   16] 잔고수량(d2)                   StartPos 1124, Length 16
    char    sellableqty         [  16];    // [long  ,   16] 매도주문가능수량               StartPos 1140, Length 16
    char    unercsellordqty     [  16];    // [long  ,   16] 미체결매도주문수량             StartPos 1156, Length 16
    char    avrpchsprc          [  13];    // [long  ,   13] 평균매입가                     StartPos 1172, Length 13
    char    pchsant             [  16];    // [long  ,   16] 매입금액                       StartPos 1185, Length 16
    char    deposit             [  16];    // [long  ,   16] 예수금                         StartPos 1201, Length 16
    char    substamt            [  16];    // [long  ,   16] 대용금                         StartPos 1217, Length 16
    char    csgnmnymgn          [  16];    // [long  ,   16] 위탁증거금현금                 StartPos 1233, Length 16
    char    csgnsubstmgn        [  16];    // [long  ,   16] 위탁증거금대용                 StartPos 1249, Length 16
    char    crdtpldgruseamt     [  16];    // [long  ,   16] 신용담보재사용금               StartPos 1265, Length 16
    char    ordablemny          [  16];    // [long  ,   16] 주문가능현금                   StartPos 1281, Length 16
    char    ordablesubstamt     [  16];    // [long  ,   16] 주문가능대용                   StartPos 1297, Length 16
    char    ruseableamt         [  16];    // [long  ,   16] 재사용가능금액                 StartPos 1313, Length 16
} SC1_OutBlock;

//------------------------------------------------------------------------------
// 주식 주문 정정 실시간 정보 (SC2)
//------------------------------------------------------------------------------
typedef struct _SC2_OutBlock
{
    char    lineseq             [  10];    // [long  ,   10] 라인일련번호                   StartPos 0, Length 10
    char    accno               [  11];    // [string,   11] 계좌번호                       StartPos 10, Length 11
    char    user                [   8];    // [string,    8] 조작자ID                       StartPos 21, Length 8
    char    len                 [   6];    // [long  ,    6] 헤더길이                       StartPos 29, Length 6
    char    gubun               [   1];    // [string,    1] 헤더구분                       StartPos 35, Length 1
    char    compress            [   1];    // [string,    1] 압축구분                       StartPos 36, Length 1
    char    encrypt             [   1];    // [string,    1] 암호구분                       StartPos 37, Length 1
    char    offset              [   3];    // [long  ,    3] 공통시작지점                   StartPos 38, Length 3
    char    trcode              [   8];    // [string,    8] TRCODE                         StartPos 41, Length 8
    char    compid              [   3];    // [string,    3] 이용사번호                     StartPos 49, Length 3
    char    userid              [  16];    // [string,   16] 사용자ID                       StartPos 52, Length 16
    char    media               [   2];    // [string,    2] 접속매체                       StartPos 68, Length 2
    char    ifid                [   3];    // [string,    3] I/F일련번호                    StartPos 70, Length 3
    char    seq                 [   9];    // [string,    9] 전문일련번호                   StartPos 73, Length 9
    char    trid                [  16];    // [string,   16] TR추적ID                       StartPos 82, Length 16
    char    pubip               [  12];    // [string,   12] 공인IP                         StartPos 98, Length 12
    char    prvip               [  12];    // [string,   12] 사설IP                         StartPos 110, Length 12
    char    pcbpno              [   3];    // [string,    3] 처리지점번호                   StartPos 122, Length 3
    char    bpno                [   3];    // [string,    3] 지점번호                       StartPos 125, Length 3
    char    termno              [   8];    // [string,    8] 단말번호                       StartPos 128, Length 8
    char    lang                [   1];    // [string,    1] 언어구분                       StartPos 136, Length 1
    char    proctm              [   9];    // [long  ,    9] AP처리시간                     StartPos 137, Length 9
    char    msgcode             [   4];    // [string,    4] 메세지코드                     StartPos 146, Length 4
    char    outgu               [   1];    // [string,    1] 메세지출력구분                 StartPos 150, Length 1
    char    compreq             [   1];    // [string,    1] 압축요청구분                   StartPos 151, Length 1
    char    funckey             [   4];    // [string,    4] 기능키                         StartPos 152, Length 4
    char    reqcnt              [   4];    // [long  ,    4] 요청레코드개수                 StartPos 156, Length 4
    char    filler              [   6];    // [string,    6] 예비영역                       StartPos 160, Length 6
    char    cont                [   1];    // [string,    1] 연속구분                       StartPos 166, Length 1
    char    contkey             [  18];    // [string,   18] 연속키값                       StartPos 167, Length 18
    char    varlen              [   2];    // [long  ,    2] 가변시스템길이                 StartPos 185, Length 2
    char    varhdlen            [   2];    // [long  ,    2] 가변해더길이                   StartPos 187, Length 2
    char    varmsglen           [   2];    // [long  ,    2] 가변메시지길이                 StartPos 189, Length 2
    char    trsrc               [   1];    // [string,    1] 조회발원지                     StartPos 191, Length 1
    char    eventid             [   4];    // [string,    4] I/F이벤트ID                    StartPos 192, Length 4
    char    ifinfo              [   4];    // [string,    4] I/F정보                        StartPos 196, Length 4
    char    filler1             [  41];    // [string,   41] 예비영역                       StartPos 200, Length 41
    char    ordxctptncode       [   2];    // [string,    2] 주문체결유형코드               StartPos 241, Length 2
    char    ordmktcode          [   2];    // [string,    2] 주문시장코드                   StartPos 243, Length 2
    char    ordptncode          [   2];    // [string,    2] 주문유형코드                   StartPos 245, Length 2
    char    mgmtbrnno           [   3];    // [string,    3] 관리지점번호                   StartPos 247, Length 3
    char    accno1              [  11];    // [string,   11] 계좌번호                       StartPos 250, Length 11
    char    accno2              [   9];    // [string,    9] 계좌번호                       StartPos 261, Length 9
    char    acntnm              [  40];    // [string,   40] 계좌명                         StartPos 270, Length 40
    char    Isuno               [  12];    // [string,   12] 종목번호                       StartPos 310, Length 12
    char    Isunm               [  40];    // [string,   40] 종목명                         StartPos 322, Length 40
    char    ordno               [  10];    // [long  ,   10] 주문번호                       StartPos 362, Length 10
    char    orgordno            [  10];    // [long  ,   10] 원주문번호                     StartPos 372, Length 10
    char    execno              [  10];    // [long  ,   10] 체결번호                       StartPos 382, Length 10
    char    ordqty              [  16];    // [long  ,   16] 주문수량                       StartPos 392, Length 16
    char    ordprc              [  13];    // [long  ,   13] 주문가격                       StartPos 408, Length 13
    char    execqty             [  16];    // [long  ,   16] 체결수량                       StartPos 421, Length 16
    char    execprc             [  13];    // [long  ,   13] 체결가격                       StartPos 437, Length 13
    char    mdfycnfqty          [  16];    // [long  ,   16] 정정확인수량                   StartPos 450, Length 16
    char    mdfycnfprc          [  16];    // [long  ,   16] 정정확인가격                   StartPos 466, Length 16
    char    canccnfqty          [  16];    // [long  ,   16] 취소확인수량                   StartPos 482, Length 16
    char    rjtqty              [  16];    // [long  ,   16] 거부수량                       StartPos 498, Length 16
    char    ordtrxptncode       [   4];    // [long  ,    4] 주문처리유형코드               StartPos 514, Length 4
    char    mtiordseqno         [  10];    // [long  ,   10] 복수주문일련번호               StartPos 518, Length 10
    char    ordcndi             [   1];    // [string,    1] 주문조건                       StartPos 528, Length 1
    char    ordprcptncode       [   2];    // [string,    2] 호가유형코드                   StartPos 529, Length 2
    char    nsavtrdqty          [  16];    // [long  ,   16] 비저축체결수량                 StartPos 531, Length 16
    char    shtnIsuno           [   9];    // [string,    9] 단축종목번호                   StartPos 547, Length 9
    char    opdrtnno            [  12];    // [string,   12] 운용지시번호                   StartPos 556, Length 12
    char    cvrgordtp           [   1];    // [string,    1] 반대매매주문구분               StartPos 568, Length 1
    char    unercqty            [  16];    // [long  ,   16] 미체결수량(주문)               StartPos 569, Length 16
    char    orgordunercqty      [  16];    // [long  ,   16] 원주문미체결수량               StartPos 585, Length 16
    char    orgordmdfyqty       [  16];    // [long  ,   16] 원주문정정수량                 StartPos 601, Length 16
    char    orgordcancqty       [  16];    // [long  ,   16] 원주문취소수량                 StartPos 617, Length 16
    char    ordavrexecprc       [  13];    // [long  ,   13] 주문평균체결가격               StartPos 633, Length 13
    char    ordamt              [  16];    // [long  ,   16] 주문금액                       StartPos 646, Length 16
    char    stdIsuno            [  12];    // [string,   12] 표준종목번호                   StartPos 662, Length 12
    char    bfstdIsuno          [  12];    // [string,   12] 전표준종목번호                 StartPos 674, Length 12
    char    bnstp               [   1];    // [string,    1] 매매구분                       StartPos 686, Length 1
    char    ordtrdptncode       [   2];    // [string,    2] 주문거래유형코드               StartPos 687, Length 2
    char    mgntrncode          [   3];    // [string,    3] 신용거래코드                   StartPos 689, Length 3
    char    adduptp             [   2];    // [string,    2] 수수료합산코드                 StartPos 692, Length 2
    char    commdacode          [   2];    // [string,    2] 통신매체코드                   StartPos 694, Length 2
    char    Loandt              [   8];    // [string,    8] 대출일                         StartPos 696, Length 8
    char    mbrnmbrno           [   3];    // [long  ,    3] 회원/비회원사번호              StartPos 704, Length 3
    char    ordacntno           [  20];    // [string,   20] 주문계좌번호                   StartPos 707, Length 20
    char    agrgbrnno           [   3];    // [string,    3] 집계지점번호                   StartPos 727, Length 3
    char    mgempno             [   9];    // [string,    9] 관리사원번호                   StartPos 730, Length 9
    char    futsLnkbrnno        [   3];    // [string,    3] 선물연계지점번호               StartPos 739, Length 3
    char    futsLnkacntno       [  20];    // [string,   20] 선물연계계좌번호               StartPos 742, Length 20
    char    futsmkttp           [   1];    // [string,    1] 선물시장구분                   StartPos 762, Length 1
    char    regmktcode          [   2];    // [string,    2] 등록시장코드                   StartPos 763, Length 2
    char    mnymgnrat           [   7];    // [long  ,    7] 현금증거금률                   StartPos 765, Length 7
    char    substmgnrat         [   9];    // [long  ,    9] 대용증거금률                   StartPos 772, Length 9
    char    mnyexecamt          [  16];    // [long  ,   16] 현금체결금액                   StartPos 781, Length 16
    char    ubstexecamt         [  16];    // [long  ,   16] 대용체결금액                   StartPos 797, Length 16
    char    cmsnamtexecamt      [  16];    // [long  ,   16] 수수료체결금액                 StartPos 813, Length 16
    char    crdtpldgexecamt     [  16];    // [long  ,   16] 신용담보체결금액               StartPos 829, Length 16
    char    crdtexecamt         [  16];    // [long  ,   16] 신용체결금액                   StartPos 845, Length 16
    char    prdayruseexecval    [  16];    // [long  ,   16] 전일재사용체결금액             StartPos 861, Length 16
    char    crdayruseexecval    [  16];    // [long  ,   16] 금일재사용체결금액             StartPos 877, Length 16
    char    spotexecqty         [  16];    // [long  ,   16] 실물체결수량                   StartPos 893, Length 16
    char    stslexecqty         [  16];    // [long  ,   16] 공매도체결수량                 StartPos 909, Length 16
    char    strtgcode           [   6];    // [string,    6] 전략코드                       StartPos 925, Length 6
    char    grpId               [  20];    // [string,   20] 그룹Id                         StartPos 931, Length 20
    char    ordseqno            [  10];    // [long  ,   10] 주문회차                       StartPos 951, Length 10
    char    ptflno              [  10];    // [long  ,   10] 포트폴리오번호                 StartPos 961, Length 10
    char    bskno               [  10];    // [long  ,   10] 바스켓번호                     StartPos 971, Length 10
    char    trchno              [  10];    // [long  ,   10] 트렌치번호                     StartPos 981, Length 10
    char    itemno              [  10];    // [long  ,   10] 아이템번호                     StartPos 991, Length 10
    char    orduserId           [  16];    // [string,   16] 주문자Id                       StartPos 1001, Length 16
    char    brwmgmtYn           [   1];    // [long  ,    1] 차입관리여부                   StartPos 1017, Length 1
    char    frgrunqno           [   6];    // [string,    6] 외국인고유번호                 StartPos 1018, Length 6
    char    trtzxLevytp         [   1];    // [string,    1] 거래세징수구분                 StartPos 1024, Length 1
    char    lptp                [   1];    // [string,    1] 유동성공급자구분               StartPos 1025, Length 1
    char    exectime            [   9];    // [string,    9] 체결시각                       StartPos 1026, Length 9
    char    rcptexectime        [   9];    // [string,    9] 거래소수신체결시각             StartPos 1035, Length 9
    char    rmndLoanamt         [  16];    // [long  ,   16] 잔여대출금액                   StartPos 1044, Length 16
    char    secbalqty           [  16];    // [long  ,   16] 잔고수량                       StartPos 1060, Length 16
    char    spotordableqty      [  16];    // [long  ,   16] 실물가능수량                   StartPos 1076, Length 16
    char    ordableruseqty      [  16];    // [long  ,   16] 재사용가능수량(매도)           StartPos 1092, Length 16
    char    flctqty             [  16];    // [long  ,   16] 변동수량                       StartPos 1108, Length 16
    char    secbalqtyd2         [  16];    // [long  ,   16] 잔고수량(d2)                   StartPos 1124, Length 16
    char    sellableqty         [  16];    // [long  ,   16] 매도주문가능수량               StartPos 1140, Length 16
    char    unercsellordqty     [  16];    // [long  ,   16] 미체결매도주문수량             StartPos 1156, Length 16
    char    avrpchsprc          [  13];    // [long  ,   13] 평균매입가                     StartPos 1172, Length 13
    char    pchsant             [  16];    // [long  ,   16] 매입금액                       StartPos 1185, Length 16
    char    deposit             [  16];    // [long  ,   16] 예수금                         StartPos 1201, Length 16
    char    substamt            [  16];    // [long  ,   16] 대용금                         StartPos 1217, Length 16
    char    csgnmnymgn          [  16];    // [long  ,   16] 위탁증거금현금                 StartPos 1233, Length 16
    char    csgnsubstmgn        [  16];    // [long  ,   16] 위탁증거금대용                 StartPos 1249, Length 16
    char    crdtpldgruseamt     [  16];    // [long  ,   16] 신용담보재사용금               StartPos 1265, Length 16
    char    ordablemny          [  16];    // [long  ,   16] 주문가능현금                   StartPos 1281, Length 16
    char    ordablesubstamt     [  16];    // [long  ,   16] 주문가능대용                   StartPos 1297, Length 16
    char    ruseableamt         [  16];    // [long  ,   16] 재사용가능금액                 StartPos 1313, Length 16
} SC2_OutBlock;

//------------------------------------------------------------------------------
// 주식 주문 취소 실시간 정보 (SC3)
//------------------------------------------------------------------------------
typedef struct {
    char    lineseq[10];    //[long  ,   10] 라인일련번호   StartPos 0, Length 10
    char    accno[11];    //[string,   11] 계좌번호   StartPos 10, Length 11
    char    user[8];    //[string,    8] 조작자ID   StartPos 21, Length 8
    char    len[6];    //[long  ,    6] 헤더길이   StartPos 29, Length 6
    char    gubun[1];    //[string,    1] 헤더구분   StartPos 35, Length 1
    char    compress[1];    //[string,    1] 압축구분   StartPos 36, Length 1
    char    encrypt[1];    //[string,    1] 암호구분   StartPos 37, Length 1
    char    offset[3];    //[long  ,    3] 공통시작지점   StartPos 38, Length 3
    char    trcode[8];    //[string,    8] TRCODE   StartPos 41, Length 8
    char    compid[3];    //[string,    3] 이용사번호   StartPos 49, Length 3
    char    userid[16];    //[string,   16] 사용자ID   StartPos 52, Length 16
    char    media[2];    //[string,    2] 접속매체   StartPos 68, Length 2
    char    ifid[3];    //[string,    3] I/F일련번호   StartPos 70, Length 3
    char    seq[9];    //[string,    9] 전문일련번호   StartPos 73, Length 9
    char    trid[16];    //[string,   16] TR추적ID   StartPos 82, Length 16
    char    pubip[12];    //[string,   12] 공인IP   StartPos 98, Length 12
    char    prvip[12];    //[string,   12] 사설IP   StartPos 110, Length 12
    char    pcbpno[3];    //[string,    3] 처리지점번호   StartPos 122, Length 3
    char    bpno[3];    //[string,    3] 지점번호   StartPos 125, Length 3
    char    termno[8];    //[string,    8] 단말번호   StartPos 128, Length 8
    char    lang[1];    //[string,    1] 언어구분   StartPos 136, Length 1
    char    proctm[9];    //[long  ,    9] AP처리시간   StartPos 137, Length 9
    char    msgcode[4];    //[string,    4] 메세지코드   StartPos 146, Length 4
    char    outgu[1];    //[string,    1] 메세지출력구분   StartPos 150, Length 1
    char    compreq[1];    //[string,    1] 압축요청구분   StartPos 151, Length 1
    char    funckey[4];    //[string,    4] 기능키   StartPos 152, Length 4
    char    reqcnt[4];    //[long  ,    4] 요청레코드개수   StartPos 156, Length 4
    char    filler[6];    //[string,    6] 예비영역   StartPos 160, Length 6
    char    cont[1];    //[string,    1] 연속구분   StartPos 166, Length 1
    char    contkey[18];    //[string,   18] 연속키값   StartPos 167, Length 18
    char    varlen[2];    //[long  ,    2] 가변시스템길이   StartPos 185, Length 2
    char    varhdlen[2];    //[long  ,    2] 가변해더길이   StartPos 187, Length 2
    char    varmsglen[2];    //[long  ,    2] 가변메시지길이   StartPos 189, Length 2
    char    trsrc[1];    //[string,    1] 조회발원지   StartPos 191, Length 1
    char    eventid[4];    //[string,    4] I/F이벤트ID   StartPos 192, Length 4
    char    ifinfo[4];    //[string,    4] I/F정보   StartPos 196, Length 4
    char    filler1[41];    //[string,   41] 예비영역   StartPos 200, Length 41
    char    ordxctptncode[2];    //[string,    2] 주문체결유형코드   StartPos 241, Length 2
    char    ordmktcode[2];    //[string,    2] 주문시장코드   StartPos 243, Length 2
    char    ordptncode[2];    //[string,    2] 주문유형코드   StartPos 245, Length 2
    char    mgmtbrnno[3];    //[string,    3] 관리지점번호   StartPos 247, Length 3
    char    accno1[11];    //[string,   11] 계좌번호   StartPos 250, Length 11
    char    accno2[9];    //[string,    9] 계좌번호   StartPos 261, Length 9
    char    acntnm[40];    //[string,   40] 계좌명   StartPos 270, Length 40
    char    Isuno[12];    //[string,   12] 종목번호   StartPos 310, Length 12
    char    Isunm[40];    //[string,   40] 종목명   StartPos 322, Length 40
    char    ordno[10];    //[long  ,   10] 주문번호   StartPos 362, Length 10
    char    orgordno[10];    //[long  ,   10] 원주문번호   StartPos 372, Length 10
    char    execno[10];    //[long  ,   10] 체결번호   StartPos 382, Length 10
    char    ordqty[16];    //[long  ,   16] 주문수량   StartPos 392, Length 16
    char    ordprc[13];    //[long  ,   13] 주문가격   StartPos 408, Length 13
    char    execqty[16];    //[long  ,   16] 체결수량   StartPos 421, Length 16
    char    execprc[13];    //[long  ,   13] 체결가격   StartPos 437, Length 13
    char    mdfycnfqty[16];    //[long  ,   16] 정정확인수량   StartPos 450, Length 16
    char    mdfycnfprc[16];    //[long  ,   16] 정정확인가격   StartPos 466, Length 16
    char    canccnfqty[16];    //[long  ,   16] 취소확인수량   StartPos 482, Length 16
    char    rjtqty[16];    //[long  ,   16] 거부수량   StartPos 498, Length 16
    char    ordtrxptncode[4];    //[long  ,    4] 주문처리유형코드   StartPos 514, Length 4
    char    mtiordseqno[10];    //[long  ,   10] 복수주문일련번호   StartPos 518, Length 10
    char    ordcndi[1];    //[string,    1] 주문조건   StartPos 528, Length 1
    char    ordprcptncode[2];    //[string,    2] 호가유형코드   StartPos 529, Length 2
    char    nsavtrdqty[16];    //[long  ,   16] 비저축체결수량   StartPos 531, Length 16
    char    shtnIsuno[9];    //[string,    9] 단축종목번호   StartPos 547, Length 9
    char    opdrtnno[12];    //[string,   12] 운용지시번호   StartPos 556, Length 12
    char    cvrgordtp[1];    //[string,    1] 반대매매주문구분   StartPos 568, Length 1
    char    unercqty[16];    //[long  ,   16] 미체결수량(주문)   StartPos 569, Length 16
    char    orgordunercqty[16];    //[long  ,   16] 원주문미체결수량   StartPos 585, Length 16
    char    orgordmdfyqty[16];    //[long  ,   16] 원주문정정수량   StartPos 601, Length 16
    char    orgordcancqty[16];    //[long  ,   16] 원주문취소수량   StartPos 617, Length 16
    char    ordavrexecprc[13];    //[long  ,   13] 주문평균체결가격   StartPos 633, Length 13
    char    ordamt[16];    //[long  ,   16] 주문금액   StartPos 646, Length 16
    char    stdIsuno[12];    //[string,   12] 표준종목번호   StartPos 662, Length 12
    char    bfstdIsuno[12];    //[string,   12] 전표준종목번호   StartPos 674, Length 12
    char    bnstp[1];    //[string,    1] 매매구분   StartPos 686, Length 1
    char    ordtrdptncode[2];    //[string,    2] 주문거래유형코드   StartPos 687, Length 2
    char    mgntrncode[3];    //[string,    3] 신용거래코드   StartPos 689, Length 3
    char    adduptp[2];    //[string,    2] 수수료합산코드   StartPos 692, Length 2
    char    commdacode[2];    //[string,    2] 통신매체코드   StartPos 694, Length 2
    char    Loandt[8];    //[string,    8] 대출일   StartPos 696, Length 8
    char    mbrnmbrno[3];    //[long  ,    3] 회원/비회원사번호   StartPos 704, Length 3
    char    ordacntno[20];    //[string,   20] 주문계좌번호   StartPos 707, Length 20
    char    agrgbrnno[3];    //[string,    3] 집계지점번호   StartPos 727, Length 3
    char    mgempno[9];    //[string,    9] 관리사원번호   StartPos 730, Length 9
    char    futsLnkbrnno[3];    //[string,    3] 선물연계지점번호   StartPos 739, Length 3
    char    futsLnkacntno[20];    //[string,   20] 선물연계계좌번호   StartPos 742, Length 20
    char    futsmkttp[1];    //[string,    1] 선물시장구분   StartPos 762, Length 1
    char    regmktcode[2];    //[string,    2] 등록시장코드   StartPos 763, Length 2
    char    mnymgnrat[7];    //[long  ,    7] 현금증거금률   StartPos 765, Length 7
    char    substmgnrat[9];    //[long  ,    9] 대용증거금률   StartPos 772, Length 9
    char    mnyexecamt[16];    //[long  ,   16] 현금체결금액   StartPos 781, Length 16
    char    ubstexecamt[16];    //[long  ,   16] 대용체결금액   StartPos 797, Length 16
    char    cmsnamtexecamt[16];    //[long  ,   16] 수수료체결금액   StartPos 813, Length 16
    char    crdtpldgexecamt[16];    //[long  ,   16] 신용담보체결금액   StartPos 829, Length 16
    char    crdtexecamt[16];    //[long  ,   16] 신용체결금액   StartPos 845, Length 16
    char    prdayruseexecval[16];    //[long  ,   16] 전일재사용체결금액   StartPos 861, Length 16
    char    crdayruseexecval[16];    //[long  ,   16] 금일재사용체결금액   StartPos 877, Length 16
    char    spotexecqty[16];    //[long  ,   16] 실물체결수량   StartPos 893, Length 16
    char    stslexecqty[16];    //[long  ,   16] 공매도체결수량   StartPos 909, Length 16
    char    strtgcode[6];    //[string,    6] 전략코드   StartPos 925, Length 6
    char    grpId[20];    //[string,   20] 그룹Id   StartPos 931, Length 20
    char    ordseqno[10];    //[long  ,   10] 주문회차   StartPos 951, Length 10
    char    ptflno[10];    //[long  ,   10] 포트폴리오번호   StartPos 961, Length 10
    char    bskno[10];    //[long  ,   10] 바스켓번호   StartPos 971, Length 10
    char    trchno[10];    //[long  ,   10] 트렌치번호   StartPos 981, Length 10
    char    itemno[10];    //[long  ,   10] 아이템번호   StartPos 991, Length 10
    char    orduserId[16];    //[string,   16] 주문자Id   StartPos 1001, Length 16
    char    brwmgmtYn[1];    //[long  ,    1] 차입관리여부   StartPos 1017, Length 1
    char    frgrunqno[6];    //[string,    6] 외국인고유번호   StartPos 1018, Length 6
    char    trtzxLevytp[1];    //[string,    1] 거래세징수구분   StartPos 1024, Length 1
    char    lptp[1];    //[string,    1] 유동성공급자구분   StartPos 1025, Length 1
    char    exectime[9];    //[string,    9] 체결시각   StartPos 1026, Length 9
    char    rcptexectime[9];    //[string,    9] 거래소수신체결시각   StartPos 1035, Length 9
    char    dummy_rmndLoanamt[16];    //[long  ,   16] 잔여대출금액   StartPos 1044, Length 16
    char    dummy_secbalqty[16];    //[long  ,   16] 잔고수량   StartPos 1060, Length 16
    char    dummy_spotordableqty[16];    //[long  ,   16] 실물가능수량   StartPos 1076, Length 16
    char    dummy_ordableruseqty[16];    //[long  ,   16] 재사용가능수량(매도)   StartPos 1092, Length 16
    char    flctqty[16];    //[long  ,   16] 변동수량   StartPos 1108, Length 16
    char    dummy_secbalqtyd2[16];    //[long  ,   16] 잔고수량(d2)   StartPos 1124, Length 16
    char    dummy_sellableqty[16];    //[long  ,   16] 매도주문가능수량   StartPos 1140, Length 16
    char    dummy_unercsellordqty[16];    //[long  ,   16] 미체결매도주문수량   StartPos 1156, Length 16
    char    dummy_avrpchsprc[13];    //[long  ,   13] 평균매입가   StartPos 1172, Length 13
    char    dummy_pchsant[16];    //[long  ,   16] 매입금액   StartPos 1185, Length 16
    char    deposit[16];    //[long  ,   16] 예수금   StartPos 1201, Length 16
    char    substamt[16];    //[long  ,   16] 대용금   StartPos 1217, Length 16
    char    csgnmnymgn[16];    //[long  ,   16] 위탁증거금현금   StartPos 1233, Length 16
    char    csgnsubstmgn[16];    //[long  ,   16] 위탁증거금대용   StartPos 1249, Length 16
    char    crdtpldgruseamt[16];    //[long  ,   16] 신용담보재사용금   StartPos 1265, Length 16
    char    ordablemny[16];    //[long  ,   16] 주문가능현금   StartPos 1281, Length 16
    char    ordablesubstamt[16];    //[long  ,   16] 주문가능대용   StartPos 1297, Length 16
    char    ruseableamt[16];    //[long  ,   16] 재사용가능금액   StartPos 1313, Length 16
} SC3_OutBlock;

//------------------------------------------------------------------------------
// 주식 주문 거부 실시간 정보 (SC4)
//------------------------------------------------------------------------------
typedef struct {
    char    lineseq[10];    //[long  ,   10] 라인일련번호   StartPos 0, Length 10
    char    accno[11];    //[string,   11] 계좌번호   StartPos 10, Length 11
    char    user[8];    //[string,    8] 조작자ID   StartPos 21, Length 8
    char    len[6];    //[long  ,    6] 헤더길이   StartPos 29, Length 6
    char    gubun[1];    //[string,    1] 헤더구분   StartPos 35, Length 1
    char    compress[1];    //[string,    1] 압축구분   StartPos 36, Length 1
    char    encrypt[1];    //[string,    1] 암호구분   StartPos 37, Length 1
    char    offset[3];    //[long  ,    3] 공통시작지점   StartPos 38, Length 3
    char    trcode[8];    //[string,    8] TRCODE   StartPos 41, Length 8
    char    compid[3];    //[string,    3] 이용사번호   StartPos 49, Length 3
    char    userid[16];    //[string,   16] 사용자ID   StartPos 52, Length 16
    char    media[2];    //[string,    2] 접속매체   StartPos 68, Length 2
    char    ifid[3];    //[string,    3] I/F일련번호   StartPos 70, Length 3
    char    seq[9];    //[string,    9] 전문일련번호   StartPos 73, Length 9
    char    trid[16];    //[string,   16] TR추적ID   StartPos 82, Length 16
    char    pubip[12];    //[string,   12] 공인IP   StartPos 98, Length 12
    char    prvip[12];    //[string,   12] 사설IP   StartPos 110, Length 12
    char    pcbpno[3];    //[string,    3] 처리지점번호   StartPos 122, Length 3
    char    bpno[3];    //[string,    3] 지점번호   StartPos 125, Length 3
    char    termno[8];    //[string,    8] 단말번호   StartPos 128, Length 8
    char    lang[1];    //[string,    1] 언어구분   StartPos 136, Length 1
    char    proctm[9];    //[long  ,    9] AP처리시간   StartPos 137, Length 9
    char    msgcode[4];    //[string,    4] 메세지코드   StartPos 146, Length 4
    char    outgu[1];    //[string,    1] 메세지출력구분   StartPos 150, Length 1
    char    compreq[1];    //[string,    1] 압축요청구분   StartPos 151, Length 1
    char    funckey[4];    //[string,    4] 기능키   StartPos 152, Length 4
    char    reqcnt[4];    //[long  ,    4] 요청레코드개수   StartPos 156, Length 4
    char    filler[6];    //[string,    6] 예비영역   StartPos 160, Length 6
    char    cont[1];    //[string,    1] 연속구분   StartPos 166, Length 1
    char    contkey[18];    //[string,   18] 연속키값   StartPos 167, Length 18
    char    varlen[2];    //[long  ,    2] 가변시스템길이   StartPos 185, Length 2
    char    varhdlen[2];    //[long  ,    2] 가변해더길이   StartPos 187, Length 2
    char    varmsglen[2];    //[long  ,    2] 가변메시지길이   StartPos 189, Length 2
    char    trsrc[1];    //[string,    1] 조회발원지   StartPos 191, Length 1
    char    eventid[4];    //[string,    4] I/F이벤트ID   StartPos 192, Length 4
    char    ifinfo[4];    //[string,    4] I/F정보   StartPos 196, Length 4
    char    filler1[41];    //[string,   41] 예비영역   StartPos 200, Length 41
    char    ordxctptncode[2];    //[string,    2] 주문체결유형코드   StartPos 241, Length 2
    char    ordmktcode[2];    //[string,    2] 주문시장코드   StartPos 243, Length 2
    char    ordptncode[2];    //[string,    2] 주문유형코드   StartPos 245, Length 2
    char    mgmtbrnno[3];    //[string,    3] 관리지점번호   StartPos 247, Length 3
    char    accno1[11];    //[string,   11] 계좌번호   StartPos 250, Length 11
    char    accno2[9];    //[string,    9] 계좌번호   StartPos 261, Length 9
    char    acntnm[40];    //[string,   40] 계좌명   StartPos 270, Length 40
    char    Isuno[12];    //[string,   12] 종목번호   StartPos 310, Length 12
    char    Isunm[40];    //[string,   40] 종목명   StartPos 322, Length 40
    char    ordno[10];    //[long  ,   10] 주문번호   StartPos 362, Length 10
    char    orgordno[10];    //[long  ,   10] 원주문번호   StartPos 372, Length 10
    char    execno[10];    //[long  ,   10] 체결번호   StartPos 382, Length 10
    char    ordqty[16];    //[long  ,   16] 주문수량   StartPos 392, Length 16
    char    ordprc[13];    //[long  ,   13] 주문가격   StartPos 408, Length 13
    char    execqty[16];    //[long  ,   16] 체결수량   StartPos 421, Length 16
    char    execprc[13];    //[long  ,   13] 체결가격   StartPos 437, Length 13
    char    mdfycnfqty[16];    //[long  ,   16] 정정확인수량   StartPos 450, Length 16
    char    mdfycnfprc[16];    //[long  ,   16] 정정확인가격   StartPos 466, Length 16
    char    canccnfqty[16];    //[long  ,   16] 취소확인수량   StartPos 482, Length 16
    char    rjtqty[16];    //[long  ,   16] 거부수량   StartPos 498, Length 16
    char    ordtrxptncode[4];    //[long  ,    4] 주문처리유형코드   StartPos 514, Length 4
    char    mtiordseqno[10];    //[long  ,   10] 복수주문일련번호   StartPos 518, Length 10
    char    ordcndi[1];    //[string,    1] 주문조건   StartPos 528, Length 1
    char    ordprcptncode[2];    //[string,    2] 호가유형코드   StartPos 529, Length 2
    char    nsavtrdqty[16];    //[long  ,   16] 비저축체결수량   StartPos 531, Length 16
    char    shtnIsuno[9];    //[string,    9] 단축종목번호   StartPos 547, Length 9
    char    opdrtnno[12];    //[string,   12] 운용지시번호   StartPos 556, Length 12
    char    cvrgordtp[1];    //[string,    1] 반대매매주문구분   StartPos 568, Length 1
    char    unercqty[16];    //[long  ,   16] 미체결수량(주문)   StartPos 569, Length 16
    char    orgordunercqty[16];    //[long  ,   16] 원주문미체결수량   StartPos 585, Length 16
    char    orgordmdfyqty[16];    //[long  ,   16] 원주문정정수량   StartPos 601, Length 16
    char    orgordcancqty[16];    //[long  ,   16] 원주문취소수량   StartPos 617, Length 16
    char    ordavrexecprc[13];    //[long  ,   13] 주문평균체결가격   StartPos 633, Length 13
    char    ordamt[16];    //[long  ,   16] 주문금액   StartPos 646, Length 16
    char    stdIsuno[12];    //[string,   12] 표준종목번호   StartPos 662, Length 12
    char    bfstdIsuno[12];    //[string,   12] 전표준종목번호   StartPos 674, Length 12
    char    bnstp[1];    //[string,    1] 매매구분   StartPos 686, Length 1
    char    ordtrdptncode[2];    //[string,    2] 주문거래유형코드   StartPos 687, Length 2
    char    mgntrncode[3];    //[string,    3] 신용거래코드   StartPos 689, Length 3
    char    adduptp[2];    //[string,    2] 수수료합산코드   StartPos 692, Length 2
    char    commdacode[2];    //[string,    2] 통신매체코드   StartPos 694, Length 2
    char    Loandt[8];    //[string,    8] 대출일   StartPos 696, Length 8
    char    mbrnmbrno[3];    //[long  ,    3] 회원/비회원사번호   StartPos 704, Length 3
    char    ordacntno[20];    //[string,   20] 주문계좌번호   StartPos 707, Length 20
    char    agrgbrnno[3];    //[string,    3] 집계지점번호   StartPos 727, Length 3
    char    mgempno[9];    //[string,    9] 관리사원번호   StartPos 730, Length 9
    char    futsLnkbrnno[3];    //[string,    3] 선물연계지점번호   StartPos 739, Length 3
    char    futsLnkacntno[20];    //[string,   20] 선물연계계좌번호   StartPos 742, Length 20
    char    futsmkttp[1];    //[string,    1] 선물시장구분   StartPos 762, Length 1
    char    regmktcode[2];    //[string,    2] 등록시장코드   StartPos 763, Length 2
    char    mnymgnrat[7];    //[long  ,    7] 현금증거금률   StartPos 765, Length 7
    char    substmgnrat[9];    //[long  ,    9] 대용증거금률   StartPos 772, Length 9
    char    mnyexecamt[16];    //[long  ,   16] 현금체결금액   StartPos 781, Length 16
    char    ubstexecamt[16];    //[long  ,   16] 대용체결금액   StartPos 797, Length 16
    char    cmsnamtexecamt[16];    //[long  ,   16] 수수료체결금액   StartPos 813, Length 16
    char    crdtpldgexecamt[16];    //[long  ,   16] 신용담보체결금액   StartPos 829, Length 16
    char    crdtexecamt[16];    //[long  ,   16] 신용체결금액   StartPos 845, Length 16
    char    prdayruseexecval[16];    //[long  ,   16] 전일재사용체결금액   StartPos 861, Length 16
    char    crdayruseexecval[16];    //[long  ,   16] 금일재사용체결금액   StartPos 877, Length 16
    char    spotexecqty[16];    //[long  ,   16] 실물체결수량   StartPos 893, Length 16
    char    stslexecqty[16];    //[long  ,   16] 공매도체결수량   StartPos 909, Length 16
    char    strtgcode[6];    //[string,    6] 전략코드   StartPos 925, Length 6
    char    grpId[20];    //[string,   20] 그룹Id   StartPos 931, Length 20
    char    ordseqno[10];    //[long  ,   10] 주문회차   StartPos 951, Length 10
    char    ptflno[10];    //[long  ,   10] 포트폴리오번호   StartPos 961, Length 10
    char    bskno[10];    //[long  ,   10] 바스켓번호   StartPos 971, Length 10
    char    trchno[10];    //[long  ,   10] 트렌치번호   StartPos 981, Length 10
    char    itemno[10];    //[long  ,   10] 아이템번호   StartPos 991, Length 10
    char    orduserId[16];    //[string,   16] 주문자Id   StartPos 1001, Length 16
    char    brwmgmtYn[1];    //[long  ,    1] 차입관리여부   StartPos 1017, Length 1
    char    frgrunqno[6];    //[string,    6] 외국인고유번호   StartPos 1018, Length 6
    char    trtzxLevytp[1];    //[string,    1] 거래세징수구분   StartPos 1024, Length 1
    char    lptp[1];    //[string,    1] 유동성공급자구분   StartPos 1025, Length 1
    char    exectime[9];    //[string,    9] 체결시각   StartPos 1026, Length 9
    char    rcptexectime[9];    //[string,    9] 거래소수신체결시각   StartPos 1035, Length 9
    char    dummy_rmndLoanamt[16];    //[long  ,   16] 잔여대출금액   StartPos 1044, Length 16
    char    dummy_secbalqty[16];    //[long  ,   16] 잔고수량   StartPos 1060, Length 16
    char    dummy_spotordableqty[16];    //[long  ,   16] 실물가능수량   StartPos 1076, Length 16
    char    dummy_ordableruseqty[16];    //[long  ,   16] 재사용가능수량(매도)   StartPos 1092, Length 16
    char    flctqty[16];    //[long  ,   16] 변동수량   StartPos 1108, Length 16
    char    dummy_secbalqtyd2[16];    //[long  ,   16] 잔고수량(d2)   StartPos 1124, Length 16
    char    dummy_sellableqty[16];    //[long  ,   16] 매도주문가능수량   StartPos 1140, Length 16
    char    dummy_unercsellordqty[16];    //[long  ,   16] 미체결매도주문수량   StartPos 1156, Length 16
    char    dummy_avrpchsprc[13];    //[long  ,   13] 평균매입가   StartPos 1172, Length 13
    char    dummy_pchsant[16];    //[long  ,   16] 매입금액   StartPos 1185, Length 16
    char    deposit[16];    //[long  ,   16] 예수금   StartPos 1201, Length 16
    char    substamt[16];    //[long  ,   16] 대용금   StartPos 1217, Length 16
    char    csgnmnymgn[16];    //[long  ,   16] 위탁증거금현금   StartPos 1233, Length 16
    char    csgnsubstmgn[16];    //[long  ,   16] 위탁증거금대용   StartPos 1249, Length 16
    char    crdtpldgruseamt[16];    //[long  ,   16] 신용담보재사용금   StartPos 1265, Length 16
    char    ordablemny[16];    //[long  ,   16] 주문가능현금   StartPos 1281, Length 16
    char    ordablesubstamt[16];    //[long  ,   16] 주문가능대용   StartPos 1297, Length 16
    char    ruseableamt[16];    //[long  ,   16] 재사용가능금액   StartPos 1313, Length 16
} SC4_OutBlock;


//------------------------------------------------------------------------------
// 주식 현재가 호가 조회 (t1101)
//------------------------------------------------------------------------------
typedef struct {
    char    shcode[6];      char    _shcode;    //[string,    6] 단축코드   StartPos 0, Length 6
} T1101InBlock;

typedef struct {
    char hname[20];         char _hname;        //[string,   20] 한글명   StartPos 0, Length 20
    char price[8];          char _price;        //[long  ,    8] 현재가   StartPos 21, Length 8
    char sign[1];           char _sign;         //[string,    1] 전일대비구분   StartPos 30, Length 1
    char change[8];         char _change;       //[long  ,    8] 전일대비   StartPos 32, Length 8
    char diff[6];           char _diff;         //[float ,  6.2] 등락율   StartPos 41, Length 6
    char volume[12];        char _volume;       //[long  ,   12] 누적거래량   StartPos 48, Length 12
    char jnilclose[8];      char _jnilclose;    //[long  ,    8] 전일종가   StartPos 61, Length 8
    char offerho1[8];       char _offerho1;     //[long  ,    8] 매도호가1   StartPos 70, Length 8
    char bidho1[8];         char _bidho1;       //[long  ,    8] 매수호가1   StartPos 79, Length 8
    char offerrem1[12];     char _offerrem1;    //[long  ,   12] 매도호가수량1   StartPos 88, Length 12
    char bidrem1[12];       char _bidrem1;      //[long  ,   12] 매수호가수량1   StartPos 101, Length 12
    char preoffercha1[12];  char _preoffercha1; //[long  ,   12] 직전매도대비수량1   StartPos 114, Length 12
    char prebidcha1[12];    char _prebidcha1;   //[long  ,   12] 직전매수대비수량1   StartPos 127, Length 12
    char offerho2[8];       char _offerho2;     //[long  ,    8] 매도호가2   StartPos 140, Length 8
    char bidho2[8];         char _bidho2;       //[long  ,    8] 매수호가2   StartPos 149, Length 8
    char offerrem2[12];     char _offerrem2;    //[long  ,   12] 매도호가수량2   StartPos 158, Length 12
    char bidrem2[12];       char _bidrem2;      //[long  ,   12] 매수호가수량2   StartPos 171, Length 12
    char preoffercha2[12];  char _preoffercha2; //[long  ,   12] 직전매도대비수량2   StartPos 184, Length 12
    char prebidcha2[12];    char _prebidcha2;   //[long  ,   12] 직전매수대비수량2   StartPos 197, Length 12
    char offerho3[8];       char _offerho3;     //[long  ,    8] 매도호가3   StartPos 210, Length 8
    char bidho3[8];         char _bidho3;       //[long  ,    8] 매수호가3   StartPos 219, Length 8
    char offerrem3[12];     char _offerrem3;    //[long  ,   12] 매도호가수량3   StartPos 228, Length 12
    char bidrem3[12];       char _bidrem3;      //[long  ,   12] 매수호가수량3   StartPos 241, Length 12
    char preoffercha3[12];  char _preoffercha3; //[long  ,   12] 직전매도대비수량3   StartPos 254, Length 12
    char prebidcha3[12];    char _prebidcha3;   //[long  ,   12] 직전매수대비수량3   StartPos 267, Length 12
    char offerho4[8];       char _offerho4;     //[long  ,    8] 매도호가4   StartPos 280, Length 8
    char bidho4[8];         char _bidho4;       //[long  ,    8] 매수호가4   StartPos 289, Length 8
    char offerrem4[12];     char _offerrem4;    //[long  ,   12] 매도호가수량4   StartPos 298, Length 12
    char bidrem4[12];       char _bidrem4;      //[long  ,   12] 매수호가수량4   StartPos 311, Length 12
    char preoffercha4[12];  char _preoffercha4; //[long  ,   12] 직전매도대비수량4   StartPos 324, Length 12
    char prebidcha4[12];    char _prebidcha4;   //[long  ,   12] 직전매수대비수량4   StartPos 337, Length 12
    char offerho5[8];       char _offerho5;     //[long  ,    8] 매도호가5   StartPos 350, Length 8
    char bidho5[8];         char _bidho5;       //[long  ,    8] 매수호가5   StartPos 359, Length 8
    char offerrem5[12];     char _offerrem5;    //[long  ,   12] 매도호가수량5   StartPos 368, Length 12
    char bidrem5[12];       char _bidrem5;      //[long  ,   12] 매수호가수량5   StartPos 381, Length 12
    char preoffercha5[12];  char _preoffercha5; //[long  ,   12] 직전매도대비수량5   StartPos 394, Length 12
    char prebidcha5[12];    char _prebidcha5;   //[long  ,   12] 직전매수대비수량5   StartPos 407, Length 12
    char offerho6[8];       char _offerho6;     //[long  ,    8] 매도호가6   StartPos 420, Length 8
    char bidho6[8];         char _bidho6;       //[long  ,    8] 매수호가6   StartPos 429, Length 8
    char offerrem6[12];     char _offerrem6;    //[long  ,   12] 매도호가수량6   StartPos 438, Length 12
    char bidrem6[12];       char _bidrem6;      //[long  ,   12] 매수호가수량6   StartPos 451, Length 12
    char preoffercha6[12];  char _preoffercha6; //[long  ,   12] 직전매도대비수량6   StartPos 464, Length 12
    char prebidcha6[12];    char _prebidcha6;   //[long  ,   12] 직전매수대비수량6   StartPos 477, Length 12
    char offerho7[8];       char _offerho7;     //[long  ,    8] 매도호가7   StartPos 490, Length 8
    char bidho7[8];         char _bidho7;       //[long  ,    8] 매수호가7   StartPos 499, Length 8
    char offerrem7[12];     char _offerrem7;    //[long  ,   12] 매도호가수량7   StartPos 508, Length 12
    char bidrem7[12];       char _bidrem7;      //[long  ,   12] 매수호가수량7   StartPos 521, Length 12
    char preoffercha7[12];  char _preoffercha7; //[long  ,   12] 직전매도대비수량7   StartPos 534, Length 12
    char prebidcha7[12];    char _prebidcha7;   //[long  ,   12] 직전매수대비수량7   StartPos 547, Length 12
    char offerho8[8];       char _offerho8;     //[long  ,    8] 매도호가8   StartPos 560, Length 8
    char bidho8[8];         char _bidho8;       //[long  ,    8] 매수호가8   StartPos 569, Length 8
    char offerrem8[12];     char _offerrem8;    //[long  ,   12] 매도호가수량8   StartPos 578, Length 12
    char bidrem8[12];       char _bidrem8;      //[long  ,   12] 매수호가수량8   StartPos 591, Length 12
    char preoffercha8[12];  char _preoffercha8; //[long  ,   12] 직전매도대비수량8   StartPos 604, Length 12
    char prebidcha8[12];    char _prebidcha8;   //[long  ,   12] 직전매수대비수량8   StartPos 617, Length 12
    char offerho9[8];       char _offerho9;     //[long  ,    8] 매도호가9   StartPos 630, Length 8
    char bidho9[8];         char _bidho9;       //[long  ,    8] 매수호가9   StartPos 639, Length 8
    char offerrem9[12];     char _offerrem9;    //[long  ,   12] 매도호가수량9   StartPos 648, Length 12
    char bidrem9[12];       char _bidrem9;      //[long  ,   12] 매수호가수량9   StartPos 661, Length 12
    char preoffercha9[12];  char _preoffercha9; //[long  ,   12] 직전매도대비수량9   StartPos 674, Length 12
    char prebidcha9[12];    char _prebidcha9;   //[long  ,   12] 직전매수대비수량9   StartPos 687, Length 12
    char offerho10[8];      char _offerho10;    //[long  ,    8] 매도호가10   StartPos 700, Length 8
    char bidho10[8];        char _bidho10;      //[long  ,    8] 매수호가10   StartPos 709, Length 8
    char offerrem10[12];    char _offerrem10;   //[long  ,   12] 매도호가수량10   StartPos 718, Length 12
    char bidrem10[12];      char _bidrem10;     //[long  ,   12] 매수호가수량10   StartPos 731, Length 12
    char preoffercha10[12]; char _preoffercha10; //[long  ,   12] 직전매도대비수량10   StartPos 744, Length 12
    char prebidcha10[12];   char _prebidcha10;  //[long  ,   12] 직전매수대비수량10   StartPos 757, Length 12
    char offer[12];         char _offer;        //[long  ,   12] 매도호가수량합   StartPos 770, Length 12
    char bid[12];           char _bid;          //[long  ,   12] 매수호가수량합   StartPos 783, Length 12
    char preoffercha[12];   char _preoffercha;  //[long  ,   12] 직전매도대비수량합   StartPos 796, Length 12
    char prebidcha[12];     char _prebidcha;    //[long  ,   12] 직전매수대비수량합   StartPos 809, Length 12
    char hotime[8];         char _hotime;       //[string,    8] 수신시간   StartPos 822, Length 8
    char yeprice[8];        char _yeprice;      //[long  ,    8] 예상체결가격   StartPos 831, Length 8
    char yevolume[12];      char _yevolume;     //[long  ,   12] 예상체결수량   StartPos 840, Length 12
    char yesign[1];         char _yesign;       //[string,    1] 예상체결전일구분   StartPos 853, Length 1
    char yechange[8];       char _yechange;     //[long  ,    8] 예상체결전일대비   StartPos 855, Length 8
    char yediff[6];         char _yediff;       //[float ,  6.2] 예상체결등락율   StartPos 864, Length 6
    char tmoffer[12];       char _tmoffer;      //[long  ,   12] 시간외매도잔량   StartPos 871, Length 12
    char tmbid[12];         char _tmbid;        //[long  ,   12] 시간외매수잔량   StartPos 884, Length 12
    char ho_status[1];      char _ho_status;    //[string,    1] 동시구분   StartPos 897, Length 1
    char shcode[6];         char _shcode;       //[string,    6] 단축코드   StartPos 899, Length 6
    char uplmtprice[8];     char _uplmtprice;   //[long  ,    8] 상한가   StartPos 906, Length 8
    char dnlmtprice[8];     char _dnlmtprice;   //[long  ,    8] 하한가   StartPos 915, Length 8
    char open[8];           char _open;         //[long  ,    8] 시가   StartPos 924, Length 8
    char high[8];           char _high;         //[long  ,    8] 고가   StartPos 933, Length 8
    char low[8];            char _low;          //[long  ,    8] 저가   StartPos 942, Length 8
} T1101OutBlock;

//------------------------------------------------------------------------------
// 주식 현재가 시세 조회 (t1102)
//------------------------------------------------------------------------------
typedef struct {
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 0, Length 6
} T1102InBlock;

typedef struct {
    char    hname               [  20];    char    _hname               ;    // [string,   20] 한글명                          StartPos 0, Length 20
    char    price               [   8];    char    _price               ;    // [long  ,    8] 현재가                          StartPos 21, Length 8
    char    sign                [   1];    char    _sign                ;    // [string,    1] 전일대비구분                    StartPos 30, Length 1
    char    change              [   8];    char    _change              ;    // [long  ,    8] 전일대비                        StartPos 32, Length 8
    char    diff                [   6];    char    _diff                ;    // [float ,  6.2] 등락율                          StartPos 41, Length 6
    char    volume              [  12];    char    _volume              ;    // [long  ,   12] 누적거래량                      StartPos 48, Length 12
    char    recprice            [   8];    char    _recprice            ;    // [long  ,    8] 기준가(평가가격)                StartPos 61, Length 8
    char    avg                 [   8];    char    _avg                 ;    // [long  ,    8] 가중평균                        StartPos 70, Length 8
    char    uplmtprice          [   8];    char    _uplmtprice          ;    // [long  ,    8] 상한가(최고호가가격)            StartPos 79, Length 8
    char    dnlmtprice          [   8];    char    _dnlmtprice          ;    // [long  ,    8] 하한가(최저호가가격)            StartPos 88, Length 8
    char    jnilvolume          [  12];    char    _jnilvolume          ;    // [long  ,   12] 전일거래량                      StartPos 97, Length 12
    char    volumediff          [  12];    char    _volumediff          ;    // [long  ,   12] 거래량차                        StartPos 110, Length 12
    char    open                [   8];    char    _open                ;    // [long  ,    8] 시가                            StartPos 123, Length 8
    char    opentime            [   6];    char    _opentime            ;    // [string,    6] 시가시간                        StartPos 132, Length 6
    char    high                [   8];    char    _high                ;    // [long  ,    8] 고가                            StartPos 139, Length 8
    char    hightime            [   6];    char    _hightime            ;    // [string,    6] 고가시간                        StartPos 148, Length 6
    char    low                 [   8];    char    _low                 ;    // [long  ,    8] 저가                            StartPos 155, Length 8
    char    lowtime             [   6];    char    _lowtime             ;    // [string,    6] 저가시간                        StartPos 164, Length 6
    char    high52w             [   8];    char    _high52w             ;    // [long  ,    8] 52최고가                        StartPos 171, Length 8
    char    high52wdate         [   8];    char    _high52wdate         ;    // [string,    8] 52최고가일                      StartPos 180, Length 8
    char    low52w              [   8];    char    _low52w              ;    // [long  ,    8] 52최저가                        StartPos 189, Length 8
    char    low52wdate          [   8];    char    _low52wdate          ;    // [string,    8] 52최저가일                      StartPos 198, Length 8
    char    exhratio            [   6];    char    _exhratio            ;    // [float ,  6.2] 소진율                          StartPos 207, Length 6
    char    per                 [   6];    char    _per                 ;    // [float ,  6.2] PER                             StartPos 214, Length 6
    char    pbrx                [   6];    char    _pbrx                ;    // [float ,  6.2] PBRX                            StartPos 221, Length 6
    char    listing             [  12];    char    _listing             ;    // [long  ,   12] 상장주식수(천)                  StartPos 228, Length 12
    char    jkrate              [   8];    char    _jkrate              ;    // [long  ,    8] 증거금율                        StartPos 241, Length 8
    char    memedan             [   5];    char    _memedan             ;    // [string,    5] 수량단위                        StartPos 250, Length 5
    char    offernocd1          [   3];    char    _offernocd1          ;    // [string,    3] 매도증권사코드1                 StartPos 256, Length 3
    char    bidnocd1            [   3];    char    _bidnocd1            ;    // [string,    3] 매수증권사코드1                 StartPos 260, Length 3
    char    offerno1            [   6];    char    _offerno1            ;    // [string,    6] 매도증권사명1                   StartPos 264, Length 6
    char    bidno1              [   6];    char    _bidno1              ;    // [string,    6] 매수증권사명1                   StartPos 271, Length 6
    char    dvol1               [   8];    char    _dvol1               ;    // [long  ,    8] 총매도수량1                     StartPos 278, Length 8
    char    svol1               [   8];    char    _svol1               ;    // [long  ,    8] 총매수수량1                     StartPos 287, Length 8
    char    dcha1               [   8];    char    _dcha1               ;    // [long  ,    8] 매도증감1                       StartPos 296, Length 8
    char    scha1               [   8];    char    _scha1               ;    // [long  ,    8] 매수증감1                       StartPos 305, Length 8
    char    ddiff1              [   6];    char    _ddiff1              ;    // [float ,  6.2] 매도비율1                       StartPos 314, Length 6
    char    sdiff1              [   6];    char    _sdiff1              ;    // [float ,  6.2] 매수비율1                       StartPos 321, Length 6
    char    offernocd2          [   3];    char    _offernocd2          ;    // [string,    3] 매도증권사코드2                 StartPos 328, Length 3
    char    bidnocd2            [   3];    char    _bidnocd2            ;    // [string,    3] 매수증권사코드2                 StartPos 332, Length 3
    char    offerno2            [   6];    char    _offerno2            ;    // [string,    6] 매도증권사명2                   StartPos 336, Length 6
    char    bidno2              [   6];    char    _bidno2              ;    // [string,    6] 매수증권사명2                   StartPos 343, Length 6
    char    dvol2               [   8];    char    _dvol2               ;    // [long  ,    8] 총매도수량2                     StartPos 350, Length 8
    char    svol2               [   8];    char    _svol2               ;    // [long  ,    8] 총매수수량2                     StartPos 359, Length 8
    char    dcha2               [   8];    char    _dcha2               ;    // [long  ,    8] 매도증감2                       StartPos 368, Length 8
    char    scha2               [   8];    char    _scha2               ;    // [long  ,    8] 매수증감2                       StartPos 377, Length 8
    char    ddiff2              [   6];    char    _ddiff2              ;    // [float ,  6.2] 매도비율2                       StartPos 386, Length 6
    char    sdiff2              [   6];    char    _sdiff2              ;    // [float ,  6.2] 매수비율2                       StartPos 393, Length 6
    char    offernocd3          [   3];    char    _offernocd3          ;    // [string,    3] 매도증권사코드3                 StartPos 400, Length 3
    char    bidnocd3            [   3];    char    _bidnocd3            ;    // [string,    3] 매수증권사코드3                 StartPos 404, Length 3
    char    offerno3            [   6];    char    _offerno3            ;    // [string,    6] 매도증권사명3                   StartPos 408, Length 6
    char    bidno3              [   6];    char    _bidno3              ;    // [string,    6] 매수증권사명3                   StartPos 415, Length 6
    char    dvol3               [   8];    char    _dvol3               ;    // [long  ,    8] 총매도수량3                     StartPos 422, Length 8
    char    svol3               [   8];    char    _svol3               ;    // [long  ,    8] 총매수수량3                     StartPos 431, Length 8
    char    dcha3               [   8];    char    _dcha3               ;    // [long  ,    8] 매도증감3                       StartPos 440, Length 8
    char    scha3               [   8];    char    _scha3               ;    // [long  ,    8] 매수증감3                       StartPos 449, Length 8
    char    ddiff3              [   6];    char    _ddiff3              ;    // [float ,  6.2] 매도비율3                       StartPos 458, Length 6
    char    sdiff3              [   6];    char    _sdiff3              ;    // [float ,  6.2] 매수비율3                       StartPos 465, Length 6
    char    offernocd4          [   3];    char    _offernocd4          ;    // [string,    3] 매도증권사코드4                 StartPos 472, Length 3
    char    bidnocd4            [   3];    char    _bidnocd4            ;    // [string,    3] 매수증권사코드4                 StartPos 476, Length 3
    char    offerno4            [   6];    char    _offerno4            ;    // [string,    6] 매도증권사명4                   StartPos 480, Length 6
    char    bidno4              [   6];    char    _bidno4              ;    // [string,    6] 매수증권사명4                   StartPos 487, Length 6
    char    dvol4               [   8];    char    _dvol4               ;    // [long  ,    8] 총매도수량4                     StartPos 494, Length 8
    char    svol4               [   8];    char    _svol4               ;    // [long  ,    8] 총매수수량4                     StartPos 503, Length 8
    char    dcha4               [   8];    char    _dcha4               ;    // [long  ,    8] 매도증감4                       StartPos 512, Length 8
    char    scha4               [   8];    char    _scha4               ;    // [long  ,    8] 매수증감4                       StartPos 521, Length 8
    char    ddiff4              [   6];    char    _ddiff4              ;    // [float ,  6.2] 매도비율4                       StartPos 530, Length 6
    char    sdiff4              [   6];    char    _sdiff4              ;    // [float ,  6.2] 매수비율4                       StartPos 537, Length 6
    char    offernocd5          [   3];    char    _offernocd5          ;    // [string,    3] 매도증권사코드5                 StartPos 544, Length 3
    char    bidnocd5            [   3];    char    _bidnocd5            ;    // [string,    3] 매수증권사코드5                 StartPos 548, Length 3
    char    offerno5            [   6];    char    _offerno5            ;    // [string,    6] 매도증권사명5                   StartPos 552, Length 6
    char    bidno5              [   6];    char    _bidno5              ;    // [string,    6] 매수증권사명5                   StartPos 559, Length 6
    char    dvol5               [   8];    char    _dvol5               ;    // [long  ,    8] 총매도수량5                     StartPos 566, Length 8
    char    svol5               [   8];    char    _svol5               ;    // [long  ,    8] 총매수수량5                     StartPos 575, Length 8
    char    dcha5               [   8];    char    _dcha5               ;    // [long  ,    8] 매도증감5                       StartPos 584, Length 8
    char    scha5               [   8];    char    _scha5               ;    // [long  ,    8] 매수증감5                       StartPos 593, Length 8
    char    ddiff5              [   6];    char    _ddiff5              ;    // [float ,  6.2] 매도비율5                       StartPos 602, Length 6
    char    sdiff5              [   6];    char    _sdiff5              ;    // [float ,  6.2] 매수비율5                       StartPos 609, Length 6
    char    fwdvl               [  12];    char    _fwdvl               ;    // [long  ,   12] 외국계매도합계수량              StartPos 616, Length 12
    char    ftradmdcha          [  12];    char    _ftradmdcha          ;    // [long  ,   12] 외국계매도직전대비              StartPos 629, Length 12
    char    ftradmddiff         [   6];    char    _ftradmddiff         ;    // [float ,  6.2] 외국계매도비율                  StartPos 642, Length 6
    char    fwsvl               [  12];    char    _fwsvl               ;    // [long  ,   12] 외국계매수합계수량              StartPos 649, Length 12
    char    ftradmscha          [  12];    char    _ftradmscha          ;    // [long  ,   12] 외국계매수직전대비              StartPos 662, Length 12
    char    ftradmsdiff         [   6];    char    _ftradmsdiff         ;    // [float ,  6.2] 외국계매수비율                  StartPos 675, Length 6
    char    vol                 [   6];    char    _vol                 ;    // [float ,  6.2] 회전율                          StartPos 682, Length 6
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 689, Length 6
    char    value               [  12];    char    _value               ;    // [long  ,   12] 누적거래대금                    StartPos 696, Length 12
    char    jvolume             [  12];    char    _jvolume             ;    // [long  ,   12] 전일동시간거래량                StartPos 709, Length 12
    char    highyear            [   8];    char    _highyear            ;    // [long  ,    8] 연중최고가                      StartPos 722, Length 8
    char    highyeardate        [   8];    char    _highyeardate        ;    // [string,    8] 연중최고일자                    StartPos 731, Length 8
    char    lowyear             [   8];    char    _lowyear             ;    // [long  ,    8] 연중최저가                      StartPos 740, Length 8
    char    lowyeardate         [   8];    char    _lowyeardate         ;    // [string,    8] 연중최저일자                    StartPos 749, Length 8
    char    target              [   8];    char    _target              ;    // [long  ,    8] 목표가                          StartPos 758, Length 8
    char    capital             [  12];    char    _capital             ;    // [long  ,   12] 자본금                          StartPos 767, Length 12
    char    abscnt              [  12];    char    _abscnt              ;    // [long  ,   12] 유동주식수                      StartPos 780, Length 12
    char    parprice            [   8];    char    _parprice            ;    // [long  ,    8] 액면가                          StartPos 793, Length 8
    char    gsmm                [   2];    char    _gsmm                ;    // [string,    2] 결산월                          StartPos 802, Length 2
    char    subprice            [   8];    char    _subprice            ;    // [long  ,    8] 대용가                          StartPos 805, Length 8
    char    total               [  12];    char    _total               ;    // [long  ,   12] 시가총액                        StartPos 814, Length 12
    char    listdate            [   8];    char    _listdate            ;    // [string,    8] 상장일                          StartPos 827, Length 8
    char    name                [  10];    char    _name                ;    // [string,   10] 전분기명                        StartPos 836, Length 10
    char    bfsales             [  12];    char    _bfsales             ;    // [long  ,   12] 전분기매출액                    StartPos 847, Length 12
    char    bfoperatingincome   [  12];    char    _bfoperatingincome   ;    // [long  ,   12] 전분기영업이익                  StartPos 860, Length 12
    char    bfordinaryincome    [  12];    char    _bfordinaryincome    ;    // [long  ,   12] 전분기경상이익                  StartPos 873, Length 12
    char    bfnetincome         [  12];    char    _bfnetincome         ;    // [long  ,   12] 전분기순이익                    StartPos 886, Length 12
    char    bfeps               [  13];    char    _bfeps               ;    // [float , 13.2] 전분기EPS                       StartPos 899, Length 13
    char    name2               [  10];    char    _name2               ;    // [string,   10] 전전분기명                      StartPos 913, Length 10
    char    bfsales2            [  12];    char    _bfsales2            ;    // [long  ,   12] 전전분기매출액                  StartPos 924, Length 12
    char    bfoperatingincome2  [  12];    char    _bfoperatingincome2  ;    // [long  ,   12] 전전분기영업이익                StartPos 937, Length 12
    char    bfordinaryincome2   [  12];    char    _bfordinaryincome2   ;    // [long  ,   12] 전전분기경상이익                StartPos 950, Length 12
    char    bfnetincome2        [  12];    char    _bfnetincome2        ;    // [long  ,   12] 전전분기순이익                  StartPos 963, Length 12
    char    bfeps2              [  13];    char    _bfeps2              ;    // [float , 13.2] 전전분기EPS                     StartPos 976, Length 13
    char    salert              [   7];    char    _salert              ;    // [float ,  7.2] 전년대비매출액                  StartPos 990, Length 7
    char    opert               [   7];    char    _opert               ;    // [float ,  7.2] 전년대비영업이익                StartPos 998, Length 7
    char    ordrt               [   7];    char    _ordrt               ;    // [float ,  7.2] 전년대비경상이익                StartPos 1006, Length 7
    char    netrt               [   7];    char    _netrt               ;    // [float ,  7.2] 전년대비순이익                  StartPos 1014, Length 7
    char    epsrt               [   7];    char    _epsrt               ;    // [float ,  7.2] 전년대비EPS                     StartPos 1022, Length 7
    char    info1               [  10];    char    _info1               ;    // [string,   10] 락구분                          StartPos 1030, Length 10
    char    info2               [  10];    char    _info2               ;    // [string,   10] 관리/급등구분                   StartPos 1041, Length 10
    char    info3               [  10];    char    _info3               ;    // [string,   10] 정지/연장구분                   StartPos 1052, Length 10
    char    info4               [  12];    char    _info4               ;    // [string,   12] 투자/불성실구분                 StartPos 1063, Length 12
    char    janginfo            [  10];    char    _janginfo            ;    // [string,   10] 장구분                          StartPos 1076, Length 10
    char    t_per               [   6];    char    _t_per               ;    // [float ,  6.2] T.PER                           StartPos 1087, Length 6
    char    tonghwa             [   3];    char    _tonghwa             ;    // [string,    3] 통화ISO코드                     StartPos 1094, Length 3
    char    dval1               [  18];    char    _dval1               ;    // [long  ,   18] 총매도대금1                     StartPos 1098, Length 18
    char    sval1               [  18];    char    _sval1               ;    // [long  ,   18] 총매수대금1                     StartPos 1117, Length 18
    char    dval2               [  18];    char    _dval2               ;    // [long  ,   18] 총매도대금2                     StartPos 1136, Length 18
    char    sval2               [  18];    char    _sval2               ;    // [long  ,   18] 총매수대금2                     StartPos 1155, Length 18
    char    dval3               [  18];    char    _dval3               ;    // [long  ,   18] 총매도대금3                     StartPos 1174, Length 18
    char    sval3               [  18];    char    _sval3               ;    // [long  ,   18] 총매수대금3                     StartPos 1193, Length 18
    char    dval4               [  18];    char    _dval4               ;    // [long  ,   18] 총매도대금4                     StartPos 1212, Length 18
    char    sval4               [  18];    char    _sval4               ;    // [long  ,   18] 총매수대금4                     StartPos 1231, Length 18
    char    dval5               [  18];    char    _dval5               ;    // [long  ,   18] 총매도대금5                     StartPos 1250, Length 18
    char    sval5               [  18];    char    _sval5               ;    // [long  ,   18] 총매수대금5                     StartPos 1269, Length 18
    char    davg1               [   8];    char    _davg1               ;    // [long  ,    8] 총매도평단가1                   StartPos 1288, Length 8
    char    savg1               [   8];    char    _savg1               ;    // [long  ,    8] 총매수평단가1                   StartPos 1297, Length 8
    char    davg2               [   8];    char    _davg2               ;    // [long  ,    8] 총매도평단가2                   StartPos 1306, Length 8
    char    savg2               [   8];    char    _savg2               ;    // [long  ,    8] 총매수평단가2                   StartPos 1315, Length 8
    char    davg3               [   8];    char    _davg3               ;    // [long  ,    8] 총매도평단가3                   StartPos 1324, Length 8
    char    savg3               [   8];    char    _savg3               ;    // [long  ,    8] 총매수평단가3                   StartPos 1333, Length 8
    char    davg4               [   8];    char    _davg4               ;    // [long  ,    8] 총매도평단가4                   StartPos 1342, Length 8
    char    savg4               [   8];    char    _savg4               ;    // [long  ,    8] 총매수평단가4                   StartPos 1351, Length 8
    char    davg5               [   8];    char    _davg5               ;    // [long  ,    8] 총매도평단가5                   StartPos 1360, Length 8
    char    savg5               [   8];    char    _savg5               ;    // [long  ,    8] 총매수평단가5                   StartPos 1369, Length 8
    char    ftradmdval          [  18];    char    _ftradmdval          ;    // [long  ,   18] 외국계매도대금                  StartPos 1378, Length 18
    char    ftradmsval          [  18];    char    _ftradmsval          ;    // [long  ,   18] 외국계매수대금                  StartPos 1397, Length 18
    char    ftradmdavg          [   8];    char    _ftradmdavg          ;    // [long  ,    8] 외국계매도평단가                StartPos 1416, Length 8
    char    ftradmsavg          [   8];    char    _ftradmsavg          ;    // [long  ,    8] 외국계매수평단가                StartPos 1425, Length 8
    char    info5               [   8];    char    _info5               ;    // [string,    8] 투자주의환기                    StartPos 1434, Length 8
    char    spac_gubun          [   1];    char    _spac_gubun          ;    // [string,    1] 기업인수목적회사여부            StartPos 1443, Length 1
    char    issueprice          [   8];    char    _issueprice          ;    // [long  ,    8] 발행가격                        StartPos 1445, Length 8
    char    alloc_gubun         [   1];    char    _alloc_gubun         ;    // [string,    1] 배분적용구분코드(1:배분발생2:배 StartPos 1454, Length 1
    char    alloc_text          [   8];    char    _alloc_text          ;    // [string,    8] 배분적용구분                    StartPos 1456, Length 8
    char    shterm_text         [  10];    char    _shterm_text         ;    // [string,   10] 단기과열/VI발동                 StartPos 1465, Length 10
    char    svi_uplmtprice      [   8];    char    _svi_uplmtprice      ;    // [long  ,    8] 정적VI상한가                    StartPos 1476, Length 8
    char    svi_dnlmtprice      [   8];    char    _svi_dnlmtprice      ;    // [long  ,    8] 정적VI하한가                    StartPos 1485, Length 8
    char    low_lqdt_gu         [   1];    char    _low_lqdt_gu         ;    // [string,    1] 저유동성종목여부                StartPos 1494, Length 1
    char    abnormal_rise_gu    [   1];    char    _abnormal_rise_gu    ;    // [string,    1] 이상급등종목여부                StartPos 1496, Length 1
    char    lend_text           [   8];    char    _lend_text           ;    // [string,    8] 대차불가표시                    StartPos 1498, Length 8
} T1102OutBlock;

//------------------------------------------------------------------------------
// 기간별 주가 (t1305)
//------------------------------------------------------------------------------
typedef struct {
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 0, Length 6
    char    dwmcode             [   1];    char    _dwmcode             ;    // [long  ,    1] 일주월구분                      StartPos 7, Length 1
    char    date                [   8];    char    _date                ;    // [string,    8] 날짜                            StartPos 9, Length 8
    char    idx                 [   4];    char    _idx                 ;    // [long  ,    4] IDX                             StartPos 18, Length 4
    char    cnt                 [   4];    char    _cnt                 ;    // [long  ,    4] 건수                            StartPos 23, Length 4
} T1305InBlock;

typedef struct {
    char    cnt                 [   4];    char    _cnt                 ;    // [long  ,    4] CNT                             StartPos 0, Length 4
    char    date                [   8];    char    _date                ;    // [string,    8] 날짜                            StartPos 5, Length 8
    char    idx                 [   4];    char    _idx                 ;    // [long  ,    4] IDX                             StartPos 14, Length 4
} T1305OutBlock;

typedef struct {
    char    date                [   8];    char    _date                ;    // [string,    8] 날짜                            StartPos 0, Length 8
    char    open                [   8];    char    _open                ;    // [long  ,    8] 시가                            StartPos 9, Length 8
    char    high                [   8];    char    _high                ;    // [long  ,    8] 고가                            StartPos 18, Length 8
    char    low                 [   8];    char    _low                 ;    // [long  ,    8] 저가                            StartPos 27, Length 8
    char    close               [   8];    char    _close               ;    // [long  ,    8] 종가                            StartPos 36, Length 8
    char    sign                [   1];    char    _sign                ;    // [string,    1] 전일대비구분                    StartPos 45, Length 1
    char    change              [   8];    char    _change              ;    // [long  ,    8] 전일대비                        StartPos 47, Length 8
    char    diff                [   6];    char    _diff                ;    // [float ,  6.2] 등락율                          StartPos 56, Length 6
    char    volume              [  12];    char    _volume              ;    // [long  ,   12] 누적거래량                      StartPos 63, Length 12
    char    diff_vol            [  10];    char    _diff_vol            ;    // [float , 10.2] 거래증가율                      StartPos 76, Length 10
    char    chdegree            [   6];    char    _chdegree            ;    // [float ,  6.2] 체결강도                        StartPos 87, Length 6
    char    sojinrate           [   6];    char    _sojinrate           ;    // [float ,  6.2] 소진율                          StartPos 94, Length 6
    char    changerate          [   6];    char    _changerate          ;    // [float ,  6.2] 회전율                          StartPos 101, Length 6
    char    fpvolume            [  12];    char    _fpvolume            ;    // [long  ,   12] 외인순매수                      StartPos 108, Length 12
    char    covolume            [  12];    char    _covolume            ;    // [long  ,   12] 기관순매수                      StartPos 121, Length 12
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 종목코드                        StartPos 134, Length 6
    char    value               [  12];    char    _value               ;    // [long  ,   12] 누적거래대금(단위:백만)         StartPos 141, Length 12
    char    ppvolume            [  12];    char    _ppvolume            ;    // [long  ,   12] 개인순매수                      StartPos 154, Length 12
    char    o_sign              [   1];    char    _o_sign              ;    // [string,    1] 시가대비구분                    StartPos 167, Length 1
    char    o_change            [   8];    char    _o_change            ;    // [long  ,    8] 시가대비                        StartPos 169, Length 8
    char    o_diff              [   6];    char    _o_diff              ;    // [float ,  6.2] 시가기준등락율                  StartPos 178, Length 6
    char    h_sign              [   1];    char    _h_sign              ;    // [string,    1] 고가대비구분                    StartPos 185, Length 1
    char    h_change            [   8];    char    _h_change            ;    // [long  ,    8] 고가대비                        StartPos 187, Length 8
    char    h_diff              [   6];    char    _h_diff              ;    // [float ,  6.2] 고가기준등락율                  StartPos 196, Length 6
    char    l_sign              [   1];    char    _l_sign              ;    // [string,    1] 저가대비구분                    StartPos 203, Length 1
    char    l_change            [   8];    char    _l_change            ;    // [long  ,    8] 저가대비                        StartPos 205, Length 8
    char    l_diff              [   6];    char    _l_diff              ;    // [float ,  6.2] 저가기준등락율                  StartPos 214, Length 6
    char    marketcap           [  12];    char    _marketcap           ;    // [long  ,   12] 시가총액(단위:백만)             StartPos 221, Length 12
} T1305OutBlock1;

//------------------------------------------------------------------------------
// 현물 당일전일분틱조회 (t1310)
//------------------------------------------------------------------------------
typedef struct {
    char    daygb               [   1];    char    _daygb               ;    // [string,    1] 당일전일구분                    StartPos 0, Length 1
    char    timegb              [   1];    char    _timegb              ;    // [string,    1] 분틱구분                        StartPos 2, Length 1
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 4, Length 6
    char    endtime             [   4];    char    _endtime             ;    // [string,    4] 종료시간                        StartPos 11, Length 4
    char    cts_time            [  10];    char    _cts_time            ;    // [string,   10] 시간CTS                         StartPos 16, Length 10
} T1310InBlock;

typedef struct {
    char    cts_time            [  10];    char    _cts_time            ;    // [string,   10] 시간CTS                         StartPos 0, Length 10
} T1310OutBlock;

typedef struct {
    char    chetime             [  10];    char    _chetime             ;    // [string,   10] 시간                            StartPos 0, Length 10
    char    price               [   8];    char    _price               ;    // [long  ,    8] 현재가                          StartPos 11, Length 8
    char    sign                [   1];    char    _sign                ;    // [string,    1] 전일대비구분                    StartPos 20, Length 1
    char    change              [   8];    char    _change              ;    // [long  ,    8] 전일대비                        StartPos 22, Length 8
    char    diff                [   6];    char    _diff                ;    // [float ,  6.2] 등락율                          StartPos 31, Length 6
    char    cvolume             [  12];    char    _cvolume             ;    // [long  ,   12] 체결수량                        StartPos 38, Length 12
    char    chdegree            [   8];    char    _chdegree            ;    // [float ,  8.2] 체결강도                        StartPos 51, Length 8
    char    volume              [  12];    char    _volume              ;    // [long  ,   12] 거래량                          StartPos 60, Length 12
    char    mdvolume            [  12];    char    _mdvolume            ;    // [long  ,   12] 매도체결수량                    StartPos 73, Length 12
    char    mdchecnt            [   8];    char    _mdchecnt            ;    // [long  ,    8] 매도체결건수                    StartPos 86, Length 8
    char    msvolume            [  12];    char    _msvolume            ;    // [long  ,   12] 매수체결수량                    StartPos 95, Length 12
    char    mschecnt            [   8];    char    _mschecnt            ;    // [long  ,    8] 매수체결건수                    StartPos 108, Length 8
    char    revolume            [  12];    char    _revolume            ;    // [long  ,   12] 순체결량                        StartPos 117, Length 12
    char    rechecnt            [   8];    char    _rechecnt            ;    // [long  ,    8] 순체결건수                      StartPos 130, Length 8
} T1310OutBlock1;

//------------------------------------------------------------------------------
// 시간 조회 (t0167)
//------------------------------------------------------------------------------
typedef struct {
    char    date[8];                            // 일자(YYYYMMDD)
    char    time[12];                           // 시간(HHMMSSssssss)
} T0167OutBlock;

//------------------------------------------------------------------------------
// ETF 현재가(시세) 조회 (t1901)
//------------------------------------------------------------------------------
typedef struct {
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 0, Length 6
} T1901InBlock;

typedef struct {
    char    hname               [  20];    char    _hname               ;    // [string,   20] 한글명                          StartPos 0, Length 20
    char    price               [   8];    char    _price               ;    // [long  ,    8] 현재가                          StartPos 21, Length 8
    char    sign                [   1];    char    _sign                ;    // [string,    1] 전일대비구분                    StartPos 30, Length 1
    char    change              [   8];    char    _change              ;    // [long  ,    8] 전일대비                        StartPos 32, Length 8
    char    diff                [   6];    char    _diff                ;    // [float ,  6.2] 등락율                          StartPos 41, Length 6
    char    volume              [  12];    char    _volume              ;    // [float ,   12] 누적거래량                      StartPos 48, Length 12
    char    recprice            [   8];    char    _recprice            ;    // [long  ,    8] 기준가                          StartPos 61, Length 8
    char    avg                 [   8];    char    _avg                 ;    // [long  ,    8] 가중평균                        StartPos 70, Length 8
    char    uplmtprice          [   8];    char    _uplmtprice          ;    // [long  ,    8] 상한가                          StartPos 79, Length 8
    char    dnlmtprice          [   8];    char    _dnlmtprice          ;    // [long  ,    8] 하한가                          StartPos 88, Length 8
    char    jnilvolume          [  12];    char    _jnilvolume          ;    // [float ,   12] 전일거래량                      StartPos 97, Length 12
    char    volumediff          [  12];    char    _volumediff          ;    // [long  ,   12] 거래량차                        StartPos 110, Length 12
    char    open                [   8];    char    _open                ;    // [long  ,    8] 시가                            StartPos 123, Length 8
    char    opentime            [   6];    char    _opentime            ;    // [string,    6] 시가시간                        StartPos 132, Length 6
    char    high                [   8];    char    _high                ;    // [long  ,    8] 고가                            StartPos 139, Length 8
    char    hightime            [   6];    char    _hightime            ;    // [string,    6] 고가시간                        StartPos 148, Length 6
    char    low                 [   8];    char    _low                 ;    // [long  ,    8] 저가                            StartPos 155, Length 8
    char    lowtime             [   6];    char    _lowtime             ;    // [string,    6] 저가시간                        StartPos 164, Length 6
    char    high52w             [   8];    char    _high52w             ;    // [long  ,    8] 52최고가                        StartPos 171, Length 8
    char    high52wdate         [   8];    char    _high52wdate         ;    // [string,    8] 52최고가일                      StartPos 180, Length 8
    char    low52w              [   8];    char    _low52w              ;    // [long  ,    8] 52최저가                        StartPos 189, Length 8
    char    low52wdate          [   8];    char    _low52wdate          ;    // [string,    8] 52최저가일                      StartPos 198, Length 8
    char    exhratio            [   6];    char    _exhratio            ;    // [float ,  6.2] 소진율                          StartPos 207, Length 6
    char    flmtvol             [  12];    char    _flmtvol             ;    // [float ,   12] 외국인보유수량                  StartPos 214, Length 12
    char    per                 [   6];    char    _per                 ;    // [float ,  6.2] PER                             StartPos 227, Length 6
    char    listing             [  12];    char    _listing             ;    // [long  ,   12] 상장주식수(천)                  StartPos 234, Length 12
    char    jkrate              [   8];    char    _jkrate              ;    // [long  ,    8] 증거금율                        StartPos 247, Length 8
    char    vol                 [   6];    char    _vol                 ;    // [float ,  6.2] 회전율                          StartPos 256, Length 6
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 263, Length 6
    char    value               [  12];    char    _value               ;    // [long  ,   12] 누적거래대금                    StartPos 270, Length 12
    char    highyear            [   8];    char    _highyear            ;    // [long  ,    8] 연중최고가                      StartPos 283, Length 8
    char    highyeardate        [   8];    char    _highyeardate        ;    // [string,    8] 연중최고일자                    StartPos 292, Length 8
    char    lowyear             [   8];    char    _lowyear             ;    // [long  ,    8] 연중최저가                      StartPos 301, Length 8
    char    lowyeardate         [   8];    char    _lowyeardate         ;    // [string,    8] 연중최저일자                    StartPos 310, Length 8
    char    upname              [  20];    char    _upname              ;    // [string,   20] 업종명                          StartPos 319, Length 20
    char    upcode              [   3];    char    _upcode              ;    // [string,    3] 업종코드                        StartPos 340, Length 3
    char    upprice             [   7];    char    _upprice             ;    // [float ,  7.2] 업종현재가                      StartPos 344, Length 7
    char    upsign              [   1];    char    _upsign              ;    // [string,    1] 업종전일비구분                  StartPos 352, Length 1
    char    upchange            [   6];    char    _upchange            ;    // [float ,  6.2] 업종전일대비                    StartPos 354, Length 6
    char    updiff              [   6];    char    _updiff              ;    // [float ,  6.2] 업종등락율                      StartPos 361, Length 6
    char    futname             [  20];    char    _futname             ;    // [string,   20] 선물최근월물명                  StartPos 368, Length 20
    char    futcode             [   8];    char    _futcode             ;    // [string,    8] 선물최근월물코드                StartPos 389, Length 8
    char    futprice            [   6];    char    _futprice            ;    // [float ,  6.2] 선물현재가                      StartPos 398, Length 6
    char    futsign             [   1];    char    _futsign             ;    // [string,    1] 선물전일비구분                  StartPos 405, Length 1
    char    futchange           [   6];    char    _futchange           ;    // [float ,  6.2] 선물전일대비                    StartPos 407, Length 6
    char    futdiff             [   6];    char    _futdiff             ;    // [float ,  6.2] 선물등락율                      StartPos 414, Length 6
    char    nav                 [   8];    char    _nav                 ;    // [float ,  8.2] NAV                             StartPos 421, Length 8
    char    navsign             [   1];    char    _navsign             ;    // [string,    1] NAV전일대비구분                 StartPos 430, Length 1
    char    navchange           [   8];    char    _navchange           ;    // [float ,  8.2] NAV전일대비                     StartPos 432, Length 8
    char    navdiff             [   6];    char    _navdiff             ;    // [float ,  6.2] NAV등락율                       StartPos 441, Length 6
    char    cocrate             [   6];    char    _cocrate             ;    // [float ,  6.2] 추적오차율                      StartPos 448, Length 6
    char    kasis               [   6];    char    _kasis               ;    // [float ,  6.2] 괴리율                          StartPos 455, Length 6
    char    subprice            [  10];    char    _subprice            ;    // [long  ,   10] 대용가                          StartPos 462, Length 10
    char    offerno1            [   6];    char    _offerno1            ;    // [string,    6] 매도증권사코드1                 StartPos 473, Length 6
    char    bidno1              [   6];    char    _bidno1              ;    // [string,    6] 매수증권사코드1                 StartPos 480, Length 6
    char    dvol1               [   8];    char    _dvol1               ;    // [long  ,    8] 총매도수량1                     StartPos 487, Length 8
    char    svol1               [   8];    char    _svol1               ;    // [long  ,    8] 총매수수량1                     StartPos 496, Length 8
    char    dcha1               [   8];    char    _dcha1               ;    // [long  ,    8] 매도증감1                       StartPos 505, Length 8
    char    scha1               [   8];    char    _scha1               ;    // [long  ,    8] 매수증감1                       StartPos 514, Length 8
    char    ddiff1              [   6];    char    _ddiff1              ;    // [float ,  6.2] 매도비율1                       StartPos 523, Length 6
    char    sdiff1              [   6];    char    _sdiff1              ;    // [float ,  6.2] 매수비율1                       StartPos 530, Length 6
    char    offerno2            [   6];    char    _offerno2            ;    // [string,    6] 매도증권사코드2                 StartPos 537, Length 6
    char    bidno2              [   6];    char    _bidno2              ;    // [string,    6] 매수증권사코드2                 StartPos 544, Length 6
    char    dvol2               [   8];    char    _dvol2               ;    // [long  ,    8] 총매도수량2                     StartPos 551, Length 8
    char    svol2               [   8];    char    _svol2               ;    // [long  ,    8] 총매수수량2                     StartPos 560, Length 8
    char    dcha2               [   8];    char    _dcha2               ;    // [long  ,    8] 매도증감2                       StartPos 569, Length 8
    char    scha2               [   8];    char    _scha2               ;    // [long  ,    8] 매수증감2                       StartPos 578, Length 8
    char    ddiff2              [   6];    char    _ddiff2              ;    // [float ,  6.2] 매도비율2                       StartPos 587, Length 6
    char    sdiff2              [   6];    char    _sdiff2              ;    // [float ,  6.2] 매수비율2                       StartPos 594, Length 6
    char    offerno3            [   6];    char    _offerno3            ;    // [string,    6] 매도증권사코드3                 StartPos 601, Length 6
    char    bidno3              [   6];    char    _bidno3              ;    // [string,    6] 매수증권사코드3                 StartPos 608, Length 6
    char    dvol3               [   8];    char    _dvol3               ;    // [long  ,    8] 총매도수량3                     StartPos 615, Length 8
    char    svol3               [   8];    char    _svol3               ;    // [long  ,    8] 총매수수량3                     StartPos 624, Length 8
    char    dcha3               [   8];    char    _dcha3               ;    // [long  ,    8] 매도증감3                       StartPos 633, Length 8
    char    scha3               [   8];    char    _scha3               ;    // [long  ,    8] 매수증감3                       StartPos 642, Length 8
    char    ddiff3              [   6];    char    _ddiff3              ;    // [float ,  6.2] 매도비율3                       StartPos 651, Length 6
    char    sdiff3              [   6];    char    _sdiff3              ;    // [float ,  6.2] 매수비율3                       StartPos 658, Length 6
    char    offerno4            [   6];    char    _offerno4            ;    // [string,    6] 매도증권사코드4                 StartPos 665, Length 6
    char    bidno4              [   6];    char    _bidno4              ;    // [string,    6] 매수증권사코드4                 StartPos 672, Length 6
    char    dvol4               [   8];    char    _dvol4               ;    // [long  ,    8] 총매도수량4                     StartPos 679, Length 8
    char    svol4               [   8];    char    _svol4               ;    // [long  ,    8] 총매수수량4                     StartPos 688, Length 8
    char    dcha4               [   8];    char    _dcha4               ;    // [long  ,    8] 매도증감4                       StartPos 697, Length 8
    char    scha4               [   8];    char    _scha4               ;    // [long  ,    8] 매수증감4                       StartPos 706, Length 8
    char    ddiff4              [   6];    char    _ddiff4              ;    // [float ,  6.2] 매도비율4                       StartPos 715, Length 6
    char    sdiff4              [   6];    char    _sdiff4              ;    // [float ,  6.2] 매수비율4                       StartPos 722, Length 6
    char    offerno5            [   6];    char    _offerno5            ;    // [string,    6] 매도증권사코드5                 StartPos 729, Length 6
    char    bidno5              [   6];    char    _bidno5              ;    // [string,    6] 매수증권사코드5                 StartPos 736, Length 6
    char    dvol5               [   8];    char    _dvol5               ;    // [long  ,    8] 총매도수량5                     StartPos 743, Length 8
    char    svol5               [   8];    char    _svol5               ;    // [long  ,    8] 총매수수량5                     StartPos 752, Length 8
    char    dcha5               [   8];    char    _dcha5               ;    // [long  ,    8] 매도증감5                       StartPos 761, Length 8
    char    scha5               [   8];    char    _scha5               ;    // [long  ,    8] 매수증감5                       StartPos 770, Length 8
    char    ddiff5              [   6];    char    _ddiff5              ;    // [float ,  6.2] 매도비율5                       StartPos 779, Length 6
    char    sdiff5              [   6];    char    _sdiff5              ;    // [float ,  6.2] 매수비율5                       StartPos 786, Length 6
    char    fwdvl               [  12];    char    _fwdvl               ;    // [long  ,   12] 외국계매도합계수량              StartPos 793, Length 12
    char    ftradmdcha          [  12];    char    _ftradmdcha          ;    // [long  ,   12] 외국계매도직전대비              StartPos 806, Length 12
    char    ftradmddiff         [   6];    char    _ftradmddiff         ;    // [float ,  6.2] 외국계매도비율                  StartPos 819, Length 6
    char    fwsvl               [  12];    char    _fwsvl               ;    // [long  ,   12] 외국계매수합계수량              StartPos 826, Length 12
    char    ftradmscha          [  12];    char    _ftradmscha          ;    // [long  ,   12] 외국계매수직전대비              StartPos 839, Length 12
    char    ftradmsdiff         [   6];    char    _ftradmsdiff         ;    // [float ,  6.2] 외국계매수비율                  StartPos 852, Length 6
    char    upname2             [  20];    char    _upname2             ;    // [string,   20] 참고지수명                      StartPos 859, Length 20
    char    upcode2             [   3];    char    _upcode2             ;    // [string,    3] 참고지수코드                    StartPos 880, Length 3
    char    upprice2            [   7];    char    _upprice2            ;    // [float ,  7.2] 참고지수현재가                  StartPos 884, Length 7
    char    jnilnav             [   8];    char    _jnilnav             ;    // [float ,  8.2] 전일NAV                         StartPos 892, Length 8
    char    jnilnavsign         [   1];    char    _jnilnavsign         ;    // [string,    1] 전일NAV전일대비구분             StartPos 901, Length 1
    char    jnilnavchange       [   8];    char    _jnilnavchange       ;    // [float ,  8.2] 전일NAV전일대비                 StartPos 903, Length 8
    char    jnilnavdiff         [   6];    char    _jnilnavdiff         ;    // [float ,  6.2] 전일NAV등락율                   StartPos 912, Length 6
    char    etftotcap           [  12];    char    _etftotcap           ;    // [long  ,   12] 순자산총액(억원)                StartPos 919, Length 12
    char    spread              [   6];    char    _spread              ;    // [float ,  6.2] 스프레드                        StartPos 932, Length 6
    char    leverage            [   2];    char    _leverage            ;    // [long  ,    2] 레버리지                        StartPos 939, Length 2
    char    taxgubun            [   1];    char    _taxgubun            ;    // [string,    1] 과세구분                        StartPos 942, Length 1
    char    opcom_nmk           [  20];    char    _opcom_nmk           ;    // [string,   20] 운용사                          StartPos 944, Length 20
    char    lp_nm1              [  20];    char    _lp_nm1              ;    // [string,   20] LP1                             StartPos 965, Length 20
    char    lp_nm2              [  20];    char    _lp_nm2              ;    // [string,   20] LP2                             StartPos 986, Length 20
    char    lp_nm3              [  20];    char    _lp_nm3              ;    // [string,   20] LP3                             StartPos 1007, Length 20
    char    lp_nm4              [  20];    char    _lp_nm4              ;    // [string,   20] LP4                             StartPos 1028, Length 20
    char    lp_nm5              [  20];    char    _lp_nm5              ;    // [string,   20] LP5                             StartPos 1049, Length 20
    char    etf_cp              [  10];    char    _etf_cp              ;    // [string,   10] 복제방법                        StartPos 1070, Length 10
    char    etf_kind            [  10];    char    _etf_kind            ;    // [string,   10] 상품유형                        StartPos 1081, Length 10
    char    vi_gubun            [  10];    char    _vi_gubun            ;    // [string,   10] VI발동해제                      StartPos 1092, Length 10
    char    etn_kind_cd         [  20];    char    _etn_kind_cd         ;    // [string,   20] ETN상품분류                     StartPos 1103, Length 20
    char    lastymd             [   8];    char    _lastymd             ;    // [string,    8] ETN만기일                       StartPos 1124, Length 8
    char    payday              [   8];    char    _payday              ;    // [string,    8] ETN지급일                       StartPos 1133, Length 8
    char    lastdate            [   8];    char    _lastdate            ;    // [string,    8] ETN최종거래일                   StartPos 1142, Length 8
    char    issuernmk           [  20];    char    _issuernmk           ;    // [string,   20] ETN발행시장참가자               StartPos 1151, Length 20
    char    last_sdate          [   8];    char    _last_sdate          ;    // [string,    8] ETN만기상환가격결정시작일       StartPos 1172, Length 8
    char    last_edate          [   8];    char    _last_edate          ;    // [string,    8] ETN만기상환가격결정종료일       StartPos 1181, Length 8
    char    lp_holdvol          [  12];    char    _lp_holdvol          ;    // [string,   12] ETNLP보유수량                   StartPos 1190, Length 12
} T1901OutBlock;

//------------------------------------------------------------------------------
// ETF 시간별 추이 (t1902)
//------------------------------------------------------------------------------
typedef struct {
    char    shCode[6];  char _shcode;       //[string,    6] 단축코드   StartPos 0, Length 6
    char    time[6];    char _time;         //[string,    6] 시간   StartPos 7, Length 6
} T1902InBlock;

typedef struct {
    char    time[6];    char _time;         //[string,    6] 시간   StartPos 0, Length 6
    char    hName[20];  char _hname;        //[string,   20] 종목명   StartPos 7, Length 20
    char    upName[20]; char _upname;       //[string,   20] 업종지수명   StartPos 28, Length 20
} T1902OutBlock;

typedef struct {    // occurs
    char    time[8];    char _time;         //[string,    8] 시간   StartPos 0, Length 8
    char    price[8];   char _price;        //[long  ,    8] 현재가   StartPos 9, Length 8
    char    sign[1];    char _sign;         //[string,    1] 전일대비구분   StartPos 18, Length 1
    char    change[8];  char _change;       //[long  ,    8] 전일대비   StartPos 20, Length 8
    char    volume[12]; char _volume;       //[float ,   12] 누적거래량   StartPos 29, Length 12
    char    navDiff[9]; char _navdiff;      //[float ,  9.2] NAV대비   StartPos 42, Length 9
    char    nav[9];     char _nav;          //[float ,  9.2] NAV   StartPos 52, Length 9
    char    navChange[9];   char _navchange;    //[float ,  9.2] 전일대비   StartPos 62, Length 9
    char    crate[9];   char _crate;        //[float ,  9.2] 추적오차   StartPos 72, Length 9
    char    grate[9];   char _grate;        //[float ,  9.2] 괴리   StartPos 82, Length 9
    char    jisu[8];    char _jisu;         //[float ,  8.2] 지수   StartPos 92, Length 8
    char    jiChange[8];    char _jichange; //[float ,  8.2] 전일대비   StartPos 101, Length 8
    char    jiRate[8];  char _jirate;       //[float ,  8.2] 전일대비율   StartPos 110, Length 8
} T1902OutBlock1;

//------------------------------------------------------------------------------
// 증시 주변 자금 추이 (t8428)
//------------------------------------------------------------------------------
typedef struct {
    char    fdate               [   8];    char    _fdate               ;    // [string,    8] from일자                        StartPos 0, Length 8
    char    tdate               [   8];    char    _tdate               ;    // [string,    8] to일자                          StartPos 9, Length 8
    char    gubun               [   1];    char    _gubun               ;    // [string,    1] 구분                            StartPos 18, Length 1
    char    keyDate            [   8];    char    _key_date            ;    // [string,    8] 날짜                            StartPos 20, Length 8
    char    upcode              [   3];    char    _upcode              ;    // [string,    3] 업종코드                        StartPos 29, Length 3
    char    cnt                 [   3];    char    _cnt                 ;    // [string,    3] 조회건수                        StartPos 33, Length 3
} T8428InBlock;

typedef struct {
    char    date                [   8];    char    _date                ;    // [string,    8] 날짜CTS                         StartPos 0, Length 8
    char    idx                 [   4];    char    _idx                 ;    // [long  ,    4] IDX                             StartPos 9, Length 4
} T8428OutBlock;

typedef struct {
    char    date                [   8];    char    _date                ;    // [string,    8] 일자                            StartPos 0, Length 8
    char    jisu                [   7];    char    _jisu                ;    // [float ,  7.2] 지수                            StartPos 9, Length 7
    char    sign                [   1];    char    _sign                ;    // [string,    1] 대비구분                        StartPos 17, Length 1
    char    change              [   6];    char    _change              ;    // [float ,  6.2] 대비                            StartPos 19, Length 6
    char    diff                [   6];    char    _diff                ;    // [float ,  6.2] 등락율                          StartPos 26, Length 6
    char    volume              [  12];    char    _volume              ;    // [long  ,   12] 거래량                          StartPos 33, Length 12
    char    custmoney           [  12];    char    _custmoney           ;    // [long  ,   12] 고객예탁금_억원                 StartPos 46, Length 12
    char    yecha               [  12];    char    _yecha               ;    // [long  ,   12] 예탁증감_억원                   StartPos 59, Length 12
    char    vol                 [   6];    char    _vol                 ;    // [float ,  6.2] 회전율                          StartPos 72, Length 6
    char    outmoney            [  12];    char    _outmoney            ;    // [long  ,   12] 미수금_억원                     StartPos 79, Length 12
    char    trjango             [  12];    char    _trjango             ;    // [long  ,   12] 신용잔고_억원                   StartPos 92, Length 12
    char    futymoney           [  12];    char    _futymoney           ;    // [long  ,   12] 선물예수금_억원                 StartPos 105, Length 12
    char    stkmoney            [   8];    char    _stkmoney            ;    // [long  ,    8] 주식형_억원                     StartPos 118, Length 8
    char    mstkmoney           [   8];    char    _mstkmoney           ;    // [long  ,    8] 혼합형_억원(주식)               StartPos 127, Length 8
    char    mbndmoney           [   8];    char    _mbndmoney           ;    // [long  ,    8] 혼합형_억원(채권)               StartPos 136, Length 8
    char    bndmoney            [   8];    char    _bndmoney            ;    // [long  ,    8] 채권형_억원                     StartPos 145, Length 8
    char    bndsmoney           [   8];    char    _bndsmoney           ;    // [long  ,    8] 필러(구.단기채권)               StartPos 154, Length 8
    char    mmfmoney            [   8];    char    _mmfmoney            ;    // [long  ,    8] MMF_억원(주식)                  StartPos 163, Length 8
} T8428OutBlock1;

//------------------------------------------------------------------------------
// 주식종목조회 API용 (t8436)
//------------------------------------------------------------------------------
typedef struct {
    char    gubun[1];    //[string,    1] 구분(0:전체1:코스피2:코스닥)
} T8436InBlock;

typedef struct {
    char    hName[20];    //[string,   20] 종목명
    char    shCode[6];    //[string,    6] 단축코드
    char    expCode[12];    //[string,   12] 확장코드
    char    etfGubun[1];    //[string,    1] ETF구분(1:ETF2:ETN)
    char    upLmtPrice[8];    //[long  ,    8] 상한가
    char    dnLmtPrice[8];    //[long  ,    8] 하한가
    char    jnilClose[8];    //[long  ,    8] 전일가
    char    meMeDan[5];    //[string,    5] 주문수량단위
    char    recPrice[8];    //[long  ,    8] 기준가
    char    gubun[1];    //[string,    1] 구분(1:코스피2:코스닥)
    char    bu12Gubun[2];    //[string,    2] 증권그룹
    char    spacGubun[1];    //[string,    1] 기업인수목적회사여부(Y/N)
    char    filler[32];    //[string,   32] filler(미사용)
} T8436OutBlock;

//------------------------------------------------------------------------------
// 코스피 호가 잔량 (H1_)
//------------------------------------------------------------------------------
typedef struct {
    char shcode[6];     char _shcode;    // [string, 6] 단축코드 StartPos 0, Length 6
} H1_InBlock;

typedef struct {
    char hotime[6];     char _hotime;      //[string, 6] 호가시간 StartPos 0, Length 6
    char offerho1[7];   char _offerho1;    //[long , 7] 매도호가1 StartPos 7, Length 7
    char bidho1[7];     char _bidho1;      //[long , 7] 매수호가1 StartPos 15, Length 7
    char offerrem1[9];  char _offerrem1;   //[long , 9] 매도호가잔량1 StartPos 23, Length 9
    char bidrem1[9];    char _bidrem1;     //[long , 9] 매수호가잔량1 StartPos 33, Length 9
    char offerho2[7];   char _offerho2;    //[long , 7] 매도호가2 StartPos 43, Length 7
    char bidho2[7];     char _bidho2;      //[long , 7] 매수호가2 StartPos 51, Length 7
    char offerrem2[9];  char _offerrem2;   //[long , 9] 매도호가잔량2 StartPos 59, Length 9
    char bidrem2[9];    char _bidrem2;     //[long , 9] 매수호가잔량2 StartPos 69, Length 9
    char offerho3[7];   char _offerho3;    //[long , 7] 매도호가3 StartPos 79, Length 7
    char bidho3[7];     char _bidho3;      //[long , 7] 매수호가3 StartPos 87, Length 7
    char offerrem3[9];  char _offerrem3;   //[long , 9] 매도호가잔량3 StartPos 95, Length 9
    char bidrem3[9];    char _bidrem3;     //[long , 9] 매수호가잔량3 StartPos 105, Length 9
    char offerho4[7];   char _offerho4;    //[long , 7] 매도호가4 StartPos 115, Length 7
    char bidho4[7];     char _bidho4;      //[long , 7] 매수호가4 StartPos 123, Length 7
    char offerrem4[9];  char _offerrem4;   //[long , 9] 매도호가잔량4 StartPos 131, Length 9
    char bidrem4[9];    char _bidrem4;     //[long , 9] 매수호가잔량4 StartPos 141, Length 9
    char offerho5[7];   char _offerho5;    //[long , 7] 매도호가5 StartPos 151, Length 7
    char bidho5[7];     char _bidho5;      //[long , 7] 매수호가5 StartPos 159, Length 7
    char offerrem5[9];  char _offerrem5;   //[long , 9] 매도호가잔량5 StartPos 167, Length 9
    char bidrem5[9];    char _bidrem5;     //[long , 9] 매수호가잔량5 StartPos 177, Length 9
    char offerho6[7];   char _offerho6;    //[long , 7] 매도호가6 StartPos 187, Length 7
    char bidho6[7];     char _bidho6;      //[long , 7] 매수호가6 StartPos 195, Length 7
    char offerrem6[9];  char _offerrem6;   //[long , 9] 매도호가잔량6 StartPos 203, Length 9
    char bidrem6[9];    char _bidrem6;     //[long , 9] 매수호가잔량6 StartPos 213, Length 9
    char offerho7[7];   char _offerho7;    //[long , 7] 매도호가7 StartPos 223, Length 7
    char bidho7[7];     char _bidho7;      //[long , 7] 매수호가7 StartPos 231, Length 7
    char offerrem7[9];  char _offerrem7;   //[long , 9] 매도호가잔량7 StartPos 239, Length 9
    char bidrem7[9];    char _bidrem7;     //[long , 9] 매수호가잔량7 StartPos 249, Length 9
    char offerho8[7];   char _offerho8;    //[long , 7] 매도호가8 StartPos 259, Length 7
    char bidho8[7];     char _bidho8;      //[long , 7] 매수호가8 StartPos 267, Length 7
    char offerrem8[9];  char _offerrem8;   //[long , 9] 매도호가잔량8 StartPos 275, Length 9
    char bidrem8[9];    char _bidrem8;     //[long , 9] 매수호가잔량8 StartPos 285, Length 9
    char offerho9[7];   char _offerho9;    //[long , 7] 매도호가9 StartPos 295, Length 7
    char bidho9[7];     char _bidho9;      //[long , 7] 매수호가9 StartPos 303, Length 7
    char offerrem9[9];  char _offerrem9;   //[long , 9] 매도호가잔량9 StartPos 311, Length 9
    char bidrem9[9];    char _bidrem9;     //[long , 9] 매수호가잔량9 StartPos 321, Length 9
    char offerho10[7];  char _offerho10;   //[long , 7] 매도호가10 StartPos 331, Length 7
    char bidho10[7];    char _bidho10;     //[long , 7] 매수호가10 StartPos 339, Length 7
    char offerrem10[9]; char _offerrem10;  //[long , 9] 매도호가잔량10 StartPos 347, Length 9
    char bidrem10[9];   char _bidrem10;    //[long , 9] 매수호가잔량10 StartPos 357, Length 9
    char totofferrem[9]; char _totofferrem; //[long , 9] 총매도호가잔량 StartPos 367, Length 9
    char totbidrem[9];  char _totbidrem;   //[long , 9] 총매수호가잔량 StartPos 377, Length 9
    char donsigubun[1]; char _donsigubun;  //[string, 1] 동시호가구분 StartPos 387, Length 1
    char shcode[6];     char _shcode;      //[string, 6] 단축코드 StartPos 389, Length 6
    char alloc_gubun[1]; char _alloc_gubun; //[string, 1] 배분적용구분 StartPos 396, Length 1
} H1_OutBlock;

//------------------------------------------------------------------------------
// 코스피 시간외 호가 잔량 (H2_)
//------------------------------------------------------------------------------
typedef struct {
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 0, Length 6
} H2_InBlock;

typedef struct {
    char    hotime              [   6];    char    _hotime              ;    // [string,    6] 호가시간                        StartPos 0, Length 6
    char    tmofferrem          [  12];    char    _tmofferrem          ;    // [long  ,   12] 시간외매도잔량                  StartPos 7, Length 12
    char    tmbidrem            [  12];    char    _tmbidrem            ;    // [long  ,   12] 시간외매수잔량                  StartPos 20, Length 12
    char    pretmoffercha       [  12];    char    _pretmoffercha       ;    // [long  ,   12] 시간외매도수량직전대비          StartPos 33, Length 12
    char    pretmbidcha         [  12];    char    _pretmbidcha         ;    // [long  ,   12] 시간외매수수량직전대비          StartPos 46, Length 12
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 59, Length 6
} H2_OutBlock;

//------------------------------------------------------------------------------
// 코스피 체결 (S3_)
//------------------------------------------------------------------------------
typedef struct {
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 0, Length 6
} S3_InBlock;

typedef struct {
    char    chetime             [   6];    char    _chetime             ;    // [string,    6] 체결시간                        StartPos 0, Length 6
    char    sign                [   1];    char    _sign                ;    // [string,    1] 전일대비구분                    StartPos 7, Length 1
    char    change              [   8];    char    _change              ;    // [long  ,    8] 전일대비                        StartPos 9, Length 8
    char    drate               [   6];    char    _drate               ;    // [float ,  6.2] 등락율                          StartPos 18, Length 6
    char    price               [   8];    char    _price               ;    // [long  ,    8] 현재가                          StartPos 25, Length 8
    char    opentime            [   6];    char    _opentime            ;    // [string,    6] 시가시간                        StartPos 34, Length 6
    char    open                [   8];    char    _open                ;    // [long  ,    8] 시가                            StartPos 41, Length 8
    char    hightime            [   6];    char    _hightime            ;    // [string,    6] 고가시간                        StartPos 50, Length 6
    char    high                [   8];    char    _high                ;    // [long  ,    8] 고가                            StartPos 57, Length 8
    char    lowtime             [   6];    char    _lowtime             ;    // [string,    6] 저가시간                        StartPos 66, Length 6
    char    low                 [   8];    char    _low                 ;    // [long  ,    8] 저가                            StartPos 73, Length 8
    char    cgubun              [   1];    char    _cgubun              ;    // [string,    1] 체결구분                        StartPos 82, Length 1
    char    cvolume             [   8];    char    _cvolume             ;    // [long  ,    8] 체결량                          StartPos 84, Length 8
    char    volume              [  12];    char    _volume              ;    // [long  ,   12] 누적거래량                      StartPos 93, Length 12
    char    value               [  12];    char    _value               ;    // [long  ,   12] 누적거래대금                    StartPos 106, Length 12
    char    mdvolume            [  12];    char    _mdvolume            ;    // [long  ,   12] 매도누적체결량                  StartPos 119, Length 12
    char    mdchecnt            [   8];    char    _mdchecnt            ;    // [long  ,    8] 매도누적체결건수                StartPos 132, Length 8
    char    msvolume            [  12];    char    _msvolume            ;    // [long  ,   12] 매수누적체결량                  StartPos 141, Length 12
    char    mschecnt            [   8];    char    _mschecnt            ;    // [long  ,    8] 매수누적체결건수                StartPos 154, Length 8
    char    cpower              [   9];    char    _cpower              ;    // [float ,  9.2] 체결강도                        StartPos 163, Length 9
    char    wAvrg              [   8];    char    _w_avrg              ;    // [long  ,    8] 가중평균가                      StartPos 173, Length 8
    char    offerho             [   8];    char    _offerho             ;    // [long  ,    8] 매도호가                        StartPos 182, Length 8
    char    bidho               [   8];    char    _bidho               ;    // [long  ,    8] 매수호가                        StartPos 191, Length 8
    char    status              [   2];    char    _status              ;    // [string,    2] 장정보                          StartPos 200, Length 2
    char    jnilvolume          [  12];    char    _jnilvolume          ;    // [long  ,   12] 전일동시간대거래량              StartPos 203, Length 12
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 216, Length 6
} S3_OutBlock;

//------------------------------------------------------------------------------
// 코스피 예상 체결 (YS3)
//------------------------------------------------------------------------------
typedef struct {
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 0, Length 6
} YS3InBlock;

typedef struct {
    char    hotime              [   6];    char    _hotime              ;    // [string,    6] 호가시간                        StartPos 0, Length 6
    char    yeprice             [   8];    char    _yeprice             ;    // [long  ,    8] 예상체결가격                    StartPos 7, Length 8
    char    yevolume            [  12];    char    _yevolume            ;    // [long  ,   12] 예상체결수량                    StartPos 16, Length 12
    char    jnilysign           [   1];    char    _jnilysign           ;    // [string,    1] 예상체결가전일종가대비구분      StartPos 29, Length 1
    char    preychange          [   8];    char    _preychange          ;    // [long  ,    8] 예상체결가전일종가대비          StartPos 31, Length 8
    char    jnilydrate          [   6];    char    _jnilydrate          ;    // [float ,  6.2] 예상체결가전일종가등락율        StartPos 40, Length 6
    char    yofferho0           [   8];    char    _yofferho0           ;    // [long  ,    8] 예상매도호가                    StartPos 47, Length 8
    char    ybidho0             [   8];    char    _ybidho0             ;    // [long  ,    8] 예상매수호가                    StartPos 56, Length 8
    char    yofferrem0          [  12];    char    _yofferrem0          ;    // [long  ,   12] 예상매도호가수량                StartPos 65, Length 12
    char    ybidrem0            [  12];    char    _ybidrem0            ;    // [long  ,   12] 예상매수호가수량                StartPos 78, Length 12
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 91, Length 6
} YS3OutBlock;

//------------------------------------------------------------------------------
// 코스피 ETF종목 실시간 NAV (I5_)
//------------------------------------------------------------------------------
typedef struct {
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 0, Length 6
} I5_InBlock;

typedef struct _I5__OutBlock
{
    char    time                [   8];    char    _time                ;    // [string,    8] 시간                            StartPos 0, Length 8
    char    price               [   8];    char    _price               ;    // [long  ,    8] 현재가                          StartPos 9, Length 8
    char    sign                [   1];    char    _sign                ;    // [string,    1] 전일대비구분                    StartPos 18, Length 1
    char    change              [   8];    char    _change              ;    // [long  ,    8] 전일대비                        StartPos 20, Length 8
    char    volume              [  12];    char    _volume              ;    // [float ,   12] 누적거래량                      StartPos 29, Length 12
    char    navdiff             [   9];    char    _navdiff             ;    // [float ,  9.2] NAV대비                         StartPos 42, Length 9
    char    nav                 [   9];    char    _nav                 ;    // [float ,  9.2] NAV                             StartPos 52, Length 9
    char    navchange           [   9];    char    _navchange           ;    // [float ,  9.2] 전일대비                        StartPos 62, Length 9
    char    crate               [   9];    char    _crate               ;    // [float ,  9.2] 추적오차                        StartPos 72, Length 9
    char    grate               [   9];    char    _grate               ;    // [float ,  9.2] 괴리                            StartPos 82, Length 9
    char    jisu                [   8];    char    _jisu                ;    // [float ,  8.2] 지수                            StartPos 92, Length 8
    char    jichange            [   8];    char    _jichange            ;    // [float ,  8.2] 전일대비                        StartPos 101, Length 8
    char    jirate              [   8];    char    _jirate              ;    // [float ,  8.2] 전일대비율                      StartPos 110, Length 8
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 119, Length 6
} I5_OutBlock;

//------------------------------------------------------------------------------
// 주식 VI발동해제 (VI_)
//------------------------------------------------------------------------------
typedef struct {
    char    shcode[6];      char    _shcode;    //[string,    6] 단축코드
} VI_InBlock;

typedef struct {
    char    vi_gubun            [   1];    char    _vi_gubun            ;    // [string,    1] 구분(0:해제 1:정적발동 2:동적발 StartPos 0, Length 1
    char    svi_recprice        [   8];    char    _svi_recprice        ;    // [long  ,    8] 정적VI발동기준가격              StartPos 2, Length 8
    char    dvi_recprice        [   8];    char    _dvi_recprice        ;    // [long  ,    8] 동적VI발동기준가격              StartPos 11, Length 8
    char    vi_trgprice         [   8];    char    _vi_trgprice         ;    // [long  ,    8] VI발동가격                      StartPos 20, Length 8
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드(KEY)                   StartPos 29, Length 6
    char    ref_shcode          [   6];    char    _ref_shcode          ;    // [string,    6] 참조코드                        StartPos 36, Length 6
    char    time                [   6];    char    _time                ;    // [string,    6] 시간                            StartPos 43, Length 6
} VI_OutBlock;

//------------------------------------------------------------------------------
// 시간외 단일가 VI발동해제 (DVI)
//------------------------------------------------------------------------------
typedef struct {
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드(KEY)                   StartPos 0, Length 6
} DVIInBlock;

typedef struct {
    char    vi_gubun            [   1];    char    _vi_gubun            ;    // [string,    1] 구분(0:해제 1:정적발동 2:동적발 StartPos 0, Length 1
    char    svi_recprice        [   8];    char    _svi_recprice        ;    // [long  ,    8] 정적VI발동기준가격              StartPos 2, Length 8
    char    dvi_recprice        [   8];    char    _dvi_recprice        ;    // [long  ,    8] 동적VI발동기준가격              StartPos 11, Length 8
    char    vi_trgprice         [   8];    char    _vi_trgprice         ;    // [long  ,    8] VI발동가격                      StartPos 20, Length 8
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드(KEY)                   StartPos 29, Length 6
    char    ref_shcode          [   6];    char    _ref_shcode          ;    // [string,    6] 참조코드(미사용)                StartPos 36, Length 6
    char    time                [   6];    char    _time                ;    // [string,    6] 시간                            StartPos 43, Length 6
} DVIOutBlock;

//------------------------------------------------------------------------------
// 장 운영 정보 (JIF)
//------------------------------------------------------------------------------
typedef struct {
    char    jangubun[1];    //[string,    1] 장구분   StartPos 0, Length 1
} JIFInBlock;

typedef struct {
    char    jangubun[1];    //[string,    1] 장구분   StartPos 0, Length 1
    char    jstatus[2];    //[string,    2] 장상태   StartPos 1, Length 2
} JIFOutBlock;

//------------------------------------------------------------------------------
// 업종별 투자자별 매매현황 (BM_)
//------------------------------------------------------------------------------
typedef struct {
    char    upCode              [   3];    char    _upcode              ;    // [string,    3] 업종코드                        StartPos 0, Length 3
} BM_InBlock;

typedef struct {
    char    tjjCode             [   4];    char    _tjjcode             ;    // [string,    4] 투자자코드                      StartPos 0, Length 4
    char    tjjTime             [   8];    char    _tjjtime             ;    // [string,    8] 수신시간                        StartPos 5, Length 8
    char    msVolume            [   8];    char    _msvolume            ;    // [long  ,    8] 매수 거래량                     StartPos 14, Length 8
    char    mdVolume            [   8];    char    _mdvolume            ;    // [long  ,    8] 매도 거래량                     StartPos 23, Length 8
    char    msVol               [   8];    char    _msvol               ;    // [long  ,    8] 거래량 순매수                   StartPos 32, Length 8
    char    pMsVol             [   8];    char    _p_msvol             ;    // [long  ,    8] 거래량 순매수 직전대비          StartPos 41, Length 8
    char    msValue             [   6];    char    _msvalue             ;    // [long  ,    6] 매수 거래대금                   StartPos 50, Length 6
    char    mdValue             [   6];    char    _mdvalue             ;    // [long  ,    6] 매도 거래대금                   StartPos 57, Length 6
    char    msVal               [   6];    char    _msval               ;    // [long  ,    6] 거래대금 순매수                 StartPos 64, Length 6
    char    pMsVal             [   6];    char    _p_msval             ;    // [long  ,    6] 거래대금 순매수 직전대비        StartPos 71, Length 6
    char    upCode              [   3];    char    _upcode              ;    // [string,    3] 업종코드                        StartPos 78, Length 3
} BM_OutBlock;

// 앞서 '#pragma pack(push, 1)'로 1바이트 단위로 설정한
// 구조체 메모리 저장방식을 원래대로 되돌림.
#pragma pack(pop)

