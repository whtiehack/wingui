windres -i ui/ui.rc -O coff -o ui.syso
set GOPROXY=https://goproxy.cn
go generate
go build  -buildmode=c-shared -o test.dll
exit