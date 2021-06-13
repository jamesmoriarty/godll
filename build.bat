REM config
set GOARCH=386
set GOOS=windows
set CGO_ENABLED=1
set PATH=%PATH%;C:\MinGW\bin

REM build example.dll^
go build -buildmode=c-archive -o example.a example.go
gcc -shared -o example.dll example.def example.a -Wl,--subsystem,windows

REM build example.exe^
go build -v -o example.exe main.go

REM run example.exe
.\example.exe