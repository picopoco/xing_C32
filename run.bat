@echo off

call %GOPATH%\src\github.com\ghts\dep\batch_scripts\32.bat
cd %GOPATH%\src\github.com\ghts\xing_C32

go run xing_C32.go