windres -i ui/ui.rc -O coff -o ui.syso
set GOPROXY=https://goproxy.cn
go generate
go build -ldflags="-s -w -H windowsgui"

