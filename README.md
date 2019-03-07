# xing_C32

이베스트투자증권의 Xing API 호출을 대행하는 패키지
  - Xing API 호출 결과물을 네트워크를 통해서 전달.
  - 64비트에서 32비트 DLL 호출할 수 없는 문제점을 우회 해결.
  - Xing API의 DLL을 호출한 결과물을 그대로 전달하며, 추가적인 해석 및 데이터 변환은 xing 패키지에서 진행됨.

설치 준비물
  - Go언어 
    : https://golang.org/dl/
  - Rtools 패키지 (32비트 및 64비트 C언어 컴파일러를 1번에 간편하게 설치)  
    : https://cran.r-project.org/bin/windows/Rtools/index.html
  - Git 소스코드 관리 시스템
    : https://git-scm.com/download/win 

설치법
    go get github.com/ghts/xing_C32
    
사용법
  - 간접 실행 : 대개의 경우 xing패키지를 사용하면 자동으로 실행됨.
  - 직접 실행 : %GOPATH%\src\github.com\ghts\xing_C32\xing_C32.bat 
 
참고 링크.
  - 이베스트 투자증권 : https://www.ebestsec.co.kr
  - xing 패키지 : https://github.com/ghts/xing
  - nanomsg : https://nanomsg.org
  - mangos : https://github.com/nanomsg/mangos
 