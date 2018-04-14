@echo off

IF NOT DEFINED GOPATH (
    SET GOPATH = %USERPROFILE%\Go
)

if EXIST *.exe (
    del *.exe
)

call %GOPATH%\src\github.com\ghts\ghts_dependency\batch_scripts\32.bat
cd %GOPATH%\src\github.com\ghts\xing_C32

go run xing_C32.go