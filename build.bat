@echo off

IF NOT DEFINED GOPATH (
    SET GOPATH = %USERPROFILE%\Go
)

if EXIST *.exe (
    del *.exe
)

call %GOPATH%\src\github.com\ghts\dep\batch_scripts\32.bat
cd %GOPATH%\src\github.com\ghts\xing_C32

go build
