@echo off


REM *********** 
REM *  32Bit  *
REM ***********
call %GOPATH%\dep\batch_scripts\32.bat

cls
cd %PROJECT_ROOT%\xing_C32\internal
copy type_c.orig type_1.go
go tool cgo -godefs type_1.go > type_2.go
sed -e 's/uint8/byte/g' type_2.go > type_3.go
sed -e 's/int8/byte/g' type_3.go > type_c_386.go
del type_1.go
del type_2.go
del type_3.go